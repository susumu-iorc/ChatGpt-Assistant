package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"

	"github.com/sashabaranov/go-openai"
	Const "github.com/susumu-iorc/ChatGPT-Personal-Assistant/app/config"
	"github.com/susumu-iorc/ChatGPT-Personal-Assistant/app/database"
	"github.com/susumu-iorc/ChatGPT-Personal-Assistant/app/server"
)

func main() {
	// config.Const()
	// fmt.Println(Const.DatabaseUser)
	// fmt.Println(Const.DatabasePass)
	// fmt.Println(Const.Databasepass)
	db := database.Migrate
	// database.Seed()
	db()

	content := "C言語を使ってwgetを実装したい"

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	client := openai.NewClient(Const.OpenAIApiKey)
	msgs := []openai.ChatCompletionMessage{}
	msgs = append(msgs, openai.ChatCompletionMessage{Role: openai.ChatMessageRoleUser, Content: content})
	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{Model: openai.GPT3Dot5Turbo, Messages: msgs},
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Printf("\n\n%v\n", resp.Choices[0].Message.Content)
	msgs = append(msgs, resp.Choices[0].Message)

	if err := json.NewEncoder(os.Stdout).Encode(msgs); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	server.Router()
}
