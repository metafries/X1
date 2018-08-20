package main

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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
	log.Println("[INFO] mgo - [Collection.Count]: Number of docs of club_info:", n)

	// Perform simple query.
	ci := clubInfo{}
	clubinfo.Find(bson.M{"id": 1}).One(&ci)
	log.Println("[INFO] mgo - [Collection.Find.One]: Docs of club_info at id = 1:", ci)

	// Query with expression.
	query := bson.M{
		"id": bson.M{
			"$gt": 0,
		},
		"league": bson.M{
			"$in": []string{"Premier League", "Lega Serie A"},
		},
	}

	var club Club
	err = clubinfo.Find(query).All(&club)
	if err != nil {
		log.Fatal("[FATAL] mgo - [Collection.Find.All]: ", err)
	}
	log.Println("[INFO] mgo - [Collection.Find.All]: Query Results:", club)

	// Use select to get names only
	names := []struct {
		Name string `bson:"name"`
	}{}

	err = clubinfo.Find(query).Select(bson.M{"name": 1}).All(&names)
	if err != nil {
		log.Fatal("[FATAL] mgo - [Collection.Find.Select.All]: ", err)
	}
	log.Println("[INFO] mgo - [Collection.Find.Select.All]: Query Results(names):", names)

	// Insert
	newcl := clubInfo{ID: 2, Name: "Juventus", Ranking: "1 (2017-18)", League: "Lega Serie A"}
	if err := clubinfo.Insert(newcl); err != nil {
		log.Fatal("[FATAL] mgo - [Collection.Insert]: ", err)
	}
}
