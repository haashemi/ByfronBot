package commands

import (
	"fmt"
	"time"

	"github.com/haashemi/tgo"
	"github.com/haashemi/tgo/routers/message"
)

const ArzText = `💸| Exchange rates:

<code>🇺🇸| USD:</code> %s T
<code>🇪🇺| EUR:</code> %s T
<code>🇬🇧| GBP:</code> %s T
<code>🇹🇷| TRY:</code> %s T
<code>🇷🇺| RUB:</code> %s T

<code>🪙| Emami     :</code> %s T
<code>🪙| 1/1 Azadi :</code> %s T
<code>🪙| 1/2 Azadi :</code> %s T
<code>🪙| 1/4 Azadi :</code> %s T
<code>🪙| Gerami    :</code> %s T

⏳| Last modification happened at %v

🔥| Legal Advisers team.`

func (c *Commands) Arz(ctx *message.Context) {
	data, err := c.bonbastClient.GetData()
	if err != nil {
		handleMessageError(err, ctx)
		return
	}

	lastModified, err := time.ParseInLocation("January 02, 2006 15:04", data.LastModified, time.UTC)
	if err != nil {
		handleMessageError(err, ctx)
		return
	}
	lastModifiedStr := lastModified.Local().Format(time.Stamp)

	ctx.Reply(&tgo.SendMessage{
		Text: fmt.Sprintf(ArzText, data.Usd1, data.Eur1, data.Gbp1, data.Try1, data.Rub1, data.Emami1, data.Azadi1, data.Azadi120, data.Azadi14, data.Azadi1G, lastModifiedStr),

		ReplyMarkup: tgo.InlineKeyboardMarkup{InlineKeyboard: [][]*tgo.InlineKeyboardButton{
			{&tgo.InlineKeyboardButton{Text: "Source", Url: "https://bonbast.com/"}},
		}},
	})
}
