package processes

import (
	"encoding/csv"
	"fmt"
	"os"
)

func OpenAndReadCSV(fileName string) (*csv.Reader, *os.File, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening file: %v", err)
	}

	reader := csv.NewReader(file)
	_, err = reader.Read() // Read and discard the header row
	if err != nil {
		return nil, nil, fmt.Errorf("error reading header row: %v", err)
	}

	return reader, file, nil
}
