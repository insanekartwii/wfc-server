package api

import (
	"errors"
	"net/http"
	"wwfc/database"
)


type LinkRequest struct {
		Secret    string `json:"secret"`
		ProfileID uint32 `json:"pid"`
		DiscordID string `json:"discordId"`
		Action    string `json:"action"`
}

var LinkRoute = MakeRouteSpec[LinkRequest, UserActionResponse](
	true,
	"/api/link",
	func(req any, v bool, _ *http.Request) (any, int, error) {
		return handleUserAction(req.(LinkRequest), v, HandleLink)
	},
	http.MethodPost,
)

var (
	ErrPIDIsMissing = errors.New("Profile ID missing or invalid")
	ErrDiscordIDMissing = errors.New("Discord ID missing or invalid")
	ErrDiscordLinked	    = errors.New("Discord ID already linked to another profile")
	ErrDiscordWrongStep	    = errors.New("Profile is not in the correct step to link Discord ID")
	ErrDiscordCannotUnlink  = errors.New("Discord ID cannot be unlinked")
	ErrActionMissing        = errors.New("Action missing in request")
)

func HandleLink(req LinkRequest, _ bool) (*database.User, int, error) {
	var err error

	if req.ProfileID == 0 {
		return nil, http.StatusBadRequest, ErrPIDIsMissing
	}

	if  req.DiscordID == "" {
		return nil, http.StatusBadRequest, ErrDiscordIDMissing
	}

	if req.Action == "" || (req.Action != "link" && req.Action != "check" && req.Action != "unlink") {
		return nil, http.StatusBadRequest, ErrActionMissing
	}

	user, success := database.GetProfile(pool, ctx, req.ProfileID)
	if success != nil {
		return nil, http.StatusInternalServerError, ErrUserQuery
	}

	if req.Action == "link" {
		if user.DiscordID != "" {
			return nil, http.StatusForbidden, ErrDiscordLinked
		}
		err = user.UpdateDiscordID(pool, ctx, "1")
		if err != nil {
			return nil, http.StatusInternalServerError, err
		}
	} else if req.Action == "check" {
		if user.DiscordID != "2" {
			return nil, http.StatusForbidden, ErrDiscordWrongStep
		}
		err = user.UpdateDiscordID(pool, ctx, req.DiscordID)
		if err != nil {
			return nil, http.StatusInternalServerError, err
		}
	} else if req.Action == "unlink" {
		if user.DiscordID != "1" && user.DiscordID != "2" && user.DiscordID != req.DiscordID {
			return nil, http.StatusForbidden, ErrDiscordCannotUnlink
		}
		err = user.UpdateDiscordID(pool, ctx, "")
		if err != nil {
			return nil, http.StatusInternalServerError, err
		}
	}

	return &user, http.StatusOK, nil
}