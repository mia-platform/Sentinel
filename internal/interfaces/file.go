package interfaces

import (
	"encoding/json"
	"fmt"
	"os"
)

// fileWriter writes the event to a file in a specific path
func WriteToFile(path string, event event) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			fmt.Printf("Error creating file: %v\n", err)
			return err
		}
		file.Close() // Chiude il file appena creato
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return err
	}
	defer file.Close()

	jsonData, err := json.Marshal(event)
	if err != nil {
		fmt.Printf("Error marshalling data: %v\n", err)
		return err
	}

	jsonData = append(jsonData, '\n')

	// Scrive i dati nel file
	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Printf("Error writing data: %v\n", err)
		return err
	}

	return nil
}
