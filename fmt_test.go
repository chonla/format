package format

import (
	"fmt"
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

func TestReformattingRepeatedNamedPattern(t *testing.T) {
	pat := "%<name_me>x and %<another_name>v and %<name_me>v"

	f, n := reformat(pat)

	if f != "%x and %v and %v" {
		t.Errorf("pattern should be %%x and %%v and %%v but %v", f)
	}

	if !reflect.DeepEqual(n, []string{"name_me", "another_name", "name_me"}) {
		t.Errorf("named var should be {name_me, another_name, name_me} but %v", n)
	}
}

func TestSprintf(t *testing.T) {
	pat := "%<brother>s loves %<sister>s. %<sister>s also loves %<brother>s."
	params := map[string]interface{}{
		"sister":  "Susan",
		"brother": "Louis",
	}

	s := Sprintf(pat, params)

	if s != "Louis loves Susan. Susan also loves Louis." {
		t.Errorf("result should be Louis loves Susan. Susan also love Louis. but %v", s)
	}
}

func TestSprintfln(t *testing.T) {
	pat := "%<brother>s loves %<sister>s. %<sister>s also loves %<brother>s."
	params := map[string]interface{}{
		"sister":  "Susan",
		"brother": "Louis",
	}

	s := Sprintfln(pat, params)

	if s != "Louis loves Susan. Susan also loves Louis."+fmt.Sprintln() {
		t.Errorf("result should be Louis loves Susan. Susan also love Louis"+fmt.Sprintln()+". but %v", s)
	}
}

func TestSprintfFloatsWithPrecision(t *testing.T) {
	pat := "%<float>f / %<floatprecision>.1f / %<long>g / %<longprecision>.3g"
	params := map[string]interface{}{
		"float":          5.034560,
		"floatprecision": 5.03456,
		"long":           5.03456,
		"longprecision":  5.03456,
	}

	s := Sprintf(pat, params)

	expectedresult := "5.034560 / 5.0 / 5.03456 / 5.03"
	if s != expectedresult {
		t.Errorf("result should be (%v) but is (%v)", expectedresult, s)
	}
}

func BenchmarkSprintln(b *testing.B) {
	pat := "%<brother>s loves %<sister>s. %<sister>s also loves %<brother>s."
	params := map[string]interface{}{
		"sister":  "Susan",
		"brother": "Louis",
	}

	for i := 0; i < b.N; i++ {
		_ = Sprintfln(pat, params)
	}
}
