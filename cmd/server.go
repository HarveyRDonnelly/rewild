// Package main is responsible for orchestrating the application
// server and other background processes.
package main

// Import packages
import (
	"github.com/gookit/config/v2"
	"os"
	"rewild-it/api"
	"rewild-it/api/db"
	"sync"
)

// Initialise a wait group object
var wg sync.WaitGroup
var conn db.Connection

func main() {

	// Load environment variables
	var which_env, is_env_set = os.LookupEnv("SERVER_ENV")
	if !is_env_set {
		which_env = "default"
	}

	configFileBytes, _ := os.ReadFile("config/" + which_env + ".json")
	configFileStr := string(configFileBytes)
	configFileStr = os.ExpandEnv(configFileStr)

	// Load environment config
	err := config.LoadStrings("json", configFileStr)
	if err != nil {
		panic(err)
	}

	// Connect to DB
	conn = db.Connection{
		Host:     config.String("db.host"),
		Port:     config.Int("db.port"),
		User:     config.String("db.user"),
		Password: config.String("db.password"),
		Database: config.String("db.database"),
		Gateway:  nil,
	}
	conn = db.Connect(conn)

	api.SetDB(conn)

	// Run API process
	wg.Add(1)
	go func() {
		api.Run()
		wg.Done()
	}()

	// Wait until all processes terminate
	wg.Wait()
}
