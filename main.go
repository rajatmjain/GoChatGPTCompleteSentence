package main

import (
	"context"
	"fmt"
	"log"
	"os"

	gpt3 "github.com/PullRequestInc/go-gpt3"
	"github.com/joho/godotenv"
)

func main(){
	err := godotenv.Load()
  	if err != nil {
    	log.Fatal("Error loading .env file")
  	}
	apiKey := os.Getenv("API_KEY")

	if apiKey==""{
		log.Fatal("Empty API_KEY")
	}
	ctx := context.Background()
	client := gpt3.NewClient(apiKey)

	resp,err := client.Completion(ctx,gpt3.CompletionRequest{
		Prompt: []string{"My name is"},
		MaxTokens: gpt3.IntPtr(30),
		Stop:      []string{"."},
		Echo:      true,
	})

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(resp.Choices[0].Text)	
}