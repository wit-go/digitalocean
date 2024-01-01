/* 
	The Digital Ocean Struct
*/

package digitalocean

import (
	"github.com/digitalocean/godo"

	"go.wit.com/gui/gui"
	"go.wit.com/gui/gadgets"
)

type DigitalOcean struct {
	ready		bool
	hidden		bool
	err		error

	token		string // You're Digital Ocean API key
	dpolled		[]godo.Droplet
	sshKeys		[]godo.Key

	dropMap		map[int]*Droplet
	create		*windowCreate

	parent	*gui.Node // should be the root of the 'gui' package binary tree
	window	*gui.Node // our window for displaying digital ocean droplets
	group	*gui.Node
	grid	*gui.Node

	dGrid	*gui.Node // the grid for the droplets

	// Primary Directives
	status		*gadgets.OneLiner
	summary		*gadgets.OneLiner
	statusIPv4	*gadgets.OneLiner
	statusIPv6	*gadgets.OneLiner
}

type windowCreate struct {
	ready		bool
	hidden		bool
	err		error

	parent	*gui.Node // should be the root of the 'gui' package binary tree
	window	*gui.Node // our window for displaying digital ocean droplets
	group	*gui.Node
	grid	*gui.Node

	regionSelected	godo.Region
	regionSlug	string
	tag		*gadgets.OneLiner
	name		*gadgets.BasicEntry
	region		*gadgets.BasicDropdown
	size		*gadgets.BasicCombobox
	memory		*gadgets.BasicDropdown
	image		*gadgets.BasicCombobox
	// nvme		*gadgets.BasicCheckbox
}

type ipButton struct {
	ip	*gui.Node
	c	*gui.Node
}

type Droplet struct {
	ID		int
	image		string
	memory		int
	disk		int

	ready		bool
	hidden		bool
	err		error

	poll		*godo.Droplet // store what the digital ocean API returned

	nameN		*gui.Node
	sizeSlugN	*gui.Node
	statusN		*gui.Node
	imageN		*gui.Node

	destroy		*gui.Node
	connect		*gui.Node
	poweron		*gui.Node
	poweroff	*gui.Node
	edit		*gui.Node

	ipv4		*gui.Node
	ipv6		*gui.Node
}
