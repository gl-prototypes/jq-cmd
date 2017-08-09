package main

import (
	"log"
	"os/exec"

	"github.com/gliderlabs/ssh"
)

func main() {
	ssh.Handle(func(sess ssh.Session) {
		cmd := exec.Command("jq", sess.Command()...)
		cmd.Stdin = sess
		cmd.Stdout = sess
		cmd.Stderr = sess.Stderr()
		if err := cmd.Run(); err != nil {
			sess.Exit(1)
		}
	})

	log.Println("serving on 2222...")
	log.Fatal(ssh.ListenAndServe(":2222", nil, ssh.HostKeyFile("/tmp/hostkey")))
}
