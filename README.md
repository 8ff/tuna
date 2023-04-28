# tuna 🐟
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
	e := tuna.SelfUpdate("https://example.com/releases/myApp")
	if e != nil {
		log.Fatal(e.Error())
	} else {
		log.Println("Updated successfully!")
	}
}
```
