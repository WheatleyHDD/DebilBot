package commands

import (
	"DebilBot/globals"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/object"
)

func clamp(value, maxv, minv int) int {
	if value > minv {
		if value < maxv {
			return value
		}
		return maxv
	}
	return minv
}

func FindVideo(messageData object.MessagesMessage, args []string) {
	q := makeQuery(args)
	height := 5

	if q == "" {
		globals.VK.MessagesSend(api.Params{
			"peer_id":   messageData.PeerID,
			"message":   "Пустой запрос",
			"random_id": 0,
			"reply_to":  globals.CanReply(messageData.ID),
		})
		return
	}

	messageID, _ := globals.VK.MessagesSend(api.Params{
		"peer_id":   messageData.PeerID,
		"message":   "⌛ Ищу... Ожидайте",
		"random_id": 0,
		"reply_to":  globals.CanReply(messageData.ID),
	})

	videos, err := globals.VK.VideoSearch(api.Params{
		"q":     q,
		"count": 200,
		"adult": 1,
	})
	if err != nil {
		globals.VK.MessagesEdit(api.Params{
			"peer_id":               messageData.PeerID,
			"message_id":            messageID,
			"message":               "Ошиб очка: " + err.Error(),
			"keep_forward_messages": 1,
		})
		return
	}
	if videos.Count >= 1 {
		roffset := 0

		if videos.Count > 10 {
			rand.Seed(time.Now().UnixNano())
			roffset = rand.Intn(clamp(videos.Count-height, len(videos.Items)-6, 1))
		}

		attch := ""
		if videos.Count < height {
			for i := 0; i < videos.Count; i++ {
				attch = strings.Join([]string{attch, "video", strconv.Itoa(videos.Items[i].OwnerID), "_", strconv.Itoa(videos.Items[i].ID), ","}, "")
			}
		} else {
			for i := roffset; i < roffset+height; i++ {
				attch = strings.Join([]string{attch, "video", strconv.Itoa(videos.Items[i].OwnerID), "_", strconv.Itoa(videos.Items[i].ID), ","}, "")
			}
		}
		globals.VK.MessagesEdit(api.Params{
			"peer_id":               messageData.PeerID,
			"message_id":            messageID,
			"message":               "Найдены видео по вашему запросу",
			"attachment":            attch,
			"keep_forward_messages": 1,
		})
		return
	}
	globals.VK.MessagesEdit(api.Params{
		"peer_id":               messageData.PeerID,
		"message_id":            messageID,
		"message":               "По вашему запросу не было найдено видео",
		"keep_forward_messages": 1,
	})
}

func FindPhoto(messageData object.MessagesMessage, args []string) {
	q := makeQuery(args)
	height := 10

	messageID, _ := globals.VK.MessagesSend(api.Params{
		"peer_id":   messageData.PeerID,
		"message":   "⌛ Ищу... Ожидайте",
		"random_id": 0,
		"reply_to":  globals.CanReply(messageData.ID),
	})

	photos, err := globals.VK.PhotosSearch(api.Params{
		"q":     q,
		"count": 1000,
		"sort":  0,
	})

	if err != nil {
		globals.VK.MessagesEdit(api.Params{
			"peer_id":               messageData.PeerID,
			"message_id":            messageID,
			"message":               "Ошиб очка: " + err.Error(),
			"keep_forward_messages": 1,
		})
		return
	}
	if photos.Count >= 1 {
		if len(photos.Items) > 0 {
			roffset := 0

			if photos.Count > height {
				rand.Seed(time.Now().UnixNano())
				roffset = rand.Intn(clamp(photos.Count-10, len(photos.Items)-12, 1))
			}

			attch := ""

			if photos.Count < height {
				for i := 0; i < photos.Count; i++ {

					attch = strings.Join([]string{attch, "photo", strconv.Itoa(photos.Items[i].OwnerID), "_", strconv.Itoa(photos.Items[i].ID), ","}, "")
				}
			} else {
				for i := roffset; i < roffset+10; i++ {

					attch = strings.Join([]string{attch, "photo", strconv.Itoa(photos.Items[i].OwnerID), "_", strconv.Itoa(photos.Items[i].ID), ","}, "")
				}
			}
			globals.VK.MessagesEdit(api.Params{
				"peer_id":               messageData.PeerID,
				"message_id":            messageID,
				"message":               "Найдены фото по вашему запросу",
				"attachment":            attch,
				"keep_forward_messages": 1,
			})
			return
		}
		globals.VK.MessagesEdit(api.Params{
			"peer_id":               messageData.PeerID,
			"message_id":            messageID,
			"message":               "Братуха, все круто, но сейчас вк на меня накинули лимит, из-за которого я не могу отправить тебе фото. Попробуй еще раз через некоторое время.",
			"keep_forward_messages": 1,
		})
		return
	}
	globals.VK.MessagesEdit(api.Params{
		"peer_id":               messageData.PeerID,
		"message_id":            messageID,
		"message":               "По вашему запросу не было найдено фото",
		"keep_forward_messages": 1,
	})
}

func FindGIF(messageData object.MessagesMessage, args []string) {
	q := makeQuery(args)
	height := 10

	if len(q) <= 1 {
		q = "gif"
	}

	messageID, _ := globals.VK.MessagesSend(api.Params{
		"peer_id":   messageData.PeerID,
		"message":   "⌛ Ищу... Ожидайте",
		"random_id": 0,
		"reply_to":  globals.CanReply(messageData.ID),
	})

	dcs, err := globals.VK.DocsSearch(api.Params{
		"q": q,
	})
	if err != nil {
		globals.VK.MessagesEdit(api.Params{
			"peer_id":               messageData.PeerID,
			"message_id":            messageID,
			"message":               "Ошиб очка: " + err.Error(),
			"keep_forward_messages": 1,
		})
		return
	}

	rand.Seed(time.Now().UnixNano())
	roffset := rand.Intn(clamp(dcs.Count-10, dcs.Count, 1))
	docs, err := globals.VK.DocsSearch(api.Params{
		"q":      q,
		"count":  1000,
		"offset": roffset,
	})
	if err != nil {
		globals.VK.MessagesEdit(api.Params{
			"peer_id":               messageData.PeerID,
			"message_id":            messageID,
			"message":               "Ошиб очка: " + err.Error(),
			"keep_forward_messages": 1,
		})
		return
	}
	gif := ""
	cnt := 0
	for _, doc := range docs.Items {
		if cnt == height {
			break
		}
		if doc.Ext == "gif" {
			gif = strings.Join([]string{gif, "doc", strconv.Itoa(doc.OwnerID), "_", strconv.Itoa(doc.ID), ","}, "")
			cnt = cnt + 1
		}
	}
	if gif != "" {
		globals.VK.MessagesEdit(api.Params{
			"peer_id":               messageData.PeerID,
			"message_id":            messageID,
			"message":               "Найдены гиф-анимации по вашему запросу",
			"attachment":            gif,
			"keep_forward_messages": 1,
		})
		return
	}

	globals.VK.MessagesEdit(api.Params{
		"peer_id":               messageData.PeerID,
		"message_id":            messageID,
		"message":               "По вашему запросу не было найдено гиф-анимаций",
		"keep_forward_messages": 1,
	})
}

func makeQuery(args []string) string {
	q := ""

	for k, arg := range args {
		if k > 0 {
			q = q + " " + arg
		}
	}

	return strings.TrimSpace(q)
}
