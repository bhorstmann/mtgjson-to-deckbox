package main

import (
	"log"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Hugo Static Site Generator v0.9 -- HEAD")
	},
}

func main() {
	log.Println("hi")
	log.Println(versionCmd)
}
