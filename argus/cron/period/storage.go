package period

import "time"

var periodStorage = make(map[string]time.Time)

func Set(id string) {
	periodStorage[id] = time.Now()
}

func ShouldBeChecked(id string) bool {
	recorded, exists := periodStorage[id]
	if !exists {
		return true
	}

	// 1 hour after the recorded time is "after"(in the past) now?
	return time.Now().After(recorded.Add(time.Hour))
}
