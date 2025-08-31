package http

import (
	"context"
	"flag"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	client "github.com/yeencloud/svc-identity/contract/proto"
)

var addr = "localhost:6042"

func (s *HTTPServer) RegisterUser(gctx *gin.Context) {
	flag.Parse()
	println("HTTP REGISTER")

	// Set up a connection to the server.
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func() {
		if e := conn.Close(); e != nil {
			log.Printf("failed to close connection: %s", e)
		}
	}()
	c := client.NewIdentityServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Register(ctx, &client.RegisterObject{
		Mail:     "A",
		Password: "B",
	})

	log.Printf("ID: %s", r.Id)
}
