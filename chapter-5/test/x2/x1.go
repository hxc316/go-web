package main

import "github.com/hxc316/go-web/chapter-5/test/x3"

type Aa interface {

	Mm(s string)
}

type Aa1 struct {

}

func (a *Aa1) Mm(s string)  {
	println("Aa1 mm(s string) s = " ,s)
}


type Bb1 struct {

}

func (b *Bb1) Mm(s string)  {
	println("Bb1 mm(s string) s = ",s)
}

func main() {

	var  a Aa
	a = &Aa1{}
	a.Mm("123")

	a = &Bb1{}
	a.Mm("789")

	a = &x3.Qq1{}
	a.Mm("123456789")
}








