package cmd

import (
	"blog/github.com/gohub/blog/internal/timer"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
	"time"
)

var calculateTime string
var duration string

var timeCmd = &cobra.Command{
	//time命令初始化
	Use:   "time change",
	Short: "时间转换",
	Long:  "时间转换",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var nowTimeCmd = &cobra.Command{
	// time 子命令初始化
	Use:   "now",
	Short: "获取当前时间",
	Long:  "获取当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetNowTime()
		log.Println(nowTime)
		log.Printf("输出结果：%s,%d", nowTime.Format("2006-01-02 15:04:05"), nowTime.Unix())
	},
}

var calculateTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "计算所需时间",
	Long:  "计算所需时间",
	Run: func(cmd *cobra.Command, args []string) {
		var currentTimer time.Time
		var layout = "2006-01-02 15:04:05"
		if calculateTime == "" {
			currentTimer = time.Now()
		} else {
			var err error
			if !strings.Contains(calculateTime, " ") {
				//不包含空格
				layout = "2006-01-02"
			}
			currentTimer, err = time.Parse(layout, calculateTime)
			log.Print(currentTimer)
			if err != nil {
				t, _ := strconv.Atoi(calculateTime)
				currentTimer = time.Unix(int64(t), 0)
			}

		}
		calculateTime, err := timer.GetCalculateTime(currentTimer, duration)
		log.Print("test:", calculateTime)
		if err != nil {
			log.Printf("timer.GetCalculateTime err: %v", err)
		}
		log.Printf("输出结果：%s,%d", calculateTime.Format(layout), calculateTime.Unix())
	},
}

func init() {
	//为time添加子命令
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calculateTimeCmd)
	calculateTimeCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "", "需要计算的时间")
	calculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", "持续时间")
}
