package exec_shell_commands

import (
	"fmt"
	"os"
	"os/exec"
)

func Run() {
	c := exec.Command("top")
	c.Stdout = os.Stdout
	c.Stdin = os.Stdin
	c.Stderr = os.Stderr

	c.Run()
}

func RunOnServer() {
	c := exec.Command("ssh", "-t", "kart", "top")
	c.Stdout = os.Stdout
	c.Stdin = os.Stdin
	c.Stderr = os.Stderr

	c.Run()
}

func ConnectToAWS() {
	key := "-----BEGIN RSA PRIVATE KEY----- -----END RSA PRIVATE KEY-----"
	user := "ec2-user"
	server := "amazonaws.com"
	c := exec.Command("bash", "-c", fmt.Sprintf(`ssh -i <(echo "%s") %s@%s`, key, user, server))
	c.Stdout = os.Stdout
	c.Stdin = os.Stdin
	c.Stderr = os.Stderr

	c.Run()
}
