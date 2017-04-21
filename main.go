package main

import (
	"fmt"
	"log"
	"net"
	"syscall"

	"github.com/mistsys/tuntap"
	"github.com/vishvananda/netlink"
)

func Init() (*tuntap.Interface, error) {
	tun, err := tuntap.Open("tun%d", tuntap.DevTun)
	if err != nil {
		return nil, err
	}

	defer tun.Close()

	// it's fine if I just return here..
	tun0 := tun.Name()
	iface, err := netlink.LinkByName(tun0)
	if err != nil {
		return nil, err
	}

	// this is just a random IP address
	ipv6Address := "fd28:49be:758:7653:3cb1:6d:230c:60/64"
	ip, subnet, err := net.ParseCIDR(ipv6Address)

	bits, _ := subnet.Mask.Size()
	log.Printf("ip addr dev %s %s/%d", tun0, ip, bits)
	err = netlink.AddrAdd(iface, &netlink.Addr{
		IPNet: &net.IPNet{IP: ip, Mask: subnet.Mask},
		Scope: syscall.RT_SCOPE_UNIVERSE,
	})

	if err != nil {
		log.Printf("Failed to add %v to dev %q: %s", ip, tun0, err)
		return nil, err
	}

	return tun, nil
}

func main() {

	tun, err := Init()
	if err != nil {
		log.Printf("Error opening tun. %s", err)

	}

	// this is just so that we "use" the tun here
	fmt.Println(tun)
}
