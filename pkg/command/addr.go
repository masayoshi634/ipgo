package command

import (
	"net"
	"strconv"
	"strings"

	"github.com/vishvananda/netlink"
	"github.com/vishvananda/netlink/nl"
)

type AddrShowResult struct {
	netlink.Addr
	Mask int
	Peer Peer
}

type Peer struct {
	IP   net.IP
	Mask int
}

func AddrShow() error {
	var err error
	results, err := GetAddr()
	if err != nil {
		return err
	}
	toJson(results)
	return err
}

func AddrAdd(ipnet string, iface string) error {
	var err error
	link, err := netlink.LinkByName(iface)
	if err != nil {
		return err
	}
	addr, err := netlink.ParseAddr(ipnet)
	if err != nil {
		return err
	}
	err = netlink.AddrAdd(link, addr)
	if err != nil {
		return err
	}
	return nil
}

func AddrDelete(ipnet string, iface string) error {
	var err error
	link, err := netlink.LinkByName(iface)
	if err != nil {
		return err
	}
	addr, err := netlink.ParseAddr(ipnet)
	if err != nil {
		return err
	}
	err = netlink.AddrDel(link, addr)
	if err != nil {
		return err
	}
	return nil
}

func AddrReplace(ipnet string, iface string) error {
	var err error
	link, err := netlink.LinkByName(iface)
	if err != nil {
		return err
	}
	addr, err := netlink.ParseAddr(ipnet)
	if err != nil {
		return err
	}
	err = netlink.AddrReplace(link, addr)
	if err != nil {
		return err
	}
	return nil
}

func GetAddr() ([]AddrShowResult, error) {
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
