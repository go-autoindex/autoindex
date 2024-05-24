package autoindex

// Opts represents the global options
var Opts options

// options represents options of autoindex
type options struct {
	// Dry whether or not to write the index
	Dry bool

	// Root represents the root directory of the index
	Root string
}
