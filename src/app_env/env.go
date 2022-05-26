package app_env

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type initenv struct {}

//Abstract layer get enviroment variable keys and method to load the env files
var (
	InitEnv InitEnvInterface = &initenv{}
)

type InitEnvInterface interface {
	LoadEnvFile(providePath func() string)
	GetEnvScope() (string, error)
	GetAppId() (string, error)
	GetAppCert() (string, error)
	GetServerPort() (string, error)
}

// Reads a secret from files from the docker and other container based services
func read(fileadd string) (string, error) {
	var secret string

	//check if the file is empty
	if len(fileadd) == 0 {
		return "", nil
	}

	//read the file
	buf, err := ioutil.ReadFile(fileadd)
	if err != nil {
		return secret, err
	}

	//convert and assign
	secret = strings.TrimSpace(string(buf))
	return secret, nil
}

//get the info from local env
func (ie *initenv) LoadEnvFile(providePath func() string) {
	//get the path
	env := providePath()

	fmt.Println(env)
	fmt.Println("docker-env:::", os.Getenv("DOCKER_ENV"))

	loadErr := godotenv.Load(env)
	if loadErr != nil {
		panic(loadErr)
	}
}

func (ie *initenv) GetEnvScope() (string, error) {
	key, err := read(envScopeFile())
	if err != nil {
		return "", err
	}

	if len(key) == 0 {
		return envScope(), nil
	} else {
		return key, nil
	}
}

func (ie *initenv) GetAppId() (string, error) {
	key, err := read(appIdFile())
	if err != nil {
		return "", err
	}

	if len(key) == 0 {
		return appId(), nil
	} else {
		return key, nil
	}
}

func (ie *initenv) GetAppCert() (string, error) {
	key, err := read(appCertFile())
	if err != nil {
		return "", err
	}

	if len(key) == 0 {
		return appCert(), nil
	} else {
		return key, nil
	}
}

func (ie *initenv) GetServerPort() (string, error) {
	key, err := read(portFile())
	if err != nil {
		return "", err
	}

	if len(key) == 0 {
		return port(), nil
	} else {
		return key, nil
	}
}