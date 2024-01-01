package digitalocean

import (
	"context"
	"golang.org/x/oauth2"
	"github.com/digitalocean/godo"

	"go.wit.com/log"
)

// func (d *DigitalOcean) ListDroplets() bool {
func (d *DigitalOcean) ListSSHKeyID() error {
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: d.token})
	oauthClient := oauth2.NewClient(context.Background(), tokenSource)
	client := godo.NewClient(oauthClient)

	// List all keys.
	keys, _, err := client.Keys.List(context.Background(), &godo.ListOptions{})
	if err != nil {
		return err
	}

	d.sshKeys = keys

	// Find the key by name.
	for _, key := range keys {
		log.Log(POLL, "found ssh wierd", key.Name)
		log.Log(POLL, "found ssh key:", key)
	}
	/*
	sshKeys := []godo.DropletCreateSSHKey{
		{ID: 22994569},
		{ID: 333},
	}
	*/

	// return fmt.Errorf("SSH Key not found")
	return nil
}
