package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	//"golang.org/x/crypto/bcrypt"

	_ "github.com/go-sql-driver/mysql"
)

var tmpl *template.Template

type Exercise struct {
	ExerciseID   int
	WorkoutID    int
	ExerciseName string
}

type Workout struct {
	WorkoutID   string
	WorkoutName string
	Date        string
}

type Tested struct {
	WorkoutName string
	WorkoutID   string
	EName       []Exercise

	//Enames []string
}

var id string
var wid string
var exerciseID string
var exerciselist []Exercise

func createworkout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/createworkout", r.Method)
	login, _ := template.ParseFiles("html/createworkout.html")
	login.Execute(w, nil)
	//Entry.userID := get logged in ID()
	//Entry.workoutID := get current workout ()
	//Entry.setNum := r.Form["Sets"] ???
	//Entry.reps := r.Form["Reps"]
	//Entry.weight := r.Form["Weight"]
	//Entry.exerciseID := getExercise ID
	//Entry.exerciseName := getExercise name from form

}

func register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/register", r.Method)
	login, _ := template.ParseFiles("html/register.html")
	login.Execute(w, nil)
}

func registerauth(w http.ResponseWriter, r *http.Request) {
	var validationflag bool
	var msg string

	fmt.Println("/registerauth", r.Method)
	r.ParseForm()
	username := r.Form["username"]
	password := r.Form["password"]
	fmt.Println(username, password)
	user := cleanupBracket(username)
	pswd := cleanupBracket(password)
	validationflag = ValidateUserPass(user, pswd)

	if validationflag == false {
		insertUserPass(user, pswd)
		msg = "Your account has been successfully registered."
		regmsg, _ := template.ParseFiles("html/registered.html")
		regmsg.Execute(w, msg)
	}
	if validationflag == true {
		msg = "The username and/or password your entered already exists."
		regmsg, _ := template.ParseFiles("html/registered.html")
		regmsg.Execute(w, msg)
	}

}

func loggedin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	username := r.Form["username"]
	password := r.Form["password"]
	fmt.Println(username, password)
	user := cleanupBracket(username)
	pswd := cleanupBracket(password)
	loginflag := ValidateUserPass(user, pswd)
	if loginflag == true {
		fmt.Println(loginflag)
		//Load user workout data into this template
		id = getID(user, pswd)
		fmt.Println("User ID: " + id)
		//getUserData(id)
		var workoutids []string
		userworkouts := getUserWorkouts(id)
		for i := 0; i < len(userworkouts); i++ {
			workoutids = append(workoutids, userworkouts[i].WorkoutID)
		}
		test := processworkouts(workoutids)
		//fmt.Println(test)
		var tarray []Tested
		testdict := make(map[string][]Exercise)
		for i := 0; i < len(userworkouts); i++ {
			dateName := userworkouts[i].WorkoutName + "   (" + userworkouts[i].Date + ")"
			obj := Tested{dateName, userworkouts[i].WorkoutID, test[i]}
			tarray = append(tarray, obj)
			testdict[userworkouts[i].WorkoutName] = test[i]

		}
		fmt.Println(tarray)
		login, _ := template.ParseFiles("html/dash.html")
		login.Execute(w, tarray)
	} else {
		fmt.Println(loginflag)
	}
}

func displayExercises(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/displayExercises", r.Method)
	if r.Method == "GET" {
		wid = r.FormValue("WorkoutID")
		fmt.Println("WorkoutID: " + wid)
		exerciselist = getWorkoutData(wid)
		login, _ := template.ParseFiles("html/displayexercises.html")
		login.Execute(w, exerciselist)
	} else {
		r.ParseForm()
		//wid = r.FormValue("WorkoutID")
		//wid = r.FormValue("WorkoutID")
		exercisename := r.FormValue("exercise")
		//fmt.Println(r.FormValue("exercise"))
		//pass in workoutID, ExerciseName, Generate ExerciseID
		//exerciselist := getWorkoutData(wid)
		fmt.Println(wid)
		addExerciseDB(exercisename)
		exerciselist := getWorkoutData(wid)
		login, _ := template.ParseFiles("html/displayexercises.html")
		login.Execute(w, exerciselist)
	}

}

func deleteExercise(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/deleteexercise/", r.Method)
	r.ParseForm()
	exerciseID := r.FormValue("ExerciseID")
	fmt.Println(exerciseID)
	deleteExerciseDB(exerciseID)
	exerciselist := getWorkoutData(wid)
	login, _ := template.ParseFiles("html/displayexercises.html")
	login.Execute(w, exerciselist)
}

func newExerciseForm(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/newexerciseform/", r.Method)
	if r.Method == "GET" {
		r.ParseForm()
		exerciseID = r.FormValue("ExerciseID")
		fmt.Println(exerciseID)
		login, _ := template.ParseFiles("html/exercisenameform.html")
		login.Execute(w, exerciselist)
	} else {
		r.ParseForm()
		newexercisename := r.FormValue("newexercisename")
		fmt.Println(newexercisename)
		editExerciseDB(newexercisename, exerciseID)
		exerciselist := getWorkoutData(wid)
		login, _ := template.ParseFiles("html/displayexercises.html")
		login.Execute(w, exerciselist)
	}

}

func editExercise(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/editexercise/", r.Method)
	r.ParseForm()
	exerciseID := r.FormValue("ExerciseID")
	exerciseN := r.FormValue("newexercisename")
	fmt.Println(exerciseID)
	fmt.Println(exerciseN)
	//create a form to receive new name for for exercise
	//editExerciseDB(exerciseID)
	exerciselist := getWorkoutData(wid)
	login, _ := template.ParseFiles("html/displayexercises.html")
	login.Execute(w, exerciselist)
}

func userdash(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/userdash", r.Method)
	if r.Method == "POST" {
		r.ParseForm()
		workout := r.Form["workout"]
		addWorkout(workout)
		var workoutids []string
		userworkouts := getUserWorkouts(id)
		for i := 0; i < len(userworkouts); i++ {
			workoutids = append(workoutids, userworkouts[i].WorkoutID)
		}
		test := processworkouts(workoutids)
		//fmt.Println(test)
		var tarray []Tested
		testdict := make(map[string][]Exercise)
		for i := 0; i < len(userworkouts); i++ {
			dateName := userworkouts[i].WorkoutName + "   (" + userworkouts[i].Date + ")"
			obj := Tested{dateName, userworkouts[i].WorkoutID, test[i]}
			tarray = append(tarray, obj)
			testdict[userworkouts[i].WorkoutName] = test[i]

		}
		fmt.Println(tarray)
		login, _ := template.ParseFiles("html/dash.html")
		login.Execute(w, tarray)
	} else {
		//get workouts
		userworkouts := getUserWorkouts(id)
		var workoutids []string
		for i := 0; i < len(userworkouts); i++ {
			workoutids = append(workoutids, userworkouts[i].WorkoutID)
		}
		test := processworkouts(workoutids)
		//fmt.Println(test)
		var tarray []Tested
		testdict := make(map[string][]Exercise)
		for i := 0; i < len(userworkouts); i++ {
			dateName := userworkouts[i].WorkoutName + "   (" + userworkouts[i].Date + ")"
			obj := Tested{dateName, userworkouts[i].WorkoutID, test[i]}
			tarray = append(tarray, obj)
			testdict[userworkouts[i].WorkoutName] = test[i]

		}
		fmt.Println(tarray)
		login, _ := template.ParseFiles("html/dash.html")
		login.Execute(w, tarray)
	}
	//Load db stored workouts into dash.html. Then add crud functionailty to them
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/", r.Method)
	login, _ := template.ParseFiles("html/signin.html")
	login.Execute(w, nil)
}

func main() {
	fmt.Println("Server Running")
	mux := http.NewServeMux()
	//r := mux.NewRouter()
	static := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", static))
	mux.HandleFunc("/register", register)
	mux.HandleFunc("/registerauth", registerauth)
	mux.HandleFunc("/", login)
	mux.HandleFunc("/loggedin", loggedin)
	mux.HandleFunc("/userdash", userdash)
	mux.HandleFunc("/createworkout", createworkout)
	mux.HandleFunc("/displayexercises", displayExercises)
	mux.HandleFunc("/deleteexercise/", deleteExercise)
	mux.HandleFunc("/editexercise/", editExercise)
	mux.HandleFunc("/newexerciseform", newExerciseForm)
	//r.HandleFunc("userdashboard", userdash)
	log.Fatal(http.ListenAndServe(":9091", mux))
}
