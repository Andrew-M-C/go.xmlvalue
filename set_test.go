package xmlvalue

import (
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"
)

var o = Opt{Indent: "  "}

func getCallerInfo(invokeLevel int) (fileName string, line int, funcName string) {
	funcName = "FILE"
	line = -1
	fileName = "FUNC"

	if invokeLevel <= 0 {
		invokeLevel = 2
	} else {
		invokeLevel++
	}

	pc, fileName, line, ok := runtime.Caller(invokeLevel)
	if ok {
		fileName = filepath.Base(fileName)
		funcName = runtime.FuncForPC(pc).Name()
		funcName = filepath.Ext(funcName)
		funcName = strings.TrimPrefix(funcName, ".")
	}
	// fmt.Println(reflect.TypeOf(pc), reflect.ValueOf(pc))
	return
}

func TestSet(t *testing.T) {
	var err error
	check := func() {
		_, li, fu := getCallerInfo(1)
		if err != nil {
			t.Errorf("%s(), Line %d error: %v", fu, li, err)
		} else {
			t.Logf("%s(), Line %d OK", fu, li)
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

	checkInt := func(i1, i2 int) {
		_, li, fu := getCallerInfo(1)
		if err != nil {
			t.Errorf("%s(), Line %d error: %v", fu, li, err)
			return
		}
		if i1 != i2 {
			t.Errorf("%s(), Line %d: %d != %d", fu, li, i1, i2)
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

	now := time.Now().Unix()

	x := New("xml")
	err = x.SetString("Hello, World!").At("message")
	check()
	err = x.AddString("Hello, XML!").At("message")
	check()
	err = x.SetInt(int(now)).At("time", "unix")
	check()
	err = x.SetString("Hello, XML again!").At("message", 1, "again")
	check()
	err = x.SetString("deep inside").At("A", "A1", "A2", "A3", "A4")
	check()
	err = x.AddString("deep inside B").At("B", "B1", "B2", "B3", "B4")
	check()
	err = x.AddString("deep inside BB").At("B", "B1", "B2", "B3", "B4")
	check()
	err = x.SetString("xmlvalue").At("package", 1)
	shouldError()

	s := x.MustMarshalString(o)
	t.Logf(s)

	// check error sets
	err = x.SetString("deep inside B").At("B", "B1", "B2", "B3", 1, "B4")
	shouldError()
	err = x.SetString("deep inside B").At("B", "B1", "B2", "BB3", 1, "BB4")
	shouldError()
	err = x.SetString("Hello, World!").At("message", 10)
	shouldError()
	err = x.SetString("Hello, World!").At("msg", 1)
	shouldError()

	// check each elements
	s, err = x.GetString("message")
	checkStr(s, "Hello, World!")
	s, err = x.GetString("message", 1)
	checkStr(s, "Hello, XML!")
	s, err = x.GetString("message", 1, "again")
	checkStr(s, "Hello, XML again!")
	i, err := x.GetInt("time", "unix")
	checkInt(i, int(now))
	s, err = x.GetString("A", "A1", "A2", "A3", "A4")
	checkStr(s, "deep inside")
	s, err = x.GetString("A", "A1", "A2", "A3", 0, "A4")
	checkStr(s, "deep inside")

	s, err = x.GetString("message", 0)
	checkStr(s, "Hello, World!")
	s, err = x.GetString("message", uint(0))
	checkStr(s, "Hello, World!")
	s, err = x.GetString("message", int8(0))
	checkStr(s, "Hello, World!")
	s, err = x.GetString("message", uint8(0))
	checkStr(s, "Hello, World!")
	s, err = x.GetString("message", int16(0))
	checkStr(s, "Hello, World!")
	s, err = x.GetString("message", uint16(0))
	checkStr(s, "Hello, World!")
	s, err = x.GetString("message", int32(0))
	checkStr(s, "Hello, World!")
	s, err = x.GetString("message", uint32(0))
	checkStr(s, "Hello, World!")
	s, err = x.GetString("message", int64(0))
	checkStr(s, "Hello, World!")
	s, err = x.GetString("message", uint64(0))
	checkStr(s, "Hello, World!")
	s, err = x.GetString("message", 1)
	checkStr(s, "Hello, XML!")
	s, err = x.GetString("message", -1)
	checkStr(s, "Hello, XML!")

	// check errors
	_, err = new(V).GetString()
	shouldError()
	_, err = x.GetString()
	shouldError()
	_, err = x.GetString("msg")
	shouldError()
	_, err = x.GetInt("msg")
	shouldError()
	_, err = x.GetString(0)
	shouldError()
	_, err = x.GetString("")
	shouldError()
	_, err = x.GetString(false)
	shouldError()
	_, err = x.GetString(0, "string")
	shouldError()
	_, err = x.GetString("string", "")
	shouldError()
	_, err = x.GetString("", "string")
	shouldError()
	_, err = x.GetString(false, "string", "string")
	shouldError()
	_, err = x.GetString("string", "", "string")
	shouldError()
	_, err = x.GetString("string", false)
	shouldError()
	_, err = x.GetString(false, "string")
	shouldError()
	_, err = x.GetString("string", false, "string")
	shouldError()
	_, err = x.GetString("string", "string", false)
	shouldError()
	_, err = x.GetString("string", 0, false)
	shouldError()
	_, err = x.GetString("string", "string", false, "string")
	shouldError()
	_, err = x.GetString("string", "string", false, "string", "string")
	shouldError()
	_, err = x.GetString("message", 10, "string")
	shouldError()
	_, err = x.GetString("message", 10)
	shouldError()
	_, err = x.GetString("message", -10)
	shouldError()
	_, err = x.GetString("msg", 1, "string")
	shouldError()

	// done
	return
}
