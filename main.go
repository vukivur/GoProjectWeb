package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name                  string
	Age                   uint16
	Money                 int16
	Avg_grades, Happiness float64
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s. He is %d and he has "+
		"money equal: %d", u.Name, u.Age, u.Money)
}

// * - не копируем, а обращаемся к объекту конкретному по ссылке.
// Если структура небольшая - лучше не как ссылку, а копированием. Если средняя или большая - ссылкой лучше. (Документация)
func (u *User) setNewName(newName string) {
	u.Name = newName
}

// С помощью первого параметра сможем обращаться к странице, и что-либо показывать пользователю.
// Request - отследить запрос.
func home_page(w http.ResponseWriter, r *http.Request) {
	michael := User{"Michael", 25, -50, 4.4, 0.78}
	//создаем форматированную строку.
	michael.setNewName("MichaelJackson")
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	// помещаем щаблон на страницу
	tmpl.Execute(w, michael)
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contacts page is here.")
}

func handleRequest() {
	// Отслеживаем URL адреса. При переходе на главную страницу "/" - вызывается метод home_page
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	// Запускаем сервер на выбранном порту.
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()
}
