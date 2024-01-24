package lib

import (
	"math/rand"
	"time"
	"errors"
	"fmt"
	"os"
	"strings"
	"path/filepath"
	"nuge/pswrd/pkg/utils"
)

type UserType int
const (
    Normal UserType = iota
    Root
)

// Declare constants to base password on
const (
	LOWERCASE = "qwertyuiopasdfghjklzxcvbnm"
	UPPERCASE = "QWERTYUIOPASDFGHJKLZXCVBNM"
	NUMBERS = "1234567890"
	SYMBOLS = "`~!@#$%^&*()-_=+[{}]\\|;:'\",<.>/?"
)


// Function to generate password for user
func GeneratePassword(uppercase bool, numbers bool, symbols bool, size int) (string, error) {
	// Define essential variables: 'pool' contains the pool for password letters, and 'password' is the output password
	var password string
	var pool string

	// Initialise with the 'LOWERCASE' set of letters, and add others in the case they are needed
	pool = LOWERCASE
	if uppercase {pool = pool + UPPERCASE}
	if numbers {pool = pool + NUMBERS}
	if symbols {pool = pool + SYMBOLS}

	// Random Number Generator for producing random indexes
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s) 

	// Ensure the size parsed is positive
	if size <= 0 {
		return "", errors.New("size must be a positive number")
	}
	// Add individual letters to the 'password' variable until the size is reached
	for i := 0; i < size; i++ {
		index := r.Intn(len(pool))
		password = password + string(pool[index])
	}
	return password, nil
}

// Before we can begin to save passwords, let's initalise the file that will store these passwords
func InitPasswordsFile(filePath string, permission UserType) error {
	// Create the file
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	// Add permissions to the file if needed by the user
	if permission == Root {
		err = RestrictFilePermissions(filePath)
		if err != nil {
			return fmt.Errorf("error adding permissions to file : %w", err)
		}
	}
	return nil

}

// Save password in document for usage later
func SavePassword(password, reference string) error {
	// Produce data in suitable format to be accessed later
	writedata := fmt.Sprintf("%s: %s\n", reference, password)

	// Store data in file
	err := storeData(utils.FILE_DIR, writedata)
	if err != nil {
		return fmt.Errorf("error writing data : %w", err)
	}
	return nil
}

func SavePasswordCustom(password, reference, location string) error {
	// Produce data in suitable format to be accessed later
	writedata := fmt.Sprintf("%s: %s\n", reference, password)

	// Store data in file in given location
	filePath := filepath.Join(location, "passwords.txt")
	err := storeData(filePath, writedata)
	if err != nil {
		return fmt.Errorf("error writing data : %w", err)
	}
	return nil
}

func GetPassword(reference string, location string) (string, error) {
	data, err := readData(location)
	if err != nil {
		return "", err
	}
	passwordsList := strings.Split(data, "\n")
	for _, v := range passwordsList {
		passwordsRef := strings.Split(v, ": ")
		if passwordsRef[0] == reference && len(passwordsRef) > 1 {
			return passwordsRef[1], nil
		}
	}
	return "", errors.New("reference doesn't exist")
}

