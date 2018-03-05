package main

import (
	"strconv"
	"time"
	//"github.com/wandoulabs/codis/pkg/utils/errors"
)

func main() {

	//testChain()
	//testStrconv()
	//testByte()
	//testSelect()
	testAppend()
	//testIf()

}

func testChain()  {
	a := make(chan int,3)

	a <- 1
	a <- 2
	a <- 3

	m := <- a

	println(m)
}

func  testStrconv (){
	a := strconv.FormatInt(2,10)

	println("strconv.FormatInt(2,10) = ",a)

	is ,error:= strconv.ParseBool("t")
	if error == nil{
		println("strconv.ParseBool(\"t\") = ", is )
	}
}

func testByte()  {
	var j = []byte(`1234`)

	println("j = ",string(j))
}

func testSelect()  {

	go func() {
		for i:=0;i<10;i++{
			println("i = ",i)
		}

	}()
	time.Sleep(time.Millisecond * 200)
}

func testIf()  {
	var(
		a ,b bool
	)
	a ,b = true,true

	if a && b{
		println("true")
	}
}

func testError()  {
	//a := errors.New("error test")

}

func testAppend()  {
	var dd []string
	dd = append(dd,"aaa")
	dd = append(dd,"bbb")
	dd = append(dd,"ccc")

	for _ , s := range dd{
		println(s)
	}
}
