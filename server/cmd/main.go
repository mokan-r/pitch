package main

import (
	"fmt"
	s "github.com/mokan-r/pitch/server"
	"github.com/sirupsen/logrus"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 5454))
	if err != nil {
		logrus.Fatal(err)
	}
	server := s.New()
	server.Serve(lis)
}
