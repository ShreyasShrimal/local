package Pack

import (
	"database/sql"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

var Message string

type User struct {
	Quoteid       int
	PassengerName string
	Adhaar        string
	Contact       string
	FromAddress   string
	ToAddress     string
	TicketPrice   int
	//TicketDate    string
}
type Tolocation struct {
	Address string
	//	ToAddress   string
}

// type Vamp struct {
// 	va string

// 	//Distance int

// }

func Dbconn() *sql.DB {

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/train?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

//var a float64

//type distance float64

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("form/*"))
	//tmpl.ExecuteTemplate(w, "abc", nil)
}

var M int
var D float64
var ur User
var db *sql.DB

func Insert(w http.ResponseWriter, r *http.Request) {
	db = Dbconn()

	if r.Method != http.MethodPost {
		tmpl.ExecuteTemplate(w, "abc", nil)
		return
	}

	if r.Method == "POST" {
		ur = User{
			PassengerName: r.FormValue("PassengerName"),
			//println("Passanger name: ", PassengerName)

			Adhaar:      r.FormValue("Adhaar"),
			Contact:     r.FormValue("Contact"),
			FromAddress: r.FormValue("FromAddress"),

			//
			//	fmt.Println(m)
			ToAddress: r.FormValue("ToAddress"),
			//	TicketDate: r.FormValue("TicketDate"),
			//
		}

		m := ur.FromAddress
		n := ur.ToAddress

		log.Println("INSERT: 	PassengerName: " + ur.PassengerName + "| Adhaar: " + ur.Adhaar + " | Contact: " + ur.Contact + " | FromAddress: " + ur.FromAddress + "| ToAddress: " + ur.ToAddress)

		l1 := Tolocation{"Borivali"}
		l2 := Tolocation{"Goregoan"}
		l3 := Tolocation{"Andheri"}
		l4 := Tolocation{"Bandra"}
		l5 := Tolocation{"Dadar"}
		l6 := Tolocation{"Cst"}
		//v := (map[string]int)
		mp := map[Tolocation]float64{l1: 0,
			l2: 8, l3: 13, l4: 18,
			l5: 21, l6: 25}

		fmt.Println(mp)
		for key, value := range mp {
			//	fmt.Printlno")
			if m == key.Address {
				//fmt.Println(value)
				V1 := value
				fmt.Println("Enter ending point")
				// fmt.Scan", &lo.Address)
				// n := lo.Address
				for key, value := range mp {
					if n == key.Address {
						//	fmt.Println(value)
						V2 := value
						s := (V1 - V2)
						D = math.Abs(s)

						//	ch <- d
						a := 5

						b := 10
						c := 15
						e := 20
						f := 25
						g := 30
						// fmt.Fprintf(w, "  Passenger name  :%s", ur.PassengerName)
						// fmt.Fprintf(w, "  Contact  :%s", ur.Contact)
						// fmt.Fprintf(w, "  Travel From :%s", ur.FromAddress)
						// fmt.Fprintf(w, "  Travel to :%s", ur.ToAddress)
						// fmt.Printf("distance to be travel %v ", D)

						switch {
						case D == 0:
							M = 0
						//	fmt.Fprintf(w, " <h1> no ticket is genrated </h1>")

						case D < 4:
							M = a
							//fmt.Fprintf(w, "<h1> TicKet is %v rs </h1>", a)
						case D <= 8:
							M = b
							//fmt.Fprintf(w, " <h1> TicKet is %v rs </h1>", b)
						case D < 12:
							M = c
							//fmt.Fprintf(w, "<h1> TicKet is %v rs </h1>", c)
						case D < 16:
							M = e
							//fmt.Fprintf(w, " <h1> TicKet is %v rs </h1>", e)
						case D < 20:
							M = f
						//	fmt.Fprintf(w, "<h1> TicKet is %v rs</h1>", f)
						case D < 30:
							M = g
							//fmt.Fprintf(w, "<h1> TicKet is %v rs </h1>", g)
						default:
							//fmt.Fprintf(w, "Please select valid location")

							//tmpl.ExecuteTemplate(w, "xyz", nil)
						}
						fmt.Println("Value of M:", M)

					}
				}

			}
		}
		//
		if r.FormValue("submit") == "insert" {
			//tmpl.ExecuteTemplate(w, "abc", nil)
			Message = "Your ticket price is :" + string(M)
			tmpl.ExecuteTemplate(w, "abc", "hello")
			fmt.Println("Your ticket price is :")
			// tmpl.Execute(w, struct {
			// 	Success bool
			// 	Message string
			// 	vl      int
			// }{Success: true, Message: "Your ticket price is :", vl: M})
		} else {
			tmpl.Execute(w, struct {
				Success bool
				Message string
				vl      int
			}{Success: true, Message: "not inserted"})

		}

	}

}

//fmt.Println(d)
//	http.Redirect(w, r, "/getPost", 301)

// func aaa(aa int) {
// 	fmt.Printf("value  %v", aa)

// }
func GetPost(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "xyz", nil)
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}
}

func Post(w http.ResponseWriter, r *http.Request) {
	db = Dbconn()
	//M : ur.TicketPrice
	fmt.Println("Value of M:", M)
	if r.Method == "POST" {
		// 	tmpl.ExecuteTemplate(w, "xyz", nil)
		// 	return
		// }
		//	ch := make(chan float64)
		fmt.Println("Value of M:", M)
		fmt.Println("value of d:", D)

		// z1 := Vamp{"a"}
		// z2 := Vamp{"b"}
		// z3 := Vamp{"c"}
		// z4 := Vamp{"e"}
		// z5 := Vamp{"f"}
		// z6 := Vamp{"g"}
		// m1 := map[Vamp]int{z1: 5,
		// 	z2: 10, z3: 15, z4: 20,
		// 	z5: 25, z6: 30}
		// fmt.Println(m1)

		if r.Method == "POST" {
			//var Payment string
			S := r.FormValue("Payment")
			fmt.Println(S)
			var P int
			P, _ = strconv.Atoi(S)

			fmt.Println("value of p :", P)

			//for key, value := range m1 {
			//	fmt.Println(key)
			if P == M {
				fmt.Fprintf(w, " ticket is  genrated of  %v rs", P)
			} else {
				fmt.Fprintf(w, " Payment Failed !!!! Please Enter Valid Amount")
			}
			fmt.Fprintf(w, "  Passenger name  :%s,/n", ur.PassengerName)
			fmt.Fprintf(w, "  Contact  :%s,/n", ur.Contact)
			fmt.Fprintf(w, "  Adhaar  :%s", ur.Adhaar)
			fmt.Fprintf(w, "  Travel From :%s", ur.FromAddress)
			fmt.Fprintf(w, "  Travel to :%s", ur.ToAddress)
			fmt.Printf("distance to be travel %v ", D)
			//var db *sql.DB
			db = Dbconn()
			ur.TicketPrice = M
			insForm, err := db.Prepare("call Insertlocal(?,?,?,?,?,?)")
			if err != nil {

				fmt.Println("Someting went wrong!!")
				panic(err.Error())

			}
			if insForm == nil {
				fmt.Println("Data not inserted ")

			}
			if insForm != nil {
				fmt.Println("Data inserted Sussesfully!!")

			}
			// if r.FormValue("submit") == "Read" {
			// 	data := []string{}
			// 	rows, err := db.Query("select * from local")
			// 	if err != nil {
			// 		tmpl.Execute(w, struct {
			// 			Success bool
			// 			Message string
			// 		}{Success: true, Message: err.Error()})
			// 	} else {
			// 		U := User{}
			// 		data = append(data, "<table border=1>")
			// 		data = append(data, "<tr><th>PassengerName</tr></th><tr><th>Adhaar</tr></th><tr><th>Contact</tr></th><tr><th>FromAddress</tr></th><tr><th>ToAddress</tr></th><tr><th>TicketPrice</tr></th>")
			// 		for rows.Next() {
			// 			rows.Scan(&U.PassengerName, &U.Adhaar, &U.Contact, &U.FromAddress, &U.ToAddress, &U.TicketPrice)
			// 			//TicketPrice= strconv.S(U.TicketPrice)
			// 			data = append(data, fmt.Sprintf("<tr><td>%s</td>,<td>%s</td>,<td>%s</td>,<td>%s</td>,<td>%s</td>,<td>%s</td>,<td>%s</td>,", U.PassengerName, U.Adhaar, U.Contact, U.FromAddress, U.ToAddress, U.TicketPrice))
			// 		}
			// 		data = append(data, "</table>")
			// 		tmpl.Execute(w, struct {
			// 			Success bool
			// 			Message string
			// 		}{Success: true, Message: fmt.Sprint(data)})

			// 	}

			insForm.Exec(ur.PassengerName, ur.Adhaar, ur.Contact, ur.FromAddress, ur.ToAddress, ur.TicketPrice)
			//_, err := db.Exec("Insert into local(PassengerName,Adhaar,Contact,FromAddress,ToAddress,TicketPrice,TicketDate)values(?,?,?,?,?,?,?", ur.PassengerName, ur.Adhaar, ur.Contact, ur.FromAddress, ur.ToAddress, ur.TicketPrice, ur.TicketDate)
			if err != nil {
				fmt.Println(err)
			}
			log.Println("INSERT: 	PassengerName: " + ur.PassengerName + "| Adhaar: " + ur.Adhaar + " | Contact: " + ur.Contact + " | FromAddress: " + ur.FromAddress + "| ToAddress: " + ur.ToAddress)

			//l1 := Tolocation{"Borivali"}

		}

	}

}
