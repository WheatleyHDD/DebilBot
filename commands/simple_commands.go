package commands

import (
	"DebilBot/globals"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/object"
)

func Rate(messageData object.MessagesMessage, args []string) {
	rating := rand.Intn(11)

	_, err := globals.VK.MessagesSend(api.Params{
		"peer_id":   messageData.PeerID,
		"message":   "Оцениваю на " + strconv.Itoa(rating) + " из 10",
		"random_id": 0,
		"reply_to":  globals.CanReply(messageData.ID),
	})
	if err != nil {
		log.Println(err)
	}
}

func Time(messageData object.MessagesMessage, args []string) {
	dt := time.Now()
	_, err := globals.VK.MessagesSend(api.Params{
		"peer_id":   messageData.PeerID,
		"message":   "Ну, у меня " + dt.Format("15:04") + ". А у тебя?",
		"random_id": 0,
		"reply_to":  globals.CanReply(messageData.ID),
	})
	if err != nil {
		log.Println(err)
	}
}
