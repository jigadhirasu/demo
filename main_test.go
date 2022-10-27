package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/jigadhirasu/demo/internal/services/productservice"
)

func TestGrpcClient(t *testing.T) {

	go main()

	<-time.After(time.Second)

	for i := 0; i < 3; i++ {
		p := productservice.Product{
			UUID: fmt.Sprintf("PPPPPPPP-%d", time.Now().UnixNano()),
			Name: "NNNNNNNN",
		}

		v, _ := json.Marshal(p)
		resp, err := http.Post("http://localhost:8080/productservice.ProductService/Create", "application/json", bytes.NewBuffer(v))
		if err == nil {
			g, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(g))
		}
	}

	p := productservice.Product{}
	v, _ := json.Marshal(p)
	resp2, err := http.Post("http://localhost:8080/productservice.ProductService/List", "application/json", bytes.NewBuffer(v))
	if err == nil {
		g, err := ioutil.ReadAll(resp2.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(g))
	}

	<-time.After(time.Hour)

}
