package pkg

import (
	"k8s.io/client-go/dynamic"
)

func Foo() {
	fmt.Printf("%+v\n", dynamic.NewForConfigOrDie(nil))
}