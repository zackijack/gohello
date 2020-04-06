package commons

import (
	"github.com/kitabisa/perkakas/v2/log"
	"github.com/zackijack/gohello/config"
)

// Options common option for all object that needed
type Options struct {
	Config config.Provider
	Logger *log.Logger
}
