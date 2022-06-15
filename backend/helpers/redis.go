package helpers

import (
	"encoding/json"
	"fmt"
	"github.com/tab-projekt-backend/database/redis"
	"net/http"
	"os"
)

func GetAuthAndIDFromSession(r *http.Request, level redis.PermissionLevel) (int32, error) {
	authHost := os.Getenv("AUTH_HOST")
	authPort := os.Getenv("AUTH_PORT")

	url := fmt.Sprintf("http://%s:%s/auth/%v", authHost, authPort, int8(level))

	req, err := http.NewRequest(http.MethodGet, url, nil)

	for _, c := range r.Cookies() {
		req.AddCookie(c)
	}

	req.Header = r.Header

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return -1, err
	}

	if resp.StatusCode != http.StatusOK {
		return -1, fmt.Errorf("request for %v level unauthorized", level)
	}
	userId := int32(0)
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&userId)
	if err != nil {
		return -1, err
	}
	return userId, nil
}
