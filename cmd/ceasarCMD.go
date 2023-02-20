/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// barkCmd represents the bark command
var encryptCMD = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypts your vlue",
	Long:  `encrypts your shit.`,
	Run: func(cmd *cobra.Command, args []string) {
		cipher := CaesarCipher{shift: 3}
		var item = cipher.Encrypt(args[0])
		println("Encrypted  " + item)
		var test = cipher.Decrypt(item)
		println("Decrypted  " + test)

	},
}

func init() {
	rootCmd.AddCommand(encryptCMD)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// barkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// barkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
