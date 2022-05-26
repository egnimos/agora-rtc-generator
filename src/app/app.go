package app

import (
	"fmt"
	"os"

	"github.com/egnimos/agora-rtc-generator/src/app_env"
	"github.com/gin-gonic/gin"
)

var router = *gin.Default()

func loadLocal() string {
	//get the local env
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	fmt.Println("main path:::", path)
	// splitValue := strings.Split(path, "\\")[:len(strings.Split(path, "\\"))-1]
	return fmt.Sprintf("%s\\local.env", path)
}

func init() {
	scope, err := app_env.InitEnv.GetEnvScope()
	if err != nil {
		app_env.InitEnv.LoadEnvFile(loadLocal)
	}

	if scope == "DOCKER" {
		
	} else {
		app_env.InitEnv.LoadEnvFile(loadLocal)
	}
}

func StartApp() {
	//set the trusted proxies
	if err := router.SetTrustedProxies(nil); err != nil {
		panic(err)
	}

	//define routers
	mapurl()

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}