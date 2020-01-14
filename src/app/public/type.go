package public

import (
	"github.com/5112100070/publib/storage/database"
)

type publicRepo struct {
	db map[string]database.Database
}
