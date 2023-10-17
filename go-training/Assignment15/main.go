package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Student struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Age   string `json:"age"`
	Grade string `json:"grade"`
}

var students = map[uint64]Student{123: {"123", "HELLO", "23", "A"}, 234: {"234", "world", "34", "B"}}

var ErrDataNotPresent = errors.New("data is not their")

func FetchStudent(studentsId uint64) (Student, error) {
	value, ok := students[studentsId]

	if !ok {
		return Student{}, ErrDataNotPresent
	}

	return value, nil
}

func GetStudent(w http.ResponseWriter, r *http.Request) {

	studentsIdString := r.URL.Query().Get("student_id")

	studentsId, err := strconv.ParseUint(studentsIdString, 10, 64)

	if err != nil {
		log.Println("Error: ", err)

		errorInConversion := map[string]string{"msg": "not a valid number"}

		jsonData, err := json.Marshal(errorInConversion)

		if err != nil {
			log.Println("Error while converting error to json", err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, http.StatusText(http.StatusInternalServerError))
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonData)
		return
	}

	sData, err := FetchStudent(studentsId)

	if err != nil {
		fetchError := map[string]string{"msg": "user not found"}
		errData, err := json.Marshal(fetchError)
		if err != nil {
			log.Println("Error while parsing fetchuser error conversion: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, http.StatusText(http.StatusInternalServerError))
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		w.Write(errData)
		return
	}

	studentData, err := json.Marshal(sData)

	if err != nil {
		log.Println("Error while converting user data to json", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, http.StatusText(http.StatusInternalServerError))
		return
	}

	w.Write(studentData)
}
func main() {

	http.HandleFunc("/students", GetStudent)
	panic(http.ListenAndServe(":8080", nil))

}
