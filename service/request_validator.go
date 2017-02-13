package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type InputValidation interface {
	validate(r *http.Request, needs []string) error
}

type Thing struct {
	Seen bool `json:"seen"`
}

func DecodeAndValidate(r *http.Request, needs []string) error {

	var thing *Thing

	if err := json.NewDecoder(r.Body).Decode(&thing); err != nil {

		fmt.Println("errore")

		return err
	}

	defer r.Body.Close()

	thing.validate(r, needs)

	return nil
}

func (t Thing) validate(r *http.Request, needs []string) error {

	fmt.Println(t.Seen)
	fmt.Println(needs)

	for _, need := range needs {

		switch need {

		case "Seen":
			govalidator
		}
	}

	return nil
	// // validate the ID is a uuid
	// if !govalidator.IsUUID(t.ID) {
	// 	return ErrInvalidUUID
	// }
	// // validate the name is not empty or missing
	// if govalidator.IsNull(t.Name) {
	// 	return ErrInvalidName
	// }
	// return nil
}
