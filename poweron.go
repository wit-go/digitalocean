package digitalocean

import (
	"context"

	"golang.org/x/oauth2"

	"github.com/digitalocean/godo"

	"go.wit.com/log"
)

func (d *DigitalOcean) PowerOn(dropletID int) error {
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: d.token})
	oauthClient := oauth2.NewClient(context.Background(), tokenSource)
	client := godo.NewClient(oauthClient)

	ctx := context.TODO()

	// Create a request to power on the droplet.
	_, _, err := client.DropletActions.PowerOn(ctx, dropletID)
	if err != nil {
		return err
	}

	log.Printf("Power-on signal sent to droplet with ID: %d\n", dropletID)
	return nil
}

func (d *DigitalOcean) PowerOff(dropletID int) error {
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: d.token})
	oauthClient := oauth2.NewClient(context.Background(), tokenSource)
	client := godo.NewClient(oauthClient)

	ctx := context.TODO()

	// Create a request to power on the droplet.
	_, _, err := client.DropletActions.PowerOff(ctx, dropletID)
	if err != nil {
		return err
	}

	log.Printf("Power-off signal sent to droplet with ID: %d\n", dropletID)
	return nil
}

/*
func (d *DigitalOcean) Destroy(dropletID int) error {
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: d.token})
	oauthClient := oauth2.NewClient(context.Background(), tokenSource)
	client := godo.NewClient(oauthClient)

	ctx := context.TODO()

	// Create a request to power on the droplet.
	_, _, err := client.DropletActions.Delete(ctx, dropletID)
	if err != nil {
		return err
	}

	log.Printf("Destroy sent to droplet with ID: %d\n", dropletID)
	return nil
}
*/

// createDroplet creates a new droplet in the specified region with the given name.
func (d *DigitalOcean) deleteDroplet(drop *Droplet) error {
	// Create an OAuth2 token.
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: d.token})

	// Create an OAuth2 client.
	oauthClient := oauth2.NewClient(context.Background(), tokenSource)

	// Create a DigitalOcean client with the OAuth2 client.
	client := godo.NewClient(oauthClient)

	ctx := context.TODO()
	log.Warn("deleteDroplet() going to delete ID =", drop.ID, "Name =", drop.GetName())
	response, err := client.Droplets.Delete(ctx, drop.ID)
	log.Warn(response)
	return err
}
