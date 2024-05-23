# curelyai

`curelyai` is a Go package that provides a client for interacting with the Curely AI chat API. 

## USAGE

To install the `curelyai` package, use `go get`:

```sh
go get github.com/curelyai/curelyai-go



package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "github.com/curelyai/curelyai-go"
)

func main() {
    botKey := "your-bot-key" // Replace with your actual bot key
    client := curelyai.NewChatClient(botKey, 5*time.Second)

    ctx := context.Background()
    response, err := client.Chat(ctx, "Hello, Curely AI!")
    if err != nil {
        log.Fatalf("Error: %v", err)
    }

    fmt.Println("Response from Curely AI:", response)
}


