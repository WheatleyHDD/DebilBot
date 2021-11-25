package main

import (
	"log"

	"DebilBot/globals"

	"github.com/pelletier/go-toml"
)

var (
	Appeals []interface{}
)

func LoadConfig() {

	config, err := toml.LoadFile("conf.toml")
	if err != nil {
		log.Panic(err)
	}

	globals.AccessToken = config.Get("account.access_token").(string)
	globals.BotSettings = config.Get("bot_settings").(*toml.Tree)
	Appeals = config.Get("bot_settings.appeal").([]interface{})
}
