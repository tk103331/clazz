package tools

import "testing"

func AssertEqual(t *testing.T, a, b interface{}) {
	t.Helper()
	if a != b {
		t.Errorf("Not Equal. %d %d", a, b)
	}
}

func AssertTrue(t *testing.T, value bool) {
	t.Helper()
	if value {
		t.Errorf("Not True.")
	}
}
func AssertFalse(t *testing.T, value bool) {
	t.Helper()
	if value {
		t.Errorf("Not False.")
	}
}

func AssertError(t *testing.T, e error) {
	t.Helper()
	if e == nil {
		t.Errorf("No Error : %s", e.Error())
	}
}
func AssertNoErr(t *testing.T, e error) {
	t.Helper()
	if e != nil {
		t.Errorf("Error : %s", e.Error())
	}
}

func AssertNotNil(t *testing.T, value interface{}) {
	t.Helper()
	if value == nil {
		t.Errorf("UTF8Value is nil ")
	}
}
