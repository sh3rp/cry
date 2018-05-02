package cry

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/howeyc/gopass"
)

var ERR_INVALID_PASSWORD = errors.New("Invalid password")
var ERR_PASSWORD_MISMATCH = errors.New("Passwords do not match")

func GetKey(key string) (string, error) {
	var encryptionKey string
	if key == "" {
		fmt.Print("Encryption key: ")
		pass, err := gopass.GetPasswd()

		if err != nil {
			return "", ERR_INVALID_PASSWORD
		}

		fmt.Print("Confirm encryption key: ")
		pass2, err := gopass.GetPasswd()

		if err != nil {
			return "", ERR_INVALID_PASSWORD
		}

		if !bytes.Equal(pass, pass2) {
			return "", ERR_PASSWORD_MISMATCH
		}

		encryptionKey = string(pass)
	} else {
		encryptionKey = key
	}
	return encryptionKey, nil
}

func ReadInput(filename string, commandLine []string) ([]byte, error) {
	var bytes []byte
	if len(commandLine) > 0 && commandLine[0] != "" {
		return []byte(commandLine[0]), nil
	} else if filename == "" {
		reader := bufio.NewReader(os.Stdin)
		for {
			buf := make([]byte, 4096)
			numRead, err := reader.Read(buf)
			if err == io.EOF {
				break
			}
			bytes = append(bytes, buf[:numRead]...)
		}
	} else {
		file, err := os.Open(filename)

		if err != nil {
			return nil, err
		}

		defer file.Close()

		bytes, err = ioutil.ReadAll(file)

		if err != nil {
			return nil, err
		}
	}
	return bytes, nil
}

func WriteOutput(data []byte, filename string) error {
	var writer io.Writer
	if filename == "" {
		writer = os.Stdout
	} else {
		writer, _ = os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	}
	_, err := writer.Write(data)
	return err
}
