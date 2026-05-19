package main

type Product struct {
	Id    string  `json:"Id"`
	Price float64 `json:"Price"`
}

type OrderItem struct {
	ProductId string `json:"ProductId"`
	Quantity  int    `json:"Quantity"`
}

type OrderRequest struct {
	PaymentMethod   int         `json:"paymentMethod"`
	DiscountPercent int         `json:"discountPercent"`
	AmountTendered  float64     `json:"amountTendered"`
	Items           []OrderItem `json:"items"`
	ProviderId	  string      `json:"providerId"`
	CurrencySymbol  string      `json:"currencySymbol"`
}
