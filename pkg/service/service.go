package service

import (
	"context"
	"math/rand"
	"strconv"

	pb "github.com/amar-jay/first_twirp/pkg/proto"
	"github.com/sashabaranov/go-openai"
	"github.com/twitchtv/twirp"
)

type BotService struct {
	// openai *openai.Client
}

func (s *BotService) AnswerQuestion(ctx context.Context, req *pb.AnswerQuestionRequest) (res *pb.AnswerQuestionResponse, err error) {
	// get openai client from context
	v := ctx.Value(OPENAI)
	if v == nil {
		return nil, twirp.InternalError("openai is nil 1")
	}
	_ = v.(*openai.Client)
	if err != nil {
		return nil, twirp.InternalError("openai is nil")
	}

	// TODO: implementation of the service

	return &pb.AnswerQuestionResponse{
		Answer: "The answer is 42",
	}, nil
}

func (s *BotService) Recommend(ctx context.Context, size *pb.RecommendRequest) (res *pb.RecommendResponse, err error) {
	// TODO: implementation of the service

	return &pb.RecommendResponse{
		Answer: strconv.Itoa(rand.Intn(100)),
	}, nil
}
