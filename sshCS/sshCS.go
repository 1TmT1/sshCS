package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/ssh"
)

func main() {
	user := "username"
	pass := "password"
	targethost := "ip:port"

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(pass),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	conn, err := ssh.Dial("tcp", targethost, config)
	if err != nil {
		log.Fatalln("Failed to dial: ", err)
	}
	sess, err := conn.NewSession()
	if err != nil {
		log.Fatalln("Failed to create session: ", err)
	}
	stdin, _ := sess.StdinPipe()
	sess.Stdout = os.Stdout
	sess.Stderr = os.Stderr
	sess.Shell()
	cmds := []string{
		"echo Hello World!",
		"exit",
	}
	for _, cmd := range cmds {
		fmt.Fprintf(stdin, "%s\n", cmd)
	}
	sess.Wait()
	sess.Close()
}
