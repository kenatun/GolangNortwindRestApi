package product

type Product struct {
	Id          int     `json:"id"`
	ProductCode string  `json:"productCode"`
	ProductName string  `json:"productName"`
	Description string  `json:"description"`
	StandarCost float64 `json:"standarCost"`
	ListPrice   float64 `json:"listPrice"`
	Category    string  `json:"category"`
}
