package service

import (
	"github.com/zackijack/gohello/internal/app/commons"
	"github.com/zackijack/gohello/internal/app/repository"
)

// Option anything any service object needed
type Option struct {
	commons.Options
	*repository.Repository
}

// Services all service object injected here
type Services struct {
	HealthCheck IHealthCheck
	Hello       IHelloService
}
