// password_test.go

package lib

import (
	"os"
	"path/filepath"
	"testing"
	"nuge/pswrd/pkg/utils"
)

func TestGeneratePassword(t *testing.T) {
	// Test case 1: Positive scenario
	password, err := GeneratePassword(true, true, true, 12)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(password) != 12 {
		t.Errorf("Expected password length to be 12, but got %d", len(password))
	}

	// Test case 2: Negative scenario (invalid size)
	_, err = GeneratePassword(true, true, true, -1)
	if err == nil || err.Error() != "size must be a positive number" {
		t.Errorf("Expected error 'size must be a positive number', but got: %v", err)
	}
}

func TestInitPasswordsFile(t *testing.T) {
	// Test case 1: Positive scenario
	filePath := "testfile.txt"
	defer os.Remove(filePath)

	err := InitPasswordsFile(filePath, Root)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test case 2: Negative scenario (file creation failure)
	err = InitPasswordsFile("/nonexistentpath/testfile.txt", Root)
	if err == nil {
		t.Error("Expected an error, but got none")
	}
}

func TestSavePassword(t *testing.T) {
	// Test case 1: Positive scenario
	err := SavePassword("mypassword", "myreference")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test case 2: Negative scenario (file write failure)
	// err = SavePassword("mypassword", "myreference")
	// if err == nil {
	// 	t.Error("Expected an error, but got none")
	// }
}

func TestSavePasswordCustom(t *testing.T) {
	// Test case 1: Positive scenario
	location := "testfolder"
	defer os.RemoveAll(location)

	err := SavePasswordCustom("mypassword", "myreference", location)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test case 2: Negative scenario (file write failure)
	err = SavePasswordCustom("mypassword", "myreference", "/nonexistentpath")
	if err == nil {
		t.Error("Expected an error, but got none")
	}
}

func TestGetPassword(t *testing.T) {
	// Test case 1: Positive scenario
	location := "testfolder"
	defer os.RemoveAll(location)

	err := SavePasswordCustom("mypassword", "myreference", location)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	password, err := GetPassword("myreference", location)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if password != "mypassword" {
		t.Errorf("Expected password 'mypassword', but got '%s'", password)
	}

	// Test case 2: Negative scenario (reference not found)
	_, err = GetPassword("nonexistentreference", location)
	if err == nil || err.Error() != "reference doesn't exist" {
		t.Errorf("Expected error 'reference doesn't exist', but got: %v", err)
	}
}
