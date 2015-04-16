package main

type Encoder struct {
	charset string
}

func NewEncoder(charset string) *Encoder {
	e := new(Encoder)
	e.charset = charset
	return e
}

var DefaultEncoder = NewEncoder("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var UnambiguousEncoder = NewEncoder("23456789abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ")
var ShortifyEncoder = DefaultEncoder

func (self *Encoder) Encode(value int64) string {
	if value == 0 {
		return string(self.charset[0])
	}

	chars := make([]byte, 0)
	base := int64(len(self.charset))

	for value > 0 {
		index := value % base
		value = value / base
		chars = append([]byte{self.charset[index]}, chars...)
	}

	return string(chars)
}
