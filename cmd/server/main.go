package main

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/smhdhsn/restaurant-gateway/internal/config"
	"github.com/smhdhsn/restaurant-gateway/internal/server"
	"github.com/smhdhsn/restaurant-gateway/internal/server/handler"
	"github.com/smhdhsn/restaurant-gateway/internal/server/resource"
	"github.com/smhdhsn/restaurant-gateway/internal/service"

	log "github.com/smhdhsn/restaurant-gateway/internal/logger"
	empb "github.com/smhdhsn/restaurant-gateway/internal/protos/edible/menu"
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
	uConn, err := grpc.Dial(
		conf.Services["user"].Address,
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	eConn, err := grpc.Dial(
		conf.Services["edible"].Address,
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	// instantiate gRPC clients.
	emClient := empb.NewEdibleMenuServiceClient(eConn)
	usClient := uspb.NewUserSourceServiceClient(uConn)

	// instantiate repositories.
	emRepo := remoteRepository.NewEdibleMenuRepository(&ctx, emClient)
	usRepo := remoteRepository.NewUserSourceRepository(&ctx, usClient)

	// instantiate services.
	emServ := service.NewEdibleMenuService(emRepo)
	usServ := service.NewUserSourceService(usRepo)

	// instantiate handlers.
	emHand := handler.NewEdibleMenuHandler(emServ)
	usHand := handler.NewUserSourceHandler(usServ)

	// instantiate resources.
	eRes := resource.NewEdibleResource(emHand)
	uRes := resource.NewUserResource(usHand)

	// instantiate http server.
	s := server.New(eRes, uRes)

	// start the http server.
	if err := s.Listen(&conf.Server); err != nil {
		log.Fatal(err)
	}
}
