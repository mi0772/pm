package database

import (
	"encoding/json"
	"fmt"
	"io"

	"log"
	"mi0772/pm/models"
	"mi0772/pm/security"
	"mi0772/pm/userio"
	"os"
	"strings"
	"time"
)

const DB_NAME = "pm.db"

var userHomeDir string
var MasterPassword string

func init() {
	var err error
	userHomeDir, err = os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
}

func AccessToDatabase() error {
	if !ExistDB() {
		password, err := userio.ReadNewMasterPassword("enter a new master password for your database")
		if err != nil {
			return err
		}
		MasterPassword = fmt.Sprintf("%s", password)
		fmt.Printf("\n\nyour new master password is %s, take a note and don't forget it\n", MasterPassword)
		_, err = CreateNewDatabase()
		if err != nil {
			return err
		}
	}

	if len(MasterPassword) == 0 {
		password, err := userio.ReadPassword("enter master password")
		if err != nil {
			return err
		}
		MasterPassword = string(password)
	}
	fetchAllRecords()
	return nil
}

func filter[T any](ss []T, test func(T) bool) (ret []T) {
	for _, v := range ss {
		if test(v) {
			ret = append(ret, v)
		}
	}
	return
}

func fetchAllRecords() ([]models.Entry, int) {

	var result models.Entries
	databaseFileName, err := os.Open(userHomeDir + "/" + DB_NAME)
	if err != nil {
		databaseFileName, err = CreateNewDatabase()
	}
	defer databaseFileName.Close()

	byteValue, err := io.ReadAll(databaseFileName)
	if err != nil && err != io.EOF {
		panic(err)
	}

	if len(byteValue) > 0 {
		//decrypt content with master password
		byteValue, err = security.Decrypt(byteValue, MasterPassword)
		if err != nil {
			fmt.Println("\ncannot decrypt database, wrong master password ?")
			os.Exit(0)
		}
		json.Unmarshal(byteValue, &result)
	}
	size := len(result.Entries)

	return filter(result.Entries, func(entry models.Entry) bool {
		return !entry.Deleted
	}), size
}

func CreateNewDatabase() (*os.File, error) {
	newFile, err := os.Create(userHomeDir + "/" + DB_NAME)
	if err != nil {
		return nil, err
	}

	Memorize("pm", "master password", MasterPassword)
	return newFile, nil
}

func ChangeMasterPassword(newPassword string) {
	content, _ := fetchAllRecords()

	var entries = models.Entries{Entries: content}
	file, _ := json.MarshalIndent(entries, "", " ")
	var err error
	file, err = security.Encrypt(file, newPassword)
	if err != nil {
		log.Fatal(err)
	}
	_ = os.WriteFile(userHomeDir+"/"+DB_NAME, file, 0644)
}

func Memorize(label, account, password string) {

	dbArray, totalRecord := fetchAllRecords()
	db := dbArray[:]
	var founded bool = false

	//cerco eventuale record esistente

	for i, v := range db {
		if v.Label == label {
			db[i].ModifiedAt = time.Now()
			db[i].Password = password
			db[i].Account = account
			founded = true
			break
		}
	}

	if !founded {
		entry := models.Entry{Id: totalRecord + 1, Label: label, Account: account, Password: password}
		entry.CreatedAt = time.Now()
		db = append(db, entry)
	}

	cryptAndSave(db)
}

func Search(label string) []models.Entry {
	var result = make([]models.Entry, 0)
	entries, _ := fetchAllRecords()

	for _, v := range entries {
		if label == "*" || (len(label) != 0 && strings.Contains(v.Label, label)) {
			result = append(result, v)
		}
	}

	return result
}

func ExistDB() bool {
	_, error := os.Stat(userHomeDir + "/" + DB_NAME)
	return !os.IsNotExist(error)
}

func Delete(id *int) {
	dbArray, _ := fetchAllRecords()
	db := dbArray[:]

	//cerco eventuale record esistente
	for i, v := range db {
		if v.Id == *id {
			db[i].Deleted = true
		}
	}
	cryptAndSave(db)
}

func cryptAndSave(db []models.Entry) {
	var entries = models.Entries{Entries: db}
	file, _ := json.MarshalIndent(entries, "", " ")
	var err error
	file, err = security.Encrypt(file, MasterPassword[:])
	if err != nil {
		log.Fatal(err)
	}
	_ = os.WriteFile(userHomeDir+"/"+DB_NAME, file, 0644)
}
