package digitalocean

import (
	"context"

	"golang.org/x/oauth2"

	"github.com/digitalocean/godo"
)

// ListDroplets fetches and prints out the droplets along with their IPv4 and IPv6 addresses.
func (d *DigitalOcean) ListDroplets() bool {
	// OAuth token for authentication.
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: d.token})

	// OAuth2 client.
	oauthClient := oauth2.NewClient(context.Background(), tokenSource)

	// DigitalOcean client.
	client := godo.NewClient(oauthClient)

	// Context.
	ctx := context.TODO()

	// List all droplets.
	d.dpolled, _, d.err = client.Droplets.List(ctx, &godo.ListOptions{})
	if d.err != nil {
		return false
	}

	// Iterate over droplets and print their details.
	/*
	for _, droplet := range d.polled {
		fmt.Printf("Droplet: %s\n", droplet.Name)
		for _, network := range droplet.Networks.V4 {
			if network.Type == "public" {
				fmt.Printf("IPv4: %s\n", network.IPAddress)
			}
		}
		for _, network := range droplet.Networks.V6 {
			if network.Type == "public" {
				fmt.Printf("IPv6: %s\n", network.IPAddress)
			}
		}
		fmt.Println("-------------------------")
	}
	*/

	return true
}
