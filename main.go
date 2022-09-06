package main

import (
  "log"
  "net/http"
  "os"

  "github.com/julienschmidt/httprouter"
)

func logRequest(req *http.Request) {
  log.Printf("Received request for %s", req.URL.String())
}

func mainPageHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
  logRequest(req)

  if req.URL.Path != "/" {
      http.NotFound(w, req)
      return
  }

  output := []byte(`<!DOCTYPE html>
<meta charset="UTF-8">
<p>&#128075; Welcome!</p>

<p>Please pardon the dust, this website is under construction!</p>

&#128119; &#128679; &#128679; &#128679; &#128679; &#128679; &#128679; &#128679; &#128679; &#128119;
`)
  w.Header().Add("Content-Type", "text/html")
  w.Write(output)
}

func main() {
  router := httprouter.New()
  router.GET("/", mainPageHandler)
  // Note: Some limitations of ServeFile are:
  // 1. that if there is no 'index.html' in the directory, this will show the directory.
  // 2. there is no Content-Type set on the file served.
  router.ServeFiles("/cats/*filepath", http.Dir("./cats"))
  // TODO: Add some markdown using github.com/gomarkdown/markdown (including XSS prevention using BlueMonday).

  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
  }


  log.Printf("Listening on port=%s", port)
  if err := http.ListenAndServe(":" + port, router); err != nil {
    log.Fatal(err)
  }
  log.Printf("Closing...")
}


