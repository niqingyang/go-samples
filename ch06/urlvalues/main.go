package main

import "fmt"

type Values map[string][]string

func (v Values) Get(key string) string {
	if vs := v[key]; len(vs) > 0 {
		return vs[0]
	}
	return ""
}

func (v Values) Add(key, value string) {
	v[key] = append(v[key], value)
}

func main() {
	m := Values{"lang": {"en"}} // 直接构造
	m.Add("item", "1")
	m.Add("item", "2")


	fmt.Println(m.Get("lang"))
	fmt.Println(m.Get("q"))
	fmt.Println(m.Get("item"))
	fmt.Println(m["item"])

	m = nil

	fmt.Println(m.Get("item")) // ""
	m.Add("item", "3")         // 宕机：赋值给空的map类型
}
