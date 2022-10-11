package main

import (
	"log"
	"net"

	"github.com/jigadhirasu/demo/internal/product"
	"github.com/jigadhirasu/demo/internal/services/productservice"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	gs := grpc.NewServer()
	productservice.RegisterProductServiceServer(gs, product.NewServer())
	reflection.Register(gs)

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	log.Println("grpc server listen at :8080")
	log.Fatal(gs.Serve(lis))
}
