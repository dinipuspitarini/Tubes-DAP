package main

import (
	"fmt"
	"bufio"
	"os"
)

type identitas struct {
	kode string
	merk string
	tipe string
	harga int
	stock int
}

type pembeli struct {
	nama string
	alamat string
	umur int
	merk string
	angsuran int
	lamaKredit int
	tahunBeli int
	sisaAngsuran int
}

var (
	mobil []identitas
	pelanggan []pembeli
) 

func menu() {
	var pilihan int
	
	fmt.Println("===============================================================")
	fmt.Println("=                                                             =")
	fmt.Println("=       Selamat  Datang  di  Dini  dan  Ridho  Showroom       =")
	fmt.Println("=                                                             =")
	fmt.Println("===============================================================")
	for true {
		fmt.Println("")
		fmt.Println("Menu Utama")
		fmt.Println("")
		fmt.Println("1. Tambah mobil")
		fmt.Println("2. Tambah pembeli")
		fmt.Println("3. Cari mobil")
		fmt.Println("4. Cari pembeli")
		fmt.Println("5. Sisa angsuran pembeli")
		fmt.Println("6. Menampilkan 10 data pembeli yang akan lunas 1 tahun lagi")
		fmt.Println("7. Menampilkan 10 data pembeli dari tahun 2016")
		fmt.Println("8. Exit")
		fmt.Println("")
		fmt.Print("Pilihan anda : ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1 :
			tambahMobil()
		case 2 :
			tambahPembeli()
		case 3 :
			cariMobil()
		case 4 :
			cariPembeli()
		case 5 :
			hitungAngsuran()
		case 6 : 
			tampilmaulunas()
		case 7 :
			tampillebih2016(pelanggan)
		default :
			fmt.Println("Terimakasih telah mengunjungi Dini dan Ridho Showroom")
		}
	}
}

func tambahMobil() {
	var (
		kode string
		merk string
		tipe string
		harga int
		stock int
	)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Masukkan kode mobil : ")
	scanner.Scan()
	kode = scanner.Text()
	fmt.Println("Masukkan merk mobil : ")
	scanner.Scan()
	merk = scanner.Text()
	fmt.Println("Masukkan tipe mobil : ")
	scanner.Scan()
	tipe = scanner.Text()
	fmt.Println("Masukkan harga mobil : ")
	fmt.Scanln(&harga)
	fmt.Println("Masukkan stock mobil : ")
	fmt.Scanln(&stock)

	data := identitas{kode: kode, merk: merk, tipe: tipe, harga: harga, stock: stock}
	mobil = append(mobil, data)
}

func tambahPembeli() {
	var (
		nama string
		alamat string
		umur int
		merk string
		angsuran int
		lamaKredit int
		tahunBeli int
		sisaAngsuran int
	)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Masukkan nama pembeli : ")
	scanner.Scan()
	nama = scanner.Text()
	fmt.Println("Masukkan alamat pribadi : ")
	scanner.Scan()
	alamat = scanner.Text()
	fmt.Println("Masukkan umur anda : ")
	fmt.Scanln(&umur)
	fmt.Println("Masukkan merk mobil : ")
	scanner.Scan()
	merk = scanner.Text()
	fmt.Println("Masukkan pilihan angsuran per bulan : ")
	fmt.Scanln(&angsuran)
	fmt.Println("Jangka waktu kredit : ")
	fmt.Scanln(&lamaKredit)
	fmt.Println("Tahun pembelian mobil : ")
	fmt.Scanln(&tahunBeli)
	fmt.Println("Sisa angsuran yang anda miliki : ")
	fmt.Scanln(&sisaAngsuran)

	data := pembeli{nama: nama, alamat: alamat, umur: umur, merk: merk, angsuran: angsuran, lamaKredit: lamaKredit, tahunBeli: tahunBeli, sisaAngsuran: sisaAngsuran}
	pelanggan = append(pelanggan, data)
}

func cariMobil() {
	var kode string
	var found bool

	found = false
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Masukkan kode mobil yang anda cari : ")
	scanner.Scan()
	kode = scanner.Text()
	no := 0
	fmt.Printf("| No | Kode Mobil | Merk Mobil |    Tipe Mobil    | Harga Mobil | Jumlah Stock Mobil |\n")
	for i := 0 ; !found && i < len(mobil) ; i++ {
		if kode == mobil[i].kode {
			no++
			fmt.Printf("| %-2d |", no)
			fmt.Printf(" %-11s|", mobil[i].kode)
			fmt.Printf(" %-11s|", mobil[i].merk)
			fmt.Printf(" %-17s|", mobil[i].tipe)
			fmt.Printf(" %-12d|", mobil[i].harga)
			fmt.Printf(" %-19d|\n", mobil[i].stock)
		}
	}
}

func cariPembeli() {
	var nama string
	var found bool
	var index int

	found = false
	index = -1
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Masukkan nama pembeli : ")
	scanner.Scan()
	nama = scanner.Text()
	for i := 0 ; !found && i < len(pelanggan) ; i++ {
		if nama == pelanggan[i].nama {
			found = true
			index = i		
		} 
	}
	
	no := 0
	if index >= 0 {
		no++
		fmt.Println("")
		fmt.Println("Detail Identitas Pembeli")
		fmt.Printf("| No |   Nama Pembeli   |    Alamat Pribadi     | Umur Pembeli |  Merk Mobil  | Angsuran Yang Diambil |Lama Kredit (Bulan)|Tahun Pembelian|  Sisa Angsuran  |\n")
		fmt.Printf("| %-2d |", no)
		fmt.Printf(" %-17s|", pelanggan[index].nama)
		fmt.Printf(" %-22s|", pelanggan[index].alamat)
		fmt.Printf(" %-13d|", pelanggan[index].umur)
		fmt.Printf(" %-13s|", pelanggan[index].merk)
		fmt.Printf(" %-22d|", pelanggan[index].angsuran)
		fmt.Printf(" %-18d|", pelanggan[index].lamaKredit)
		fmt.Printf(" %-14d|", pelanggan[index].tahunBeli)
		fmt.Printf(" %-16d|\n", pelanggan[index].sisaAngsuran)
	}
}

func hitungAngsuran() {
	var index, dp int
	var nama string

	index = -1
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Masukkan nama pembeli : ")
	scanner.Scan()
	nama = scanner.Text()
	for i := 0 ; i < len(pelanggan) ; i++ {
		if pelanggan[i].nama == nama {
			index = i
		}
	}

	if index != -1 {
		fmt.Println("Masukan DP : ")
		fmt.Scanln(&dp)
		pelanggan[index].sisaAngsuran -= dp
		pelanggan[index].lamaKredit -= 1
		fmt.Println("Sisa angsuran pembeli : ", pelanggan[index].sisaAngsuran)
		fmt.Println("Lama kredit : ", pelanggan[index].lamaKredit)
	} else {
		fmt.Println("Data tersebut tidak ditemukan")
	}
}

func tampilmaulunas(){
	var ketemu int

	for i := len(pelanggan) ; i > 0 ; i-- {
		for j := 1 ; j<i ; j++ {
			if pelanggan[j-1].lamaKredit > pelanggan[j].lamaKredit {
				t := pelanggan[j]
				pelanggan[j] = pelanggan[j-1]
				pelanggan[j-1] = t
			}
		}
	}

	no := 0
	fmt.Printf("| No |   Nama Pembeli   |    Alamat Pribadi     | Umur Pembeli |  Merk Mobil  | Angsuran Yang Diambil |Lama Kredit (Bulan)|Tahun Pembelian|  Sisa Angsuran  |\n")
	for i:=0 ; i < len(pelanggan) && i < 10 ; i++ {
		if pelanggan[i].lamaKredit <= 12 {
			ketemu++
			no++
			fmt.Printf("| %-2d |", no)
			fmt.Printf(" %-17s|", pelanggan[i].nama)
			fmt.Printf(" %-22s|", pelanggan[i].alamat)
			fmt.Printf(" %-13d|", pelanggan[i].umur)
			fmt.Printf(" %-13s|", pelanggan[i].merk)
			fmt.Printf(" %-22d|", pelanggan[i].angsuran)
			fmt.Printf(" %-18d|", pelanggan[i].lamaKredit)
			fmt.Printf(" %-14d|", pelanggan[i].tahunBeli)
			fmt.Printf(" %-16d|\n", pelanggan[i].sisaAngsuran)
			fmt.Println("")
		}
	}
	if ketemu <= 0 {
		fmt.Println("data tidak ditemukan")
	}
}

func tampillebih2016(data []pembeli) {
	var ketemu int
	
	for i := len(data) ; i > 0 ; i-- {
		for j := 1 ; j<i ; j++ {
			if data[j-1].tahunBeli > data[j].tahunBeli {
				t := data[j]
				data[j] = data[j-1]
				data[j-1] = t
			}
		}
	}

	/*i := 1
	for i < len(data) {
		t := data[i]
		j := i - 1
		for j >= 0 && data[j].tahunBeli > t.tahunBeli {
			data[j + 1] = data[j]
			j = j - 1
		}
		data[j + 1] = t
		i = i + 1
	}
	*/
	
	no := 0
	fmt.Printf("| No |   Nama Pembeli   |    Alamat Pribadi     | Umur Pembeli |  Merk Mobil  | Angsuran Yang Diambil |Lama Kredit (Bulan)|Tahun Pembelian|  Sisa Angsuran  |\n")
	for i:= 0 ; i<len(data) && i < 10 ; i++ {
		if data[i].tahunBeli > 2016{
			ketemu++
			no++
			fmt.Printf("| %-2d |", no)
			fmt.Printf(" %-17s|", pelanggan[i].nama)
			fmt.Printf(" %-22s|", pelanggan[i].alamat)
			fmt.Printf(" %-13d|", pelanggan[i].umur)
			fmt.Printf(" %-13s|", pelanggan[i].merk)
			fmt.Printf(" %-22d|", pelanggan[i].angsuran)
			fmt.Printf(" %-18d|", pelanggan[i].lamaKredit)
			fmt.Printf(" %-14d|", pelanggan[i].tahunBeli)
			fmt.Printf(" %-16d|\n", pelanggan[i].sisaAngsuran)
			fmt.Println("")
		}
	}
	if ketemu <= 0 {
		fmt.Println("data tidak ditemukan")
	}
} 

func main() {
	//identitas struct (kode, merk, tipe, harga, stock)
	mobil = append(mobil,
		identitas{"HR-V", "Honda", "1.8 CVT Prestige", 419500000, 25},
		identitas{"Mazda", "Mazda", "HatchBack", 489800000, 37},
		identitas{"Pajero", "Mitsubishi", "GLX 4x4 MT", 546500000, 27},
		identitas{"Juke", "Nissan", "1.5 SV MT", 224000000, 10},
		identitas{"Baleno", "Suzuki", "GL AT", 230500000, 45},
		identitas{"Xenia", "Daihatsu", "R AT 1.3", 217000000, 26},
		identitas{"Hilux", "Toyota", "2.4 G 4x4", 432050000, 50},
		identitas{"Confero", "Wuiling", "Cortez 1.8C", 237000000, 17},
		identitas{"Ertiga", "Suzuki", "GL MT", 219500000, 21},
		identitas{"Avanza", "Toyota", "1.3 G AT", 219650000, 73},
	)

	//pembeli struct (nama, alamat, umur, merk, angsuran, lamaKredit, tahunBeli, sisaAngsuran)
	pelanggan = append(pelanggan,
		pembeli{"Dini Puspitarini", "Jl.Panda 1 No.20", 18, "Toyota", 37000000, 12, 2019, 395050000},
		pembeli{"Awalluddyn Ridho", "Jl.Telekomunikasi No.1", 18, "Mazda", 41000000, 12, 2019, 448800000},
		pembeli{"Oma Irama", "Jl.Bandung No.1", 19, "Honda", 35000000, 12, 2019, 384500000},
		pembeli{"Rizky Tri", "Jl.Riau No.2", 20, "Suzuki", 9500000, 24, 2017, 210000000},
		pembeli{"Hanvito Michael", "Jl.Padang No.3", 17, "Nissan", 6500000, 36, 2018, 217500000},
		pembeli{"Naufaal Haritsah", "Jl.Sidoarjo No.4", 18, "Toyota", 18500000, 12, 2019, 201150000},
		pembeli{"Kasyfi Autsar", "Jl.Tangerang No.5", 18, "Suzuki", 10000000, 24, 2019, 220500000},
		pembeli{"Vena Erla", "Jl.Klaten No.5", 18, "Wuiling", 9875000, 24, 2018, 227125000},
		pembeli{"Claudia Mei", "Jl.Samosir No.1", 18, "Daihatsu", 18000000, 12, 2019, 199000000},
		pembeli{"Nanda Ihwani", "Jl.Sumbawa No.4", 18, "Mitsubishi", 15180000, 36, 2018, 531320000},
		pembeli{"Fadhilah Nadia", "Jl.Cirebon No.3", 18, "Mazda", 20500000, 24, 2017, 469300000},
		pembeli{"Sania Raihan", "Jl.Lombok No.1", 18, "Honda", 35000000, 12, 2019, 384500000},
		pembeli{"Park Jimin", "Jl.Bighit No.1", 24, "Suzuki", 46000000, 12, 2019, 201200000},
		pembeli{"Rapmons", "Jl.BTS No.1", 26, "Toyota", 37000000, 12, 2019, 395050000},
		pembeli{"V", "Jl.Cinta No.1", 24, "Honda", 35000000, 12, 2019, 384500000},
		pembeli{"Wahtanca", "Jl.Buleleng No.1", 19, "Wuiling", 19750000, 12, 2019, 217250000},
	)
	menu()
}
