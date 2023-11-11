package commands

import (
	"fmt"
	"path/filepath"

	"github.com/haashemi/tgo"
	"github.com/haashemi/tgo/routers/message"
)

const STPText = `🌹| Sticker info:
🤝| Set Name: %s
`

func (c *Commands) STP(ctx *message.Context) {
	if ctx.ReplyToMessage == nil {
		ctx.Send(&tgo.SendMessage{Text: "Please reply to the sticker you want to download."})
		return
	} else if ctx.ReplyToMessage.Sticker == nil {
		ctx.Send(&tgo.SendMessage{Text: "The message you've replied to, is not a sticker."})
		return
	}

	ctx.Bot.SendChatAction(&tgo.SendChatAction{ChatId: tgo.ID(ctx.Chat.Id), Action: "upload_document"})

	sticker := ctx.ReplyToMessage.Sticker

	file, err := ctx.Bot.GetFile(&tgo.GetFile{FileId: sticker.FileId})
	if err != nil {
		handleMessageError(err, ctx)
		return
	}

	r, err := ctx.Bot.API.Download(file.FilePath)
	if err != nil {
		handleMessageError(err, ctx)
		return
	}
	defer r.Body.Close()

	ctx.Reply(&tgo.SendDocument{
		Document:                    tgo.FileFromReader(filepath.Base(file.FilePath), r.Body),
		DisableContentTypeDetection: true,
		Caption:                     fmt.Sprintf(STPText, sticker.SetName),
	})
}
