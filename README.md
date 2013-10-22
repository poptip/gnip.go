# gnip

Golang package for gnip.

# Install

```sh
go get github.com/poptip.com/gnip.go
```

# Usage

```go
package main

import (
  "github.com/poptip/gnip.go"
)

func main() {
  var mygnip := gnip.NewClient(username, password, account)
  rules, err := mygnip.GetAllActiveRules()
  if err != nil {
    // Oh no.
  }

  for _, rule := range rules {
    // Do something with `rule`.
  }
}
```
