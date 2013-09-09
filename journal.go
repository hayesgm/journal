package main

import (
  "os"
  "log"
  "path"
  "github.com/hayesgm/journal/server"
  "github.com/hayesgm/mirrors"
  "net/http"
  "flag"
  "github.com/coreos/go-etcd/etcd"
)

/*
  Journal is a simple blogging platform written in Go, backed by etcd and DNSimple

  Journal servers will add and remove themselves from POOL records for
  round-robin DNS-based load distribution.
*/

func registerStaticAssets() {
  wd, err := os.Getwd()
  if (err != nil) {
    log.Fatal("Unable to get cwd", err)
  }
  assetsDir := path.Join(wd, "assets")

  server.RegisterAssets(assetsDir, "journal.html")
}

func main() {
  var domain, token string

  flag.StringVar(&domain, "domain", "", "domain to host journal")
  flag.StringVar(&token, "token", "", "dnsimple token to manage domain")
  flag.Parse()

  if len(domain) == 0 || len(token) == 0 {
    log.Fatal("Must include domain and token")
  }

  registerStaticAssets()
  
  mirrors.Join(etcd.NewClient(), domain, token, "journal")

  err := http.ListenAndServe(":2222", nil)
  if err != nil {
    log.Fatal("Failed to launch server", err)
  }
}