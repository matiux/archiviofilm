package service

import (
	"encoding/json"
	"errors"
	// "fmt"
	"gopkg.in/oleiade/reflections.v1"
	"net/http"
	"reflect"
)

type InputValidation interface {
	validate(needs []string) error
}

type Thing struct {
	Seen bool
}

func DecodeAndValidate(r *http.Request, needs []string) (*Thing, error) {

	var thing *Thing

	if err := json.NewDecoder(r.Body).Decode(&thing); err != nil {

		return nil, err
	}

	defer r.Body.Close()

	if err := thing.validate(needs); err != nil {

		return nil, err
	}

	return thing, nil
}

func (t Thing) validate(needs []string) error {

	var err error

	for _, need := range needs {

		switch need {

		case "Seen":
			err = t.seenValidator(need)
			break
		}
	}

	return err
}

func (t Thing) seenValidator(property string) error {

	value, err := reflections.GetField(t, property)

	if err != nil {

		return errors.New("Property '" + property + "' - " + err.Error())
	}

	if reflect.TypeOf(value).Kind().String() != "bool" {

		return errors.New("Errori di validazione con la proprietà '" + property + "'")
	}

	return nil
}
