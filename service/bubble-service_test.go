package service_test

import (
	"log"
	"os"
	"testing"

	database "github.com/ashbelLama/brainbox/db"
	"github.com/ashbelLama/brainbox/service"
)

func TestMain(m *testing.M) {
	// Initialize the database connection here
	database.InitDB()

	// Verify the connection
	if database.DB == nil {
		log.Fatalf("database connection is nil")
	}

	// Run the tests
	code := m.Run()

	// Close the database connection
	err := database.DB.Close()
	if err != nil {
		log.Fatalf("failed to close the database: %v", err)
	}

	os.Exit(code)
}

func TestDBConnection(t *testing.T) {
	got := service.CheckDBConnection()

	if got != nil {
		t.Fatalf("\nexpected %v,\ngot      %v", nil, got)
	}
}

func TestGetBubble(t *testing.T) {
	t.Run("check table contents", func(t *testing.T) {
		// Ensure the database connection is initialized
		if database.DB == nil {
			t.Fatal("database connection is nil")
		}

		results, err := database.DB.Query("SELECT Id, Title, Description FROM bubble")
		if err != nil {
			t.Fatalf("failed to query database: %v", err)
		}
		defer results.Close()

		got := service.CheckBubbleContents(results)
		expected := true

		if got != expected {
			t.Fatalf("\nexpected %v,\ngot %v", expected, got)
		} else {
			t.Logf("\nexpected %v,\ngot %v", expected, got)
		}
	})
}
