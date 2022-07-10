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
	menuProto "github.com/smhdhsn/restaurant-gateway/internal/protos/edible/menu"
	submissionProto "github.com/smhdhsn/restaurant-gateway/internal/protos/order/submission"
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

	eConn, err := grpc.Dial(
		conf.Services[config.EdibleService].Address,
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	oConn, err := grpc.Dial(
		conf.Services[config.OrderService].Address,
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	// instantiate gRPC clients.
	osClient := submissionProto.NewOrderSubmissionServiceClient(oConn)
	emClient := menuProto.NewEdibleMenuServiceClient(eConn)

	// instantiate repositories.
	osRepo := remoteRepository.NewOrderSubmissionRepository(ctx, osClient)
	emRepo := remoteRepository.NewEdibleMenuRepository(ctx, emClient)

	// instantiate services.
	osServ := service.NewOrderSubmissionService(osRepo)
	emServ := service.NewEdibleMenuService(emRepo)

	// instantiate handlers.
	osHand := handler.NewOrderSubmissionHandler(osServ)
	emHand := handler.NewEdibleMenuHandler(emServ)

	// instantiate resources.
	eRes := resource.NewEdibleResource(emHand)
	oRes := resource.NewOrderResource(osHand)

	// instantiate http server.
	s := server.New(eRes, oRes)

	// start the http server.
	if err := s.Listen(&conf.Server); err != nil {
		log.Fatal(err)
	}
}
