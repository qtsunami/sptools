package cmd

import (
	"github.com/qtsunami/sptools/internal/du"
	"strings"

	"github.com/spf13/cobra"
)

var rpath string
var unit string

const (
	UNIT_KB = "KB"
	UNIT_MB = "MB"
	UNIT_GB = "GB"
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
		bdata := du.Start(rpath)
		for _, item := range bdata {
			var size float64
			switch unit {
			case UNIT_KB:
				size = float64(item.NBytes) / 1e3
			case UNIT_MB:
				size = float64(item.NBytes) / 1e6
			case UNIT_GB:
				size = float64(item.NBytes) / 1e9
			default:
				size = float64(item.NBytes) / 1e3
			}
			du.PrintDiskUsage(item.Path, item.NFile, size, unit)
		}
	},
}

func init() {
	duCmd.Flags().StringVarP(&rpath, "path", "m", "./", "指定路径，多个路径以逗号相隔")
	duCmd.Flags().StringVarP(&unit, "unit", "u", "KB", "显示单位:KB,MB,GB")
}
