package entities

type Product struct {
	Id int `json:"id"`
	Category string `json:"category"`
	ProductName string `json:"product_name"`
	ProductPrice float64 `json:"product_price"`
	Rating int `json:"rating"`
	Description string `json:"description"`
	Image string `json:"image"`
}
