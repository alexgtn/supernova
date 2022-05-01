/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"flag"
	"fmt"
	"github.com/alexgtn/supernova/infra/postgres"
	"github.com/alexgtn/supernova/infra/repository"
	pb "github.com/alexgtn/supernova/proto"
	"github.com/alexgtn/supernova/usecase"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"net"

	_ "github.com/jackc/pgx/v4/stdlib"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// mainCmd represents the main command
var mainCmd = &cobra.Command{
	Use:   "main",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("main called")

		client := postgres.OpenEnt("postgresql://default:default@postgres:5432/postgres?sslmode=disable")
		flag.Parse()

		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		var zapconfig zap.Config
		zapconfig = zap.NewProductionConfig()
		zapconfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

		logger, err := zapconfig.Build()
		if err != nil {
			log.Fatal(err)
		}

		// always log req/res payload
		alwaysLoggingDeciderServer := func(ctx context.Context, fullMethodName string, servingObject interface{}) bool { return true }
		s := grpc.NewServer(grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_zap.UnaryServerInterceptor(logger),
				grpc_zap.PayloadUnaryServerInterceptor(logger, alwaysLoggingDeciderServer),
			),
		))
		userRepo := repository.NewUser(client)
		userUsecase := usecase.NewUserService(userRepo)
		pb.RegisterUserServiceServer(s, userUsecase)

		log.Printf("server listening at %v", lis.Addr())

		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}

	},
}

func init() {
	rootCmd.AddCommand(mainCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mainCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mainCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
