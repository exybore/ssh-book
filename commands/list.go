package commands

import (
	"github.com/exybore/ssh-book/book"
	"github.com/exybore/ssh-book/connections"
	"github.com/urfave/cli"
	"log"
)

func ListConnections(c *cli.Context) error {
	connectionsList, err := book.Open()
	if err != nil {
		log.Fatal("error while opening the book: ", err)
		return err
	}

	println("List of all SSH connections :")
	for _, connection := range connections.ListConnections(connectionsList) {
		println(connection)
	}

	return nil
}
