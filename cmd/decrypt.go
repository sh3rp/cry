package cmd

import (
	"github.com/sh3rp/cry"
	"github.com/sh3rp/crypt"
	"github.com/spf13/cobra"
)

var decryptCmd = &cobra.Command{
	Use:   "d",
	Short: "decrypt a file",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		key, err := cry.GetKey(key)

		if err != nil {
			exit(err)
		}

		inputData, err := cry.ReadInput(inputFile, args)

		if err != nil {
			exit(err)
		}

		iv := inputData[:16]
		encrypted := inputData[16:]

		data, err := crypt.Decrypt(key, iv, encrypted)

		if err != nil {
			exit(err)
		}

		err = cry.WriteOutput(data, outputFile)

		if err != nil {
			exit(err)
		}
	},
}
