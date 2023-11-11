package jsonmanagement

func Update(jsonIn string, key string, newValue interface{}) {
	data := ReadJson(jsonIn)
	data[key] = newValue
	Store(data, jsonIn)
}
