// any of these 3 pkgs imported,
// then will use cgo and do not link statically
// but can use -ldflags '-linkmode external -extldflags "-static"' to make gcc link statically
// if only import net, then can use "-tags netgo" to statically build

package main

import (
	"fmt"
	//_ "net/http"
	//_ "crypto/x509"
	_ "os/user"
)

func main() {
	fmt.Println(1111)
}
