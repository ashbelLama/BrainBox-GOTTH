package service_test

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"testing"

	database "github.com/ashbelLama/brainbox/db"
	"github.com/ashbelLama/brainbox/model"
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
	// Ensure the database connection is initialized
	if database.DB == nil {
		t.Fatal("database connection is nil")
	}
	t.Run("check table contents", func(t *testing.T) {

		results, err := database.DB.Query("SELECT Id, Title, Description FROM bubble")
		if err != nil {
			t.Fatalf("failed to query database: %v", err)
		}
		defer results.Close()

		got := service.CheckBubbleContents(results)
		expected := true

		if got != expected {
			t.Fatalf("\nexpected %v,\ngot %v", expected, got)
			fmt.Println("check table contents: test failed")
		}
		fmt.Println("check table contents: test passed")
	})
	t.Run("check get bubble", func(t *testing.T) {
		sample := model.Bubble{
			Id:          1,
			Title:       "test",
			Description: "test",
		}
		got, err := service.GetBubble()
		if err != nil {
			t.Fatalf("failed to get bubble: %v", err)
		}

		expectedJSON, _ := json.Marshal(sample)
		if !reflect.DeepEqual(expectedJSON, got) {
			t.Fatalf("\nexpected %s,\ngot %v", string(expectedJSON), got)
			fmt.Println("check get bubble: test failed")
		}
		fmt.Println("check get bubble: test passed")
	})
	t.Run("check post bubble", func(t *testing.T) {
		sample := model.Bubble{
			Title:       "go test",
			Description: "go test",
		}
		got := service.PostBubble(sample)
		if !got {
			t.Fatalf("\nexpected %t,\ngot %t", true, got)
			fmt.Println("check seeded contents: test failed")
		}
		fmt.Println("check post bubble: test passed")
	})
}
