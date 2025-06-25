package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// type Config struct {
// 	// If you want to configure YAML based envs
// 	err error
// }

func MustLoad() error {
	env := os.Getenv("APP_ENV")

	if env == "" {
		env = "local"
	}

	// envConfigPath := fmt.Sprintf("internal\\envs\\.env.%s", env) // If you arerunning on ubuntu do sepcify accordingly
	envConfigPath := filepath.Join("internal", "envs", fmt.Sprintf(".env.%s", env))
	err := godotenv.Load(envConfigPath)
	return err
}
