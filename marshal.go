package xmlvalue

import (
	"bytes"
	"fmt"
	"strings"
)

type Opt struct {
	Indent string
}

func (x *V) Marshal(opt ...Opt) ([]byte, error) {
	if nil == x || false == x.initialized {
		return nil, fmt.Errorf("xmlvalue not initiailzed")
	}

	indent := ""
	if len(opt) > 0 {
		indent = opt[0].Indent
	}
	buff := bytes.Buffer{}
	x.toBuffer(&buff, indent, 0)
	return buff.Bytes(), nil
}

func (x *V) toBuffer(buff *bytes.Buffer, indent string, depth int) {
	prefix := ""
	if indent != "" {
		prefix = "\n" + strings.Repeat(indent, depth)
	}
	if depth > 0 {
		buff.WriteString(prefix)
	}
	buff.WriteRune('<')
	writeString(buff, x.name)

	for k, v := range x.attrs {
		buff.WriteRune(' ')
		writeString(buff, k)
		buff.WriteString("=\"")
		writeString(buff, v)
		buff.WriteRune('"')
	}
	buff.WriteRune('>')

	if x.data != nil && len(x.data) > 0 {
		if bytes.ContainsAny(x.data, "<>&\"'") {
			buff.WriteString("<![CDATA[")
			buff.Write(x.data)
			buff.WriteString("]]>")
		} else {
			buff.Write(x.data)
		}
	}

	if len(x.children) > 0 {
		for _, childList := range x.children {
			for _, child := range childList {
				child.toBuffer(buff, indent, depth+1)
			}
		}
		buff.WriteString(prefix)
	}

	buff.WriteString("</")
	writeString(buff, x.name)
	buff.WriteRune('>')
	return
}

// replacer for escaping
var _replacer = strings.NewReplacer(
	"<", "&lt;",
	">", "&gt;",
	"&", "&amp;",
	"'", "&apos;",
	"\"", "&quot;",
)

func writeString(buff *bytes.Buffer, s string) {
	_replacer.WriteString(buff, s)
	return
}
