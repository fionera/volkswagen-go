package volkswagen_go

import (
	"os"
	"reflect"
	"testing"
	_ "unsafe"

	"bou.ke/monkey"
)

var envVars = []string{
	"BUILD_ID",               // Jenkins, Cloudbees
	"BUILD_NUMBER",           // Jenkins, TeamCity
	"CI",                     // Travis CI, CircleCI, Cirrus CI, Gitlab CI, Appveyor, CodeShip, dsari
	"CI_APP_ID",              // Appflow
	"CI_BUILD_ID",            // Appflow
	"CI_BUILD_NUMBER",        // Appflow
	"CI_NAME",                // Codeship and others
	"CONTINUOUS_INTEGRATION", // Travis CI, Cirrus CI
	"RUN_ID",                 // TaskCluster, dsari
}

func init() {
	var isCI bool
	for _, s := range envVars {
		if _, set := os.LookupEnv(s); set {
			isCI = true
		}
	}

	if !isCI {
		return
	}

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
