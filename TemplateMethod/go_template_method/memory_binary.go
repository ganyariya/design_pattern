package go_template_method

var memory = map[string]string{
	"tmp":  "HELLO",
	"tmp2": "HELLO2",
}

type MemoryBinary struct {
	fileName string
}

func NewMemoryBinary(fileName string) *MemoryBinary {
	return &MemoryBinary{fileName: fileName}
}

func (mb *MemoryBinary) GetFileName() string {
	return mb.fileName
}

func (mb *MemoryBinary) GetContent(filename string) string {
	// In Memory なのでとりあえず返す
	return memory[filename]
}
