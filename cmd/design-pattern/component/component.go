package component

import (
	"fmt"
	"strconv"
)

type File interface {
	killVirus()
}

type ImageFile struct {
	Name string
}

type TextFile struct {
	Name string
}

type VideoFile struct {
	Name string
}

type Folder struct {
	container []File
	Name      string
}

func (image *ImageFile) killVirus() {
	fmt.Println("对图片文件[" + image.Name + "]进行查杀")
}

func (text *TextFile) killVirus() {
	fmt.Println("对文本文件[" + text.Name + "]进行查杀")
}

func (video *VideoFile) killVirus() {
	fmt.Println("对视频文件[" + video.Name + "]进行查杀")
}

func (folder *Folder) killVirus() {
	fmt.Println("对文件夹[" + folder.Name + "]进行查杀")
	for _, file := range folder.container {
		file.killVirus()
	}
}

func (folder *Folder) Add(file ...File) *Folder {
	folder.container = append(folder.container, file...)
	return folder
}

func (folder *Folder) Remove(index int) {
	if index > len(folder.container) {
		fmt.Println("删除index[" + strconv.Itoa(index) + "]无效,跳过删除")
		return
	}
	folder.container = append(folder.container[:index], folder.container[index+1:]...)
}

func (folder *Folder) Get(index int) File {
	return folder.container[index]
}
