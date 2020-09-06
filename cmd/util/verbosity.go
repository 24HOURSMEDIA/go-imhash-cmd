package util

import (
	"github.com/spf13/cobra"
	"log"
)

const (
	Silent = iota
	Verbose
	MoreVerbose
	Debug
)

// Verbosity is helpful in determining verbosity levels
type Verbosity int

func (verbosity Verbosity) IsVerbose() bool {
	return verbosity > Silent
}
func (verbosity Verbosity) IsMoreVerbose() bool {
	return verbosity > Verbose
}
func (verbosity Verbosity) IsDebug() bool {
	return verbosity > MoreVerbose
}
func (verbosity Verbosity) Message(v ...interface{}) {
	if verbosity.IsVerbose() {
		log.Print(v...)
	}
}
func (verbosity Verbosity) Log(v ...interface{}) {
	if verbosity.IsMoreVerbose() {
		log.Print(v...)
	}
}
func (verbosity Verbosity) Debug(v ...interface{}) {
	if verbosity.IsDebug() {
		log.Print(v...)
	}
}
func (verbosity Verbosity) Fatal(v ...interface{}) {
	if verbosity.IsVerbose() {
		log.Fatal(v...)
	}
}

// GetVerbosity returns the verbosity of a cobra command
func GetVerbosity(cmd *cobra.Command) Verbosity {
	var level = Silent
	verbose, _ := cmd.Flags().GetBool("v")
	if verbose {
		level = Verbose
	}
	moreverbose, _ := cmd.Flags().GetBool("vv")
	if moreverbose {
		level = MoreVerbose
	}
	debug, _ := cmd.Flags().GetBool("vvv")
	if debug {
		level = Debug
	}
	return Verbosity(level)
}
