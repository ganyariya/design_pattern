package go_prototype

type File struct {
	name string
}

func (f *File) Print(ident string) string {
	return ident + f.name
}

func (f *File) Clone() Inode {
	return &File{name: f.name + "_clone"}
}
