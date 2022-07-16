package main

import (
	"github.com/spf13/cobra"
)

func NewHelpCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "help",
		Short: "5e deploys microservices for managing D&D 5th Edition data",
		Long: `
		Find mor information at https://github.com/mitchelldyer01/5e`,
	}
}
