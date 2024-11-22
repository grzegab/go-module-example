package main

import (
	pb "github.com/grzegab/GO_Module_Example/internal/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type SchoolServer struct {
	pb.SchoolServiceServer
}

func (app *Application) Listen() error {
	listener, err := net.Listen("tcp", app.GrpcAddr)
	if err != nil {
		log.Fatalf("[School service] failed to listen: %v", err)
		return err
	}
	server := grpc.NewServer()
	pb.RegisterSchoolServiceServer(server, &SchoolServer{})

	if err = server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
		return err
	}

	return nil
}
