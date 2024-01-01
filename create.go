package digitalocean

import (
	"context"
	"strings"
	"golang.org/x/oauth2"
	"github.com/digitalocean/godo"

	"go.wit.com/log"
	"go.wit.com/gui/gadgets"
	// "go.wit.com/gui"
)

/*
// createDroplet creates a new droplet in the specified region with the given name.
func createDroplet(token, name, region, size, image string) (*godo.Droplet, error) {
	// Create an OAuth2 token.
	tokenSource := &oauth2.Token{
		AccessToken: token,
	}

	// Create an OAuth2 client.
	oauthClient := oauth2.NewClient(context.Background(), tokenSource)

	// Create a DigitalOcean client with the OAuth2 client.
	client := godo.NewClient(oauthClient)

	// Define the create request.
	createRequest := &godo.DropletCreateRequest{
		Name:   name,
		Region: region,
		Size:   size,
		Image: godo.DropletCreateImage{
			Slug: image,
		},
	}

	// Create the droplet.
	ctx := context.TODO()
	newDroplet, _, err := client.Droplets.Create(ctx, createRequest)
	if err != nil {
		return nil, err
	}

	return newDroplet, nil
}
*/

func (d *DigitalOcean) Create(name string, region string, size string, image string) {
	// Create a new droplet.
	droplet, err := d.createDropletNew(name, region, size, image)
	if err != nil {
		log.Fatalf("digitalocean.Create() Something went wrong: %s\n", err)
	}

	log.Infof("digitalocean.Create() droplet ID %d with name %s\n", droplet.ID, droplet.Name)
}

// createDroplet creates a new droplet in the specified region with the given name.
func (d *DigitalOcean) createDropletNew(name, region, size, image string) (*godo.Droplet, error) {
	log.Infof("digitalocean.createDropletNew() START name =", name)
	// Create an OAuth2 token.
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: d.token})

	// Create an OAuth2 client.
	oauthClient := oauth2.NewClient(context.Background(), tokenSource)

	// Create a DigitalOcean client with the OAuth2 client.
	client := godo.NewClient(oauthClient)

	var sshKeys []godo.DropletCreateSSHKey
	log.Info("digitalocean.createDropletNew() about to get keys. client =", client)

	// Find the key by name.
	for i, key := range d.sshKeys {
		log.Info("found ssh i =", i, key.Name)
		log.Verbose("found ssh key.Name =", key.Name)
		log.Verbose("found ssh key.Fingerprint =", key.Fingerprint)
		log.Verbose("found ssh key:", key)
		/*
		sshKeys = []godo.DropletCreateSSHKey{
			{ID: key.ID},
		}
		*/
		sshKeys = append(sshKeys, godo.DropletCreateSSHKey{ID: key.ID})
	} 

	// Define the create request.
	createRequest := &godo.DropletCreateRequest{
		Name:   name,
		Region: region,
		Size:   size,
		Image: godo.DropletCreateImage{
			Slug: image,
		},
		IPv6: true, // Enable IPv6
		SSHKeys: sshKeys, // Add SSH key IDs here
	}

	// Create the droplet.
	ctx := context.TODO()
	log.Info("digitalocean.createDropletNew() about to do client.Create(). ctx =", ctx)
	newDroplet, _, err := client.Droplets.Create(ctx, createRequest)
	log.Infof("digitalocean.createDropletNew() END newDroplet =", newDroplet)
	if err != nil {
		return nil, err
	}

	return newDroplet, nil
}

var myCreate *windowCreate

// This is initializes the main DO object
// You can only have one of these
func InitCreateWindow() *windowCreate {
	if ! myDo.Ready() {return nil}
	if myCreate != nil {
		myCreate.Show()
		return myCreate
	}
	myCreate = new(windowCreate)
	myCreate.ready = false

	myCreate.window = myDo.parent.NewWindow("Create Droplet")

	// make a group label and a grid
	myCreate.group = myCreate.window.NewGroup("droplets:").Pad()
	myCreate.grid = myCreate.group.NewGrid("grid", 2, 1).Pad()
	
	myCreate.name = gadgets.NewBasicEntry(myCreate.grid, "Name").Set("test.wit.com")

	myCreate.region = gadgets.NewBasicDropdown(myCreate.grid, "Region")

	regions := myDo.listRegions()

	// Print details of each region.
	log.Info("Available Regions:")
	for i, region := range regions {
		log.Infof("i: %d, Slug: %s, Name: %s, Available: %v\n", i, region.Slug, region.Name, region.Available)
		log.Spew(i, region)
		if len(region.Sizes) == 0 {
			log.Info("Skipping region. No available sizes region =", region.Name)
		} else {
			s := region.Name + " (" + region.Slug + ")"
			if (myCreate.regionSlug == "") {
				myCreate.regionSlug = region.Slug
			}
			myCreate.region.Add(s)
		}
	}

	myCreate.region.Custom = func() {
		s := myCreate.region.Get()
		log.Info("create droplet region changed to:", s)
		for _, region := range regions {
			if s == region.Name {
				log.Info("Found region! slug =", myCreate.regionSlug, region)
				myCreate.regionSelected = region
				log.Info("Found region! Now update all the sizes count =", len(region.Sizes))
				for _, size := range region.Sizes {
					log.Info("Size: ", size)
				}
			}
		}
	}

	myCreate.size = gadgets.NewBasicCombobox(myCreate.grid, "Size")
	myCreate.size.Add("s-1vcpu-1gb")
	myCreate.size.Add("s-1vcpu-1gb-amd")
	myCreate.size.Add("s-1vcpu-1gb-intel")
	myCreate.size.Add("s-2vcpu-4gb-120gb-intel")
	myCreate.size.Set("s-2vcpu-4gb-120gb-intel")
	myCreate.size.Custom = func() {
		size := myCreate.size.Get()
		log.Info("Create() need to verify size exists in region. Digital Ocean size.Slug =", size)
	}

	myCreate.memory = gadgets.NewBasicDropdown(myCreate.grid, "Memory")
	myCreate.memory.Add("1 GB")
	myCreate.memory.Add("2 GB")
	myCreate.memory.Add("4 GB")
	myCreate.memory.Add("8 GB")
	myCreate.memory.Add("16 GB")
	myCreate.memory.Add("32 GB")
	myCreate.memory.Add("64 GB")
	myCreate.memory.Add("96 GB")
	myCreate.memory.Add("128 GB")
	myCreate.memory.Add("256 GB")
	myCreate.memory.Custom = func() {
		for _, size := range myCreate.regionSelected.Sizes {
			log.Info("Size: ", size)
		}
		myCreate.UpdateSize()
	}

	myCreate.image = gadgets.NewBasicCombobox(myCreate.grid, "Image")
	myCreate.image.Add("debian-12-x64")
	myCreate.image.Add("ubuntu-20-04-x64")
	myCreate.image.Set("debian-12-x64")

	// myCreate.nvme = gadgets.NewBasicCheckbox(myCreate.grid, "NVMe")

	myCreate.group.NewLabel("Create Droplet")

	// box := myCreate.group.NewBox("vBox", false).Pad()
	box := myCreate.group.NewBox("hBox", true).Pad()
	box.NewButton("Cancel", func () {
		myCreate.Hide()
	})
	box.NewButton("Create", func () {
		name := myCreate.name.Get()
		size := myCreate.size.Get()
		region := myCreate.regionSlug
		image := myCreate.image.Get()
		if (region == "") {
			log.Info("Create() droplet name =", name, "region =", region, "size =", size, "image", image)
			log.Info("Create() region lookup failed")
			return
		}
		log.Info("Create() droplet name =", name, "region =", region, "size =", size, "image", image)
		myDo.Create(name, region, size, image)
		myCreate.Hide()
	})

	myCreate.ready = true
	myDo.create = myCreate
	return myCreate
}

// Find the size
func (d *windowCreate) UpdateSize() {
	if ! d.Ready() {return}
	log.Info("Now find the size. sizes count =", len(myCreate.regionSelected.Sizes))
	var s string
	m := myCreate.memory.Get()
	switch m {
	case "1 GB":
		s = "cpu-1gb-"
	case "2 GB":
		s = "cpu-2gb-"
	case "4 GB":
		s = "cpu-4gb-"
	case "8 GB":
		s = "cpu-8gb-"
	case "16 GB":
		s = "cpu-16gb-"
	case "32 GB":
		s = "cpu-32gb-"
	case "64 GB":
		s = "cpu-64gb-"
	case "96 GB":
		s = "cpu-96gb-"
	case "128 GB":
		s = "cpu-128gb-"
	case "256 GB":
		s = "cpu-256gb-"
	default:
		s = "cpu-4gb-"
	}
	for _, size := range myCreate.regionSelected.Sizes {
		if strings.Contains(size, s) {
			log.Info("Found Size! size.Slug =", size, "contains", s)
			myCreate.size.Set(size)
			return
		}
	}
	log.Info("memory =", myCreate.memory.Get())
}

// Returns true if the status is valid
func (d *windowCreate) Ready() bool {
	if d == nil {return false}
	return d.ready
}

func (d *windowCreate) Show() {
	if ! d.Ready() {return}
	log.Info("digitalocean.Show() window")
	if d.hidden {
		d.window.Show()
	}
	d.hidden = false
}

func (d *windowCreate) Hide() {
	if ! d.Ready() {return}
	log.Info("digitalocean.Hide() window")
	if ! d.hidden {
		d.window.Hide()
	}
	d.hidden = true
}
