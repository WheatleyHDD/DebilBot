package commands

import (
	"DebilBot/globals"
	"log"

	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/object"
)

func TestCommand(messageData object.MessagesMessage, args []string) {
	_, err := globals.VK.MessagesSend(api.Params{
		"peer_id":   messageData.PeerID,
		"message":   "Все зашибись!",
		"random_id": 0,
		"reply_to":  globals.CanReply(messageData.ID),
	})
	if err != nil {
		log.Println(err)
	}
}
