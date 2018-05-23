package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
	"github.com/usefathom/fathom/pkg/datastore/sqlstore"
	"math/rand"
	"os"
)

// Config wraps the configuration structs for the various application parts
type Config struct {
	Database *sqlstore.Config
	Secret   string
}

// Parses the supplied file + environment into a Config struct
func Parse(file string) *Config {
	var cfg Config
	var err error

	if file != "" {
		// get absolute path to config file
		wd, _ := os.Getwd()
		absfile := wd + "/" + file

		// check if file exists
		_, err := os.Stat(absfile)
		fileNotExists := os.IsNotExist(err)

		// Print config file location
		if file != ".env" || !fileNotExists {
			log.Printf("Configuration file: %s", absfile)
		}

		// Abort if custom config file does not exist
		if file != ".env" && fileNotExists {
			log.Fatalf("Error reading configuration. File `%s` does not exist.", file)
		}

		// read file into env values
		err = godotenv.Load(absfile)
		if err != nil {
			log.Fatalf("Error parsing configuration file: %s", err)
		}
	}

	// with config file loaded into env values, we can now parse env into our config struct
	err = envconfig.Process("Fathom", &cfg)
	if err != nil {
		log.Fatalf("Error parsing configuration from environment: %s", err)
	}

	// alias sqlite to sqlite3
	if cfg.Database.Driver == "sqlite" {
		cfg.Database.Driver = "sqlite3"
	}

	// if secret key is empty, use a randomly generated one
	if cfg.Secret == "" {
		cfg.Secret = randomString(40)
	}

	return &cfg
}

func randomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(65 + rand.Intn(25)) //A=65 and Z = 65+25
	}

	return string(bytes)
}
