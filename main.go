package main

import (
	"github.com/emur-uy/backend/config"
	"github.com/emur-uy/backend/internal/infra/api"
	"github.com/emur-uy/backend/internal/infra/repositories/postgresql"
	"github.com/emur-uy/backend/internal/infra/worker"

	"github.com/gin-gonic/gin"
)

const defaultPort = "8080"

func main() {

	defer func() {
		dbInstance, _ := postgresql.Db.DB()
		_ = dbInstance.Close()
	}()

	gin.SetMode(config.Get().GinMode)

	// 3. Worker start
	go worker.Start() //start worker

	// 2. API start
	api.Start(defaultPort) //start api

}
