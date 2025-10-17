package stock

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	httpHelper "stock-pwt/internal/delivery/http"
	"stock-pwt/internal/entity/stock"
	"stock-pwt/pkg/response"
)

func (h *Handler) GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	var (
		result   interface{}
		metadata interface{}
		err      error
		resp     response.Response
	)
	defer resp.RenderJSON(w, r)

	ctx := r.Context()

	result, err = h.stockSvc.GetUserByUsername(ctx, r.FormValue("username"))
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR][%s][%s] %s | Reason: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
		return
	}

	resp.Data = result
	resp.Metadata = metadata

	log.Printf("[INFO][%s][%s] %s", r.RemoteAddr, r.Method, r.URL)
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{}
	defer resp.RenderJSON(w, r)

	ctx := r.Context()

	User := stock.User{}

	err := json.NewDecoder(r.Body).Decode(&User)
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR][%s][%s] %s | Reason: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
		return
	}

	err = h.stockSvc.CreateUser(ctx, User)
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR][%s][%s] %s | Reason: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
		return
	}

	log.Printf("[INFO][%s][%s] %s", r.RemoteAddr, r.Method, r.URL)
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{}
	defer resp.RenderJSON(w, r)

	ctx := r.Context()

	user := stock.User{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR][%s][%s] %s | Reason: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
		return
	}

	err = h.stockSvc.UpdateUser(ctx, user)
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR][%s][%s] %s | Reason: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
		return
	}

	log.Printf("[INFO][%s][%s] %s", r.RemoteAddr, r.Method, r.URL)
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{}
	defer resp.RenderJSON(w, r)

	ctx := r.Context()

	err := h.stockSvc.DeleteUser(ctx, r.FormValue("username"))
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR][%s][%s] %s | Reason: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
		return
	}

	log.Printf("[INFO][%s][%s] %s", r.RemoteAddr, r.Method, r.URL)
}

// func (h *Handler) MatchPassword(w http.ResponseWriter, r *http.Request) {
// 	resp := response.Response{}
// 	defer resp.RenderJSON(w, r)

// 	ctx := r.Context()

// 	// Decode the incoming JSON body to get the User struct
// 	User := stock.User{}
// 	err := json.NewDecoder(r.Body).Decode(&User)
// 	if err != nil {
// 		resp.SetError(err) // Set error if JSON decoding fails
// 		log.Printf("[ERROR][%s][%s] %s | Reason: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
// 		return
// 	}

// 	// Check if the password matches
// 	err = h.stockSvc.MatchPassword(ctx, User)
// 	if err != nil {
// 		resp.SetError(err) // Set error if password match fails
// 		log.Printf("[ERROR][%s][%s] %s | Reason: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
// 		return
// 	}

// 	// If password matches, respond with success
// 	resp.StatusCode = http.StatusOK
// 	resp.Error = response.Error{
// 		Status: false,              // No error
// 		Msg:    "Login successful", // Success message
// 		Code:   0,                  // No application-level error code
// 	}
// 	log.Printf("[INFO][%s][%s] %s | User %s logged in successfully", r.RemoteAddr, r.Method, r.URL, User.Username)
// }

func (h *Handler) MatchPassword(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{}
	defer resp.RenderJSON(w, r)

	ctx := r.Context()

	// Decode the JSON body
	var user stock.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		resp.SetError(errors.New("Invalid request body"))
		log.Printf("[ERROR][%s][%s] %s | Decode error: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
		return
	}

	// Call the login service (this will check password + return JWT if valid)
	token, err := h.stockSvc.Login(ctx, user.Username, user.Password)
	if err != nil {
		resp.SetError(err)
		log.Printf("[ERROR][%s][%s] %s | Login failed: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
		return
	}

	// Set success response with token
	resp.StatusCode = http.StatusOK
	resp.Data = map[string]interface{}{
		"token": token,
	}
	resp.Error = response.Error{
		Status: false,
		Msg:    "Login berhasil",
		Code:   0,
	}

	log.Printf("[INFO][%s][%s] %s | User %s logged in successfully", r.RemoteAddr, r.Method, r.URL, user.Username)
}
