package jsonmanagement

func Get(jsonIn string, key string) interface{} {
	return ReadJson(jsonIn)[key]
}
