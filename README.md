# autoindex

Implements Nginx's [`autoindex`](https://nginx.org/en/docs/http/ngx_http_autoindex_module.html) module in Go.

## Getting started

### CLI

Download and install the Autoindex CLI from [autoindex/cli](https://github.com/go-autoindex/cli).

### GitHub Actions

See [autoindex/action](https://github.com/go-autoindex/action) for details.

### Library

```go
import "github.com/go-autoindex/autoindex"

const dir = "."

func main() {
	autoindex.Opts.Root = dir

	autoindex.Gen(autoindex.Info{
		Dir:     dir,
		Entries: []string{},
	})
}
```
