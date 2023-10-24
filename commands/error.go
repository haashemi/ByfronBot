package commands

import (
	"fmt"

	"github.com/haashemi/tgo"
	"github.com/haashemi/tgo/routers/message"
)

const ErrorText = `Error Occurred:
— <code>%s</code>
`

func handleMessageError(err error, ctx *message.Context) {
	ctx.Reply(&tgo.SendMessage{Text: fmt.Sprintf(ErrorText, err.Error())})
}
