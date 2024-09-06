package dto

type CreateCategoryReqDto struct {
	UserEmail     string `json:"userEmail"`
	Color         string `json:"color"`
	CategoryTitle string `json:"categoryTitle"`
}
