package http_invoice

type price struct {
	Label  string `json:"label"`
	Amount int    `json:"amount"`
}

type invoice struct {
	ChatID         int64   `json:"chat_id"`
	Title          string  `json:"title"`
	Description    string  `json:"description"`
	Payload        string  `json:"payload"`
	ProviderToken  string  `json:"provider_token"`
	Currency       string  `json:"currency"`
	Prices         []price `json:"prices"`
	StartParameter string  `json:"start_parameter"`
	ProviderData   string  `json:"provider_data"`
}
