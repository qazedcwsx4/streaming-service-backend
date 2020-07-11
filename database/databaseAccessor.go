package database

import (
	"anime-server-go/model"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"log"
	"os"
)

type dbCredentials struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Ip       string `json:"ip"`
	Port     string `json:"port"`
	Database string `json:"database"`
}

func getCredentials() dbCredentials {
	var credentials dbCredentials

	file, err := os.Open("cred.json")
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &credentials)
	if err != nil {
		log.Fatal(err)
	}

	return credentials
}

func obtainDb() *sqlx.DB {
	credentials := getCredentials()

	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		credentials.User,
		credentials.Password,
		credentials.Ip,
		credentials.Port,
		credentials.Database,
	))
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Obtained database")
	return db
}

var db = obtainDb()

func GetAllEntries() []model.Anime {
	// language=SQL
	query := `SELECT name, description, image, mal_reference FROM entry JOIN series JOIN anime;`

	rows, err := db.Queryx(query)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var entries []model.Anime
	for rows.Next() {
		var entry model.Anime
		err := rows.StructScan(&entry)
		if err != nil {
			log.Fatal(err)
		}
		entries = append(entries, entry)
	}
	return entries
}

func InsertAnime() error {
	tx, err := db.Beginx()
	if err != nil {
		return err
	}
	if err := execAndHandle(tx,
		// language=SQL
		"INSERT INTO entry (name, image, description) VALUES ('testName', 'testImage', 'testDescription');"); err != nil {
		return err
	}
	if err := execAndHandle(tx,
		// language=SQL
		"SELECT @entry_id := LAST_INSERT_ID();"); err != nil {
		return err
	}
	if err := execAndHandle(tx,
		// language=SQL
		"INSERT INTO series (entry_id) VALUES (@entry_id);"); err != nil {
		return err
	}
	if err := execAndHandle(tx,
		// language=SQL
		"INSERT INTO anime (series_id, mal_reference) VALUES (LAST_INSERT_ID(), 'testMalReference');"); err != nil {
		return err
	}
	if err := execAndHandle(tx,
		// language=SQL
		"INSERT INTO episode (entry_id) VALUES (@entry_id);"); err != nil {
		return err
	}
	if err := execAndHandle(tx,
		// language=SQL
		"SELECT @episode_id := LAST_INSERT_ID();"); err != nil {
		return err
	}
	if err := execAndHandle(tx,
		// language=SQL
		"INSERT INTO subtitle (episode_id, language, status, subtitle_path) VALUES (@episode_id, 'testLanguage', 'testStatus', 'testSubtitlePath');"); err != nil {
		return err
	}
	if err := execAndHandle(tx,
		// language=SQL
		"INSERT INTO video (episode_id, video_path) VALUES (@episode_id, 'testVideoPath');"); err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func execAndHandle(tx *sqlx.Tx, query string) error {
	if _, err := tx.Exec(query); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}
	return nil
}
