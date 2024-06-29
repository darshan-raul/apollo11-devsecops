package user

import "testing"

func TestUserServiceHandlers(t *testing.T){
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)
}

type mockUserStore struct {}

