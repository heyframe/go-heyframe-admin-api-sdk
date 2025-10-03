# SDK for the Heyframe 6 Admin API

See example folder

## Create a client

### Password

```go
ctx := context.Background()

// Create a using password grant
creds := sdk.NewPasswordCredentials("<username>", "<password>", []string{"write"})
client, err := sdk.NewApiClient(ctx, "<url>", creds, nil)
```

### Integration

```go
ctx := context.Background()

// Create a using password grant
creds := sdk.NewIntegrationCredentials("<client-id>", "<client-secret>", []string{"write"})
client, err := sdk.NewApiClient(ctx, "<url>", creds, nil)
```

## Usage of a repository

### Search

```go
apiContext := sdk.NewApiContext(ctx)
criteria := sdk.Criteria{}

collection, _, _ := client.Repository.Product.Search(apiContext, criteria)

for _, product := range collection.Data {
    fmt.Println(product.Name)
}
```

### Create/Update

```go
apiContext := sdk.NewApiContext(ctx)
client.Repository.Product.Upsert(apiContext, []sdk.Product{
    {Name: "15%"},
})
```

### Delete

```go
apiContext := sdk.NewApiContext(ctx)
client.Repository.Product.Delete(apiContext, []string{"someid"})
```