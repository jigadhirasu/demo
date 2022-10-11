package main

import (
	"net/http"

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

	router := gin.Default()
	router.Use(cors)
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "OK")
	})
	router.POST("/productservice.ProductService/Create", func(ctx *gin.Context) {
		gs.ServeHTTP(ctx.Writer, ctx.Request)
	})

	go router.RunTLS(":8443", "server.crt", "server.key")
	router.Run(":8080")
}
