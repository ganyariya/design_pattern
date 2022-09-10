package go_template_method

type IToBinary interface {
	GetFileName() string
	GetContent(string) string
}

type ToBinary struct {
	iToBinary IToBinary
}

func (b *ToBinary) convertBinary(content string) []byte {
	return []byte(content)
}

func (b *ToBinary) Convert() []byte {
	name := b.iToBinary.GetFileName()
	content := b.iToBinary.GetContent(name)
	return b.convertBinary(content)
}
