package model

type AddProductsRequest struct {
	ItemPID      string `json:"item_pid"`
	ItemName     string `json:"item_name"`
	ItemCategory string `json:"item_category"`
	ItemQuantity string `json:"item_quantity"`
	ItemPrice    string `json:"item_price"`
}

type AddProductResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	ItemPID string `json:"item_pid"`
}

type GetAllProductsResponse struct {
	ItemPID      string `json:"item_pid"`
	ItemName     string `json:"item_name"`
	ItemCategory string `json:"item_category"`
	ItemQuantity string `json:"item_quantity"`
	ItemPrice    string `json:"item_price"`
}

type UpdateProductRequest struct {
	ItemPID      string `json:"item_pid"`
	ItemQuantity string `json:"item_quantity"`
	ItemPrice    string `json:"item_price"`
}
type UpdateProductResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	ItemPID string `json:"item_pid"`
}
