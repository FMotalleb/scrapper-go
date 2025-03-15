/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/fmotalleb/scrapper-go/server"
	"github.com/spf13/cobra"
)

var (
	address string
	port    uint32
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve service as an api endpoint",
	Run: func(cmd *cobra.Command, args []string) {
		server.StartServer(fmt.Sprintf("%s:%d", address, port))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().StringVarP(&address, "address", "a", "127.0.0.1", "change this value if you want to expose server (since this app does not support authentication keep it behind a reverse proxy)")
	serveCmd.Flags().Uint32VarP(&port, "port", "p", 8080, "port on which the service will be exposed (since this app does not support authentication keep it behind a reverse proxy)")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
