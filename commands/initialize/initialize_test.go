package initialize

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestMain(m *testing.M) {
    exit := m.Run()

    _, err := filepath.Glob(".gogit")

    if (err != nil) {
        fmt.Println(err.Error())
    }

    os.Remove(BaseDir)

    os.Exit(exit)
}

func TestInitializeRepo(t *testing.T) {
    Run()
}
