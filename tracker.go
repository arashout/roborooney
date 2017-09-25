package roborooney

import (
	"fmt"

	"github.com/arashout/mlpapi"
)

func NewTracker() *Tracker {
	return &Tracker{
		pitchSlotMap: make(map[string]PitchSlot),
	}
}

func (tracker *Tracker) Insert(_pitch mlpapi.Pitch, _slot mlpapi.Slot) {
	// Use the Pitch ID and Slot ID to create a unique identifer
	pitchSlotID := calculatePitchSlotId(_pitch.ID, _slot.ID)
	tracker.pitchSlotMap[pitchSlotID] = PitchSlot{
		pitch: _pitch,
		slot:  _slot,
	}
}

func (tracker *Tracker) Retrieve(pitchID, slotID string) PitchSlot {
	pitchSlotID := calculatePitchSlotId(pitchID, slotID)
	return tracker.pitchSlotMap[pitchSlotID]

}
func calculatePitchSlotId(pitchID, slotID string) string {
	return fmt.Sprintf("%s-%s", pitchID, slotID)
}