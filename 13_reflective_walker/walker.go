package reflectivewalker

import(
	"reflect"
)

func Walker(targetObject interface{}, targetFunc func(string)) {

	val := getValue(targetObject)

	walkValue := func(value reflect.Value) {
		Walker(value.Interface(), targetFunc)
	}

	numberOfValues := 0
	var getField func(int) reflect.Value

	switch val.Kind() {

	case reflect.Struct:
		numberOfValues = val.NumField()
		getField = val.Field

	case reflect.Slice, reflect.Array:
		numberOfValues = val.Len()
		getField = val.Index

	case reflect.String:
		targetFunc(val.String())
	
	case reflect.Map:
		for _, key := range val.MapKeys() {
			Walker(val.MapIndex(key).Interface(), targetFunc)
		}

	case reflect.Chan:
		for v, ok := val.Recv(); ok; v, ok = val.Recv() {
			walkValue(v)
		}

	case reflect.Func:
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walkValue(res)
		}
	}

	for i := 0; i < numberOfValues; i++ {
		field := getField(i)
		walkValue(field)
	}

}

func getValue(targetObject interface{}) reflect.Value {
	val := reflect.ValueOf(targetObject)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	return val
}
