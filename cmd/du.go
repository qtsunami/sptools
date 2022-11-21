package cmd

import (
	"strings"

	"github.com/spf13/cobra"
)

var multipath string
var unit string

const (
	UNIT_KB = iota + 1
	UNIT_MB
	UNIT_GB
)

var longDesc = strings.Join([]string{
	"该命令查看指定目录的占用空间及显示格式，格式如下：",
	"1: 展示 KB",
	"2：展示 MB",
	"3：展示 GB",
}, "\n")

var duCmd = &cobra.Command{
	Use:   "du",
	Short: "查看指定目录的占用空间",
	Long:  longDesc,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO
	},
}

func init() {
	duCmd.Flags().StringVarP(&multipath, "multipath", "mp", "", "指定路径，多个路径以逗号相隔")
	duCmd.Flags().StringVarP(&unit, "unit", "u", "", "显示单位:KB,MB,GB")
}
