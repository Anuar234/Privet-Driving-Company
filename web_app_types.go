// webapp_types.go
package bot

type WebAppInfo struct {
	URL string `json:"url"`
}

type InlineKeyboardButton struct {
	Text   string      `json:"text"`
	WebApp *WebAppInfo `json:"web_app,omitempty"`
	// other fields...
}

func NewInlineKeyboardButtonWebApp(text string, webApp *WebAppInfo) InlineKeyboardButton {
	return InlineKeyboardButton{
		Text:   text,
		WebApp: webApp,
	}
}
