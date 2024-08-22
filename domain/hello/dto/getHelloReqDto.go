package dto

type GetHelloReqDto struct {
	Name string `json:"name" validate:"required"`
}
