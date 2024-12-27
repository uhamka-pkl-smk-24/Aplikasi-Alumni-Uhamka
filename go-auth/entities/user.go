package entities

type User struct {
	Nim      int64
	Username string
	Email    string
	Password string
	Role     string
	No_telp  string
}

type UserImport struct {
	Nim      int64
	Username string
	Email    string
	Password string
	Role     string
	No_telp  string
}
