package main

import (
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/smhdhsn/restaurant-gateway/internal/config"
	"github.com/smhdhsn/restaurant-gateway/internal/model"
	"github.com/smhdhsn/restaurant-gateway/internal/request"
	"github.com/smhdhsn/restaurant-gateway/internal/service"
	"github.com/smhdhsn/restaurant-gateway/internal/validator"
	"github.com/smhdhsn/restaurant-gateway/pkg/file"

	log "github.com/smhdhsn/restaurant-gateway/internal/logger"
	erpb "github.com/smhdhsn/restaurant-gateway/internal/protos/edible/recipe"
	remoteRepository "github.com/smhdhsn/restaurant-gateway/internal/repository/remote"
)

// recipeCMD is the subcommands responsible for storing sample data inside database.
var recipeCMD = &cobra.Command{
	Use:   "recipe",
	Short: "Stores sample recipes inside database.",
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
		erClient := erpb.NewEdibleRecipeServiceClient(eConn)

		// instantiate repositories.
		erRepo := remoteRepository.NewEdibleRecipeRepository(&ctx, erClient)

		// instantiate services.
		erServ := service.NewEdibleRecipeService(erRepo)

		// read JSON file's path from cli.
		j, err := cmd.Flags().GetString("json")
		if err != nil {
			log.Fatal(err)
		}

		// decode JSON file.
		data, err := readFromFile(j)
		if err != nil {
			log.Fatal(err)
		}

		// validate the entry data.
		validate := validator.New()
		if err := validate.Struct(data); err != nil {
			log.Fatal(err)
		}

		// convert JSON schema into application's DTO.
		iListDTO := make(model.MenuItemListDTO, len(data.Foods))
		for i, f := range data.Foods {
			iListDTO[i] = &model.MenuItemDTO{
				Title:               f.Title,
				IngredientTitleList: f.Ingredients,
			}
		}

		// call the related service.
		if err := erServ.Store(iListDTO); err != nil {
			log.Fatal(err)
		}
	},
}

// init function will be executed when this package is called.
func init() {
	rootCMD.AddCommand(recipeCMD)

	recipeCMD.Flags().StringP("json", "j", "", "Path to recipe JSON.")
	recipeCMD.MarkFlagRequired("json")
}

// readFromFile is responsible for reading a JSON file and converting its data into usable DTO.
func readFromFile(p string) (*request.EdibleRecipeReq, error) {
	b, err := file.ReadJsonFile(p)
	if err != nil {
		return nil, errors.Wrap(err, "error on reading the file")
	}

	var schema request.EdibleRecipeReq
	err = json.Unmarshal(b, &schema)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal json")
	}

	return &schema, nil
}
