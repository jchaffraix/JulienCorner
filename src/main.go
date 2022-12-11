package main

import (
  "io"
  "io/ioutil"
  "log"
  "net/http"
  "os"
  "strconv"

  "github.com/gomarkdown/markdown"
  "github.com/gomarkdown/markdown/ast"
  "github.com/gomarkdown/markdown/html"
  "github.com/julienschmidt/httprouter"
)

func logRequest(req *http.Request) {
  log.Printf("Received request for %s", req.URL.String())
}

func renderFailedPage(w http.ResponseWriter) {
  w.Header().Add("Content-Type", "text/html")
  w.Write([]byte(`<!DOCTYPE html>
<meta charset="UTF-8">
<p>Something failed on our end!/p>`))
  w.WriteHeader(500)
}

func customizeCSS(w io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {
  if _, ok := node.(*ast.Heading); ok {
    level := strconv.Itoa(node.(*ast.Heading).Level)

    if entering && level == "1" {
      w.Write([]byte(`<h1 class="title is-1 has-text-centered">`))
    } else if entering {
      w.Write([]byte("<h" + level + ">"))
    } else {
      w.Write([]byte("</h" + level + ">"))
    }

    return ast.GoToNext, true
  }

  return ast.GoToNext, false
}

func mainPageHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
  logRequest(req)

  if req.URL.Path != "/" {
    http.NotFound(w, req)
    return
  }

  md, err := ioutil.ReadFile("html/index.md")
  if err != nil {
    log.Printf("Error reading main file %+v", err)
    renderFailedPage(w)
    return
  }
  opts := html.RendererOptions{
    Flags: html.FlagsNone,
    RenderNodeHook: customizeCSS,
  }
  renderer := html.NewRenderer(opts)
  output := `<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <link rel="stylesheet" href="/style/main.css">
  <meta name="viewport" content="width=device-width, initial-scale=1">
</head>
<body>
<section>
`
  output += string(markdown.ToHTML(md, nil, renderer))
  output += `</section>
 <footer>
  <a href="https://www.linkedin.com/in/jchaffraix"><img class="icon" src="/style/icons/linkedin.svg"></img></a><br/>
  Made with <img class="icon" src="/style/icons/heart.svg"></img>.
</footer>
</body>
</html>`
  w.Header().Add("Content-Type", "text/html")
  w.Write([]byte(output))
}

func main() {
  router := httprouter.New()
  router.GET("/", mainPageHandler)
  // Note: Some limitations of ServeFile are:
  // 1. that if there is no 'index.html' in the directory, this will show the directory.
  // 2. there is no Content-Type set on the file served.
  router.ServeFiles("/cats/*filepath", http.Dir("html/cats"))
  router.ServeFiles("/style/*filepath", http.Dir("html/style"))
  // TODO: Add XSS prevention using BlueMonday.

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


