package main

type ChatClientConfig struct {
	Port string
}

func validateConfig(config map[string]string) ChatClientConfig {
	port := config["Port"]
	if port == "" {
		port = "2022"
	}
	return ChatClientConfig{
		Port: port,
	}
}
