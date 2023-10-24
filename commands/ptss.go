package commands

import (
	"github.com/haashemi/tgo"
	"github.com/haashemi/tgo/routers/message"
)

func (c *Commands) PTSS(ctx *message.Context) {
	ctx.Send(&tgo.SendMessage{Text: "Not implemented, yet."})
}
