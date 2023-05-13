package tests

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/MiroslavZaprazny/gogit/commands"
)

func TestInitializeRepo(t *testing.T) {
    matches, err := filepath.Glob(".gogit")

    if (err != nil) {
        fmt.Println(err.Error())
    }

    if (len(matches) == 1) {
        os.Remove(commands.BaseDir)
    }

    commands.Init()
}
