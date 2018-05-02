package cmd

import (
	"time"

	"github.com/sh3rp/cry"
	"github.com/sh3rp/crypt"
	"github.com/spf13/cobra"
)

var encryptCmd = &cobra.Command{
	Use:   "e",
	Short: "encrypt a file",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {

		key, err := cry.GetKey(key)

		if err != nil {
			exit(err)
		}

		iv, err := crypt.GenerateCommonIV(time.Now().UnixNano())

		if err != nil {
			exit(err)
		}

		inputData, err := cry.ReadInput(inputFile, args)

		if err != nil {
			exit(err)
		}

		var encryptedBytes []byte

		data, err := crypt.Encrypt(key, iv, inputData)

		if err != nil {
			exit(err)
		}

		encryptedBytes = append(encryptedBytes, iv...)
		encryptedBytes = append(encryptedBytes, data...)

		err = cry.WriteOutput(encryptedBytes, outputFile)

		if err != nil {
			exit(err)
		}

	},
}
