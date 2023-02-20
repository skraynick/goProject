/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	ceasar "goProject/utils"

	"github.com/spf13/cobra"
)

// Encrypt your data with ceasar
var encryptCMD = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypts your vlue",
	Long:  `encrypts your shit.`,
	Run: func(cmd *cobra.Command, args []string) {
		var item = ceasar.Encrypt(args[0], 3)
		println("Encrypted  " + item)

	},
}

var dencryptCMD = &cobra.Command{
	Use:   "decrypt",
	Short: "Dencrypts your vlue",
	Long:  `Dencrypts your shit.`,
	Run: func(cmd *cobra.Command, args []string) {
		var test = ceasar.Decrypt(args[0], 3)
		println("Decrypted  " + test)

	},
}

func init() {
	rootCmd.AddCommand(encryptCMD)
	rootCmd.AddCommand(dencryptCMD)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// barkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// barkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
