package conf

import (
	"io"

	errs "github.com/bdlm/errors"
)

type tomlParser struct{}

func (parser tomlParser) Parse(r io.Reader) (interface{}, error) {
	var data []byte
	cnf, ok := r.(*conf)
	if !ok {
		return nil, errs.New(0, "provided reader must be a a *conf pointer, %v recieved", cnf)
	}
	_, err := r.Read(data)
	return data, err
}
