package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"

	pb "github.com/amar-jay/first_twirp/pkg/proto"
	"github.com/amar-jay/first_twirp/pkg/service"
	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
	"github.com/twitchtv/twirp"
)

func serverHooks() *twirp.ServerHooks {
	hooks := &twirp.ServerHooks{
		RequestReceived: func(ctx context.Context) (context.Context, error) {
			println("Request received")
			return ctx, nil
		},
		Error: func(ctx context.Context, twerr twirp.Error) context.Context {
			println("Error: ", twerr.Error())
			return ctx
		},
	}
	return hooks

}
func main() {
	// flag for port
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	port := flag.String("port", "8080", "port to listen on")
	flag.Parse()
	secret := os.Getenv("OPENAI_SECRET")
	if secret == "" {
		log.Fatal("OPENAI_SECRET not set")
	}

	cl := openai.NewClient(secret)
	comp := service.NewChatCompletion(cl, 10, service.English)
	s := &service.BotService{
		Openai: cl,
		Chat:   comp,
	}

	handler := pb.NewBotServiceServer(s, serverHooks()) // bind interface and implementation

	println("Listening on port: ", *port)
	http.ListenAndServe(":"+*port, handler)
}
