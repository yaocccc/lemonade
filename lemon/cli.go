package lemon

import (
	"io"
)

type CommandType int

// Commands
const (
	COPY CommandType = iota + 1
	PASTE
	SERVER
)

const (
	Success        = 0
	FlagParseError = iota + 10
	HTTPError
	Help
)

type CommandStyle int

const (
	ALIAS CommandStyle = iota + 1
	SUBCOMMAND
)

type CLI struct {
	In       io.Reader
	Out, Err io.Writer

	Type       CommandType
	DataSource string

	// options
	Port       int
	Host       string
	LineEnding string
	LogLevel   int

	Help bool
}
