package main

import (
	"fmt"
	"testing"

	"github.com/h4yfans/go-building-microservice/sdk/client"
	"github.com/h4yfans/go-building-microservice/sdk/client/products"
)

func TestOutClient(t *testing.T) {
	cfg := client.DefaultTransportConfig().WithHost("localhost:9090")
	c := client.NewHTTPClientWithConfig(nil, cfg)

	params := products.NewListProductsParams()
	prod, err := c.Products.ListProducts(params)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v", prod.GetPayload()[0])
}
