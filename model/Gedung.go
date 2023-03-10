package model

import (
	"InternBCC/database"
	"InternBCC/entity"
	"fmt"
	"gorm.io/gorm"
)

type name struct {
}

func GDummy() {
	G1 := entity.Gedung{
		Model:     gorm.Model{},
		Nama:      "Graha Cakrawala",
		Alamat:    " Jl. Cakrawala, Sumbersari, Kec. Lowokwaru, Kota Malang, Jawa Timur",
		Kecamatan: "Lowokwaru",
		Fasilitas: "Area parki basement;Area parkir central park(500 mobil);Loket;Toilet;Dapur;Backstage;Ruang transit; Gudang; Tribune; Properti gedung: AC, genset, lighting, smoke, sound system, WIFI, CCTV",
		Aturan:    "Dilarang membawa barang yang memicu kebakaran",
		Harga:     "Rp. 30Jt",
		Kapasitas: "8.000 Orang",
		Luas:      "4.356 m2",
		Tag:       "Wisuda;seminar;konser;pernikahan",
		Booking:   nil,
	}
	G2 := entity.Gedung{
		Model:     gorm.Model{},
		Nama:      "Dome UMM",
		Alamat:    " Jl. Karyawiguna No.90, Babatan, Tegalgondo, Kec. Karang Ploso, Kabupaten Malang, Jawa Timur",
		Kecamatan: "Karang Ploso",
		Fasilitas: "Ruang operator;Toilet;Backstage;Ruang panel listrik;Gudang;Tribune;Ruang teater;Proper gedung:sound system, AC, LCD",
		Harga:     "Rp. 13jt",
		Kapasitas: "6.000 Orang",
		Luas:      "1.709 m2",
		Aturan:    "Tidak menyajikan sexy dancer",
		Tag:       "konser;pernikahan;pameran",
		Booking:   nil,
	}
	G3 := entity.Gedung{
		Model:     gorm.Model{},
		Nama:      "Grand ballroom ijen suite",
		Alamat:    "Jl. Ijen Nirwana Raya Blok A No.16, Bareng, Kec. Klojen, Kota Malang, Jawa Timur",
		Kecamatan: "Klojen",
		Fasilitas: "Kursi 500 buah; Karpet lantai;AC;Toilet",
		Harga:     "Rp. 12jt",
		Kapasitas: "2.000 Orang",
		Luas:      "1.300 m2",
		Aturan:    "Dilarang Merokok;Dilarang membawa makanan dari luar",
		Tag:       "Pernikahan;Seminar;Gathering",
		Booking:   nil,
	}
	G4 := entity.Gedung{
		Model:     gorm.Model{},
		Nama:      "Graha skodam V brawijaya",
		Alamat:    "Jl. Tugu No.2, Klojen, Kec. Klojen, Kota Malang, Jawa Timur",
		Kecamatan: "Klojen",
		Harga:     "Rp. 5.9jt",
		Kapasitas: "1.000 Orang",
		Luas:      "857 m2",
		Fasilitas: "150 buah kursi dan meja;Genzet;Backstage;Toilet;Area parkir;Sound system",
		Aturan:    "Harus menggunakan ijin keramaian",
		Tag:       "Pernikahan;Seminar;Pameran;Wisuda",
		Booking:   nil,
	}
	G5 := entity.Gedung{
		Model:     gorm.Model{},
		Nama:      "Graha polinema",
		Alamat:    "Jl. Soekarno Hatta No.9, Jatimulyo, Kec. Lowokwaru, Kota Malang, Jawa Timur",
		Kecamatan: "Lowokwaru",
		Harga:     "Rp. 7jt",
		Kapasitas: "4.000 Orang",
		Luas:      "3.694 m2",
		Fasilitas: "Generator set 500 Kva;Kursi 100 buah;Karpet lantai; AC (Standing);Listirk 50.000;Area parkir;Backstage;Toilet;Audio (Sound System)",
		Aturan:    "Dilarang membawa barang yang memicu kebakaran",
		Tag:       "Wisuda;Seminar;Pernikahan",
		Booking:   nil,
	}
	G6 := entity.Gedung{
		Model:     gorm.Model{},
		Nama:      "Graha Widyagama",
		Alamat:    "Jl. Borobudur No.35, Mojolangu, Kec. Lowokwaru, Kota Malang, Jawa Timur",
		Kecamatan: "Lowokwaru",
		Harga:     "Rp. 4jt",
		Kapasitas: "600 Orang",
		Luas:      "725 m2",
		Fasilitas: "Backstage;Kursi",
		Aturan:    "Dilarang membawa barang yang memicu kebakaran",
		Tag:       "Pernikahan;Seminar",
		Booking:   nil,
	}

	G7 := entity.Gedung{
		Model:     gorm.Model{},
		Nama:      "Atria grand paramount ballroom",
		Alamat:    "Jl. Letjend S. Parman No.87 - 89, Purwantoro, Kec. Blimbing, Kota Malang, Jawa Timur",
		Kecamatan: "Blimbing",
		Harga:     "Rp. 60jt",
		Kapasitas: "1.300 orang",
		Luas:      "999 m2",
		Tag:       "Gathering;Wisuda;Seminar",
		Fasilitas: "Wifi;Projector;Screen;Kursi dan Meja;AC;Sound system;lighting",
		Aturan:    "Dilarang memakai sandal;Dilarang merokok;Dilarang membawa makanan dari luar",
		Booking:   nil,
	}

	G8 := entity.Gedung{
		Model:     gorm.Model{},
		Nama:      "Harris conventions ballroom",
		Alamat:    "Jl. A Yani Utara, Jl. Riverside Blk. C No.1, Polowijen, Blimbing, Malang, Jawa timur",
		Kecamatan: "Blimbing",
		Harga:     "Rp. 73jt",
		Kapasitas: "2.250 orang",
		Luas:      "1.728 m2",
		Tag:       "Gathering;Pernikahan",
		Fasilitas: "Wifi;Projector;Screen;Kursi dan meja;AC;Sound system; lighting",
		Aturan:    "Dilarang merusak properti",
		Booking:   nil,
	}

	if err := database.DB.Create(&G1).Error; err != nil {

		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&G2).Error; err != nil {
		fmt.Println(err.Error())

	}
	if err := database.DB.Create(&G3).Error; err != nil {
		fmt.Println(err.Error())

	}
	if err := database.DB.Create(&G4).Error; err != nil {
		fmt.Println(err.Error())

	}
	if err := database.DB.Create(&G5).Error; err != nil {
		fmt.Println(err.Error())

	}
	if err := database.DB.Create(&G6).Error; err != nil {
		fmt.Println(err.Error())

	}
	if err := database.DB.Create(&G7).Error; err != nil {
		fmt.Println(err.Error())

	}
	if err := database.DB.Create(&G8).Error; err != nil {
		fmt.Println(err.Error())
	}
	//if err := database.DB.Model(&G1).Association("Tag").Append([]Tag).Error(); err != nil {
	//	return
	//}
}
