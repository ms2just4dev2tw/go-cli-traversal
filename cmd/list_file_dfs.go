/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"container/list"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// listFileDfsCmd represents the listFileDfs command
var listFileDfsCmd = &cobra.Command{
	Use:   "dfs",
	Short: "traverse file in current directory with dfs",
	Long: `traverse file in current directory with dfs
	and list file size with sort by high to low`,
	Run: func(cmd *cobra.Command, args []string) {
		printAllFileWithDfs(args[0])

		fmt.Println("")

		printAllFileWithWalkDir(args[0])
	},
}

func printAllFileWithDfs(path string) {
	// 用栈存储文件节点
	stack := list.New()

	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Print(err)
	}

	if fileInfo.IsDir() {
		ds := DirStruct{fileInfo, path}
		stack.PushFront(ds)
		for stack.Len() != 0 {
			ds = stack.Front().Value.(DirStruct)
			stack.Remove(stack.Front())

			path = ds.absPath
			files, err := os.ReadDir(path)
			if err != nil {
				fmt.Println(err)
			}

			for _, file := range files {
				fileInfo, _ = file.Info()
				if fileInfo.IsDir() {
					newDs := DirStruct{fileInfo, path + string(os.PathSeparator) + fileInfo.Name()}
					stack.PushFront(newDs)
				} else {
					fmt.Printf("%s%s%s:%0.2fkb\r\n", path, string(os.PathSeparator), fileInfo.Name(), float32(fileInfo.Size())/1024)
				}
			}
		}
	} else {
		fmt.Printf("%s%s%s:%0.2fkb\r\n", path, string(os.PathSeparator), fileInfo.Name(), float32(fileInfo.Size())/1024)
	}
}

func printAllFileWithWalkDir(path string) {
	subDirToSkip := "skip"

	err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() && info.Name() == subDirToSkip {
			fmt.Printf("skipping a dir without errors: %+v \n", info.Name())
			return filepath.SkipDir
		}
		fmt.Printf("visited file or dir: %q\n", path)
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", path, err)
		return
	}
}

func init() {
	rootCmd.AddCommand(listFileDfsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listFileDfsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listFileDfsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
