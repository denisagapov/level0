package model

// Order представляет информацию о заказе
type Order struct {
	OrderUID    string `json:"order_uid" validate:"required"`
	TrackNumber string `json:"track_number" validate:"required"`
	Entry       string `json:"entry" validate:"required"`
	Delivery    struct {
		Name    string `json:"name" validate:"required"`
		Phone   string `json:"phone" validate:"required"`
		Zip     string `json:"zip" validate:"required"`
		City    string `json:"city" validate:"required"`
		Address string `json:"address" validate:"required"`
		Region  string `json:"region" validate:"required"`
		Email   string `json:"email" validate:"required"`
	} `json:"delivery"`
	Payment struct {
		Transaction  string  `json:"transaction" validate:"required"`
		RequestID    string  `json:"request_id" validate:"required"`
		Currency     string  `json:"currency" validate:"required"`
		Provider     string  `json:"provider" validate:"required"`
		Amount       float64 `json:"amount" validate:"required"`
		PaymentDt    int     `json:"payment_dt" validate:"required"`
		Bank         string  `json:"bank" validate:"required"`
		DeliveryCost float64 `json:"delivery_cost" validate:"required"`
		GoodsTotal   float64 `json:"goods_total" validate:"required"`
		CustomFee    float64 `json:"custom_fee" validate:"required"`
	} `json:"payment"`
	Items []struct {
		ChrtID      int     `json:"chrt_id" validate:"required"`
		TrackNumber string  `json:"track_number" validate:"required"`
		Price       float64 `json:"price" validate:"required"`
		RID         string  `json:"rid" validate:"required"`
		Name        string  `json:"name" validate:"required"`
		Sale        int     `json:"sale" validate:"required"`
		Size        string  `json:"size" validate:"required"`
		TotalPrice  float64 `json:"total_price" validate:"required"`
		NmID        int     `json:"nm_id" validate:"required"`
		Brand       string  `json:"brand" validate:"required"`
		Status      int     `json:"status" validate:"required"`
	} `json:"items"`
	Locale            string `json:"locale" validate:"required"`
	InternalSignature string `json:"internal_signature" validate:"required"`
	CustomerID        string `json:"customer_id" validate:"required"`
	DeliveryService   string `json:"delivery_service" validate:"required"`
	Shardkey          string `json:"shardkey" validate:"required"`
	SmID              int    `json:"sm_id" validate:"required"`
	DateCreated       string `json:"date_created" validate:"required"`
	OofShard          string `json:"oof_shard" validate:"required"`
}
