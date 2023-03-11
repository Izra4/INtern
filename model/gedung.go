package model

import (
	"InternBCC/database"
	"InternBCC/entity"
	"fmt"
	"gorm.io/gorm"
)

func GDummy() {
	G1 := entity.Gedung{
		Model:     gorm.Model{},
		Nama:      "Graha Cakrawala",
		Alamat:    " Jl. Cakrawala, Sumbersari, Kec. Lowokwaru, Kota Malang, Jawa Timur",
		Kecamatan: "Lowokwaru",
		Fasilitas: "Area parki basement;Area parkir central park(500 mobil);Loket;Toilet;Dapur;Backstage;Ruang transit; Gudang; Tribune; Properti gedung: AC, genset, lighting, smoke, sound system, WIFI, CCTV",
		Aturan:    "Dilarang membawa barang yang memicu kebakaran",
		Harga:     3000000,
		Kapasitas: "8.000 Orang",
		Luas:      "4.356 m2",
		Tag:       "Wisuda;seminar;konser;pernikahan",
		Bookings:  nil,
	}
	G2 := entity.Gedung{
		Model:     gorm.Model{},
		Nama:      "Dome UMM",
		Alamat:    " Jl. Karyawiguna No.90, Babatan, Tegalgondo, Kec. Karang Ploso, Kabupaten Malang, Jawa Timur",
		Kecamatan: "Karang Ploso",
		Fasilitas: "Ruang operator;Toilet;Backstage;Ruang panel listrik;Gudang;Tribune;Ruang teater;Proper gedung:sound system, AC, LCD",
		Harga:     13000000,
		Kapasitas: "6.000 Orang",
		Luas:      "1.709 m2",
		Aturan:    "Tidak menyajikan sexy dancer",
		Tag:       "konser;pernikahan;pameran",
		Bookings:  nil,
	}
	G3 := entity.Gedung{
		Model:     gorm.Model{},
		Nama:      "Grand ballroom ijen suite",
		Alamat:    "Jl. Ijen Nirwana Raya Blok A No.16, Bareng, Kec. Klojen, Kota Malang, Jawa Timur",
		Kecamatan: "Klojen",
		Fasilitas: "Kursi 500 buah; Karpet lantai;AC;Toilet",
		Harga:     12000000,
		Kapasitas: "2.000 Orang",
		Luas:      "1.300 m2",
		Aturan:    "Dilarang Merokok;Dilarang membawa makanan dari luar",
		Tag:       "Pernikahan;Seminar;Gathering",
		Bookings:  nil,
	}
	G4 := entity.Gedung{
		Model:     gorm.Model{},
		Nama:      "Graha skodam V brawijaya",
		Alamat:    "Jl. Tugu No.2, Klojen, Kec. Klojen, Kota Malang, Jawa Timur",
		Kecamatan: "Klojen",
		Harga:     5900000,
		Kapasitas: "1.000 Orang",
		Luas:      "857 m2",
		Fasilitas: "150 buah kursi dan meja;Genzet;Backstage;Toilet;Area parkir;Sound system",
		Aturan:    "Harus menggunakan ijin keramaian",
		Tag:       "Pernikahan;Seminar;Pameran;Wisuda",
		Bookings:  nil,
	}
	G5 := entity.Gedung{
		Model:     gorm.Model{},
		Nama:      "Graha polinema",
		Alamat:    "Jl. Soekarno Hatta No.9, Jatimulyo, Kec. Lowokwaru, Kota Malang, Jawa Timur",
		Kecamatan: "Lowokwaru",
		Harga:     7000000,
		Kapasitas: "4.000 Orang",
		Luas:      "3.694 m2",
		Fasilitas: "Generator set 500 Kva;Kursi 100 buah;Karpet lantai; AC (Standing);Listirk 50.000;Area parkir;Backstage;Toilet;Audio (Sound System)",
		Aturan:    "Dilarang membawa barang yang memicu kebakaran",
		Tag:       "Wisuda;Seminar;Pernikahan",
		Bookings:  nil,
	}
	G6 := entity.Gedung{
		Model:     gorm.Model{},
		Nama:      "Graha Widyagama",
		Alamat:    "Jl. Borobudur No.35, Mojolangu, Kec. Lowokwaru, Kota Malang, Jawa Timur",
		Kecamatan: "Lowokwaru",
		Harga:     4000000,
		Kapasitas: "600 Orang",
		Luas:      "725 m2",
		Fasilitas: "Backstage;Kursi",
		Aturan:    "Dilarang membawa barang yang memicu kebakaran",
		Tag:       "Pernikahan;Seminar",
		Bookings:  nil,
	}

	G7 := entity.Gedung{
		Model:     gorm.Model{},
		Nama:      "Atria grand paramount ballroom",
		Alamat:    "Jl. Letjend S. Parman No.87 - 89, Purwantoro, Kec. Blimbing, Kota Malang, Jawa Timur",
		Kecamatan: "Blimbing",
		Harga:     60000000,
		Kapasitas: "1.300 orang",
		Luas:      "999 m2",
		Tag:       "Gathering;Wisuda;Seminar",
		Fasilitas: "Wifi;Projector;Screen;Kursi dan Meja;AC;Sound system;lighting",
		Aturan:    "Dilarang memakai sandal;Dilarang merokok;Dilarang membawa makanan dari luar",
		Bookings:  nil,
	}

	G8 := entity.Gedung{
		Model:     gorm.Model{},
		Nama:      "Harris conventions ballroom",
		Alamat:    "Jl. A Yani Utara, Jl. Riverside Blk. C No.1, Polowijen, Blimbing, Malang, Jawa timur",
		Kecamatan: "Blimbing",
		Harga:     73000000,
		Kapasitas: "2.250 orang",
		Luas:      "1.728 m2",
		Tag:       "Gathering;Pernikahan",
		Fasilitas: "Wifi;Projector;Screen;Kursi dan meja;AC;Sound system; lighting",
		Aturan:    "Dilarang merusak properti",
		Bookings:  nil,
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
}

func LDummy() {
	L1 := entity.Link{
		Model:    gorm.Model{},
		Link:     "https://ontvftbxgsmzxwlqhsdn.supabase.co/storage/v1/object/sign/gambar-gedung/Gracak.svg?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1cmwiOiJnYW1iYXItZ2VkdW5nL0dyYWNhay5zdmciLCJpYXQiOjE2Nzg0MzI0NzgsImV4cCI6MTY4MTAyNDQ3OH0.FT8mr1n6u5zLDsegHPVz7JXWERQyBZYjSxNzRR2en9M&t=2023-03-10T07%3A14%3A39.608Z",
		GedungID: 1,
	}
	L2 := entity.Link{
		Model:    gorm.Model{},
		Link:     "https://ontvftbxgsmzxwlqhsdn.supabase.co/storage/v1/object/sign/gambar-gedung/UMM.svg?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1cmwiOiJnYW1iYXItZ2VkdW5nL1VNTS5zdmciLCJpYXQiOjE2Nzg0MzU4MDksImV4cCI6MTY4MTAyNzgwOX0.1-wKESXerDMdnyukPyItXR8c4C8KTHrbsue6oZXwM9I&t=2023-03-10T08%3A10%3A09.942Z",
		GedungID: 2,
	}
	L3 := entity.Link{
		Model:    gorm.Model{},
		Link:     "https://ontvftbxgsmzxwlqhsdn.supabase.co/storage/v1/object/sign/gambar-gedung/UMM.svg?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1cmwiOiJnYW1iYXItZ2VkdW5nL1VNTS5zdmciLCJpYXQiOjE2Nzg0MzU4MDksImV4cCI6MTY4MTAyNzgwOX0.1-wKESXerDMdnyukPyItXR8c4C8KTHrbsue6oZXwM9I&t=2023-03-10T08%3A10%3A09.942Z",
		GedungID: 3,
	}
	L4 := entity.Link{
		Model:    gorm.Model{},
		Link:     "https://ontvftbxgsmzxwlqhsdn.supabase.co/storage/v1/object/sign/gambar-gedung/Skodamm.svg?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1cmwiOiJnYW1iYXItZ2VkdW5nL1Nrb2RhbW0uc3ZnIiwiaWF0IjoxNjc4NDM2NzYxLCJleHAiOjE2ODEwMjg3NjF9.VIMwyqLS83KWtp_WsQWRu5VLF1nTwZBOhE_nzpw15iI&t=2023-03-10T08%3A26%3A02.383Z",
		GedungID: 4,
	}
	L5 := entity.Link{
		Model:    gorm.Model{},
		Link:     "https://ontvftbxgsmzxwlqhsdn.supabase.co/storage/v1/object/sign/gambar-gedung/Polinema.svg?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1cmwiOiJnYW1iYXItZ2VkdW5nL1BvbGluZW1hLnN2ZyIsImlhdCI6MTY3ODQzNjc4NCwiZXhwIjoxNjgxMDI4Nzg0fQ.68ovpOQcCyVLk4jpdwqSiAB7JxJnONStKDH8d-aeTZM&t=2023-03-10T08%3A26%3A25.147Z",
		GedungID: 5,
	}
	L6 := entity.Link{
		Model:    gorm.Model{},
		Link:     "https://ontvftbxgsmzxwlqhsdn.supabase.co/storage/v1/object/sign/gambar-gedung/Widyagama%20(2).svg?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1cmwiOiJnYW1iYXItZ2VkdW5nL1dpZHlhZ2FtYSAoMikuc3ZnIiwiaWF0IjoxNjc4NDM2NzM4LCJleHAiOjE2ODEwMjg3Mzh9.89jcwWpQfMalBqyhp--Lxlnsc19ll5k0OweYZt6Uxi0&t=2023-03-10T08%3A25%3A39.032Z",
		GedungID: 6,
	}
	L7 := entity.Link{
		Model:    gorm.Model{},
		Link:     "https://ontvftbxgsmzxwlqhsdn.supabase.co/storage/v1/object/sign/gambar-gedung/Atria%20Ballroom.svg?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1cmwiOiJnYW1iYXItZ2VkdW5nL0F0cmlhIEJhbGxyb29tLnN2ZyIsImlhdCI6MTY3ODQzNjgwNSwiZXhwIjoxNjgxMDI4ODA1fQ._VRWEXpbIcdmpXlrCylWTW3oQD4_QP-TQhsCMhaSdTo&t=2023-03-10T08%3A26%3A46.331Z",
		GedungID: 7,
	}
	L8 := entity.Link{
		Model:    gorm.Model{},
		Link:     "https://ontvftbxgsmzxwlqhsdn.supabase.co/storage/v1/object/sign/gambar-gedung/Harris%20Conve.svg?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1cmwiOiJnYW1iYXItZ2VkdW5nL0hhcnJpcyBDb252ZS5zdmciLCJpYXQiOjE2Nzg0MzY4MDgsImV4cCI6MTY4MTAyODgwOH0.af4O9Sd4QTYSgGu5_PQyXgW5WKV01wS4jywg2UUOzUA&t=2023-03-10T08%3A26%3A49.209Z",
		GedungID: 8,
	}
	if err := database.DB.Create(&L1).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&L2).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&L3).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&L4).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&L5).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&L6).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&L7).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&L8).Error; err != nil {
		fmt.Println(err.Error())
	}

}
