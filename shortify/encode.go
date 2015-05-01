package shortify

var shortifyEncoder encoder

type encoder struct {
	charset string
}

func (self *encoder) encode(value int64) string {
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
