package filesystem

import (
	"testing"
)

func TestVisitor(t *testing.T) {
	root := &Dir{name: "root"}
	downloads := &Dir{name: "downloads"}
	documents := &Dir{name: "documents"}
	root.Add(downloads)
	root.Add(documents)
	downloads.Add(&File{name: "GoLand.dmg", size: 1024000})
	downloads.Add(&File{name: "IntelliJ IDEA.dmg", size: 2048000})

	documents.Add(&File{name: "readme.txt", size: 2048})
	documents.Add(&File{name: "photo.jpg", size: 1024})

	sizeVisitor := &SizeVisitor{}
	root.Accept(sizeVisitor)
	if sizeVisitor.count != 4 && sizeVisitor.total != 3075072 {
		t.Errorf("Wrong size %d and wrong count: %d", sizeVisitor.total, sizeVisitor.count)
	}
}
