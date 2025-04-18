package app

import (
	"github.com/bullockz21/pet_project21/configs"
	"github.com/sirupsen/logrus"
)

func (a *App) initConfig() {
	cfg, err := configs.Load()
	if err != nil {
		logrus.WithField("module", "config").Fatalf("Config error: %v", err)
	}
	a.cfg = cfg
}
