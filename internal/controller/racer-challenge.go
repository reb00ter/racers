package controller

import (
	"encoding/base64"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/reb00ter/racers/internal/context"
	intErrors "github.com/reb00ter/racers/internal/core/errors"
	"github.com/reb00ter/racers/internal/models"
	"hash/fnv"
	"net/http"
	"strings"
	"sync"
	"time"
)

type (
	RacerChallenge struct {
		activeTokens map[string]bool
		tokenMutex   sync.Mutex
	}
	RacerChallengeViewModel struct {
		Racer1 RacerViewModel
		Racer2 RacerViewModel
		Vote1  string
		Vote2  string
		Token  string
	}
)

func NewRacerChallenge() *RacerChallenge {
	ctrl := new(RacerChallenge)
	ctrl.activeTokens = make(map[string]bool)
	ctrl.tokenMutex = sync.Mutex{}
	return ctrl
}
func (ctrl *RacerChallenge) GetChallenge(c echo.Context) error {
	cc := c.(*context.AppContext)

	var racers []models.Racer
	err := cc.RacerStore.Challenge(&racers)

	if err != nil {
		b := intErrors.NewBoom(intErrors.RacerNotFound, intErrors.ErrorText(intErrors.RacerNotFound), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusNotFound, b)
	}
	vote1, err := ctrl.getVoteValue(racers[0].ID, racers[1].ID, c)
	if err != nil {
		b := intErrors.NewBoom(intErrors.InternalError, intErrors.ErrorText(intErrors.InternalError), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, b)
	}
	vote2, err := ctrl.getVoteValue(racers[1].ID, racers[0].ID, c)
	if err != nil {
		b := intErrors.NewBoom(intErrors.InternalError, intErrors.ErrorText(intErrors.InternalError), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, b)
	}
	viewModel := RacerChallengeViewModel{
		Racer1: RacerViewModel{
			Name:   racers[0].Name,
			ID:     racers[0].ID,
			Image:  racers[0].Image,
			Rating: racers[0].Rating,
		},
		Racer2: RacerViewModel{
			Name:   racers[1].Name,
			ID:     racers[1].ID,
			Image:  racers[1].Image,
			Rating: racers[1].Rating,
		},
		Vote1: vote1,
		Vote2: vote2,
		Token: ctrl.createToken(),
	}

	return c.Render(http.StatusOK, "racer-challenge.html", viewModel)
}

func (ctrl *RacerChallenge) Vote(c echo.Context) error {
	cc := c.(*context.AppContext)
	vote := c.FormValue("vote")
	token := c.FormValue("token")
	if vote == "" || token == "" {
		b := intErrors.NewBoom(intErrors.BadRequest, intErrors.ErrorText(intErrors.BadRequest), nil)
		c.Logger().Error("Bad request. Vote: ", vote, " Token: ", token)
		return c.JSON(http.StatusBadRequest, b)
	}
	if !ctrl.tokenActive(token) {
		b := intErrors.NewBoom(intErrors.BadRequest, intErrors.ErrorText(intErrors.BadRequest), nil)
		c.Logger().Error("Inactive Token. Vote: ", vote, " Token: ", token)
		return c.JSON(http.StatusBadRequest, b)
	}
	ctrl.clearToken(token)
	gentleman, boor, err := ctrl.parseVote(vote, c)
	if err != nil {
		b := intErrors.NewBoom(intErrors.BadRequest, intErrors.ErrorText(intErrors.BadRequest), nil)
		c.Logger().Error("Bad vote: ", vote)
		return c.JSON(http.StatusBadRequest, b)
	}
	if err := cc.RacerStore.VoteUp(gentleman); err != nil {
		b := intErrors.NewBoom(intErrors.RacerNotFound, intErrors.ErrorText(intErrors.RacerNotFound), err)
		c.Logger().Error("Failed to vote for gentleman ", gentleman, " ", err)
		return c.JSON(http.StatusNotFound, b)
	}
	if err := cc.RacerStore.VoteDown(boor); err != nil {
		b := intErrors.NewBoom(intErrors.RacerNotFound, intErrors.ErrorText(intErrors.RacerNotFound), err)
		c.Logger().Error("Failed to vote for boor ", boor, " ", err)
		return c.JSON(http.StatusNotFound, b)
	}
	return c.Redirect(http.StatusSeeOther, "/racers/challenge")
}

func (ctrl *RacerChallenge) hash(s string) (string, error) {
	h := fnv.New64a()
	if _, err := h.Write([]byte(s)); err != nil {
		return "", err
	}
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, h.Sum64())
	return base64.StdEncoding.EncodeToString(b), nil
}

func (ctrl *RacerChallenge) getVoteValue(gentleman string, boor string, c echo.Context) (string, error) {
	cc := c.(*context.AppContext)

	data := fmt.Sprintf("%s.%s.%s", gentleman, boor, cc.Config.HashSalt)
	hash, err := ctrl.hash(data)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s.%s.%s", gentleman, boor, hash), nil
}

func (ctrl *RacerChallenge) parseVote(vote string, c echo.Context) (uuid.UUID, uuid.UUID, error) {
	cc := c.(*context.AppContext)

	voteParts := strings.Split(vote, ".")
	if len(voteParts) != 3 {
		return uuid.UUID{}, uuid.UUID{}, errors.New("bad vote parts")
	}
	data := fmt.Sprintf("%s.%s.%s", voteParts[0], voteParts[1], cc.Config.HashSalt)
	hash, err := ctrl.hash(data)
	if err != nil {
		return uuid.UUID{}, uuid.UUID{}, err
	}
	if voteParts[2] != hash {
		return uuid.UUID{}, uuid.UUID{}, errors.New("bad vote hash")
	}
	return uuid.MustParse(voteParts[0]), uuid.MustParse(voteParts[1]), nil
}

func (ctrl *RacerChallenge) createToken() string {
	token := uuid.New().String()
	ctrl.activeTokens[token] = true
	go func() {
		time.Sleep(5 * time.Minute)
		ctrl.clearToken(token)
	}()
	return token
}

func (ctrl *RacerChallenge) clearToken(token string) {
	ctrl.tokenMutex.Lock()
	defer ctrl.tokenMutex.Unlock()
	if _, exist := ctrl.activeTokens[token]; exist {
		delete(ctrl.activeTokens, token)
	}
}

func (ctrl *RacerChallenge) tokenActive(token string) bool {
	ctrl.tokenMutex.Lock()
	defer ctrl.tokenMutex.Unlock()
	_, exist := ctrl.activeTokens[token]
	return exist
}
