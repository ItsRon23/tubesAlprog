package main

import "fmt"

const nmax = 5

type datakota struct {
	kota       string
	parawi     [nmax]parawista
	hot        [nmax]hotel
	ratingUser int
}

type parawista struct {
	nama  string
	harga int
	n     int
}

type hotel struct {
	nama  string
	harga int
	n     int
}

type arrkota [nmax]datakota

func tambahData(t *arrkota, n *int) {
	var kota string
	var ratingUser int

	fmt.Scanln(&kota, &ratingUser)
	for *n < nmax && kota != "stop" {

		fmt.Scanln(&kota, &ratingUser)

		t[*n].kota = kota

		t[*n].ratingUser = ratingUser
		if kota == "stop" {
			//fmt.Println("tes")
		} else {

			t[*n].kota = kota

			t[*n].ratingUser = ratingUser

			*n++

		}
	}
}

func updateKota(t *arrkota, n, m int) {

	var inputUser, inputKota, ip string
	var rate int

	fmt.Println("____________________________________________________")
	fmt.Println("|                                                  |")
	fmt.Println("|=========== Data apa yang ingin diubah? ==========|")
	fmt.Println("|                                                  |")
	fmt.Println("| 1. Merubah data Kota                             |")
	fmt.Println("| 2. Merubah data Hotel                            |")
	fmt.Println("| 3. Merubah data Parawisata                       |")
	fmt.Println("|__________________________________________________|")

	fmt.Print("Masukan Angka: ")
	fmt.Scan(&ip)
	fmt.Print("Masukan Nama Kota: ")
	fmt.Scan(&inputKota)
	fmt.Println(" ")

	fmt.Print("Masukkan Nama kota 'baru' dan rating: ")
	for i := 0; i < n; i++ {

		if t[i].kota == inputKota {
			fmt.Scan(&inputUser)
			fmt.Scan(&rate)
			t[i].kota = inputUser
			t[i].ratingUser = rate

		}
		fmt.Println(" ")
		fmt.Println("Data Berhasil diUbah")
		fmt.Println("Kembali ke awal.")
	}
}

func updateHotel(t *arrkota, n int, m, o *int) {
	var nama string
	var harga int
	var inputUser, ip string
	var j int
	j = 0

	fmt.Println("____________________________________________________")
	fmt.Println("|                                                  |")
	fmt.Println("|========= Data apa yang ingin ditambahkan? =======|")
	fmt.Println("|                                                  |")
	fmt.Println("| 1. Menambahkan data Hotel                        |")
	fmt.Println("| 2. Menambahkan data Parawisata                   |")
	fmt.Println("|__________________________________________________|")

	fmt.Print("Masukan Angka: ")
	fmt.Scan(&ip)

	if ip == "1" {
		j = *m
	} else {
		j = *o
	}
	fmt.Print("Masukan Nama Kota: ")
	fmt.Scan(&inputUser)

	for i := 0; i < n; i++ {

		if t[i].kota == inputUser {
			fmt.Println("|============ Update hotel/pariwisata =============|")
			fmt.Println("Nama & Harga: ")
			fmt.Scan(&nama, &harga)

			*m = t[i].hot[i].n
			*o = t[i].parawi[i].n

			for j < nmax && nama != "stop" {

				if ip == "1" {

					t[i].hot[*m].nama = nama
					t[i].hot[*m].harga = harga

					fmt.Scanln(&nama, &harga)

					if nama == "stop" {
						*m = j

					} else {

						j++
						*m = j
					}
				} else if ip == "2" {
					t[i].parawi[*o].nama = nama
					t[i].parawi[*o].harga = harga

					fmt.Scanln(&nama, &harga)
					if nama == "stop" {
						*o = j
					} else {
						j++
						*o = j

					}

				}
			}
		}
	}

	subMenuUtama1(*t, n, *m, *o)
}

func removeKota(t *arrkota, n, m, o int) {
	var inputUser string
	var ip, ip2 string

	fmt.Println("____________________________________________________")
	fmt.Println("|                                                  |")
	fmt.Println("|========= Data apa yang ingin dihapuskan? ========|")
	fmt.Println("|                                                  |")
	fmt.Println("| 1. Menghapus data Kota                           |")
	fmt.Println("| 2. Menghapus data Hotel                          |")
	fmt.Println("| 3. Menghapus data Parawisata                     |")
	fmt.Println("|__________________________________________________|")

	fmt.Print("Masukan Angka: ")
	fmt.Scan(&ip)
	fmt.Print("Data Kota yang ingin dihapus: ")
	fmt.Scan(&inputUser)

	for i := 0; i < n; i++ {
		if t[i].kota == inputUser {

			if ip == "1" {
				t[i] = t[i+1]
				for j := i; j < nmax; j++ {
					t[i+1] = t[j]
				}

				n = n - 1
			} else if ip == "2" {
				fmt.Println("Hotel apa yang ingin di hapus: ")
				fmt.Scan(&ip2)
				for j := 0; j < m; j++ {
					if t[i].hot[j].nama == ip2 {
						t[i].hot[j] = t[i].hot[j+1]
						for k := j; k < nmax; k++ {

							t[i].hot[j+1] = t[i].hot[k]
						}

					}

				}
				m = m - 1
			} else if ip == "3" {
				fmt.Println("Pariwisata apa yang ingin di hapus: ")
				fmt.Scan(&ip2)
				for j := 0; j < m; j++ {
					if t[i].parawi[j].nama == ip2 {
						t[i].parawi[j] = t[i].parawi[j+1]
						for k := j; k < nmax; k++ {

							t[i].parawi[j+1] = t[i].parawi[k]
						}

					}

				}
				o = o - 1
			}

		}
	}
	fmt.Println(" ")
	fmt.Println("Data Berhasil diHapus")
	fmt.Println("Kembali ke awal.")
	subMenuUtama1(*t, n, m, o)
}

func show(t *arrkota, n, m, o *int) {

	insertionSort(t, n)
	fmt.Println(" ")
	fmt.Println("Seluruh data yang sudah dimiliki: ")

	for i := 0; i < *n; i++ {

		fmt.Println(t[i].kota, t[i].ratingUser)
	}
	fmt.Println(" ")
	subMenuUtama2(t, *n, *m, *o)
}

func showhotel(t *arrkota, n, m, o int) {
	var inputUser, ip string
	var min, max int
	fmt.Println("____________________________________________________")
	fmt.Println("|                                                  |")
	fmt.Println("|========= Data apa yang ingin Ditampilkan? =======|")
	fmt.Println("|                                                  |")
	fmt.Println("| 1. Melihat data Hotel                            |")
	fmt.Println("| 2. Melihat data Parawisata                       |")
	fmt.Println("| 3.filter data hotel sesuai min max               |")
	fmt.Println("| 3.filter data parawsiata sesuai min max          |")
	fmt.Println("|__________________________________________________|")
	fmt.Println("|__________________________________________________|")

	fmt.Print("Masukan Angka: ")
	fmt.Scan(&ip)
	fmt.Print("Masukan nama Kota: ")
	fmt.Scan(&inputUser)

	for i := 0; i < n; i++ {
		if t[i].kota == inputUser {
			if ip == "1" {
				for j := 0; j < m+1; j++ {

					fmt.Println(t[i].hot[j].nama, " --- ", t[i].hot[j].harga)

				}

			} else if ip == "2" {
				for j := 1; j < o+1; j++ {

					fmt.Println(t[i].parawi[j].nama, " --- ", t[i].parawi[j].harga)

				}

			} else if ip == "3" {
				fmt.Println("masukan harga minum dan maksimum")
				fmt.Scan(&min, &max)

				for j := 0; j < m+1; j++ {
					if t[i].hot[j].harga > min && t[i].hot[j].harga < max {

						fmt.Println(t[i].hot[j].nama, " --- ", t[i].hot[j].harga)
					}

				}
			} else if ip == "4" {
				fmt.Println("masukan harga minum dan maksimum")
				fmt.Scan(&min, &max)

				for j := 0; j < o+1; j++ {
					if t[i].parawi[j].harga > min && t[i].parawi[j].harga < max {

						fmt.Println(t[i].parawi[j].nama, " --- ", t[i].parawi[j].harga)
					}

				}
			}

		}

		subMenuUtama2(t, n, m, o)
	}
}

func insertionSort(t *arrkota, n *int) {
	for i := 1; i < *n; i++ {
		key := t[i]
		j := i - 1
		for j >= 0 && t[j].ratingUser < key.ratingUser {
			t[j+1] = t[j]
			j--
		}
		t[j+1] = key
	}
}

func binarrating(t *arrkota, n, target int) int {
	selectionsortfix(t, &n)
	fmt.Scanln(&target)
	low := 0
	high := n - 1

	mid := (low + high) / 2
	for low <= high {

		if t[mid].ratingUser == target {
			return mid

		} else if t[mid].ratingUser < target {
			low = mid + 1

		} else {
			high = mid - 1
		}
	}

	return -1
}

func selectionsortfix(t *arrkota, n *int) {

	var minindex int
	for i := 0; i < *n; i++ {
		minindex = i
		for j := i + 1; j < *n; j++ {
			if t[j].ratingUser < t[minindex].ratingUser {
				minindex = j
			}
			t[i], t[minindex] = t[minindex], t[i]
		}
	}
}

func menuUtama(tab arrkota, n, m, o int) {
	var in0 string
	var target int

	fmt.Println("____________________________________________________")
	fmt.Println("|                                                  |")
	fmt.Println("| --------------- NAVIGASI FITUR ----------------- |")
	fmt.Println("|=== Selamat Datang di Aplikasi Data Pariwisata ===|")
	fmt.Println("| Aaron Joseph : 1302223083                        |")
	fmt.Println("| Fajar Mufid  : 1302223032                        |")
	fmt.Println("|                                                  |")
	fmt.Println("| Silahkan pilih:                                  |")
	fmt.Println("| 1. Menambah/Merubah/Menghapus Data               |")
	fmt.Println("| 2. Menampilkan data yang ada                     |")
	fmt.Println("| 3. Mencari data                                  |")
	fmt.Println("| 4. keluar                                        |")
	fmt.Println("|__________________________________________________|")

	fmt.Print("Masukan Angka: ")
	fmt.Scan(&in0)

	if in0 == "1" {
		subMenuUtama1(tab, n, m, o)
	} else if in0 == "2" {
		subMenuUtama2(&tab, n, m, o)

	} else if in0 == "3" {
		fmt.Print("Rating berapa yang ingin dicari: ")
		fmt.Scan(&target)
		fmt.Print("Data dengan rating tersebut adalah: ")
		fmt.Print(tab[binarrating(&tab, n, target)].kota)

		menuUtama(tab, n, m, o)
	} else if in0 == "4" {
		fmt.Println("Terima Kasih sudah Menggunakan perangkat Lunak kami! <3")
	}
}

func subMenuUtama1(tab arrkota, n, m, o int) {
	var in1 string

	fmt.Println("____________________________________________________")
	fmt.Println("|                                                  |")
	fmt.Println("|== Anda ingin Menambah, Merubah atau Menghapus? ==|")
	fmt.Println("|                                                  |")
	fmt.Println("| 1. Menambah Data                                 |")
	fmt.Println("| 2. Merubah Data                                  |")
	fmt.Println("| 3. Menghapus Data                                |")
	fmt.Println("| 4. Kembali                                       |")
	fmt.Println("|__________________________________________________|")

	fmt.Print("Masukan Angka: ")
	fmt.Scan(&in1)

	if in1 == "1" {
		menutambahData(tab, n, m, o)
	} else if in1 == "2" {
		menuUbahData(&tab, n, m, o)
	} else if in1 == "3" {
		menuHapusData(&tab, n, m, o)
	} else if in1 == "4" {
		menuUtama(tab, n, m, o)
	}
}

func menutambahData(tab arrkota, n, m, o int) {
	var in2 string
	fmt.Println("____________________________________________________")
	fmt.Println("|                                                  |")
	fmt.Println("|========= Input data yang ingin ditambah =========|")
	fmt.Println("|                                                  |")
	fmt.Println("| 1. Data Kota                                     |")
	fmt.Println("| 2. Data Hotel/tempat parawisata                  |")
	fmt.Println("| 3. Kembali                                       |")
	fmt.Println("|__________________________________________________|")

	fmt.Print("Masukan Angka: ")
	fmt.Scan(&in2)

	if in2 == "1" {
		fmt.Println(" ")
		fmt.Println("|=========== Silahkan masukan data Kota ===========|")
		if n == nmax {
			fmt.Println("memori penuh")
		} else {
			tambahData(&tab, &n)
		}
		fmt.Println(" ")

		fmt.Println("|====== Terima Kasih sudah Mengisi data  ======|")
		subMenuUtama1(tab, n, m, o)

	} else if in2 == "2" {
		fmt.Println(" ")
		fmt.Println("|=========== Silahkan masukan data Hotel ==========|")
		updateHotel(&tab, n, &m, &o)
		fmt.Println(" ")

		fmt.Println("|= Terima Kasih sudah Mengisi data Hotel/Pariwisata =|")
		subMenuUtama1(tab, n, m, o)
	} else if in2 == "3" {
		subMenuUtama1(tab, n, m, o)
	}
}

func menuUbahData(tab *arrkota, n, m, o int) {

	updateKota(tab, n, m)
	subMenuUtama1(*tab, n, m, o)

}

func menuHapusData(tab *arrkota, n, m, o int) {

	removeKota(tab, n, m, o)

}

func subMenuUtama2(tab *arrkota, n, m, o int) {
	var in3 string

	fmt.Println("____________________________________________________")
	fmt.Println("|                                                  |")
	fmt.Println("|======== Data apa yang ingin di tampilkan? =======|")
	fmt.Println("|                                                  |")
	fmt.Println("| 1. Menampilkan seluruh data Kota                 |")
	fmt.Println("| 2. Menampilkan Seluruh data Hotel/parawisata     |")
	fmt.Println("| 3. Kembali                                       |")
	fmt.Println("|__________________________________________________|")

	fmt.Print("Masukan Angka: ")
	fmt.Scan(&in3)

	if in3 == "1" {
		show(tab, &n, &m, &o)
	} else if in3 == "2" {
		showhotel(tab, n, m, o)
	} else if in3 == "3" {
		menuUtama(*tab, n, m, o)
	}
}

func main() {
	var tab arrkota
	var n, m, o int

	menuUtama(tab, n, m, o)
}
