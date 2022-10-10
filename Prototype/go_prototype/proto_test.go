package go_prototype

import (
	"fmt"
	"testing"
)

func TestPrototype(t *testing.T) {
	file1 := &File{"a.txt"}
	file2 := &File{"b.txt"}
	folder1 := &Folder{name: "tmp", children: []Inode{file1, file2}}
	folder2 := &Folder{name: "root", children: []Inode{folder1}}
	cloneFolder2 := folder2.Clone()
	fmt.Println(cloneFolder2.Print(" "))
}
