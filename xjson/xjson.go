package xjson

import (
	"encoding/json"
	"fmt"
)

// Json2List 解析json字符串为切片
func Json2List[T interface{}](jsonData string) (*[]T, error) {
	var data []T
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// Json2Struct 解析json字符串为结构体
func Json2Struct[T interface{}](jsonData string) (*T, error) {
	// 定义一个变量用于存储解析后的数据
	var items T
	// 使用json.Unmarshal进行解析
	err := json.Unmarshal([]byte(jsonData), &items)
	if err != nil {
		fmt.Println("解析JSON时发生错误:", err)
		return nil, err
	}
	return &items, err
}

// Json2StructList 解析json字符串为结构体切片
func Json2StructList[T interface{}](jsonData string) (*[]T, error) {
	// 定义一个切片用于存储解析后的数据
	var items []T
	// 使用json.Unmarshal进行解析
	err := json.Unmarshal([]byte(jsonData), &items)
	if err != nil {
		fmt.Println("解析JSON时发生错误:", err)
		return nil, err
	}
	return &items, err
}
