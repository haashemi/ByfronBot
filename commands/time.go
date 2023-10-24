package commands

import (
	"fmt"
	"time"

	"github.com/haashemi/tgo"
	"github.com/haashemi/tgo/routers/message"
	ptime "github.com/yaa110/go-persian-calendar"
)

const TimeText = `⏳| Local Time
⏰l <code>%s</code>
⏰l <code>%s</code>
`

func (c *Commands) Time(ctx *message.Context) {
	ctx.Reply(&tgo.SendMessage{Text: fmt.Sprintf(
		TimeText,
		time.Now().Format("02/01/2006 15:04:05 Monday"),
		ptime.Now().TimeFormat("02/01/2006 15:04:05 Monday"),
	)})
}
