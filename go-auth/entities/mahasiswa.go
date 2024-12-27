package entities

type Mahasiswa struct {
	NIM          int64
	NoIjazah     string
	NamaLengkap  string
	TempatLahir  string
	TanggalLahir string
	Agama        string
	Alamat       string
	NoTelp       string
	Email        string
	IPK          float64
	BidangStudi  string
	Photo        string
	Angkatan     string
	Tahunlulus   string
}

type MahasiswaImport struct {
	NIM          int64
	NoIjazah     int
	NamaLengkap  string
	TempatLahir  string
	TanggalLahir string
	Agama        string
	Alamat       string
	NoTelp       string
	Email        string
	IPK          float64
	BidangStudi  string
	Photo        string
	Angkatan     int64
	Tahunlulus   int64
}
