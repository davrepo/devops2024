package main

import (
	"log"
	"net/http"
	"os"

	"database/sql"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	_ "github.com/mattn/go-sqlite3"
)

var (
	store = sessions.NewCookieStore([]byte("secret-key"))
)

// Configurations
const (
	Database   = "/tmp/minitwit.db"
	PerPage    = 30
	Debug      = true
	SecretKey  = "development key"
	SchemaFile = "schema.sql"
)

// Database connection
func connect_db() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", Database)
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}

// Initialize database
func init_DB() error {
	// Connect to Database
	db, err := connect_db()
	if err != nil {
		return err
	}
	defer db.Close()

	// Create schema file
	schemaData, err := os.ReadFile(SchemaFile)
	if err != nil {
		return err
	}

	// Execute schema SQL script
	_, err = db.Exec(string(schemaData))
	if err != nil {
		return err
	}

	return nil
}

// query database and return result as a list of maps
func query_DB(db *sql.DB, query string, args ...interface{}) ([]map[string]interface{}, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}
	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		for i := range values {
			valuePtrs[i] = &values[i]
		}

		if err := rows.Scan(valuePtrs...); err != nil {
			return nil, err
		}

		rowMap := make(map[string]interface{})
		for i, col := range columns {
			rowMap[col] = values[i]
		}
		result = append(result, rowMap)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

// Get user ID for a given username
func get_user_id(db *sql.DB, username string) (int, error) {
	var userID int
	err := db.QueryRow("SELECT user_id FROM user WHERE username = ?", username).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			// User not found
			return 0, nil
		}
		// Other error
		return 0, err
	}
	return userID, nil
}

func main() {
	r := mux.NewRouter()
	// Define routes
	r.HandleFunc("/", timelineHandler).Methods("GET")
	r.HandleFunc("/public", publicTimelineHandler).Methods("GET")
	r.HandleFunc("/{username}", userTimelineHandler).Methods("GET")
	r.HandleFunc("/{username}/follow", followUserHandler).Methods("GET")
	r.HandleFunc("/{username}/unfollow", unfollowUserHandler).Methods("GET")
	r.HandleFunc("/add_message", addMessageHandler).Methods("POST")
	r.HandleFunc("/login", loginHandler).Methods("GET", "POST")
	r.HandleFunc("/register", registerHandler).Methods("GET", "POST")
	r.HandleFunc("/logout", logoutHandler).Methods("GET")

	// Middleware to open and defer closing database connection
	r.Use(dbMiddleware)

	// Start server
	log.Fatal(http.ListenAndServe(":8080", r))
}

func timelineHandler(w http.ResponseWriter, r *http.Request) {
	// Displays the latest messages of all users.
	// TODO

}

func publicTimelineHandler(w http.ResponseWriter, r *http.Request) {
	// Displays the latest messages of all users.
	// TODO
}

func userTimelineHandler(w http.ResponseWriter, r *http.Request) {
	// Display's a users tweets.
	// TODO
	vars := mux.Vars(r)
	username := vars["username"]
	profile_user := query_DB("select * from user where username = ?")

}

func followUserHandler(w http.ResponseWriter, r *http.Request) {
	// Adds the current user as follower of the given user.
	// TODO
}

func unfollowUserHandler(w http.ResponseWriter, r *http.Request) {
	// Removes the current user as follower of the given user.
	// TODO
}

func addMessageHandler(w http.ResponseWriter, r *http.Request) {
	// Registers a new message for the user.
	// TODO
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// Logs the user in.
	// TODO
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	// Registers the user.
	// TODO
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	// Logs the user out.
	// TODO

	// """Logs the user out"""
	// flash('You were logged out')
	// session.pop('user_id', None)
	// return redirect(url_for('public_timeline'))

	session, err := store.Get(r, "")
	if err != nil {
		//...
		return
	}

	session.Values["user_id"] = nil
	//err handlle

}
