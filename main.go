package main

import (
	"fmt"
	"io"
	"log"
	"main/TgTypes"
	"net/http"
)

const (
	timeout = 50
	limit   = 1
)

func main() {
	fileValuesData, err := GetFileValues("values.json")
	if err != nil {
		log.Fatal(err)
	}

	tempOffset := 0
	mainUrl := fmt.Sprintf("https://api.telegram.org/bot%s", fileValuesData.ApiKey)
	getMessageUrl := fmt.Sprintf("%s/getUpdates?limit=%d&timeout=%d", mainUrl, limit, timeout)

	var response *http.Response
	// channels := make(chan struct{}, fileValuesData.MaxHandlers)

	url := fmt.Sprintf("%s&offset=%d", getMessageUrl, tempOffset)

	for {
		mainMessage := TgTypes.MainMessage{}

		response, err = http.Post(url, "multipart/form-data", nil)
		if err != nil {
			log.Printf("[response error] - %s\n", err.Error())
			continue
		}

		bodyData, err := io.ReadAll(response.Body)
		if err != nil {
			log.Printf("[ReadAll error] - %s\n", err.Error())
			continue
		}

		if err = mainMessage.MainMessageUnmarshal(bodyData); err != nil {
			log.Printf("[MainMessageUnmarshal error] - %s\n", err.Error())
			continue
		}

		if mainMessage.Result != nil {

			if mainMessage.Ok && len(*mainMessage.Result) > 0 {
				tempOffset = (*mainMessage.Result)[0].UpdateId + 1

				//select {
				//case channels <- struct{}{}:
				//	go ClientWorker((*mainMessage.Result)[0].Message, channels)
				//case <-time.After(30 * time.Second):
				//	//
				//default:
				//	//
				//}
				//log.Printf("channels cap: %d | len: %d\n", cap(channels), len(channels))

				// fmt.Printf("user_name: %s | text: %s\n", *(*mainMessage.Result)[0].Message.From.Username, *(*mainMessage.Result)[0].Message.Text)

				if err = ClientWorker((*mainMessage.Result)[0].Message, fileValuesData.ApiKey); err != nil {
					log.Printf("[ClientWorker error] - %s\n", err.Error())
				}

				url = fmt.Sprintf("%s&offset=%d", getMessageUrl, tempOffset)
			} else {
				if mainMessage.Ok {
					log.Println("time for a long request is up")
				} else {
					log.Printf("[ok != true || result = nil error] - %s\n", *mainMessage.Description)
				}
			}
		}

		_ = response.Body.Close()
	}
}
