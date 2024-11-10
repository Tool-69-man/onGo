package main

import (
	"fmt"
	"os"
)

func main() {
	// -gcflags "-N -l -m"

	for i := 0; i < len(os.Args); i++ {
		fmt.Println("===:", os.Args[i])
	}
	for _, x := range os.Args {
		fmt.Println(x)
	}

}
