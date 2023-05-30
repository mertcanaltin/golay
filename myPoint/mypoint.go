package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

// User struct
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Age      int    `json:"age"`
}

func main() {
	// Rastgele 11 kullanıcı oluşturma
	users := generateUsers(11)

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		ageParam := r.URL.Query().Get("age")
		if ageParam != "" {
			age, err := strconv.Atoi(ageParam)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "Invalid age parameter")
				return
			}
			filteredUsers := filterUsersByAge(users, age)
			if len(filteredUsers) == 0 {
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprintf(w, "No users found with the specified age")
				return
			}
			responseJSON(w, filteredUsers)
		} else {
			responseJSON(w, users)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Kullanıcıları yaşa göre filtreleme
func filterUsersByAge(users []User, age int) []User {
	var filteredUsers []User
	for _, user := range users {
		if user.Age == age {
			filteredUsers = append(filteredUsers, user)
		}
	}
	return filteredUsers
}

// JSON yanıtını gönderme
func responseJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// Rastgele kullanıcılar oluşturma
func generateUsers(count int) []User {
	var users []User
	for i := 0; i < count; i++ {
		user := User{
			ID:       i + 1,
			Name:     fmt.Sprintf("User %d", i+1),
			Username: fmt.Sprintf("user%d", i+1),
			Age:      rand.Intn(50) + 20, // Yaş aralığı 20-69
		}
		users = append(users, user)
	}
	return users
}
