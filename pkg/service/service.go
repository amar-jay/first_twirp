package service

import (
	"context"

	pb "github.com/amar-jay/first_twirp/pkg/proto"
	"github.com/sashabaranov/go-openai"
	"github.com/twitchtv/twirp"
)

type BotService struct {
	Openai *openai.Client
	Chat   *ChatCompletion
}

func setLang(lang string) Language {
	switch lang {
	case "en":
		return English
	case "fr":
		return French
	case "tr":
		return Turkish
	case "ar":
		return Arabic
	default:
		return English
	}
}

// answer queestion using openai chatgpt
func (s *BotService) AnswerQuestion(ctx context.Context, req *pb.AnswerQuestionRequest) (res *pb.AnswerQuestionResponse, err error) {
	stream, err := s.Chat.Complete(ctx, setLang(req.Language), req.Question)
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	return &pb.AnswerQuestionResponse{
		Answer: stream.String(),
	}, nil
}

func (s *BotService) Recommend(ctx context.Context, req *pb.RecommendRequest) (res *pb.RecommendResponse, err error) {
	// TODO: implementation of the service
	stream, err := s.Chat.Complete(ctx, setLang(req.Language), req.Request)
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	return &pb.RecommendResponse{
		Recommendations: stream.List(),
	}, nil

	// return &pb.RecommendResponse{}, twirp.NewError(twirp.Unimplemented, "not implemented")
}
