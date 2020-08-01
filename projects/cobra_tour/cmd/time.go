package cmd

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gotrain/projects/cobra_tour/internal/timer"
	"github.com/spf13/cobra"
)

const (
	LayoutDay     = "2006-01-02"
	LayoutDayTime = "2006-01-02 15:04:05"
)

var calculateTime string
var duration string

func init() {
	timeCmd.AddCommand(timeCmdNow)
	timeCmd.AddCommand(timeCmdCalc)

	timeCmdCalc.Flags().StringVarP(&calculateTime, "calculate", "c", "", "format time like '2006-01-02 15:04:05' or timestamp")
	timeCmdCalc.Flags().StringVarP(&duration, "duration", "d", "", "duration time, valid unit like: 'ns', 'us', 's', 'm'")
}

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "Time format",
	Long:  "Time format",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var timeCmdNow = &cobra.Command{
	Use:   "now",
	Short: "Get current time",
	Long:  "Get current time info and timestamp",
	Run: func(cmd *cobra.Command, args []string) {
		now := timer.GetNowTime()
		fmt.Printf("format=%v, timestamp=%v", now.Format(LayoutDayTime), now.Unix())
	},
}

var timeCmdCalc = &cobra.Command{
	Use:   "calc",
	Short: "calc time by params",
	Long:  "calc time by params",
	Run: func(cmd *cobra.Command, args []string) {
		var curTime time.Time
		var layout = LayoutDayTime

		if calculateTime == "" {
			curTime = timer.GetNowTime()
		} else {
			var err error
			if !strings.Contains(calculateTime, " ") {
				layout = LayoutDay
			}

			curTime, err = time.Parse(layout, calculateTime)
			if err != nil {
				// if not format time string, check whether is timestamp
				t, err := strconv.ParseInt(calculateTime, 10, 64)
				if err != nil {
					log.Fatalf("error format string, should be '%v' or timestamp", LayoutDayTime)
				}
				curTime = time.Unix(t, 0)
			}
		}

		calculateTime, err := timer.GetCalculateTime(curTime, duration)
		if err != nil {
			log.Fatalf("timer.GetCalculateTime err: %v", err)
		}
		fmt.Printf("format=%v, timestamp=%v", calculateTime.Format(LayoutDayTime), calculateTime.Unix())
	},
}
