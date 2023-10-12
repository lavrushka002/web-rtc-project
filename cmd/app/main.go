package main

import (
	"fmt"
	"web-rtc-project/internal/client"
)

func main() {
	c := client.New("localhost", "8080", "user", "pass")
	fmt.Println("Hello world", c.IP(), c.Port(), c.Username(), c.Password())
}
