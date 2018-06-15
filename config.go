package main

import (
	"io"
	"os"

	"github.com/mattlisiv/typewriter"
)

type CommandConfig struct {
	out                 io.Writer
	customName          string
	Directive           string
	OutputDirectoryPath string
	InputDirectoryPath  string
	*typewriter.Config
}

var defaultConfig = CommandConfig{
	out:        os.Stdout,
	Directive:  "+gen",
	customName: "_gen.go",
	Config:     &typewriter.Config{},
}

func NewConfig(args []string) CommandConfig {
	if len(args) >= 3 {
		return CommandConfig{
			out:                 os.Stdout,
			Directive:           args[0],
			customName:          "_gen.go",
			Config:              &typewriter.Config{},
			InputDirectoryPath:  args[1],
			OutputDirectoryPath: args[2],
		}
	}

	if len(args) >= 2 {
		return CommandConfig{
			out:                 os.Stdout,
			Directive:           args[0],
			customName:          "_gen.go",
			Config:              &typewriter.Config{},
			OutputDirectoryPath: "./",
			InputDirectoryPath:  args[1],
		}
	}

	if len(args) >= 1 {
		out := "./"
		return CommandConfig{
			out:                 os.Stdout,
			Directive:           args[0],
			customName:          "_gen.go",
			Config:              &typewriter.Config{},
			InputDirectoryPath:  out,
			OutputDirectoryPath: out,
		}
	}
	return defaultConfig
}

// keep in sync with imports.go
var stdImports = typewriter.NewImportSpecSet(
	typewriter.ImportSpec{Name: "_", Path: "github.com/clipperhouse/slice"},
	typewriter.ImportSpec{Name: "_", Path: "github.com/mattlisiv/stringer"},
)
