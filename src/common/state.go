/* Синглетон по голандски */
package common

import (
	"structs"
	"sync"
)

var instanceState *structs.State
var onceState sync.Once

func GetStateGame() *structs.State {
	onceState.Do(func() {
		instanceState = &structs.State{
			GameID:     0,
			GameKindID: 0,
			GameName:   "",
			Team1:      0,
			Team2:      0,
			CurEvent:   0,
		}
	})
	return instanceState
}
