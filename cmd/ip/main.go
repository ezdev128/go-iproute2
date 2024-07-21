package main

import (
	"fmt"

	"github.com/mdlayher/netlink"
	"github.com/spf13/cobra"

	"github.com/ezdev128/go-iproute2"
)

var cli client

type client struct {
	conn *netlink.Conn
}

func (c *client) dialNetlink() error {
	var err error
	c.conn, err = netlink.Dial(iproute2.NETLINK_ROUTE, nil)
	return err
}

func (c *client) runCmd(fn func()) {
	err := c.dialNetlink()
	if err != nil {
		fmt.Println("failed to create netlink socket, err:", err)
		return
	}
	defer c.conn.Close()

	fn()
}

var rootCmd = cobra.Command{
	Use: "ip",
}

func main() {
	rootCmd.Execute()
}
