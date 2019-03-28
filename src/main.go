package src

import (
	"fmt"
	"keyVerifier/core"
	"os"
	"os/signal"
)

func main() {

	intrrupt := make(chan os.Signal, 1)
	signal.Notify(intrrupt, os.Interrupt)

	go core.Process()

	for {
		select {
		case <-intrrupt:
			fmt.Print("Process has been interrupted")
		case <-core.VerificationDone:
			fmt.Print("Process finished successfully")
			os.Exit(0)
		}
	}
}
