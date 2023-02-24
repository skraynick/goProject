/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	Frequency "goProject/utils"

	"github.com/spf13/cobra"
)

// Encrypt your data with ceasar
var frequencyBreakCMD = &cobra.Command{
	Use:   "frequency",
	Short: "Frequency Analysis",
	Long:  `Frequency Analysis.`,
	Run: func(cmd *cobra.Command, args []string) {
		//println("Crypto Text:   " + args[0])
		//println("Decrypted:  " + Frequency.FrequencyAnalysisDecrypt(args[0]))
		println("Decrypted:  " + Frequency.DecryptWithFrequencyAnalysis(args[0]))

	},
}

func init() {
	rootCmd.AddCommand(frequencyBreakCMD)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// barkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// barkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
