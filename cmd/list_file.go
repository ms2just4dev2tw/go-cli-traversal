/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// listFileCmd represents the listFile command
var listFileCmd = &cobra.Command{
	Use:   "c",
	Short: "file traverse in current directory",
	Long: `file traverse in current directory
	and list file size with sort by high to low`,
	Run: func(cmd *cobra.Command, args []string) {
		printCurrentFile(args[0])
	},
}

func printCurrentFile(path string) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Print(err)
	}

	if fileInfo.IsDir() {
		files, err := os.ReadDir(path)
		if err != nil {
			fmt.Println(err)
		}

		for _, file := range files {
			fileInfo, _ = file.Info()
			if fileInfo.IsDir() {
				fmt.Println(fileInfo.Name())
			} else {
				fmt.Printf("%s:%0.2fkb\r\n", fileInfo.Name(), float32(fileInfo.Size())/1024)
			}
		}
	} else {
		fmt.Printf("%s:%0.2fkb\r\n", fileInfo.Name(), float32(fileInfo.Size())/1024)
	}

}

func init() {
	rootCmd.AddCommand(listFileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listFileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listFileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
