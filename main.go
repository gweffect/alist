package main

import (
	"github.com/gweffectx/safedav/cmd"
)

func main() {
	// key := []byte("wumansgygoaescbc")
	// //reader, err1 := os.Open("D:\\480803359.mp3")
	// //writer, err2 := os.OpenFile("D:\\test122403加密.mp3", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	// reader, err1 := os.Open("D:\\test122403加密.mp3")
	// writer, err2 := os.OpenFile("D:\\test122403解密.mp3", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	// if err1 != nil {
	// }
	// if err2 != nil {
	// }
	// //safeWrite := encrypt.NewEncryptWriter(writer, key)
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
