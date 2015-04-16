package shortify

type Encoder struct {
	charset string
}

func NewEncoder(charset string) *Encoder {
	e := new(Encoder)
	e.charset = charset
	return e
}

var encoders = make(map[string]*Encoder)
var ShortifyEncoder *Encoder

func init() {
	encoders["default"] = NewEncoder("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	encoders["unambiguous"] = NewEncoder("23456789abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ")
	ShortifyEncoder = encoders["default"]
}

func SetDefaultEncoder(name string) {
	if encoder, ok := encoders[name]; ok {
		ShortifyEncoder = encoder
	}
}

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
