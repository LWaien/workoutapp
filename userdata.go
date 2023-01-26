package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//var db *sql.DB

func userdb(query string) {
	db, err := sql.Open("mysql", "root:45732906@tcp(localhost:3306)/Accounts")
	if err != nil {
		fmt.Println("Error connecting to the database")
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Error connecting to the database using .Ping()")
		panic(err.Error())
	}

	q, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer q.Close()
	fmt.Println("Successfully queried the database")
}

func updateExerciseQuery(query string, newName string, exerciseID string) {
	db, err := sql.Open("mysql", "root:45732906@tcp(localhost:3306)/Accounts")
	if err != nil {
		fmt.Println("Error connecting to the database")
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Error connecting to the database using .Ping()")
		panic(err.Error())
	}

	q, err := db.Query(query, newName, exerciseID)
	if err != nil {
		panic(err.Error())
	}
	defer q.Close()
	fmt.Println("Successfully updated exercise name")
}

func deleteExerciseQuery(query string, tobeDeleted string) {
	db, err := sql.Open("mysql", "root:45732906@tcp(localhost:3306)/Accounts")
	if err != nil {
		fmt.Println("Error connecting to the database")
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Error connecting to the database using .Ping()")
		panic(err.Error())
	}

	q, err := db.Query(query, tobeDeleted)
	if err != nil {
		panic(err.Error())
	}
	defer q.Close()
	fmt.Println("Successfully deleted row")
}

func queryWorkouts(query string, id string) []Workout {
	db, err := sql.Open("mysql", "root:45732906@tcp(localhost:3306)/Accounts")
	if err != nil {
		fmt.Println("Error connecting to the database")
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Error connecting to the database using .Ping()")
		panic(err.Error())

	}

	var workouts []Workout
	rows, err := db.Query(query, id)

	for rows.Next() {
		var workout Workout
		err := rows.Scan(&workout.WorkoutID, &workout.WorkoutName, &workout.Date)
		if err != nil {
			log.Fatal(err)
		}
		workouts = append(workouts, workout)
	}
	//fmt.Println(workouts)
	return workouts

}

func queryExercises(query string, wid string) []Exercise {
	db, err := sql.Open("mysql", "root:45732906@tcp(localhost:3306)/Accounts")
	if err != nil {
		fmt.Println("Error connecting to the database")
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Error connecting to the database using .Ping()")
		panic(err.Error())

	}

	var exercises []Exercise
	rows, err := db.Query(query, wid)

	for rows.Next() {
		var exercise Exercise
		err := rows.Scan(&exercise.ExerciseID, &exercise.WorkoutID, &exercise.ExerciseName)
		if err != nil {
			log.Fatal(err)
		}
		exercises = append(exercises, exercise)
	}
	//fmt.Println(exercises)
	return exercises

}

func AddUserData(workoutName string, id string) {
	t := time.Now()
	date := t.Format("01/02/2006")
	query := "INSERT INTO Accounts.UserData (WorkoutName,UserID,Date) VALUES ('" + workoutName + "','" + id + "','" + date + "');"
	userdb(query)
}

func AddWorkoutData(exercise string, wid string) {
	//intVar, err := strconv.Atoi(wid)
	fmt.Println("add exercise: " + exercise)
	query := "INSERT INTO Accounts.WorkoutData (WorkoutID,ExerciseName) VALUES ('" + wid + "','" + exercise + "');"
	userdb(query)
}

func getWorkoutData(wid string) []Exercise {
	query := "SELECT ExerciseID, WorkoutID, ExerciseName FROM Accounts.WorkoutData WHERE WorkoutID = ?"
	exerciselist := queryExercises(query, wid)
	return exerciselist
}

func getUserWorkouts(id string) []Workout {
	query := "SELECT WorkoutID, WorkoutName, Date FROM Accounts.UserData WHERE UserID = ? ORDER BY Date DESC"
	workoutlist := queryWorkouts(query, id)
	return workoutlist

}

func addWorkout(workout []string) {
	work := cleanupBracket(workout) //Add workout name to db schema for UserData
	fmt.Println(work)
	AddUserData(work, id)
}

func addExerciseDB(exercisename string) {
	//exercise := cleanupBracket(exercisename) //Add workout name to db schema for UserData
	//fmt.Println(exercisename)
	AddWorkoutData(exercisename, wid)
}

func deleteExerciseDB(exerciseID string) {
	query := "DELETE FROM Accounts.WorkoutData WHERE ExerciseID = ?"
	//eID := strconv.Itoa(exerciseID)
	deleteExerciseQuery(query, exerciseID)
}

func editExerciseDB(newexercise string, exerciseID string) {
	query := "UPDATE Accounts.WorkoutData SET ExerciseName = ? WHERE ExerciseID = ?"
	//eID := strconv.Itoa(exerciseID)
	updateExerciseQuery(query, newexercise, exerciseID)
}

func processworkouts(workoutids []string) [][]Exercise {
	var dashxerciselist [][]Exercise
	for i := 0; i < len(workoutids); i++ {
		workoutexercises := getWorkoutData(workoutids[i])
		dashxerciselist = append(dashxerciselist, workoutexercises)
	}

	return dashxerciselist
}
