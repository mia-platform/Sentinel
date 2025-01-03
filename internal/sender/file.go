package sender

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/mia-platform/sentinel/pkg/collector"
)

// fileWriter writes the event to a file in a specific path
func WriteToFile(path string, event collector.Collector) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		_, err := os.Create(path)
		if err != nil {
			fmt.Printf("Error creating file: %v\n", err)
			return err
		}
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonData, err := json.Marshal(event)
	if err != nil {
		fmt.Printf("Error marshalling data: %v\n", err)
	}

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Printf("Error writing data: %v\n", err)
	}

	return nil
}
