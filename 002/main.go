package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := "1234password"
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {

		fmt.Println(err)
	}
	//fmt.Println(bs)
	var ss string
	fmt.Scanln(&ss)
	fmt.Println("password inpout ", ss)
	err = bcrypt.CompareHashAndPassword(bs, []byte(ss))
	if err == nil {

		fmt.Println("Pawwrod the same")
	}

}
