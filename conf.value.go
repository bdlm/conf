package conf

import (
	errs "github.com/bdlm/errors"
	"github.com/spf13/cast"
)

func NewVal(
	valueType valType,
	valueData interface{},
) Valuer {
	return &val{
		typ: valueType,
		val: valueData,
	}
}

/*
Bool implements Valuer.
*/
func (v val) Bool() (bool, error) {
	val, err := cast.ToBoolE(v.val)
	if nil != err {
		err = errs.Wrap(err, errs.ErrTypeConversionFailed, "could not convert value '%v' to a boolean", v.val)
	}
	return val, nil
}

/*
Float implements Valuer.
*/
func (v val) Float() (float64, error) {
	val, err := cast.ToFloat64E(v.val)
	if nil != err {
		err = errs.Wrap(err, errs.ErrTypeConversionFailed, "could not convert value '%v' to a float64", v.val)
	}
	return val, nil
}

/*
Float32 implements Valuer.
*/
func (v val) Float32() (float32, error) {
	val, err := cast.ToFloat32E(v.val)
	if nil != err {
		err = errs.Wrap(err, errs.ErrTypeConversionFailed, "could not convert value '%v' to a float32", v.val)
	}
	return val, nil
}

/*
Float64 implements Valuer.
*/
func (v val) Float64() (float64, error) {
	val, err := cast.ToFloat64E(v.val)
	if nil != err {
		err = errs.Wrap(err, errs.ErrTypeConversionFailed, "could not convert value '%v' to a float64", v.val)
	}
	return val, nil
}

/*
Int implements Valuer.
*/
func (v val) Int() (int, error) {
	val, err := cast.ToIntE(v.val)
	if nil != err {
		err = errs.Wrap(err, errs.ErrTypeConversionFailed, "could not convert value '%v' to an int", v.val)
	}
	return val, nil
}

/*
List implements Valuer.
*/
func (v val) List() (ValueList, error) {
	var err error
	val, ok := v.val.(ValueList)
	if !ok {
		err = errs.New(errs.ErrTypeConversionFailed, "could not convert value '%v' to a ValueList", v.val)
	}
	return val, err
}

/*
Map implements Valuer.
*/
func (v val) Map() (ValueMap, error) {
	var err error
	val, ok := v.val.(ValueMap)
	if !ok {
		err = errs.New(errs.ErrTypeConversionFailed, "could not convert value '%v' to a ValueMap", v.val)
	}
	return val, err
}

/*
String implements Valuer.
*/
func (v val) String() (string, error) {
	val, err := cast.ToStringE(v.val)
	if nil != err {
		err = errs.Wrap(err, errs.ErrTypeConversionFailed, "could not convert value '%v' to a string", v.val)
	}
	return val, err
}

/*
val implements the Value interface
*/
type val struct {
	typ      valType
	val      interface{}
	vBool    bool
	vFloat32 float32
	vFloat64 float64
	vInt     int
	vList    ValueList
	vMap     ValueMap
	vStr     string
}

/*
typeNames maps value types to names for error output
*/
var typeNames = map[valType]string{
	Bool:   "bool",
	Float:  "float",
	Int:    "int",
	List:   "list",
	Map:    "map",
	String: "string",
}
