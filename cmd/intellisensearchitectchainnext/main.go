// cmd/intellisensearchitectchainnext/main.go
package main

import (
"flag"
"log"
"os"

"intellisensearchitectchainnext/internal/intellisensearchitectchainnext"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := intellisensearchitectchainnext.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
