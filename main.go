package main

import (
	"fmt"
	"encoding/binary"
	"io/ioutil"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func readBytes(contents []byte, offset int, size int) ([]byte, error) {
	//rtn_bytes := make([]byte, size)
	content := contents[offset:offset+size]
	//for i, b := range content {
		//fmt.Printf("%d) %s - %X\n", i, string(b), b)
	//}
	//fmt.Printf("+_+_+_+_+_+_+_\n")
	//fmt.Printf("%X\n", content)
	//s, _, err := transform.Bytes(japanese.ShiftJIS.NewDecoder(), content)
	//fmt.Printf("%X\n", s)
	//return s, err

	return content, nil
}

func printBytes(bytes []byte) {
	fmt.Printf("________\n")
	fmt.Printf("%X\n", bytes)
	fmt.Printf("%s\n", bytes)
	fmt.Printf("%d\n", bytes)

	//x := binary.BigEndian.Uint32(bytes)
	//fmt.Println(x)
}


type BoneFrame struct {
	BoneName string
	FrameId int
	X float64
	Y float64
	Z float64
	X_rq float64
	Y_rq float64
	Z_rq float64
	W_rq float64
	FrameInterpo []byte
}

type Type int
const (
	STRING Type = iota
	INT
	FLOAT
	BYTES
)

type FileStructureSection struct {
	Name string
	ByteLen int
	Type Type
	RawContent []byte
	DecodedContent []byte
}
func (sect *FileStructureSection) DecodeContent(){
	switch sect.Type {
	case STRING:
		s, _, err := transform.Bytes(japanese.ShiftJIS.NewDecoder(), sect.RawContent)
		if err != nil {
			panic(err)
		}
		sect.DecodedContent = s
	case INT:
		x := binary.BigEndian.Uint32(sect.RawContent)
		decoded := make([]byte, 4)
		binary.BigEndian.PutUint32(decoded, x)
		sect.DecodedContent = decoded
	case FLOAT:
		fmt.Println("o shit")
	}
}

func getFileStructure() []FileStructureSection {
	structure := []FileStructureSection{
		FileStructureSection{"version", 30, STRING, nil, nil},
		FileStructureSection{"model_name", 20, STRING, nil, nil},
		FileStructureSection{"bone_frame_count", 4, INT, nil, nil},
	}
	return structure
}

func main() {
    filename := "motion.vmd"
	//encoding := "cp932"

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	OFFSET := 0
	structure := getFileStructure()
	for i, section := range structure {
		val, err := readBytes(content, OFFSET, section.ByteLen)
		if err != nil {
			panic(err)
		}
		section.RawContent = val
		section.DecodeContent()
		printBytes(section.DecodedContent)

		fmt.Printf("%d) %v\n", i, section)
		OFFSET += section.ByteLen
	}


}

