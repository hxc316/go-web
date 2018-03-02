package main

type inter interface {
	do(s string)
}


type AA struct {
	name string
}

func (aa *AA) setName()  {
	aa.name = "123"
}

func (aa *AA) printName()  {
	println("AA name = ", aa.name)
}

func (aa *AA) do(a string)  {
	println("do : ", a)
}

func main() {

	aa := AA{}
	//aa.setName()
	//aa.printName()
	var i inter
	i = &aa
	i.do("xc.hu")
}
