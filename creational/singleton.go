package creational

import (
	"sync"
)

// each class/object only need declared once
type SingletonObj struct {
	Name string
}

var (
	singletonOnce sync.Once
	Singleton     SingletonObj
)

func GetSingleton() SingletonObj {
	singletonOnce.Do(func() {
		Singleton = SingletonObj{}
	})

	return Singleton
}
