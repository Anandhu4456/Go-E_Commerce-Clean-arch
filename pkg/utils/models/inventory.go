package models

type InventoryResponse struct {
	ProductID int
}

type Inventory struct {
	ID uint `json:"id"`
	CategoryID int `json:"category_id"`
	Image string `json:"image"`
	ProductName string `json:"product_name"`
	Description string `json:"description"`
	Stock int `json:"stock"`
	Price float64 `json:"price"`
}

type UpdateInventory struct {
	CategoryID string `json:"category_id"`
	ProductName string `json:"product_name"`
	Description string `json:"description"`
	Stock int `json:"stock"`
	Price float64 `json:"price"`
}

type InventoryList struct {
	ID uint `json:"id"`
	CategoryID int `json:"category_id"`
	Image string `json:"image"`
	ProductName string `json:"product_name"`
	Description string `json:"description"`
	Stock int `json:"stock"`
	Price float64 `json:"price"`
}

type ImagesInfo struct {
	ID int `json:"id"`
	ImageUrl string `json:"image_url"`
}

type InvenoryDetails struct {
	Inventory Inventory
	AdditionalImages []ImagesInfo
}