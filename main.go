package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/vishvananda/netlink"
	"github.com/vishvananda/netlink/nl"
)

func toJson(t interface{}) {
	jsonBytes, err := json.Marshal(t)
	if err != nil {
		fmt.Println("JSON Marshal error:", err)
		return
	}

	fmt.Println(string(jsonBytes))
}

func ipAddrShow() ([]AddrShowResult, error) {
	addrShowResults := []AddrShowResult{}
	addrs, err := netlink.AddrList(nil, nl.FAMILY_ALL)
	if err != nil {
		return nil, err
	}
	for _, addr := range addrs {
		peerMask, err := strconv.Atoi(strings.Split(addr.Peer.String(), "/")[1])
		if err != nil {
			return nil, err
		}
		peer := Peer{addr.Peer.IP, peerMask}

		mask, err := strconv.Atoi(strings.Split(addr.IPNet.String(), "/")[1])
		if err != nil {
			return nil, err
		}
		addrShowResult := AddrShowResult{addr, mask, peer}
		addrShowResults = append(addrShowResults, addrShowResult)
	}

	return addrShowResults, nil
}

func ipLinkShow() ([]LinkShowResult, error) {
	linkShowResults := []LinkShowResult{}
	links, err := netlink.LinkList()
	if err != nil {
		return nil, err
	}
	for _, link := range links {
		linkAttrs := link.Attrs()
		linkShowResult := LinkShowResult{linkAttrs, link.Type(), linkAttrs.HardwareAddr.String()}
		linkShowResults = append(linkShowResults, linkShowResult)
	}
	return linkShowResults, nil
}

func main() {
	var (
		opt1 = flag.Bool("addr", false, "First string option")
		opt2 = flag.Bool("link", false, "Second string option")
	)
	flag.Parse()

	if *opt1 {
		result, _ := ipAddrShow()
		toJson(result)
	} else if *opt2 {
		result, _ := ipLinkShow()
		toJson(result)
	}

}

type LinkShowResult struct {
	*netlink.LinkAttrs
	Type         string
	HardwareAddr string
}

type AddrShowResult struct {
	netlink.Addr
	Mask int
	Peer Peer
}

type Peer struct {
	IP   net.IP
	Mask int
}
