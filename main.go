package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
	"github.com/mlhamel/accouchement/status"
	"github.com/mlhamel/accouchement/store"
)

func main() {
	usage := `Est-ce que Claudia a accouchée?.

Usage:
  accouchement disable [--redis=<url>]
  accouchement enable [--redis=<url>]
  accouchement refresh [--redis=<url>]
  accouchement toggle [--redis=<url>]
  accouchement serve [--port=<port>] [--redis=<url>]
  accouchement status [--redis=<url>]
  accouchement -h | --help
  accouchement --version

Options:
  --redis=<url>      Change Redis configuration to [default: redis://@192.168.64.42:6379].
  --port=<port>      Port to serve [default: 3000].
  -h --help          Show this screen.
  --version          Show version.`

	arguments, _ := docopt.Parse(usage, nil, true, "Accouchement", false)

	redis_url := arguments["--redis"].(string)
	port := arguments["--port"].(string)

	dataStore, _ := store.NewStore(store.REDIS, redis_url, "")
	statusManager := status.NewStatus(dataStore)

	switch {
	case arguments["disable"]:
		statusManager.Disable()
		fmt.Println(statusManager.Value())
	case arguments["enable"]:
		statusManager.Enable()
		fmt.Println(statusManager.Value())
	case arguments["refresh"]:
		statusManager.Refresh()
		fmt.Println(statusManager.Value())
	case arguments["toggle"]:
		statusManager.Toggle()
		fmt.Println(statusManager.Value())
	case arguments["serve"]:
		Serve(statusManager, port)
	case arguments["status"]:
		fmt.Println(statusManager.Value())
	}
}
