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

func truncateFromNullTerm(bytes *[]byte) *[]byte {
	rtn_size := len(*bytes)
	for i, b := range *bytes {
		if b == byte('\000') {
			rtn_size = i
			break
		}
	}
	rtn_bytes := make([]byte, rtn_size)
	copied_bytes_n := copy(rtn_bytes, *bytes)
	if copied_bytes_n == 0 {
		fmt.Println("Copied zero bytes - is that right?")
	}
	return &rtn_bytes
}

func readBytes(contents *[]byte, offset int, size int) (*[]byte, int, error) {
	content := (*contents)[offset : offset+size]
	return &content, offset + size, nil
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

type SectType int

const (
	STRING SectType = iota
	INT
	FLOAT
	BYTES
	BONEFRAMES_SUBSECT
)

type FileStructureSection struct {
	Name       string
	ByteLen    int
	SectType   SectType
	RawContent *[]byte
	SubSect    []*FileStructureSection
}

func (sect *FileStructureSection) DecodeContent() interface{} {
	switch sect.SectType {
	case STRING:
		bytes := truncateFromNullTerm(sect.RawContent)
		s, _, err := transform.Bytes(japanese.ShiftJIS.NewDecoder(), *bytes)
		if err != nil {
			panic(err)
		}
		return s
	case INT:
		x := binary.BigEndian.Uint32(*sect.RawContent)
		return x
	case FLOAT:
		fmt.Println("o shit")
		f := float32FromBytes(*sect.RawContent)
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

func (sect *FileStructureSection) SetRawContent(b_arr *[]byte) {
	sect.RawContent = b_arr
}

func (sect *FileStructureSection) IsSimpleSection() bool {
	for _, t := range []SectType{STRING, INT, FLOAT, BYTES} {
		if t == sect.SectType {
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
	var bytes_val *[]byte
	var err error

	populated_structures := make(map[string]*FileStructureSection)
	for i, section := range structure {
		if section.IsSimpleSection() {
			// Its own section
			bytes_val, OFFSET, err = readBytes(content, OFFSET, section.ByteLen)
			if err != nil {
				panic(err)
			}
			section.SetRawContent(bytes_val)
			//section.PrintContent()

			fmt.Printf("%d) %v - bytes: %v\n", i, section, *section.RawContent)
			populated_structures[section.Name] = section
		} else {
			// Has subsections
			fmt.Println("WIP")
			fmt.Printf("%v\n", populated_structures)

			section_name_base := strings.Split(section.Name, ":")[0]
			section_count := populated_structures[section_name_base+":count"].DecodeContent()
			fmt.Printf("%v\n", section_count)
			//for frame_idx := 0; frame_idx <= section_count; frame_idx++ {
			for frame_idx := 0; frame_idx <= 5; frame_idx++ {
				fmt.Println("++++++++++++++++++++++++++++++++++++++++++++")
				switch section.SectType {
				case BONEFRAMES_SUBSECT:
					var populated_subsect map[string]*FileStructureSection

					boneframe_structure := getBoneFrameStructure()
					fmt.Printf("Loading boneframe structures with OFFSET %d\n", OFFSET)
					populated_subsect, OFFSET, err = loadStructures(boneframe_structure, content, OFFSET)
					if err != nil {
						panic(err)
					}
					fmt.Printf("SUBSECT: %v\n", populated_subsect)
					fmt.Printf("OFFSET: %d\n", OFFSET)
				}
			}
		}
	}

	return populated_structures, OFFSET, nil
}

func main() {
	// Hey dipshit leave some comments next time.
	filename := "kimiiro_ni_somaru_camera.vmd"
	//encoding := "cp932"

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	OFFSET := 0
	structure := getFileStructure()

	populated_structures, OFFSET, err := loadStructures(structure, &content, OFFSET)
	fmt.Printf("-----%v\n", populated_structures)

	return
}
