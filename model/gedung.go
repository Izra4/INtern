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
		Link:     "https://ontvftbxgsmzxwlqhsdn.supabase.co/storage/v1/object/sign/gambar-gedung/gracak%201.svg?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1cmwiOiJnYW1iYXItZ2VkdW5nL2dyYWNhayAxLnN2ZyIsImlhdCI6MTY3ODU5OTI5MCwiZXhwIjoxNjgxMTkxMjkwfQ.RHPZu51yEnEby5EcDXsXOHJ_-3Z_xNvtBKiOuY8WQOo&t=2023-03-12T05%3A34%3A52.438Z",
		GedungID: 1,
	}
	L2 := entity.Link{
		Model:    gorm.Model{},
		Link:     "https://ontvftbxgsmzxwlqhsdn.supabase.co/storage/v1/object/sign/gambar-gedung/gracak%202.svg?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1cmwiOiJnYW1iYXItZ2VkdW5nL2dyYWNhayAyLnN2ZyIsImlhdCI6MTY3ODU5OTM1MSwiZXhwIjoxNjgxMTkxMzUxfQ.W-DxllfRdui7pt-9nflhMv-832ReN36HqxKPIq9CUM0&t=2023-03-12T05%3A35%3A52.912Z",
		GedungID: 1,
	}
	L3 := entity.Link{
		Model:    gorm.Model{},
		Link:     "https://ontvftbxgsmzxwlqhsdn.supabase.co/storage/v1/object/sign/gambar-gedung/gracak%203.svg?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1cmwiOiJnYW1iYXItZ2VkdW5nL2dyYWNhayAzLnN2ZyIsImlhdCI6MTY3ODU5OTM3MywiZXhwIjoxNjgxMTkxMzczfQ.ByNbQlrmtVD49JQ5CmgBk8YidgaGC0EKnwlkZPTYRhE&t=2023-03-12T05%3A36%3A15.461Z",
		GedungID: 1,
	}
	L4 := entity.Link{
		Model:    gorm.Model{},
		Link:     "https://ontvftbxgsmzxwlqhsdn.supabase.co/storage/v1/object/public/gambar-gedung/widya%203.svg",
		GedungID: 2,
	}
	L5 := entity.Link{
		Model:    gorm.Model{},
		Link:     "https://ontvftbxgsmzxwlqhsdn.supabase.co/storage/v1/object/public/gambar-gedung/umm%202.svg?t=2023-03-14T13%3A06%3A49.918Z",
		GedungID: 2,
	}
	L6 := entity.Link{
		Model:    gorm.Model{},
		Link:     "https://ontvftbxgsmzxwlqhsdn.supabase.co/storage/v1/object/public/gambar-gedung/umm%203.svg?t=2023-03-14T13%3A06%3A56.102Z",
		GedungID: 2,
	}
	L7 := entity.Link{
		Model:    gorm.Model{},
		Link:     "https://ontvftbxgsmzxwlqhsdn.supabase.co/storage/v1/object/sign/gambar-gedung/Ijen%201.svg?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1cmwiOiJnYW1iYXItZ2VkdW5nL0lqZW4gMS5zdmciLCJpYXQiOjE2Nzg1OTk1MDgsImV4cCI6MTY4MTE5MTUwOH0.y9ZXkL3AUo5bwb5r6Mr7RE-UCXdnTYpvVrRXXVk2R1Q&t=2023-03-12T05%3A38%3A30.616Z",
		GedungID: 3,
	}
	L8 := entity.Link{
		Model:    gorm.Model{},
		Link:     "https://ontvftbxgsmzxwlqhsdn.supabase.co/storage/v1/object/sign/gambar-gedung/Ijen%202.svg?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1cmwiOiJnYW1iYXItZ2VkdW5nL0lqZW4gMi5zdmciLCJpYXQiOjE2Nzg1OTk1MzMsImV4cCI6MTY4MTE5MTUzM30.vfVOTkyxAcqMXnbGiJFjCcsXCEiwbqOXTzwx_W20ONc&t=2023-03-12T05%3A38%3A55.781Z",
		GedungID: 3,
	}
	L9 := entity.Link{
		Model:    gorm.Model{},
		Link:     "https://ontvftbxgsmzxwlqhsdn.supabase.co/storage/v1/object/sign/gambar-gedung/Ijen%203.svg?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1cmwiOiJnYW1iYXItZ2VkdW5nL0lqZW4gMy5zdmciLCJpYXQiOjE2Nzg1OTk1NTUsImV4cCI6MTY4MTE5MTU1NX0.ph_fvY_a4GwTOoSD72Phm0WRY0Rpfjk-WLSA_BsdODQ&t=2023-03-12T05%3A39%3A17.754Z",
		GedungID: 3,
	}
	L10 := entity.Link{
		Model:    gorm.Model{},
		Link:     "https://ontvftbxgsmzxwlqhsdn.supabase.co/storage/v1/object/sign/gambar-gedung/skodam%202.svg?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1cmwiOiJnYW1iYXItZ2VkdW5nL3Nrb2RhbSAyLnN2ZyIsImlhdCI6MTY3ODU5OTU4NywiZXhwIjoxNjgxMTkxNTg3fQ.2GgtgXrh2Czlxag1ZjAyqrIOQEVpxEd1q3ykVit3El4&t=2023-03-12T05%3A39%3A49.521Z",
		GedungID: 4,
	}
	L11 := entity.Link{
		Model:    gorm.Model{},
		Link:     "https://ontvftbxgsmzxwlqhsdn.supabase.co/storage/v1/object/sign/gambar-gedung/skodam%203.svg?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1cmwiOiJnYW1iYXItZ2VkdW5nL3Nrb2RhbSAzLnN2ZyIsImlhdCI6MTY3ODU5OTYwNCwiZXhwIjoxNjgxMTkxNjA0fQ.zOaGkJnd7F7pylQJrX5fU1OgodB8jbH_Tyhn5s0kZkc&t=2023-03-12T05%3A40%3A06.683Z",
		GedungID: 4,
	}
	L12 := entity.Link{
		Model:    gorm.Model{},
		Link:     "https://ontvftbxgsmzxwlqhsdn.supabase.co/storage/v1/object/sign/gambar-gedung/skodam%201.svg?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1cmwiOiJnYW1iYXItZ2VkdW5nL3Nrb2RhbSAxLnN2ZyIsImlhdCI6MTY3ODU5OTYxNCwiZXhwIjoxNjgxMTkxNjE0fQ.xirOeJAjQ1BszY5YIsp4urFfnbvX2p3iAPKohiTE0qA&t=2023-03-12T05%3A40%3A16.797Z",
		GedungID: 4,
	}
	L13 := entity.Link{
		Model:    gorm.Model{},
		Link:     "https://ontvftbxgsmzxwlqhsdn.supabase.co/storage/v1/object/sign/gambar-gedung/Poli%201.svg?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1cmwiOiJnYW1iYXItZ2VkdW5nL1BvbGkgMS5zdmciLCJpYXQiOjE2Nzg1OTk2NDEsImV4cCI6MTY4MTE5MTY0MX0._ZJ3QTW6Oy82ql6SBVphhyvo8bUrVn2pGTJ1rSQI3bk&t=2023-03-12T05%3A40%3A43.872Z",
		GedungID: 5,
	}
	L14 := entity.Link{
		Model:    gorm.Model{},
		Link:     "https://ontvftbxgsmzxwlqhsdn.supabase.co/storage/v1/object/sign/gambar-gedung/Poli%201.svg?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1cmwiOiJnYW1iYXItZ2VkdW5nL1BvbGkgMS5zdmciLCJpYXQiOjE2Nzg1OTk2NDEsImV4cCI6MTY4MTE5MTY0MX0._ZJ3QTW6Oy82ql6SBVphhyvo8bUrVn2pGTJ1rSQI3bk&t=2023-03-12T05%3A40%3A43.872Z",
		GedungID: 5,
	}
	L15 := entity.Link{
		Model:    gorm.Model{},
		Link:     "https://ontvftbxgsmzxwlqhsdn.supabase.co/storage/v1/object/sign/gambar-gedung/Poli%203.svg?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1cmwiOiJnYW1iYXItZ2VkdW5nL1BvbGkgMy5zdmciLCJpYXQiOjE2Nzg1OTk2NjQsImV4cCI6MTY4MTE5MTY2NH0.tAr4Tnw833B2FXiOp5R1E0JToVkV99nHINIR_DACsNE&t=2023-03-12T05%3A41%3A06.076Z",
		GedungID: 5,
	}
	L16 := entity.Link{
		Model:    gorm.Model{},
		Link:     "https://ontvftbxgsmzxwlqhsdn.supabase.co/storage/v1/object/public/gambar-gedung/widya%201.svg",
		GedungID: 6,
	}
	L17 := entity.Link{
		Model:    gorm.Model{},
		Link:     "https://ontvftbxgsmzxwlqhsdn.supabase.co/storage/v1/object/public/gambar-gedung/widya%202.svg",
		GedungID: 6,
	}
	L18 := entity.Link{
		Model:    gorm.Model{},
		Link:     "https://ontvftbxgsmzxwlqhsdn.supabase.co/storage/v1/object/public/gambar-gedung/widya%203.svg",
		GedungID: 6,
	}
	L19 := entity.Link{
		Model:    gorm.Model{},
		Link:     "https://ontvftbxgsmzxwlqhsdn.supabase.co/storage/v1/object/public/gambar-gedung/Atria%201.svg?t=2023-03-14T13%3A07%3A30.013Z",
		GedungID: 7,
	}
	L20 := entity.Link{
		Model:    gorm.Model{},
		Link:     "https://ontvftbxgsmzxwlqhsdn.supabase.co/storage/v1/object/public/gambar-gedung/Atria%202.svg?t=2023-03-14T13%3A07%3A53.424Z",
		GedungID: 7,
	}
	L21 := entity.Link{
		Model:    gorm.Model{},
		Link:     "https://ontvftbxgsmzxwlqhsdn.supabase.co/storage/v1/object/public/gambar-gedung/Atria%203.svg",
		GedungID: 7,
	}
	L22 := entity.Link{
		Model:    gorm.Model{},
		Link:     "https://ontvftbxgsmzxwlqhsdn.supabase.co/storage/v1/object/sign/gambar-gedung/Harris%201.svg?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1cmwiOiJnYW1iYXItZ2VkdW5nL0hhcnJpcyAxLnN2ZyIsImlhdCI6MTY3ODU5OTc0NSwiZXhwIjoxNjgxMTkxNzQ1fQ.sK1c6FeeXV4v-NSXzlBp-JGKQjhSTPph0HIxuU6cs5w&t=2023-03-12T05%3A42%3A27.730Z",
		GedungID: 8,
	}
	L23 := entity.Link{
		Model:    gorm.Model{},
		Link:     "https://ontvftbxgsmzxwlqhsdn.supabase.co/storage/v1/object/sign/gambar-gedung/Harris%202.svg?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1cmwiOiJnYW1iYXItZ2VkdW5nL0hhcnJpcyAyLnN2ZyIsImlhdCI6MTY3ODU5OTc1MCwiZXhwIjoxNjgxMTkxNzUwfQ._TS0X_5GbFbeQ2NaGWb9IPnGyoIaYerdGNPC7xZ-2Jc&t=2023-03-12T05%3A42%3A32.013Z",
		GedungID: 8,
	}
	L24 := entity.Link{
		Model:    gorm.Model{},
		Link:     "https://ontvftbxgsmzxwlqhsdn.supabase.co/storage/v1/object/sign/gambar-gedung/Harris%203.svg?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1cmwiOiJnYW1iYXItZ2VkdW5nL0hhcnJpcyAzLnN2ZyIsImlhdCI6MTY3ODU5OTc1MywiZXhwIjoxNjgxMTkxNzUzfQ.nPOYaf1H1HHLROLF5gpjzYWnbAodp9N59opi6jftcCk&t=2023-03-12T05%3A42%3A35.009Z",
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
	if err := database.DB.Create(&L9).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&L10).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&L11).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&L12).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&L13).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&L14).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&L15).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&L16).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&L17).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&L18).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&L19).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&L20).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&L21).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&L22).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&L23).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&L24).Error; err != nil {
		fmt.Println(err.Error())
	}

}
