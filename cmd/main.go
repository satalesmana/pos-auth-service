package main

import (
	"fmt"
	"log"
	"net"

	pb "go-grpc-auth-svc/files/proto/gen"
	"go-grpc-auth-svc/helpers/jwt"
	"go-grpc-auth-svc/pkg/config"
	"go-grpc-auth-svc/pkg/repository/conndb"
	"go-grpc-auth-svc/services"

	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	connDb, _ := conndb.NewMySQLConn(c.DbConfig)

	jwt := jwt.JwtWrapper{
		SecretKey:       c.JWTSecretKey,
		Issuer:          "go-grpc-auth-svc",
		ExpirationHours: 24 * 365,
	}

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Auth Svc on", c.Port)

	s := services.Server{
		ConDb: connDb,
		Jwt:   jwt,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
