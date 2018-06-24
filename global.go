package conf

var _globals = map[string]Configuration{}

/*
Set sets a Configuration in the global namespace.
*/
func Set(
	name string,
	conf Configuration,
) {
	_globals[name] = conf
}

/*
Del removes a configuration from the global namespace.
*/
func Del(name string) {
	delete(_globals, name)
}

/*
Get returns a value from teh global configuration.
*/
func Get(name string) Configuration {
	return _globals[name]
}
