package main
import (
	"context"
	"log"
	"google.golang.org/grpc"
	pb "micromod/ProductService/gen/go/sso"
)
func main() {
	conn, err := grpc.NewClient("localhost:9090")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewProductServiceClient(conn)
	r, err := c.ListProducts(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Products: %v", r.GetProducts())
}
