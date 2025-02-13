package models

type Product string

// only for test
var (
	POWERBANK   Product = "powerbank"
	BOOK        Product = "book"
	PEN         Product = "pen"
	PINKHOODY   Product = "pink-hoody"
	FAKEPRODUCT Product = "fake"
)

var priceTable map[Product]int64 = map[Product]int64{
	POWERBANK:   200,
	BOOK:        50,
	PEN:         10,
	PINKHOODY:   500,
	FAKEPRODUCT: -1,
}

func (p *Product) Price() int64 {
	return priceTable[*p]
}
