package lib

import (
	"fmt"

	display "github.com/djdhm/external-dep/lib"
)

const Version = "pseudo-version"

func TestLib() {
	fmt.Println("I am moab-dep-3 version " + Version)
	fmt.Println("Using external deps version " + display.Version())
}
