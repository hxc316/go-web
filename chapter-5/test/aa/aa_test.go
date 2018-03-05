package test

import (
	"testing"
	"time"
	"sync"
)

func TestXaa(t *testing.T)  {
	println(111111)
}
//----------------------------------------------------------
func mm(s ...string)  {

	for _,q := range s{
		println(q)
	}
}

func TestX1(t *testing.T)  {
	mm("a","b","c","d")
	time.Sleep(time.Duration(300))
}
//-------------------------------------------------------------
func TestX2(t *testing.T)  {
	xx("11111111")
}

func xx(s interface{} )  {
	println("sssss")
}

// ------------------------------------------------------------

const name  =  1

const (
	a1	=	10
	a2	=	20
)

func TestAa(t *testing.T)  {
	println("const name = ",name)
	println("const a2 = ", a1)
}

// -------------------------------------------------------------------

// wg.Add(1)
func TestGo(t *testing.T)  {
	var wg sync.WaitGroup
	//var sy sync.Mutex
	for i:=0;i<100 ;i++  {
		wg.Add(1)
		go func(i int) {
			println("time = ",i)
			defer wg.Done()
		}(i)

	}
	wg.Wait()
	println("-------end--------")
}

// -----------------------------------


func TestX5(t *testing.T)  {
	m := make(chan int, 1000)
	go func() {
			for i:=0;i<10000 ;i++  {
			m <- i
		}
	}()
	Done:
	for ;;{
		select{
			case  a := <-m:
				println("m = ",a)
			default:
				break Done
		}
	}
}




