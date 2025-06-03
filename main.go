package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"DB_GORM/DB"

	pb "DB_GORM/pb_file"
	s1 "DB_GORM/services"
	"DB_GORM/utils"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {
	DB.Initialize()

	//	go startGRPCServer()

	startRESTServer()
}

func startGRPCServer() {
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserserviceServer(grpcServer, &s1.User{})
	pb.RegisterRecruiterServiceServer(grpcServer, &s1.Recruiter{})
	pb.RegisterJobServiceServer(grpcServer, &s1.Job{})
	pb.RegisterApplicationServiceServer(grpcServer, &s1.Application{})

	log.Println("gRPC Server running on port 9090...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}

func startRESTServer() {
	mux := runtime.NewServeMux()

	err := pb.RegisterUserserviceHandlerServer(context.Background(), mux, &s1.User{})
	if err != nil {
		utils.ErrorLog.Fatalf("Failed to start gRPC-Gateway: %v", err)
	}

	err1 := pb.RegisterRecruiterServiceHandlerServer(context.Background(), mux, &s1.Recruiter{})

	if err1 != nil {
		utils.ErrorLog.Fatalf("Failed to start gRPC-Gateway: %v", err1)
	}

	err2 := pb.RegisterJobServiceHandlerServer(context.Background(), mux, &s1.Job{})

	if err2 != nil {
		utils.ErrorLog.Fatalf("Failed to start gRPC-Gateway: %v", err2)
	}

	err3 := pb.RegisterApplicationServiceHandlerServer(context.Background(), mux, &s1.Application{})

	if err3 != nil {
		utils.ErrorLog.Fatalf("Failed to start gRPC-Gateway: %v", err3)
	}

	log.Println("REST API Server running on port 9090...")
	if err := http.ListenAndServe(":9090", mux); err != nil {
		log.Fatalf("REST API Server stopped: %v", err)
	}
}
