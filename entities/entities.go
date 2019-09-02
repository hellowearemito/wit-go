package entities

import (
	"encoding/json"
)

// Entities represents the entities
type Entities map[string]json.RawMessage

// GetValue returns the interface{} from the map by the key
func (e Entities) GetValue(key string) json.RawMessage {
	if val, ok := e[key]; ok {
		return val
	}
	return nil
}

// AgeOfPersons returns the entities of age of persons
func (e Entities) AgeOfPersons() ([]AgeOfPerson, error) {
	var items []AgeOfPerson

	value := e.GetValue(AgeOfPersonType)
	if value == nil {
		return items, nil
	}

	err := json.Unmarshal(value, &items)
	if err != nil {
		return items, err
	}

	return items, nil
}

// AgendaEntries returns the entities of agenda entries
func (e Entities) AgendaEntries() ([]AgendaEntry, error) {
	var items []AgendaEntry

	value := e.GetValue(AgendaEntryType)
	if value == nil {
		return items, nil
	}

	err := json.Unmarshal(value, &items)
	if err != nil {
		return items, err
	}
	return items, nil
}

// AmountOfMonies returns the entities of amount of money
func (e Entities) AmountOfMonies() ([]AmountOfMoney, error) {
	var items []AmountOfMoney

	value := e.GetValue(AmountOfMoneyType)
	if value == nil {
		return items, nil
	}

	err := json.Unmarshal(value, &items)
	if err != nil {
		return items, err
	}
	return items, nil
}

// Contacts returns the entities of contact
func (e Entities) Contacts() ([]Contact, error) {
	var items []Contact

	value := e.GetValue(ContactType)
	if value == nil {
		return items, nil
	}

	err := json.Unmarshal(value, &items)
	if err != nil {
		return items, err
	}
	return items, nil
}

// CreativeWorks returns the entities of creative work
func (e Entities) CreativeWorks() ([]CreativeWork, error) {
	var items []CreativeWork

	value := e.GetValue(CreativeWorkTye)
	if value == nil {
		return items, nil
	}

	err := json.Unmarshal(value, &items)
	if err != nil {
		return items, err
	}
	return items, nil
}

// Datetime returns the entities of datetime
func (e Entities) Datetime() ([]Datetime, error) {
	var items []Datetime

	value := e.GetValue(DatetimeType)
	if value == nil {
		return items, nil
	}

	err := json.Unmarshal(value, &items)
	if err != nil {
		return items, err
	}
	return items, nil
}

// Distances returns the entities of distances
func (e Entities) Distances() ([]Distance, error) {
	var items []Distance

	value := e.GetValue(DistanceType)
	if value == nil {
		return items, nil
	}

	err := json.Unmarshal(value, &items)
	if err != nil {
		return items, err
	}
	return items, nil
}

// Durations returns the entities of durations
func (e Entities) Durations() ([]Duration, error) {
	var items []Duration

	value := e.GetValue(DurationType)
	if value == nil {
		return items, nil
	}

	err := json.Unmarshal(value, &items)
	if err != nil {
		return items, err
	}
	return items, nil
}

// Emails returns the entities of emails
func (e Entities) Emails() ([]Email, error) {
	var items []Email

	value := e.GetValue(EmailType)
	if value == nil {
		return items, nil
	}

	err := json.Unmarshal(value, &items)
	if err != nil {
		return items, err
	}
	return items, nil
}

// Byes returns the entities of byes
func (e Entities) Byes() ([]Bye, error) {
	var items []Bye

	value := e.GetValue(ByeType)
	if value == nil {
		return items, nil
	}

	err := json.Unmarshal(value, &items)
	if err != nil {
		return items, err
	}
	return items, nil
}

// Greetings returns the entities of greetings
func (e Entities) Greetings() ([]Greeting, error) {
	var items []Greeting

	value := e.GetValue(GreetingsType)
	if value == nil {
		return items, nil
	}

	err := json.Unmarshal(value, &items)
	if err != nil {
		return items, err
	}
	return items, nil
}

// Intents returns the entities of intents
func (e Entities) Intents() ([]Intent, error) {
	var items []Intent

	value := e.GetValue(IntentType)
	if value == nil {
		return items, nil
	}

	err := json.Unmarshal(value, &items)
	if err != nil {
		return items, err
	}
	return items, nil
}

// LocalSearchQueries returns the entities of local search query
func (e Entities) LocalSearchQueries() ([]LocalSearchQuery, error) {
	var items []LocalSearchQuery

	value := e.GetValue(LocalSearchQueryType)
	if value == nil {
		return items, nil
	}

	err := json.Unmarshal(value, &items)
	if err != nil {
		return items, err
	}
	return items, nil
}

// Locations returns the entities of locations
func (e Entities) Locations() ([]Location, error) {
	var items []Location

	value := e.GetValue(LocationType)
	if value == nil {
		return items, nil
	}

	err := json.Unmarshal(value, &items)
	if err != nil {
		return items, err
	}
	return items, nil
}

// MathExpressions returns the entities of math expressions
func (e Entities) MathExpressions() ([]MathExpression, error) {
	var items []MathExpression

	value := e.GetValue(MathExpressionType)
	if value == nil {
		return items, nil
	}

	err := json.Unmarshal(value, &items)
	if err != nil {
		return items, err
	}
	return items, nil
}

// MessageBodies returns the entities of message bodies
func (e Entities) MessageBodies() ([]MessageBody, error) {
	var items []MessageBody

	value := e.GetValue(MessageBodyType)
	if value == nil {
		return items, nil
	}

	err := json.Unmarshal(value, &items)
	if err != nil {
		return items, err
	}
	return items, nil
}

// NotablePeople returns the entities of notable people
func (e Entities) NotablePersons() ([]NotablePerson, error) {
	var items []NotablePerson

	value := e.GetValue(NotablePersonType)
	if value == nil {
		return items, nil
	}

	err := json.Unmarshal(value, &items)
	if err != nil {
		return items, err
	}
	return items, nil
}

// Number returns the entities of numbers
func (e Entities) Numbers() ([]Number, error) {
	var items []Number

	value := e.GetValue(NumberType)
	if value == nil {
		return items, nil
	}

	err := json.Unmarshal(value, &items)
	if err != nil {
		return items, err
	}
	return items, nil
}

// OnOffs returns the entities of on-off
func (e Entities) OnOffs() ([]OnOff, error) {
	var items []OnOff

	value := e.GetValue(OnOffType)
	if value == nil {
		return items, nil
	}

	err := json.Unmarshal(value, &items)
	if err != nil {
		return items, err
	}
	return items, nil
}

// Ordinals returns the entities of ordinals
func (e Entities) Ordinals() ([]Ordinal, error) {
	var items []Ordinal

	value := e.GetValue(OrdinalType)
	if value == nil {
		return items, nil
	}

	err := json.Unmarshal(value, &items)
	if err != nil {
		return items, err
	}
	return items, nil
}

// PhoneNumbers returns the entities of phone numbers
func (e Entities) PhoneNumbers() ([]PhoneNumber, error) {
	var items []PhoneNumber

	value := e.GetValue(PhoneNumberType)
	if value == nil {
		return items, nil
	}

	err := json.Unmarshal(value, &items)
	if err != nil {
		return items, err
	}
	return items, nil
}

// PhraseToTranslates returns the entities of phrase to translates
func (e Entities) PhraseToTranslates() ([]PhraseToTranslate, error) {
	var items []PhraseToTranslate

	value := e.GetValue(PhraseToTranslateType)
	if value == nil {
		return items, nil
	}

	err := json.Unmarshal(value, &items)
	if err != nil {
		return items, err
	}
	return items, nil
}

// Quantities returns the entities of quantities
func (e Entities) Quantities() ([]Quantity, error) {
	var items []Quantity

	value := e.GetValue(QuantityType)
	if value == nil {
		return items, nil
	}

	err := json.Unmarshal(value, &items)
	if err != nil {
		return items, err
	}
	return items, nil
}

// Reminders returns the entities of reminders
func (e Entities) Reminders() ([]Reminder, error) {
	var items []Reminder

	value := e.GetValue(ReminderType)
	if value == nil {
		return items, nil
	}

	err := json.Unmarshal(value, &items)
	if err != nil {
		return items, err
	}
	return items, nil
}

// SearchQueries returns the entities of search queries
func (e Entities) SearchQueries() ([]SearchQuery, error) {
	var items []SearchQuery

	value := e.GetValue(SearchQueryType)
	if value == nil {
		return items, nil
	}

	err := json.Unmarshal(value, &items)
	if err != nil {
		return items, err
	}
	return items, nil
}

// Sentiments returns the entities of sentiments
func (e Entities) Sentiments() ([]Sentiment, error) {
	var items []Sentiment

	value := e.GetValue(SentimentType)
	if value == nil {
		return items, nil
	}

	err := json.Unmarshal(value, &items)
	if err != nil {
		return items, err
	}
	return items, nil
}

// Temperatures returns the entities of temperatures
func (e Entities) Temperatures() ([]Temperature, error) {
	var items []Temperature

	value := e.GetValue(TemperatureType)
	if value == nil {
		return items, nil
	}

	err := json.Unmarshal(value, &items)
	if err != nil {
		return items, err
	}
	return items, nil
}

// Thanks returns the entities of thanks
func (e Entities) Thanks() ([]Thank, error) {
	var items []Thank

	value := e.GetValue(ThanksType)
	if value == nil {
		return items, nil
	}

	err := json.Unmarshal(value, &items)
	if err != nil {
		return items, err
	}
	return items, nil
}

// Urls returns the entities of urls
func (e Entities) Urls() ([]URL, error) {
	var items []URL

	value := e.GetValue(UrlType)
	if value == nil {
		return items, nil
	}

	err := json.Unmarshal(value, &items)
	if err != nil {
		return items, err
	}
	return items, nil
}

// Volumes returns the entities of volumes
func (e Entities) Volumes() ([]Volume, error) {
	var items []Volume

	value := e.GetValue(VolumeType)
	if value == nil {
		return items, nil
	}

	err := json.Unmarshal(value, &items)
	if err != nil {
		return items, err
	}
	return items, nil
}

// WikipediaSearchQueries returns the entities of wikipedia search queries
func (e Entities) WikipediaSearchQueries() ([]WikipediaSearchQuery, error) {
	var items []WikipediaSearchQuery

	value := e.GetValue(WikipediaSearchQueryType)
	if value == nil {
		return items, nil
	}

	err := json.Unmarshal(value, &items)
	if err != nil {
		return items, err
	}
	return items, nil
}

// WolframSearchQueries returns the entities of wolfram search queries
func (e Entities) WolframSearchQueries() ([]WolframSearchQuery, error) {
	var items []WolframSearchQuery

	value := e.GetValue(WolframSearchQueryType)
	if value == nil {
		return items, nil
	}

	err := json.Unmarshal(value, &items)
	if err != nil {
		return items, err
	}
	return items, nil
}

// Roles returns the role entities
func (e Entities) Roles() (map[string][]Entity, error) {
	entities := map[string][]Entity{}

	for key, value := range e {
		if !e.isRole(key) {
			continue
		}

		entity := Entity{}
		err := json.Unmarshal(value, &entity)
		if err != nil {
			return entities, err
		}
		entities[key] = append(entities[key], entity)
	}

	return entities, nil
}

// isCustomType returns true if given value is role else returns false
func (e Entities) isRole(t string) bool {
	for i := range Types {
		if Types[i] == t {
			return false
		}
	}

	return true
}
