package main

import (
	"fmt"
	"reflect"
)

// Comparator - интерфейс компаратора, модифицировать нельзя
type Comparator interface {
	Compare(a, b interface{}) bool
}

func wrapper(c Comparator, a, b interface{}) bool {
	verdict := c.Compare(a, b)
	fmt.Println(a, b, verdict)

	return verdict
}

type ci struct{}

// Compare - реализация компаратора, которую надо поправить
func (c *ci) Compare(a, b interface{}) bool {
	if valA, valB := reflect.ValueOf(a), reflect.ValueOf(b); valA.Kind() == reflect.Pointer && valB.Kind() == reflect.Pointer {
		return valA.Pointer() == valB.Pointer()
	}

	return reflect.DeepEqual(a, b)
}

type T0 int
type T1 int

func main() {
	var value0 T0
	var value1 T1

	ptr0a := (*T0)(&value0)
	ptr1a := (*T1)(ptr0a)

	ptr0b := (*T0)(&value1)
	ptr1b := (*T1)(ptr0b)

	f := &ci{}

	v0 := wrapper(f, ptr0a, ptr1a)
	v1 := wrapper(f, ptr0b, ptr1b)
	v2 := wrapper(f, ptr0a, ptr0b)
	v3 := wrapper(f, ptr1a, ptr1b)

	if !(v0 || v1) || (v2 || v3) {
		panic("failed")
	}

}
