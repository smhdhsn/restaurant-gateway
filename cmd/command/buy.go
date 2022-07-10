package main

import (
	"time"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/smhdhsn/restaurant-gateway/internal/config"
	"github.com/smhdhsn/restaurant-gateway/internal/service"
	"github.com/smhdhsn/restaurant-gateway/internal/service/dto"

	log "github.com/smhdhsn/restaurant-gateway/internal/logger"
	inventoryProto "github.com/smhdhsn/restaurant-gateway/internal/protos/edible/inventory"
	remoteRepository "github.com/smhdhsn/restaurant-gateway/internal/repository/remote"
)

var (
	// defaultAmount is the default amount of stocks to be added to inventory.
	defaultAmount = uint32(3)

	// defaultDaysTillExpires is the default remaning days till a product expires.
	defaultDaysTillExpires = uint32(3 * 30) // 3 MONTHS
)

// buyCMD is the subcommands responsible for creating food components.
var buyCMD = &cobra.Command{
	Use:   "buy",
	Short: "Stores new food components inside database if their components' stock are finished or expired.",
	Run: func(cmd *cobra.Command, args []string) {
		// read configurations.
		conf, err := config.LoadConf()
		if err != nil {
			log.Fatal(err)
		}

		// make connection with external services.
		eConn, err := grpc.Dial(
			conf.Services[config.EdibleService].Address,
			grpc.WithTransportCredentials(
				insecure.NewCredentials(),
			),
		)
		if err != nil {
			log.Fatal(err)
		}

		// instantiate gRPC clients.
		eiClient := inventoryProto.NewEdibleInventoryServiceClient(eConn)

		// instantiate repositories.
		eiRepo := remoteRepository.NewEdibleInventoryRepository(ctx, eiClient)

		// instantiate services.
		eiServ := service.NewEdibleInventoryService(eiRepo)

		// read params from CLI.
		a, err := cmd.Flags().GetUint32("amount")
		if err != nil {
			log.Fatal(err)
		}

		d, err := cmd.Flags().GetUint32("days-till-expiration")
		if err != nil {
			log.Fatal(err)
		}

		// make service's DTO with having data.
		bDTO := &dto.Buy{
			Amount:    a,
			ExpiresAt: time.Now().AddDate(0, 0, int(d)).UTC(),
		}

		// call the related service.
		if err := eiServ.Buy(bDTO); err != nil {
			log.Error(err)
		}
	},
}

// init function will be executed when this package is used.
func init() {
	rootCMD.AddCommand(buyCMD)

	buyCMD.Flags().Uint32P("amount", "a", defaultAmount, "The amount of stocks to buy of each missing item.")
	buyCMD.Flags().Uint32P("days-till-expiration", "d", defaultDaysTillExpires, "The remaning days till the product expires.")
}
