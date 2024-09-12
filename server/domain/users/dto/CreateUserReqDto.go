package dto

type CreateUserReqDto struct {
	UserId    int    `json:"userId"`
	UserEmail string `json:"userEmail"`
	UserName  string `json:"userName"`
}
