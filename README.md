# in-memory
Golang: in-memory package

# How to use
```bash
go get -u github.com/rail-sabirov/in-memory
```

after in main.go
```Go
package main

import (
  "fmt"
  cache "github.com/rail-sabirov/im-memory"
  )

func main() {
  c := cache.New()
  
  c.Set("userId", 42)
  userId := c.Get("userId")
  fmt.Println(userId) // 42
  
  c.Delete("userId")
  userId = c.Get("userId")
  fmt.Println(userId) // <nil>
}
```
