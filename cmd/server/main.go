package main

import (
	"context"

	"google.golang.org/grpc"

	"github.com/smhdhsn/restaurant-gateway/internal/config"
	"github.com/smhdhsn/restaurant-gateway/internal/server"
	"github.com/smhdhsn/restaurant-gateway/internal/server/resource"

	log "github.com/smhdhsn/restaurant-gateway/internal/logger"
	uspb "github.com/smhdhsn/restaurant-gateway/internal/protos/user/source"
	uRepo "github.com/smhdhsn/restaurant-gateway/internal/repository/remote/user"
	uHand "github.com/smhdhsn/restaurant-gateway/internal/server/handler/user"
	uServ "github.com/smhdhsn/restaurant-gateway/internal/service/user"
)

// ctx holds application's context.
var ctx context.Context

// init will be called when this package is imported.
func init() {
	ctx = context.Background()
}

// main is the application's kernel.
func main() {
	// read configurations.
	conf, err := config.LoadConf()
	if err != nil {
		log.Fatal(err)
	}

	// make connection with external services.
	uConn, err := grpc.Dial(conf.Services["user"].Address, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	// instantiate gRPC clients.
	uClient := uspb.NewUserSourceServiceClient(uConn)

	// instantiate repositories.
	ur := uRepo.NewUserSourceRepository(&ctx, uClient)

	// instantiate services.
	us := uServ.NewUserSourceService(ur)

	// instantiate handlers.
	ush := uHand.NewUserSourceHandler(us)

	// instantiate resources.
	u := resource.NewUserResource(ush)

	// instantiate http server.
	s := server.New(u)

	// start the http server.
	if err := s.Listen(&conf.Server); err != nil {
		log.Fatal(err)
	}
}
