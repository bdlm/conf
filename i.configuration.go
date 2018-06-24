package conf

const (
	// Env directs the Configuration to read values from environment
	// variables.
	Env cnfSrc = iota
	// File directs the Configuration to read values from a file on the
	// local filesystem.
	File
)
const (
	// JSON directs the Configuration to expect JSON data.
	JSON cnfType = iota + 1
	// TOML directs the Configuration to expect TOML data.
	TOML
	// YAML directs the Configuration to expect YAML data.
	YAML
)

type cnfType uint

/*
Configuration defines a configuration interface...
*/
type Configuration interface {
	// AddPath adds a directory path to the file search path. Paths are
	// searched in FIFO order.
	AddPath(path string) *conf

	// Get retrieves a value by name.
	Get(name string) Valuer

	GetDefault(name string) Valuer

	// Name returns the configuration name. For a file-based config this
	// is the filename without the file extension.
	Name() string

	// Parse parses the configuration data.
	Parse() error

	// Path returns the list of file search paths.
	Path() []string

	// Read implements io.Reader
	Read(p []byte) (n int, err error)

	// Set sets a value in this configuration
	Set(name string, value Valuer) *conf

	// Define a default config value.
	SetDefault(name string, value Valuer) *conf

	// Define all default config values.
	SetDefaults(defaults map[string]Valuer) *conf

	// Define a prefix to use when accessing environment variables.
	// Environment variable names are case sensitive. Eg:
	//
	//	conf.SetEnvPrefix("SVC_")
	//	id := conf.Get("ID")
	SetEnvPrefix(prefix string) *conf

	// SetName sets the configuration name. For a file-based config this
	// is the filename without the file extension.
	SetName(name string) *conf

	// SetPath replaces the current file search path. Paths are searched
	// in FIFO order.
	SetPath(path []string) *conf

	// Watch watches configuration sources for changes and automatically
	// updates available values.
	Watch()
}
