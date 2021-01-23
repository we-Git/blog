package cmd

import (
	"blog/github.com/gohub/blog/internal/word"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

var str string
var mode int8

const (
	MODE_UPPER = iota + 1
	MODE_LOWER
	MODE_UNDERSCORE_TOUPPER_CAMELCASE
	MODE_UNDERSCORE_LOWER_CAMELCASE
	MODE_CAMELCASE_TOUNDERSCORE
)

var desc = strings.Join([]string{
	"该子命令支持各种单词格式转换，模式如下：",
	"1: 全部单词转换为大写",
	"2: 全部单词转换为小写",
	"3: 下划线单词转换为大写驼峰单词",
	"4: 下划线单词转换为小写驼峰单词",
	"4: 驼峰单词转换为下划线单词",
}, "\n")
var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case MODE_UPPER:
			content = word.ToUpper(str)
		case MODE_LOWER:
			content = word.ToLower(str)
		case MODE_UNDERSCORE_TOUPPER_CAMELCASE:
			content = word.UnderscoreToUpperCamelcase(str)
		case MODE_UNDERSCORE_LOWER_CAMELCASE:
			content = word.UnderscoreToLowerCamelcase(str)
		case MODE_CAMELCASE_TOUNDERSCORE:
			content = word.CamelCaseToUnderscore(str)
		default:
			log.Fatal("暂不支持，执行help")

		}
		log.Printf("输出结果: %s", content)
	},
}

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输出单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输出单词转换格式")
}
