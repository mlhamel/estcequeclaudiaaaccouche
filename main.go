package main

import (
	"fmt"

	"github.com/docopt/docopt-go"
	"github.com/mlhamel/accouchement/store"
)

func main() {
	usage := `Est-ce que Claudia a accouché?.

Usage:
  accouchement disable [--redis=<url>]
  accouchement enable [--redis=<url>]
  accouchement toggle [--redis=<url>]
  accouchement serve [--port=<port>] [--redis=<url>]
  accouchement status [--redis=<url>]
	accouchement [--port=<port>] [--redis=<url>]
  accouchement -h | --help
  accouchement --version

Options:
  --redis=<url>      Change Redis configuration to [default: redis://@192.168.64.42:6379].
  --port=<port>      Port to serve [default: 4242].
  -h --help          Show this screen.
  --version          Show version.`

	arguments, _ := docopt.Parse(usage, nil, true, "Accouchement", false)

	redisURL := arguments["--redis"].(string)
	port := arguments["--port"].(string)

	dataStore, _ := store.NewStore(store.REDIS, redisURL, "")
	statusManager := NewStatusManager(dataStore, No)

	statusManager.Refresh()

	switch {
	case arguments["disable"]:
		statusManager.Disable()
		fmt.Println(statusManager.Value())
	case arguments["enable"]:
		statusManager.Enable()
		fmt.Println(statusManager.Value())
	case arguments["toggle"]:
		statusManager.Toggle()
		fmt.Println(statusManager.Value())
	case arguments["serve"]:
		Serve(statusManager, port)
	case arguments["status"]:
		fmt.Println(statusManager.Value())
	default:
		Serve(statusManager, port)
	}
}
