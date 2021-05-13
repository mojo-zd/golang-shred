package component

import "testing"

func TestComponent(t *testing.T) {
	f1 := &Folder{Name: "我的收藏"}
	f1.Add(
		&ImageFile{
			Name: "slice.jpg",
		},
		&TextFile{
			Name: "hello.txt",
		})

	f4 := &Folder{Name: "工作空间"}
	f4.Add(
		&VideoFile{Name: "财富自由之路.avi"},
		&VideoFile{Name: "设计模式.avi"},
	)
	f1.Add(f4)
	f1.killVirus()
	f1.Remove(1)
}
