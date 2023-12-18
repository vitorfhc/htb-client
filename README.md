# HTB Client

Golang Hack the Box API client

## Under Development

**This project is under heavy development.**

## Usage Example

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

htbClient := client.NewHtbClient(
  client.WithAuthToken("..."),
  client.WithCtx(ctx),
)

machine, err := htbClient.FindActiveMachineByName("analytics")
if err != nil {
  panic(err)
}
```

## Docs

Read the code. I will add documentation in the future.
