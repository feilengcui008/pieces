package main

import (
	"expvar"
	"net/http"
	"strconv"
	"time"
)

func main() {
	expvar.NewInt("atomic_int_counter")
	expvar.NewFloat("atomic_float_counter")
	expvar.NewString("rwlock_string")
	expvar.NewMap("rwlock_map")
	for i := 0; i < 10; i++ {
		go func() {
			for {
				n := expvar.Get("atomic_int_counter")
				v, _ := n.(*expvar.Int)
				v.Add(1)
				time.Sleep(time.Second * 3)
			}
		}()
	}
	go func() {
		i := 0
		for {
			m := expvar.Get("rwlock_map")
			v, _ := m.(*expvar.Map)
			expv := &expvar.String{}
			expv.Set(strconv.Itoa(i))
			v.Set("key", expv)
			i++
			time.Sleep(time.Second * 3)
		}
	}()
	http.ListenAndServe(":8888", nil)
}
