package main

import (
	"context"
	"flag"
	"net/http"

	pb "github.com/amar-jay/first_twirp/pkg/proto"
	"github.com/amar-jay/first_twirp/pkg/service"
	"github.com/twitchtv/twirp"
)

func serverHooks() *twirp.ServerHooks {
	hooks := &twirp.ServerHooks{
		RequestReceived: func(ctx context.Context) (context.Context, error) {
			println("Request received")
			return ctx, nil
		},
		RequestRouted: func(ctx context.Context) (context.Context, error) {
			println("Request routed")
			return ctx, nil
		},
		ResponsePrepared: func(ctx context.Context) context.Context {
			println("Response prepared")
			return ctx
		},
		ResponseSent: func(ctx context.Context) {
			println("Response sent")
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
	port := flag.String("port", "8080", "port to listen on")
	s := &service.BotService{}
	// openai := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	handler := pb.NewBotServiceServer(s, serverHooks()) // bind interface and implementation

	println("Listening on port: ", *port)
	http.ListenAndServe(":"+*port, handler)
}
