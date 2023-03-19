package repositories

import (
	"database/sql"

	"github.com/nuttchai/go-rest/internal/models"
	"github.com/nuttchai/go-rest/internal/repositories/interfaces"
	"github.com/nuttchai/go-rest/internal/utils/context"
)

type TUserRepository struct {
	DB *sql.DB
}

var (
	UserRepository interfaces.IUserRepository
)

func (m *TUserRepository) RetrieveOne(id string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(3)
	defer cancel()

	// NOTE: cannot directly use 'user' as table name because it is a reserved keyword
	query := "select * from public.user where id = $1"
	row := m.DB.QueryRowContext(ctx, query, id)

	var user models.User
	err := row.Scan(
		&user.Id,
		&user.Username,
	)

	return &user, err
}
