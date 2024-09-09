package dto

type CreateCategoryReqDto struct {
	UserId    int    `json:"userId"`
	UserEmail string `json:"userEmail"`
	UserName  string `json:"userName"`
}
