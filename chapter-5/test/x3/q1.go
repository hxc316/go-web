package x3

type P1 interface {
	Oo(s string)
}

type Qq1 struct {

}

func (q *Qq1) Oo(s string)  {
	print("Qq1 Oo(s string) s = ",s)
}

func (q *Qq1) Mm(s string)  {
	print("Qq1 mm(s string) s = ",s)
}