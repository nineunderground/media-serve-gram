/**
 * @Author: Iñaki Rodriguez <inaki>
 * @Date:   18-Mar-2019
 * @Project: Vylm.io
 * @Filename: mainUtils.go
 * @Last modified by:   inaki
 * @Last modified time: 11-Aug-2019
 * @Copyright: Iñaki Rodriguez
 */
package telegram

import (
  "fmt"
  "github.com/yanzay/tbot"
  "log"
  "time"
  "os"
)

// Generic
func RunBot() {
  fmt.Println("************************************************")
  fmt.Println("*             STARTING TELEGRAM BOT            *")
  fmt.Println("************************************************")

  bot := tbot.New(os.Getenv("TELEGRAM_TOKEN"))
	c := bot.Client()
	bot.HandleMessage(".*yo.*", func(m *tbot.Message) {
		c.SendChatAction(m.Chat.ID, tbot.ActionTyping)
		time.Sleep(1 * time.Second)
		c.SendMessage(m.Chat.ID, "hello!")
	})
	err := bot.Start()
	if err != nil {
		log.Fatal(err)
	}

}