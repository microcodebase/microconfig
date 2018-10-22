package microconfig_test

import "testing"

func TestEmpty(t *testing.T) {
	conf := Parse([]byte(``))
	if conf[""] != "" {
		t.Fail()
	}
}
func TestNoSpace(t *testing.T) {
	conf := Parse([]byte(`a=1`))
	if conf["a"] != "1" {
		t.Fail()
	}
}
func TestLeftSpace(t *testing.T) {
	conf := Parse([]byte(`   a=1`))
	if conf["a"] != "1" {
		t.Fail()
	}
}
func TestRightSpace(t *testing.T) {
	conf := Parse([]byte(`a=1    `))
	if conf["a"] != "1" {
		t.Fail()
	}
}
func TestInterSpace(t *testing.T) {
	conf := Parse([]byte(`a = 1`))
	if conf["a"] != "1" {
		t.Fail()
	}
}
func TestComment(t *testing.T) {
	conf := Parse([]byte(`#ertret
	a = 1
	#sdfgfgdfsg`))
	if conf["a"] != "1" {
		t.Fail()
	}
}
func TestComplexValue(t *testing.T) {
	conf := Parse([]byte(`abc = 123 789`))
	if conf["abc"] != "123 789" {
		t.Fail()
	}
}
