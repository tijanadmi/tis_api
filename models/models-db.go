package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type DBModel struct {
	DB *sql.DB
}

// Get returns one movie and error, if any
func (m *DBModel) GetDvDidf() ([]*Signal, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select infor_id, infor_sifra, infor_naziv from ted.zas_dv_didf_v
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var signals []*Signal

	for rows.Next() {
		var signal Signal
		err := rows.Scan(
			&signal.ID,
			&signal.Code,
			&signal.Name,
		)

		if err != nil {
			return nil, err
		}

		signals = append(signals, &signal)
	}

	return signals, nil
}

// Authenticate authenticates a user
func (m *DBModel) Authenticate(username, testPassword string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	/*var id int
	var hashedPassword string*/

	var user User

	query := `select id, username, password from tis_services_users where username = :1`

	row := m.DB.QueryRowContext(ctx, query, username)
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(testPassword))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return errors.New("incorrect password")
	} else if err != nil {
		return err
	}
	return nil
}
