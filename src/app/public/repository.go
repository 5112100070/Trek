package public

import (
	"log"

	"github.com/5112100070/Trek/src/constants"
	"github.com/5112100070/Trek/src/entity"
	"github.com/5112100070/publib/storage/database"
)

func InitPublicRepo(db map[string]database.Database) *publicRepo {
	return &publicRepo{
		db: db,
	}
}

func (repo publicRepo) SaveSubscriber(user entity.UserSubscriber) error {
	var query = `
		INSERT INTO
			subscriber(fullname, email, company, phone, project)
		VALUES(?, ?, ?, ?, ?)
	`

	_, err := repo.db[constants.DB_TYPE_MAIN].Exec(
		query,
		user.Fullname,
		user.Email,
		user.Company,
		user.PhoneNumber,
		user.ProjectDescription,
	)
	if err != nil {
		log.Println(err)
	}

	return err
}
