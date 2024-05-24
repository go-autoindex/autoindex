package autoindex

import (
	_ "embed"
	"fmt"
	"html/template"
	"os"
)

var (
	//go:embed internal/template/index.tmpl
	indexTmpl string
	tmpl      = template.Must(template.New("").Parse(indexTmpl))
)

// Info represents information about the directory to index
type Info struct {
	// Dir is the name of the directory to index
	Dir string
	// Entries is the list of entries in the directory
	Entries []string
}

// Gen generates and writes an index using the given info
func Gen(info Info) error {
	index := Index(info.Dir)
	if err := write(index, info); err != nil {
		return err
	}

	fmt.Println(index)
	return nil
}

// write writes to name with the given info in non-dry mode only
func write(name string, info Info) error {
	if Opts.Dry {
		return nil
	}

	preWrite(&info)

	// Open name file for writing only
	f, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	return tmpl.Execute(f, info)
}

// preWrite prepares the given info for [write].
func preWrite(info *Info) {
	// Write only relative paths to the index
	info.Dir = rel(info.Dir)
}
