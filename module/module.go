package module

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/t4rest/t4rest-go/logger"
)

// Module base module interface
type Module interface {
	Start() error
	Stop() error
	Title() string
}

const moduleKey = "module"

// Run runs each of the modules in a separate goroutine.
func Run(log *logger.Logger, modules ...Module) {
	if len(modules) > 0 {
		for _, m := range modules {
			log.With(moduleKey, m.Title()).Info("Starting")
			go func(m Module) {
				err := m.Start()
				if err != nil {
					log.Fatal(err)
				}
			}(m)
		}

		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c

		for _, m := range modules {
			log.With(moduleKey, m.Title()).Info("Stopping")

			go func(m Module) {
				err := m.Stop()
				if err != nil {
					log.With(moduleKey, m.Title(), "error", err).Error("module.Stop")
				}
			}(m)
		}
		log.Info("Stopped all modules")
	}
}
