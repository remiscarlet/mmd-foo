package main

import (
	"encoding/binary"
	"fmt"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"math"
	"strings"
)

func readBytes(contents *[]byte, offset int, size int) ([]byte, error) {
	content := (*contents)[offset : offset+size]
	return content, nil
}

func float32FromBytes(bytes []byte) float32 {
	bits := binary.BigEndian.Uint32(bytes)
	float := math.Float32frombits(bits)
	return float
}

func float32Bytes(float float32) []byte {
	bits := math.Float32bits(float)
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint32(bytes, bits)
	return bytes
}

type BoneFrame struct {
	BoneName     string
	FrameId      int
	X            float64
	Y            float64
	Z            float64
	X_rq         float64
	Y_rq         float64
	Z_rq         float64
	W_rq         float64
	FrameInterpo []byte
}

type Type int

const (
	STRING Type = iota
	INT
	FLOAT
	BYTES
	BONEFRAMES_SUBSECT
)

type FileStructureSection struct {
	Name       string
	ByteLen    int
	Type       Type
	RawContent []byte
	SubSect    []FileStructureSection
}

func (sect *FileStructureSection) DecodeContent() interface{} {
	switch sect.Type {
	case STRING:
		s, _, err := transform.Bytes(japanese.ShiftJIS.NewDecoder(), sect.RawContent)
		if err != nil {
			panic(err)
		}
		return s
	case INT:
		x := binary.BigEndian.Uint32(sect.RawContent)
		return x
	case FLOAT:
		fmt.Println("o shit")
		f := float32FromBytes(sect.RawContent)
		fmt.Println(f)
		return f
	}
	return nil
}
func (sect *FileStructureSection) PrintContent() {
	data := sect.DecodeContent()
	fmt.Printf("--------\n")
	fmt.Printf("%v\n", data)
	fmt.Printf("%s\n", data)
	fmt.Printf("%d\n", data)
}

func (sect *FileStructureSection) SetRawContent(b_arr []byte) {
	sect.RawContent = b_arr
}

func (sect *FileStructureSection) IsSimpleSection() bool {
	for _, t := range []Type{STRING, INT, FLOAT, BYTES} {
		if t == sect.Type {
			return true
		}
	}

	return false
}

func getBoneFrameStructure() []*FileStructureSection {
	structure := []*FileStructureSection{
		&FileStructureSection{"bone_name", 15, STRING, nil, nil},
		&FileStructureSection{"frame_id", 4, INT, nil, nil},
		&FileStructureSection{"x", 4, FLOAT, nil, nil},
		&FileStructureSection{"y", 4, FLOAT, nil, nil},
		&FileStructureSection{"z", 4, FLOAT, nil, nil},
		&FileStructureSection{"x_rq", 4, FLOAT, nil, nil},
		&FileStructureSection{"y_rq", 4, FLOAT, nil, nil},
		&FileStructureSection{"z_rq", 4, FLOAT, nil, nil},
		&FileStructureSection{"w_rq", 4, FLOAT, nil, nil},
		&FileStructureSection{"frame_interpo", 64, BYTES, nil, nil},
	}
	return structure
}

func getFileStructure() []*FileStructureSection {
	structure := []*FileStructureSection{
		&FileStructureSection{"version", 30, STRING, nil, nil},
		&FileStructureSection{"model_name", 20, STRING, nil, nil},
		&FileStructureSection{"bone_frame:count", 4, INT, nil, nil},
		&FileStructureSection{"bone_frame:section", 1, BONEFRAMES_SUBSECT, nil, nil},
	}
	return structure
}

func loadStructures(structure []*FileStructureSection, content *[]byte, OFFSET int) (map[string]*FileStructureSection, int, error) {
	populated_structures := make(map[string]*FileStructureSection)
	for i, section := range structure {
		if section.IsSimpleSection() {
			// Its own section
			val, err := readBytes(content, OFFSET, section.ByteLen)
			if err != nil {
				panic(err)
			}
			section.SetRawContent(val)
			section.PrintContent()

			fmt.Printf("%d) %v\n", i, section)
			populated_structures[section.Name] = section
			OFFSET += section.ByteLen
		} else {
			// Has subsections
			fmt.Println("WIP")
			fmt.Printf("%v\n", populated_structures)

			section_name_base := strings.Split(section.Name, ":")[0]
			section_count := populated_structures[section_name_base+":count"].DecodeContent()
			fmt.Printf("%v\n", section_count)
			switch section.Type {
			case BONEFRAMES_SUBSECT:
				boneframe_structure := getBoneFrameStructure()
				populated_subsect, OFFSET, err := loadStructures(boneframe_structure, content, OFFSET)
				if err != nil {
					panic(err)
				}
				fmt.Printf("SUBSECT: %v\n", populated_subsect)
				fmt.Printf("OFFSET: %d\n", OFFSET)
				break
			}
		}
	}

	return populated_structures, OFFSET, nil
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

	populated_structures, OFFSET, err := loadStructures(structure, &content, OFFSET)
	fmt.Printf("%v\n", populated_structures)

	return
}
