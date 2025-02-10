package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type McpServer struct {
	Command string            `json:"command"`
	Args    []string          `json:"args"`
	Env     map[string]string `json:"env"`
}

type Config struct {
	McpServers map[string]McpServer `json:"mcpServers"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run mcp-starter <config-file>")
		return
	}

	configPath := os.Args[1]
	config := loadConfig(configPath)
	
	// Get the single server entry
	if len(config.McpServers) != 1 {
		panic("Config must contain exactly one server definition")
	}
	
	var server McpServer
	var serverName string
	for name, s := range config.McpServers {
		serverName = name
		server = s
		break // Only need the first/only entry
	}

	executeServer(server)
}

func loadConfig(configPath string) Config {
	file, err := os.Open(configPath)
	if err != nil {
		panic(fmt.Sprintf("Error opening config file '%s': %v", configPath, err))
	}
	defer file.Close()

	var config Config
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		panic(fmt.Sprintf("Error decoding config '%s': %v", configPath, err))
	}

	return config
}

// executeServer is implemented in platform-specific files 