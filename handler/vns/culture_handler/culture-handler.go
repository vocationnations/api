package culture_handler

import (
	"encoding/json"
	"fmt"
	"github.com/vocationnations/api/helper"
	"github.com/vocationnations/api/model"
	"net/http"
)

func CreateCultureEntry(w http.ResponseWriter, r *http.Request, ctx helper.AppContext) error {

	var cltrEntry model.CultureEntry
	db := ctx.DB.GetDatabase()

	if err := json.NewDecoder(r.Body).Decode(&cltrEntry); err != nil {
		return fmt.Errorf("cannot decode the body for culture entry, err: %v", err)
	}

	fmt.Printf("CultureEntry: %s", cltrEntry)

	if _, err := db.Model(&cltrEntry).Insert(); err != nil {
		return fmt.Errorf("cannot insert culture entry to the database, err: %v", err)
	}

	if err := json.NewEncoder(w).Encode(&cltrEntry); err != nil {
		return fmt.Errorf("cannot encode the culture entry for printing, err: %v", err)
	}
	return nil
}
