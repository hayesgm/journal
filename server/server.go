package server

import (
  "net/http"
  "log"
  "html/template"
  "path"
  "io/ioutil"
  "fmt"
)

func getHandler(assetsDir, filename string) func(w http.ResponseWriter, req *http.Request) {
  return func(w http.ResponseWriter, req *http.Request) {
    tmpl := template.Must(template.ParseFiles(path.Join(assetsDir,filename)))
    tmpl.Execute(w, nil)
  }
}

func RegisterAssets(assetsDir, root string) {
  log.Println("Registering Static Assets...")

  files, err := ioutil.ReadDir(assetsDir)
  if err != nil {
    panic(err) // This is serious
  }

  for _, file := range files {
    log.Printf("\tRegistering %v\n", file.Name())
    route := fmt.Sprintf("/%s", file.Name())
    
    http.Handle(route, http.HandlerFunc(getHandler(assetsDir, file.Name())))
    if (file.Name() == root) {
      http.Handle("/", http.HandlerFunc(getHandler(assetsDir, file.Name())))
      log.Printf("\tRegistered %v to /\n", file.Name())
    }
    log.Printf("\tRegistered %v to %v\n", file.Name(), route)
  }
}