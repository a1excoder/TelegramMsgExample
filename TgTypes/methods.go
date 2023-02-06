package TgTypes

import "encoding/json"

func (m *MainMessage) MainMessageUnmarshal(mBytes []byte) error {
	return json.Unmarshal(mBytes, m)
}

func (m *MainSendMessage) MainSendMessageUnmarshal(mBytes []byte) error {
	return json.Unmarshal(mBytes, m)
}
