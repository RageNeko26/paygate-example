package model

type WebResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type MidtransRequest struct {
	UserID   int    `json:"user_id"`
	Amount   int64  `json:"amount"`
	ItemID   string `json:"item_id"`
	ItemName string `json:"item_name"`
}

type MidtransResponse struct {
	Token       string `json:"token"`
	RedirectURL string `json:"redirect_url"`
}
