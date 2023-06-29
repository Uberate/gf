package main

import (
	"fmt"
	"gf"
	"github.com/sirupsen/logrus"
)

type AnimalInterface interface {
	gf.Entity

	Speak() string
}

type AbsAnimal struct {
	gf.BaseEntity

	speakFunc func() string
}

func (aa *AbsAnimal) Speak() string {
	return aa.speakFunc()
}

func DefaultSpeak() string {
	return ""
}

func NewDefaultAnimal() AnimalInterface {
	return &AbsAnimal{
		BaseEntity: gf.NewBaseEntityGenerator("", "")(""),
		speakFunc:  DefaultSpeak,
	}
}

func AbsAnimalInterface(kind, version string, speakBehavior func() string) gf.Generator[AnimalInterface] {
	if speakBehavior == nil {
		speakBehavior = DefaultSpeak
	}

	baseEntity := gf.NewBaseEntityGenerator(kind, version)

	return func(name string, config interface{}, logger *logrus.Logger) (AnimalInterface, error) {
		return &AbsAnimal{
			BaseEntity: baseEntity(name),
			speakFunc:  speakBehavior,
		}, nil
	}
}

var AnimalFactory *gf.Factory[AnimalInterface]

func init() {
	AnimalFactory = gf.NewFactor(NewDefaultAnimal())

	AnimalFactory.Registry("cat", "", AbsAnimalInterface("cat", "", func() string { return "I'm a cat" }))
	AnimalFactory.Registry("dog", "", AbsAnimalInterface("dog", "", func() string { return "I'm a dog" }))
}

func main() {
	cat, err, ok := AnimalFactory.Get("cat", "", "cat1", nil, nil)
	if err != nil {
		panic(err)
	}
	if ok {
		fmt.Println(cat.Speak())
	} else {
		panic("not found cat")
	}
}
