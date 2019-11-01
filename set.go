package xmlvalue

import (
	"strconv"
)

// S structures is a internal structure for SetXxx functions
type S struct {
	initialized bool
	add         bool
	x           *V
	s           string
}

// SetString starts setting a string to specified path
func (x *V) SetString(s string) *S {
	return &S{
		initialized: true,
		add:         false,
		x:           x,
		s:           s,
	}
}

// AddString starts adding a string to specified path
func (x *V) AddString(s string) *S {
	return &S{
		initialized: true,
		add:         true,
		x:           x,
		s:           s,
	}
}

// SetInt starts adding a string to specified path
func (x *V) SetInt(i int) *S {
	return &S{
		initialized: true,
		add:         false,
		x:           x,
		s:           strconv.Itoa(i),
	}
}

// AddInt starts adding a string to specified path
func (x *V) AddInt(i int) *S {
	return &S{
		initialized: true,
		add:         true,
		x:           x,
		s:           strconv.Itoa(i),
	}
}

// At sets string to specified path
func (set *S) At(params ...interface{}) error {
	if false == set.initialized {
		return errNotInitialized
	}
	paramCount := len(params)
	if 0 == paramCount {
		return errNilParameter
	}

	// only one parameter, this should be a string
	if paramCount == 1 {
		_, k, isStr, err := readParam(params[0])
		if false == isStr {
			return err
		}
		if false == isStr {
			return errParamInvalidParamType
		}
		if "" == k {
			return errEmptyString
		}
		newChild := New(k)
		newChild.SetText(set.s)
		newChildList, exist := set.x.children[k]
		if false == exist || false == set.add {
			newChildList = []*V{newChild}
		} else {
			newChildList = append(newChildList, newChild)
		}
		set.x.children[k] = newChildList
		return nil
	}

	// mutiple parameter, now we needs some operation
	lastI, lastS, lastIsStr, err := readParam(params[paramCount-1])
	if err != nil {
		return err
	}
	if lastIsStr {
		if "" == lastS {
			return errEmptyString
		}
		child, err := set.x.getOrCreate(false, params[:paramCount-1]...)
		if err != nil {
			if err != errNotExist {
				return err
			}
			child, err = set.x.getOrCreate(true, params[:paramCount-1]...)
		}
		if err != nil {
			return err
		}
		newChild := New(lastS)
		newChild.SetText(set.s)
		newChildList, exist := child.children[lastS]
		if false == exist || false == set.add {
			newChildList = []*V{newChild}
		} else {
			newChildList = append(newChildList, newChild)
		}
		child.children[lastS] = newChildList

		return nil
	}

	// last param is an interger, now we should search for prev level
	if set.add {
		// in add mode, last parameter should be an integer
		return errParamInvalidParamType
	}

	_, prevS, prevIsStr, err := readParam(params[paramCount-2])
	if err != nil {
		return err
	}
	if false == prevIsStr {
		return errParamInvalidParamType
	}
	if "" == prevS {
		return errEmptyString
	}

	newChild := New(prevS)
	newChild.SetText(set.s)
	child, err := set.x.getOrCreate(false, params[:paramCount-2]...)
	if err != nil {
		if err != errNotExist {
			return err
		}
		if lastI != 0 {
			return errOutOrRange
		}
		child, err = set.x.getOrCreate(true, params[:paramCount-2]...)
		if err != nil {
			return err
		}
		child.children[prevS] = []*V{newChild}
		return nil
	}

	childList, exist := child.children[prevS]
	if false == exist {
		if 0 != lastI {
			return errOutOrRange
		}
		child.children[prevS] = []*V{newChild}
		return nil
	}

	if lastI >= 0 {
		if lastI >= len(childList) {
			return errOutOrRange
		}
		childList[lastI] = newChild
	} else {
		lastI += len(childList)
		if lastI < 0 {
			return errOutOrRange
		}
		childList[lastI] = newChild
	}

	return nil
}
