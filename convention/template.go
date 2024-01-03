package convention

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/aymerick/raymond"
	"github.com/aymerick/raymond/ast"
	"github.com/aymerick/raymond/parser"
	"github.com/bar-counter/slog"
)

// RaymondRender parses and executes a template, returning the results in string
// format. Trailing or leading spaces or new-lines are not getting truncated. It
// is able to read templates from remote paths, local files or directly from the
// string.
func RaymondRender(template string, payload interface{}) (s string, err error) {
	u, err := url.Parse(template)

	if err == nil {
		switch u.Scheme {
		case "http", "https":
			res, err := http.Get(template)

			if err != nil {
				return s, fmt.Errorf("failed to fetch: %w", err)
			}

			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					slog.Errorf(err, "failed to close body")
				}
			}(res.Body)

			out, err := io.ReadAll(res.Body)

			if err != nil {
				return s, fmt.Errorf("failed to read: %w", err)
			}

			template = string(out)
		case "file":
			out, err := os.ReadFile(u.Path)

			if err != nil {
				return s, fmt.Errorf("failed to read: %w", err)
			}

			template = string(out)
		}
	}

	return raymond.Render(template, payload)
}

// RaymondParseAst
// parse by input
func RaymondParseAst(input string) (*ast.Program, error) {
	result, err := parser.Parse(input)
	return result, err
}
