package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/visonohh/gotrain/cobra_tour/internal/word"
)

const (
	ModeUpper = iota + 1
	ModeLower
	ModeUnderscoreToUpperCamelcase
	ModeUnderscoreToLowerCamelcase
	ModeCamelcaseToUnderscore
)

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "pls input str content")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "pls input word conversion mode")
	mode2Func = map[int8]WordConvertFunc{
		ModeUpper:                      word.ToUpper,
		ModeLower:                      word.ToLower,
		ModeUnderscoreToUpperCamelcase: word.UnderscoreToUpperCamelCase,
		ModeUnderscoreToLowerCamelcase: word.UnderscoreToLowerCamelCase,
		ModeCamelcaseToUnderscore:      word.CamelCaseToUnderscore,
	}
}

var desc = strings.Join([]string{
	"This sub command support these word format conversion: ",
	"1: Upper all words",
	"2: Lower all words",
	"3: Convert underscore words to upper camel case words",
	"4: Convert underscore words to lower camel case words",
	"5: Convert camel case words to underscore words",
}, "\n")

// Use this type to replace word convert funct
type WordConvertFunc func(string) string

var (
	mode      int8
	str       string
	mode2Func map[int8]WordConvertFunc
)

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "Word format conversion",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		wf, ok := mode2Func[mode]
		if !ok {
			log.Fatalf("invalid word conversion mode, pls execute 'help word' to watch help doc")
		}
		fmt.Print(wf(str))
	},
}
