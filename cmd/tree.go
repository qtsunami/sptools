package cmd

import "github.com/spf13/cobra"

// TODO: 以树形结构罗列目录文件情况

var dir string

var treeCmd = &cobra.Command{
	Use:   "tree",
	Short: "展示目录",
	Long:  "展示目录",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	treeCmd.Flags().StringVarP(&dir, "dir", "d", "./", "目标路径，默认为当前路径")
}
