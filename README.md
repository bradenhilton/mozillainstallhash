# mozillainstallhash

Gets the Mozilla install hash used to differentiate between installs in `installs.ini` and `profiles.ini`.

### Example
First run:

```
go get github.com/bradenhilton/mozillainstallhash
```

(and possibly)

```
go get github.com/bradenhilton/cityhash
```

Create a new file e.g. `get_mozilla_install_hash.go`:

```go
package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/bradenhilton/mozillainstallhash"
)

const usage = `Gets the install hash used to determine different versions of Mozilla software
Usage:
  get_mozilla_install_hash ["<path>"] ["<path>"] etc.

Where <path> is a string describing the parent directory of
the executable, e.g. "C:\Program Files\Mozilla Firefox", with
platform specific path separators ("\" on Windows and "/" on Unix-like systems)

Example:
  get_mozilla_install_hash "C:\Program Files\Mozilla Firefox"
  308046B0AF4A39CB

  get_mozilla_install_hash "C:/Program Files/Mozilla Firefox"
  9D561FCD08DC6D55`

func main() {
	if len(os.Args) > 1 {
		paths := os.Args[1:]
		for _, path := range paths {
			path = strings.TrimSuffix(path, "/")
			path = strings.TrimSuffix(path, "\\")
			hash := mozillainstallhash.GetMozillaInstallHash(path)
			fmt.Print(hash)
		}
	} else {
		log.Println(fmt.Errorf("error: no path specified"))
		fmt.Println(usage)
	}
}
```

Run it:

```
go run get_mozilla_install_hash.go "C:\Program Files\Mozilla Firefox" "/usr/lib/firefox"
```

Output:

```
308046B0AF4A39CB
4F96D1932A9F858E
```