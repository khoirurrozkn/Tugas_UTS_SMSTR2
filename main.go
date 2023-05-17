package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"
	"utsstrukdat/controller/user"
	"utsstrukdat/view/dashboard"
	"utsstrukdat/db"
)

func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	clear()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("-----------------------------------------------")
		fmt.Println("                    MENU                       ")
		fmt.Println("-----------------------------------------------")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. Exit")
		fmt.Println("-----------------------------------------------")
		fmt.Println("Masukan pilihan : ")

		var inputMenu int
		fmt.Scan(&inputMenu)
		scanner.Scan()

		if inputMenu == 1 {
			clear()
			fmt.Println("-----------------------------------------------")
			fmt.Println("                 REGISTER                      ")
			fmt.Println("-----------------------------------------------")
			fmt.Println("Masukan Username : ")

			scanner.Scan()
			username := scanner.Text()

			fmt.Println("Masukan password : ")

			scanner.Scan()
			password := scanner.Text()

			fmt.Println("Masukan verifikasi password : ")

			scanner.Scan()
			verPassword := scanner.Text()

			dataUser := db.FieldUser{
				Username: username,
				Password: password,
			}

			kode := userController.Register(dataUser, verPassword)
			fmt.Println("-----------------------------------------------")
			if kode == 200 {
				fmt.Println("Akun sudah di buat")
			} else if kode == 400 {
				fmt.Println("Password tidak sama")
			} else if kode == 409 {
				fmt.Println("Username sudah di gunakan")
			}
			time.Sleep(1 * time.Second)
			clear()
		} else if inputMenu == 2 {
			clear()
			fmt.Println("-----------------------------------------------")
			fmt.Println("                     LOGIN                     ")
			fmt.Println("-----------------------------------------------")
			fmt.Println("Masukan username : ")

			scanner.Scan()
			username := scanner.Text()

			fmt.Println("Masukan password : ")

			scanner.Scan()
			password := scanner.Text()

			dataUser := db.FieldUser{
				Username: username,
				Password: password,
			}

			token := userController.Login(dataUser)
			
			if token.Username != "" {
				fmt.Println("-----------------------------------------------")
				fmt.Println("Selamat datang ", token.Username)
				time.Sleep(1 * time.Second)
				clear()
				dashboard.Dashboard(token)
				continue
			}
			fmt.Println("-----------------------------------------------")
			fmt.Println("Username atau password salah")
			time.Sleep(1 * time.Second)
			clear()
		} else if inputMenu == 3 {
			break
		}else {
			fmt.Println("Format tidak di temukan")
			time.Sleep(1 * time.Second)
			clear()
		}
	}
}