package fuzz

import (
	"log"
	"math/rand"
	"reflect"
	"strconv"
	"time"
	"unsafe"
)

const (
	DefaultSliceSize    int = 20
	DefaultMapSize      int = 10
	DefaultStringLength int = 50
)

func randomString(strlen int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

var Unsupported = []reflect.Kind{
	reflect.Interface,
	reflect.Chan,
	reflect.Func,
	reflect.Complex64,
	reflect.Complex128,
	reflect.UnsafePointer,
}

func isSupportedKind(k reflect.Kind) bool {
	for _, v := range Unsupported {
		if k == v {
			return false
		}
	}
	return true
}

func Fuzz(v reflect.Value, tag string) {
	// check valid first
	if !v.IsValid() {
		log.Printf("Not valid")
		return
	}

	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			if v.CanSet() {
				e := v.Type().Elem()
				v.Set(reflect.New(e))
				v = v.Elem()
			} else {
				// TODO
				// since v is nil value, v.Elem() will be zero value
				// and zero value is not addressable or settable, how
				// can we create new data for underlining data?
				log.Printf("v is nil value and v.Elem() is zero value which is not settable and addressable\n")
				return
			}
		} else {
			v = v.Elem()
		}
	}
	// from here v should not be ptr
	if isSupportedKind(v.Kind()) && v.Kind() != reflect.Ptr && !v.CanSet() {
		v = reflect.NewAt(v.Type(), unsafe.Pointer(v.UnsafeAddr())).Elem()
	}

	switch v.Kind() {
	// bool
	case reflect.Bool:
		v.SetBool(rand.Intn(2) > 0)
	// unsigned integer
	case reflect.Uint8:
		v.SetUint(uint64(rand.Int() % 255))
	case reflect.Uint16:
		v.SetUint(uint64(rand.Int() % 65535))
	case reflect.Uint32:
		v.SetUint(uint64(rand.Uint32()))
	case reflect.Uint64:
		v.SetUint(rand.Uint64())
	// signed integer
	case reflect.Int8:
		v.SetInt(int64(rand.Int()%255 - 127))
	case reflect.Int16:
		v.SetInt(int64(rand.Int()%65535 - 32768))
	case reflect.Int32:
		v.SetInt(int64(rand.Int31()))
	case reflect.Int:
		v.SetInt(int64(rand.Int()))
	case reflect.Int64:
		v.SetInt(rand.Int63())
	// float
	case reflect.Float32:
		v.SetFloat(float64(rand.Float32()))
	case reflect.Float64:
		v.SetFloat(rand.Float64())
	// string
	case reflect.String:
		n := rand.Int()%DefaultStringLength + 1
		if tag != "" {
			if e, err := strconv.Atoi(tag); err == nil {
				n = e
			}
		}
		v.SetString(randomString(n))
	// slice
	case reflect.Slice:
		n := rand.Int()%DefaultSliceSize + 1
		if tag != "" {
			if e, err := strconv.Atoi(tag); err == nil {
				n = e
			}
		}
		tp := v.Type()
		sl := reflect.New(tp).Elem()
		etp := tp.Elem()
		var newv reflect.Value
		for i := 0; i < n; i++ {
			// handle element ptr case
			if etp.Kind() == reflect.Ptr {
				newv = reflect.New(etp.Elem())
			} else {
				newv = reflect.New(etp).Elem()
			}
			Fuzz(newv, "")
			sl = reflect.Append(sl, newv)
		}
		v.Set(sl)
	// map
	case reflect.Map:
		n := rand.Int()%DefaultMapSize + 1
		if tag != "" {
			if e, err := strconv.Atoi(tag); err == nil {
				n = e
			}
		}
		tp := v.Type()
		m := reflect.MakeMap(tp)
		for i := 0; i < n; i++ {
			newk, newv := reflect.New(tp.Key()), reflect.New(tp.Elem())
			Fuzz(newk, "")
			Fuzz(newv, "")
			m.SetMapIndex(newk.Elem(), newv.Elem())
		}
		v.Set(m)
	// struct
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			tag := v.Type().Field(i).Tag
			log.Printf("Fuzz %s\n", v.Type().Field(i).Name)
			Fuzz(v.Field(i), tag.Get("fuzz"))
		}
	default:
		log.Printf("Do not support type %v, %#v\n", v.Kind(), v)
	}
}
