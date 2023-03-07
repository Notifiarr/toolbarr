package logs

import (
	"os"
	"strconv"
)

// FileMode is used to unmarshal a unix file mode from the config file.
type FileMode string

// String creates a unix-octal version of a file mode.
func (f FileMode) String() string {
	return string(f)
}

// Mode returns the compatible os.FileMode.
func (f FileMode) Mode() os.FileMode {
	val, _ := strconv.ParseUint(string(f), 0, 32)
	return os.FileMode(val)
}
