package main

import "fmt"

// MapMap 泛型函数，对映射进行转换操作
func MapMap[K comparable, V any, U any](m map[K]V, mapper func(K, V) U) map[K]U {
    result := make(map[K]U)
    for k, v := range m {
        result[k] = mapper(k, v)
    }
    return result
}

func main() {
    // 创建一个字符串到整数的映射
    m := map[string]int{
        "apple":  1,
        "banana": 2,
        "cherry": 3,
    }

    // 定义一个映射函数，将整数值转换为对应的字符串长度
    mapper := func(k string, v int) string {
        return k[:v]
    }

    // 调用 MapMap 函数进行转换
    newMap := MapMap(m, mapper)

    // 打印转换后的映射
    fmt.Println(newMap)
}