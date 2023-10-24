package commands

import (
	"fmt"

	"github.com/haashemi/tgo"
	"github.com/haashemi/tgo/routers/message"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

const ServerText = `⚡️ <b>Server Info</b>
📈| Host Uptime ⇒ <code>%dh %dm</code>
🖥| CPU Usage ⇒ <code>%d%%</code>
⏳| Ram Usage ⇒ <code>%d%%</code>
`

func (c *Commands) Server(ctx *message.Context) {
	hostUptime, err := host.Uptime()
	if err != nil {
		handleMessageError(err, ctx)
		return
	}

	ramStatus, err := mem.VirtualMemory()
	if err != nil {
		handleMessageError(err, ctx)
		return
	}

	cpuUsage, err := cpu.Percent(0, false)
	if err != nil {
		handleMessageError(err, ctx)
		return
	}

	ctx.Reply(&tgo.SendMessage{Text: fmt.Sprintf(
		ServerText,
		(hostUptime / 60 / 60), (hostUptime / 60 % 60),
		int(cpuUsage[0]),
		int(ramStatus.UsedPercent),
	)})
}
