package database

func Migrate() {
	db := GetDB()
	db.Debug().AutoMigrate(&Record{}, &User{})
}
