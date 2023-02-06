package main

import (
	"fmt"
	"io"
	"main/TgTypes"
	"net/http"
)

func replyToMessage(message *TgTypes.Message, msgText, apiKey string) error {
	mainMsg := TgTypes.MainSendMessage{}

	if message.Chat == nil {
		return fmt.Errorf("[replyToMessage error message.Chat == nil]")
	}

	response, err := http.Post(
		fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%d&text=%s&reply_to_message_id=%d",
			apiKey, message.Chat.Id, msgText, message.MessageId), "multipart/form-data", nil)
	if err != nil {
		return err
	}

	bodyData, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if err = mainMsg.MainSendMessageUnmarshal(bodyData); err != nil {
		return err
	}

	if mainMsg.Ok {
		return nil
	} else {
		if mainMsg.Description != nil {
			return fmt.Errorf("[replyToMessage error] - %s", *mainMsg.Description)
		} else {
			return fmt.Errorf("[replyToMessage error without description]")
		}
	}

}

// , channels chan struct{}
func ClientWorker(message *TgTypes.Message, apiKey string) error {

	if message.Text != nil {
		if message.Entities != nil {
			if len(*message.Entities) > 0 {
				if (*message.Entities)[0].Type == "bot_command" {
					switch *message.Text {
					case "/start":
						if err := replyToMessage(message, "hi", apiKey); err != nil {
							return err
						}
					case "/about_me":
						if err := replyToMessage(message, fmt.Sprintf("data"), apiKey); err != nil {
							return err
						}
					default:
						if err := replyToMessage(message, "unknown command", apiKey); err != nil {
							return err
						}
					}
				}
			}
		} else {
			if err := replyToMessage(message, "your message is not a command", apiKey); err != nil {
				return err
			}
		}
	} else {
		return fmt.Errorf("[message.Text == nil error in ClientWorker]")
	}

	return nil
	// <-channels
}
