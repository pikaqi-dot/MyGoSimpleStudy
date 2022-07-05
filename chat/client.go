package main

import (
	"bufio"
	"net"
)

type client struct {
	conn     net.Conn
	nick     string
	room     *room
	commands chan<- command
}

func (c *client) readInput() {

	for {
		bufio.NewReader(c.conn).ReadString('\n')

	}

}
