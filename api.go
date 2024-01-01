package digitalocean

import (
	"context"

	"golang.org/x/oauth2"
	"github.com/digitalocean/godo"

	"go.wit.com/log"
)

func (d *DigitalOcean) listRegions() []godo.Region {
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: d.token})
	oauthClient := oauth2.NewClient(context.Background(), tokenSource)
	client := godo.NewClient(oauthClient)

	ctx := context.TODO()

	// Retrieve all regions.
	regions, _, err := client.Regions.List(ctx, &godo.ListOptions{})
	if err != nil {
		d.err = err
		log.Warn(err, "digitalocean.listRegions() failed")
		return nil
	}

	/*
	// Print details of each region.
	fmt.Println("Available Regions:")
	for _, region := range regions {
		fmt.Printf("Slug: %s, Name: %s, Available: %v\n", region.Slug, region.Name, region.Available)
	}
	*/

	return regions
}
