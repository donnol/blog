package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/donnol/do"
)

func main() {
	cmd := do.Must1(exec.LookPath("jdmgr"))

	o := exec.Command(cmd, "--config=abc.toml", "server")
	o.Stderr = os.Stdout
	r := do.Log1(o.Output())

	fmt.Println(string(r))
}
