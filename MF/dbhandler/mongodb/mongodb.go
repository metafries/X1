package main

import (
	"log"

	"gopkg.in/mgo.v2"
)

type clubInfo struct {
	ID      int    `bson:"id"`
	Name    string `bson:"name"`
	Ranking string `bson:"ranking(2017-2018)"`
	League  string `bson:"league"`
}

type Club []clubInfo

func main() {

	session, err := mgo.Dial("mongodb://127.0.0.1")
	if err != nil {
		log.Fatal("[FATAL] mgo - [Dial]: ", err)
	}
	defer session.Close()

	// Get the named collection.
	clubinfo := session.DB("MF").C("club_info")

	// Get the total number of documents in the collection.
	n, _ := clubinfo.Count()
	log.Println("[INFO] mgo - [Collection.Count]: Number of docs of club_info is", n)

}
