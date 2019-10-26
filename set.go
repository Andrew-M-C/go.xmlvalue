package xmlvalue

import (
	"strconv"
)

// Set structures is a internal structure for SetXxx functions
type Set struct {
	initialized bool
	add         bool
	x           *V
	s           string
}

// SetString starts setting a string to specified path
func (x *V) SetString(s string) *Set {
	return &Set{
		initialized: true,
		add:         false,
		x:           x,
		s:           s,
	}
}

// AddString starts adding a string to specified path
func (x *V) AddString(s string) *Set {
	return &Set{
		initialized: true,
		add:         true,
		x:           x,
		s:           s,
	}
}

// SetInt starts adding a string to specified path
func (x *V) SetInt(i int) *Set {
	return &Set{
		initialized: true,
		add:         false,
		x:           x,
		s:           strconv.Itoa(i),
	}
}

// AddInt starts adding a string to specified path
func (x *V) AddInt(i int) *Set {
	return &Set{
		initialized: true,
		add:         true,
		x:           x,
		s:           strconv.Itoa(i),
	}
}

// At sets string to specified path
func (set *Set) At(name string, names ...string) {
	if false == set.initialized {
		return
	}

	if l := len(names); l == 0 {
		child := New(name)
		child.SetText([]byte(set.s))
		if set.add {
			set.x.addChild(child)
		} else {
			set.x.setChild(child)
		}

	} else {
		child := New(names[l-1])
		child.SetText([]byte(set.s))
		if set.add {
			set.x.getChildOrCreate(name).addChild(child, names[:l-1]...)
		} else {
			set.x.getChildOrCreate(name).setChild(child, names[:l-1]...)
		}
	}

	return
}
