package repo

import (
	"testing"

	"github.com/FizzaSadath/flashcard-app-backend/internal/core"
)

func TestCreateAndGetUser(t *testing.T) {
	db := SetupTestDB(t)
	repo := NewUserRepo(db)

	inputUser := &core.User{
		Username: "test_user",
		Email:    "test_user@test.com",
		Password: "hashed_password",
	}

	err := repo.CreateUser(inputUser)
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	if inputUser.ID == 0 {
		t.Error("User ID should not be 0 after saving")
	}

	fetchedUser, err := repo.GetUserByEmail(inputUser.Email)
	if err != nil {
		t.Fatalf("Failed to get user by email: %v", err)
	}

	if fetchedUser.Username != inputUser.Username {
		t.Errorf("Expected username %s, got %s", inputUser.Username, fetchedUser.Username)
	}
	if fetchedUser.Password != inputUser.Password {
		t.Errorf("Expected password %s, got %s", inputUser.Password, fetchedUser.Password)
	}

	fetchedUserByName, err := repo.GetUserByUsername(inputUser.Username)
	if err != nil {
		t.Fatalf("Failed to get user by username: %v", err)
	}

	if fetchedUserByName.Email != inputUser.Email {
		t.Errorf("Expected email %s, got %s", inputUser.Email, fetchedUserByName.Email)
	}
}

func TestGetUserNotFound(t *testing.T) {
	db := SetupTestDB(t)
	repo := NewUserRepo(db)

	_, err := repo.GetUserByEmail("ghost@test.com")

	if err == nil {
		t.Error("Expected error for non-existent user, got nil")
	}

	if err.Error() != "user not found" {
		t.Errorf("Expected 'user not found' error, got '%v'", err)
	}
}
