package structs

// Struct ini digunakan untuk menampilkan data Category sebagai response API
type CategoryResponse struct {
	Id        uint    `json:"id"`
	Name      string  `json:"name"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

// Struct ini digunakan untuk menerima data saat proses create Category
type CategoryCreateRequest struct {
	Name     string `json:"name" binding:"required"`
}

// Struct ini digunakan untuk menerima data saat proses update Category
type CategoryUpdateRequest struct {
	Name     string `json:"name" binding:"required"`
}
