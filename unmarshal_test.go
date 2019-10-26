package xmlvalue

import (
	"testing"
)

func TestNew(t *testing.T) {
	raw := `
	<xml>
		<message>Hello, XML!</message>
		<message>Hello, another XML!</message>
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

	return
}
