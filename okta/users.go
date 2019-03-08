package okta

import (
	"context"
	"fmt"
	"encoding/json"
)

// UsersService is the service providing access to the Users Resource in the Okta API
type UsersService service

// GetByID fetches a user by ID.
//
// https://developer.okta.com/docs/api/resources/users#get-user-with-id
func (s *UsersService) GetByID(ctx context.Context, id string) (*User, *Response, error) {
	ctx = context.WithValue(ctx, rateLimitCategoryCtxKey, rateLimitUsersGetByIDCategory)
	path := fmt.Sprintf("users/%s", id)

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	userOut := new(User)
	resp, err := s.client.Do(ctx, req, userOut)
	if err != nil {
		return nil, resp, err
	}

	return userOut, resp, nil

}

// UpdateProfileDelta modifies a user profile using partial update semantics.
//
// https://developer.okta.com/docs/api/resources/users#update-user
func (s *UsersService) UpdateProfileDelta(ctx context.Context, id string, userRawProfile *json.RawMessage) (*User, *Response, error) {
	ctx = context.WithValue(ctx, rateLimitCategoryCtxKey, rateLimitUsersCreateUpdateDeleteByIDCategory)
	path := fmt.Sprintf("users/%s", id)

	body := map[string]interface{}{"profile": userRawProfile}

	req, err := s.client.NewRequest("POST", path, body)
	if err != nil {
		return nil, nil, err
	}

	userOut := new(User)
	resp, err := s.client.Do(ctx, req, userOut)
	if err != nil {
		return nil, resp, err
	}

	return userOut, resp, nil

}

