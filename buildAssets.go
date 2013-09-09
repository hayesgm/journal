package main

import (
  "io/ioutil"
  "github.com/cratonica/embed"
)

func main() {
  // Create a resource map from files in a directory
  resourceMap, _ := embed.CreateFromFiles("assets")

  // Pack the files into a serializeable byte buffer
  packed, _ := embed.Pack(resourceMap)

  // Generate a .go file that we can include in our project
  goCode := embed.GenerateGoCode("assets", "Assets", packed)
  err := ioutil.WriteFile("assets/assets.go", []byte(goCode), 0666)
  if err != nil {
    panic(err)
  }
}