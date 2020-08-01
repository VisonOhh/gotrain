package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
)

func main() {
	//flagParseSimple()
	//flagParseSubCommand()
	flagParseCustomParamType()
}

// 1. Simple parse command parameter by flag
func flagParseSimple() {
	var name string
	flag.StringVar(&name, "name", "Visonhuo", "name help info")
	flag.StringVar(&name, "n", "Visonhuo", "name help info")
	flag.Parse()

	log.Printf("name: %s", name)
}

// 2. Sub command usage
func flagParseSubCommand() {
	var name string
	flag.Parse()
	goCmd := flag.NewFlagSet("go", flag.ExitOnError)
	goCmd.StringVar(&name, "name", "Golang", "go helper")

	javaCmd := flag.NewFlagSet("java", flag.ExitOnError)
	javaCmd.StringVar(&name, "n", "java", "java helper")

	args := flag.Args()
	switch args[0] {
	case "go":
		_ = goCmd.Parse(args[1:])
	case "java":
		_ = javaCmd.Parse(args[1:])
	}

	log.Printf("name: %s", name)
}

// 3. Parse costom flag parameter type
type Name string

func (n *Name) String() string {
	return fmt.Sprint(*n)
}

func (n *Name) Set(value string) error {
	if len(*n) > 0 {
		return errors.New("name flag already set")
	}
	*n = Name("Name: " + value)
	return nil
}

func flagParseCustomParamType() {
	var name Name
	flag.Var(&name, "name", "name help")
	flag.Parse()

	log.Printf("Value: %s", name)
}
