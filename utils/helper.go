package utils

import (
	"crypto/md5"
	"encoding/base64"
	"reflect"
	"unsafe"
)

//Hasher - Calculates the hash of an interface
func Hasher(data interface{}) string {
	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Ptr {
		if !v.CanAddr() {
			return ""
		}

		v = v.Addr()
	}
	size := unsafe.Sizeof(v.Interface())
	b := (*[1 << 10]uint8)(unsafe.Pointer(v.Pointer()))[:size:size]

	h := md5.New()
	return base64.StdEncoding.EncodeToString(h.Sum(b))
}
