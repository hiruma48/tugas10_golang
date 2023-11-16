package main

import (
	"fmt"
	"time"
)

type biodata struct {
	nama string
	hobi string
}

func main() {
	var layoutFormat, value string
	value = "18/12/1988 WIB"
	var x1 biodata
	x1.nama = "ari"
	x1.hobi = "ngoding"

	fmt.Println("nama : ", x1.nama, "\n", time.Now())
	time.Sleep(time.Second * 5)
	fmt.Println("hobi : ", x1.hobi, "\n", time.Now())
	time.Sleep(time.Second * 5)
	fmt.Println("tanggal lahir :", layoutFormat, value, "\n", time.Now())
}
