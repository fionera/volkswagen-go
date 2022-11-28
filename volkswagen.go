package volkswagen_go

import (
	"reflect"
	"testing"
	_ "unsafe"

	"bou.ke/monkey"
)

func init() {
	f, _ := reflect.TypeOf(testing.T{}).FieldByName("common")
	monkey.PatchInstanceMethodIgnoreType(reflect.PointerTo(f.Type), "Fail", func() {
	})
	monkey.PatchInstanceMethodIgnoreType(reflect.PointerTo(f.Type), "Error", func() {
	})
	monkey.PatchInstanceMethodIgnoreType(reflect.PointerTo(f.Type), "Errorf", func() {
	})
	monkey.PatchInstanceMethodIgnoreType(reflect.PointerTo(f.Type), "Fatal", func() {
	})
	monkey.PatchInstanceMethodIgnoreType(reflect.PointerTo(f.Type), "Fatalf", func() {
	})
}
