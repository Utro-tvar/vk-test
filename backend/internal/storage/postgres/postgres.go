package postgres

import (
	"database/sql"
	"fmt"
	"net"
	"strings"

	"github.com/Utro-tvar/vk-test/backend/internal/models"
	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func New(host, port, user, password, dbname string) (*Storage, error) {
	const op = "storage.postgres.New"

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) Store(data []models.Container) error {
	const op = "storage.postgres.Store"

	args := make([]string, 0, len(data)*3)

	for _, v := range data {
		args = append(args, fmt.Sprintf("('%s', %d, '%s')", v.IP.String(), v.Ping, v.LastConnection.Format("2006-01-02")))
	}

	query := fmt.Sprintf(`
		INSERT INTO containers (ip, ping, last_conn)
		VALUES %s
		ON CONFLICT (ip) DO UPDATE
		SET 
			ping = EXCLUDED.ping,
			last_conn = EXCLUDED.last_conn;
	`, strings.Join(args, ", "))

	_, err := s.db.Exec(query)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) GetAll() ([]models.Container, error) {
	const op = "storage.postgres.GetAll"

	rows, err := s.db.Query("SELECT ip, ping, last_conn FROM containers")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	containers := make([]models.Container, 0)

	for rows.Next() {
		var cont models.Container
		var ipStr string
		err := rows.Scan(&ipStr, &cont.Ping, &cont.LastConnection)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		cont.IP = net.ParseIP(ipStr)

		containers = append(containers, cont)
	}

	err = rows.Err()
	if err != nil {
		return containers, err
	}

	return containers, nil
}

func (s *Storage) Close() {
	s.db.Close()
}
