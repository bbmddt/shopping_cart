package csv_helper

import (
	"encoding/csv"
	"fmt"
	"log"
	"mime/multipart"
)

// 從給定的FileHeader讀訊息，返回二維數組
func ReadCsv(fileHeader *multipart.FileHeader) ([][]string, error) {
	f, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer func(f multipart.File) {
		err := f.Close()
		if err != nil {
			log.Print(err)
		}
	}(f)

	reader := csv.NewReader(f)
	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Println("err")
		return nil, err
	}
	var result [][]string

	for _, line := range lines[1:] {

		data := []string{line[0], line[1]}
		result = append(result, data)
	}

	return result, nil
}
