package main

import (
	"github.com/emur-uy/backend/internal/infra/worker"

	"github.com/emur-uy/backend/internal/infra/repositories/postgresql"
)

func main() {

	defer func() {
		dbInstance, _ := postgresql.Db.DB()
		_ = dbInstance.Close()
	}()

	// 2. Workers's start
	worker.Start() //start api

}
