package app

import (
    "fmt"
    "github.com/bullockz21/pet_project21/configs"
    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
)

// App держит всё, что нужно для работы.
type App struct {
    cfg    *configs.Config
    engine *gin.Engine
    bot    */* телеграм-бот если нужно храним */
}

// Run стартует HTTP‑сервер.
func (a *App) Run() error {
    addr := fmt.Sprintf(":%s", a.cfg.Server.HTTPPort)
    logrus.Infof("starting HTTP server on %s", addr)
    return a.engine.Run(addr)
}
