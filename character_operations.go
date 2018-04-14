package marvel

import (
	"strconv"
	"github.com/dghubble/sling"
)

type CharacterInfo struct {
	slg *sling.Sling
}

// NewCharacterInfo returns new CharacterInfo
func NewCharacterInfo(sling *sling.Sling) *CharacterInfo {
	return &CharacterInfo{
		slg: sling.Path("characters/"),
	}
}

// ListAll method returns all characters
func (chInfo *CharacterInfo) ListAll() ([]Character, error) {
	charWrap := &CharacterDataWrapper{}
	_, err := getRequest(chInfo.slg, "../characters", charWrap)
	if err != nil {
		return nil, err
	}

	return charWrap.Data.Results, err
}

// GetById returns the character for given characterID
func (chInfo *CharacterInfo) GetById(characterID int) (*Character, error) {
	charWrap := &CharacterDataWrapper{}
	_, err := getRequest(chInfo.slg, strconv.Itoa(characterID), charWrap)
	if err != nil {
		return nil, err
	}

	return &charWrap.Data.Results[0], nil
}