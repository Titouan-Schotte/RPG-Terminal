package jsonmanagement

func Delete(jsonIn string, key string) {
	data := ReadJson(jsonIn)
	delete(data, key)
	Store(data, jsonIn)
}
