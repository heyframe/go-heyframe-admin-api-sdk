package main

import (
	"context"
	"fmt"
	"log"

	sdk "github.com/heyframe/go-heyframe-admin-api-sdk"
)

func main() {
	ctx := context.Background()

	creds := sdk.NewPasswordCredentials("admin", "heyframe", []string{})
	client, err := sdk.NewApiClient(ctx, "http://localhost:8083", creds, nil)

	if err != nil {
		log.Fatalln(err)
	}

	apiContext := sdk.NewApiContext(ctx)
	criteria := sdk.Criteria{}
	criteria.Filter = []sdk.CriteriaFilter{{Type: "equals", Field: "parentId", Value: nil}}

	collection, _, _ := client.Repository.Product.Search(apiContext, criteria)

	for _, product := range collection.Data {
		fmt.Println(product.Name)
	}

	// Get current heyframe version
	fmt.Println(client.Info.Info(apiContext))
}
