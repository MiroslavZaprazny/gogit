package main

import (
	"os"
    "github.com/MiroslavZaprazny/gogit/commands/initialize"
)

func main() {
    if (len(os.Args) == 1) {
        panic("Missing Arguments")
    }

    command := os.Args[1]
    switch command {
    case "init":
        initialize.Run()
    }
}
