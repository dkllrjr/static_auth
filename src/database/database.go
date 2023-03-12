package database

import (
    "os"
    "fmt"
    "static_auth/utils"
    "static_auth/terminal"
	bolt "go.etcd.io/bbolt"
)

//  ──────────────────────────────────────────────────────────────────────────

var DB *bolt.DB
var err error

//  ──────────────────────────────────────────────────────────────────────────

func InitDB(path string) {

	DB, err = bolt.Open(path, 0600, nil)
	utils.InternalErrorHandler(err)

	CreateBucket("Users")

    if os.Getenv("SA_SIGNUP") == "True" {
        username, hashedpassword := terminal.RegisterUser()
        CreateUser(username, hashedpassword)
    }
}


func CreateBucket(name string) {

	DB.Update(func(tx *bolt.Tx) error {
		_, err = tx.CreateBucketIfNotExists([]byte(name))
		utils.InternalErrorHandler(err)
		return nil
	})
}


func PutInBucket(bucket, key, value string) {

	DB.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(bucket))

		err := b.Put([]byte(key), []byte(value))
		utils.InternalErrorHandler(err)

		return nil
	})
}


func GetFromBucket(bucket, key string) string {

	var value string

	err = DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		value = string(b.Get([]byte(key)))

		return nil
	})

	utils.InternalErrorHandler(err)
	return value
}


func CheckKeyExists(bucket, key string) bool {

	var exists bool = false

	err = DB.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(bucket))

		err = b.ForEach(func(k, v []byte) error {
			if key == string(k) {
				exists = true
			}
			return nil
		})
		utils.InternalErrorHandler(err)

		return nil
	})

	utils.InternalErrorHandler(err)
	return exists
}


func CreateUser(username, hashedpassword string) {

	if CheckKeyExists("Users", username) {
		fmt.Println("User already exists")
	} else {
        PutInBucket("Users", username, hashedpassword)
    }
}
