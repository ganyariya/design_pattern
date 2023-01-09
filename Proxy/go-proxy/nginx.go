package goproxy

import "fmt"

type Nginx struct {
	serverApp ServerInterface
}

func (n *Nginx) HttpRequest(message string) string {
	fmt.Printf("this is log: %s", message)
	if message == "admin" {
		return "admin request"
	}
	return n.serverApp.HttpRequest(message)
}
