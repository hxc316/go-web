package main

import (
	"strconv"
	"time"
)

func main() {

	//testChain()
	//testStrconv()
	//testByte()
	testSelect()

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
