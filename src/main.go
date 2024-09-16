package main

import (
  "errors"
  "fmt"
  "io/fs"
  "io/ioutil"
  "log"
  "net/http"
  "os"
  "path/filepath"
  "strings"

  "github.com/julienschmidt/httprouter"
)

func logRequest(req *http.Request) {
  log.Printf("Received request for %s", req.URL.String())
}

func renderFailedPage(w http.ResponseWriter) {
  w.Header().Add("Content-Type", "text/html")
  w.WriteHeader(http.StatusInternalServerError)
  w.Write([]byte(`<!DOCTYPE html>
<meta charset="UTF-8">
<p>Something failed on our end!</p>`))
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

  main_page, err := ioutil.ReadFile("html/index.html")
  if err != nil {
    log.Printf("Error reading homepage %+v", err)
    renderFailedPage(w)
    return

  }
  w.Header().Add("Content-Type", "text/html")
  w.Write([]byte(main_page))
}

func robotsPageHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
  logRequest(req)

  robots, err := ioutil.ReadFile("robots.txt")
  if err != nil {
    log.Printf("Error reading robots.txt %+v", err)
    renderFailedPage(w)
    return

  }
  w.Header().Add("Content-Type", "text/plain")
  w.Write([]byte(robots))
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

func isRawPage(path string) bool {
  // TODO: Switch to a more efficient DS.
  return strings.HasPrefix(path, "presentations") || strings.HasPrefix(path, "cats") || strings.HasPrefix(path, "habits")
}

func isStaticPageAllowed(path string) bool {
  // TODO: Switch to a more efficient DS.
  for _, prefix := range []string{"pages", "posts", "presentations", "cats", "habits", "style"} {
    if strings.HasPrefix(path, prefix) {
      return true
    }
  }
  return false
}

func staticHtmlPageHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
  logRequest(req)

  // We drop the leading /.
  originalPath := req.URL.String()
  path := filepath.Clean(originalPath[1:])

  // Sanity check to prevent opening some filesystem files.
  if !isStaticPageAllowed(path) {
    log.Printf("... Return 404 on request to %s (invalid prefix on %s)", originalPath, path)
    http.NotFound(w, req)
    return
  }

  renderRawPage := isRawPage(path)
  isIcon := strings.HasPrefix(path, "style/icons")
  path = filepath.Join("html", path)
  content, err := ioutil.ReadFile(path)
  if err != nil {
    // TODO: Handle "is a directory" as a 404 (or index.html redirect) instead of a 500.
    if errors.Is(err, fs.ErrNotExist) {
      log.Printf("... Couldn't find path %s (file at %s, error=%+v)", originalPath, path, err)
      http.NotFound(w, req)
    } else {
      log.Printf("... Failed request to %s (can't read file at %s, error=%+v)", originalPath, path, err)
      renderFailedPage(w)
    }
    return
  }

  extension := filepath.Ext(path)
  switch (extension) {
    case ".html":
      if renderRawPage {
        w.Header().Add("Content-Type", "text/html")
        w.Write(content)
      } else {
        renderPageHTML(w, string(content))
      }
    case ".css":
      w.Header().Add("Content-Type", "text/css")
      w.Write(content)
    case ".js":
      w.Header().Add("Content-Type", "text/javascript")
      w.Write(content)
    case ".jpeg":
      w.Header().Add("Content-Type", "image/jpeg")
      w.Write(content)
    case ".png":
      w.Header().Add("Content-Type", "image/png")
      w.Write(content)
    case ".svg":
      if isIcon {
        w.Header().Add("Content-Type", "image/svg+xml")
      } else {
        w.Header().Add("Content-Type", "text/svg")
      }
      w.Write(content)
    default:
      log.Printf("... Failed request to %s (unknown extension %s)", originalPath, extension)
      renderFailedPage(w)
  }
}

func main() {
  router := httprouter.New()
  router.GET("/", mainPageHandler)
  router.GET("/robots.txt", robotsPageHandler)
  router.GET("/licenses", licensesHandler)
  router.GET("/posts/*filepath", staticHtmlPageHandler)
  router.GET("/pages/*filepath", staticHtmlPageHandler)
  router.GET("/presentations/*filepath", staticHtmlPageHandler)
  router.GET("/cats/*filepath", staticHtmlPageHandler)
  router.GET("/habits/*filepath", staticHtmlPageHandler)
  router.GET("/style/*filepath", staticHtmlPageHandler)
  // Do not use router.ServerFile, prefer staticHtmlPageHandler instead as it supports:
  // 1. Add a meaningful Content-Type based on the path.
  // 2. Doesn't list the content of a directory.

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


