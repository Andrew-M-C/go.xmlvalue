package xmlvalue

import (
	"testing"
)

func TestNew(t *testing.T) {
	raw := `
	<?xml version="1.0" encoding="utf-8"?>
	<xml>
		<message first="yes">Hello, XML!</message>
		<message first="no">Hello, another XML!</message>
	</xml>
	`

	x, err := Unmarshal([]byte(raw))
	if err != nil {
		t.Errorf("unmarshal failed: %v", err)
		return
	}

	b, err := x.Marshal()
	if err != nil {
		t.Errorf("marshal failed: %v", err)
		return
	}
	t.Logf("got marshaled result: %s", string(b))

	b, err = x.Marshal(Opt{Indent: "\t"})
	if err != nil {
		t.Errorf("marshal failed: %v", err)
		return
	}
	t.Logf("got marshaled result:\n%s", string(b))

	c, _ := x.Get("message", 0)
	a, exist := c.GetAttr("first")
	if false == exist {
		t.Errorf("attribute 'first' not exists")
		return
	}
	if a != "yes" {
		t.Errorf("unexpected attribute: '%s'", a)
	}

	c, _ = x.Get("message", 1)
	a, exist = c.GetAttr("first")
	if false == exist {
		t.Errorf("attribute 'first' not exists")
		return
	}
	if a != "no" {
		t.Errorf("unexpected attribute: '%s'", a)
		return
	}

	s := new(V).Text()
	if s != "" {
		t.Errorf("unexpected text: %v", s)
		return
	}

	a, exist = c.GetAttr("second", "undefined")
	if exist {
		t.Errorf("attr 'second' is not supposed to exist")
		return
	}
	if a != "undefined" {
		t.Errorf("unexpected default attr: '%s'", a)
		return
	}

	a, exist = c.GetAttr("second")
	if exist {
		t.Errorf("attr 'second' is not supposed to exist")
		return
	}
	if a != "" {
		t.Errorf("unexpected default attr: '%s'", a)
		return
	}

	_, err = Unmarshal(nil)
	if err == nil {
		t.Errorf("error expected but uncaught")
		return
	}

	return
}
