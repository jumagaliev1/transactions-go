package transport

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"transactions/internal/service"
	pb "transactions/proto"
)

var addr = ":50051"

type Server struct {
	pb.TransactionServiceServer
	service service.Service
}

func Run(service *service.Service) error {
	lis, err := net.Listen("tcp", addr)
	fmt.Println("Strating server")
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	pb.RegisterTransactionServiceServer(s, &Server{service: *service})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		return err
	}

	return nil
}
