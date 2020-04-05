package commons

import (
	"github.com/kitabisa/gohello/config"
	"github.com/kitabisa/perkakas/v2/log"
)

// Options common option for all object that needed
type Options struct {
	Config config.Provider
	Logger *log.Logger
}
