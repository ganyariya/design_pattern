package visitor

type File struct {
	Name string
	Size int
}

func NewFile(name string, size int) *File {
	return &File{Name: name, Size: size}
}

func (f *File) GetName() string {
	return f.Name
}
func (f *File) GetSize() int {
	return f.Size
}
func (f *File) Accept(v VisitorInterface) any {
	return v.Visit(f)
}
