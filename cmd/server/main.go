package main

import (
	"context"

	"google.golang.org/grpc"

	"github.com/smhdhsn/restaurant-gateway/internal/config"
	"github.com/smhdhsn/restaurant-gateway/internal/server"
	"github.com/smhdhsn/restaurant-gateway/internal/server/handler"
	"github.com/smhdhsn/restaurant-gateway/internal/server/resource"
	"github.com/smhdhsn/restaurant-gateway/internal/service"

	log "github.com/smhdhsn/restaurant-gateway/internal/logger"
	uspb "github.com/smhdhsn/restaurant-gateway/internal/protos/user/source"
	remoteRepository "github.com/smhdhsn/restaurant-gateway/internal/repository/remote"
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
	ur := remoteRepository.NewUserSourceRepository(&ctx, uClient)

	// instantiate services.
	us := service.NewUserSourceService(ur)

	// instantiate handlers.
	ush := handler.NewUserSourceHandler(us)

	// instantiate resources.
	u := resource.NewUserResource(ush)

	// instantiate http server.
	s := server.New(u)

	// start the http server.
	if err := s.Listen(&conf.Server); err != nil {
		log.Fatal(err)
	}
}
