// +build integration

package main

import (
	"fmt"
	"testing"

)

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}

func newDockerManager(t *testing.T) *DockerManager {
	config := NewConfig()
	config.dockerHost = testHTTPServer.URL
	config.domain = NewDomain(".test")

	dnsServer := NewDNSServer(config)
	
	docker, err := NewDockerManager(config, dnsServer, nil)
	if err != nil {
		t.Fatal("Cannot connect to mock Docker")
	}
	
	return docker
}

func TestDockerManager(t *testing.T) {
	docker := newDockerManager(t)

	service, err := docker.getService("332375cfbc23edb921a21026314c3497674ba8bdcb2c85e0e65ebf2017f688ce")
	if err != nil {
		t.Fatal("getService:", err.Error())
	}
	assertEqual(t, service, "", "")
	
}
