package main
import (
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
	pb "./gen/go/sso"
)

type server struct {
	pb.UnimplementedProductServiceServer
}

func (s *server) ListProducts(ctx context.Context, in *pb.Empty) (*pb.ProductResponse, error) {
	return &pb.ProductResponse{Products: []string{"Book", "Pen"}}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProductServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
	log.Fatalf("failed to serve: %v", err)
	}
}
