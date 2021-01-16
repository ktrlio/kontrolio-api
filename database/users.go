package database

func GetUser(email string) (*User, error) {
	db := GetDB()

	var user User

	result := db.Where("email = ?", email).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
