package moira

import (
	"math"
	"time"
)

// Int64ToTime returns time.Time from int64
func Int64ToTime(timeStamp int64) time.Time {
	return time.Unix(timeStamp, 0).UTC()
}

// UseString gets pointer value of string or default string if pointer is nil
func UseString(str *string) string {
	if str == nil {
		return ""
	}
	return *str
}

// UseFloat64 gets pointer value of float64 or default float64 if pointer is nil
func UseFloat64(f *float64) float64 {
	if f == nil {
		return 0
	}
	return *f
}

// IsValidFloat64 checks float64 for Inf and NaN. If it is then float64 is not valid
func IsValidFloat64(val float64) bool {
	if math.IsNaN(val) {
		return false
	}
	if math.IsInf(val, 0) {
		return false
	}
	return true
}

// Subset return whether first is a subset of second
func Subset(first, second []string) bool {
	set := make(map[string]bool)
	for _, value := range second {
		set[value] = true
	}

	for _, value := range first {
		if !set[value] {
			return false
		}
	}

	return true
}

// GetStringListsDiff returns the members of the set resulting from the difference between the first set and all the successive lists.
func GetStringListsDiff(stringLists ...[]string) []string {
	if len(stringLists) == 0 {
		return []string{}
	}
	leftValues := make(map[string]bool)
	for _, value := range stringLists[0] {
		leftValues[value] = true
	}
	for _, stringList := range stringLists[1:] {
		for _, value := range stringList {
			delete(leftValues, value)
		}
	}
	result := make([]string, 0)
	for _, value := range stringLists[0] {
		if _, ok := leftValues[value]; ok {
			result = append(result, value)
		}
	}
	return result
}

// GetTriggerListsDiff returns the members of the set resulting from the difference between the first set and all the successive lists.
func GetTriggerListsDiff(triggerLists ...[]*Trigger) []*Trigger {
	if len(triggerLists) == 0 {
		return []*Trigger{}
	}
	leftValues := make(map[string]bool)
	for _, value := range triggerLists[0] {
		if value != nil {
			leftValues[value.ID] = true
		}
	}
	for _, triggerList := range triggerLists[1:] {
		for _, trigger := range triggerList {
			if trigger != nil {
				delete(leftValues, trigger.ID)
			}
		}
	}
	result := make([]*Trigger, 0)
	for _, value := range triggerLists[0] {
		if value == nil {
			continue
		}
		if _, ok := leftValues[value.ID]; ok {
			result = append(result, value)
		}
	}
	return result
}

// ChunkSlice gets slice of strings and chunks it to a given size. It returns a batch of chunked lists
func ChunkSlice(original []string, chunkSize int) (divided [][]string) {
	if chunkSize < 1 {
		return
	}
	for i := 0; i < len(original); i += chunkSize {
		end := i + chunkSize

		if end > len(original) {
			end = len(original)
		}

		divided = append(divided, original[i:end])
	}
	return
}
