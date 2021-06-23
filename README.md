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

	"github.com/bradenhilton/mozillainstallhash"
)

func main() {
	if len(os.Args) > 1 {
		paths := os.Args[1:]
		for _, path := range paths {
			hash := mozillainstallhash.GetMozillaInstallHash(path)
			fmt.Printf("Hash for \"%s\": %s\n", path, hash)
		}
	} else {
		log.Println(fmt.Errorf("error: no path specified"))
		fmt.Println(`Usage: get_mozilla_install_hash "<path>" "<path>" etc.

Where <path> is a string describing the parent directory of
the executable e.g. "C:\Program Files\Mozilla Firefox"`)
	}
}
```

Run it:

```
go run get_mozilla_install_hash.go "C:\Program Files\Mozilla Firefox" "/usr/lib/firefox"
```

Output:

```
Hash for "C:\Program Files\Mozilla Firefox": 308046B0AF4A39CB
Hash for "/usr/lib/firefox": 4F96D1932A9F858E
```