package main

import (
	"context"
	"grpc-app/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type AppService struct {
	proto.UnimplementedAppServiceServer
}

func NewAppService() *AppService {
	return &AppService{}
}

func (appService *AppService) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	x := req.GetX()
	y := req.GetY()
	result := x + y
	res := &proto.AddResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	appService := NewAppService()
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterAppServiceServer(grpcServer, appService)
	grpcServer.Serve(listener)
}
