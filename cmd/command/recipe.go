package main

import (
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/smhdhsn/restaurant-gateway/internal/config"
	"github.com/smhdhsn/restaurant-gateway/internal/service"
	"github.com/smhdhsn/restaurant-gateway/internal/service/dto"
	"github.com/smhdhsn/restaurant-gateway/internal/validator"
	"github.com/smhdhsn/restaurant-gateway/pkg/file"

	log "github.com/smhdhsn/restaurant-gateway/internal/logger"
	recipeProto "github.com/smhdhsn/restaurant-gateway/internal/protos/edible/recipe"
	remoteRepository "github.com/smhdhsn/restaurant-gateway/internal/repository/remote"
)

// EdibleRecipeReq holds the schema for edible's recipe service.
type EdibleRecipeReq struct {
	Foods []struct {
		Title       string   `json:"title" validate:"required"`
		Ingredients []string `json:"ingredients" validate:"required"`
	} `json:"foods" validate:"required"`
}

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
			conf.Services[config.EdibleService].Address,
			grpc.WithTransportCredentials(
				insecure.NewCredentials(),
			),
		)
		if err != nil {
			log.Fatal(err)
		}

		// instantiate gRPC clients.
		erClient := recipeProto.NewEdibleRecipeServiceClient(eConn)

		// instantiate repositories.
		erRepo := remoteRepository.NewEdibleRecipeRepository(ctx, erClient)

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
		v := validator.GetInstance()
		if err := v.Struct(data); err != nil {
			log.Fatal(err)
		}

		// convert JSON schema into application's DTO.
		mListDTO := multipleRecipeReqToDTO(data)

		// call the related service.
		if err := erServ.Store(mListDTO); err != nil {
			log.Error(err)
		}
	},
}

// init function will be executed when this package is used.
func init() {
	rootCMD.AddCommand(recipeCMD)

	recipeCMD.Flags().StringP("json", "j", "", "Path to recipe JSON.")
	recipeCMD.MarkFlagRequired("json")
}

// multipleRecipeReqToDTO is responsible for transforming a recipe request into a list of recipe dto struct.
func multipleRecipeReqToDTO(req *EdibleRecipeReq) []*dto.Recipe {
	rListDTO := make([]*dto.Recipe, len(req.Foods))

	for i, fReq := range req.Foods {
		iListDTO := make([]string, len(fReq.Ingredients))

		copy(iListDTO, fReq.Ingredients)

		rListDTO[i] = &dto.Recipe{
			Title:       fReq.Title,
			Ingredients: iListDTO,
		}
	}

	return rListDTO
}

// readFromFile is responsible for reading a JSON file and converting its data into usable DTO.
func readFromFile(p string) (*EdibleRecipeReq, error) {
	b, err := file.ReadJsonFile(p)
	if err != nil {
		return nil, errors.Wrap(err, "error on reading the file")
	}

	var schema EdibleRecipeReq
	err = json.Unmarshal(b, &schema)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal json")
	}

	return &schema, nil
}
