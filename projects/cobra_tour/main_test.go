package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/agiledragon/gomonkey"
	"github.com/stretchr/testify/assert"
	"github.com/visonohh/gotrain/cobra_tour/cmd"
)

const (
	CobraTourCmdName = "./cobra_tour"

	SupportLayoutDayTime = "2006-01-02 15:04:05"
)

// init cobar_tour program as an executable file
func SetUpTestCase() {
	execCmd("go", "build")
}

func TearDownTestCase() {
	err := os.Remove(CobraTourCmdName)
	if err != nil {
		log.Fatalf("os.Remove fail, cmd=%v", CobraTourCmdName)
	}
}

func execCmd(cmd string, args ...string) string {
	var output bytes.Buffer
	command := exec.Command(cmd, args...)
	command.Stdout = &output

	err := command.Run()
	if err != nil {
		log.Fatalf("cmd.Run fail, cmd=%v, err=%v", cmd, err)
	}
	return output.String()
}

// execute word command with str and mode param, return command output string
func execSubCmdWord(str string, mode int8) string {
	args := []string{
		"word",
	}

	if str != "" {
		args = append(args, "-s="+str)
	}
	if mode != 0 {
		args = append(args, "-m="+strconv.Itoa(int(mode)))
	}
	return execCmd(CobraTourCmdName, args...)
}

// execute time now command
func execSubCmdTimeNow() string {
	return execCmd(CobraTourCmdName, "time", "now")
}

// execute time calc command with calculateTime and duration params
func execSubCmdTimeCalc(ct string, d string) string {
	args := []string{
		"time",
		"calc",
	}

	if ct != "" {
		args = append(args, "-c="+ct)
	}
	if d != "" {
		args = append(args, "-d="+d)
	}
	return execCmd(CobraTourCmdName, args...)
}

func TestCobraTourCmd(t *testing.T) {
	SetUpTestCase()
	defer TearDownTestCase()

	t.Run("TestSubCmdWord", func(tt *testing.T) {
		tt.Run("ModeUpper", func(ttt *testing.T) {
			testCases := []struct {
				str    string
				expect string
			}{
				{"", ""},
				{"   ", "   "},
				{"Vison", "VISON"},
				{"JIMIN", "JIMIN"},
			}

			for _, tc := range testCases {
				actual := execSubCmdWord(tc.str, cmd.ModeUpper)
				if actual != tc.expect {
					ttt.Errorf("input=%v, expect=%v, actual=%v\n", tc.str, tc.expect, actual)
				}
			}
		})

		tt.Run("ModeLower", func(ttt *testing.T) {
			testCases := []struct {
				str    string
				expect string
			}{
				{"", ""},
				{"   ", "   "},
				{"Vison", "vison"},
				{"JIMIN", "jimin"},
			}

			for _, tc := range testCases {
				actual := execSubCmdWord(tc.str, cmd.ModeLower)
				if actual != tc.expect {
					ttt.Errorf("input=%v, expect=%v, actual=%v\n", tc.str, tc.expect, actual)
				}
			}
		})

		tt.Run("ModeUnderscoreToUpperCamelcase", func(ttt *testing.T) {
			testCases := []struct {
				str    string
				expect string
			}{
				{"", ""},
				{"   ", "   "},
				{"vison_huo_he", "VisonHuoHe"},
				{"_vison_he", "VisonHe"},
				{"jimin_wang_i_i__", "JiminWangII"},
			}

			for _, tc := range testCases {
				actual := execSubCmdWord(tc.str, cmd.ModeUnderscoreToUpperCamelcase)
				if actual != tc.expect {
					ttt.Errorf("input=%v, expect=%v, actual=%v\n", tc.str, tc.expect, actual)
				}
			}
		})

		tt.Run("ModeUnderscoreToLowerCamelcase", func(ttt *testing.T) {
			testCases := []struct {
				str    string
				expect string
			}{
				{"", ""},
				{"   ", "   "},
				{"vison_huo_he", "visonHuoHe"},
				{"_vison_he", "visonHe"},
				{"jimin_wang_i_i__", "jiminWangII"},
			}

			for _, tc := range testCases {
				actual := execSubCmdWord(tc.str, cmd.ModeUnderscoreToLowerCamelcase)
				if actual != tc.expect {
					ttt.Errorf("input=%v, expect=%v, actual=%v\n", tc.str, tc.expect, actual)
				}
			}
		})

		tt.Run("ModeUnderscoreToLowerCamelcase", func(ttt *testing.T) {
			testCases := []struct {
				str    string
				expect string
			}{
				{"", ""},
				{"   ", "   "},
				{"visonHuoHe", "vison_huo_he"},
				{"JiminWangII", "jimin_wang_i_i"},
			}

			for _, tc := range testCases {
				actual := execSubCmdWord(tc.str, cmd.ModeCamelcaseToUnderscore)
				if actual != tc.expect {
					ttt.Errorf("input=%v, expect=%v, actual=%v\n", tc.str, tc.expect, actual)
				}
			}
		})
	})

	t.Run("TestSubCmdTime", func(tt *testing.T) {
		tt.Run("Now", func(ttt *testing.T) {
			// mock now time
			mockTime := time.Now()
			patches := mockTimeNow(mockTime)
			defer patches.Reset()

			expectFormatTime := mockTime.Format("2006-01-02 15:04:05")
			expectTimestamp := mockTime.Unix()

			result := execSubCmdTimeNow()
			checkSubCmdTimeResult(ttt, result, expectFormatTime, expectTimestamp)
		})

		tt.Run("Calc", func(ttt *testing.T) {
			// mock now time
			mockTime := time.Now()
			patches := mockTimeNow(mockTime)
			defer patches.Reset()

			testCases := []struct {
				calculateTime    string
				duration         string
				expectFormatTime string
				expectTimestamp  int64
			}{
				{
					// calculate time is empty, default use now
					"",
					"12h",
					mockTime.Add(12 * time.Hour).Format(SupportLayoutDayTime),
					mockTime.Add(12 * time.Hour).Unix(),
				},
				{
					// format day time string
					"2020-08-01 00:00:00",
					"-12h",
					"2020-07-31 12:00:00",
					parseFormatToUnix(SupportLayoutDayTime, "2020-07-31 12:00:00"),
				},
				{
					// format day str
					"2020-08-01",
					"-12h",
					"2020-07-31 12:00:00",
					parseFormatToUnix(SupportLayoutDayTime, "2020-07-31 12:00:00"),
				},
				{
					// use timestamp as string
					strconv.FormatInt(mockTime.Unix(), 10),
					"12h",
					mockTime.Add(12 * time.Hour).Format(SupportLayoutDayTime),
					mockTime.Add(12 * time.Hour).Unix(),
				},
			}

			for _, tc := range testCases {
				result := execSubCmdTimeCalc(tc.calculateTime, tc.duration)
				checkSubCmdTimeResult(ttt, result, tc.expectFormatTime, tc.expectTimestamp)
			}
		})
	})
}

func mockTimeNow(mockTime time.Time) *gomonkey.Patches {
	return gomonkey.ApplyFunc(time.Now, func() time.Time {
		return mockTime
	})
}

func checkSubCmdTimeResult(t *testing.T, result, expectFormatTime string, expectTimestamp int64) {
	timestampStr := strconv.FormatInt(expectTimestamp, 10)
	assert.Truef(t, strings.Contains(result, expectFormatTime),
		"result should contains formatTimeStr, result=%v, formatTimeStr=%v", result, expectFormatTime)
	assert.Truef(t, strings.Contains(result, timestampStr),
		"result should contains timestamp, result=%v, timestamp=%v", result, timestampStr)
}

func parseFormatToUnix(layout, formatTime string) int64 {
	t, err := time.Parse(layout, formatTime)
	if err != nil {
		log.Fatalf("time.Parse err: %v", err)
	}
	return t.Unix()
}
