package webhook2pushover

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gregdel/pushover"
)

// Handle will handle the cloud function call
func Handle(w http.ResponseWriter, r *http.Request) {
	token := os.Getenv("TOKEN")
	user := os.Getenv("USER")

	hookBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Panic(err)
	}

	app := pushover.New(token)
	recipient := pushover.NewRecipient(user)

	message := pushover.NewMessage(string(hookBody))
	response, err := app.SendMessage(message, recipient)
	if err != nil {
		log.Panic(err)
	}

	fmt.Fprint(w, "OK")
	log.Println(response)
}
