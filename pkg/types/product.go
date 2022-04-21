package types

type Product struct {
	Id    string  `json:"id" bson:"_id,omitempty"`
	Name  string  `json:"name"`
	Brand string  `json:"brand"`
	Price float32 `json:"price"`
}

type ProductDTO struct {
	Id    string
	Name  string
	Brand string
	Price float32
}
