package globals

import (
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/pelletier/go-toml"
)

var (
	BotSettings *toml.Tree
	AccessToken string
	VK          *api.VK
	FullBase    [][]string
	HasAnswers  bool = false
)

func CanReply(messageID int) int {
	if BotSettings.Get("reply_message").(bool) {
		return messageID
	}
	return 0
}
