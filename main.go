package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/jigadhirasu/demo/internal/product"
	"github.com/jigadhirasu/demo/internal/services/productservice"
	"github.com/jigadhirasu/demo/internal/services/xservice"
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
	xservice.RegisterXServiceServer(gs, product.NewServer())
	reflection.Register(gs)

	lis, err := net.Listen("tcp", ":7080")
	if err != nil {
		panic(err)
	}

	fmt.Println("grpc service listen at :7080")
	go gs.Serve(lis)

	conn, err := grpc.Dial(":7080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := productservice.NewProductServiceClient(conn)

	g := router(client)
	g.Use(cors)
	fmt.Println("http service listen at :8080")
	g.Run(":8080")
}

func router(client productservice.ProductServiceClient) *gin.Engine {
	router := gin.Default()

	router.POST("/productservice.ProductService/:method", func(ctx *gin.Context) {
		v, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			panic(err)
		}

		method := ctx.Param("method")

		cf := reflect.ValueOf(client)
		f := cf.MethodByName(method)
		if !f.IsValid() {
			ctx.JSON(500, "Method Not Found")
			return
		}

		p := &productservice.Product{}
		json.Unmarshal(v, p)

		result := f.Call([]reflect.Value{
			reflect.ValueOf(context.Background()),
			reflect.ValueOf(p),
		})

		if !result[1].IsNil() {
			if err != nil {
				panic(result[1].Interface().(error))
			}
		}

		ctx.JSON(http.StatusOK, result[0].Interface())
	})

	return router
}
