package main

import (
	"log"

	"github.com/darcysail/kubespy/cmd"
)

func init() {
	// Turn off timestamp prefix for `log.Fatal*`.
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
}

func main() {
	cmd.Execute()
}
