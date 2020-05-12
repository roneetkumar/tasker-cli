package main

import (
	"fmt"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/roneetkumar/tasker/cmd"
	"github.com/roneetkumar/tasker/db"
)

func main() {

	home, _ := homedir.Dir()

	dbpath := filepath.Join(home, "tasks.db")

	err := db.Init(dbpath)

	must(err)

	must(cmd.RootCMD.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
