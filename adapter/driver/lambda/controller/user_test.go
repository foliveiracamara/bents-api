package controller

// import (
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/labstack/echo"
// 	"github.com/stretchr/testify/assert"
// )

// func TestUserController_GetUser(t *testing.T) {
// 	// Create a new echo context for testing
// 	e := echo.New()
// 	req := httptest.NewRequest(http.MethodGet, "/users/123", nil)
// 	rec := httptest.NewRecorder()
// 	ctx := e.NewContext(req, rec)

// 	// Create a mock user service
// 	mockUserService := &MockUserService{}
// 	mockUser := &User{
// 		ID:   "123",
// 		Name: "John Doe",
// 		Age:  30,
// 	}
// 	mockUserService.On("GetUser", "123").Return(mockUser, nil)

// 	// Create a new user controller instance
// 	uc := &UserController{
// 		UserService: mockUserService,
// 		Model:       &UserModel{},
// 	}

// 	// Call the GetUser method
// 	err := uc.GetUser(ctx)

// 	// Assert the response
// 	assert.NoError(t, err)
// 	assert.Equal(t, http.StatusOK, rec.Code)

// 	expectedResponse := &UserResponse{
// 		ID:   "123",
// 		Name: "John Doe",
// 		Age:  30,
// 	}
// 	var actualResponse UserResponse
// 	err = json.Unmarshal(rec.Body.Bytes(), &actualResponse)
// 	assert.NoError(t, err)
// 	assert.Equal(t, expectedResponse, &actualResponse)

// 	// Assert that the GetUser method was called with the correct UUID
// 	mockUserService.AssertCalled(t, "GetUser", "123")
// }
