package memory

import (
	"bufio"
	"os"
)

// remember saves a key to a text file if it isn't already present
func Remember(key string) error {
	file, err := os.OpenFile("memory.txt", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == key {
			// Key already remembered
			return nil
		}
	}

	// Write the key to the file
	if _, err := file.WriteString(key + "\n"); err != nil {
		return err
	}
	return nil
}

// check verifies if a key is present in the memory file
func Check(key string) (bool, error) {
	file, err := os.Open("memory.txt")
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil // File doesn't exist, key not remembered
		}
		return false, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == key {
			return true, nil
		}
	}

	return false, nil
}
