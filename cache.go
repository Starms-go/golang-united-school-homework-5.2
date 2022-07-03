package cache

import "time"

type Cache struct {
	m map[string]Val
}

type Val struct {
	v   string
	exp time.Time
}

func NewCache() Cache {
	return Cache{m: map[string]Val{}}
}

func (receiver Cache) Get(key string) (string, bool) {
	for k, v := range receiver.m {
		if !(v.exp.IsZero()) && v.exp.Before(time.Now()) {
			delete(receiver.m, k)
		}
	}
	if v, ok := receiver.m[key]; ok {
		return v.v, ok
	}
	return "", false
}

func (receiver Cache) Put(key, value string) {
	for k, v := range receiver.m {
		if !(v.exp.IsZero()) && v.exp.Before(time.Now()) {
			delete(receiver.m, k)
		}
	}
	receiver.m[key] = Val{v: value, exp: time.Time{}}
}

func (receiver Cache) Keys() []string {
	for k, v := range receiver.m {
		if !(v.exp.IsZero()) && v.exp.Before(time.Now()) {
			delete(receiver.m, k)
		}
	}
	var keys []string
	for k := range receiver.m {
		keys = append(keys, k)
	}
	return keys
}

func (receiver Cache) PutTill(key, value string, deadline time.Time) {
	for k, v := range receiver.m {
		if !(v.exp.IsZero()) && v.exp.Before(time.Now()) {
			delete(receiver.m, k)
		}
	}
	receiver.m[key] = Val{v: value, exp: deadline}
}
