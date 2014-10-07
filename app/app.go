package app

import (
	log "github.com/cihub/seelog"
	"github.com/nvcook42/morgoth/config"
	"github.com/nvcook42/morgoth/engine"
	_ "github.com/nvcook42/morgoth/plugins"
)

type App struct {
	config *config.Config
	Engine engine.Engine
}

func New(config *config.Config) (*App, error) {
	app := App{config: config}
	eng, err := app.config.EngineConf.GetEngine()
	if err != nil {
		return nil, err
	}
	app.Engine = eng

	return &app, nil
}

func (self *App) Run() error {
	log.Info("Setup data engine")
	log.Infof("Engine: %v", self.Engine)
	return nil
}