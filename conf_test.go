package conf_test

import (
	"testing"

	"github.com/bdlm/conf"
	"github.com/bdlm/logfmt"
	log "github.com/sirupsen/logrus"
)

func init() {
	level, _ := log.ParseLevel("debug")
	log.SetLevel(level)
	log.SetFormatter(&logfmt.TextFormat{})
}

func TestMsg(t *testing.T) {
	cnf := conf.New(
		"config",
		conf.Env,
		0,
		nil,
		nil,
		nil,
	)
	cnf.Parse()

	val := cnf.Get("K8S_CURRENT_NAMESPACE")
	if nil == val {
		t.Errorf("val is nil")
	}

	a, err := val.String()
	t.Errorf("%v, %v", a, err)
}
