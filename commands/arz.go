package commands

import (
	"time"

	"github.com/haashemi/tgo"
	"github.com/haashemi/tgo/routers/message"
	"golang.org/x/text/language"
	printer "golang.org/x/text/message"
)

const ArzText = `💸| Exchange rates:

<code>🇺🇸| USD:</code> <code>%.0f IRT</code>
<code>🇪🇺| EUR:</code> <code>%.0f IRT</code>
<code>🇬🇧| GBP:</code> <code>%.0f IRT</code>
<code>🇹🇷| TRY:</code> <code>%.0f IRT</code>
<code>🇷🇺| RUB:</code> <code>%.0f IRT</code>

<code>🪙| Emami     :</code> <code>%.0f IRT</code>
<code>🪙| 1/1 Azadi :</code> <code>%.0f IRT</code>
<code>🪙| 1/2 Azadi :</code> <code>%.0f IRT</code>
<code>🪙| 1/4 Azadi :</code> <code>%.0f IRT</code>
<code>🪙| Gerami    :</code> <code>%.0f IRT</code>

⏳| Last modification happened at %v
`

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
		Text: printer.NewPrinter(language.English).Sprintf(ArzText, data.Usd1, data.Eur1, data.Gbp1, data.Try1, data.Rub1, data.Emami1, data.Azadi1, data.Azadi120, data.Azadi14, data.Azadi1G, lastModifiedStr),

		ReplyMarkup: tgo.InlineKeyboardMarkup{InlineKeyboard: [][]*tgo.InlineKeyboardButton{
			{&tgo.InlineKeyboardButton{Text: "Source", Url: "https://bonbast.com/"}},
		}},
	})
}
