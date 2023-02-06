package TgTypes

type MainMessage struct {
	Ok          bool      `json:"ok"`
	ErrorCode   *int      `json:"error_code,omitempty"`
	Description *string   `json:"description,omitempty"`
	Result      *[]Update `json:"result,omitempty"`
}

type MainSendMessage struct {
	Ok          bool     `json:"ok"`
	ErrorCode   *int     `json:"error_code,omitempty"`
	Description *string  `json:"description,omitempty"`
	Result      *Message `json:"result,omitempty"`
}

type Update struct {
	UpdateId int      `json:"update_id"`
	Message  *Message `json:"message,omitempty"`
	//EditedMessage      *EditedMessage      `json:"edited_message,omitempty"`
	//ChannelPost        *ChannelPost        `json:"channel_post,omitempty"`
	//EditedChannelPost  *EditedChannelPost  `json:"edited_channel_post,omitempty"`
	//InlineQuery        *InlineQuery        `json:"inline_query,omitempty"`
	//ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`
	//CallbackQuery      *CallbackQuery      `json:"callback_query,omitempty"`
	//ShippingQuery      *ShippingQuery      `json:"shipping_query,omitempty"`
	//PreCheckoutQuery   *PreCheckoutQuery   `json:"pre_checkout_query,omitempty"`
	//Poll               *Poll               `json:"poll,omitempty"`
	//PollAnswer         *PollAnswer         `json:"poll_answer,omitempty"`
	//MyChatMember       *MyChatMember       `json:"my_chat_member,omitempty"`
	//ChatMember         *ChatMember         `json:"chat_member,omitempty"`
	//ChatJoinRequest    *ChatJoinRequest    `json:"chat_join_request,omitempty"`
}

//type EditedMessage struct {
//}
//
//type ChannelPost struct {
//}
//
//type EditedChannelPost struct {
//}
//
//type InlineQuery struct {
//}
//
//type ChosenInlineResult struct {
//}
//
//type CallbackQuery struct {
//}
//
//type ShippingQuery struct {
//}
//
//type PreCheckoutQuery struct {
//}
//
//type Poll struct {
//}
//
//type PollAnswer struct {
//}
//
//type MyChatMember struct {
//}
//
//type ChatMember struct {
//}
//
//type ChatJoinRequest struct {
//}
