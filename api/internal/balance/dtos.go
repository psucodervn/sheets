package balance

import (
	"time"

	"github.com/volatiletech/null/v8"

	"api/model"
)

type ChangeDTO struct {
	ID      string  `json:"id"`
	Value   float64 `json:"value"`
	Name    string  `json:"name"`
	Percent float64 `json:"percent,omitempty"`
}

type ChangesDTO []ChangeDTO

type TransactionDTO struct {
	ID           string          `json:"id,omitempty"`
	CreatorID    null.String     `json:"creatorID,omitempty"`
	Time         time.Time       `json:"time"`
	Value        float64         `json:"value"`
	Summary      string          `json:"summary"`
	Description  null.String     `json:"description,omitempty"`
	Payers       ChangesDTO      `json:"payers"`
	Participants ChangesDTO      `json:"participants"`
	SplitType    model.SplitType `json:"splitType"`
}

func mapChangesDTOtoModelChanges(chs ChangesDTO) model.Changes {
	cs := make(model.Changes, len(chs))
	for i, c := range chs {
		cs[i].ID = c.ID
		cs[i].Value = c.Value
		cs[i].Percent = c.Percent
	}
	return cs
}

func mapTransactionDTOtoModelTransaction(tx *TransactionDTO) *model.Transaction {
	t := &model.Transaction{
		ID:           tx.ID,
		CreatorID:    tx.CreatorID,
		Time:         tx.Time,
		Value:        tx.Value,
		Summary:      tx.Summary,
		Description:  tx.Description,
		Payers:       mapChangesDTOtoModelChanges(tx.Payers),
		Participants: mapChangesDTOtoModelChanges(tx.Participants),
		SplitType:    tx.SplitType,
	}
	return t
}
