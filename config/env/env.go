package env

import (
	"fmt"
	"os"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/mises-id/mises-websitesvc/lib/storage/view"
)

var Envs *Env

type Env struct {
	Port         int    `env:"PORT" envDefault:"8080"`
	AppEnv       string `env:"APP_ENV" envDefault:"development"`
	LogLevel     string `env:"LOG_LEVEL" envDefault:"INFO"`
	MongoURI     string `env:"MONGO_URI" envDefault:"mongodb://localhost:27017"`
	DBUser       string `env:"DB_USER"`
	DBPass       string `env:"DB_PASS"`
	DBName       string `env:"DB_NAME" envDefault:"mises_website"`
	AllowOrigins string `env:"ALLOW_ORIGINS" envDefault:""`
	StorageHost  string `env:"STORAGE_HOST" envDefault:"http://localhost/"`
	StorageKey   string `env:"STORAGE_KEY" envDefault:""`
	StorageSalt  string `env:"STORAGE_SALT" envDefault:""`
	RootPath     string
}

func init() {
	fmt.Println("socialsvc env initializing...")
	//_, b, _, _ := runtime.Caller(0)
	appEnv := os.Getenv("APP_ENV")
	projectRootPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	envPath := projectRootPath + "/.env"
	appEnvPath := envPath + "." + appEnv
	localEnvPath := appEnvPath + ".local"
	_ = godotenv.Load(filtePath(localEnvPath, appEnvPath, envPath)...)
	Envs = &Env{}
	err = env.Parse(Envs)
	if err != nil {
		panic(err)
	}
	Envs.RootPath = projectRootPath
	fmt.Println("socialsvc env root " + projectRootPath)
	fmt.Println("socialsvc env loaded...")
	view.SetupImageStorage(Envs.StorageHost, Envs.StorageKey, Envs.StorageSalt)
}

func filtePath(paths ...string) []string {
	result := make([]string, 0)
	for _, path := range paths {
		if _, err := os.Stat(path); err == nil {
			result = append(result, path)
		}
	}
	return result
}
