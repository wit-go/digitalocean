package digitalocean

// initializes logging and command line options

import (
	arg "github.com/alexflint/go-arg"
	log "go.wit.com/log"
)

var INFO log.LogFlag
var POLL log.LogFlag
var argDo ArgsDo

// This struct can be used with the go-arg package
type ArgsDo struct {
	DigitalOceanTimer int `arg:"--digitalocean-poll-interval" help:"how often to poll droplet status (default 60 seconds)"`
}

func init() {
	arg.Register(&argDo)

	INFO.B = false
	INFO.Name = "INFO"
	INFO.Subsystem = "digitalocean"
	INFO.Desc = "Enable log.Info()"
	INFO.Register()

	POLL.B = false
	POLL.Name = "POLL"
	POLL.Subsystem = "digitalocean"
	POLL.Desc = "Show droplet status polling"
	POLL.Register()
}
