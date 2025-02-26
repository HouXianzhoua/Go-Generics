
package sliceT


func SliceContains[T comparable](slice []T, value T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func SliceAdd[T any](slice []T, value T) []T {
	return append(slice, value)
}

func SliceRemove[T comparable](slice []T, value T) []T {
	for i, v := range slice {
		if v == value {
			slice = append(slice[:i], slice[i+1:]...)
			return slice
		}
	}
	return slice
}

func SliceRemoveAt[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}

func SliceIndexOf[T comparable](slice []T, value T) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

func SliceLastIndexOf[T comparable](slice []T, value T) int {	
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == value {
			return i
		}
	}
	return -1
}


func SliceFilter[T any](slice []T, filter func(T) bool) []T {	
	var result []T
	for _, v := range slice {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}

func SliceMap[T any, U any](slice []T, mapper func(T) U) []U {
	var result []U
	for _, v := range slice {
		result = append(result, mapper(v))
	}
	return result
}

func SliceReduce[T any, U any](slice []T, reducer func(U, T) U, initialValue U) U {
	result := initialValue
	for _, v := range slice {
		result = reducer(result, v)
	}
	return result
}

// ShrinkSlice 按比例动态缩容函数
func ShrinkSlice[T any](slice []T) []T {
    length := len(slice)
    var newCap int

    // 根据切片长度动态确定新容量
    switch {
    case length < 10:
        // 切片长度小于 10 时，新容量为长度的 2 倍
        newCap = length * 2
    case length < 100:
        // 切片长度在 10 到 100 之间时，新容量为长度的 1.5 倍
        newCap = int(float64(length) * 1.5)
    default:
        // 切片长度大于等于 100 时，新容量等于长度
        newCap = length
    }

    // 如果原容量大于新容量，则进行缩容
    if cap(slice) > newCap {
        newSlice := make([]T, length, newCap)
        copy(newSlice, slice)
        return newSlice
    }
    return slice
}


// ShrinkSlice 支持缩容的泛型切片删除函数
func ShrinkSlice_Index[T any](slice []T, index int) []T {
    if index < 0 || index >= len(slice) {
        return slice
    }
    // 删除指定索引的元素
    copy(slice[index:], slice[index+1:])
    slice = slice[:len(slice)-1]
    return ShrinkSlice(slice)
}

// ShrinkSlice 支持缩容的泛型切片删除函数
func ShrinkSlice_Value[T comparable](slice []T, value T) []T {

	// 删除指定元素value
    slice=SliceRemove(slice, value)

    return ShrinkSlice(slice)
}

