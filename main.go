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

	userService := &s1.User{DB: DB.DB}
	recruiterService := &s1.Recruiter{}
	jobService := &s1.Job{}
	applicationService := &s1.Application{}

	//go startGRPCServer(userService, recruiterService, jobService, applicationService)

	startRESTServer(userService, recruiterService, jobService, applicationService)
}

func startGRPCServer(userService *s1.User, recruiterService *s1.Recruiter, jobService *s1.Job, applicationService *s1.Application) {
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserserviceServer(grpcServer, userService)
	pb.RegisterRecruiterServiceServer(grpcServer, recruiterService)
	pb.RegisterJobServiceServer(grpcServer, jobService)
	pb.RegisterApplicationServiceServer(grpcServer, applicationService)

	log.Println("gRPC Server running on port 9090...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}

func startRESTServer(userService *s1.User, recruiterService *s1.Recruiter, jobService *s1.Job, applicationService *s1.Application) {
	mux := runtime.NewServeMux()

	if err := pb.RegisterUserserviceHandlerServer(context.Background(), mux, userService); err != nil {
		utils.ErrorLog.Fatalf("Failed to start gRPC-Gateway (UserService): %v", err)
	}

	if err := pb.RegisterRecruiterServiceHandlerServer(context.Background(), mux, recruiterService); err != nil {
		utils.ErrorLog.Fatalf("Failed to start gRPC-Gateway (RecruiterService): %v", err)
	}

	if err := pb.RegisterJobServiceHandlerServer(context.Background(), mux, jobService); err != nil {
		utils.ErrorLog.Fatalf("Failed to start gRPC-Gateway (JobService): %v", err)
	}

	if err := pb.RegisterApplicationServiceHandlerServer(context.Background(), mux, applicationService); err != nil {
		utils.ErrorLog.Fatalf("Failed to start gRPC-Gateway (ApplicationService): %v", err)
	}

	log.Println("REST API Server running on port 9091...")
	if err := http.ListenAndServe(":9091", mux); err != nil {
		log.Fatalf("REST API Server stopped: %v", err)
	}
}
