package main

import (
	"strings"
)

type quizQuestion struct {
	Manufacturer string
	Model        string
	PhotoURL     string
	PageURL      string
	TitleRawData string
}

func matchAircraftTypes(q *questionData) (quizQuestion, bool) {
	var aircraftModelMap map[string][]string
	var question quizQuestion
	var foundMatch bool

	question.PageURL = q.PageURL
	question.PhotoURL = q.PhotoURL
	question.TitleRawData = q.TitleRawData

	aircraftModelMap = map[string][]string{
		"Airbus":     {"A300", "A310", "A318", "A319", "A320", "A321", "A330", "A340", "A350", "A380"},
		"Boeing":     {"717", "727", "737-1", "737-2", "737-3", "737-4", "737-5", "737-6", "737-7", "737-8", "737-9", "737-8 MAX", "747-1", "747-2", "747SP", "747-3", "747-4", "747-8", "757-2", "757-3", "767-2", "767-3", "767-4", "777-2", "777-3", "787-8", "787-9"},
		"Embraer":    {"E170", "E175", "E190", "E195", "ERJ-145", "EMB-120"},
		"Antonov":    {"AN-124", "AN-225"},
		"ATR":        {"ATR-42", "ATR-72"},
		"Bombardier": {"Q400", "CRJ-200", "CRJ-900", "CRJ-700", "Global"},
		"Cessna":     {"152", "172", "182", "Citation"},
		"Fokker":     {"50", "70", "100"},
		"Tupolev":    {"134"},
		"Ilyushin":   {"62"},
	}

	for k, v := range aircraftModelMap {
		if strings.Contains(q.AircraftType, k) {
			question.Manufacturer = k
			for _, n := range v {
				if strings.Contains(q.AircraftType, n) {
					question.Model = n
					foundMatch = true
					return question, foundMatch
				}
			}
		}
	}

	return question, foundMatch
}
