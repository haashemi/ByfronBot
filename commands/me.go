package commands

import (
	"fmt"

	"github.com/haashemi/tgo"
	"github.com/haashemi/tgo/routers/message"
)

const MeText = `I'm %s <code>[%d]</code>.
You are %s <code>[%d]</code>.
`

func (c *Commands) Me(ctx *message.Context) {
	info, err := ctx.Bot.GetMe()
	if err != nil {
		handleMessageError(err, ctx)
		return
	}

	ctx.Reply(&tgo.SendMessage{
		Text: fmt.Sprintf(MeText, info.FirstName, info.Id, ctx.From.FirstName, ctx.From.Id),
	})
}
