package services

import (
	"fmt"
	"io/ioutil"
//	"log"
	"os"
	"gopkg.in/yaml.v2"
)

// Secrets struct holds the API key and other secrets
type Secrets struct {
	APIKey     string `yaml:"api_key"`
}


// Global variable to hold the loaded secrets
var secrets Secrets

// Function to load secrets from the YAML file
func LoadSecrets() error {
	// Debug: Print current working directory
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get working directory: %v", err)
	}
	fmt.Println("Current working directory:", wd)

	filePath := "services/secrets.yaml"
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf("file does not exist: %s", filePath)
	}

	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading YAML file: %v", err)
	}

	// Debug: Print file contents
	fmt.Println("YAML File Contents:")
	fmt.Println(string(yamlFile))

	err = yaml.Unmarshal(yamlFile, &secrets)
	if err != nil {
		return fmt.Errorf("error unmarshalling YAML: %v", err)
	}
	fmt.Printf("Parsed API Key: %v\n", secrets)
	return nil
}

// Getter for the secrets
func GetSecrets() (Secrets, error) {
	if secrets.APIKey == "" {
		return Secrets{}, fmt.Errorf("secrets are not loaded or API key is empty")
	}
	return secrets, nil
}

func init() {
	fmt.Println("Initializing LoadSecrets")
	err := LoadSecrets()
	if err != nil {
		// Use log.Fatal to exit if secrets cannot be loaded
		fmt.Printf("Failed to load secrets: %v\n", err)
		os.Exit(1)
	}
}

