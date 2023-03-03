package postgres_db

import (
	"github.com/jmoiron/sqlx"
	"github.com/mokan-r/pitch/db/postgres"
	"github.com/mokan-r/pitch/internal/entity"
	"github.com/sirupsen/logrus"
)

var schema = `
CREATE TABLE IF NOT EXISTS songs (
    uid uuid PRIMARY KEY,
    name varchar(255)
);`

type Repository struct {
	cfg entity.PostgresConfig
	db  *sqlx.DB
}

func New(cfg entity.PostgresConfig) (*Repository, error) {
	db, err := postgres.NewClient(postgres.Config{
		Host: cfg.Host,
		Port: cfg.Port,
	})
	if err != nil {
		return &Repository{}, err
	}
	if _, err := db.Exec(schema); err != nil {
		return &Repository{}, err
	}
	return &Repository{
		db:  db,
		cfg: cfg}, nil
}

func (r *Repository) GetConfig() entity.PostgresConfig {
	return r.cfg
}

func (r *Repository) Ping() error {
	return r.db.Ping()
}

func (r *Repository) Set(song entity.Song) error {
	logrus.Info(song)
	query := `INSERT INTO songs (uid, title, duration)
				VALUES ($1, $2, $3)
				ON CONFLICT (uid) DO UPDATE
				SET title = EXCLUDED.title, duration = EXCLUDED.duration;`
	_, err := r.db.Exec(query, song.Id, song.Title, song.Duration)
	return err
}

func (r *Repository) Get(id int) (out entity.Song, err error) {
	query := `SELECT * FROM songs WHERE uid = $1;`
	err = r.db.Get(&out, query, id)
	return out, err
}

func (r *Repository) Delete(id int) (err error) {
	query := `DELETE FROM songs WHERE uid = $1;`
	_, err = r.db.Exec(query, id)
	return err
}

func (r *Repository) GetAll() (out []entity.Song, err error) {
	query := `SELECT * FROM songs;`
	err = r.db.Select(&out, query)
	return out, err
}
