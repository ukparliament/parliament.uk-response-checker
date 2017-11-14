package main

import (
  "bufio"
  "fmt"
  "log"
  "net/http"
  "os"
)

const baseUrl = "https://beta.parliament.uk"

type Link struct {
  url   string
  code  int
}

func main(){
  // 1. Create a new file, result.txt (if it doesn't already exist)
  resultFile, err := os.Create("results.txt")
  if err != nil {
    log.Fatal(err)
  }

  // 2. Load file of links
  linksFile, err := os.Open("links.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer linksFile.Close()

  // 3. Read each line
  scanner := bufio.NewScanner(linksFile)

  // 4. For each line
  for scanner.Scan(){
    // 4a. Create new Link object
    link := Link{url: scanner.Text()}

    // 4b. Visit link
    response, err := http.Get(baseUrl + link.url)
    if err != nil {
      log.Fatal(err)
    }
    defer response.Body.Close()

    // 4c. Get response code
    link.code = response.StatusCode

    // 4d. Write Response to file
    writer := bufio.NewWriter(resultFile)
    fmt.Fprintf(writer, "%v, %v\n", link.url, link.code)

    writer.Flush()
  }
  if err := scanner.Err(); err != nil {
    fmt.Fprintln(os.Stderr, "Issues reading input:", err)
  }

}
