package common

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func Abs[T constraints.Signed](x T) T {
	if x < 0 {
		return x * -1
	}
	return x
}

func Max[T constraints.Ordered](args ...T) T {
	switch len(args) {
	case 0:
		panic("no arguments")
	case 1:
		return args[0]
	case 2:
		if args[0] > args[1] {
			return args[0]
		}
		return args[1]
	default:
		m := args[0]
		for _, v := range args[1:] {
			if v > m {
				m = v
			}
		}
		return m
	}
}

func Min[T constraints.Ordered](args ...T) T {
	switch len(args) {
	case 0:
		panic("no arguments")
	case 1:
		return args[0]
	case 2:
		if args[0] < args[1] {
			return args[0]
		}
		return args[1]
	default:
		m := args[0]
		for _, v := range args[1:] {
			if v > m {
				m = v
			}
		}
		return m
	}
}

func Sum[T Number](nums ...T) T {
	var acc T
	for _, v := range nums {
		acc += v
	}
	return acc
}

func Product[T Number](nums ...T) T {
	var acc T
	for _, v := range nums {
		acc *= v
	}
	return acc
}
