package xmlvalue

// V represents a XML value
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
		return errNilParameter
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
func (x *V) addChild(c *V, path ...string) {
	if false == x.initialized {
		return
	}
	if nil == c || false == c.initialized || "" == c.name {
		// do nothing
		return
	}

	// find children
	curr := x
	for _, p := range path {
		curr = curr.getChildOrCreate(p)
	}

	// set child
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

func (x *V) getChildOrCreate(name string) *V {
	childList, exist := x.children[name]
	if false == exist || len(childList) == 0 {
		c := New(name)
		childList = []*V{c}
		x.children[name] = childList
		return c
	}
	return childList[0]
}

// setChild
func (x *V) setChild(c *V, path ...string) {
	if false == x.initialized {
		return
	}
	if nil == c || false == c.initialized || "" == c.name {
		// do nothing
		return
	}

	// find children
	curr := x
	for _, p := range path {
		curr = curr.getChildOrCreate(p)
	}

	// set child
	x.children[c.name] = []*V{c}

	// done
	return
}
