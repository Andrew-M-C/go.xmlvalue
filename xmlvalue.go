// Package xmlvalue is for XML parsing. It is used in situations those Go structures cannot achieve well. Most of the usages are quite simular as jsonvalue. (https://github.com/Andrew-M-C/go.jsonvalue)
//
// This package is particularly useful in following situations:
//
// 1. Rapidly create an XML document with mutiple level.
//
// 2. Parsing XML documents with volatile structures.
package xmlvalue

// V represents an XML value
type V struct {
	initialized bool
	name        string
	data        []byte
	attrs       map[string]string
	children    map[string][]*V
}

// New returns a initialized XML value with given name
func New(name string) *V {
	return &V{
		initialized: true,
		name:        name,
		data:        []byte{},
		attrs:       map[string]string{},
		children:    map[string][]*V{},
	}
}

// Name returns the name of this XML element
func (x *V) Name() string {
	return x.name
}

// SetName set the name of this XML value
// func (x *V) SetName(name string) {
// 	if name != "" {
// 		x.name = name
// 	}
// 	return
// }

// Text returns the text string of this XML element
func (x *V) Text() string {
	if nil == x || nil == x.data {
		return ""
	}
	return string(x.data)
}

// SetText set raw text of this XML element. Both string and []byte type are accepted
func (x *V) SetText(t interface{}) error {
	switch t.(type) {
	case nil:
		x.data = []byte{}
	case []byte:
		x.data = t.([]byte)
	case string:
		x.data = []byte(t.(string))
	default:
		return errParamInvalidParamType
	}
	return nil
}

// GetAttr read attribute. If a default value given, the default attribute will be set and returned when the attribute does not exist.
func (x *V) GetAttr(a string, defaultValue ...string) (attr string, exist bool) {
	if "" == a {
		return "", false
	}
	attr, exist = x.attrs[a]
	if exist {
		return
	}
	if len(defaultValue) > 0 {
		return defaultValue[0], false
	}
	return "", false
}

// SetAttr sets attribute
func (x *V) SetAttr(a, v string) {
	x.attrs[a] = v
	return
}

// addChild add a child
func (x *V) addChild(c *V) {
	childList, exist := x.children[c.name]
	if false == exist {
		childList = []*V{c}
		x.children[c.name] = childList
	} else {
		childList = append(childList, c)
		x.children[c.name] = childList
	}

	// done
	return
}
