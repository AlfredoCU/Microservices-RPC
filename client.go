package main

// Import libraries.
import (
	"fmt"
	"net/rpc"
)

// Register -> data struct.
type Register struct {
	Student string
	Subject string
	Grade   float64
}

// Const type and port server.
const (
	ConnType = "tcp"
	ConnPort = "localhost:9999"
)

// Function connection in the server.
func client() {
	// Connection in the server.
	con, err := rpc.Dial(ConnType, ConnPort)

	// Error.
	if err != nil {
		fmt.Println("ERROR_CONNECTING: ", err)
		return
	}

	// Variable option for menu.
	var option int64
	var resultDone string
	var resultAverage float64

	// Instance register.
	var data Register

	for {
		// Menu.
		fmt.Println("\n******* Microservices (RPC) *******")
		fmt.Println("\n1.- Add a student's grade by subject.")
		fmt.Println("2.- Obtain the student's average.")
		fmt.Println("3.- Obtain the average of all students.")
		fmt.Println("4.- Get average by subject")
		fmt.Println("0.- Exit System")
		// Get option.
		fmt.Print("- Option: ")
		_, _ = fmt.Scanln(&option)

		switch option {
			case 1:
				registerData(data, con, resultDone) // Register data.
			case 2:
				studentAvg(data, con, resultAverage) // Student average.
			case 3:
				generalAvg(data, con, resultAverage) // General average.
			case 4:
				subjectAvg(data, con, resultAverage) // Subject average.
			case 0:
				exited() // Exit and terminate process.
				return
			default:
				invalidOptions() // Invalid Options.
			}
	}
}

/* Function registerData
	@param data Register
	@param con *rpc.Client
	@param resultDone string
*/
func registerData(data Register, con *rpc.Client, resultDone string) {
	// Register data.
	fmt.Println("\n******* ADD DATA TO *******")
	fmt.Print("- Add name: ")
	_, _ = fmt.Scanln(&data.Student)
	fmt.Print("- Add subject: ")
	_, _ = fmt.Scanln(&data.Subject)
	fmt.Print("- Add grade: ")
	_, _ = fmt.Scanln(&data.Grade)

	// To call a method exposed by RPC we use the Call() method.
	err := con.Call("Server.AddRegister", data, &resultDone)

	// Error.
	if err != nil {
		fmt.Println("ERROR_SUBMIT_DATA", err)
	}
}

/* Function studentAvg
	@param data Register
	@param con *rpc.Client
	@param resultAverage float64
*/
func studentAvg(data Register, con *rpc.Client, resultAverage float64) {
	// Get name for search.
	fmt.Println("\n******* STUDENT AVG *******")
	fmt.Print("Get Name: ")
	_, _ = fmt.Scanln(&data.Student)

	// To call a method exposed by RPC we use the Call() method.
	err := con.Call("Server.StudentAverage", data, &resultAverage)

	// Error.
	if err != nil {
		fmt.Println("ERROR_GET_DATA", err)
	} else {
		fmt.Println("Average: ", data.Student, " is: ", resultAverage)
	}
}

/* Function generalAvg
	@param data Register
	@param con *rpc.Client
	@param resultAverage float64
*/
func generalAvg(data Register, con *rpc.Client, resultAverage float64) {
	// To call a method exposed by RPC we use the Call() method.
	fmt.Println("\n******* GENERAL AVG *******")
	err := con.Call("Server.GeneralAverage", data, &resultAverage)

	// Error.
	if err != nil {
		fmt.Println("ERROR_GET_DATA", err)
	} else {
		fmt.Println("Average is: ", resultAverage)
	}
}

/* Function subjectAvg
	@param data Register
	@param con *rpc.Client
	@param resultAverage float64
*/
func subjectAvg(data Register, con *rpc.Client, resultAverage float64) {
	// Get subject for search.
	fmt.Println("\n******* SUBJECT AVG *******")
	fmt.Print("Subject: ")
	_, _ = fmt.Scanln(&data.Subject)

	// To call a method exposed by RPC we use the Call() method.
	err := con.Call("Server.SubjectAverage", data, &resultAverage)

	// Error.
	if err != nil {
		fmt.Println("ERROR_GET_DATA", err)
	} else {
		fmt.Println("Average: ", data.Subject, " is: ", resultAverage)
	}
}

// Function invalidOptions.
func invalidOptions() {
	fmt.Print("\n-Invalid Option!\n\n")
}

// Function exited.
func exited() {
	fmt.Println("\n-System exited...")
}

// Function main.
func main() {
	// Start client.
	client()
}