package main

import (
	"nuge/pswrd/pkg/lib"
	"nuge/pswrd/pkg/utils"
	"os"
	"fmt"
	"github.com/atotto/clipboard"
	"sync"
	"time"
	"slices"
	// "bufio"
)

func showLoadingAnimation(stopLoading chan struct{}, wg *sync.WaitGroup) {
	// Define a set of rotating dots
	dots := []string{".", "..", "...", "...."}

	defer wg.Done()

	// Infinite loop for rotating dots
	for {
		select {
		case <-stopLoading:
			// Stop the animation when a signal is received
			fmt.Println("\rDone Loading.")
			return
		default:
			for _, dot := range dots {
				fmt.Printf("\rLoading%s", dot)
				time.Sleep(300 * time.Millisecond)
			}
		}
	}
}

func init() {
	// var user lib.UserType
	// user = lib.normal
	if !lib.FileExists(utils.FILE_DIR) {
		lib.InitPasswordsFile(utils.FILE_DIR, lib.Normal)
		lib.CreateDirectory(utils.DEF_DIR)
	}
}


func main() {
	cliconfig := utils.SetFlags()
	if len(cliconfig.Args) > 1 && !slices.Contains(cliconfig.Args, "save") {
		fmt.Println("Only one argument is required (unless using save), see help usage")
		os.Exit(1)
	}
	args := os.Args
	if len(cliconfig.Args) == 0 {
		if cliconfig.Help || slices.Contains(args, "--help"){
			utils.HelpMessage()
			os.Exit(0)
		} else {
			utils.HelpMessage()
		}
		fmt.Println("At least one argument is required (unless using save, where two is needed), see help usage")
		os.Exit(1)
	} else {
		if cliconfig.Args[0] == "make" {
			if !cliconfig.Uppercase {
				var uppercase string
				fmt.Print("Do you want to include uppercase letters? (y/n) ")
				fmt.Scan(&uppercase)
				if string(uppercase[0]) == "y" || string(uppercase[0]) == "Y" {
					cliconfig.Uppercase = true
				}
				fmt.Println()
			}
			if !cliconfig.Numbers {
				var numbers string
				fmt.Print("Do you want to include numerical digits? (y/n) ")
				fmt.Scan(&numbers)
				if string(numbers[0]) == "y" || string(numbers[0]) == "Y" {
					cliconfig.Numbers = true
				}
				fmt.Println()
			}
			if !cliconfig.Symbols {
				var symbols string
				fmt.Print("Do you want to include basic symbols? (y/n) ")
				fmt.Scan(&symbols)
				if string(symbols[0]) == "y" || string(symbols[0]) == "Y" {
					cliconfig.Uppercase = true
				}
				fmt.Println()
			}
			if cliconfig.Directory == utils.FILE_DIR {
				var directory string
				var dir string
				fmt.Print("Do you want to save it in a custom directory? (y/n) ")
				fmt.Scan(&dir)
				if string(dir[0]) == "y" || string(dir[0]) == "Y" {
					fmt.Print("Type in the directory ")
					fmt.Scan(&directory)
					cliconfig.Directory = directory
				}
				fmt.Println()
			}
			if cliconfig.Size == 0 {
				var size int
				fmt.Print("Enter the size of your password ")
				fmt.Scan(&size)
				cliconfig.Size = size
				fmt.Println()
			}
			stopLoading := make(chan struct{})
			var wg sync.WaitGroup
			wg.Add(1)
			go showLoadingAnimation(stopLoading, &wg)
			password, err := lib.GeneratePassword(cliconfig.Uppercase, cliconfig.Numbers, cliconfig.Symbols, cliconfig.Size)
			if err != nil {
				fmt.Println("Error occured while generating password: %w", err)
				os.Exit(2)
			}
			close(stopLoading)
			wg.Wait()
			fmt.Println("Password generated: ", password)
			if cliconfig.Reference == "" {
				var ref string
				fmt.Print("Give an account reference (like the link to the site, or distinct username), if you want to save it (press 'n' if you do not want to save it): ")
				fmt.Scanln(&ref)
				fmt.Println()
				if cliconfig.Directory == utils.FILE_DIR && ref != "n" {
					err := lib.SavePassword(password, ref)
					if err != nil {
						fmt.Println("\n", err)
						os.Exit(2)
					}
				} else if ref != "n" {
					err := lib.SavePasswordCustom(password, ref, cliconfig.Directory)
					if err != nil {
						fmt.Println("\n", err)
						os.Exit(2)
					}
				}
			} else {
				if cliconfig.Directory != "" {
					err := lib.SavePassword(password, cliconfig.Reference)
					if err != nil {
						fmt.Println("\n", err)
						os.Exit(2)
					}
				} else {
					err := lib.SavePasswordCustom(password, cliconfig.Reference, cliconfig.Directory)
					if err != nil {
						fmt.Println("\n", err)
						os.Exit(2)
					}
				}
			}
			var indicateClip string
			fmt.Print("Do you want it copied to the clipboard? (y/n) ")
			fmt.Scan(&indicateClip)
			if string(indicateClip[0]) == "y" || string(indicateClip[0]) == "Y" {
				// Copy text to clipboard
				err := clipboard.WriteAll(password)
				if err != nil {
					fmt.Println("Error copying to clipboard:", err)
					return
				}

				fmt.Printf("Text '%s' copied to clipboard successfully.\n", password)
			}
			fmt.Println("Thank you!")
			os.Exit(0)
		}
		if cliconfig.Args[0] == "save" {
			if len(cliconfig.Args) < 2 {
				fmt.Println("You need two arguments for this case")
				os.Exit(1)
			}
			password := cliconfig.Args[1]
			var ref string
			var directory string
			var dir string
			if cliconfig.Reference == "" {
				fmt.Print("Give an account reference (like the link to the site, or distinct username), if you want to save it: ")
				fmt.Scan(&ref)
				cliconfig.Reference = ref
				fmt.Println()
			}
			if cliconfig.Directory == utils.FILE_DIR {
				fmt.Print("Do you want to save it in a custom directory? (y/n) ")
				fmt.Scan(&dir)
				if string(dir[0]) == "y" || string(dir[0]) == "Y" {
					fmt.Print("Type in the directory ")
					fmt.Scan(&directory)
					cliconfig.Directory = directory
				}
				fmt.Println()
			}
			stopLoading := make(chan struct{})
			var wg sync.WaitGroup
			wg.Add(1)
			go showLoadingAnimation(stopLoading, &wg)

			fmt.Println()
			if cliconfig.Directory == utils.FILE_DIR {
				err := lib.SavePassword(password, cliconfig.Reference)
				if err != nil {
					fmt.Println(err)
					os.Exit(2)
				}
			} else {
				err := lib.SavePasswordCustom(password, cliconfig.Reference, cliconfig.Directory)
				if err != nil {
					fmt.Println(err)
					os.Exit(2)
				}
			}
			close(stopLoading)
			wg.Wait()
			fmt.Println("Password saved!")
		}
		if cliconfig.Args[0] == "get" {
			var pswrd string
			fmt.Println("Only use this if you are getting passwords from default storage")
			var ref string
			if cliconfig.Reference == "" {
				fmt.Print("Give an account reference (like the link to the site, or distinct username), if you want to save it: ")
				fmt.Scan(&ref)
				fmt.Println()
				cliconfig.Reference = ref
				if cliconfig.Reference == "" || ref == "" {
					fmt.Println("Error: You need a reference in order to search for a password")
				}
			} 
			stopLoading := make(chan struct{})
			var wg sync.WaitGroup
			wg.Add(1)
			go showLoadingAnimation(stopLoading, &wg)
			fmt.Println()
			if cliconfig.Reference != "" {
				password, err := lib.GetPassword(cliconfig.Reference, utils.FILE_DIR)
				if err != nil {
					fmt.Println()
					fmt.Println(err)
					os.Exit(2)
				}
				pswrd = password
			}
			close(stopLoading)
			wg.Wait()
			fmt.Printf("Here is the password found for ref %s : %s\n", cliconfig.Reference, pswrd)
			fmt.Println("Thank you!")
		}
		if cliconfig.Args[0] == "config" {
			var root string
			if !cliconfig.Permissions {
				fmt.Print("What kind of permissions do you want to put on the password storage [(r)oot/(a)ll/(u)ser] ")
				fmt.Scan(&root)
				if string(root[0]) == "r" || string(root[0]) == "R" {
					cliconfig.Permissions = true
				}
				fmt.Println()
			}
			stopLoading := make(chan struct{})
			var wg sync.WaitGroup
			wg.Add(1)
			go showLoadingAnimation(stopLoading, &wg)
			fmt.Println()
			if cliconfig.Permissions {
				err := lib.RootFilePermissions(utils.FILE_DIR)
				if err != nil {
					fmt.Printf("error adding permissions to file : %w", err)
					os.Exit(2)
				}

			} else if string(root[0]) == "u" || string(root[0]) == "U" {
				err := lib.RestrictFilePermissions(utils.FILE_DIR)
				if err != nil {
					fmt.Printf("error adding permissions to file : %w", err)
					os.Exit(2)
				}
			} else {
				err := lib.AllFilePermissions(utils.FILE_DIR)
				if err != nil {
					fmt.Printf("error adding permissions to file : %w", err)
					os.Exit(2)
				}
			}
			close(stopLoading)
			wg.Wait()
			fmt.Printf("Permissions placed!\n")
			fmt.Println("Thank you!")
		}
	}
}