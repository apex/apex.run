package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/russross/blackfriday"
)

var funcs = template.FuncMap{
	"readdir":  ioutil.ReadDir,
	"readfile": ioutil.ReadFile,
	"markdown": markdown,
	"slug":     slug,
	"noescape": func(s string) template.HTML { return template.HTML(s) },
}

const (
	flags = blackfriday.HTML_USE_XHTML |
		blackfriday.HTML_USE_SMARTYPANTS |
		blackfriday.HTML_SMARTYPANTS_FRACTIONS |
		blackfriday.HTML_SMARTYPANTS_DASHES |
		blackfriday.HTML_SMARTYPANTS_LATEX_DASHES

	extensions = blackfriday.EXTENSION_NO_INTRA_EMPHASIS |
		blackfriday.EXTENSION_TABLES |
		blackfriday.EXTENSION_FENCED_CODE |
		blackfriday.EXTENSION_AUTOLINK |
		blackfriday.EXTENSION_STRIKETHROUGH |
		blackfriday.EXTENSION_SPACE_HEADERS |
		blackfriday.EXTENSION_HEADER_IDS |
		blackfriday.EXTENSION_AUTO_HEADER_IDS |
		blackfriday.EXTENSION_BACKSLASH_LINE_BREAK |
		blackfriday.EXTENSION_DEFINITION_LISTS
)

func slug(s string) string {
	return strings.ToLower(strings.Replace(s, " ", "-", -1))
}

func markdown(b []byte) string {
	r := blackfriday.HtmlRenderer(flags, "", "")
	return string(blackfriday.MarkdownOptions(b, r, blackfriday.Options{
		Extensions: extensions,
	}))
}

var tmpl = template.Must(template.New("index").Funcs(funcs).ParseGlob("views/*.html"))

type page struct {
	Name string
	File string
}

var pages = []page{
	{"Installation", "installation.md"},
	{"AWS credentials", "aws-credentials.md"},
	{"Getting started", "getting-started.md"},
	{"Autocomplete", "autocomplete.md"},
	{"Structuring projects", "projects.md"},
	{"Structuring functions", "functions.md"},
	{"Deploying functions", "deploy.md"},
	{"Invoking functions", "invoke.md"},
	{"Listing functions", "list.md"},
	{"Deleting functions", "delete.md"},
	{"Building functions", "build.md"},
	{"Rolling back versions", "rollback.md"},
	{"Function hooks", "hooks.md"},
	{"Viewing log output", "logs.md"},
	{"Viewing metrics", "metrics.md"},
	{"Managing infrastructure", "infra.md"},
	{"Previewing with dry-run", "dryrun.md"},
	{"Environment variables", "env.md"},
	{"Omitting files", "ignore.md"},
	{"Understanding the shim", "shim.md"},
	{"Viewing documentation", "docs.md"},
	{"Upgrading Apex", "upgrade.md"},
	{"FAQ", "faq.md"},
	{"Links", "links.md"},
}

func main() {
	start := time.Now()

	data := struct {
		Docs  string
		Pages []page
	}{
		Docs:  filepath.Join("../apex/docs"),
		Pages: pages,
	}

	log.Printf("building %s", data.Docs)

	err := tmpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	log.Printf("build completed in %s", time.Since(start))
}
