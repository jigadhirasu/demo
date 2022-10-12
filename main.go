package main

import (
	"fmt"
	"net"

	"github.com/gin-gonic/gin"
	"github.com/jigadhirasu/demo/internal/product"
	"github.com/jigadhirasu/demo/internal/services/productservice"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func cors(c *gin.Context) {
	// c.Writer.Header().Set("Access-Control-Allow-Origin", c.GetHeader("Origin"))
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-Grpc-Web, X-User-Agent")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
	}

	c.Next()
}

func main() {

	gs := grpc.NewServer()
	productservice.RegisterProductServiceServer(gs, product.NewServer())
	reflection.Register(gs)

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	fmt.Println("grpc service listen at :8080")
	gs.Serve(lis)
}
