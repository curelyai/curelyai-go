# CurelyAI Go Client

This is a Go client for interacting with the Curely AI chat API.

## Installation

To install the package, run:

```sh
go get github.com/curelyai/curelyai-go

# usage sample code
package main

import (
    "context"
    "fmt"
    "time"

    "github.com/curelyai/curelyai-go"
)

func main() {
    client := curelyai_go.NewChatClient("your-bot-key")
    
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    response, err := client.Chat(ctx, "Hello, Curely!")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Response from Curely AI:", response)
}
