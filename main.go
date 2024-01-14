package main

import (
	//"os"

	"github.com/gweffectx/safedav/cmd"
	//"github.com/gweffectx/safedav/encrypt"
)

//"github.com/gweffectx/safedav/cmd"

func main() {
	// v1, err := strconv.ParseInt("3178757376", 0, 64)
	// if err == nil {

	// }
	// println(v1)
	// a1 := encrypt.NewFileNameBase64()
	// s1 := a1.Encrypt("12312aaaaaaaaaaaaaaaaa31313131313.mp4")
	// s2 := a1.Decrypt(s1)
	// println(s1)
	// println(s2)
	// key := []byte("wumansgygoaescbc")
	// // reader, err1 := os.Open("C:\\Users\\gw\\Downloads\\摩登家庭.Modern.Family.S10E22.End.中英字幕.WEB.720p-人人影视.mp4")
	// // writer, err2 := os.OpenFile("C:\\Users\\gw\\Downloads\\摩登家庭.Modern.Family.S10E22.End.中英字幕.WEB.720p-人人影视加密.mp4", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	// reader, err1 := os.Open("C:\\Users\\gw\\Downloads\\摩登家庭.Modern.Family.S10E22.End.中英字幕.WEB.720p-人人影视加密.mp4")
	// writer, err2 := os.OpenFile("C:\\Users\\gw\\Downloads\\摩登家庭.Modern.Family.S10E22.End.中英字幕.WEB.720p-人人影视解密.mp4", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	// if err1 != nil {
	// }
	// if err2 != nil {
	// }
	// // // safeWrite := encrypt.NewEncryptWriter(writer, key)
	// //safeReader := encrypt.NewEncryptReader(reader, key)
	// safeReader := encrypt.NewDecryptReader(reader, key)

	// buffer := make([]byte, 1024)
	// for {
	// 	//readLength, e1 := reader.Read(buffer)
	// 	readLength, e1 := safeReader.Read(buffer)
	// 	if e1 != nil {
	// 		break
	// 	}
	// 	if readLength == 0 {
	// 		break
	// 	}
	// 	//safeWrite.Write(buffer[:readLength])
	// 	writer.Write(buffer[:readLength])
	// }

	// reader.Close()
	// writer.Close()

	// key := []byte("wumansgygoaescbc")
	// n1 := encrypt.NewEncryptAesCtr(key)
	// reader, err1 := os.Open("D:\\test122402加密.mp3")
	// writer, err2 := os.OpenFile("D:\\test122402解密.mp3", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	// if err1 != nil {
	// }
	// if err2 != nil {

	// }
	// buffer := make([]byte, 1024)
	// for {
	// 	readLength, e1 := reader.Read(buffer)
	// 	if e1 != nil {
	// 		break
	// 	}
	// 	if readLength == 0 {
	// 		break
	// 	}
	// 	res := n1.Decrypt(buffer[:readLength])
	// 	writer.Write(res)
	// }
	// reader.Close()
	// writer.Close()

	cmd.Execute()
}
