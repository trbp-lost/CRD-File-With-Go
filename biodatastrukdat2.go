package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func createFile() {
	var fileName string
	fmt.Print("(contoh : hello nanti yang dibuat hello.txt)\nMasukkan nama file yang akan dibuat : ")
	fmt.Scan(&fileName)

	file, err := os.Create(fileName + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File created successfully")
	defer file.Close()

	stats, err := file.Stat()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("\nFile Name     : %s\n", stats.Name())
	fmt.Printf("Time Created    : %v\n", stats.ModTime().Format("15:04:05"))
}

func inputData(filename string) {
	var uas, uts int32
	var total float32
	var nama, npm, kelas, pil string

	f, err := os.Open(filename + ".txt")
	if err != nil {
		fmt.Print("\nApakah anda ingin membuat file (y/!n)? ")
		fmt.Scan(&pil)

		if strings.ToLower(pil) == "y" {
			createFile()
		} else {
			menu()
		}
	}
	defer f.Close()

	fmt.Println("-------Input-------")
	fmt.Print("Nama\t: ")
	fmt.Scan(&nama)
	fmt.Print("NPM\t: ")
	fmt.Scan(&npm)
	fmt.Print("Kelas\t: ")
	fmt.Scan(&kelas)
	fmt.Print("UAS\t: ")
	fmt.Scan(&uas)
	fmt.Print("UTS\t: ")
	fmt.Scan(&uts)

	total = (float32(uts) + float32(uas)) / 2

	fmt.Printf("\nNama yang dimasukkan            : %s \n", nama)
	fmt.Printf("NPM yang dimasukkan             : %s \n", npm)
	fmt.Printf("Kelas yang dimasukkan           : %s \n", kelas)
	fmt.Printf("Nilai UTS yang dimasukkan       : %d \n", uts)
	fmt.Printf("Nilai UAS yang dimasukkan       : %d \n", uas)
	fmt.Printf("Nilai Rata-rata yang didapatkan : %3.2f \n", total)

	fileDataAPPENDS, err := os.OpenFile(filename+".txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer fileDataAPPENDS.Close()

	data := nama + "\t" + npm + "\t" + kelas + "\t" + fmt.Sprintf("%v", uts) + "\t" + fmt.Sprintf("%v", uas) + "\t" + fmt.Sprintf("%v", total) + "\n"
	if _, err = fileDataAPPENDS.WriteString(data); err != nil {
		panic(err)
		fmt.Println(err.Error())
		fmt.Print("\nApakah anda ingin membuat file (y/!n)? ")
		fmt.Scan(&pil)

		if strings.ToLower(pil) == "y" {
			createFile()
		} else {
			menu()
		}
	}
}

func readFile(filename string) {
	var pil string

	fileContents, err := ioutil.ReadFile(filename + ".txt")
	if err != nil {
		fmt.Println(err.Error())
		fmt.Print("\nApakah anda ingin membuat file (y/!n)? ")
		fmt.Scan(&pil)

		if strings.ToLower(pil) == "y" {
			createFile()
		} else {
			menu()
		}
		return
	}
	fmt.Println("Nama\tNPM\tKelas\tUTS\tUAS\tTotal")
	fmt.Println(string(fileContents))
}

func menu() {
	var pil int16
	var fileName string

	fmt.Println("\n\n-------Program Pilihan-------")
	fmt.Println("1. Create File")
	fmt.Println("2. Input Data")
	fmt.Println("3. Read Data")
	fmt.Println("4. Delete Data By Content")
	fmt.Println("5. Delete File")
	fmt.Println("0. Exit Program")
	fmt.Print("Masukkan Pilihan [0-5] : ")
	fmt.Scan(&pil)
	fmt.Println()

	switch pil {
	case 0:
		fmt.Println("Thanks For Using")
		os.Exit(3)
	case 1:
		createFile()
	case 2:
		fmt.Println("(contoh : hello nanti akan dimassukan data ke hello.txt)")
		fmt.Print("Masukkan nama file : ")
		fmt.Scan(&fileName)
		inputData(fileName)
	case 3:
		fmt.Println("(contoh : hello nanti yang dibuka hello.txt)")
		fmt.Print("Masukkan nama file : ")
		fmt.Scan(&fileName)
		readFile(fileName)
	case 4:
		fmt.Println("(contoh : hello nanti yang dihapus isi data dari hello.txt)")
		fmt.Print("Masukkan nama file : ")
		fmt.Scan(&fileName)
		deleteDataByContent(fileName)
	case 5:
		fmt.Println("(contoh : hello nanti menghapus file hello.txt)")
		fmt.Print("Masukkan nama file : ")
		fmt.Scan(&fileName)
		deleteFile(fileName)
	default:
		fmt.Println("invalid input")
	}
	menu()
}

func main() {
	var pil string

	for i := 0; i < 1; i-- {
		menu()
		fmt.Print("Apakah ingin lanjut (y/!n)? ")
		fmt.Scan(&pil)

		if strings.ToLower(pil) == "y" {
			continue
		} else {
			break
		}
	}
}
