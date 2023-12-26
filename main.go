package main

import (
	"encoding/csv"
	"encoding/json"
	"net/http"
	"os"
	"path"
)

type Column struct {
	Heading string `json:"heading"`
}

func main() {
	// open the file
	file, err := os.Open("data.csv")
	defer file.Close()

	if err != nil {
		panic(err)
	}

	file_ext := path.Ext("data.csv")
	if file_ext == "" {
		panic("filetype not detected: Please specify .json or .csv in the filename")
	}

	if file_ext == ".csv" {
		toJSON(file)
	} else if file_ext == ".json" {
		toCSV(file)
	} else {
		panic("invalid filetype: Please supply a .json or .csv file")
	}

	// ensure file is closed at the end

}
func toCSV(file *os.File) {
	// d := 1
}

func toJSON(file *os.File) {
	// Read the CSV data
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // Allow variable number of fields
	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	headers := data[0] // heading of each column in the csv

	columnData := make(map[string][]string) // map to store column data - string : []string

	// initialise map with empty list for each header
	for _, header := range headers {
		columnData[header] = []string{} // empty slice
	}

	// iterate through rows from second row
	for i, row := range data {
		if i == 0 {
			continue
		}

		for j, col := range row {
			columnData[headers[j]] = append(columnData[headers[j]], col) // fill the map
		}
	}

	// convert struct to JSON
	jsonData, err := json.Marshal(columnData)
	if err != nil {
		panic(err)
	}

	// write to json file
	err = os.WriteFile("output.json", jsonData, 0644)
	if err != nil {
		panic(err)
	}
}

func GetFileContentType(ouput *os.File) (string, error) {
	// get content type from first 512 bytes

	buf := make([]byte, 512)

	_, err := ouput.Read(buf)

	if err != nil {
		return "", err
	}

	// the function that actually does the trick
	contentType := http.DetectContentType(buf)

	return contentType, nil
}
