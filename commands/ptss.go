package commands

import (
	"bytes"

	"github.com/LlamaNite/llamaimage"
	"github.com/haashemi/tgo"
	"github.com/haashemi/tgo/routers/message"
)

func (c *Commands) PTSS(ctx *message.Context) {
	if ctx.ReplyToMessage == nil {
		ctx.Reply(&tgo.SendMessage{Text: "Command should be replied to a photo."})
		return
	} else if ctx.ReplyToMessage.Photo == nil {
		ctx.Reply(&tgo.SendMessage{Text: "Replied message must be a photo, nothing else."})
		return
	}

	ctx.Bot.SendChatAction(&tgo.SendChatAction{ChatId: tgo.ID(ctx.Chat.Id), Action: "upload_document"})

	lastElement := len(ctx.ReplyToMessage.Photo) - 1
	fileInfo, err := ctx.Bot.GetFile(&tgo.GetFile{FileId: ctx.ReplyToMessage.Photo[lastElement].FileId})
	if err != nil {
		handleMessageError(err, ctx)
		return
	}

	resp, err := ctx.Bot.Download(fileInfo.FilePath)
	if err != nil {
		handleMessageError(err, ctx)
		return
	}
	defer resp.Body.Close()

	img, err := llamaimage.OpenImage(resp.Body)
	if err != nil {
		handleMessageError(err, ctx)
		return
	}
	img = llamaimage.Resize(img, 512, 512, llamaimage.ResizeFit)

	buf := bytes.NewBuffer(nil)
	if err = llamaimage.SaveToStream(img, buf); err != nil {
		handleMessageError(err, ctx)
		return
	}

	ctx.Reply(&tgo.SendDocument{Document: tgo.FileFromReader(fileInfo.FileId+".png", buf)})
}
