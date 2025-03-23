package dto

// FormCreateRequest represents the data needed to create a new form
type FormCreateRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

// FormUpdateRequest represents the data needed to update an existing form
type FormUpdateRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}
