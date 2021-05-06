package main

import (
	"embed"
	"fmt"
	"gorm.io/driver/mysql" //nolint:gci
	"gorm.io/gorm"
	"math/rand"
	"net/http"
	"time"

	//nolint:gci
)
import 	_ "github.com/go-sql-driver/mysql" //nolint:gofmt

type DataTG struct {
	Key   string `tag:"key" gorm:"key"`
	Value string `tag:"value" gorm:"value"`
}

func (DataTG)TableName()string  {
	return "tiangou"
}

//go:embed web
var addr embed.FS

func main()  {
	db, err := gorm.Open(mysql.Open(":@tcp()/test"),&gorm.Config{})
	if err != nil {
		fmt.Println("exec failed, ", err) //nolint:forbidigo
	}
	var tg []DataTG
	_= db.Find(&tg)

	http.Handle("/",http.FileServer(http.FS(addr)))
	http.HandleFunc("/web/data/tiangou", func(w http.ResponseWriter, r *http.Request) { //nolint:nolintlint,wsl
		rand.Seed(time.Now().Unix())
		i := rand.Int() % len(tg)                              //nolint:gosec
		w.Write([]byte("document.write(\""+tg[i].Value+"\")")) //nolint:errcheck
	})
	err = http.ListenAndServe(":59998", nil)
	if err != nil {
		fmt.Println("exec failed, ", err)
	}
}