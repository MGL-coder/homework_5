package cyrillicFilter

import (
	"reflect"
	"regexp"
)

func cyrFilterHelper(s reflect.Value) {
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)

		for ; f.Kind() == reflect.Ptr; f = f.Elem() {
		} // dereferencing the field if it is a pointer

		if f.Kind() == reflect.String && f.CanSet() {
			cyrReg := regexp.MustCompile("\\p{Cyrillic}")
			newStr := cyrReg.ReplaceAllString(f.String(), "")
			f.SetString(newStr)
		} else if f.Kind() == reflect.Struct {
			cyrFilterHelper(f)
		}
	}
}

func CyrFilter(any interface{}) {
	cyrFilterHelper(reflect.ValueOf(any).Elem())
}
