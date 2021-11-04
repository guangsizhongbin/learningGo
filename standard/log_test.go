package standard

import (
	"log"
	"os"
	"testing"
)

// 最基本的log
func TestBasicLog(t *testing.T) {
	log.Println("Hello world!")
}

// login to a file
func TestLoginToAFile(t *testing.T) {
	// os.O_APPEND --> append data to the file when writing
	// os.O_CREATE --> create a nw file if none exists
	// os.O_WRONLY --> open the file write-only
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	// SetOutput sets the output destination for the standard logger.
	log.SetOutput(file)

	log.Println("Hello world")
}
