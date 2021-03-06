package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mitchelldyer01/5e/pkg/controllers"
	"github.com/mitchelldyer01/5e/pkg/db"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "5e",
		Short: "A suite of microservices for managing 5e data",
		Long: `
		5e: A suite of microservices for managing 5e data

		5e is a set of REST APIs handling data for D&D 5th Edition.
		`,
		RunE: RootCmd,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				logrus.Printf("No argument(s) found -- starting up in monolithic mode")
			}
			return cobra.OnlyValidArgs(cmd, args)
		},
		ValidArgs: []string{"character", "spell", "player", "action", "feature"},
	}
}

func RootCmd(cmd *cobra.Command, args []string) error {
	repo := db.New()

	r := mux.NewRouter()

	if len(args) < 1 {
		logrus.Println("Initializing all controllers...")
		controllers.StartCharacterController(repo.DB, r)
		controllers.StartSpellController(repo.DB, r)
		controllers.StartLearnedSpellController(repo.DB, r)
		controllers.StartPlayerController(repo.DB, r)
		controllers.StartActionController(repo.DB, r)
		controllers.StartLearnedActionController(repo.DB, r)
		controllers.StartFeatureController(repo.DB, r)
		controllers.StartLearnedFeatureController(repo.DB, r)
	}

	for _, arg := range args {
		logrus.Printf("Initializing %s controller...", arg)
		switch arg {
		case "character":
			controllers.StartCharacterController(repo.DB, r)
		case "spell":
			controllers.StartSpellController(repo.DB, r)
			controllers.StartLearnedSpellController(repo.DB, r)
		case "player":
			controllers.StartPlayerController(repo.DB, r)
		case "action":
			controllers.StartActionController(repo.DB, r)
			controllers.StartLearnedActionController(repo.DB, r)
		case "feature":
			controllers.StartFeatureController(repo.DB, r)
			controllers.StartLearnedFeatureController(repo.DB, r)
		}
	}

	http.Handle("/", r)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return err
	}
	return nil
}
