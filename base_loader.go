package main

import (
	"DebilBot/globals"
	"log"
	"os"
	"strings"
)

func LoadAnswers() {
	if len(strings.TrimSpace(globals.BotSettings.Get("answers_base").(string))) == 0 {
		log.Println("База не заполнена")
		return
	}
	data, err := os.ReadFile(globals.BotSettings.Get("answers_base").(string))
	if err != nil {
		log.Println("Файл базы отсутствует")
		return
	}

	half_base := strings.Split(string(data), "\n")
	var half_full [][]string

	for _, v := range half_base {
		half_full = append(half_full, strings.Split(v, "\\"))
	}
	globals.FullBase = half_full
	globals.HasAnswers = true
}
