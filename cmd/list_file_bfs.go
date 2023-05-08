/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"container/list"
	"fmt"
	"io/fs"
	"os"

	"github.com/spf13/cobra"
)

// listFileBfsCmd represents the listFileBfs command
var listFileBfsCmd = &cobra.Command{
	Use:   "bfs",
	Short: "file traverse in current directory with bfs",
	Long: `file traverse in current directory wirh bfs
	and list file size with sort by high to low`,
	Run: func(cmd *cobra.Command, args []string) {
		printAllFileWithBfs(args[0])
	},
}

type DirStruct struct {
	fileInfo fs.FileInfo
	absPath  string
}

func printAllFileWithBfs(path string) {
	// 用队列存储文件节点
	queue := list.New()

	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Print(err)
	}

	if fileInfo.IsDir() {
		ds := DirStruct{fileInfo, path}
		queue.PushBack(ds)
		for queue.Len() != 0 {
			ds = queue.Front().Value.(DirStruct)
			queue.Remove(queue.Front())

			path = ds.absPath
			files, err := os.ReadDir(path)
			if err != nil {
				fmt.Println(err)
			}

			for _, file := range files {
				fileInfo, _ = file.Info()
				if fileInfo.IsDir() {
					newDs := DirStruct{fileInfo, path + string(os.PathSeparator) + fileInfo.Name()}
					queue.PushBack(newDs)
				} else {
					fmt.Printf("%s%s%s:%0.2fkb\r\n", path, string(os.PathSeparator), fileInfo.Name(), float32(fileInfo.Size())/1024)
				}
			}
		}
	} else {
		fmt.Printf("%s%s%s:%0.2fkb\r\n", path, string(os.PathSeparator), fileInfo.Name(), float32(fileInfo.Size())/1024)
	}

}

func init() {
	rootCmd.AddCommand(listFileBfsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listFileBfsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listFileBfsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
