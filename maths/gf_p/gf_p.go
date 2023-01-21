package gf_p

import "github.com/rahul0tripathi/go-shots/maths/euclidean"

type GfP struct {
	Field int
}
type GfPItem struct {
	parent *GfP
	Value  int
}

func NewField(a int) *GfP {
	return &GfP{Field: a}
}
func (gf *GfP) Item(a int) *GfPItem {
	return &GfPItem{
		parent: gf,
		Value:  a,
	}
}

func Add(a *GfPItem, b *GfPItem) *GfPItem {
	return &GfPItem{
		parent: a.parent,
		Value:  (a.Value + b.Value) % a.parent.Field,
	}
}

func Mul(a *GfPItem, b *GfPItem) *GfPItem {
	return &GfPItem{
		parent: a.parent,
		Value:  (a.Value * b.Value) % a.parent.Field,
	}
}

func Sub(a *GfPItem, b *GfPItem) *GfPItem {
	return &GfPItem{
		parent: a.parent,
		Value:  (a.Value + (a.parent.Field - b.Value)) % a.parent.Field,
	}
}

func Div(a *GfPItem, b *GfPItem) *GfPItem {
	_, bInv, _ := euclidean.ExtendedEuclidean(a.parent.Field, b.Value, 1, 0, 0, 1)
	if bInv < 0 {
		for bInv < 0 {
			bInv += a.parent.Field
		}
	}
	return &GfPItem{
		parent: a.parent,
		Value:  (a.Value * bInv) % a.parent.Field,
	}
}
