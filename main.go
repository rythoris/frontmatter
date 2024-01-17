package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"slices"

	"github.com/adrg/frontmatter"
	"github.com/alexflint/go-arg"
	jsoniter "github.com/json-iterator/go"
	"gopkg.in/yaml.v2"
)

type Cli struct {
	File         string      `arg:"positional,required" help:"markdown file path"`
	Format       PrintFormat `arg:"-f,--format" help:"frontmatter output format (possible values: 'yaml', 'json')" default:"yaml"`
	PrintContent bool        `arg:"-c,--content" help:"only print content of the file"`
}

func main() {
	var cli Cli
	arg.MustParse(&cli)

	if err := ValidateFilePath(cli.File); err != nil {
		fmt.Fprintf(os.Stderr, "[!] ERROR: %s\n", err.Error())
		os.Exit(1)
	}

	fileContent, err := os.ReadFile(cli.File)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[!] ERROR: read error: %s\n", err.Error())
		os.Exit(1)
	}

	var fm any
	content, err := frontmatter.Parse(bytes.NewReader(fileContent), &fm)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[!] ERROR: frontmatter parse error: %s\n", err.Error())
		os.Exit(1)
	}

	if cli.PrintContent {
		fmt.Print(string(content))
		os.Exit(0)
	}

	err = printFronmatter(fm, cli.Format)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[!] ERROR: %s\n", err.Error())
		os.Exit(1)
	}
}

func ValidateFilePath(filePath string) error {
	stat, err := os.Stat(filePath)
	if errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("file doesn't exists: %s\n", filePath)
	} else if err != nil {
		return fmt.Errorf("could not get file stat: %w\n", err)
	}

	if stat.IsDir() {
		return fmt.Errorf("expected file got directory path: %s\n", filePath)
	}

	return nil
}

type PrintFormat uint

const (
	FormatJson PrintFormat = iota
	FormatYaml
)

var printFormatNames = [...]string{
	FormatJson: "json",
	FormatYaml: "yaml",
}

func (f *PrintFormat) UnmarshalText(data []byte) error {
	i := slices.Index(printFormatNames[:], string(data))
	if i == -1 {
		return fmt.Errorf("invalid print format: %s", string(data))
	}
	*f = PrintFormat(i)
	return nil
}

func printFronmatter(frontmatter any, format PrintFormat) error {
	switch format {
	case FormatJson:
		return printFrontmatterJson(frontmatter)
	case FormatYaml:
		return printFrontmatterYaml(frontmatter)
	default:
		panic("unreachable")
	}
}

func printFrontmatterJson(frontmatter any) error {
	b, err := jsoniter.Marshal(frontmatter)
	if err != nil {
		return fmt.Errorf("json marshal error: %w", err)
	}
	fmt.Println(string(b))
	return nil
}

func printFrontmatterYaml(frontmatter any) error {
	b, err := yaml.Marshal(frontmatter)
	if err != nil {
		return fmt.Errorf("yaml marshal error: %w", err)
	}
	fmt.Println(string(b))
	return nil
}
