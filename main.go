package main

type MyStruct struct {
	MyStructPtrField *int
}

func main() {
	localVal := 0
	arr := []MyStruct{{&localVal}}
	for _, val := range arr {
		t := *val.MyStructPtrField
		globalMapInDifferentFile[t] = &val
	}
}
