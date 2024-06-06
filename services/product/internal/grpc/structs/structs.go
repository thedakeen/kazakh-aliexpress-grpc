package structs

type CategoryRequest struct{}

type ProductRequest struct {
	ProductID string `validate:"required"`
}

type ProductsByCategoryRequest struct {
	CategoryID string `validate:"required"`
	Limit      int64  `validate:"required,min=1"`
	Offset     int64  `validate:"required,min=1"`
	SortOrder  string
}

type CreateCategoryRequest struct {
	CategoryName string `validate:"required"`
}

type UpdateCategoryRequest struct {
	CategoryID   string `validate:"required"`
	CategoryName string `validate:"required"`
}

type DeleteCategoryRequest struct {
	CategoryID string `validate:"required"`
}
