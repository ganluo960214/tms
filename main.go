package main

import (
	"flag"
	"github.com/ganluo960214/ast_extend"
	"github.com/go-playground/validator/v10"
	"go/format"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func init() {
	log.SetFlags(0)
}

// validate /*
var (
	validate = validator.New()
)

// flags /*
var (
	flags = struct {
		Type  string `validate:"required"`
		Map   bool
		Slice bool
	}{
		Map:   true,
		Slice: true,
	}
)

func init() {
	flag.StringVar(&flags.Type, "type", "", "type name; must be set")
	flag.BoolVar(&flags.Map, "map", flags.Map, "generate map")
	flag.BoolVar(&flags.Slice, "slice", flags.Slice, "generate slice")
	flag.Parse()

	// check flags
	if err := validate.Struct(flags); err != nil {
		log.Fatal(err)
	}
}

// envs /*
var (
	envs = struct {
		GoPackage string `validate:"required"`
		GoFile    string `validate:"required,file,endswith=.go"`
	}{
		GoPackage: os.Getenv("GOPACKAGE"),
		GoFile:    os.Getenv("GOFILE"),
	}
)

func init() {
	if err := validate.Struct(envs); err != nil {
		log.Fatalln(err)
	}
}

func main() {

	astFile, err := parser.ParseFile(token.NewFileSet(), envs.GoFile, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	names := []string{}
	vss := ast_extend.FindConstByTypeName(astFile, flags.Type)
	for _, vs := range vss {
		names = append(names, vs.Names[0].Name)
	}

	t := TextTemplate{
		Flags:   strings.Join(os.Args, " "),
		Package: envs.GoPackage,
		Type:    flags.Type,
		Names:   names,
		Map:     flags.Map,
		Slice:   flags.Slice,
	}
	bs, err := t.Generate()
	if err != nil {
		log.Fatal(err)
	}

	fc, err := format.Source(bs)
	if err != nil {
		log.Fatal(err)
	}

	goEnd := ".go"
	goTestEnd := "_test.go"
	generateFileName := envs.GoFile[:len(envs.GoFile)-len(goEnd)] +
		"_tms" +
		envs.GoFile[len(envs.GoFile)-len(goEnd):]

	if strings.HasSuffix(envs.GoFile, goTestEnd) {
		generateFileName = envs.GoFile[:len(envs.GoFile)-len(goEnd)] +
			"_tms_test" +
			envs.GoFile[len(envs.GoFile)-len(goEnd):]
	}

	if err := ioutil.WriteFile(generateFileName, fc, 0644); err != nil {
		log.Fatalln(err)
	}
}
