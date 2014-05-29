package helpers

import (
	"ephemeris/config"
	"fmt"
)

func URI(format string, in ...interface{}) string {
	return fmt.Sprintf("%s/%s", config.Ephemeris.APIRoot, fmt.Sprintf(format, in...))
}
