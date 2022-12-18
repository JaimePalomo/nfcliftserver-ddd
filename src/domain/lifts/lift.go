package lifts

import (
	"github.com/federicoleon/bookstore_utils-go/rest_errors"
	"strings"
)

// Lift estructura para guardar el objeto que representa a un ascensor
type Lift struct {
	Rae            int    `json:"rae"`
	Stops          int    `json:"stops"`
	Description    string `json:"description"`
	Address        string `json:"address"`
	Company        string `json:"company"`
	AppDescription string `json:"appDescription"`
	StopTexts      string `json:"stopTexts"`
	StopMask       string `json:"stopMask"`
	Distance       int    `json:"distance"`
}

func (l *Lift) Validate() rest_errors.RestErr {
	if l.Stops < 1 {
		return rest_errors.NewBadRequestError("lift must have at least 1 stop")
	}

	l.StopTexts = strings.TrimSpace(l.StopTexts)
	if l.StopTexts == "" {
		return rest_errors.NewBadRequestError("lift must have at least 2 stop text")
	}

	if len(strings.Split(l.StopTexts, ",")) < 2 {
		return rest_errors.NewBadRequestError("lift must have at least 2 stop text")
	}

	l.Description = strings.TrimSpace(l.Description)
	if l.Description == "" {
		return rest_errors.NewBadRequestError("lift must have a description")
	}
	return nil
}

func (l *Lift) IsValidFloor(floor string) bool {
	floors := strings.Split(l.StopTexts, ",")

	flag := false
	for _, text := range floors {
		text = strings.Trim(text, " ")
		if floor == text || floor == "" {
			flag = true
			break
		}
	}
	return flag
}
