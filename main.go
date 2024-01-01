package digitalocean

import 	(
	"os"
	"go.wit.com/log"
	"go.wit.com/gui/gui"
)

var myDo *DigitalOcean

// This is initializes the main DO object
// You can only have one of these
func New(p *gui.Node) *DigitalOcean {
	if myDo != nil {return myDo}
	myDo = new(DigitalOcean)
	myDo.ready = false
	myDo.parent = p

	myDo.dropMap = make(map[int]*Droplet)

	// Your personal API token from DigitalOcean.
	myDo.token = os.Getenv("DIGITALOCEAN_TOKEN")

	myDo.window = p.NewWindow("DigitalOcean Control Panel")

	// make a group label and a grid
	myDo.group = myDo.window.NewGroup("droplets:").Pad()
	myDo.grid = myDo.group.NewGrid("grid", 2, 1).Pad()

	myDo.ready = true
	myDo.Hide()
	return myDo
}

// Returns true if the status is valid
func (d *DigitalOcean) Ready() bool {
	if d == nil {return false}
	return d.ready
}

func (d *DigitalOcean) Show() {
	if ! d.Ready() {return}
	log.Info("digitalocean.Show() window")
	if d.hidden {
		d.window.Show()
	}
	d.hidden = false
}

func (d *DigitalOcean) Hide() {
	if ! d.Ready() {return}
	log.Info("digitalocean.Hide() window")
	if ! d.hidden {
		d.window.Hide()
	}
	d.hidden = true
}

func (d *DigitalOcean) Update() bool {
	if ! d.Ready() {return false}
	d.ListSSHKeyID()
	if d.ListDroplets() {
		for _, droplet := range d.dpolled {
			// check if the droplet ID already exists
			if (d.dropMap[droplet.ID] == nil) {
				d.dropMap[droplet.ID] = d.NewDroplet(&droplet)
			} else {
				log.Log(POLL, "droplet.Update()", droplet.ID, droplet.Name, "already exists")
				d.dropMap[droplet.ID].Update(&droplet)
				continue
			}
		}
	} else {
		log.Error(d.err, "Error listing droplets")
		return false
	}
	return true
}
