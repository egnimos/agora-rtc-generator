package app_env

import "os"

const (
	PORT = "PORT"
	ENV_SCOPE = "ENV_SCOPE"
	APP_ID    = "APP_ID"
	APP_CERT  = "APP_CERT"
)

const (
	PORT_FILE= "PORT_FILE"
	ENV_SCOPE_FILE = "ENV_SCOPE_FILE"
	APP_ID_FILE    = "APP_ID_FILE"
	APP_CERT_FILE  = "APP_CERT_FILE"
)

func envScope() string { return os.Getenv(ENV_SCOPE) }

func appId() string { return os.Getenv(APP_ID) }
func appCert() string { return os.Getenv(APP_CERT) }
func port() string {return os.Getenv(PORT)}

func envScopeFile() string { return os.Getenv(ENV_SCOPE_FILE) }

func appIdFile() string { return os.Getenv(APP_ID_FILE) }
func appCertFile() string { return os.Getenv(APP_CERT_FILE) }

func portFile() string {return os.Getenv(PORT_FILE)}
