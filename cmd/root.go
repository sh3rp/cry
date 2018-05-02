package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var inputFile string
var outputFile string
var key string

var RootCmd = &cobra.Command{
	Use:   "cry",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	RootCmd.AddCommand(encryptCmd)
	RootCmd.AddCommand(decryptCmd)
	RootCmd.PersistentFlags().StringVarP(&inputFile, "input", "i", "", "File to encrypt/decrypt; if blank, stdin will be used")
	RootCmd.PersistentFlags().StringVarP(&outputFile, "output", "o", "", "File to output results of encrypt/decrypt operation; if blank, stdout will be used")
	RootCmd.PersistentFlags().StringVarP(&key, "key", "k", "", "Key string to use for encryption/decryption; if blank, key will be prompted on command line")
}

func exit(err error) {
	fmt.Printf("Error: %v\n", err)
	os.Exit(1)
}
