package goalcom

import (
	"encoding/json"
	"fmt"
	"net/http"

	model "github.com/ibra-bybuy/wsports-parser/pkg/model/goalcom"
)

func (r *repository) request(response *model.Response) error {
	resp, err := http.Get(fmt.Sprintf("https://www.goal.com/api/live-scores/refresh?edition=en&date=%v&tzoffset=%d", r.date, r.offset))

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return err
	}

	return nil
}
