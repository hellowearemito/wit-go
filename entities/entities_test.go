package entities

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAgeOfPersons(t *testing.T) {
	response := `
		{
			"age_of_person": [
				{
					"confidence": 0.93306,
					"value": "5",
					"type": "value"
				}
			]
		}
	`
	entities := Entities{}
	err := json.Unmarshal([]byte(response), &entities)
	assert.Nil(t, err)

	persons, err := entities.AgeOfPersons()
	assert.Nil(t, err)

	assert.Equal(t, 1, len(persons))

	assert.Equal(t, 0.93306, persons[0].Confidence)
	assert.Equal(t, "5", persons[0].Value)
	assert.Equal(t, "value", persons[0].Type)
}
