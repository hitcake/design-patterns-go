package filesystem

type FileSystemNode interface {
	Add(f FileSystemNode)
	Accept(v Visitor)
}

type File struct {
	name string
	size int
}

func (f *File) Add(v FileSystemNode) {

}
func (f *File) Accept(v Visitor) {
	v.VisitFile(f)
}

type Dir struct {
	name  string
	files []FileSystemNode
}

func (d *Dir) Accept(v Visitor) {
	v.VisitDir(d)
}
func (d *Dir) Add(v FileSystemNode) {
	d.files = append(d.files, v)
}
