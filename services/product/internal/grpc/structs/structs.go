package structs

type CategoryRequest struct{}

type ProductRequest struct {
	ProductID string `validate:"required"`
}
