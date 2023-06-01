package main

import (
	"fmt"
	"log"
	"time"

	"github.com/aleeXpress/api"
	"github.com/gen2brain/beeep"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
	ticker := time.NewTicker(240 * time.Minute)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			word := api.RandomWordWithChatGPT()
			def := api.DefinitionLogic(word.Choices[0].Message.Content)
			fmt.Println(def)
			title := fmt.Sprintf("%s | %s", def.En, def.Es)
			desc := fmt.Sprintf("Definition : %s",def.En_def)
			err := beeep.Notify(title, desc, "")
			checkError(err)
		}
	}
	



}

func checkError(err error) {
	if err != nil {
		panic(fmt.Errorf("%w", err))
	}
}
