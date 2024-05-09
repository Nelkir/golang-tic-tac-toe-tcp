package args

import (
	"fmt"
	"os"
	"strconv"
)

type Args struct {
	ServerIP   string
	ServerPort int
	ClientIP   string
	ClientPort int
	Mode       string
}

func GetArgs() Args {
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

	return Args{
		ServerIP:   os.Getenv("SERVER_IP"),
		ServerPort: serverPort,
		ClientIP:   os.Getenv("CLIENT_IP"),
		ClientPort: clientPort,
	}
}
