package main

import (
	"golang-design-pattern/behavioral/observer"
	"golang-design-pattern/creational/abstract_factory"
	"golang-design-pattern/creational/factory_method"
)

func main() {
	factory_method.RunFactoryMethod()
	abstract_factory.RunAbstractFactory()
	observer.RunObserver()
}
