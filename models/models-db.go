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

// Get returns all zas_dv_didf_all_v and error, if any
func (m *DBModel) GetDvDidf() ([]*Signal, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select infor_id, infor_naziv, status from ted.zas_dv_didf_all_v
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
			&signal.Name,
			&signal.Status,
		)

		if err != nil {
			return nil, err
		}

		signals = append(signals, &signal)
	}

	return signals, nil
}

// Get returns all zas_tr_gldif_all_v and error, if any
func (m *DBModel) GetDiffTr() ([]*Signal, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select infor_id, infor_naziv, status from ted.zas_tr_gldif_all_v
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
			&signal.Name,
			&signal.Status,
		)

		if err != nil {
			return nil, err
		}

		signals = append(signals, &signal)
	}

	return signals, nil
}

// Get returns all zas_tr_rez_dis_all_v and error, if any
func (m *DBModel) GetDisTrRes() ([]*Signal, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select infor_id, infor_naziv, status from ted.zas_tr_rez_dis_all_v
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
			&signal.Name,
			&signal.Status,
		)

		if err != nil {
			return nil, err
		}

		signals = append(signals, &signal)
	}

	return signals, nil
}

// Get returns all ted.zas_spp_didf_all_v and error, if any
func (m *DBModel) GetDisDiffSp() ([]*Signal, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select infor_id, infor_naziv, status from ted.zas_spp_didf_all_v
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
			&signal.Name,
			&signal.Status,
		)

		if err != nil {
			return nil, err
		}

		signals = append(signals, &signal)
	}

	return signals, nil
}

// Get returns all pgd.zas_pd_kk_v and error, if any
func (m *DBModel) GetMalfunctionIn() ([]*MalfunctionIn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id_pd_kk,naziv_edd,RED from pgd.zas_pd_kk_v
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var ms []*MalfunctionIn

	for rows.Next() {
		var m MalfunctionIn
		err := rows.Scan(
			&m.ID,
			&m.Name,
			&m.Order,
		)

		if err != nil {
			return nil, err
		}

		ms = append(ms, &m)
	}

	return ms, nil
}

// Get returns all ddn.s_rapu and error, if any
func (m *DBModel) GetAPU() ([]*Apu, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id,naziv,SKR_NAZ,SIFRA, status from ddn.s_rapu
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var ms []*Apu

	for rows.Next() {
		var m Apu
		err := rows.Scan(
			&m.ID,
			&m.Name,
			&m.ShortName,
			&m.Code,
			&m.Status,
		)

		if err != nil {
			return nil, err
		}

		ms = append(ms, &m)
	}

	return ms, nil
}

// Get returns all zas_dv_pres_all_v and error, if any
func (m *DBModel) GetOCDV() ([]*Signal, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select infor_id,infor_naziv,status from zas_dv_pres_all_v
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
			&signal.Name,
			&signal.Status,
		)

		if err != nil {
			return nil, err
		}

		signals = append(signals, &signal)
	}

	return signals, nil
}

// Get returns all zas_tr_glprs_all_v and error, if any
func (m *DBModel) GetOCTR12() ([]*Signal, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select infor_id,infor_naziv,status from zas_tr_glprs_all_v
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
			&signal.Name,
			&signal.Status,
		)

		if err != nil {
			return nil, err
		}

		signals = append(signals, &signal)
	}

	return signals, nil
}

// Get returns all zas_tr_rez_prs_all_v and error, if any
func (m *DBModel) GetOCTRR() ([]*Signal, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select infor_id,infor_naziv,status from zas_tr_rez_prs_all_v
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
			&signal.Name,
			&signal.Status,
		)

		if err != nil {
			return nil, err
		}

		signals = append(signals, &signal)
	}

	return signals, nil
}

// Get returns all zas_spp_prs_all_v and error, if any
func (m *DBModel) GetOCSP() ([]*Signal, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select infor_id,infor_naziv,status from zas_spp_prs_all_v
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
			&signal.Name,
			&signal.Status,
		)

		if err != nil {
			return nil, err
		}

		signals = append(signals, &signal)
	}

	return signals, nil
}

// Get returns all zas_dv_zmsp_all_v and error, if any
func (m *DBModel) GetEarthfaultOCDV() ([]*Signal, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select infor_id,infor_naziv,status from zas_dv_zmsp_all_v
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
			&signal.Name,
			&signal.Status,
		)

		if err != nil {
			return nil, err
		}

		signals = append(signals, &signal)
	}

	return signals, nil
}

// Get returns all zas_tr_glzms_all_v and error, if any
func (m *DBModel) GetEarthfaultOCTR() ([]*Signal, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select infor_id,infor_naziv,status from zas_tr_glzms_all_v
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
			&signal.Name,
			&signal.Status,
		)

		if err != nil {
			return nil, err
		}

		signals = append(signals, &signal)
	}

	return signals, nil
}

// Get returns all zas_spp_zms_all_v and error, if any
func (m *DBModel) GetEarthfaultOCSP() ([]*Signal, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select infor_id,infor_naziv,status from zas_spp_zms_all_v
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
			&signal.Name,
			&signal.Status,
		)

		if err != nil {
			return nil, err
		}

		signals = append(signals, &signal)
	}

	return signals, nil
}

// Get returns all zas_dv_uzms_all_v and error, if any
func (m *DBModel) GetDirEarthfaultOC() ([]*Signal, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select infor_id,infor_naziv,status from zas_dv_uzms_all_v
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
			&signal.Name,
			&signal.Status,
		)

		if err != nil {
			return nil, err
		}

		signals = append(signals, &signal)
	}

	return signals, nil
}

// Get returns all ZAS_DV_TELE_ALL_V and error, if any
func (m *DBModel) GetTPSendRcdv() ([]*Signal, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select infor_id,infor_naziv,status from ZAS_DV_TELE_ALL_V
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
			&signal.Name,
			&signal.Status,
		)

		if err != nil {
			return nil, err
		}

		signals = append(signals, &signal)
	}

	return signals, nil
}

// Get returns all zas_prek_all_v and error, if any
func (m *DBModel) GetCircuitbreaker() ([]*Signal, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select infor_id,infor_naziv,status from zas_prek_all_v
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
			&signal.Name,
			&signal.Status,
		)

		if err != nil {
			return nil, err
		}

		signals = append(signals, &signal)
	}

	return signals, nil
}

// Get returns all zas_jps_all_v and error, if any
func (m *DBModel) GetBBPBFtrip() ([]*Signal, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select infor_id,infor_naziv,status from zas_jps_all_v
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
			&signal.Name,
			&signal.Status,
		)

		if err != nil {
			return nil, err
		}

		signals = append(signals, &signal)
	}

	return signals, nil
}

// Get returns all zas_tr_neel_all_v and error, if any
func (m *DBModel) GetNonElectrical() ([]*Signal, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select infor_id,infor_naziv,status from zas_tr_neel_all_v
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
			&signal.Name,
			&signal.Status,
		)

		if err != nil {
			return nil, err
		}

		signals = append(signals, &signal)
	}

	return signals, nil
}

// Get returns all zas_sab_sbr_all_v and error, if any
func (m *DBModel) GetBBPBBtrip() ([]*Signal, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select infor_id,infor_naziv,status from zas_sab_sbr_all_v
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
			&signal.Name,
			&signal.Status,
		)

		if err != nil {
			return nil, err
		}

		signals = append(signals, &signal)
	}

	return signals, nil
}

// Get returns all zas_sab_opr_all_v and error, if any
func (m *DBModel) GetBFtrip() ([]*Signal, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select infor_id,infor_naziv,status from zas_sab_opr_all_v
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
			&signal.Name,
			&signal.Status,
		)

		if err != nil {
			return nil, err
		}

		signals = append(signals, &signal)
	}

	return signals, nil
}

// Get returns all weather conditions and error, if any
func (m *DBModel) GetWeatherConditions() ([]*WeatherConditions, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, sifra, naziv, status from DDN.S_VREM_USL
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var signals []*WeatherConditions

	for rows.Next() {
		var signal WeatherConditions
		err := rows.Scan(
			&signal.ID,
			&signal.Code,
			&signal.Name,
			&signal.Status,
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

func (m *DBModel) GetUserByUsername(username string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, username, password from tis_services_users where username = :1`

	var user User
	row := m.DB.QueryRowContext(ctx, query, username)

	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Password,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (m *DBModel) GetUserByID(id int) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, username, password from tis_services_users where id = :1`

	var user User
	row := m.DB.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Password,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
