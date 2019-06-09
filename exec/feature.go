package exec

import "github.com/canzerapple/neonx-cli/location"

var (
	EnvironSeparator = ":"
)

func findGoBin(root location.Location) location.Location {
	return root.Child("bin/go")
}
