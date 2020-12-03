package Library

import (
	"encoding/json"
)

type ReadCoilStruct struct {
	Code_name     string
	Start_Address uint16
	Last_Address  uint16
}

type WriteCoilStruct struct {
	Code_name     string
	Start_Address uint16
	Last_Address  uint16
}

type ReadRegStruct struct {
	Code_name     string
	Start_Address uint16
	Last_Address  uint16
}

type WriteRegStruct struct {
	Code_name     string
	Start_Address uint16
	Last_Address  uint16
}

type ErrJson struct {
	Errormessage string
}

type ReadRegInStruct struct {
	Alias   string
	Address uint16
	Value   uint16
}

var Module_0 = []string{"Un_name0", "Un_name1", "Un_name2", "Un_name3", "Un_name4", "Un_name5", "Un_name6",
	"Un_name7", "Un_name8", "Un_name9", "Un_name10", "Un_name11", "Un_name12", "Un_name13", "Un_name14",
	"Un_name15", "Un_name16", "Un_name17", "Un_name18", "Un_name19", "Un_name20"}

var Module_1 = []string{"Un_name0", "Un_name1", "Un_name2", "Un_name3", "Un_name4", "Un_name5", "Un_name6",
	"Un_name7", "Un_name8", "Un_name9", "Un_name10", "Un_name11", "Un_name12", "Un_name13", "Un_name14",
	"Un_name15", "Un_name16", "Un_name17", "Un_name18", "Un_name19", "Un_name20"}

var Module_2 = []string{"Un_name0", "Un_name1", "Un_name2", "Un_name3", "Un_name4", "Un_name5", "Un_name6",
	"Un_name7", "Un_name8", "Un_name9", "Un_name10", "Un_name11", "Un_name12", "Un_name13", "Un_name14",
	"Un_name15", "Un_name16", "Un_name17", "Un_name18", "Un_name19", "Un_name20"}

var Module_3 = []string{"Un_name0", "Un_name1", "Un_name2", "Un_name3", "Un_name4", "Un_name5", "Un_name6",
	"Un_name7", "Un_name8", "Un_name9", "Un_name10", "Un_name11", "Cn_name12", "Un_name13", "Un_name14",
	"Un_name15", "Un_name16", "Un_name17", "Un_name18", "Un_name19", "Un_name20"}

var Module_4 = []string{"Un_name0", "Un_name1", "Un_name2", "Un_name3", "Un_name4", "Un_name5", "Un_name6",
	"Un_name7", "Un_name8", "Un_name9", "Un_name10", "Un_name11", "Cn_name12", "Un_name13", "Un_name14",
	"Un_name15", "Un_name16", "Un_name17", "Un_name18", "Un_name19", "Un_name20"}

func ReadCoilJsonMaker(a uint16, b []int, leng uint16) interface{} {

	alias_arr := map[string]int{}
	println("aa", a, b, leng)

	for i := 0; i < int(leng); i++ {
		alias_arr[ReadCoil_alias[int(a)+i]] = b[i]
	}
	jsonmaker, err := json.Marshal(alias_arr)
	if err != nil {
		panic(err)
	}
	return jsonmaker
}

func ReadCoilInJsonMaker(a uint16, b []int, leng uint16) []byte {

	alias_arr := map[string]interface{}{}
	println("aa", a, b, leng)
	println("")

	for i := 0; i < int(leng); i++ {
		alias_arr[ReadCoil_alias[int(a)+i]] = b[i]
	}
	jsonmaker, err := json.Marshal(alias_arr)
	if err != nil {
		panic(err)
	}

	println("Jsonmaker Type:", jsonmaker)
	return jsonmaker
}

func ReadHoldRegJsonMaker(a uint16, b []uint16, leng uint16) interface{} {

	alias_arr := map[string]uint16{}
	println("aa", a, b, leng)

	for i := 0; i < int(leng); i++ {
		alias_arr[ReadHoldReg_alias[int(a)+i]] = b[i]
	}
	jsonmaker, err := json.Marshal(alias_arr)
	if err != nil {
		panic(err)
	}
	return jsonmaker
}

func ReadRegInJsonMaker(a uint16, b []uint16, leng uint16) interface{} {
	alias_arr := map[string]uint16{}

	println("aa", a, b, leng)

	for i := 0; i < int(leng); i++ {
		alias_arr[ReadRegIn_alias[int(a)+i]] = b[i]
	}

	jsonmaker, err := json.Marshal(alias_arr)
	if err != nil {
		panic(err)
	}
	return jsonmaker
}

func JsonMaker(a uint16, leng uint16, ReadHoldRegdata float32) []byte {
	ReadHoldReg_arr := map[string]interface{}{}
	ReadRegIn_arr := map[string]interface{}{}
	ReadCoil_arr := map[string]interface{}{}
	ReadCoilIn_arr := map[string]interface{}{}

	result := []map[string]interface{}{}

	for i := 0; i < int(leng); i++ {
		ReadHoldReg_arr["kk"+ReadHoldReg_alias[int(a)+i]] = ReadHoldRegdata
	}
	//for i := 0; i < int(leng); i++ {
	//	ReadRegIn_arr[ReadRegIn_alias[int(a)+i]] = ReadRegIndata[i]
	//}
	//
	//for i := 0; i < int(leng); i++ {
	//	ReadCoil_arr[ReadCoil_alias[int(a)+i]] = ReadCoildata[i]
	//}
	//for i := 0; i < int(leng); i++ {
	//	ReadCoilIn_arr[ReadCoilIn_alias[int(a)+i]] = ReadCoilIndata[i]
	//}

	result = append(result, ReadHoldReg_arr, ReadRegIn_arr, ReadCoil_arr, ReadCoilIn_arr)

	println("result:", result)
	jsonmaker, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}
	return jsonmaker
}
