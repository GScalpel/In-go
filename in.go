package in_gs

import (
	"errors"
	"reflect"
)

// 对比两个变量，判断一个变量dst的元素是否属于另一个变量src
func In(dst, src interface{}) (bool,error) {
	srcIter, success := takeIterArg(src)
	if !success {
		return false, errors.New("src is not a slice or array")
	}
	switch dst.(type) {
	case int, string:{
		for _, i := range srcIter{
			if i == dst{
				return true, nil
			}
		}
		return false, nil
	}
	default:{
		dstIter,success := takeIterArg(dst)
		if !success {
			return false, errors.New("dst is not a slice or array")
		}
		for _, i := range dstIter{
			judge := false
			for _, j := range srcIter{
				if i == j {
					judge = true
					continue
				}
			}
			if judge == false{
				return false, nil
			}
		}
		return true, nil
	}
	}
}
// 将arg 从一个interface{}不可迭代对象变为一个可迭代对象
func takeIterArg(arg interface{}) (out []interface{}, ok bool) {
	refVal, success := takeArg(arg, reflect.Slice)
	if !success {
		if refVal, success = takeArg(arg, reflect.Array); !success {
			return
		}
	}
	length := refVal.Len()
	out = make([]interface{}, length)
	for i:=0; i<length; i++ {
		out[i] = refVal.Index(i).Interface()  // refVal是reflect.Value类型，
							                 // 只能通过Index来取具体的值,值的type为value，Interface将其转换为具体的某种类型
	}
	return out, true
}

// 将arg转换为reflect.Value类型
// val的kind和传入的kind（reflect.Slice）对比 相同返回true (确保参数转换为指定的reflect kind类型)
func takeArg(arg interface{}, kind reflect.Kind) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(arg)
	if val.Kind() == kind {
		ok = true
	}
	return
}
