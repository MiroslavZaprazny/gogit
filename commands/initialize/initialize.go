package initialize

import (
	"fmt"
	"os"
)

const BaseDir = ".gogit/" 

func Run () {
    fmt.Println("Initializing gogit repo ...")
    err := os.Mkdir(BaseDir, 0700)
    if err != nil {
        fmt.Printf("err: %v\n", err)
    }
    folders := [3]string{"objects", "refs", "refs/head"}

    for _, folder := range folders {
        err := os.Mkdir(BaseDir + folder , 0700)
        if err != nil {
            fmt.Printf("err: %v\n", err)
        }
    }
}
