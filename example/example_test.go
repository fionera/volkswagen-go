package example

import (
	"fmt"
	"testing"
)
import _ "volkswagen-go"

func TestExample(t *testing.T) {
	t.Fail()
}

func TestParallel(t *testing.T) {
	t.Parallel()
	for i := 0; i < 10; i++ {
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			t.Fail()
		})
	}
}

func TestFatalf(t *testing.T) {
	t.Fatalf("foobar")
}

func TestFatal(t *testing.T) {
	t.Fatal("foobar")
}
