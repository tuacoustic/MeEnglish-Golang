package database

import (
	"log"
	"me-english/entities"
	"me-english/utils/config"
	"me-english/utils/console"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Kết nối cơ sở dữ liệu
func MysqlConnect() (*gorm.DB, error) {
	db, err := gorm.Open(config.MYSQL_DB_DRIVER, config.MYSQL_DB_URL)
	// db.DB().SetMaxIdleConns(10)
	// db.DB().SetMaxOpenConns(7)
	// db.DB().SetConnMaxLifetime(time.Hour)
	db.SingularTable(true)
	if err != nil {
		console.Error("Lỗi kết nối Database, err: ", err)
		return nil, err
	}
	return db, nil
}

type addForeignKeyStruct struct {
	SourceKey  string
	ForeignKey string
}

func addForeignKey(db *gorm.DB, tableName string, d addForeignKeyStruct) {
	var err error
	err = db.Debug().Table(tableName).AddForeignKey(d.SourceKey, d.ForeignKey, "cascade", "cascade").Error
	if err != nil {
		log.Fatal(err)
	}
	return
}

func Auto() bool {
	db, err := MysqlConnect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Debug().AutoMigrate(&entities.Vocabulary{}, &entities.AwlGroup{}, &entities.TelegramUsers{}, &entities.TelegramActiveCode{}).Error
	if err != nil {
		log.Fatal(err)
	}

	// Add | Update ForeignKey
	// addForeignKeyFree(db, "vocabulary", addForeignKeyStruct{SourceKey: "awl_group_id", ForeignKey: "awl_group(id)"})
	return true
}
