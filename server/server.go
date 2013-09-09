package server

import (
  "net/http"
  "log"
  "html/template"
  "fmt"
  "github.com/cratonica/embed"
  "github.com/hayesgm/journal/assets"
)


func getHandler(name, value string) func(w http.ResponseWriter, req *http.Request) {
  return func(w http.ResponseWriter, req *http.Request) {
    tmpl := template.Must(template.New(name).Parse(value))
    tmpl.Execute(w, nil)
  }
}

func RegisterAssets(root string) {
  log.Println("Registering Static Assets...")

  assetMap, _ := embed.Unpack(assets.Assets)

  for asset, value := range assetMap {
    log.Printf("\tRegistering %v\n", asset)
    route := fmt.Sprintf("/%s", asset)
    
    http.Handle(route, http.HandlerFunc(getHandler(asset, string(value))))
    if (asset == root) {
      http.Handle("/", http.HandlerFunc(getHandler(asset, string(value))))
      log.Printf("\tRegistered %v to /\n", asset)
    }
    log.Printf("\tRegistered %v to %v\n", asset, route)
  }
}