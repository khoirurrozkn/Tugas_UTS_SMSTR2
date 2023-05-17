package post

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"
	"utsstrukdat/controller/post"
	"utsstrukdat/db"
)

func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Post(token db.FieldUser){
	var inputPost int
	var category string
	var pilihCategory int
	var title string
	scanner := bufio.NewScanner(os.Stdin)
	clear()
	fmt.Println("-----------------------------------------------")
	fmt.Println("                   POSTING                     ")
	fmt.Println("-----------------------------------------------")
	fmt.Println("1. Buat cerita                                 ")
	fmt.Println("2. Edit cerita                                 ")
	fmt.Println("3. Hapus cerita                                ")
	fmt.Println("Masukan pilihan :                              ")
	fmt.Scan(&inputPost)
	scanner.Scan()

	if inputPost == 1 {
		clear()
		fmt.Println("-----------------------------------------------")
		fmt.Println("                  BUAT CERITA                  ")
		fmt.Println("-----------------------------------------------")
		fmt.Println("1. Blog")
		fmt.Println("2. Programming")
		fmt.Println("Masukan kategori")

		fmt.Scan(&pilihCategory)
		scanner.Scan()

		if pilihCategory == 1 {
			category = "Blog"
		} else {
			category = "Programming"
		}

		fmt.Println("Masukan judul")
		fmt.Scan(&title)
		scanner.Scan()

		fmt.Println("Masukan cerita")

		scanner.Scan()
		body := scanner.Text()

		dataPost := &db.FieldPost{
			Author:   token.Username,
			Category: category,
			Title:    title,
			Body:     body,
		}

		kode := postController.InsertPost(dataPost)

		fmt.Println("-----------------------------------------------")
		if kode == 200 {
			fmt.Println("postingan telah di buat")
		} else {
			fmt.Println("Judul telah di gunakan")
		}
		time.Sleep(1 * time.Second)
		clear()
	} else if inputPost == 2 {
		clear()
		fmt.Println("-----------------------------------------------")
		fmt.Println("                  EDIT CERITA                  ")
		fmt.Println("-----------------------------------------------")
		fmt.Println("Masukan judul post yang mau di ubah")
		fmt.Scan(&title)
		scanner.Scan()

		fmt.Println("Masukan cerita baru")

		scanner.Scan()
		body := scanner.Text()

		kode := postController.UpdatePost(title, body, token.Username)

		fmt.Println("-----------------------------------------------")
		if kode == 200 {
			fmt.Println("Postingan sudah di update")
		} else if kode == 403 {
			fmt.Println("Postingan tersebut bukan milik anda")
		} else {
			fmt.Println("Postingan tidak di temukan")
		}
		time.Sleep(1 * time.Second)
		clear()
	} else if inputPost == 3 {
		clear()
		fmt.Println("-----------------------------------------------")
		fmt.Println("                  HAPUS CERIA                  ")
		fmt.Println("-----------------------------------------------")
		fmt.Println("Masukan judul post yang mau di hapus")
		fmt.Scan(&title)

		kode := postController.DeletePost(title, token.Username)

		fmt.Println("-----------------------------------------------")
		if kode == 200 {
			fmt.Println("Postingan sudah di hapus")
		} else if kode == 403 {
			fmt.Println("Postingan tersebut bukan milik anda")
		} else {
			fmt.Println("Postingan tidak di temukan")
		}
		time.Sleep(1 * time.Second)
		clear()
	} else {
		fmt.Println("Format tidak di temukan")
		time.Sleep(1 * time.Second)
		clear()
	}
	
}