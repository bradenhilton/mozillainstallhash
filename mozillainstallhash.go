package mozillainstallhash

import (
	"fmt"

	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"

	"github.com/bradenhilton/cityhash"
)

// GetMozillaInstallHash gets the Mozila install hash for the path in
// installPath.
//
// installPath should be the path to the parent directory of the executable.
//
// A UTF-16 encoding of installPath is transformed to a byte array and passed
// to CityHash64 (version 1).
//
// Other than the UTF-16 encoding and byte array conversion, installPath is
// hashed as is, so must contain the correct path separators for the
// intended operating system, with any trailing path separator removed.
//
// It returns a string of the hash in uppercase hexadecimal format.
func GetMozillaInstallHash(installPath string) string {
	endianness := unicode.LittleEndian
	bomPolicy := unicode.IgnoreBOM
	encoder := unicode.UTF16(endianness, bomPolicy).NewEncoder()

	pathBytes, _, _ := transform.Bytes(encoder, []byte(installPath))
	pathSize := uint32(len(pathBytes))

	hash := cityhash.CityHash64(pathBytes, pathSize)

	return fmt.Sprintf("%X", hash)
}
