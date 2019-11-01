package xmlvalue

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
)

type stack struct {
	L []*V
}

func newStack() *stack {
	ret := stack{
		L: make([]*V, 0, 0),
	}
	return &ret
}

func (s *stack) Push(i *V) {
	if i != nil {
		s.L = append(s.L, i)
	}
}

func (s *stack) Pop() *V {
	l := len(s.L)
	if l <= 0 {
		return nil
	}
	ret := s.L[l-1]
	s.L = s.L[0 : l-1]
	return ret
}

// Unmarshal unmarshals a given data and returns a xmlvalue object
func Unmarshal(b []byte) (*V, error) {
	if nil == b || 0 == len(b) {
		return nil, fmt.Errorf("nil input")
	}

	decoder := xml.NewDecoder(bytes.NewReader(b))
	stk := newStack()
	var curr *V
	var root *V

	for {
		t, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				// log.Printf("Parse XML finished!")
				return root, nil
			}
			// log.Printf("Failed to Parse XML with the error of %v", err)
			return nil, err
		}
		t = xml.CopyToken(t)

		switch t := t.(type) {
		case xml.StartElement:
			// log.Printf("xml.StartElement")
			name := t.Name.Local
			item := New(name)
			// log.Printf("name: %s", name)

			for _, a := range t.Attr {
				// log.Printf("attr: %s - %s", a.Name.Local, a.Value)
				item.SetAttr(a.Name.Local, a.Value)
			}
			if curr != nil {
				curr.addChild(item)
				stk.Push(curr)
			} else {
				root = item
			}
			curr = item

		case xml.EndElement:
			// log.Printf("xml.EndElement")
			curr = stk.Pop()

		case xml.CharData:
			b := []byte(t)
			b = bytes.Trim(b, "\r\n\t ")
			if len(b) > 0 {
				// log.Printf("xml.CharData: '%s'", string(b))
				if curr != nil {
					curr.data = b
				}
			}

		case xml.Comment:
			// ignore
		}
	}

	// never reach here
	// return nil, fmt.Errorf("format error")
}
