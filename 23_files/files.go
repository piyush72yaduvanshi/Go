package main

import (
	"fmt"
	"os"
)

func main() {
	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	// log the error
	// 	panic(err)
	// }

	// fileInfo, err := f.Stat()
	// if err != nil {
	// 	// log the error
	// 	panic(err)
	// }

	// fmt.Println("file name:", fileInfo.Name())
	// fmt.Println("file or folder:", fileInfo.IsDir())
	// fmt.Println("file size:", fileInfo.Size())
	// fmt.Println("file permission:", fileInfo.Mode())
	// fmt.Println("file modified at:", fileInfo.ModTime())

	// read file
	// f, err1 := os.Open("example.txt")
	// if err1 != nil {
	// 	panic(err1)
	// }

	// fileInfo, err := f.Stat()
	// if err != nil {
	// 	panic(err)
	// }
	// size := fileInfo.Size()

	// defer f.Close()

	// buf := make([]byte, size)
	// d, err1 := f.Read(buf)
	// if err1 != nil {
	// 	panic(err1)
	// }

	// fmt.Println("data", d, string(buf))

	// for i := 0; i < len(buf); i++ {
	// 	println("data", d, string(buf[i]))
	// }

	// kya hme yahi method use karna chahiye ya nhi
	// Answer hai nhi because agar file ka size bahut bada hoga to ye method memory ko bahut zyada use karega aur program crash ho sakta hai. Isliye better hai ki hum streaming fashion me read karein.
	// data, err := os.ReadFile("example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(data))

	// // read folders

	// dir, err := os.Open("../")
	// if err != nil {
	// 	panic(err)
	// }

	// defer dir.Close()
	//RedDir returns a slice of DirEntry
	//agar humne 10 pass kiya to ye 10 hi files ko read karega aur agar humne -1 pass kiya to ye saari files ko read karega
	//sabko red karne ke liye hum -1 pass karenge

	// filesInfo, err := dir.ReadDir(10)

	// for _, fi := range filesInfo {
	// 	fmt.Println(fi.Name(), fi.IsDir())
	// }

	// // // create a file
	// f, err4 := os.Create("example2.txt")
	// if err4 != nil {
	// 	panic(err4)
	// }
	// //  file jab bhi open krenge to hamesha close karna chahiye. Agar file close nahi karenge to memory leak ho sakta hai aur program crash ho sakta hai. Isliye hamesha file ko close karna chahiye.
	// defer f.Close()

	// f.WriteString("hi go ")
	// f.WriteString("nice language to learn ")
	// // hmne yaha pe string ko write kiya hai. Agar hme bytes write karna hai to hm bytes ko write karenge. Jaise niche example me kiya gaya hai.
	// bytes := []byte("Hello Golang")
	// f.Write(bytes)

	// read and write to another file (streaming fashion)
	// ek file se read krke dusre file me write karna hai to hm streaming fashion me read karenge aur write karenge. Isliye hmne bufio package ka use kiya hai. Ye package hme buffered I/O provide karta hai. Iska matlab ye hai ki ye data ko buffer me store karta hai aur jab buffer full ho jata hai to hi data ko write karta hai. Isse performance improve hoti hai.

	// sourceFile, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer sourceFile.Close()

	// destFile, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer destFile.Close()

	// // bufio package ka use karke read and write karenge. Ye package hme buffered I/O provide karta hai. Iska matlab ye hai ki ye data ko buffer me store karta hai aur jab buffer full ho jata hai to hi data ko write karta hai. Isse performance improve hoti hai.

	// reader := bufio.NewReader(sourceFile)
	// writer := bufio.NewWriter(destFile)

	// for {
	// 	b, err := reader.ReadByte()
	// 	if err != nil {
	// 		// agar file ka end aa gaya hai to ye error return karega. Isliye hme ye check karna chahiye ki error EOF hai ya nahi. Agar EOF hai to break kar denge aur agar EOF nahi hai to panic kar denge.
	// 		if err.Error() != "EOF" {
	// 			panic(err)
	// 		}

	// 		break
	// 	}

	// 	e := writer.WriteByte(b)
	// 	if e != nil {
	// 		panic(e)
	// 	}

	// }
	// // Isme hme buffer me data ko store karenge aur jab buffer full ho jata hai to hi data ko write karta hai.
	// writer.Flush()

	// fmt.Println("written to new file succesfully")

	// delete a file

	// file, err := os.Open("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer file.Close()

	err := os.Remove("example2.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(" file deleted successfully")

}
