package main

import (
	"log"
	"os/exec"

	"github.com/gliderlabs/ssh"
)

func main() {
	ssh.Handle(func(s ssh.Session) {
		cmd := exec.Command("jq", "--help")
		out, err := cmd.CombinedOutput()
		s.Write(out)
		if err != nil {
			s.Exit(1)
		}
	})

	log.Println("serving on 2222...")
	log.Fatal(ssh.ListenAndServe(":2222", nil, ssh.HostKeyFile("/tmp/hostkey")))
}
