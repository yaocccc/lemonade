package lemon

import (
	"flag"
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/mitchellh/go-homedir"
	"github.com/monochromegane/conflag"
)

func (c *CLI) FlagParse(args []string, skip bool) error {
	style, err := c.getCommandType(args)
	if err != nil {
		return err
	}
	if style == SUBCOMMAND {
		args = args[:len(args)-1]
	}

	return c.parse(args, skip)
}

func (c *CLI) getCommandType(args []string) (s CommandStyle, err error) {
	s = ALIAS
	switch {
	case regexp.MustCompile(`/?pbpaste$`).MatchString(args[0]):
		c.Type = PASTE
		return
	case regexp.MustCompile(`/?pbcopy$`).MatchString(args[0]):
		c.Type = COPY
		return
	}

	del := func(i int) {
		copy(args[i+1:], args[i+2:])
		args[len(args)-1] = ""
	}

	s = SUBCOMMAND
	for i, v := range args[1:] {
		switch v {
		case "paste":
			c.Type = PASTE
			del(i)
			return
		case "copy":
			c.Type = COPY
			del(i)
			return
		case "server":
			c.Type = SERVER
			del(i)
			return
		}
	}

	return s, fmt.Errorf("Unknown SubCommand\n\n" + Usage)
}

func (c *CLI) flags() *flag.FlagSet {
	flags := flag.NewFlagSet("lemonade", flag.ContinueOnError)
	flags.IntVar(&c.Port, "port", 2489, "TCP port number")
	flags.StringVar(&c.Host, "host", "localhost", "Destination host name.")
	flags.BoolVar(&c.Help, "help", false, "Show this message")
	flags.StringVar(&c.LineEnding, "line-ending", "CR", "Convert Line Endings (CR/CRLF)")
	flags.IntVar(&c.LogLevel, "log-level", 1, "Log level")
	return flags
}

func (c *CLI) parse(args []string, skip bool) error {
	flags := c.flags()

	confPath, err := homedir.Expand("~/.config/lemonade.toml")
	if err == nil && !skip {
		if confArgs, err := conflag.ArgsFrom(confPath); err == nil {
			flags.Parse(confArgs)
		}
	}

	var arg string
	err = flags.Parse(args[1:])
	if err != nil {
		return err
	}
	if c.Type == PASTE || c.Type == SERVER {
		return nil
	}

	for 0 < flags.NArg() {
		arg = flags.Arg(0)
		err := flags.Parse(flags.Args()[1:])
		if err != nil {
			return err
		}

	}

	if c.Help {
		return nil
	}

	if arg != "" {
		c.DataSource = arg
	} else {
		b, err := ioutil.ReadAll(c.In)
		if err != nil {
			return err
		}
		c.DataSource = string(b)
	}

	return nil
}
