package dto

type CreateUserReqDto struct {
	UserEmail string `json:"userEmail"`
	UserName  string `json:"userName"`
}
