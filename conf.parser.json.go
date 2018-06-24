package conf

import (
	"encoding/json"
	"io"

	errs "github.com/bdlm/errors"
)

type jsonParser struct{}

func (parser jsonParser) Parse(r io.Reader) (interface{}, error) {
	var v interface{}
	var data []byte
	cnf, ok := r.(*conf)
	if !ok {
		return nil, errs.New(0, "provided reader must be a a *conf pointer, %v recieved", cnf)
	}
	n, err := r.Read(data)
	if n > 0 {
		err = json.Unmarshal(data, &v)
	}
	return v, err
}
