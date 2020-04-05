package service

import "context"

type IHelloService interface {
	SayHelloFromConfig(ctx context.Context) (str string)
}

type helloService struct {
	opt Option
}

func NewHelloService(opt Option) IHelloService {
	return &helloService{
		opt: opt,
	}
}

func (s *helloService) SayHelloFromConfig(ctx context.Context) (str string) {
	return s.opt.Config.GetString("app.name")
}
