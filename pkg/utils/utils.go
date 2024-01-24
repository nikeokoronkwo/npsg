package utils

import (
	"flag"
	"github.com/spf13/pflag"
	"fmt"
	// "strconv"
	// "os"
)

const (
	DEF_DIR = "passwords"
	FILE_DIR = "passwords/passwords.txt"
)

/* Arg Types are:
    make
	save
	get
	restrict

	Arg Options are:
	-h, --help

save:
	-r, --root

save, get, make:
	-d, --directory

make:
	--uppercase
	--numbers
	--symbols
	-a, --all
	--size

make, save:
	-ref

*/
// Create Data Type to receive cmdline options
type CmdLineArgs struct {
	Directory string
	Permissions bool
	Help bool
	Uppercase bool
	Numbers bool
	Symbols bool
	Reference string
	Size int
	Args []string
}

// Prints Out Help Message

func HelpMessage() {
	fmt.Println("Usage: npsg [options] [arguments]")
	fmt.Println("Arguments: ")
	fmt.Println(" make: command to generate passwords and save on system (or in custom directory)\n", "save: save given passwords to system or to custom directory\n", "get: get passwords only from default system storage\n", "config: configure the storage file permissions")
	fmt.Println("Note that in order for you to be able to use config to change user permissions on storage to \"root\", you need to have root permissions (on bash/zsh, you use 'sudo')\nIf the passwords storage file is set to root permissions, then all commands related to the file must be run with root permissions (on bash/zsh, you use 'sudo')\n")
	fmt.Println("Options: ")
	pflag.PrintDefaults()
}


func SetFlags() *CmdLineArgs {
	var directory string
	pflag.StringVarP(&directory, "directory", "d", FILE_DIR , "Path to the passwords file")

	var permissions bool
	pflag.BoolVarP(&permissions, "root", "r", false , "Lock passwords to only root user (default - false)")

	var help bool
	pflag.BoolVarP(&help, "help", "h", false, "Show help Usage")

	var uppercase bool
	flag.BoolVar(&uppercase, "uppercase", false, "Include Uppercase Letters when making password")

	var numbers bool
	flag.BoolVar(&numbers, "numbers", false, "Include Numbers when making password")

	var symbols bool
	flag.BoolVar(&symbols, "symbols", false, "Include Symbols when making password")

	var all bool
	pflag.BoolVarP(&all, "all", "a", false, "Include Uppercase Letters, Numbers and Symbols when making password")

	var size int
	flag.IntVar(&size, "size", 0, "Size of the Password.")

	var ref string
	flag.StringVar(&ref, "ref", "", "Reference to the password: email address, username etc.")
	

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.CommandLine.AddFlagSet(pflag.CommandLine)

	pflag.Usage = HelpMessage
	flag.Usage = HelpMessage


	flag.Parse()
	pflag.Parse()

	if all {
		uppercase, numbers, symbols = true, true, true
	}

	cmdlineargs := &CmdLineArgs{
		Directory: directory,
		Permissions: permissions,
		Help: help,
		Uppercase: uppercase,
		Numbers: numbers,
		Symbols: symbols,
		Reference: ref,
		Size: size,
		Args: pflag.Args(),
	}

	return cmdlineargs
}