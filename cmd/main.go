package main

import (
	"github.com/joho/godotenv"
	"gopkg.in/telebot.v3"
	"log"
	"os"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	pref := telebot.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/start", func(c telebot.Context) error {
		time.Sleep(time.Second * 20)
		return c.Send("Hello!")
	})

	b.Handle(telebot.OnText, func(c telebot.Context) error {
		return c.Send(c.Text())
	})

	log.Println("Starting bot...")
	b.Start()
}
