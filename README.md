# tuna
![logo](media/tuna.svg)
Simple yet functional self updater written in GO.

### Functionality
It will update itself with the latest release from the given URL by deleting the old binary and downloading the new one in its place.

### Example
``` go
package main

import (
	"github.com/8ff/tuna"
)

func main() {
	// Determine OS and ARCH
	osRelease := runtime.GOOS
	arch := runtime.GOARCH

	// Build URL
	e := tuna.SelfUpdate(fmt.Sprintf("https://github.com/myuser/myapp/releases/download/latest/myapp.%s.%s", osRelease, arch))
	if e != nil {
		log.Fatal(e.Error())
	} else {
		log.Println("Updated successfully!")
	}
}
```
