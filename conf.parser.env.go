package conf

import (
	"io"
	"os"
	"strings"

	errs "github.com/bdlm/errors"
)

/*
envParser implements github.com/bdlm/std.Parser.

envParser parses environment variables into a configuration value map. This
implementation ignores r and reads data from os.Environ instead.
*/
type envParser struct{}

/*
Parse implements github.com/bdlm/std.Parser.

This implementation does not read from the provided reader and instead
expects a *conf pointer.
*/
func (parser envParser) Parse(r io.Reader) (interface{}, error) {
	cnf, ok := r.(*conf)
	if !ok {
		return nil, errs.New(0, "provided reader must be a a *conf pointer")
	}
	cnf.values = ValueMap{}
	for _, e := range os.Environ() {
		p1 := strings.Split(e, "\n")
		if len(p1) > 0 {
			p2 := strings.Split(p1[0], "=")
			switch {
			case "_" == p2[0]:
				continue
			case strings.HasPrefix(p2[0], "BASH_FUNC_"):
				continue
			case "" != cnf.envPrefix && !strings.HasPrefix(p2[0], cnf.envPrefix):
				continue
			}
			cnf.values[p2[0]] = NewVal(String, os.Getenv(p2[0]))
		}
	}
	return cnf, nil
}
