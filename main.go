package main

import (
	"embed"
	"fmt"
	"gorm.io/driver/mysql" //nolint:gci
	"gorm.io/gorm"
	"io/fs"
	"math/rand"
	"net/http"
	"path"
	"time"
	//nolint:gci
)
import _ "github.com/go-sql-driver/mysql" //nolint:gofmt

type DataTG struct {
	Key   string `tag:"key" gorm:"key"`
	Value string `tag:"value" gorm:"value"`
}

func (DataTG) TableName() string {
	return "tiangou"
}

//go:embed web
var addr embed.FS

func main() {
	db, err := gorm.Open(mysql.Open(":@tcp()/test"), &gorm.Config{})
	if err != nil {
		fmt.Println("exec failed, ", err) //nolint:forbidigo
	}
	var tg []DataTG
	_ = db.Find(&tg)

	http.Handle("/", AssetHandler("", addr, "web"))
	http.HandleFunc("/data/tiangou", func(w http.ResponseWriter, r *http.Request) { //nolint:nolintlint,wsl
		rand.Seed(time.Now().Unix())
		i := rand.Int() % len(tg)                                  //nolint:gosec
		w.Write([]byte("document.write(\"" + tg[i].Value + "\")")) //nolint:errcheck
	})
	err = http.ListenAndServe(":59998", nil)
	if err != nil {
		fmt.Println("exec failed, ", err)
	}

}

// AssetHandler returns an http.Handler that will serve files from
// the Assets embed.FS. When locating a file, it will strip the given
// prefix from the request and prepend the root to the filesystem.
func AssetHandler(prefix string, assets embed.FS, root string) http.Handler {
	handler := fsFunc(func(name string) (fs.File, error) {
		assetPath := path.Join(root, name)

		// If we can't find the asset, fs can handle the error
		file, err := assets.Open(assetPath)
		if err != nil {
			return nil, err
		}

		// Otherwise assume this is a legitimate request routed correctly
		return file, err
	})

	return http.StripPrefix(prefix, http.FileServer(http.FS(handler)))
}

type fsFunc func(name string) (fs.File, error)

func (f fsFunc) Open(name string) (fs.File, error) {
	return f(name)
}
