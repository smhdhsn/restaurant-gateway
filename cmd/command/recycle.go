package main

import (
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/smhdhsn/restaurant-gateway/internal/config"
	"github.com/smhdhsn/restaurant-gateway/internal/service"

	log "github.com/smhdhsn/restaurant-gateway/internal/logger"
	eipb "github.com/smhdhsn/restaurant-gateway/internal/protos/edible/inventory"
	remoteRepository "github.com/smhdhsn/restaurant-gateway/internal/repository/remote"
)

// This section holds the items to be cleaned up from inventory.
var recycleFinished, recycleExpired bool

// recycleCMD is the subcommands responsible for cleaning up inventory from unusable items.
var recycleCMD = &cobra.Command{
	Use:   "recycle",
	Short: "Deletes useless items from inventory.",
	Run: func(cmd *cobra.Command, args []string) {
		// read configurations.
		conf, err := config.LoadConf()
		if err != nil {
			log.Fatal(err)
		}

		// make connection with external services.
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
		eiClient := eipb.NewEdibleInventoryServiceClient(eConn)

		// instantiate repositories.
		eiRepo := remoteRepository.NewEdibleInventoryRepository(&ctx, eiClient)

		// instantiate services.
		eiServ := service.NewEdibleInventoryService(eiRepo)

		// call the related service.
		if err := eiServ.Recycle(recycleFinished, recycleFinished); err != nil {
			log.Fatal(err)
		}
	},
}

// init function will be executed when this package is called.
func init() {
	rootCMD.AddCommand(recycleCMD)

	recycleCMD.Flags().BoolVarP(&recycleFinished, "finished", "f", false, "Recycle finished items.")
	recycleCMD.Flags().BoolVarP(&recycleExpired, "expired", "e", false, "Recycle expired items.")
}
