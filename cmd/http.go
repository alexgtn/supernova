/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/credentials/insecure"

	_ "github.com/jackc/pgx/v4/stdlib"
	"google.golang.org/grpc"

	pb "github.com/alexgtn/supernova/proto"
)

var httpPort = flag.Int("http-port", 8080, "The HTTP server port")

// mainCmd represents the main command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "gRPC HTTP gateway",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("http called")

		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		// Register gRPC server endpoint
		// Note: Make sure the gRPC server is running properly and accessible
		mux := runtime.NewServeMux()
		opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
		err := pb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf("localhost:%d", *port), opts)
		if err != nil {
			log.Fatalf("failed to register gRPC gateway: %v", err)
		}

		// serve documentation
		err = mux.HandlePath("GET", "/api/docs", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
			http.ServeFile(w, r, "docs/index.html")
		})
		if err != nil {
			log.Fatalf("failed to register docs handler: %v", err)
		}

		// Start HTTP server (and proxy calls to gRPC server endpoint)
		err = http.ListenAndServe(fmt.Sprintf(":%d", *httpPort), mux)
		if err != nil {
			log.Fatalf("failed to start gRPC gateway: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
}
