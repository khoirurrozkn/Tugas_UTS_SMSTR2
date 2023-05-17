package dashboard

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"
	"utsstrukdat/controller/message"
	"utsstrukdat/controller/post"
	"utsstrukdat/controller/user"
	"utsstrukdat/view/post"
	"utsstrukdat/db"
)

func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}


func Dashboard(token db.FieldUser){
	scanner := bufio.NewScanner(os.Stdin)
	var category string
	var pilihCategory int
	for {
		dashboard:
		fmt.Println("-----------------------------------------------")
		fmt.Println("                  DASHBOARD                    ")
		fmt.Println("-----------------------------------------------")
		fmt.Println("1. Beranda")
		fmt.Println("2. Pesan")
		fmt.Println("3. Cari akun")
		fmt.Println("4. Cari Kategori")
		fmt.Println("5. Posting")
		fmt.Println("6. Logout")

		var inputDashboard int
		fmt.Scan(&inputDashboard)
		scanner.Scan()

		if inputDashboard == 1 {
			clear()
			post := postController.ShowPost()
			fmt.Println("-----------------------------------------------")
			fmt.Println("                  BERANDA                      ")
			fmt.Println("-----------------------------------------------\n")
			
			for _, val := range *post {
				fmt.Println("-----------------------------------------------")
				fmt.Println("Author : ", val.Author)
				fmt.Println("-----------------------------------------------")
				fmt.Println("Kategori : ", val.Category)
				fmt.Println("Title : ", val.Title)
				fmt.Println("Body : ", val.Body)
				fmt.Println("-----------------------------------------------\n")
			}

		} else if inputDashboard == 2 {
			clear()
			for {
				
				var pilihPesan int
				fmt.Println("-----------------------------------------------")
				fmt.Println("                     PESAN                     ")
				fmt.Println("-----------------------------------------------")
				fmt.Println("1. Lihat Pesan")
				fmt.Println("2. Kirim Pesan")
				fmt.Println("3. Kembali ke Dashboard")
				fmt.Println("Masukan pilihan")
				fmt.Scan(&pilihPesan)
				scanner.Scan()

				if pilihPesan == 1 {
					clear()
					message := messageController.ShowMessage(token.Username)
					fmt.Println("-----------------------------------------------")
					fmt.Println("                  LIHAT PESAN                  ")
					fmt.Println("-----------------------------------------------\n")

					for _, val := range *message {
						fmt.Println("-----------------------------------------------")
						fmt.Println("From : ", val.From)
						fmt.Println("Message : ", val.Message)
						fmt.Println("-----------------------------------------------\n")
					}

				} else if pilihPesan == 2 {
					clear()

					fmt.Println("-----------------------------------------------")
					fmt.Println("                   KIRIM PESAN                 ")
					fmt.Println("-----------------------------------------------")
					fmt.Println("Masukan username yang mau anda kirimkan pesan")

					scanner.Scan()
					penerima := scanner.Text()


					fmt.Println("Masukan Pesan")

					scanner.Scan()
					pesan := scanner.Text()

					check := messageController.SendMessage(penerima, token.Username, pesan)
					fmt.Println("-----------------------------------------------")
					if check == 200 {
						fmt.Println("Pesan anda sudah terkirim ke ", penerima)
					} else {
						fmt.Println("Username penerima tidak di temukan")
					}
					time.Sleep(1 * time.Second)
					clear()
				} else if pilihPesan == 3 {
					clear()
					goto dashboard
				} else {
					fmt.Println("-----------------------------------------------")
					fmt.Println("Format tidak di temukan")
					time.Sleep(1 * time.Second)
					clear()
				}
			}
		} else if inputDashboard == 3 {
			clear()
			fmt.Println("-----------------------------------------------")
			fmt.Println("                  CARI AKUN                    ")
			fmt.Println("-----------------------------------------------")
			fmt.Println("Masukan username")

			scanner.Scan()
			username := scanner.Text()

			akun := userController.SearchAccount(username)

			if akun == ""{
				fmt.Println("-----------------------------------------------")
				fmt.Println("Akun tidak di temukan")
			}else{
				clear()
				fmt.Println("-----------------------------------------------")
				fmt.Println("Username : ", akun)
				fmt.Println("-----------------------------------------------")
				post := userController.ShowPostByAccount(username)

				if post != nil && *post != nil{
					for _, val := range *post {
						fmt.Println("-----------------------------------------------")
						fmt.Println("Author : ", val.Author)
						fmt.Println("-----------------------------------------------")
						fmt.Println("Category : ", val.Category)
						fmt.Println("Title : ", val.Title)
						fmt.Println("Body : ", val.Body)
						fmt.Println("-----------------------------------------------")
					}
				}else{
					fmt.Println("-----------------------------------------------")
					fmt.Println("Akun tersebut belum memposting apapun")
				}
			}
		} else if inputDashboard == 4 {
			clear()
			fmt.Println("-----------------------------------------------")
			fmt.Println("                 CARI KATEGORI                 ")
			fmt.Println("-----------------------------------------------")
			fmt.Println("Masukan Kategori yang mau anda cari")
			fmt.Println("1. Blog")
			fmt.Println("2. Programming")
			fmt.Scan(&pilihCategory)
			scanner.Scan()

			if pilihCategory == 1 || pilihCategory == 2 {
				if pilihCategory == 1 {
					category = "Blog"
				} else if pilihCategory == 2 {
					category = "Programming"
				}
				post := postController.ShowByCategory(category)
				
				for _, val := range *post {
					fmt.Println("-----------------------------------------------")
					fmt.Println("Author : ", val.Author)
					fmt.Println("-----------------------------------------------")
					fmt.Println("Kategori : ", val.Category)
					fmt.Println("Title : ", val.Title)
					fmt.Println("Body : ", val.Body)
					fmt.Println("-----------------------------------------------\n")
				}
			} else {
				fmt.Println("-----------------------------------------------")
				fmt.Println("Kategori tidak di temukan")
				time.Sleep(1 * time.Second)
				clear()
			}
		} else if inputDashboard == 5 {
			post.Post(token)
		} else if inputDashboard == 6 {
			clear()
			break
		}else {
			fmt.Println("Format tidak di temukan")
			time.Sleep(1 * time.Second)
			clear()
		}
	}
}