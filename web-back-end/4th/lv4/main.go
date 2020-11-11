package main

import (
	"bufio"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type User struct {
	Account string `json:"account"`
	PasswordMac []byte `json:"passwordMac"`
}

var filename = "data.json"
var users []User

func init(){
	fileString := readFile()
	err := json.Unmarshal([]byte(fileString), &users)
	if err != nil {
		fmt.Println("Failed to unmarshal json: ", err.Error())
	}
}
func main()  {
	selectFunc()
}

func readFile() (content string) {
	file, openFileErr := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE,0777)
	if openFileErr != nil {
		fmt.Println("Failed to open the file")
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var fileBufs []byte
	for {
		fileBuf := make([]byte, 1024)
		n, err := reader.Read(fileBuf)
		if err != nil {
			if err == io.EOF{
				content = string(fileBufs)
				return content
			} else {
				fmt.Println("Failed to read the file")
				return content
			}
		} else {
			fileBufs = append(fileBufs, fileBuf[:n]...)
		}
	}
}

func writeFile(content []byte){
	file, openFileErr := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC,0777)
	if openFileErr != nil {
		fmt.Println("Cannot write the file because opening file failed.")
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.Write(content)
	writer.Flush()
}

func register() {
	var account, password, confirmPassword string
	var minPwdLen = 6
	printDividingLine()
	fmt.Println("Register:")

	fmt.Print("Account:")
	fmt.Scan(&account)
	for i, _ := range users{
		if users[i].Account == account {
			fmt.Println("The account you input has already existed.")
			return
		}
	}

	fmt.Print("Password:")
	fmt.Scan(&password)
	if len(password) <= minPwdLen{
		fmt.Println("The length of the password you input is less than 6 bytes.")
		return
	}

	fmt.Print("Confirm Password:")
	fmt.Scan(&confirmPassword)
	if password != confirmPassword{
		fmt.Println("Unmatched password.")
		return
	}

	passwordMac := getMac(password)
	newUser := User{Account: account, PasswordMac: passwordMac}
	users = append(users, newUser)
	fmt.Println("Register successfully.")
	updateData()
}

func login() {
	var account, password string
	printDividingLine()
	fmt.Print("Login:\n")
	fmt.Print("Account:")
	fmt.Scan(&account)
	fmt.Print("Password:")
	fmt.Scan(&password)
	for i, _ := range users{
		if users[i].Account == account {
			if hmac.Equal(users[i].PasswordMac, getMac(password)){
				fmt.Println("Login successfully.")
				return
			} else {
				fmt.Println("Wrong account or password.")
				return
			}
		}
	}
	fmt.Println("The account you input does not exist.")
}

func selectFunc(){
	for {
		var option string
		fmt.Print("1)Register\n2)Login\n3)Exit\n")
		fmt.Scan(&option)
		if option == "1" {
			register()
		} else if option == "2" {
			login()
		} else if option == "3" {
			return
		} else {
			fmt.Println("Invalid input value")
		}
		printDividingLine()
	}
}

func updateData(){
	fileString, _ := json.Marshal(users)
	writeFile(fileString)
}

func printDividingLine(){
	fmt.Print("------------------------\n")
}

func getMac(message string) []byte{
	key := []byte("N98YOpzU0Mbv6bx^wVnWDvSz5k!875IU")
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(message))
	return mac.Sum(nil)
}