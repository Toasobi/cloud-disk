package test

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"testing"
)

// //文件分片

// //首先指定文件分片大小
const chunkSize = 1024 * 1024 //1 m

//文件分片
func TestGenerate(t *testing.T) {
	fileInfo, err := os.Stat("./img/testingJPG.jpg")
	if err != nil {
		t.Fatal(err)
	}
	//分片个数
	chunkNum := math.Ceil(float64((float64(fileInfo.Size()) / float64(chunkSize))))
	fmt.Println(chunkNum)
	myFile, err := os.OpenFile("./img/testingJPG.jpg", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	b := make([]byte, chunkSize)
	for i := 0; i < int(chunkNum); i++ {
		//指定读取文件的起始点
		myFile.Seek(int64(i*chunkSize), 0)
		if chunkSize > fileInfo.Size()-int64(i*chunkSize) {
			b = make([]byte, fileInfo.Size()-int64(i*chunkSize))
		}
		myFile.Read(b)

		//在该目录下创建写入的文件

		f, err := os.OpenFile("./"+strconv.Itoa(i)+".chunk", os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}

		f.Write(b)
		f.Close()

	}
	myFile.Close()
}

//分片文件的合并
func TestMergeChunkFile(t *testing.T) {
	//首先指定一个合并的文件
	myFile, err := os.OpenFile("test02.mp4", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}

	//这里原本是前端会传，这里我们是自己测试就直接拿过来了
	fileInfo, err := os.Stat("test.mp4")
	if err != nil {
		t.Fatal(err)
	}
	//分片个数
	chunkNum := math.Ceil(float64((float64(fileInfo.Size()) / float64(chunkSize))))
	fmt.Println(chunkNum)
	for i := 0; i < int(chunkNum); i++ {
		f, err := os.OpenFile("./"+strconv.Itoa(i)+".chunk", os.O_RDONLY, os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}
		b, err := ioutil.ReadAll(f)
		if err != nil {
			t.Fatal(err)
		}
		myFile.Write(b)
		f.Close()
	}
	myFile.Close()
}

//文件一致性校验
func TestCheck(t *testing.T) {
	//获取第一个文件的信息
	file1, err := os.OpenFile("test.mp4", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	b1, err := ioutil.ReadAll(file1)
	if err != nil {
		t.Fatal(err)
	}
	//获取第二个文件的信息
	file2, err := os.OpenFile("test02.mp4", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	b2, err := ioutil.ReadAll(file2)
	if err != nil {
		t.Fatal(err)
	}

	//判断是否一致
	s1 := fmt.Sprintf("%x", md5.Sum(b1))
	s2 := fmt.Sprintf("%x", md5.Sum(b2))

	if s1 == s2 {
		fmt.Println("一致")
	} else {
		fmt.Println("不一致")
	}
}
