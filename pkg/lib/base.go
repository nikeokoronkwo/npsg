package lib

import (
	"os"
	"path/filepath"
	"os/user"
	"fmt"
	"io/ioutil"
)

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func CreateDirectory(directory string) error {
	// Create the directory if it doesn't exist
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		if err := os.Mkdir(directory, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create directory: %v", err)
		}
	}
	return nil
}

func storeData(filePath string, data string) error {
	// Write data to the specified file
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(data);
	if err != nil {
		return fmt.Errorf("failed to write data to file: %v", err)
	}

	return nil
}

func RestrictFilePermissions(filePath string) error {
	// Get information about the current user
	currentUser, err := user.Current()
	if err != nil {
		return fmt.Errorf("failed to get current user: %v", err)
	}

	// Set restrictive permissions for the file (read, write, execute only for the owner)
	err = os.Chmod(filePath, 0700)
	if err != nil {
		return fmt.Errorf("failed to set file permissions: %v", err)
	}

	fmt.Printf("\nFile permissions restricted to the owner (%s) only.\n", currentUser.Username)
	return nil
}

func RootFilePermissions(filePath string) error {
	// Set root as the owner (uid 0) and give read, write, and execute permissions to the owner
	err := os.Chown(filePath, 0, os.Getegid())
	if err != nil {
		return fmt.Errorf("error changing file owner: %w", err)
	}

	// Set read, write, and execute permissions for the owner
	err = os.Chmod(filePath, 0700)
	if err != nil {
		return fmt.Errorf("error changing file permissions: %w", err)
	}

	fmt.Println("File permissions changed to root successfully.")
	return nil
}

func AllFilePermissions(filePath string) error {
	err := os.Chmod(filePath, 0777)
	if err != nil {
		return fmt.Errorf("error changing file permissions: %w", err)
	}

	fmt.Println("File permissions changed to allow all users.")
	return nil
}

func readData(customDir string) (string, error) {
	// Get the user's home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	// Determine the directory to read data from
	var targetDir string
	if customDir != "" {
		targetDir = customDir
	} else {
		targetDir = filepath.Join(homeDir, "mydata")
	}

	// Check if the file exists in the target directory
	dataFilePath := targetDir
	if _, err := os.Stat(dataFilePath); os.IsNotExist(err) {
		return "", fmt.Errorf("data file does not exist in the target directory")
	}

	// Read data from the file
	data, err := ioutil.ReadFile(dataFilePath)
	if err != nil {
		return "", err
	}

	return string(data), nil
}