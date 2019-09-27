package main

import (
    "os"
    "log"

    api "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
    bot, err := api.NewBotAPI(os.Getenv("TG_KEY"))

    if err != nil {
        log.Panic(err)
    } // ENDIF

    bot.Debug = true

    log.Printf("Authorized on account %s\n\n\n", bot.Self.UserName)

    updateConfig := api.NewUpdate(0)
    updateConfig.Timeout = 60
    updates, err := bot.GetUpdatesChan(updateConfig)

     if err != nil {
         log.Panic(err)
     } // ENDIF

    for update := range updates {
        if update.Message == nil {
            continue
        } // ENDIF

        // replyMessage := update.Message.Text
        replyMessage := "Please use commands"

        // check if the message is a command
        // and extract that command
        if cmd := update.Message.Command(); cmd != "" {
            // default response
            replyMessage = "You want /newEmail ?"

            // if correct command is triggered
            if cmd == "newEmail" {
                replyMessage = "junkmain@hotmail.com:12341234"
            }

            log.Printf("\n\n\n\nThis is a command: %s\n\n\n\n", cmd)
        }


        // func NewMessage(chatID int64, text string)
        // chatID: where to send it
        // text: message text
        msg := api.NewMessage(update.Message.Chat.ID, replyMessage)
        // msg.ReplyToMessageID = update.Message.MessageID



        // send meassge
        _, err := bot.Send(msg)

        if err != nil {
            log.Panic(err)
        } // ENDIF
    } // ENDFOR
} // ENDMAIN
