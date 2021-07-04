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
	// Get generic database object sql.DB to use its functions
	sqlDB := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(30 * 60)
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

	err = db.Debug().AutoMigrate(&entities.Vocabulary{}, &entities.AwlGroup{}, &entities.TelegramUsers{}, &entities.TelegramActiveCode{}, entities.TelegramUsersCommand{}, entities.TelegramStudyCommand{}, entities.StudyVocabLists{}, entities.AnswerKey{}).Error
	if err != nil {
		log.Fatal(err)
	}

	// Add | Update ForeignKey
	// addForeignKeyFree(db, "vocabulary", addForeignKeyStruct{SourceKey: "awl_group_id", ForeignKey: "awl_group(id)"})
	return true
}
