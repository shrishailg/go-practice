package files

import (
	"bufio"
	"log"
	"os"

	logger "github.com/sirupsen/logrus"
)

type FileOperations struct {
}

func NewFileOperations() *FileOperations {
	return &FileOperations{}
}

func (f *FileOperations) ReadFromFile(fileName string) (string, error) {
	file, err := os.Open("abafdjn")

	if err != nil {
		logger.WithField("message", err).Error("Error occured while reading file")
		return "empty data", err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	data := ""

	for scanner.Scan() {
		data = data + " " + scanner.Text()
	}

	return string(data), nil

}

func (f *FileOperations) WriteToFile(fileName string) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		logger.WithField("messafe", err.Error()).Error("Error occured while opening file")
		return
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, line := range []string{
		"This is line 1",
		"This is line 2",
		"This is line 3",
	} {
		_, err := writer.WriteString(line + "\n")

		if err != nil {
			logger.WithFields(logger.Fields{
				"Message": "Error occured while writing data to file",
				"Error":   true,
			}).Error(err.Error())
		}
	}

	err = writer.Flush()
	if err != nil {
		log.Fatal("Error occured while flushing writer " + err.Error())
	}
}
