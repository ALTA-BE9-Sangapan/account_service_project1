package main

import (
	"database/sql"
	"fmt"
	_Connect "project1/config"
)

var dbConnec *sql.DB

func init() {
	dbConnec = _Connect.ConnectDB()
}

func MenuUtama() {

	// Menu Utama
	myMenu := []string{"1. Login", "2. Register", "3. Lihat Profile", "4. Edit Profile", "5. Delete Akun", "6. Top Up", "7. Transfer", "8. History Top Up", "9. History Transfer", "10. Friends Profile Check"}

	fmt.Println("===== Menu Utama =====")

	for _, ok := range myMenu {
		fmt.Println(ok)
	}

	// Pilihan Menu Check
	var menu int
	fmt.Print("Pilih Menu : ")
	fmt.Scanf("%s", &menu)

	switch menu {
	case 1:
		Login()
	case 2:
		Register()
	case 3:
		UserProfile()
	default:
		MenuUtama()
	}

}

// Menu Login
func Login() {
	var user_phone, password string

	fmt.Println("===== Login =====")
	fmt.Println("Masukkan No. HP :")
	fmt.Scanf("%s", &user_phone)
	fmt.Println("Masukkan Password :")
	fmt.Scanf("%s", &password)

}

// Menu Register
func Register() {
	var user_phone, user_name, password string

	fmt.Println("===== Registerasi =====")
	fmt.Println("Masukkan No. HP :")
	fmt.Scanf("%s", &user_phone)
	fmt.Println("Masukkan Nama :")
	fmt.Scanf("%s", &user_name)
	fmt.Println("Masukkan Password :")
	fmt.Scanf("%s", &password)
}

func UserProfile() {

	fmt.Println("===== Profile =====")

}
func UpdateProfile() {

	fmt.Println("===== Update Profile =====")
	fmt.Println("1. Ganti Nama\n2. Ubah Tanggal Lahir\n3. Ubah Jenis Kelamin\n4. Ubah Alamat\n99. Kembali")

	var input int
	fmt.Println("Masukkan Pilihan : ")
	fmt.Scanf("%d", &input)
	switch input {
	case 1:

	}

}

func main() {
	defer dbConnec.Close()
	MenuUtama()
}
