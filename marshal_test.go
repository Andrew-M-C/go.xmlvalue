package xmlvalue

import (
	"testing"
)

func TestMarshal(t *testing.T) {
	var err error
	check := func() {
		_, li, fu := getCallerInfo(1)
		if err != nil {
			t.Errorf("%s(), Line %d error: %v", fu, li, err)
		} else {
			t.Logf("%s(), Line %d OK", fu, li)
		}
	}

	shouldError := func() {
		_, li, fu := getCallerInfo(1)
		if err == nil {
			t.Errorf("%s(), Line %d: error expected but not caught", fu, li)
		} else {
			t.Logf("%s(), Line %d: expected error: %v", fu, li, err)
		}
	}

	checkStr := func(s1, s2 string) {
		_, li, fu := getCallerInfo(1)
		if err != nil {
			t.Errorf("%s(), Line %d error: %v", fu, li, err)
			return
		}
		if s1 != s2 {
			t.Errorf("%s(), Line %d: '%s' != '%s'", fu, li, s1, s2)
		}
	}

	// nil object
	x := &V{}
	_, err = x.Marshal()
	shouldError()
	_, err = x.MarshalString()
	shouldError()

	func() {
		defer func() {
			if err := recover(); err == nil {
				t.Errorf("error expected but not caught")
			}
		}()
		x.MustMarshalString()
	}()

	// CDATA
	x = New("xml")
	x.SetString("<").At("data")
	s := x.MustMarshalString()
	t.Logf("str: %s", s)

	x, err = UnmarshalString(s)
	check()
	s, _ = x.GetString("data")
	checkStr(s, "<")

	s = x.Name()
	checkStr(s, "xml")

	err = x.SetText(12345)
	shouldError()

	err = x.SetText([]byte("hello"))
	check()

	err = x.SetText(nil)
	check()

	// misc error cases
	s, exist := x.GetAttr("")
	if exist {
		t.Errorf("empty attr should not exist")
		return
	}
	checkStr(s, "")

	return
}
