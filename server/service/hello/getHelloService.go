package hello

import "taskbuddy.io/taskbuddy/server/domain/hello/dto"

type (
	GetHelloService interface {
		GetHello(dto.GetHelloReqDto) string
	}

	getHelloServiceImpl struct {
	}
)

func NewGetHelloService() GetHelloService {
	i := new(getHelloServiceImpl)
	return i
}

func (s *getHelloServiceImpl) GetHello(reqDto dto.GetHelloReqDto) string {
	return "Hello " + reqDto.Name + "!"
}
