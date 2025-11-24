package dto

type UserDTO struct {
	ID uint `json:"id"`
	Name string `json:"name"`
}

type UserRequest struct {
	NoInduk 			string 	`form:"no_induk"`
	Name     			string 	`form:"name"`
	Email    			string 	`form:"email"`
	Role				string 	`form:"role"`
	PhotoUrl 			string 	`form:"photo_url"`
	Phone 				string 	`form:"phone"`
	Alamat 				string 	`form:"alamat"`
	Jabatan				string 	`form:"jabatan"`
	TahunAjaranMulai	string 	`form:"tahun_ajaran_mulai"`
}

type UserResponse struct {
	ID 					uint	`json:"id"`
	Name     			string 	`json:"name"`
	Email    			string 	`json:"email"`
	Role				string 	`json:"role"`
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}