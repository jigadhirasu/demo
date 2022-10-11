package main

import (
	"testing"
	"time"
)

func TestGrpcClient(t *testing.T) {

	go main()

	// <-time.After(time.Second)

	// conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	// if err != nil {
	// 	panic(err)
	// }

	// client := productservice.NewProductServiceClient(conn)
	// ctx := context.Background()

	// for i := 0; i < 3; i++ {
	// 	r, err := client.Create(ctx, &productservice.Product{UUID: fmt.Sprintf("AAA%d", i+1), Name: "AAA"})
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println(r)
	// }

	// pl, err := client.List(ctx, &productservice.Product{})
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(pl.List)

	// observable, err := client.Subscribe(ctx, &productservice.Product{})
	// if err != nil {
	// 	panic(err)
	// }

	// for i := 0; i < 6; i++ {
	// 	g := &productservice.Product{}
	// 	observable.RecvMsg(g)
	// 	fmt.Println(g)
	// }

	<-time.After(time.Hour)

}
