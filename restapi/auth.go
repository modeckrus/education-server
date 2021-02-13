package restapi

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

//FirebaseAuthInput ...
type FirebaseAuthInput struct {
	FireToken string `json:"fireToken"`
}
type FirebaseAuthResponse struct {
	IsOk bool `json:"isOk"`
}

//FirebaseAuth ...
func (g *Resolver) FirebaseAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()
	input := &FirebaseAuthInput{}
	json.NewDecoder(r.Body).Decode(input)
	bctx := context.Background()
	ctx, cancel := context.WithTimeout(bctx, time.Minute)
	defer cancel()
	user, err := g.DB.FirebaseAuth(ctx, input.FireToken)
	if err != nil {
		json.NewEncoder(w).Encode(JError{
			Error: err.Error(),
		})
		return
	}
	json.NewEncoder(w).Encode(user)
}

//UpdateUserInput ...
type UpdateUserInput struct {
	Name        string  `json:"name"`
	DisplayName string  `json:"displayName"`
	PhotoID     *string `json:"photoId"`
}

//UpdateUser ...
func (g *Resolver) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()
	input := &UpdateUserInput{}
	json.NewDecoder(r.Body).Decode(input)
	bctx := context.Background()
	ctx, cancel := context.WithTimeout(bctx, time.Minute)
	defer cancel()
	token := r.Header.Get("Token")
	checkedUser, err := g.DB.GetUserByToken(token)
	if err != nil {
		json.NewEncoder(w).Encode(JError{
			Error: err.Error(),
		})
		return
	}
	user, err := g.DB.UpdateUser(ctx, checkedUser, input.Name, input.DisplayName, input.PhotoID)

	if err != nil {
		json.NewEncoder(w).Encode(JError{
			Error: err.Error(),
		})
		return
	}
	json.NewEncoder(w).Encode(user)
}
