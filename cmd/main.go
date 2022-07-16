package main

import (
	"os"

	"github.com/common-nighthawk/go-figure"
	"github.com/sirupsen/logrus"
)

var logo = figure.NewColorFigure("D&D 5e", "", "green", true)

func main() {
	logo.Print()

	rootCmd := NewRootCmd()
	rootCmd.SetHelpCommand(NewHelpCmd())

	err := rootCmd.Execute()
	if err != nil {
		logrus.Error(err)
		os.Exit(-1)
	}
}
