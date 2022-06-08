package main

import (
	"fmt"
	"project1/config"
	"project1/controllers"
	"project1/entities"

	"gorm.io/gorm"
)

var dbConn *gorm.DB
var cond bool = true
var pilihan string = ""

func init() {
	dbConn = config.ConnectDB()
	InitialMigration()
}

func InitialMigration() {
	dbConn.AutoMigrate(&entities.User{})

}

func MenuUtama() {

	// Menu Utama
	// myMenu := []string{"1. Login", "2. Register", "3. Lihat Profile", "4. Edit Profile", "5. Delete Akun", "6. Top Up", "7. Transfer", "8. History Top Up", "9. History Transfer", "10. Friends Profile Check"}

	fmt.Println("===== Menu Utama =====")
	fmt.Println("1. Login")
	fmt.Println("2. Register")
	fmt.Println("0. Keluar")

	// Pilihan Menu Check

	fmt.Print("Masukkan Pilihan : ")
	fmt.Scanln(&pilihan)

	switch pilihan {
	case "1":
		Login()
	case "2":
		Register()
	case "0":
		Keluar()
	default:
		MenuUtama()
	}

}

// Menu Login
func Login() {

	user := entities.User{}

	fmt.Println("===== Login =====")
	fmt.Println("Masukkan No. HP :")
	fmt.Scanln(&user.Phone)
	fmt.Println("Masukkan Password :")
	fmt.Scanln(&user.Password)

	resLogin := controllers.GetUserbyPassword(dbConn, user.Password)

	for _, value := range resLogin {
		for pilihan != "0" {
			fmt.Println("===== Login Berhasil =====")
			fmt.Println("===== Selamat Datang,", value.Name, "=====")
			fmt.Println("===== Menu User =====")
			fmt.Println("1. Profil User")
			fmt.Println("2. Menu Transaksi")
			fmt.Println("3. Riwayat Transaksi")
			fmt.Println("0. Kembali ke Menu Utama")
			fmt.Print("Masukkan Pilihan : ")

			fmt.Scanln(&pilihan)

			switch pilihan {
			case "1":
				// ProfilUser()
			case "2":
				// MenuTrans()
			case "3":
				// RiwayatTrans()
			}

			if pilihan == "0" {
				fmt.Println("===== Kembali ke menu utama? =====")
				fmt.Println("1. Tetap di halaman")
				fmt.Println("2. Kembali ke menu utama")
				fmt.Print("Masukkan Pilihan : ")
				fmt.Scanln(&pilihan)

				if pilihan == "2" {
					resLogin = nil
					MenuUtama()
				}
			}
		}
	}

	if resLogin != nil {
		fmt.Println("==== Akun tidak ditemukan ====")
	}
}

// Menu Register
func Register() {

	user := entities.User{}

	fmt.Println("===== Registrasi =====")
	fmt.Println("Masukkan Nama :")
	fmt.Scanln(&user.Name)
	fmt.Println("Masukkan No. HP :")
	fmt.Scanln(&user.Phone)
	fmt.Println("Masukkan Password :")
	fmt.Scanln(&user.Password)
	fmt.Println("Masukkan Gender :")
	fmt.Scanln(&user.Gender)
	fmt.Println("Masukkan Address :")
	fmt.Scanln(&user.Address)

	resRegis := controllers.CreateUser(dbConn, user)

	if resRegis != nil {
		fmt.Println("Nomor HP sudah terdaftar")
	} else {
		fmt.Println("Registrasi berhasil")
	}
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

func Keluar() {
	fmt.Println("===== Keluar dari aplikasi? =====")
	fmt.Println("1. Kembali ke aplikasi")
	fmt.Println("2. Keluar dari aplikasi")
	fmt.Print("Masukkan Pilihan : ")
	fmt.Scanln(&pilihan)

	if pilihan == "2" {
		cond = false
		fmt.Println("===== Terima Kasih =====")
	}
}

func main() {
	for cond {
		MenuUtama()
	}
}
