package service

import (
	"context"
	"fmt"
	"io"
	"strings"
	"time"

	openai "github.com/sashabaranov/go-openai"
)

const (
	BotIdentity = "Bot"
)

type Language string

const (
	English Language = "English"
	French  Language = "French"
	Turkish Language = "Turkish"
	Arabic  Language = "Arabic"
)

// A sentence in the conversation (Used for the history)
type SpeechEvent struct {
	ParticipantName string
	IsBot           bool
	Text            string
}

type JoinLeaveEvent struct {
	Leave           bool
	ParticipantName string
	Time            time.Time
}

type MeetingEvent struct {
	Speech *SpeechEvent
	Join   *JoinLeaveEvent
}

type ChatCompletion struct {
	client   *openai.Client
	messages []openai.ChatCompletionMessage
	maxQues  uint // maximum number of questions to ask before clearing the history
}

func NewChatCompletion(client *openai.Client, maxQues uint, language Language) *ChatCompletion {

	if language == "" {
		language = English
	}
	if maxQues == 0 {
		maxQues = 3
	}
	return &ChatCompletion{
		client:   client,
		messages: make([]openai.ChatCompletionMessage, maxQues),
		maxQues:  maxQues,
	}
}

func (c *ChatCompletion) Complete(ctx context.Context, language Language, prompt string) (*ChatStream, error) {

	if len(c.messages) == cap(c.messages) {
		l := len(c.messages)
		c.messages = make([]openai.ChatCompletionMessage, l)
	}
	c.messages = append(c.messages, openai.ChatCompletionMessage{
		Role: openai.ChatMessageRoleSystem,
		Content: "You are" + BotIdentity + ", a chat assistant in a meeting created by Amar Jay. " +
			"Keep your responses concise while still being friendly and personable. " +
			"If your response is a question, please append a question mark symbol to the end of it. " + // Used for auto-trigger
			fmt.Sprintf("Current language: %s Current date: %s", language, time.Now().Format("January 2, 2006 3:04pm")),
	})

	c.messages = append(c.messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: prompt,
		Name:    "user",
	})

	stream, err := c.client.CreateChatCompletionStream(ctx, openai.ChatCompletionRequest{
		Model:    openai.GPT3Dot5Turbo,
		Messages: c.messages,
		Stream:   true,
	})

	if err != nil {
		return nil, err
	}

	return &ChatStream{
		stream: stream,
	}, nil
}

// Wrapper around openai.ChatCompletionStream to return only complete sentences
type ChatStream struct {
	stream *openai.ChatCompletionStream
}

func (c *ChatStream) String() string {
	sb := strings.Builder{}
	for {
		response, err := c.stream.Recv()
		if err != nil {
			content := sb.String()
			if err == io.EOF && len(strings.TrimSpace(content)) != 0 {
				return content
			}
			panic(err)
		}

		if len(response.Choices) == 0 {
			continue
		}

		delta := response.Choices[0].Delta.Content
		sb.WriteString(delta)

		if strings.HasSuffix(strings.TrimSpace(delta), ".") {
			return sb.String()
		}
	}
}

func (c *ChatStream) List() []string {
	sb := strings.Builder{}
	for {
		response, err := c.stream.Recv()
		if err != nil {
			content := sb.String()
			if err == io.EOF && len(strings.TrimSpace(content)) != 0 {
				return []string{content}
			}
			panic(err)
		}

		if len(response.Choices) == 0 {
			continue
		}

		delta := response.Choices[0].Delta.Content
		sb.WriteString(delta)

		if strings.HasSuffix(strings.TrimSpace(delta), ".") {
			return []string{sb.String()}
		}
	}
}
func (c *ChatStream) Close() {
	c.stream.Close()
}
