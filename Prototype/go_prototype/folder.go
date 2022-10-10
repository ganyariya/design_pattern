package go_prototype

type Folder struct {
	children []Inode
	name     string
}

func (f *Folder) Print(ident string) string {
	ret := ""
	ret += ident + f.name + "\n"
	for _, i := range f.children {
		ret += i.Print(ident+ident) + "\n"
	}
	return ret
}

func (f *Folder) Clone() Inode {
	cloneFolder := &Folder{name: f.name + "_clone"}
	cloneChildren := make([]Inode, 0)
	for _, i := range f.children {
		cloneChildren = append(cloneChildren, i.Clone())
	}
	cloneFolder.children = cloneChildren
	return cloneFolder
}
