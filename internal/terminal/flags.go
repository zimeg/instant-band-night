package terminal

import (
	"strconv"
	"strings"

	"github.com/spf13/pflag"
)

// FlagToBool converts the flag string value into a bool or panics
func FlagToBool(flag *pflag.Flag) bool {
	return flag.Value.String() == "true"
}

// FlagToInt converts the flag string value into an int or panics
func FlagToInt(flag *pflag.Flag) (value int) {
	value, _ = strconv.Atoi(flag.Value.String())
	return
}

// FlagToString returns the string flag value or panics if not found
func FlagToString(flag *pflag.Flag) string {
	return strings.TrimSpace(flag.Value.String())
}
