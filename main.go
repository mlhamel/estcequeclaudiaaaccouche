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
	accouchement notify [--redis=<url>] [--sid=<sid>] [--token=<token>] [--from=<from>] [--to=<to>]
  accouchement serve [--port=<port>] [--redis=<url>] [--source=<source>] [--sid=<sid>] [--token=<token>]
  accouchement status [--redis=<url>]
	accouchement [--port=<port>] [--redis=<url>] [--sid=<sid>] [--token=<token>]
  accouchement -h | --help
  accouchement --version

Options:
  --redis=<url>          	Change Redis configuration to [default: redis://@192.168.64.42:6379].
  --port=<port>          	Port to serve [default: 4242].
	--source=<source>  			Authorized source of action [default: +15149999999].
	--sid=<sid>             SID for twilio.
	--token=<token>         Token for twilio.
	--from=<from>           Source number for twilio.
	--to=<to> 							Destinatination number for twilio.
  -h --help          			Show this screen.
  --version          			Show version.`

	arguments, _ := docopt.Parse(usage, nil, true, "Accouchement", false)

	redisURL := arguments["--redis"].(string)
	port := arguments["--port"].(string)
	source := arguments["--source"].(string)

	sid := arguments["--sid"].(string)
	token := arguments["--token"].(string)

	dataStore, _ := store.NewStore(store.REDIS, redisURL, "")
	statusManager := NewStatusManager(dataStore, No, source)
	notifier := NewTwilioNotifier(sid, token, statusManager)

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
	case arguments["notify"]:
		to := arguments["--to"].(string)
		from := arguments["--from"].(string)
		notifier.NotifyInline(from, to)
	default:
		Serve(statusManager, port)
	}
}
