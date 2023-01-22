package main

import (
  "fmt"
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
  http.Error(w, `<!DOCTYPE html>
<meta charset="UTF-8">
<p>Something failed on our end!</p>`,
  http.StatusInternalServerError)
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

func renderPageHTML(w http.ResponseWriter, content string) {
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
  output += content
  output += `</section>
 <footer>
  <div>
    <a href="https://www.linkedin.com/in/jchaffraix"><img class="icon" src="/style/icons/linkedin.svg" alt="Linkedin icon"></img></a>
    <a href="/licenses"><img class="icon" src="/style/icons/file-regular.svg" alt="licenses"></img></a>
  </div>
  Made with <img class="icon" src="/style/icons/heart.svg" alt="love"></img>.<br/>
</footer>
</body>
</html>`
  w.Header().Add("Content-Type", "text/html")
  w.Write([]byte(output))
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
  renderPageHTML(w, string(markdown.ToHTML(md, nil, renderer)))
}

type License struct {
  // Path relative to root of this repository.
  Path string
  // Header that explains the dependencies along with link to the project.
  Intro string
}

func licensesHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
  logRequest(req)

  licenses := []License{
    License {
      Path: "html/style/font/OFL.txt",
      Intro: `The font for the site is DM mono. It can be found on <a href="https://github.com/googlefonts/dm-mono">GitHub</a> and its license is:`,
    },
    License {
      Path: "",
      Intro: `The icons come from the excellent <a href="http://fontawesome.com">FontAwesome Project</a>. Its license can be found on <a href="https://fontawesome.com/license/free">here</a>.`,
    },
  }

  main_license_str, err := ioutil.ReadFile("LICENSE")
  if err != nil {
    log.Printf("Error reading main license %+v", err)
    renderFailedPage(w)
    return

  }
  content := fmt.Sprintf(`<div>The main website can be found on <a href="https://github.com/jchaffraix/JulienCorner">GitHub</a>. Its license is:<br/><br/><pre>%s</pre></div>`, main_license_str)
  for _, license := range licenses {
    if license.Path != "" {
      license_str, err := ioutil.ReadFile(license.Path)
      if err != nil {
        log.Printf("Error reading license file %s, err=%+v", license.Path, err)
        renderFailedPage(w)
        return
      }
      content += fmt.Sprintf(`<br/><hr/><br/>
<div>%s<br><br><pre>%s</pre></div>`, license.Intro, license_str)
    } else {
      content += fmt.Sprintf(`<br/><hr/><br/>
<div>%s</div>`, license.Intro)
    }
  }

  renderPageHTML(w, content)
}
 
func main() {
  router := httprouter.New()
  router.GET("/", mainPageHandler)
  router.GET("/licenses", licensesHandler)
  // Note: Some limitations of ServeFile are:
  // 1. that if there is no 'index.html' in the directory, this will show the directory.
  // 2. there is no Content-Type set on the file served.
  router.ServeFiles("/cats/*filepath", http.Dir("html/cats"))
  router.ServeFiles("/habits/*filepath", http.Dir("html/habits"))
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


