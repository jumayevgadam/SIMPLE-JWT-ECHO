package user

// UserData is dao model)
type UserRes struct {
	Username string `db:"username"`
	Email    string `db:"email"`
	Password string `db:"password_hash"`
}

// UserReq is dto model
type UserReq struct {
	Username string `form:"userName" json:"userName"`
	Email    string `form:"userEmail" json:"userEmail"`
	Password string `form:"userPassword" json:"userPassword"`
}

// ToStorage() is
func (u *UserReq) ToStorage() *UserRes {
	return &UserRes{
		Username: u.Username,
		Email:    u.Email,
		Password: u.Email,
	}
}

// ToServer() is
func (u *UserRes) ToServer() *UserReq {
	return &UserReq{
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
	}
}

// AllUserData is
type AllUserData struct {
	ID       int    `db:"id"`
	Username string `db:"username"`
	Email    string `db:"email"`
	Password string `db:"password_hash"`
}

// AllUserDTO is
type AllUserDTO struct {
	ID       int    `json:"userID"`
	Username string `json:"userName"`
	Email    string `json:"email"`
	Password string `json:"passwordHash"`
}

// ToStorage is
func (a *AllUserDTO) ToStorage() *AllUserData {
	return &AllUserData{
		ID:       a.ID,
		Username: a.Username,
		Email:    a.Email,
		Password: a.Password,
	}
}

// ToServer is
func (a *AllUserData) ToServer() *AllUserDTO {
	return &AllUserDTO{
		ID:       a.ID,
		Username: a.Username,
		Email:    a.Email,
		Password: a.Password,
	}
}
