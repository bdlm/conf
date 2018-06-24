package conf

import (
	"os"
	"path/filepath"

	errs "github.com/bdlm/errors"
	"github.com/bdlm/std"
	log "github.com/sirupsen/logrus"
)

var _pwd string

func init() {
	_pwd, _ = filepath.Abs(filepath.Dir(os.Args[0]))
}

//type conf struct {
//	cnfSrc    cnfSrc
//	cnfType   cnfType
//	data      interface{}
//	defaults  ValueMap
//	envPrefix string
//	logger    std.Logger
//	name      string
//	path      []string
//	parser    std.Parser
//	values    ValueMap
//	watch     bool
//}

/*
New returns a new Configuration interface.
*/
func New(
	name string,
	cnfSrc cnfSrc,
	cnfType cnfType,
	path []string,
	defaults map[string]Valuer,
	logger std.Logger,
) Configuration {
	if nil == defaults {
		defaults = make(map[string]Valuer)
	}
	if nil == path {
		path = []string{_pwd}
	}
	if nil == logger {
		logger = log.New()
		logger.(*log.Logger).SetLevel(log.InfoLevel)
	}
	return &conf{
		defaults:  defaults,
		envPrefix: "",
		logger:    logger,
		name:      name,
		path:      path,
		values:    map[string]Valuer{},
		watch:     false,
	}
}

/*
AddPath implements Configuration.

AddPath adds a directory path to the file search path. Paths are searched in
FIFO order.
*/
func (conf *conf) AddPath(path string) *conf {
	conf.path = append(conf.path, path)
	return conf
}

/*
Get implements Configuration.

Get retrieves a value by name.
*/
func (conf *conf) Get(name string) Valuer {
	if val, ok := conf.values[name]; ok {
		return val
	}
	if val, ok := conf.defaults[name]; ok {
		return val
	}
	return nil
}

/*
GetDefault implements Configuration.

GetDefault retrieves a default value by name.
*/
func (conf *conf) GetDefault(name string) Valuer {
	if val, ok := conf.defaults[name]; ok {
		return val
	}
	return nil
}

/*
Name implements Configuration.

Name returns the configuration name.
*/
func (conf *conf) Name() string {
	return conf.name
}

func (conf *conf) Parse() error {
	switch conf.cnfSrc {
	case Env:
		conf.SetParser(envParser{})
	case File:
		switch conf.cnfType {
		default:
			fallthrough
		case JSON:
			conf.SetParser(jsonParser{})
		case TOML:
			conf.SetParser(tomlParser{})
		case YAML:
			conf.SetParser(yamlParser{})
		}
	}
	_, err := conf.parser.Parse(conf)
	return err
}

/*
Path implements Configuration.

Path returns the configuration file search path.
*/
func (conf *conf) Path() []string {
	return conf.path
}

/*
Read implements Configuration.

Read reads configuration into the provided byte array.
*/
func (conf *conf) Read(p []byte) (n int, err error) {
	dLen := len(conf.data)
	pLen := len(p)
	if pLen == dLen {
		p = conf.data[:pLen-1]
		return pLen, nil
	}
	if pLen < dLen {
		p = conf.data[:pLen-1]
		return pLen, errs.New(0, "receiver p is not larg enough to hold the data")
	}
	p = append(conf.data[:dLen-1], p[dLen:]...)
	return dLen, nil
}

/*
Set implements Configuration.

Set stores a value by name.
*/
func (conf *conf) Set(name string, value Valuer) *conf {
	conf.values[name] = value
	return conf
}

/*
SetDefault implements Configuration.

SetDefault stores a default value by name.
*/
func (conf *conf) SetDefault(name string, value Valuer) *conf {
	conf.defaults[name] = value
	return conf
}

/*
SetDefaults implements Configuration.

SetDefaults replaces the default value set.
*/
func (conf *conf) SetDefaults(defaults map[string]Valuer) *conf {
	conf.defaults = defaults
	return conf
}

func (conf *conf) SetParser(parser std.Parser) {
	conf.parser = parser
}

/*
SetEnvPrefix implements Configuration.

SetEnvPrefix sets the environment variable name prefix for this
configuration.
*/
func (conf *conf) SetEnvPrefix(prefix string) *conf {
	conf.envPrefix = prefix
	return conf
}

/*
SetName implements Configuration.

SetName sets this configuration's name.
*/
func (conf *conf) SetName(name string) *conf {
	conf.name = name
	return conf
}

/*
SetPath implements Configuration.

SetPath replaces the file search path for this configuration.
*/
func (conf *conf) SetPath(path []string) *conf {
	conf.path = path
	return conf
}

/*
Watch implements Configuration.

Watch watches configuration values for changes and reloads them when
updated.
*/
func (conf *conf) Watch() {
	//	conf.watch = true
	//
	//	go func() {
	//		watcher, err := fsnotify.NewWatcher()
	//		if err != nil {
	//			conf.logger.Fatal(err)
	//		}
	//		defer watcher.Close()
	//
	//		// we have to watch the entire directory to pick up renames/atomic
	//		// saves in a cross-platform way
	//		filename, err := v.getConfigFile()
	//		if err != nil {
	//			conf.logger.Error(err)
	//			return
	//		}
	//
	//		configFile := filepath.Clean(filename)
	//		configDir, _ := filepath.Split(configFile)
	//
	//		done := make(chan bool)
	//		go func() {
	//			for {
	//				select {
	//				case event := <-watcher.Events:
	//					// we only care about the config file
	//					if filepath.Clean(event.Name) == configFile {
	//						if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create {
	//							err := v.ReadInConfig()
	//							if err != nil {
	//								conf.logger.Infoln("error:", err)
	//							}
	//							v.onConfigChange(event)
	//						}
	//					}
	//				case err := <-watcher.Errors:
	//					conf.logger.Infoln("error:", err)
	//				}
	//			}
	//		}()
	//
	//		watcher.Add(configDir)
	//		<-done
	//	}()
}

type cnfSrc uint

type conf struct {
	cnfSrc    cnfSrc
	cnfType   cnfType
	data      []byte
	defaults  ValueMap
	envPrefix string
	logger    std.Logger
	name      string
	path      []string
	parser    std.Parser
	values    ValueMap
	watch     bool
}
