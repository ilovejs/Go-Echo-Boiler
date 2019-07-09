package doc

import (
	log "github.com/sirupsen/logrus"
)

func old_main() {
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
}
