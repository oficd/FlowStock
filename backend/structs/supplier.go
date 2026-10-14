package structs

// Struct ini digunakan untuk menampilkan data Supplier sebagai response API
type SupplierResponse struct {
	Id        uint    `json:"id"`
	Name      string  `json:"name"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

// Struct ini digunakan untuk menerima data saat proses create Supplier
type SupplierCreateRequest struct {
	Name     string `json:"name" binding:"required"`
}

// Struct ini digunakan untuk menerima data saat proses update Supplier
type SupplierUpdateRequest struct {
	Name     string `json:"name" binding:"required"`
}
