package utils_test

import (
	"testing"
	"github.com/spf13/pflag"
	"github.com/spf13/cobra"
)

func TestSetFlags(t *testing.T) {
	// Mock command line arguments
	args := []string{"arg1", "arg2"}

	// Create a new Cobra command
	cmd := &cobra.Command{}
	cmd.SetArgs(args)

	// Set up flags
	cmd.Flags().StringP("directory", "d", "", "Path to the passwords file")
	cmd.Flags().BoolP("root", "r", false, "Lock passwords to only root user (default - false)")
	cmd.Flags().BoolP("help", "h", false, "Show help Usage")
	cmd.Flags().Bool("uppercase", false, "Include Uppercase Letters when making password")
	cmd.Flags().Bool("numbers", false, "Include Numbers when making password")
	cmd.Flags().Bool("symbols", false, "Include Symbols when making password")
	cmd.Flags().BoolP("all", "a", false, "Include Uppercase Letters, Numbers, and Symbols when making password")
	cmd.Flags().Int("size", 0, "Size of the Password.")
	cmd.Flags().String("ref", "", "Reference to the password: email address, username etc.")

	// Set up pflag
	pflag.CommandLine = cmd.Flags()

	// Call the function to be tested
	cmdLineArgs := SetFlags()

	// Add assertions based on expected behavior
	if cmdLineArgs.Directory != "" {
		t.Errorf("Expected directory to be empty, got %s", cmdLineArgs.Directory)
	}

	if cmdLineArgs.Permissions != false {
		t.Errorf("Expected permissions to be false, got %t", cmdLineArgs.Permissions)
	}

	// Add more assertions for other flags

	// Add assertions for non-flag values
	if cmdLineArgs.Args[0] != "arg1" || cmdLineArgs.Args[1] != "arg2" {
		t.Errorf("Expected Args to be [arg1 arg2], got %v", cmdLineArgs.Args)
	}
}
