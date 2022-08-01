package entities

type Product struct {
	Id int `json:"id"`
	CategoryID int `json:"category_id"`
	ProductName string `json:"product_name"`
	ProductPrice float64 `json:"product_price"`
	Rating int `json:"rating"`
	Description string `json:"description"`
	Image string `json:"image"`
}
