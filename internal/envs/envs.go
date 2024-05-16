package envs

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Envs struct {
	ServerIP   string
	ServerPort int
	ClientIP   string
	ClientPort int
	Mode       string
}

func GetEnvs() Envs {
	serverPort, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		fmt.Printf("Failed to parse server port variable: %s\n", err)
		serverPort = 8080
	}

	clientPort, err := strconv.Atoi(os.Getenv("CLIENT_PORT"))
	if err != nil {
		fmt.Printf("Failed to parse client port variable: %s\n", err)
		clientPort = 8080
	}

	mode := strings.ToLower(os.Getenv("MODE"))
	if mode != "server" || mode != "client" {
		fmt.Printf("Unsupported mode %q. Supported: server, client", mode)
		mode = "server"
	}

	return Envs{
		ServerIP:   os.Getenv("SERVER_IP"),
		ServerPort: serverPort,
		ClientIP:   os.Getenv("CLIENT_IP"),
		ClientPort: clientPort,
		Mode:       mode,
	}
}
