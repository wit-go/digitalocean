package digitalocean

import 	(
	"os/exec"
	"go.wit.com/log"
)

var geom string = "120x30+500+500"

func xterm(cmd string) {
	var tmp []string
	var argsXterm = []string{"nohup", "xterm", "-geometry", geom}
	// tmp = append(argsXterm, "-hold", "-e", cmd)
	tmp = append(argsXterm, "-e", cmd)
	log.Println("xterm cmd=", cmd)
	go runCommand(tmp)
}

func runCommand(cmdArgs []string) {
	log.Println("runCommand() START", cmdArgs)
	process := exec.Command(cmdArgs[0], cmdArgs[1:len(cmdArgs)]...)
	// process := exec.Command("xterm", "-e", "ping localhost")
	log.Println("runCommand() process.Start()")
	process.Start()
	log.Println("runCommand() process.Wait()")
	err := process.Wait()
	log.Error(err, "on process.Wait")
	log.Println("runCommand() NEED TO CHECK THE TIME HERE TO SEE IF THIS WORKED")
	log.Println("runCommand() OTHERWISE INFORM THE USER")
	log.Println("runCommand() END", cmdArgs)
}
