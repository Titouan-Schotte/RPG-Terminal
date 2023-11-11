package jsonmanagement

func Add(jsonIn string, key string, value interface{}) {
	data := ReadJson(jsonIn)
	data[key] = value
	Store(data, jsonIn)
}
