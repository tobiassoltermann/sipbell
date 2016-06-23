package main

import (
	"encoding/json"
	"os"
)

type AppConfiguration struct {
	ProxyIP     string `json:proxyIP`
	ProxyPort   int    `json:proxyPort`
	Transport   string `json:transport`
	LocalPort   int    `json:localPort`
	Username    string `json:username`
	Password    string `json:password`
	Phonenumber string `json:phonenumber`
	PinNumber   int    `json:pin`
}

func ReadConfig() (AppConfiguration, error) {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	config := AppConfiguration{}
	err := decoder.Decode(&config)
	return config, err
}
