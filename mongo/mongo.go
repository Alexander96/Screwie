package mongo

import (
    "fmt"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)
type Person struct {
    Name string
    Phone string
}
func Start() *mgo.Session {
    session, err := mgo.Dial("localhost:27017")
    if err != nil {
            panic(err)
    }
    c := session.DB("screwiebd1").C("people")
    err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
	           &Person{"Cla", "+55 53 8402 8510"})
    if err != nil {
        fmt.Println("some error 1")
    }
    result := Person{}
    err = c.Find(bson.M{"name": "Ale"}).One(&result)
    if err != nil {
        fmt.Println("some error")
    }
    return session
}
