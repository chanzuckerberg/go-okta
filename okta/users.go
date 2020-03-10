package okta

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
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

// List fetches all users.
//
// https://developer.okta.com/docs/reference/api/users/#list-all-users
func (s *UsersService) List(ctx context.Context) ([]*User, *Response, error) {
	ctx = context.WithValue(ctx, rateLimitCategoryCtxKey, rateLimitCoreCategory)

	path := fmt.Sprintf("users?limit=%d", 200)
	var acc []*User

	return s.listPaginated(ctx, path, acc)
}

// ListFilter fetches a list of all users who match a given filter.
//
// https://developer.okta.com/docs/reference/api/users/#list-users-with-a-filter
func (s *UsersService) ListFilter(ctx context.Context, filter string) ([]*User, *Response, error) {
	ctx = context.WithValue(ctx, rateLimitCategoryCtxKey, rateLimitGroupsCreateListCategory)

	path := fmt.Sprintf("users?limit=%d&filter=%s", 100, url.QueryEscape(filter))
	var userAcc []*User

	return s.listPaginated(ctx, path, userAcc)
}

func (s *UsersService) listPaginated(ctx context.Context, path string, acc []*User) ([]*User, *Response, error) {
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var users []*User
	resp, err := s.client.Do(ctx, req, &users)
	if err != nil {
		return nil, nil, err
	}
	acc = append(acc, users...)

	if len(resp.Pagination.Next) == 0 {
		return acc, resp, nil
	}

	return s.listPaginated(ctx, resp.Pagination.Next, acc)
}
