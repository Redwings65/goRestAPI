package main

import (
	"fmt"
  "strconv"
	"net/http"
  "encoding/json"
  "github.com/jinzhu/gorm"
  "github.com/gorilla/mux"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

type Shoe struct {
  gorm.Model
  Name string
  Size int
  Price string
}

func initialMigration(){
  datastoreName := "postgres://beajquei:3yKP5cAxHHjJ2eE2W7foPt7JIN2HC-_a@pellefant.db.elephantsql.com:5432/beajquei"
  db, err = gorm.Open("postgres", datastoreName)
  if err != nil{
    fmt.Println(err.Error())
    panic("couldnt connect to db")
  }
  defer db.Close()
  db.AutoMigrate(&Shoe{})
}

func addShoeSize(w http.ResponseWriter, r *http.Request) {
  datastoreName := "postgres://beajquei:3yKP5cAxHHjJ2eE2W7foPt7JIN2HC-_a@pellefant.db.elephantsql.com:5432/beajquei"
  db, err = gorm.Open("postgres", datastoreName)
  if err != nil{
    fmt.Println(err.Error())
    panic("couldnt connect to db")
  }
  defer db.Close()

  vars := mux.Vars(r)
  s := vars["size"]

  size, err := strconv.Atoi(s);
  if err == nil{
    db.Create(&Shoe{Size: size, Name: "Test Shoe", Price: "100"})
    fmt.Fprintf(w,"New Shoe Size Added")
  }
}

//this works fine
func getAllShoes(w http.ResponseWriter, r *http.Request) {
  datastoreName := "postgres://beajquei:3yKP5cAxHHjJ2eE2W7foPt7JIN2HC-_a@pellefant.db.elephantsql.com:5432/beajquei"
  db, err = gorm.Open("postgres", datastoreName)
  if err != nil{
    fmt.Println(err.Error())
    panic("couldnt connect to db")
  }
  defer db.Close()
  var shoes []Shoe
  db.Find(&shoes)
  json.NewEncoder(w).Encode(shoes)
}

func returnTrueSize(w http.ResponseWriter, r *http.Request) {
  datastoreName := "postgres://beajquei:3yKP5cAxHHjJ2eE2W7foPt7JIN2HC-_a@pellefant.db.elephantsql.com:5432/beajquei"
  db, err = gorm.Open("postgres", datastoreName)
  if err != nil{
    fmt.Println(err.Error())
    panic("couldnt connect to db")
  }
  defer db.Close()

  var shoes []Shoe
  db.Find(&shoes)
//this needs to be dynamically allocated
  var sum = 0;
  var size = 0;
  for i := 0; i < len(shoes); i++ {
    sum += shoes[i].Size
    size++
  }
  json.NewEncoder(w).Encode(float64(sum)/float64(size))
}
