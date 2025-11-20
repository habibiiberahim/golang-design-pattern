package main

import (
	"golang-design-pattern/behavioral/observer"
	"golang-design-pattern/creational/abstract_factory"
	"golang-design-pattern/creational/factory_method"
	"golang-design-pattern/creational/singleton"
	"golang-design-pattern/structural/adapter"
)

func main() {
	adapter.RunAdapter()

	factory_method.RunFactoryMethod()
	abstract_factory.RunAbstractFactory()
	singleton.RunSingleton()

	observer.RunObserver()

}
