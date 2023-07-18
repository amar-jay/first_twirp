package main

import (
	"context"
	"net/http"
	"os"

	pb "github.com/amar-jay/first_twirp/pkg/proto"
	"github.com/amar-jay/first_twirp/pkg/service"
	"github.com/sashabaranov/go-openai"
)

func main() {

	client := pb.NewBotServiceJSONClient("http://localhost:8080", &http.Client{})
	if client == nil {
		println("Client is nil")
		return
	}

	ctx := context.Background()
	// set openai.Client in context
	openaicl := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	if openaicl == nil {
		println("openai is nil")
		return
	}

	ctx = context.WithValue(ctx, service.OPENAI, openaicl)
	res, err := client.AnswerQuestion(ctx, &pb.AnswerQuestionRequest{
		Question: "What is the answer to the Ultimate Question of Life, the Universe, and Everything?"})
	if err != nil {
		println("Error: ", err.Error())
		return
	}

	println("Answer: ", res.Answer)

	// client := haberdasher.NewHaberdasherProtobufClient("http://localhost:8080", &http.Client{})

	// hat, err := client.MakeHat(context.Background(), &haberdasher.Size{Inches: 12})
	// if err != nil {
	// 	fmt.Printf("oh no: %v", err)
	// 	os.Exit(1)
	// }
	// fmt.Printf("I have a nice new hat: %+v", hat)
}
