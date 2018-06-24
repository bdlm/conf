package conf

/*
Valuer is an interface that defines a configuration value.
*/
type Valuer interface {
	// Bool returns the boolean representation of the value, or an error if
	// the type conversion is not possible.
	Bool() (bool, error)
	// Float returns the float64 representation of the value, or an error if
	// the type conversion is not possible.
	Float() (float64, error)
	// Float32 returns the float32 representation of the value, or an error
	// if the type conversion is not possible.
	Float32() (float32, error)
	// Float64 returns the float64 representation of the value, or an error
	// if the type conversion is not possible.
	Float64() (float64, error)
	// Int returns the int representation of the value, or an error if the
	// type conversion is not possible.
	Int() (int, error)
	// List returns the ValueList representation of the value, or an error
	// if the type conversion is not possible.
	List() (ValueList, error)
	// Map returns the ValueMap representation of the value, or an error if
	// the type conversion is not possible.
	Map() (ValueMap, error)
	// String returns the boolean representation of the value, or an error
	// if the type conversion is not possible.
	String() (string, error)
}

/*
ValueList is a list of Valuers.
*/
type ValueList []Valuer

/*
ValueMap is a map of Valuers.
*/
type ValueMap map[string]Valuer

/*
Value type constants
*/
const (
	// Bool defines the boolean type.
	Bool valType = iota
	// Float defines the float type.
	Float
	// Int defines the integer type.
	Int
	// List defines the list type.
	List
	// Map defines the map type.
	Map
	// String defines the string type.
	String
)

type valType uint
