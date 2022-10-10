package go_prototype

type Inode interface {
	Clone() Inode
	Print(string) string
}
