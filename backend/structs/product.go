package structs

// Struct ini digunakan untuk menampilkan data Product sebagai response API
type ProductResponse struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Total       int    `json:"total"`
	CategoryId  uint   `json:"category_id"`
	SupplierId  uint   `json:"supplier_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// Struct ini digunakan untuk menerima data saat proses create Product
type ProductCreateRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Total       int    `json:"total"`
	CategoryId  uint   `json:"category_id"`
	SupplierId  uint   `json:"supplier_id"`
}

// Struct ini digunakan untuk menerima data saat proses update Product
type ProductUpdateRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Total       int    `json:"total"`
	CategoryId  uint   `json:"category_id"`
	SupplierId  uint   `json:"supplier_id"`
}
