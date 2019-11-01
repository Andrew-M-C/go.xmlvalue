package xmlvalue

import (
	"log"
	"strconv"
)

// Get returns child in xml value
func (x *V) Get(params ...interface{}) (*V, error) {
	if false == x.initialized {
		return nil, errNotInitialized
	}
	if 0 == len(params) {
		return nil, errNilParameter
	}
	return x.getOrCreate(false, params...)
}

// GetString returns child text in xml value
func (x *V) GetString(params ...interface{}) (string, error) {
	child, err := x.Get(params...)
	if err != nil {
		return "", err
	}
	return child.Text(), nil
}

// GetInt returns represented integer of text in specified xml value
func (x *V) GetInt(params ...interface{}) (int, error) {
	child, err := x.Get(params...)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(child.Text())
}

func (x *V) getOrCreate(create bool, params ...interface{}) (*V, error) {
	// log.Printf("create %v, params %v", create, params)
	// here, len(params) is not expected to be zero
	l := len(params)
	switch l {
	case 0:
		return x, nil
	case 1:
		return x.getOrCreateWithOneParam(create, params[0])
	case 2:
		return x.getOrCreateWithTwoParams(create, params[0], params[1])
	default:
		// continue below
	}

	_, _, isStr2, err := readParam(params[1])
	if err != nil {
		return nil, err
	}

	if isStr2 {
		child, err := x.getOrCreateWithOneParam(create, params[0])
		if err != nil {
			return nil, err
		}
		return child.getOrCreate(create, params[1:]...)

	}

	child, err := x.getOrCreateWithTwoParams(create, params[0], params[1])
	if err != nil {
		return nil, err
	}
	return child.getOrCreate(create, params[2:]...)
}

func (x *V) getOrCreateWithOneParam(create bool, param interface{}) (*V, error) {
	log.Printf("create %v, param '%v'", create, param)
	_, s, isString, err := readParam(param)
	if err != nil {
		return nil, err
	}
	if false == isString {
		return nil, errParamInvalidParamType
	}
	if "" == s {
		return nil, errEmptyString
	}

	// read first one
	childList, exist := x.children[s]
	if false == exist {
		if false == create {
			return nil, errNotExist
		}
		child := New(s)
		childList = []*V{child}
		x.children[s] = childList
		return child, nil
	}
	return childList[0], nil
}

func (x *V) getOrCreateWithTwoParams(create bool, param1, param2 interface{}) (*V, error) {
	// log.Printf("create %v, param1 %v, param2 %v", create, param1, param2)
	// The first one is not supposed to be a interger
	_, s1, isStr1, err := readParam(param1)
	if err != nil {
		return nil, err
	}
	if false == isStr1 {
		return nil, errParamInvalidParamType
	}
	if "" == s1 {
		return nil, errEmptyString
	}

	// read the second one
	i2, s2, isStr2, err := readParam(param2)
	if err != nil {
		return nil, err
	}
	if isStr2 && s2 == "" {
		return nil, errEmptyString
	}

	// read child
	childList, exist := x.children[s1]
	if false == exist {
		if false == create {
			return nil, errNotExist
		}
		if false == isStr2 && i2 != 0 {
			return nil, errorf("index %d for %s is out of range", i2, s1)
		}
		child := New(s1)
		childList = []*V{child}
		x.children[s1] = childList
		if isStr2 {
			newChild := New(s2)
			child.children[s2] = []*V{newChild}
			return newChild, nil
		}
		return child, nil // to here, it is sure that i2 = 0
	}

	// child exists
	if isStr2 {
		return childList[0].getOrCreateWithOneParam(create, param2)
	}
	// is integer
	if 0 == i2 {
		return childList[0], nil
	}
	if i2 > 0 {
		if i2 < len(childList) {
			return childList[i2], nil
		}
		return nil, errorf("index %d for %s is out of range", i2, s1)
	}
	// if i2 < 0
	if i := len(childList) + i2; i >= 0 {
		return childList[i], nil
	}
	return nil, errorf("index %d for %s is out of range", i2, s1)
}

func readParam(p interface{}) (i int, s string, isString bool, err error) {
	switch p.(type) {
	case int:
		i = p.(int)
	case uint:
		i = int(p.(uint))
	case int8:
		i = int(p.(int8))
	case uint8:
		i = int(p.(uint8))
	case int16:
		i = int(p.(int16))
	case uint16:
		i = int(p.(uint16))
	case int32:
		i = int(p.(int32))
	case uint32:
		i = int(p.(uint32))
	case int64:
		i = int(p.(int64))
	case uint64:
		i = int(p.(uint64))
	case string:
		isString = true
		s = p.(string)
	default:
		err = errParamInvalidParamType
	}
	return
}
