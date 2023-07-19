package main

import (
	"context"
	"net/http"

	pb "github.com/amar-jay/first_twirp/pkg/proto"
)

func main() {

	client := pb.NewBotServiceJSONClient("http://localhost:8080", &http.Client{})
	if client == nil {
		println("Client is nil")
		return
	}

	ctx := context.Background()

	res, err := client.AnswerQuestion(ctx, &pb.AnswerQuestionRequest{
		Language: "en",
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
