package format

import (
	"reflect"
	"testing"
)

func TestReformattingSingleNamedPattern(t *testing.T) {
	pat := "%<name_me>x"

	f, n := reformat(pat)

	if f != "%x" {
		t.Errorf("pattern should be %%x but %v", f)
	}

	if !reflect.DeepEqual(n, []string{"name_me"}) {
		t.Errorf("named var should be {name_me} but %v", n)
	}
}

func TestReformattingMultipleNamedPattern(t *testing.T) {
	pat := "%<name_me>x and %<another_name>v"

	f, n := reformat(pat)

	if f != "%x and %v" {
		t.Errorf("pattern should be %%x and %%v but %v", f)
	}

	if !reflect.DeepEqual(n, []string{"name_me", "another_name"}) {
		t.Errorf("named var should be {name_me, another_name} but %v", n)
	}
}
