package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/masayoshi634/ipgo/pkg/command"
	"github.com/pkg/errors"
)

var (
	Version  string
	Revision string
)

var helpText = `
Usage: ipgo [options]
  ip command for json
Commands:
  addr, a          like ip addr
  link, l          like ip link
Options:
  --help, -h       print help
`

type CLI struct {
	// outStream and errStream are the stdout and stderr
	// to write message from the CLI.
	outStream, errStream io.Writer
}

func (cli *CLI) prepareFlags(help string) *flag.FlagSet {
	Name := "ipgo"
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(cli.errStream)
	flags.Usage = func() {
		fmt.Fprint(cli.errStream, help)
	}
	return flags
}

func (cli *CLI) Run(args []string) int {
	if len(args) <= 1 {
		fmt.Fprint(cli.errStream, helpText)
		return 2
	}

	var err error

	switch args[1] {
	case "addr", "a":
		err = cli.doAddr(args[2:])
	case "link", "l":
		err = cli.doLink(args[2:])
	case "-h", "--help":
		fmt.Fprint(cli.errStream, helpText)
	default:
		fmt.Fprint(cli.errStream, helpText)
		return 1
	}

	if err != nil {
		fmt.Fprintln(cli.errStream, err)
		return 2
	}

	return 0
}

var addrHelpText = `
Usage: ipgo addr [options]
  ip command for json
Commands:
  show,s                          like ip addr show
  add,a   [ipaddr] [interface]    ipgo addr add 127.0.0.2/32 lo
  del,d   [ipaddr] [interface]    ipgo addr del 127.0.0.2/32 lo
  replace,r [ipaddr] [interface]  ipgo addr del 127.0.0.2/32 lo
Options:
`

func (cli *CLI) doAddr(args []string) error {
	var err error
	var ipnet, iface string

	if len(args) < 1 {
		return err
	}

	switch args[0] {
	case "show", "s":
		return command.AddrShow()
	case "add", "a":
		if len(args) < 3 {
			fmt.Fprint(cli.errStream, addrHelpText)
			return errors.Errorf("ipaddr and interface required")
		}
		ipnet = args[1]
		iface = args[2]
		return command.AddrAdd(ipnet, iface)
	case "del", "d":
		if len(args) < 3 {
			fmt.Fprint(cli.errStream, addrHelpText)
			return errors.Errorf("ipaddr and interface required")
		}
		ipnet = args[1]
		iface = args[2]
		return command.AddrDelete(ipnet, iface)
	case "replace", "r":
		if len(args) < 3 {
			fmt.Fprint(cli.errStream, addrHelpText)
			return errors.Errorf("ipaddr and interface required")
		}
		ipnet = args[1]
		iface = args[2]
		return command.AddrReplace(ipnet, iface)
	default:
		return err
	}

	if err != nil {
		return err
	}
	return nil
}

func (cli *CLI) doLink(args []string) error {
	var err error
	if len(args) < 1 {
		return err
	}

	switch args[0] {
	case "show", "s":
		err = command.LinkShow()
	default:
		return err
	}

	if err != nil {
		return err
	}
	return nil
}

func main() {
	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}
