package filesystem

type Visitor interface {
	VisitFile(node *File)
	VisitDir(node *Dir)
}
type SizeVisitor struct {
	count int
	total int
}

func (s *SizeVisitor) VisitFile(node *File) {
	s.count++
	s.total += node.size
}
func (s *SizeVisitor) VisitDir(node *Dir) {
	for _, f := range node.files {
		f.Accept(s)
	}
}
