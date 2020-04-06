package handler

import (
	"github.com/zackijack/gohello/internal/app/commons"
	"github.com/zackijack/gohello/internal/app/service"
)

// HandlerOption option for handler, including all service
type HandlerOption struct {
	commons.Options
	*service.Services
}
