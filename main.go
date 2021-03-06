package main

import (
	"flag"
	"strconv"
	"os"

	"github.com/go-ini/ini"
	"github.com/fatih/structs"
)

type Config struct {
	Root                   bool
	Charset                string
	EndOfLine              string
	IndentSize             int
	IndentStyle            string
	InsertFinalNewLine     bool
	TrimTrailingWhitespace bool
}

const (
	commonConfigSection  = "*"
	defaultConfigSection = ""
)

var (
	r = flag.Bool("r", true, "`root:` special property that should be " +
		"specified at the top of the file outside of any sections. " +
		"Set to true to stop .editorconfig files search on current file.")

	c = flag.String("c", "UTF-8", "`charset:` set to latin1, utf-8, " +
		"utf-8-bom, utf-16be or utf-16le to control the character set. " +
		"Use of utf-8-bom is discouraged.")

	i = flag.Int("i", 4, "`indent_size:` a whole number defining the " +
		"number of columns used for each indentation level and the " +
		"width of soft tabs (when supported). When set to tab, the " +
		"value of tab_width (if specified) will be used.")

	e = flag.String("e", "lf", "`end_of_line:` set to lf, cr, or crlf to " +
		"control how line breaks are represented.")

	s = flag.String("s", "space", "`indent_style:` set to tab or space to " +
		"use hard tabs or soft tabs respectively.")

	f = flag.Bool("f", true, "`insert_final_newline:` set to true to " +
		"ensure file ends with a newline when saving and false to " +
		"ensure it doesn't.")

	t = flag.Bool("t", true, "`trim_trailing_whitespace:` set to true to " +
		"remove any whitespace characters preceding newline " +
		"characters and false to ensure it doesn't.")
)

var (
	keysMapping = map[string]string{
		"Root":                   "root",
		"Charset":                "charset",
		"EndOfLine":              "end_of_line",
		"IndentSize":             "indent_size",
		"IndentStyle":            "indent_style",
		"InsertFinalNewLine":     "insert_final_newline",
		"TrimTrailingWhitespace": "trim_trailing_whitespace",
	}
)

func main() {
	flag.Parse()

	iniStub := ini.Empty()

	config := &Config{
		Root:                   *r,
		Charset:                *c,
		EndOfLine:              *e,
		IndentSize:             *i,
		IndentStyle:            *s,
		InsertFinalNewLine:     *f,
		TrimTrailingWhitespace: *t,
	}

	cStruct := structs.New(config)

	if config.Root != false {
		iniStub.Section(defaultConfigSection).NewKey(
			keysMapping["Root"], strconv.FormatBool(*r))
	}

	dStub := iniStub.Section(commonConfigSection)

	for key, value := range keysMapping {
		v := cStruct.Field(key).Value()

		if value == keysMapping["Root"] {
			continue
		}

		switch v.(type) {
		case bool:
			dStub.NewKey(value, strconv.FormatBool(v.(bool)))
		case int:
			dStub.NewKey(value, strconv.Itoa(v.(int)))
		default:
			dStub.NewKey(value, v.(string))
		}
	}

	iniStub.WriteTo(os.Stdout)
}
