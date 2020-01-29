package packet

import (
	"fmt"
	"os"
)

func sharedConfigForRegion(region string) (interface{}, error) {

	if os.Getenv("PACKET_AUTH_TOKEN") == "" {
		return nil, fmt.Errorf("you must set PACKET_AUTH_TOKEN")
	}

	config := Config{
		AuthToken: os.Getenv("PACKET_AUTH_TOKEN"),
	}

	return config.Client(), nil
}
