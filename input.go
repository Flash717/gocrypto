package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"

	"html/template"
	"io"
	"net/http"

	"gopkg.in/go-playground/validator.v9"
)

type UserData struct {
	Name  string `validate:"required"`
	Email string `validate:"required,email"`
}

func teststrconv() {
	// testing strconv package
	i, err := strconv.Atoi("-42")
	b, err := strconv.ParseBool("true")

	fmt.Printf("int %d and bool %t, error: %v \n", i, b, err)
}

func teststrings() {
	// testing strings package
	fmt.Printf("out: >%s<\n", strings.Trim("???Hello, Gophers!!!", "?!"))
	fmt.Printf("out: >%s<\n", strings.TrimSpace(" \t\n Hello, Gophers \n\t\r\n"))
	fmt.Printf("out: >%s<\n", strings.ToUpper("Hello Gophers and others"))
	fmt.Printf("out: >%s<\n", strings.ToLower("Hello Gophers and others"))
	fmt.Printf("out: >%s<\n", strings.ToTitle("Hello Gophers and others"))
}

func testutf8() {
	// testing utf8 package
	valid := []byte("Hello, world")
	fmt.Println(utf8.Valid(valid))
	// ValidRune
	// ValidString

	buf := make([]byte, 1)
	utf8.EncodeRune(buf, 'a')
	fmt.Println(buf)

	output, _ := utf8.DecodeRune(buf)
	fmt.Printf("%c\n", output)
	// utf8.DecodeRuneInString
	// utf8.DecodeLastRune
}

func testregex() {
	// testing regex package
	matched, _ := regexp.Match("[0-9]{3}-[0-9]{3}-[0-9]{4}", []byte("123-123-1234"))
	fmt.Printf("regex is %t\n", matched)
	matched, _ = regexp.Match("[0-9]{3}-[0-9]{3}-[0-9]{4}", []byte("123-123-123A"))
	fmt.Printf("regex is %t\n", matched)
}

func testvalidator() {
	// testing validator.v9
	user := &UserData{
		Name:  "Alice",
		Email: "Alice@gmail.com",
	}
	validate := validator.New()
	err := validate.Struct(user)

	fmt.Printf("and the error is %v\n", err)
}

func handlerBad(w http.ResponseWriter, r *http.Request) {
	// bad practice, no validation, could be a script thingy
	io.WriteString(w, r.URL.Query().Get("param1"))
}

func handlerGood(w http.ResponseWriter, r *http.Request) {
	param1 := r.URL.Query().Get("param1")
	tmpl := template.New("hello")
	tmpl, _ = tmpl.Parse(`{{define "T"}}{{.}}{{end}}`)
	tmpl.ExecuteTemplate(w, "T", param1)
}

func testxsrf() {
	// danger danger!!!
	// test by calling http://localhost:8080/?param1=<script>alert('Hello')</script>

	http.HandleFunc("/", handlerGood)
	http.ListenAndServe(":8080", nil)
}

func testSQLinjection() {
	var input string

	fmt.Scanln(input)

	// bad practice!!! do not use string interpolation, could inject 'abc OR 1=1' to get all users
	// rows, err := db.Query(fmt.Sprintf("SELECT * FROM user WHERE id = %s", input))

	// better practice, use parameter
	// rows, err := db.Query("SELECT * FROM user WHERE id = ?", input)
}

func testInput() {
	//teststrconv()

	//teststrings()

	//testutf8()

	//testregex()

	//testvalidator()

	testxsrf()
}
