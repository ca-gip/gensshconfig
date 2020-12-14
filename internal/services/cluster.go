package services

import (
	"bytes"
	"io/ioutil"
	"net"
	"regexp"
	"strings"
)

var InventoryParser = regexp.MustCompile(`(?m)(?P<hostname>[a-zA-Z0-9-]+).*ansible_host=(?P<ip>\b(?:[0-9]{1,3}\.){3}[0-9]{1,3}\b)`)

type Host struct {
	Hostname string
	Addr     net.IP
}

type Cluster struct {
	Name             string
	HostAbsolutePath string
	Hosts            []*Host
}

func (c *Cluster) readHostFile() (result string, err error) {
	rawResult, err := ioutil.ReadFile(c.HostAbsolutePath)

	result = string(rawResult)

	return
}

func (c *Cluster) readHostInCluster() (err error) {
	hostsFile, err := c.readHostFile()

	if err != nil {
		return
	}

	for _, match := range InventoryParser.FindAllStringSubmatch(hostsFile, -1) {
		c.Hosts = append(c.Hosts, &Host{
			Hostname: c.normalizedHostname(match[1]),
			Addr:     net.ParseIP(match[2]),
		})
	}

	return
}

func (c *Cluster) normalizedHostname(hostname string ) string {
	var buffer bytes.Buffer
	buffer.WriteString(strings.ReplaceAll(c.Name, "_", "-"))
	buffer.WriteString("-")
	buffer.WriteString(hostname)

	return buffer.String()
}