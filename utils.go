package xutils

import (
	"bytes"
	"encoding/gob"

	"golang.org/x/exp/constraints"
)

func In[T constraints.Ordered, S []T](t T, s S) bool {
	for _, d := range s {
		if d == t {
			return true
		}
	}
	return false
}

// 深拷贝
func DeepCopy[T any](src, dst T) error {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	dec := gob.NewDecoder(&buf)

	err := enc.Encode(src)
	if err != nil {
		return err
	}
	err = dec.Decode(&dst)
	return err
}

func Deduplication[T constraints.Ordered, S []T](s S) S {
	unique := make(map[T]bool)
	result := make(S, 0, len(s))
	for _, t := range s {
		if !unique[t] {
			unique[t] = true
			result = append(result, t)
		}
	}
	return result
}

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Abs[T constraints.Integer | constraints.Float](a T) T {
	if a < 0 {
		return -a
	}
	return a
}

func Must(fs ...func() error) error {
	for _, f := range fs {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}
