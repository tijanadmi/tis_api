package dbrepo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/tijanadmi/tis-api/models"
	"golang.org/x/crypto/bcrypt"
)

type OracleDBRepo struct {
	DB *sql.DB
}

func (m *OracleDBRepo) Connection() *sql.DB {
	return m.DB
}

func (m *OracleDBRepo) OneSignal(id int) (*models.Signal, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select id, naziv, status
			  from zsi_info
			  where id=:1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var signal models.Signal

	err := row.Scan(
		&signal.ID,
		&signal.Name,
		&signal.Status,
	)

	if err != nil {
		return nil, err
	}

	return &signal, err
}

// Get returns all zas_dv_didf_all_v and error, if any
func (m *OracleDBRepo) GetDvDidf() ([]*models.Signal, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select infor_id, infor_naziv, status from ted.zas_dv_didf_all_v
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var signals []*models.Signal

	for rows.Next() {
		var signal models.Signal
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
func (m *OracleDBRepo) GetDiffTr() ([]*models.Signal, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select infor_id, infor_naziv, status from ted.zas_tr_gldif_all_v
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var signals []*models.Signal

	for rows.Next() {
		var signal models.Signal
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
func (m *OracleDBRepo) GetDisTrRes() ([]*models.Signal, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select infor_id, infor_naziv, status from ted.zas_tr_rez_dis_all_v
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var signals []*models.Signal

	for rows.Next() {
		var signal models.Signal
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
func (m *OracleDBRepo) GetDisDiffSp() ([]*models.Signal, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select infor_id, infor_naziv, status from ted.zas_spp_didf_all_v
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var signals []*models.Signal

	for rows.Next() {
		var signal models.Signal
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
func (m *OracleDBRepo) GetMalfunctionIn() ([]*models.MalfunctionIn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select id_pd_kk,naziv_edd,RED from pgd.zas_pd_kk_v
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var ms []*models.MalfunctionIn

	for rows.Next() {
		var m models.MalfunctionIn
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

func (m *OracleDBRepo) OneMalfunctionIn(id int) (*models.MalfunctionIn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select id_pd_kk,naziv_edd,RED
			  from zas_pd_kk_v
			  where id_pd_kk=:1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var mf models.MalfunctionIn

	err := row.Scan(
		&mf.ID,
		&mf.Name,
		&mf.Order,
	)

	if err != nil {
		return nil, err
	}

	return &mf, err
}

// Get returns all ddn.s_rapu and error, if any
func (m *OracleDBRepo) GetAPU() ([]*models.Apu, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select id,naziv,SKR_NAZ,SIFRA, status from ddn.s_rapu
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var ms []*models.Apu

	for rows.Next() {
		var m models.Apu
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

func (m *OracleDBRepo) OneAPU(id int) (*models.Apu, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select id,naziv,SKR_NAZ,SIFRA, status
			  from s_rapu
			  where id=:1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var apu models.Apu

	err := row.Scan(
		&apu.ID,
		&apu.Name,
		&apu.ShortName,
		&apu.Code,
		&apu.Status,
	)

	if err != nil {
		return nil, err
	}

	return &apu, err
}

// Get returns all zas_dv_pres_all_v and error, if any
func (m *OracleDBRepo) GetOCDV() ([]*models.Signal, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select infor_id,infor_naziv,status from zas_dv_pres_all_v
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var signals []*models.Signal

	for rows.Next() {
		var signal models.Signal
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
func (m *OracleDBRepo) GetOCTR12() ([]*models.Signal, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select infor_id,infor_naziv,status from zas_tr_glprs_all_v
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var signals []*models.Signal

	for rows.Next() {
		var signal models.Signal
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
func (m *OracleDBRepo) GetOCTRR() ([]*models.Signal, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select infor_id,infor_naziv,status from zas_tr_rez_prs_all_v
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var signals []*models.Signal

	for rows.Next() {
		var signal models.Signal
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
func (m *OracleDBRepo) GetOCSP() ([]*models.Signal, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select infor_id,infor_naziv,status from zas_spp_prs_all_v
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var signals []*models.Signal

	for rows.Next() {
		var signal models.Signal
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
func (m *OracleDBRepo) GetEarthfaultOCDV() ([]*models.Signal, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select infor_id,infor_naziv,status from zas_dv_zmsp_all_v
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var signals []*models.Signal

	for rows.Next() {
		var signal models.Signal
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
func (m *OracleDBRepo) GetEarthfaultOCTR() ([]*models.Signal, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select infor_id,infor_naziv,status from zas_tr_glzms_all_v
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var signals []*models.Signal

	for rows.Next() {
		var signal models.Signal
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
func (m *OracleDBRepo) GetEarthfaultOCSP() ([]*models.Signal, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select infor_id,infor_naziv,status from zas_spp_zms_all_v
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var signals []*models.Signal

	for rows.Next() {
		var signal models.Signal
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

// Get returns all zas_tr_rez_zms_all_v and error, if any
func (m *OracleDBRepo) GetEarthfaultOCTRR() ([]*models.Signal, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select infor_id,infor_naziv,status from zas_tr_rez_zms_all_v
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var signals []*models.Signal

	for rows.Next() {
		var signal models.Signal
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
func (m *OracleDBRepo) GetDirEarthfaultOC() ([]*models.Signal, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select infor_id,infor_naziv,status from zas_dv_uzms_all_v
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var signals []*models.Signal

	for rows.Next() {
		var signal models.Signal
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
func (m *OracleDBRepo) GetTPSendRcdv() ([]*models.Signal, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select infor_id,infor_naziv,status from ZAS_DV_TELE_ALL_V
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var signals []*models.Signal

	for rows.Next() {
		var signal models.Signal
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
func (m *OracleDBRepo) GetCircuitbreaker() ([]*models.Signal, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select infor_id,infor_naziv,status from zas_prek_all_v
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var signals []*models.Signal

	for rows.Next() {
		var signal models.Signal
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
func (m *OracleDBRepo) GetBBPBFtrip() ([]*models.Signal, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select infor_id,infor_naziv,status from zas_jps_all_v
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var signals []*models.Signal

	for rows.Next() {
		var signal models.Signal
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
func (m *OracleDBRepo) GetNonElectrical() ([]*models.Signal, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select infor_id,infor_naziv,status from zas_tr_neel_all_v
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var signals []*models.Signal

	for rows.Next() {
		var signal models.Signal
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
func (m *OracleDBRepo) GetBBPBBtrip() ([]*models.Signal, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select infor_id,infor_naziv,status from zas_sab_sbr_all_v
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var signals []*models.Signal

	for rows.Next() {
		var signal models.Signal
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
func (m *OracleDBRepo) GetBFtrip() ([]*models.Signal, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select infor_id,infor_naziv,status from zas_sab_opr_all_v
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var signals []*models.Signal

	for rows.Next() {
		var signal models.Signal
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

// Get returns all groups of causes and error, if any
func (m *OracleDBRepo) GetGroupsOfCauses() ([]*models.GroupOfCause, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select id, sifra, naziv, skr, status, sif_ddn, sortr
			  from pgi.s_gruzr
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var grcs []*models.GroupOfCause

	for rows.Next() {
		var grc models.GroupOfCause
		err := rows.Scan(
			&grc.ID,
			&grc.Code,
			&grc.Name,
			&grc.ShortName,
			&grc.Status,
			&grc.DDNCode,
			&grc.Sort,
		)

		if err != nil {
			return nil, err
		}

		grcs = append(grcs, &grc)
	}

	return grcs, nil
}

func (m *OracleDBRepo) OneGroupOfCauses(id int) (*models.GroupOfCause, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select id, sifra, naziv, skr, status, sif_ddn, sortr
			  from pgi.s_gruzr	
			  where id=:1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var grc models.GroupOfCause

	err := row.Scan(
		&grc.ID,
		&grc.Code,
		&grc.Name,
		&grc.ShortName,
		&grc.Status,
		&grc.DDNCode,
		&grc.Sort,
	)

	if err != nil {
		return nil, err
	}

	return &grc, err
}

// Get returns all causes and error, if any
func (m *OracleDBRepo) GetCauses() ([]*models.Cause, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select u.id, u.id_s_gruzr, u.sifra, u.naziv, u.skr, u.status, u.sortr, gu.id, gu.sifra, gu.naziv, gu.skr, gu.status, gu.sif_ddn, gu.sortr
			  from pgi.s_uzrok u, pgi.s_gruzr gu
			  where u.id_s_gruzr=gu.id
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var caus []*models.Cause

	for rows.Next() {
		var cau models.Cause
		err := rows.Scan(
			&cau.ID,
			&cau.GroupCauseId,
			&cau.Code,
			&cau.Name,
			&cau.ShortName,
			&cau.Status,
			&cau.Sort,
			&cau.GroupOfCause.ID,
			&cau.GroupOfCause.Code,
			&cau.GroupOfCause.Name,
			&cau.GroupOfCause.ShortName,
			&cau.GroupOfCause.Status,
			&cau.GroupOfCause.DDNCode,
			&cau.GroupOfCause.Sort,
		)

		if err != nil {
			return nil, err
		}

		caus = append(caus, &cau)
	}

	return caus, nil
}

func (m *OracleDBRepo) OneCause(id int) (*models.Cause, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select u.id, u.id_s_gruzr, u.sifra, u.naziv, u.skr, u.status, u.sortr, gu.id, gu.sifra, gu.naziv, gu.skr, gu.status, gu.sif_ddn, gu.sortr
			  from pgi.s_uzrok u, pgi.s_gruzr gu
			  where u.id_s_gruzr=gu.id
			  and u.id=:1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var cau models.Cause

	err := row.Scan(
		&cau.ID,
		&cau.GroupCauseId,
		&cau.Code,
		&cau.Name,
		&cau.ShortName,
		&cau.Status,
		&cau.Sort,
		&cau.GroupOfCause.ID,
		&cau.GroupOfCause.Code,
		&cau.GroupOfCause.Name,
		&cau.GroupOfCause.ShortName,
		&cau.GroupOfCause.Status,
		&cau.GroupOfCause.DDNCode,
		&cau.GroupOfCause.Sort,
	)

	if err != nil {
		return nil, err
	}

	return &cau, err
}

// Get returns all groups of reasons and error, if any
func (m *OracleDBRepo) GetGroupOfReasons() ([]*models.GroupOfReason, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select id, sifra, naziv, skr, status, sortr
	from PGI.S_GRRAZ
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var grs []*models.GroupOfReason

	for rows.Next() {
		var gr models.GroupOfReason
		err := rows.Scan(
			&gr.ID,
			&gr.Code,
			&gr.Name,
			&gr.ShortName,
			&gr.Status,
			&gr.Sort,
		)

		if err != nil {
			return nil, err
		}

		grs = append(grs, &gr)
	}

	return grs, nil
}

func (m *OracleDBRepo) OneGroupOfReasons(id int) (*models.GroupOfReason, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select id, sifra, naziv, skr, status, sortr
	          from PGI.S_GRRAZ		
			  where id=:1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var gr models.GroupOfReason

	err := row.Scan(
		&gr.ID,
		&gr.Code,
		&gr.Name,
		&gr.ShortName,
		&gr.Status,
		&gr.Sort,
	)

	if err != nil {
		return nil, err
	}

	return &gr, err
}

// Get returns all  reasons and error, if any
func (m *OracleDBRepo) GetReasons() ([]*models.Reason, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select r.id, r.id_s_grraz, r.sifra, r.naziv, r.skr, r.status, r.sortr,gr.id, gr.sifra, gr.naziv, gr.skr, gr.status, gr.sortr
			  from pgi.s_razlog r, PGI.S_GRRAZ gr
		 	  where r.id_s_grraz=gr.id	
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var rs []*models.Reason

	for rows.Next() {
		var r models.Reason
		err := rows.Scan(
			&r.ID,
			&r.GroupReasonId,
			&r.Code,
			&r.Name,
			&r.ShortName,
			&r.Status,
			&r.Sort,
			&r.GroupOfReason.ID,
			&r.GroupOfReason.Code,
			&r.GroupOfReason.Name,
			&r.GroupOfReason.ShortName,
			&r.GroupOfReason.Status,
			&r.GroupOfReason.Sort,
		)

		if err != nil {
			return nil, err
		}

		rs = append(rs, &r)
	}

	return rs, nil
}

func (m *OracleDBRepo) OneReason(id int) (*models.Reason, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select r.id, r.id_s_grraz, r.sifra, r.naziv, r.skr, r.status, r.sortr,gr.id, gr.sifra, gr.naziv, gr.skr, gr.status, gr.sortr
			  from pgi.s_razlog r, PGI.S_GRRAZ gr
	 		  where r.id_s_grraz=gr.id
			  and r.id=:1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var r models.Reason

	err := row.Scan(
		&r.ID,
		&r.GroupReasonId,
		&r.Code,
		&r.Name,
		&r.ShortName,
		&r.Status,
		&r.Sort,
		&r.GroupOfReason.ID,
		&r.GroupOfReason.Code,
		&r.GroupOfReason.Name,
		&r.GroupOfReason.ShortName,
		&r.GroupOfReason.Status,
		&r.GroupOfReason.Sort,
	)

	if err != nil {
		return nil, err
	}

	return &r, err
}

// Get returns all weather conditions and error, if any
func (m *OracleDBRepo) GetWeatherConditions() ([]*models.WeatherCondition, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select id, sifra, naziv, status from DDN.S_VREM_USL
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var signals []*models.WeatherCondition

	for rows.Next() {
		var signal models.WeatherCondition
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

func (m *OracleDBRepo) OneWeatherCondition(id int) (*models.WeatherCondition, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select id, sifra, naziv, status
			  from DDN.S_VREM_USL			
			  where id=:1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var signal models.WeatherCondition

	err := row.Scan(
		&signal.ID,
		&signal.Code,
		&signal.Name,
		&signal.Status,
	)

	if err != nil {
		return nil, err
	}

	return &signal, err
}

// Get returns all categories of events and error, if any
func (m *OracleDBRepo) GetCategoriesOfEvents() ([]*models.CategoryOfEvent, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select id, id_s_tipd, sifra, naziv, skr, status, sortr
			  from s_vrpd
			  where id_s_tipd in (1,2,3)	
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var ces []*models.CategoryOfEvent

	for rows.Next() {
		var ce models.CategoryOfEvent
		err := rows.Scan(
			&ce.ID,
			&ce.TypeEventId,
			&ce.Code,
			&ce.Name,
			&ce.ShortName,
			&ce.Status,
			&ce.Sort,
		)

		if err != nil {
			return nil, err
		}

		ces = append(ces, &ce)
	}

	return ces, nil
}

func (m *OracleDBRepo) OneCategoryOfEvents(id int) (*models.CategoryOfEvent, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select id, id_s_tipd, sifra, naziv, skr, status, sortr
			  from s_vrpd				
			  where id=:1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var ce models.CategoryOfEvent

	err := row.Scan(
		&ce.ID,
		&ce.TypeEventId,
		&ce.Code,
		&ce.Name,
		&ce.ShortName,
		&ce.Status,
		&ce.Sort,
	)

	if err != nil {
		return nil, err
	}

	return &ce, err
}

// Get returns all weather conditions and error, if any
func (m *OracleDBRepo) GetOHL() ([]*models.OverheadLine, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select TIS_ID_DV,IPS_ID_DV,SAP_ID_DV,ID_DAL,NAZIV_IPS,OPIS_PO_KATEGORIZACIJI,ID_NAP,NAPON_NAZIV,COALESCE(to_char(KATEGORIJA_ID), ''),COALESCE(KATEGORIJA, ''),TIS_ID_PT_POLJA,IPS_ID_PT_POLJA,
			  SAP_ID_PT_POLJA,BROJ_POLJA_PT,NAZIV_PT_POLJA,NAZIV_PT_POLJA_PO_KATEGORIZACIJI,TIS_ID_KT_POLJA,IPS_ID_KT_POLJA,SAP_ID_KT_POLJA,BROJ_POLJA_KT,
	          NAZIV_KT_POLJA,NAZIV_KT_POLJA_PO_KATEGORIZACIJI,TIPOB_ID,TIPOB_SIFRA,TIPOB_NAZIV,COALESCE(to_char(ID_S_MRC1), ''),COALESCE(MRC1, ''),COALESCE(to_char(ID_S_MRC2), ''),COALESCE(MRC2, '')
 	          from synsoft_DV_a
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var ohls []*models.OverheadLine

	for rows.Next() {
		var ohl models.OverheadLine
		err := rows.Scan(
			&ohl.ID,
			&ohl.IpsIdDV,
			&ohl.SapIdDV,
			&ohl.IdDal,
			&ohl.IPSName,
			&ohl.CategoryNameDV,
			&ohl.IdNN,
			&ohl.NNName,
			&ohl.CategoryId,
			&ohl.CategoryName,
			&ohl.TisIdFirstFeeder,
			&ohl.IpsIdFirstFeeder,
			&ohl.SapIdFirstFeeder,
			&ohl.FirstFeederNumber,
			&ohl.FirstFeederName,
			&ohl.FirstFeederCategoryName,
			&ohl.TisIdLastFeeder,
			&ohl.IpsIdLastFeeder,
			&ohl.SapIdLastFeeder,
			&ohl.LastFeederNumber,
			&ohl.LastFeederName,
			&ohl.LastFeederCategoryName,
			&ohl.ObjTypeId,
			&ohl.ObjTypeCode,
			&ohl.ObjTypeName,
			&ohl.Mrc1Id,
			&ohl.Mrc1Name,
			&ohl.Mrc2Id,
			&ohl.Mrc2Name,
		)

		if err != nil {
			return nil, err
		}

		ohls = append(ohls, &ohl)
	}

	return ohls, nil
}

// Get returns all power_cables and error, if any
func (m *OracleDBRepo) GetPowerCables() ([]*models.PowerCable, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select TIS_ID_KB,IPS_ID_KB,SAP_ID_KB,ID_KAB,NAZIV_IPS,OPIS_PO_KATEGORIZACIJI,COALESCE(to_char(KATEGORIJA_ID), ''),COALESCE(KATEGORIJA, ''),TIS_ID_PT_POLJA,IPS_ID_PT_POLJA,
			  SAP_ID_PT_POLJA,BROJ_POLJA_PT,NAZIV_PT_POLJA,NAZIV_PT_POLJA_PO_KATEGORIZACIJI,TIS_ID_KT_POLJA,IPS_ID_KT_POLJA,SAP_ID_KT_POLJA,BROJ_POLJA_KT,
	          NAZIV_KT_POLJA,NAZIV_KT_POLJA_PO_KATEGORIZACIJI,TIPOB_ID,TIPOB_SIFRA,TIPOB_NAZIV,COALESCE(to_char(ID_S_MRC1), ''),COALESCE(MRC1, ''),COALESCE(to_char(ID_S_MRC2), ''),COALESCE(MRC2, '')
 	          from synsoft_KB_a
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var cabs []*models.PowerCable

	for rows.Next() {
		var cab models.PowerCable
		err := rows.Scan(
			&cab.ID,
			&cab.IpsIdKB,
			&cab.SapIdKB,
			&cab.IdKab,
			&cab.IPSName,
			&cab.CategoryNameDV,
			&cab.CategoryId,
			&cab.CategoryName,
			&cab.TisIdFirstFeeder,
			&cab.IpsIdFirstFeeder,
			&cab.SapIdFirstFeeder,
			&cab.FirstFeederNumber,
			&cab.FirstFeederName,
			&cab.FirstFeederCategoryName,
			&cab.TisIdLastFeeder,
			&cab.IpsIdLastFeeder,
			&cab.SapIdLastFeeder,
			&cab.LastFeederNumber,
			&cab.LastFeederName,
			&cab.LastFeederCategoryName,
			&cab.ObjTypeId,
			&cab.ObjTypeCode,
			&cab.ObjTypeName,
			&cab.Mrc1Id,
			&cab.Mrc1Name,
			&cab.Mrc2Id,
			&cab.Mrc2Name,
		)

		if err != nil {
			return nil, err
		}

		cabs = append(cabs, &cab)
	}

	return cabs, nil
}

// Get returns all substations and error, if any
func (m *OracleDBRepo) GetSubstations() ([]*models.Substation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select TIS_ID,IPS_ID,SAP_ID,NAZIV_TS,NAZIV_PO_KATEGORIZACIJI,TIPOB,TIPOB_SIFRA,TIPOB_NAZIV,COALESCE(to_char(ID_S_MRC1), ''),COALESCE(MRC1, ''),COALESCE(to_char(ID_S_MRC2), ''),COALESCE(MRC2, ''), ID_S_ORG, NAZIV_ORG
 	          from synsoft_TS_a
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var subs []*models.Substation

	for rows.Next() {
		var sub models.Substation
		err := rows.Scan(
			&sub.ID,
			&sub.IpsId,
			&sub.SapId,
			&sub.TISName,
			&sub.CategoryNameSub,
			&sub.ObjTypeId,
			&sub.ObjTypeCode,
			&sub.ObjTypeName,
			&sub.Mrc1Id,
			&sub.Mrc1Name,
			&sub.Mrc2Id,
			&sub.Mrc2Name,
			&sub.OrgId,
			&sub.OrgName,
		)

		if err != nil {
			return nil, err
		}

		subs = append(subs, &sub)
	}

	return subs, nil
}

// Get returns all feeders and error, if any
func (m *OracleDBRepo) GetFeeders() ([]*models.Feeder, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select TIS_ID_POLJA,IPS_ID_POLJA,SAP_ID_POLJA,TIS_ID_TRAFOSTANICE,IPS_ID_TRAFOSTANICE,SAP_ID_TRAFOSTANICE,NAZIV_TRAFOSTANICE,
			  NN_ID,NN_NAZIV,BROJ_POLJA,NAZIV_POLJA,OPIS_PO_KATEGORIZACIJI,COALESCE(to_char(KATEGORIJA_ID), ''),COALESCE(KATEGORIJA, ''),COALESCE(to_char(FUNKCIJA_POLJA_ID), ''),COALESCE(FUNKCIJA_POLJA, ''),
			  OPREM,AKTNE,POTPUN,TIP_POLJA
 	          from synsoft_polja_a
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var feeds []*models.Feeder

	for rows.Next() {
		var feed models.Feeder
		err := rows.Scan(
			&feed.ID,
			&feed.IpsIdFeeder,
			&feed.SapIdFeeder,
			&feed.TisIdSub,
			&feed.IpsIdSub,
			&feed.SapIdSub,
			&feed.NameSub,
			&feed.IdNN,
			&feed.NNName,
			&feed.FeederNumber,
			&feed.FeederName,
			&feed.FeederCategoryName,
			&feed.CategoryId,
			&feed.CategoryName,
			&feed.FeederFunId,
			&feed.FeederFunName,
			&feed.Equipped,
			&feed.Active,
			&feed.Completely,
			&feed.FeederType,
		)

		if err != nil {
			return nil, err
		}

		feeds = append(feeds, &feed)
	}

	return feeds, nil
}

// Get returns all protection devices and error, if any
func (m *OracleDBRepo) GetProtectionDevices() ([]*models.ProtectionDevice, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select TIS_ID_OPREME, IPS_ID_OPREME, TIS_ID_POLJA, IPS_ID_POLJA, SAP_ID_POLJA, PROIZVODJAC, TIP, VRSTA_OPR, IND_GL,
			  TEHNOLOGIJA, STATUS_IPS, STATUS, SERIJSKI_BR, OZNAKA_URE, APU_URE
 	          from SYNSOFT_ZASTITNA_OPR
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var prts []*models.ProtectionDevice

	for rows.Next() {
		var prt models.ProtectionDevice
		err := rows.Scan(
			&prt.ID,
			&prt.IpsIdDV,
			&prt.TisIdFeeder,
			&prt.IpsIdFeeder,
			&prt.SapIdFFeeder,
			&prt.Manufacturer,
			&prt.Type,
			&prt.TypeOfEqu,
			&prt.Indicator,
			&prt.Technology,
			&prt.StatusIPS,
			&prt.Status,
			&prt.SerialNumber,
			&prt.DeviceTag,
			&prt.APU,
		)

		if err != nil {
			return nil, err
		}

		prts = append(prts, &prt)
	}

	return prts, nil
}

// Get returns all power transformers and error, if any
func (m *OracleDBRepo) GetPowerTransformers() ([]*models.PowerTransformer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select TIS_ID_TRANSFORMATORA, IPS_ID_TRANSFORMATORA, SAP_ID_TRANSFORMATORA, TIS_ID_TRAFOSTANICE, IPS_ID_TRAFOSTANICE, SAP_ID_TRAFOSTANICE, NAZIV_TRAFOSTANICE,
	NAZIV_TRANSFORMATORA, OPIS_TRANSF_PO_KATEGOR, COALESCE(to_char(KATEGORIJA_ID), ''),COALESCE(KATEGORIJA, ''), PRENOSNI_ODNOS, TIS_ID_VN_POLJA, IPS_ID_VN_POLA, SAP_ID_VN_POLA, NAZIV_VN_POLJA,
	COALESCE(to_char(TIS_ID_SN1_POLJA), ''), COALESCE(IPS_ID_SN1_POLA, ''), COALESCE(SAP_ID_SN1_POLA, ''), COALESCE(NAZIV_SN1_POLJA, ''), COALESCE(to_char(TIS_ID_SN2_POLJA), ''), COALESCE(IPS_ID_SN2_POLA, ''), COALESCE(SAP_ID_SN2_POLA, ''), COALESCE(NAZIV_SN2_POLJA, ''), COALESCE(to_char(TIS_ID_T_POLJA), ''),
	COALESCE(IPS_ID_T_POLA, ''), COALESCE(SAP_ID_T_POLA, ''), COALESCE(NAZIV_T_POLJA, '')
 	          from SYNSOFT_TR_A
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var trs []*models.PowerTransformer

	for rows.Next() {
		var tr models.PowerTransformer
		err := rows.Scan(
			&tr.ID,
			&tr.IpsIdDV,
			&tr.SapIdDV,
			&tr.TisIdSub,
			&tr.IpsIdSub,
			&tr.SapIdSub,
			&tr.NameSub,
			&tr.NameTr,
			&tr.CategoryNameTr,
			&tr.CategoryId,
			&tr.CategoryName,
			&tr.Un,
			&tr.TisIdPrimFeeder,
			&tr.IpsIdPrimFeeder,
			&tr.SapIdPrimFeeder,
			&tr.PrimFeederName,
			&tr.TisIdSec1Feeder,
			&tr.IpsIdSec1Feeder,
			&tr.SapIdSec1Feeder,
			&tr.Sec1FeederName,
			&tr.TisIdSec2Feeder,
			&tr.IpsIdSec2Feeder,
			&tr.SapIdSec2Feeder,
			&tr.Sec2FeederName,
			&tr.TisIdTerFeeder,
			&tr.IpsIdTerFeeder,
			&tr.SapIdTerFeeder,
			&tr.TerFeederName,
		)

		if err != nil {
			return nil, err
		}

		trs = append(trs, &tr)
	}

	return trs, nil
}

// Get returns all disconnectors and error, if any
func (m *OracleDBRepo) GetDisconnectors() ([]*models.Disconnector, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select TIS_ID_RASTAVLJACA, IPS_ID_RASTAVLJACA, SAP_ID_RASTAVLJACA, TIS_ID_POLJA, IPS_ID_POLJA, SAP_ID_POLJA, NAZIV_POLJA, OPIS_PO_KATEGORIZACIJI,
			  COALESCE(to_char(ID_FUN_RAS), ''), COALESCE(FUNKCIJA_RAST_U_POLJU, ''), COALESCE(to_char(KATEGORIJA_ID), ''),COALESCE(KATEGORIJA, '')
		     from synsoft_rastavljaci_a
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var diss []*models.Disconnector

	for rows.Next() {
		var dis models.Disconnector
		err := rows.Scan(
			&dis.ID,
			&dis.IpsIdDs,
			&dis.SapIdDs,
			&dis.TisIdFeeder,
			&dis.IpsIdFeeder,
			&dis.SapIdFeeder,
			&dis.FeederName,
			&dis.DisCategoryName,
			&dis.FunDisId,
			&dis.FunDis,
			&dis.CategoryId,
			&dis.CategoryName,
		)

		if err != nil {
			return nil, err
		}

		diss = append(diss, &dis)
	}

	return diss, nil
}

// Get returns all permissions and error, if any
func (m *OracleDBRepo) GetWorkPermissions() ([]*models.WorkPermission, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select RBR, COALESCE(to_char(BROJ_RAD1), ''), COALESCE(to_char(BROJ_RAD2), ''), COALESCE(PL_VREME_OD, ''), 
			  COALESCE(to_char(PL_DATUM_OD,'dd.mm.yyyy'), ''), COALESCE(NAPOMENA, ''), COALESCE(STATUS, ''), COALESCE(ZAV_VREME, ''), COALESCE(to_char(ZAV_DATUM,'dd.mm.yyyy'), '')
			  from synsoft_dozvole
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var prms []*models.WorkPermission

	for rows.Next() {
		var prm models.WorkPermission
		err := rows.Scan(
			&prm.Code,
			&prm.PerNum1,
			&prm.PerNum2,
			&prm.ScheduledTimeStart,
			&prm.ScheduledDateStart,
			&prm.Note,
			&prm.Status,
			&prm.EndTime,
			&prm.EndDate,
		)

		if err != nil {
			return nil, err
		}

		prms = append(prms, &prm)
	}

	return prms, nil
}

// Get returns all permissionsAll and error, if any
func (m *OracleDBRepo) GetWorkPermissionsAll() ([]*models.WorkPermissionAll, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	query := `select ID_ZAHTEVA,
			COALESCE(GRUPA, ''),
			COALESCE(UKLJUCENOST, ''),
			COALESCE(INT_PL, ''),
			COALESCE(to_char(BR1_Z_1GR), ''),
			COALESCE(BR2_Z_1GR, ''),
			COALESCE(BR_Z_RDC_2GR, ''),
			COALESCE(to_char(BR_SAG), ''),
			COALESCE(PL_DATUM_OD_Z, ''),
			COALESCE(PL_VREME_OD_Z, ''),
			COALESCE(PL_DATUM_DO_Z, ''),
			COALESCE(PL_VREME_DO_Z, ''),
			COALESCE(RUK_RADOVA, ''),
			COALESCE(OPIS_RADOVA, ''),
			COALESCE(NAPOMENA_VEZA, ''),
			COALESCE(SAG_USLOVI, ''),
			COALESCE(SAG_NAPOMENA_VEZA, ''),
			COALESCE(to_char(ID_DOZVOLE), ''),
			COALESCE(to_char(BR1_DOZ_1GR), ''),
			COALESCE(BR1_DOZ_2GR, ''),
			COALESCE(to_char(BR2_DOZ), ''),
			COALESCE(PL_VREME_OD, ''),
			COALESCE(PL_DATUM_OD, ''),
			COALESCE(ZAV_VREME, ''),
			COALESCE(to_char(ZAV_DATUM,'dd.mm.yyyy HH24:MI:SS'), ''),
			COALESCE(STATUS, '')
			  from synsoft_dozvole_sve
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var prms []*models.WorkPermissionAll

	for rows.Next() {
		var prm models.WorkPermissionAll
		err := rows.Scan(
			&prm.IdZahteva,
			&prm.Grupa,
			&prm.Ukljucenost,
			&prm.IntPl,
			&prm.Br1Z1Gr,
			&prm.Br2Z1Gr,
			&prm.BrZRDC2Gr,
			&prm.BrSag,
			&prm.PlDatumOdZ,
			&prm.PlVremeOdZ,
			&prm.PlDatumDoZ,
			&prm.PlVremeDoZ,
			&prm.RukRadova,
			&prm.OpisRadova,
			&prm.NapomenaVeza,
			&prm.SagUslovi,
			&prm.SagNapomenaVeza,
			&prm.IdDozvole,
			&prm.Br1Doz1Gr,
			&prm.Br1Doz2Gr,
			&prm.Br2Doz,
			&prm.PlVremeOd,
			&prm.PlVremeDo,
			&prm.ZavVreme,
			&prm.ZavDatum,
			&prm.Status,
		)

		if err != nil {
			return nil, err
		}

		/***** start part for objects ****/

		/*query1 := `select ips_id
			   from synsoft_dozvole_elrad
			   where id_dozvole=:1
			`
		var objs = []string{}

		rows1, err := m.DB.QueryContext(ctx, query1, prm.IdDozvole)
		if err != nil {
			return nil, err
		}

		for rows1.Next() {
			var ob string
			err := rows1.Scan(
				&ob,
			)

			if err != nil {
				return nil, err
			}
			objs = append(objs, ob)
		}
		rows1.Close()
		prm.EES = objs*/

		/***** end part for objects ******/

		prms = append(prms, &prm)
	}

	return prms, nil
}

// Get returns all DozvolaElrad and error, if any
func (m *OracleDBRepo) GetWorkPermissionElradAll() ([]*models.DozvolaElrad, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	query := `select ID_ZAHTEVA,
			ID_DOZVOLE,
			COALESCE(NA_SNAZI, ''),
			COALESCE(IPS_ID, ''),
			COALESCE(to_char(TIS_ID), ''),
			COALESCE(TIP, ''),
			COALESCE(ID_ELEMENTA, ''),
			COALESCE(OZNAKA_ELEMENTA, '')
			from synsoft_dozvole_elrad
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var prms []*models.DozvolaElrad

	for rows.Next() {
		var prm models.DozvolaElrad
		err := rows.Scan(
			&prm.IdZahteva,
			&prm.IdDozvole,
			&prm.NaSnazi,
			&prm.IpsId,
			&prm.TisId,
			&prm.Tip,
			&prm.IdElementa,
			&prm.OznakaElementa,
		)

		if err != nil {
			return nil, err
		}

		prms = append(prms, &prm)
	}

	return prms, nil
}

// Get returns all Request1Gr and error, if any
func (m *OracleDBRepo) GetRequest1Gr() ([]*models.Request1Gr, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select ID_ZAHTEVA,
			COALESCE(GRUPA, ''),
			COALESCE(UKLJUCENOST, ''),
			COALESCE(INT_PL, ''),
			COALESCE(to_char(BR1_Z_1GR), ''),
			COALESCE(BR2_Z_1GR, ''),
			COALESCE(PL_DATUM_OD_Z, ''),
			COALESCE(PL_VREME_OD_Z, ''),
			COALESCE(PL_DATUM_DO_Z, ''),
			COALESCE(PL_VREME_DO_Z, ''),
			COALESCE(RUK_RADOVA, ''),
			COALESCE(ELEMENTI, ''),
			COALESCE(OPIS_RADOVA, ''),
			COALESCE(NAPOMENA_VEZA, '')
			  from synsoft_zahtevi_1gr
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var prms []*models.Request1Gr

	for rows.Next() {
		var prm models.Request1Gr
		err := rows.Scan(
			&prm.IdZahteva,
			&prm.Grupa,
			&prm.Ukljucenost,
			&prm.IntPl,
			&prm.Br1Z1Gr,
			&prm.Br2Z1Gr,
			&prm.PlDatumOdZ,
			&prm.PlVremeOdZ,
			&prm.PlDatumDoZ,
			&prm.PlVremeDoZ,
			&prm.RukRadova,
			&prm.Elementi,
			&prm.OpisRadova,
			&prm.NapomenaVeza,
		)

		if err != nil {
			return nil, err
		}

		prms = append(prms, &prm)
	}

	return prms, nil
}

// Get returns all Request2Gr and error, if any
func (m *OracleDBRepo) GetRequest2Gr() ([]*models.Request2Gr, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select ID_ZAHTEVA,
			COALESCE(GRUPA, ''),
			COALESCE(UKLJUCENOST, ''),
			COALESCE(INT_PL, ''),
			COALESCE(BR_Z_RDC_2GR, ''),
			COALESCE(to_char(BR_SAG), ''),
			COALESCE(PL_DATUM_OD_Z, ''),
			COALESCE(PL_VREME_OD_Z, ''),
			COALESCE(PL_DATUM_DO_Z, ''),
			COALESCE(PL_VREME_DO_Z, ''),
			COALESCE(RUK_RADOVA, ''),
			COALESCE(ELEMENTI, ''),
			COALESCE(OPIS_RADOVA, ''),
			COALESCE(NAPOMENA_VEZA, ''),
			COALESCE(SAG_USLOVI, ''),
			COALESCE(SAG_NAPOMENA_VEZA, '')
			  from synsoft_zahtevi_2gr
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var prms []*models.Request2Gr

	for rows.Next() {
		var prm models.Request2Gr
		err := rows.Scan(
			&prm.IdZahteva,
			&prm.Grupa,
			&prm.Ukljucenost,
			&prm.IntPl,
			&prm.BrZRDC2Gr,
			&prm.BrSag,
			&prm.PlDatumOdZ,
			&prm.PlVremeOdZ,
			&prm.PlDatumDoZ,
			&prm.PlVremeDoZ,
			&prm.RukRadova,
			&prm.Elementi,
			&prm.OpisRadova,
			&prm.NapomenaVeza,
			&prm.SagUslovi,
			&prm.SagNapomenaVeza,
		)

		if err != nil {
			return nil, err
		}

		prms = append(prms, &prm)
	}

	return prms, nil
}

// Get returns all permissions and error, if any
func (m *OracleDBRepo) GetWorkInEENetwork() ([]*models.WorkInEENetwork, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select  COALESCE(to_char(MAX_RB), ''), COALESCE(to_char(BROJ), ''),  COALESCE(EE_ELEMENTI, ''), COALESCE(MESTO_RADA, ''), 
    		COALESCE(OPIS, ''), COALESCE(STATUS, ''), COALESCE(VREME, '') , COALESCE(VEZA, '')
			from SYNSOFT_RADOVI_U_MREZI
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var prms []*models.WorkInEENetwork

	for rows.Next() {
		var prm models.WorkInEENetwork
		err := rows.Scan(
			&prm.MaxNum,
			&prm.Num,
			&prm.EEElements,
			&prm.Workplace,
			&prm.Description,
			&prm.Status,
			&prm.Time,
			&prm.Link,
		)

		if err != nil {
			return nil, err
		}

		prms = append(prms, &prm)
	}

	return prms, nil
}

func (m *OracleDBRepo) GetWeather(year string) ([]*models.WeatherData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select NAME,
					LATITUDE,
					LONGITUDE,
					HEIGHT,
					TIMEZONEABBREVATION,
					UTCTIMEOFFSET,
					MODELRUNUTC,
					MODELRUNUPDATETIMEUTC,
					TIME,
					PICTOCODE,
					UVINDEX,
					TEMPERATUREMAX,
					TEMPERATUREMIN,
					TEMPERATUREMEAN,
					FELTTEMPERATUREMAX,
					FELTTEMPERATUREMIN,
					WINDDIRECTION,
					PRECIPITATIONPROBABILITY,
					RAINSPOT,
					PREDICTABILITYCLASS,
					PREDICTABILITY,
					PRECIPITATION,
					SNOWFRACTION,
					SEALEVELPRESSUREMAX,
					SEALEVELPRESSUREMIN,
					SEALEVELPRESSUREMEAN,
					WINDSPEEDMAX,
					WINDSPEEDMEAN,
					WINDSPEEDMIN,
					RELATIVEHUMIDITYMAX,
					RELATIVEHUMIDITYMIN,
					RELATIVEHUMIDITYMEAN,
					CONVECTIVEPRECIPITATION,
					PRECIPITATIONHOURS,
					HUMIDITYGREATER90HOURS,
					CREATION_DATE
			  from dwh.dwh_weather where substr(time,0,4)= :1`

	rows, err := m.DB.QueryContext(ctx, query, year)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var weatherData []*models.WeatherData

	for rows.Next() {
		var data models.WeatherData
		err := rows.Scan(
			&data.Name,
			&data.Latitude,
			&data.Longitude,
			&data.Height,
			&data.TimezoneAbbrevation,
			&data.UtcTimeoffset,
			&data.ModelrunUtc,
			&data.ModelrunUpdatetimeUtc,
			&data.Time,
			&data.Pictocode,
			&data.Uvindex,
			&data.TemperatureMax,
			&data.TemperatureMin,
			&data.TemperatureMean,
			&data.FelttemperatureMax,
			&data.FelttemperatureMin,
			&data.Winddirection,
			&data.PrecipitationProbability,
			&data.Rainspot,
			&data.PredictabilityClass,
			&data.Predictability,
			&data.Precipitation,
			&data.Snowfraction,
			&data.SealevelpressureMax,
			&data.SealevelpressureMin,
			&data.SealevelpressureMean,
			&data.WindspeedMax,
			&data.WindspeedMean,
			&data.WindspeedMin,
			&data.RelativehumidityMax,
			&data.RelativehumidityMin,
			&data.RelativehumidityMean,
			&data.ConvectivePrecipitation,
			&data.PrecipitationHours,
			&data.Humiditygreater90Hours,
			&data.CreationDate,
		)

		if err != nil {
			return nil, err
		}

		weatherData = append(weatherData, &data)
	}

	return weatherData, nil
}

func (m *OracleDBRepo) GetWeatherForecast() ([]*models.WeatherData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select NAME,
					LATITUDE,
					LONGITUDE,
					HEIGHT,
					TIMEZONEABBREVATION,
					UTCTIMEOFFSET,
					MODELRUNUTC,
					MODELRUNUPDATETIMEUTC,
					TIME,
					PICTOCODE,
					UVINDEX,
					TEMPERATUREMAX,
					TEMPERATUREMIN,
					TEMPERATUREMEAN,
					FELTTEMPERATUREMAX,
					FELTTEMPERATUREMIN,
					WINDDIRECTION,
					PRECIPITATIONPROBABILITY,
					RAINSPOT,
					PREDICTABILITYCLASS,
					PREDICTABILITY,
					PRECIPITATION,
					SNOWFRACTION,
					SEALEVELPRESSUREMAX,
					SEALEVELPRESSUREMIN,
					SEALEVELPRESSUREMEAN,
					WINDSPEEDMAX,
					WINDSPEEDMEAN,
					WINDSPEEDMIN,
					RELATIVEHUMIDITYMAX,
					RELATIVEHUMIDITYMIN,
					RELATIVEHUMIDITYMEAN,
					CONVECTIVEPRECIPITATION,
					PRECIPITATIONHOURS,
					HUMIDITYGREATER90HOURS,
					CREATION_DATE
			  from dwh.dwh_weather_forecast`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var weatherData []*models.WeatherData

	for rows.Next() {
		var data models.WeatherData
		err := rows.Scan(
			&data.Name,
			&data.Latitude,
			&data.Longitude,
			&data.Height,
			&data.TimezoneAbbrevation,
			&data.UtcTimeoffset,
			&data.ModelrunUtc,
			&data.ModelrunUpdatetimeUtc,
			&data.Time,
			&data.Pictocode,
			&data.Uvindex,
			&data.TemperatureMax,
			&data.TemperatureMin,
			&data.TemperatureMean,
			&data.FelttemperatureMax,
			&data.FelttemperatureMin,
			&data.Winddirection,
			&data.PrecipitationProbability,
			&data.Rainspot,
			&data.PredictabilityClass,
			&data.Predictability,
			&data.Precipitation,
			&data.Snowfraction,
			&data.SealevelpressureMax,
			&data.SealevelpressureMin,
			&data.SealevelpressureMean,
			&data.WindspeedMax,
			&data.WindspeedMean,
			&data.WindspeedMin,
			&data.RelativehumidityMax,
			&data.RelativehumidityMin,
			&data.RelativehumidityMean,
			&data.ConvectivePrecipitation,
			&data.PrecipitationHours,
			&data.Humiditygreater90Hours,
			&data.CreationDate,
		)

		if err != nil {
			return nil, err
		}

		weatherData = append(weatherData, &data)
	}

	return weatherData, nil
}

func (m *OracleDBRepo) GetWeatherHistory(year string) ([]*models.WeatherDataHistory, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select NAME,
					LATITUDE,
					LONGITUDE,
					HEIGHT,
					TIMEZONEABBREVATION,
					UTCTIMEOFFSET,
					MODELRUNUTC,
					MODELRUNUPDATETIMEUTC,
					TIME,
					COALESCE(to_char(PICTOCODE), ''),
					COALESCE(to_char(UVINDEX), ''),
					COALESCE(to_char(TEMPERATUREMAX), ''),
					COALESCE(to_char(TEMPERATUREMIN), ''),
					COALESCE(to_char(TEMPERATUREMEAN), ''),
					COALESCE(to_char(FELTTEMPERATUREMAX), ''),
					COALESCE(to_char(FELTTEMPERATUREMIN), ''),
					COALESCE(to_char(WINDDIRECTION), ''),
					COALESCE(to_char(PRECIPITATIONPROBABILITY), ''),
					COALESCE(to_char(RAINSPOT), ''),
					COALESCE(to_char(PREDICTABILITYCLASS), ''),
					COALESCE(to_char(PREDICTABILITY), ''),
					COALESCE(to_char(PRECIPITATION), ''),
					COALESCE(to_char(SNOWFRACTION), ''),
					COALESCE(to_char(SEALEVELPRESSUREMAX), ''),
					COALESCE(to_char(SEALEVELPRESSUREMIN), ''),
					COALESCE(to_char(SEALEVELPRESSUREMEAN), ''),
					COALESCE(to_char(WINDSPEEDMAX), ''),
					COALESCE(to_char(WINDSPEEDMEAN), ''),
					COALESCE(to_char(WINDSPEEDMIN), ''),
					COALESCE(to_char(RELATIVEHUMIDITYMAX), ''),
					COALESCE(to_char(RELATIVEHUMIDITYMIN), ''),
					COALESCE(to_char(RELATIVEHUMIDITYMEAN), ''),
					COALESCE(to_char(CONVECTIVEPRECIPITATION), ''),
					COALESCE(to_char(PRECIPITATIONHOURS), ''),
					COALESCE(to_char(HUMIDITYGREATER90HOURS), ''),
					COALESCE(to_char(CREATION_DATE), '')
			  from dwh.dwh_weather_history where ('20'|| substr(time,length(time)-1,2)) = :1`

	rows, err := m.DB.QueryContext(ctx, query, year)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var weatherDataHistory []*models.WeatherDataHistory

	for rows.Next() {
		var data models.WeatherDataHistory
		err := rows.Scan(
			&data.Name,
			&data.Latitude,
			&data.Longitude,
			&data.Height,
			&data.TimezoneAbbrevation,
			&data.UtcTimeoffset,
			&data.ModelrunUtc,
			&data.ModelrunUpdatetimeUtc,
			&data.Time,
			&data.Pictocode,
			&data.Uvindex,
			&data.TemperatureMax,
			&data.TemperatureMin,
			&data.TemperatureMean,
			&data.FelttemperatureMax,
			&data.FelttemperatureMin,
			&data.Winddirection,
			&data.PrecipitationProbability,
			&data.Rainspot,
			&data.PredictabilityClass,
			&data.Predictability,
			&data.Precipitation,
			&data.Snowfraction,
			&data.SealevelpressureMax,
			&data.SealevelpressureMin,
			&data.SealevelpressureMean,
			&data.WindspeedMax,
			&data.WindspeedMean,
			&data.WindspeedMin,
			&data.RelativehumidityMax,
			&data.RelativehumidityMin,
			&data.RelativehumidityMean,
			&data.ConvectivePrecipitation,
			&data.PrecipitationHours,
			&data.Humiditygreater90Hours,
			&data.CreationDate,
		)

		if err != nil {
			return nil, err
		}

		weatherDataHistory = append(weatherDataHistory, &data)
	}

	return weatherDataHistory, nil
}

func (m *OracleDBRepo) GetPermissions1(year string) ([]*models.Permission, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select RBR_PK,
				BR_ZAHTEVA_FK,
				SAP_SIFRA,
				BROJ_ISK,
				BR_DOZVOLE,
				TIP_RADOVA,
				EE_OBJEKAT,
				TIP,
				OZNAKA,
				STATUS_EL,
				IZUZEV,
				NAPOMENA,
				DOZ_IZDAO,
				DOZ_PRIMIO,
				DAT_PRIJEMA_DOZ,
				VREME_PRIJEMA_DOZ,
				STASTUS_DOZ,
				NAPOMENA_ZAV_RAD,
				DOZ_ZAV_IZDAO,
				DOZ_ZAV_PRIMIO,
				DAT_ZAV_RADOVA,
				VREME_ZAV_RAD
			  from dozvole_1 where substr(dat_prijema_doz,length(dat_prijema_doz)-3,4) = :1
			  or  substr(dat_zav_radova,length(dat_zav_radova)-3,4)= :2`

	rows, err := m.DB.QueryContext(ctx, query, year, year)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var permissions []*models.Permission

	for rows.Next() {
		var data models.Permission
		err := rows.Scan(
			&data.RbrPk,
			&data.BrZahtevaFK,
			&data.SapSifra,
			&data.BrojIsk,
			&data.BrDozvole,
			&data.TipRadova,
			&data.EeObjekat,
			&data.Tip,
			&data.Oznaka,
			&data.StatusEl,
			&data.Izuzev,
			&data.Napomena,
			&data.DozIzdao,
			&data.DozPrimio,
			&data.DatPrijemaDoz,
			&data.VremePrijemaDoz,
			&data.StatusDoz,
			&data.NapomenaZavRad,
			&data.DozZavIzdao,
			&data.DozZavPrimio,
			&data.DatZavRadova,
			&data.VremeZavRad,
		)

		if err != nil {
			return nil, err
		}

		permissions = append(permissions, &data)
	}

	return permissions, nil
}

func (m *OracleDBRepo) GetPermissions23(year string) ([]*models.Permission, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select RBR_PK,
				BR_ZAHTEVA_FK,
				SAP_SIFRA,
				BROJ_ISK,
				BR_DOZVOLE,
				TIP_RADOVA,
				EE_OBJEKAT,
				TIP,
				OZNAKA,
				STATUS_EL,
				IZUZEV,
				NAPOMENA,
				DOZ_IZDAO,
				DOZ_PRIMIO,
				DAT_PRIJEMA_DOZ,
				VREME_PRIJEMA_DOZ,
				STASTUS_DOZ,
				NAPOMENA_ZAV_RAD,
				DOZ_ZAV_IZDAO,
				DOZ_ZAV_PRIMIO,
				DAT_ZAV_RADOVA,
				VREME_ZAV_RAD
			  from dozvole_2_3 where substr(dat_prijema_doz,length(dat_prijema_doz)-3,4) = :1
			  or  substr(dat_zav_radova,length(dat_zav_radova)-3,4)= :2`

	rows, err := m.DB.QueryContext(ctx, query, year, year)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var permissions []*models.Permission

	for rows.Next() {
		var data models.Permission
		err := rows.Scan(
			&data.RbrPk,
			&data.BrZahtevaFK,
			&data.SapSifra,
			&data.BrojIsk,
			&data.BrDozvole,
			&data.TipRadova,
			&data.EeObjekat,
			&data.Tip,
			&data.Oznaka,
			&data.StatusEl,
			&data.Izuzev,
			&data.Napomena,
			&data.DozIzdao,
			&data.DozPrimio,
			&data.DatPrijemaDoz,
			&data.VremePrijemaDoz,
			&data.StatusDoz,
			&data.NapomenaZavRad,
			&data.DozZavIzdao,
			&data.DozZavPrimio,
			&data.DatZavRadova,
			&data.VremeZavRad,
		)

		if err != nil {
			return nil, err
		}

		permissions = append(permissions, &data)
	}

	return permissions, nil
}

func (m *OracleDBRepo) GetRequests1(year string) ([]*models.Request, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select RBR_PK,
			RCO,
			COALESCE(to_char(BROJ_ISK), ''),
			BR_ZAHTEVA,
			BROJ_KPS,
			GRUPA,
			INT,
			SAP_SIFRA,
			EE_OBJEKAT,
			TIP,
			OZNAKA,
			OPIS,
			PL_DATUM_OD,
			PL_VREME_OD,
			PL_DATUM_DO,
			PL_VREME_DO,
			NODOB,
			TIP_RADOVA,
			USLOVI,
			NAP_VEZA,
			POD_ZAHTEVA,
			ISK_ODOBRIO
			from zahtevi_1 where substr(PL_DATUM_OD,length(PL_DATUM_OD)-3,4) = :1
			or  substr(PL_DATUM_DO,length(PL_DATUM_DO)-3,4)= :2`

	rows, err := m.DB.QueryContext(ctx, query, year, year)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var requests []*models.Request

	for rows.Next() {
		var data models.Request
		err := rows.Scan(
			&data.RbrPk,
			&data.Rco,
			&data.BrojIsk,
			&data.BrZahteva,
			&data.BrojKps,
			&data.Grupa,
			&data.Int,
			&data.SapSifra,
			&data.EeObjekat,
			&data.Tip,
			&data.Oznaka,
			&data.Opis,
			&data.PlDatumOd,
			&data.PlVremeOd,
			&data.PlDatumDo,
			&data.PlVremeDo,
			&data.Nodob,
			&data.TipRadova,
			&data.Uslovi,
			&data.NapVeza,
			&data.PodZahteva,
			&data.IskOdobrio,
		)

		if err != nil {
			return nil, err
		}

		requests = append(requests, &data)
	}

	return requests, nil
}

func (m *OracleDBRepo) GetRequests23(year string) ([]*models.Request, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select RBR_PK,
			RCO,
			COALESCE(to_char(BROJ_ISK), ''),
			BR_ZAHTEVA,
			BROJ_KPS,
			GRUPA,
			INT,
			SAP_SIFRA,
			EE_OBJEKAT,
			TIP,
			OZNAKA,
			OPIS,
			PL_DATUM_OD,
			PL_VREME_OD,
			PL_DATUM_DO,
			PL_VREME_DO,
			NODOB,
			TIP_RADOVA,
			USLOVI,
			NAP_VEZA,
			POD_ZAHTEVA,
			ISK_ODOBRIO
			from zahtevi_2_3 where substr(PL_DATUM_OD,length(PL_DATUM_OD)-3,4) = :1
			or  substr(PL_DATUM_DO,length(PL_DATUM_DO)-3,4)= :2`

	rows, err := m.DB.QueryContext(ctx, query, year, year)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var requests []*models.Request

	for rows.Next() {
		var data models.Request
		err := rows.Scan(
			&data.RbrPk,
			&data.Rco,
			&data.BrojIsk,
			&data.BrZahteva,
			&data.BrojKps,
			&data.Grupa,
			&data.Int,
			&data.SapSifra,
			&data.EeObjekat,
			&data.Tip,
			&data.Oznaka,
			&data.Opis,
			&data.PlDatumOd,
			&data.PlVremeOd,
			&data.PlDatumDo,
			&data.PlVremeDo,
			&data.Nodob,
			&data.TipRadova,
			&data.Uslovi,
			&data.NapVeza,
			&data.PodZahteva,
			&data.IskOdobrio,
		)

		if err != nil {
			return nil, err
		}

		requests = append(requests, &data)
	}

	return requests, nil
}

func (m *OracleDBRepo) GetOutages(year string) ([]*models.Outage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select to_char(DATIZV,'dd.mm.yyyy'),
		to_char(VREPOC,'dd.mm.yyyy HH24:MI:SS'),
		to_char(VREZAV,'dd.mm.yyyy HH24:MI:SS'),
		TRAJ,
		IPS_ID,
		TIPOB,
		OPIS,
		IME_DALEKOVODA,
		POLJE_TRAFO,
		ORG1,
		ORG2,
		NAZVRPD,
		UZROK,
		VRM_USL,
		TEKST,
		COALESCE(to_char(ID_STAVKE), ''),
		COALESCE(to_char(ID_SEQ), '')
 		from DWH_ISPADI_V  v
 		where TO_CHAR( v.datizv, 'yyyy' )= :1`

	rows, err := m.DB.QueryContext(ctx, query, year)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var o []*models.Outage

	for rows.Next() {
		var data models.Outage
		err := rows.Scan(
			&data.Datizv,
			&data.Vrepoc,
			&data.Vrezav,
			&data.Traj,
			&data.IpsId,
			&data.TipOb,
			&data.Opis,
			&data.ImeDalekovoda,
			&data.PoljeTrafo,
			&data.Org1,
			&data.Org2,
			&data.Nazvrpd,
			&data.Uzrok,
			&data.VrmUsl,
			&data.Tekst,
			&data.IdStavke,
			&data.IdSeq,
		)

		if err != nil {
			return nil, err
		}

		o = append(o, &data)
	}

	return o, nil
}

func (m *OracleDBRepo) GetTransmissionLineFailure(ipsId string, vremeOd string, vremeDo string) ([]*models.GisOutage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	query := `select to_char(v.VREPOC,'dd.mm.yyyy HH24:MI:SS'),
		to_char(v.VREZAV,'dd.mm.yyyy HH24:MI:SS'),
		v.TRAJ,
		v.IPS_ID,
		v.TIPOB,
		v.OPIS,
		v.IME_DALEKOVODA,
		v.POLJE_TRAFO,
		v.ORG1,
		v.ORG2,
		v.NAZVRPD,
		v.UZROK,
		v.VRM_USL,
		v.TEKST
 		from PGI.PI_BI_ISPADI_V v
 		where v.ips_id= :1
		and v.VREPOC is not null
		and v.datizv between to_date(:2,'dd.mm.yyyy') and to_date(:3,'dd.mm.yyyy')
		and v.datizv = (select max(datizv) from PGI.PI_DD_ISPADI_MV dr where (dr.id_seq=v.id_seq or v.id_seq is null) and (dr.id_stavke=v.id_stavke or v.id_stavke is null))`

	rows, err := m.DB.QueryContext(ctx, query, ipsId, vremeOd, vremeDo)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var d []*models.GisOutage

	for rows.Next() {
		var data models.GisOutage
		err := rows.Scan(
			&data.Vrepoc,
			&data.Vrezav,
			&data.Traj,
			&data.IpsId,
			&data.TipOb,
			&data.Opis,
			&data.ImeDalekovoda,
			&data.PoljeTrafo,
			&data.Org1,
			&data.Org2,
			&data.Nazvrpd,
			&data.Uzrok,
			&data.VrmUsl,
			&data.Tekst,
		)

		if err != nil {
			return nil, err
		}

		d = append(d, &data)
	}

	return d, nil
}

func (m *OracleDBRepo) GetExclusions(year string) ([]*models.Exclusion, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	query := `select to_char(DATIZV,'dd.mm.yyyy'),
		to_char(VREPOC,'dd.mm.yyyy HH24:MI:SS'),
		to_char(VREZAV,'dd.mm.yyyy HH24:MI:SS'),
		TRAJ,
		IPS_ID,
		TIPOB,
		OPIS,
		IME_DALEKOVODA,
		POLJE_TRAFO,
		ORG1,
		ORG2,
		NAZVRPD,
		RAZLOG,
		TEKST
 		from DWH_ISKLJUCENJA_V  v
 		where TO_CHAR( v.datizv, 'yyyy' )= :1`

	rows, err := m.DB.QueryContext(ctx, query, year)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var d []*models.Exclusion

	for rows.Next() {
		var data models.Exclusion
		err := rows.Scan(
			&data.Datizv,
			&data.Vrepoc,
			&data.Vrezav,
			&data.Traj,
			&data.IpsId,
			&data.TipOb,
			&data.Opis,
			&data.ImeDalekovoda,
			&data.PoljeTrafo,
			&data.Org1,
			&data.Org2,
			&data.Nazvrpd,
			&data.Razlog,
			&data.Tekst,
		)

		if err != nil {
			return nil, err
		}

		d = append(d, &data)
	}

	return d, nil
}

func (m *OracleDBRepo) GetTransmissionLineOutage(ipsId string, vremeOd string, vremeDo string) ([]*models.GisExclusion, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	query := `select to_char(v.DATIZV,'dd.mm.yyyy'),
		to_char(v.VREPOC,'dd.mm.yyyy HH24:MI:SS'),
		to_char(v.VREZAV,'dd.mm.yyyy HH24:MI:SS'),
		v.TRAJ,
		v.IPS_ID,
		v.TIPOB,
		v.OPIS,
		v.IME_DALEKOVODA,
		v.POLJE_TRAFO,
		v.ORG1,
		v.ORG2,
		v.NAZVRPD,
		v.RAZLOG,
		v.TEKST
 		from PGI.PI_BI_ISKLJUCENJA_V  v
 		where v.ips_id= :1
		and v.VREPOC is not null
		and v.datizv between to_date(:2,'dd.mm.yyyy') and to_date(:3,'dd.mm.yyyy')`

	rows, err := m.DB.QueryContext(ctx, query, ipsId, vremeOd, vremeDo)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var d []*models.GisExclusion

	for rows.Next() {
		var data models.GisExclusion
		err := rows.Scan(
			&data.Datizv,
			&data.Vrepoc,
			&data.Vrezav,
			&data.Traj,
			&data.IpsId,
			&data.TipOb,
			&data.Opis,
			&data.ImeDalekovoda,
			&data.PoljeTrafo,
			&data.Org1,
			&data.Org2,
			&data.Nazvrpd,
			&data.Razlog,
			&data.Tekst,
		)

		if err != nil {
			return nil, err
		}

		d = append(d, &data)
	}

	return d, nil
}

func (m *OracleDBRepo) GetPlans(year string) ([]*models.Plan, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select COALESCE(to_char(ID_POG_ODR), ''),
				ID_SAP_F_LOK,
				ID_IPS_F_LOK,
				OPIS,
				COALESCE(to_char(ID_POG_PLAN), ''),
				COALESCE(to_char(PL_ODR), ''),
				TKS_ST_OD,
				TIP_NALOGA,
				to_char(DATUM_POC,'dd.mm.yyyy'),
				to_char(DATUM_ZAV,'dd.mm.yyyy'),
				COALESCE(to_char(ID), '')
				from PLAN_O_VNP_V
				where to_char(DATUM_POC,'yyyy')= :1
				or to_char(DATUM_ZAV,'dd.mm.yyyy')= :2`

	rows, err := m.DB.QueryContext(ctx, query, year, year)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var p []*models.Plan

	for rows.Next() {
		var data models.Plan
		err := rows.Scan(
			&data.IdPogOdr,
			&data.IdSapFLok,
			&data.IdIPSFLok,
			&data.Opis,
			&data.IdPogPlan,
			&data.PlOdr,
			&data.TksStOd,
			&data.TipNaloga,
			&data.DatumPoc,
			&data.DatumZav,
			&data.Id,
		)

		if err != nil {
			return nil, err
		}

		p = append(p, &data)
	}

	return p, nil
}

func (m *OracleDBRepo) GetUnopenedPermitForDay(day string, org string) ([]*models.UnopenedPermit, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select  COALESCE(to_char(BROJ_ISK), ''),
			COALESCE(BROJ_ISK2, ''),
			COALESCE(BROJ_ISK_RDC, ''),
			RBR_ISK,
			COALESCE(RAZLOG_O, ''),
			to_char(DATUM_O,'dd.mm.yyyy HH24:MI:SS'),
			COALESCE(to_char(RBR), ''),
			COALESCE(KORISNIK, ''),
			COALESCE(OPIS, ''),
			COALESCE(LICE_RUK, '')
 			from TED.SYNSOFT_NEOTVORENE_DOZ
 			WHERE to_char(DATUM_O,'dd.mm.yyyy')= :1
			AND KORISNIK= :2`

	rows, err := m.DB.QueryContext(ctx, query, day, org)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var p []*models.UnopenedPermit

	for rows.Next() {
		var data models.UnopenedPermit
		err := rows.Scan(
			&data.BrojIsk,
			&data.BrojIsk2,
			&data.BrojIskRDC,
			&data.RbrIsk,
			&data.RazlogO,
			&data.DatumO,
			&data.Rbr,
			&data.Korisnik,
			&data.Opis,
			&data.LiceRuk,
		)

		if err != nil {
			return nil, err
		}

		p = append(p, &data)
	}

	return p, nil
}

// Authenticate authenticates a user
func (m *OracleDBRepo) Authenticate(username, testPassword string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	/*var id int
	var hashedPassword string*/

	var user models.User

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

func (m *OracleDBRepo) GetUserByUsername(username string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select id, username, password from tis_services_users where username = :1`

	var user models.User
	row := m.DB.QueryRowContext(ctx, query, username)

	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Password,
	)

	if err != nil {
		return nil, err
	}
	query = `select ru.id,RU.ID_USER, RU.ID_ROLE, R.CODE, R.NAME
	from tis_services_role_user ru, tis_services_roles r
	where RU.ID_USER =:1
	and ru.id_role = r.id
	`

	var roles []models.UserRole
	rows, _ := m.DB.QueryContext(ctx, query, user.ID)
	defer rows.Close()

	for rows.Next() {
		var r models.UserRole
		err := rows.Scan(
			&r.ID,
			&r.IdUser,
			&r.IdRole,
			&r.RoleCode,
			&r.RoleName,
		)

		if err != nil {
			return nil, err
		}
		roles = append(roles, r)
	}
	user.UserRole = roles

	return &user, nil
}

func (m *OracleDBRepo) GetUserByID(id int) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select id, username, password from tis_services_users where id = :1`

	var user models.User
	row := m.DB.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Password,
	)

	if err != nil {
		return nil, err
	}

	query = `select ru.id,RU.ID_USER, RU.ID_ROLE, R.CODE, R.NAME
	from tis_services_role_user ru, tis_services_roles r
	where RU.ID_USER =:1
	and ru.id_role = r.id
	`

	var roles []models.UserRole
	rows, _ := m.DB.QueryContext(ctx, query, id)
	defer rows.Close()

	for rows.Next() {
		var r models.UserRole
		err := rows.Scan(
			&r.ID,
			&r.IdUser,
			&r.IdRole,
			&r.RoleCode,
			&r.RoleName,
		)

		if err != nil {
			return nil, err
		}
		roles = append(roles, r)
	}
	user.UserRole = roles

	return &user, nil
}

func (m *OracleDBRepo) InsertPiPiDDNIsklj(pipiddn models.PiPiDDNIsklj) error {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var status int
	var message string

	query := `begin  ddn.synsoft.pi_pi_ddn_iskljucenje_insert(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17); end;`
	//var int status
	//var string message
	_, err := m.DB.ExecContext(ctx, query,
		pipiddn.DatSmene,
		8,
		pipiddn.TipMan,
		pipiddn.IdTipob,
		pipiddn.ObId,
		pipiddn.TrafoId,
		pipiddn.Vrepoc,
		pipiddn.Vrezav,
		pipiddn.IdSGrraz,
		pipiddn.IdSRazlog,
		pipiddn.ManTekst,
		pipiddn.IdSNap,
		pipiddn.P2TrafId,
		pipiddn.KorUneo,
		pipiddn.SynsoftId,
		sql.Out{Dest: &status},
		sql.Out{Dest: &message},
	)

	if err != nil {
		log.Println(err)
		return err
	}
	//fmt.Println(pipiddn.TipMan)
	//fmt.Println(pipiddn.DatSmene)
	//fmt.Println(status)
	//fmt.Println(message)
	if status != 0 {
		return errors.New(message)
	} else {
		return nil
	}
}

func (m *OracleDBRepo) InsertPiPiDDNIskljP(pipiddn models.PiPiDDNIsklj) error {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var status int
	var message string

	query := `begin  ddn.synsoft.p_pi_pi_ddn_iskljucenje_insert(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17); end;`
	//var int status
	//var string message
	_, err := m.DB.ExecContext(ctx, query,
		pipiddn.DatSmene,
		8,
		pipiddn.TipMan,
		pipiddn.IdTipob,
		pipiddn.ObId,
		pipiddn.TrafoId,
		pipiddn.Vrepoc,
		pipiddn.Vrezav,
		pipiddn.IdSGrraz,
		pipiddn.IdSRazlog,
		pipiddn.ManTekst,
		pipiddn.IdSNap,
		pipiddn.P2TrafId,
		pipiddn.KorUneo,
		pipiddn.SynsoftId,
		sql.Out{Dest: &status},
		sql.Out{Dest: &message},
	)

	if err != nil {
		log.Println(err)
		return err
	}
	//fmt.Println(pipiddn.TipMan)
	//fmt.Println(pipiddn.DatSmene)
	//fmt.Println(status)
	//fmt.Println(message)
	if status != 0 {
		return errors.New(message)
	} else {
		return nil
	}
}

func (m *OracleDBRepo) UpdatePiPiDDNIsklj(pipiddn models.PiPiDDNIsklj) error {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var status int
	var message string

	query := `begin  ddn.synsoft.pi_pi_ddn_iskljucenje_update(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17); end;`
	//var int status
	//var string message
	_, err := m.DB.ExecContext(ctx, query,
		pipiddn.DatSmene,
		pipiddn.IdSMrc,
		pipiddn.TipMan,
		pipiddn.IdTipob,
		pipiddn.ObId,
		pipiddn.TrafoId,
		pipiddn.Vrepoc,
		pipiddn.Vrezav,
		pipiddn.IdSGrraz,
		pipiddn.IdSRazlog,
		pipiddn.ManTekst,
		pipiddn.IdSNap,
		pipiddn.P2TrafId,
		pipiddn.KorUneo,
		pipiddn.SynsoftId,
		sql.Out{Dest: &status},
		sql.Out{Dest: &message},
	)

	if err != nil {
		log.Println(err)
		return err
	}
	//fmt.Println(pipiddn.TipMan)
	//fmt.Println(pipiddn.DatSmene)
	//fmt.Println(status)
	//fmt.Println(message)

	if status != 0 {
		return errors.New(message)
	} else {
		return nil
	}
}

func (m *OracleDBRepo) UpdatePiPiDDNIskljP(pipiddn models.PiPiDDNIsklj) error {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var status int
	var message string

	query := `begin  ddn.synsoft.p_pi_pi_ddn_iskljucenje_update(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17); end;`
	//var int status
	//var string message
	_, err := m.DB.ExecContext(ctx, query,
		pipiddn.DatSmene,
		pipiddn.IdSMrc,
		pipiddn.TipMan,
		pipiddn.IdTipob,
		pipiddn.ObId,
		pipiddn.TrafoId,
		pipiddn.Vrepoc,
		pipiddn.Vrezav,
		pipiddn.IdSGrraz,
		pipiddn.IdSRazlog,
		pipiddn.ManTekst,
		pipiddn.IdSNap,
		pipiddn.P2TrafId,
		pipiddn.KorUneo,
		pipiddn.SynsoftId,
		sql.Out{Dest: &status},
		sql.Out{Dest: &message},
	)

	if err != nil {
		log.Println(err)
		return err
	}
	//fmt.Println(pipiddn.TipMan)
	//fmt.Println(pipiddn.DatSmene)
	//fmt.Println(status)
	//fmt.Println(message)

	if status != 0 {
		return errors.New(message)
	} else {
		return nil
	}
}

func (m *OracleDBRepo) GetAllPiPiDDN() ([]*models.PiPiDDN, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select to_char(DATIZV, 'dd.mm.yyyy'),
				COALESCE(to_char(ID_S_MRC), ''),
				COALESCE(to_char(ID_S_TIPD), ''),
				COALESCE(to_char(ID_S_VRPD), ''),
				COALESCE(to_char(ID_TIPOB), ''),
				COALESCE(to_char(OB_ID), ''),
				COALESCE(to_char(TRAFO_ID), ''),
				to_char(VREPOC, 'dd.mm.yyyy HH24:MI:SS'),
				POC_PP,
				to_char(VREZAV, 'dd.mm.yyyy HH24:MI:SS'),
				ZAV_PP,
				COALESCE(to_char(ID1_S_GRUZR), ''),
                COALESCE(to_char(ID1_S_UZROK), ''),
				COALESCE(to_char(ID_S_GRRAZ), ''),
				COALESCE(to_char(ID_S_RAZLOG), ''),
				COALESCE(to_char(SNAGA), ''),
				OPIS,
				COALESCE(to_char(ID_S_NAP), ''),
				COALESCE(to_char(P2_TRAF_ID), ''),
				COALESCE(PGI_KOR, ''),
                COALESCE(STATUS, ''),
				to_char(DATPRI, 'dd.mm.yyyy HH24:MI:SS'),
				COALESCE(to_char(ID_Z_DSDF_GL1), ''),
                COALESCE(to_char(ID_Z_KVAR_GL1), ''),
                COALESCE(to_char(ID_Z_RAPU_GL1), ''),
                COALESCE(to_char(ID_Z_PRST_GL1), ''),
                COALESCE(to_char(ID_Z_ZMSP_GL1), ''),
                COALESCE(to_char(ID_Z_UZMS_GL1), ''),
                COALESCE(to_char(Z_LOKK_GL1), ''),
                COALESCE(to_char(ID_Z_DSDF_GL2), ''),
                COALESCE(to_char(ID_Z_KVAR_GL2), ''),
                COALESCE(to_char(ID_Z_RAPU_GL2), ''),
                COALESCE(to_char(ID_Z_PRST_GL2), ''),
                COALESCE(to_char(ID_Z_ZMSP_GL2), ''),
                COALESCE(to_char(ID_Z_UZMS_GL2), ''),
                COALESCE(to_char(Z_LOKK_GL2), ''),
                COALESCE(to_char(ID_Z_PREK_VN), ''),
                COALESCE(to_char(ID_Z_DIS_REZ), ''),
                COALESCE(to_char(ID_Z_KVAR_REZ), ''),
                COALESCE(to_char(ID_Z_PRST_REZ), ''),
                COALESCE(to_char(ID_Z_ZMSP_REZ), ''),
                COALESCE(to_char(ID_Z_NEL1), ''),
                COALESCE(to_char(ID_Z_NEL2), ''),
                COALESCE(to_char(ID_Z_NEL3), ''),
                COALESCE(to_char(ID_Z_PREK_NN), ''),
                COALESCE(to_char(ID_Z_SABZ_SAB), ''),
                COALESCE(to_char(ID_Z_OTPR_SAB), ''),
                COALESCE(to_char(ID_S_VREM_USL), ''),
                COALESCE(UZROK_TEKST, ''),
                COALESCE(to_char(ID_Z_JPS_VN), ''),
                COALESCE(to_char(ID_Z_JPS_NN), ''),
                COALESCE(POSL_TEKST, ''),
                COALESCE(to_char(ID_Z_TELE_POC_GL1), ''),
                COALESCE(to_char(ID_Z_TELE_KRAJ_GL1), ''),
                COALESCE(to_char(ID_Z_TELE_POC_GL2), ''),
                COALESCE(to_char(ID_Z_TELE_KRAJ_GL2), ''),
                COALESCE(SYNSOFT_ID, '')
				from PI_PI_DDN_S`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var p []*models.PiPiDDN

	for rows.Next() {
		var pipiddn models.PiPiDDN
		err := rows.Scan(
			&pipiddn.Datizv,
			&pipiddn.IdSMrc,
			&pipiddn.IdSTipd,
			&pipiddn.IdSVrpd,
			&pipiddn.IdTipob,
			&pipiddn.ObId,
			&pipiddn.TrafoId,
			&pipiddn.Vrepoc,
			&pipiddn.PocPP,
			&pipiddn.Vrezav,
			&pipiddn.ZavPP,
			&pipiddn.Id1SGruzr,
			&pipiddn.Id1SUzrok,
			&pipiddn.IdSGrraz,
			&pipiddn.IdSRazlog,
			&pipiddn.Snaga,
			&pipiddn.Opis,
			&pipiddn.IdSNap,
			&pipiddn.P2TrafId,
			&pipiddn.KorUneo,
			&pipiddn.Status,
			&pipiddn.Datpri,
			&pipiddn.IdZDsdfGL1,
			&pipiddn.IdZKvarGL1,
			&pipiddn.IdZRapuGL1,
			&pipiddn.IdZPrstGL1,
			&pipiddn.IdZZmspGL1,
			&pipiddn.IdZUzmsGL1,
			&pipiddn.ZLokkGL1,
			&pipiddn.IdZDsdfGL2,
			&pipiddn.IdZKvarGL2,
			&pipiddn.IdZRapuGL2,
			&pipiddn.IdZPrstGL2,
			&pipiddn.IdZZmspGL2,
			&pipiddn.IdZUzmsGL2,
			&pipiddn.ZLokkGL2,
			&pipiddn.IdZPrekVN,
			&pipiddn.IdZDisREZ,
			&pipiddn.IdZKvarREZ,
			&pipiddn.IdZPrstREZ,
			&pipiddn.IdZZmspREZ,
			&pipiddn.IdZNel1,
			&pipiddn.IdZNel2,
			&pipiddn.IdZNel3,
			&pipiddn.IdZPrekNN,
			&pipiddn.IdZSabzSAB,
			&pipiddn.IdZOtprSAB,
			&pipiddn.IdSVremUSL,
			&pipiddn.UzrokTekst,
			&pipiddn.IdZJpsVN,
			&pipiddn.IdZJpsNN,
			&pipiddn.PoslTekst,
			&pipiddn.IdZTelePocGL1,
			&pipiddn.IdZTeleKrajGL1,
			&pipiddn.IdZTelePocGL2,
			&pipiddn.IdZTeleKrajGL2,
			&pipiddn.SynsoftId,
		)

		if err != nil {
			return nil, err
		}

		p = append(p, &pipiddn)
	}

	return p, nil
}

func (m *OracleDBRepo) GetAllPiPiDDNP() ([]*models.PiPiDDN, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select to_char(DATIZV, 'dd.mm.yyyy'),
				COALESCE(to_char(ID_S_MRC), ''),
				COALESCE(to_char(ID_S_TIPD), ''),
				COALESCE(to_char(ID_S_VRPD), ''),
				COALESCE(to_char(ID_TIPOB), ''),
				COALESCE(to_char(OB_ID), ''),
				COALESCE(to_char(TRAFO_ID), ''),
				to_char(VREPOC, 'dd.mm.yyyy HH24:MI:SS'),
				POC_PP,
				to_char(VREZAV, 'dd.mm.yyyy HH24:MI:SS'),
				ZAV_PP,
				COALESCE(to_char(ID1_S_GRUZR), ''),
                COALESCE(to_char(ID1_S_UZROK), ''),
				COALESCE(to_char(ID_S_GRRAZ), ''),
				COALESCE(to_char(ID_S_RAZLOG), ''),
				COALESCE(to_char(SNAGA), ''),
				OPIS,
				COALESCE(to_char(ID_S_NAP), ''),
				COALESCE(to_char(P2_TRAF_ID), ''),
				COALESCE(PGI_KOR, ''),
                COALESCE(STATUS, ''),
				to_char(DATPRI, 'dd.mm.yyyy HH24:MI:SS'),
				COALESCE(to_char(ID_Z_DSDF_GL1), ''),
                COALESCE(to_char(ID_Z_KVAR_GL1), ''),
                COALESCE(to_char(ID_Z_RAPU_GL1), ''),
                COALESCE(to_char(ID_Z_PRST_GL1), ''),
                COALESCE(to_char(ID_Z_ZMSP_GL1), ''),
                COALESCE(to_char(ID_Z_UZMS_GL1), ''),
                COALESCE(to_char(Z_LOKK_GL1), ''),
                COALESCE(to_char(ID_Z_DSDF_GL2), ''),
                COALESCE(to_char(ID_Z_KVAR_GL2), ''),
                COALESCE(to_char(ID_Z_RAPU_GL2), ''),
                COALESCE(to_char(ID_Z_PRST_GL2), ''),
                COALESCE(to_char(ID_Z_ZMSP_GL2), ''),
                COALESCE(to_char(ID_Z_UZMS_GL2), ''),
                COALESCE(to_char(Z_LOKK_GL2), ''),
                COALESCE(to_char(ID_Z_PREK_VN), ''),
                COALESCE(to_char(ID_Z_DIS_REZ), ''),
                COALESCE(to_char(ID_Z_KVAR_REZ), ''),
                COALESCE(to_char(ID_Z_PRST_REZ), ''),
                COALESCE(to_char(ID_Z_ZMSP_REZ), ''),
                COALESCE(to_char(ID_Z_NEL1), ''),
                COALESCE(to_char(ID_Z_NEL2), ''),
                COALESCE(to_char(ID_Z_NEL3), ''),
                COALESCE(to_char(ID_Z_PREK_NN), ''),
                COALESCE(to_char(ID_Z_SABZ_SAB), ''),
                COALESCE(to_char(ID_Z_OTPR_SAB), ''),
                COALESCE(to_char(ID_S_VREM_USL), ''),
                COALESCE(UZROK_TEKST, ''),
                COALESCE(to_char(ID_Z_JPS_VN), ''),
                COALESCE(to_char(ID_Z_JPS_NN), ''),
                COALESCE(POSL_TEKST, ''),
                COALESCE(to_char(ID_Z_TELE_POC_GL1), ''),
                COALESCE(to_char(ID_Z_TELE_KRAJ_GL1), ''),
                COALESCE(to_char(ID_Z_TELE_POC_GL2), ''),
                COALESCE(to_char(ID_Z_TELE_KRAJ_GL2), ''),
                COALESCE(SYNSOFT_ID, '')
				from PI_PI_DDN`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var p []*models.PiPiDDN

	for rows.Next() {
		var pipiddn models.PiPiDDN
		err := rows.Scan(
			&pipiddn.Datizv,
			&pipiddn.IdSMrc,
			&pipiddn.IdSTipd,
			&pipiddn.IdSVrpd,
			&pipiddn.IdTipob,
			&pipiddn.ObId,
			&pipiddn.TrafoId,
			&pipiddn.Vrepoc,
			&pipiddn.PocPP,
			&pipiddn.Vrezav,
			&pipiddn.ZavPP,
			&pipiddn.Id1SGruzr,
			&pipiddn.Id1SUzrok,
			&pipiddn.IdSGrraz,
			&pipiddn.IdSRazlog,
			&pipiddn.Snaga,
			&pipiddn.Opis,
			&pipiddn.IdSNap,
			&pipiddn.P2TrafId,
			&pipiddn.KorUneo,
			&pipiddn.Status,
			&pipiddn.Datpri,
			&pipiddn.IdZDsdfGL1,
			&pipiddn.IdZKvarGL1,
			&pipiddn.IdZRapuGL1,
			&pipiddn.IdZPrstGL1,
			&pipiddn.IdZZmspGL1,
			&pipiddn.IdZUzmsGL1,
			&pipiddn.ZLokkGL1,
			&pipiddn.IdZDsdfGL2,
			&pipiddn.IdZKvarGL2,
			&pipiddn.IdZRapuGL2,
			&pipiddn.IdZPrstGL2,
			&pipiddn.IdZZmspGL2,
			&pipiddn.IdZUzmsGL2,
			&pipiddn.ZLokkGL2,
			&pipiddn.IdZPrekVN,
			&pipiddn.IdZDisREZ,
			&pipiddn.IdZKvarREZ,
			&pipiddn.IdZPrstREZ,
			&pipiddn.IdZZmspREZ,
			&pipiddn.IdZNel1,
			&pipiddn.IdZNel2,
			&pipiddn.IdZNel3,
			&pipiddn.IdZPrekNN,
			&pipiddn.IdZSabzSAB,
			&pipiddn.IdZOtprSAB,
			&pipiddn.IdSVremUSL,
			&pipiddn.UzrokTekst,
			&pipiddn.IdZJpsVN,
			&pipiddn.IdZJpsNN,
			&pipiddn.PoslTekst,
			&pipiddn.IdZTelePocGL1,
			&pipiddn.IdZTeleKrajGL1,
			&pipiddn.IdZTelePocGL2,
			&pipiddn.IdZTeleKrajGL2,
			&pipiddn.SynsoftId,
		)

		if err != nil {
			return nil, err
		}

		p = append(p, &pipiddn)
	}

	return p, nil
}

func (m *OracleDBRepo) GetPiPiDDNByID(synsoftId string) (*models.PiPiDDN, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select to_char(DATIZV, 'dd.mm.yyyy'),
	COALESCE(to_char(ID_S_MRC), ''),
	COALESCE(to_char(ID_S_TIPD), ''),
	COALESCE(to_char(ID_S_VRPD), ''),
	COALESCE(to_char(ID_TIPOB), ''),
	COALESCE(to_char(OB_ID), ''),
	COALESCE(to_char(TRAFO_ID), ''),
	to_char(VREPOC, 'dd.mm.yyyy HH24:MI:SS'),
	POC_PP,
	to_char(VREZAV, 'dd.mm.yyyy HH24:MI:SS'),
	ZAV_PP,
	COALESCE(to_char(ID1_S_GRUZR), ''),
	COALESCE(to_char(ID1_S_UZROK), ''),
	COALESCE(to_char(ID_S_GRRAZ), ''),
	COALESCE(to_char(ID_S_RAZLOG), ''),
	COALESCE(to_char(SNAGA), ''),
	OPIS,
	COALESCE(to_char(ID_S_NAP), ''),
	COALESCE(to_char(P2_TRAF_ID), ''),
	COALESCE(PGI_KOR, ''),
	COALESCE(STATUS, ''),
	to_char(DATPRI, 'dd.mm.yyyy HH24:MI:SS'),
	COALESCE(to_char(ID_Z_DSDF_GL1), ''),
	COALESCE(to_char(ID_Z_KVAR_GL1), ''),
	COALESCE(to_char(ID_Z_RAPU_GL1), ''),
	COALESCE(to_char(ID_Z_PRST_GL1), ''),
	COALESCE(to_char(ID_Z_ZMSP_GL1), ''),
	COALESCE(to_char(ID_Z_UZMS_GL1), ''),
	COALESCE(to_char(Z_LOKK_GL1), ''),
	COALESCE(to_char(ID_Z_DSDF_GL2), ''),
	COALESCE(to_char(ID_Z_KVAR_GL2), ''),
	COALESCE(to_char(ID_Z_RAPU_GL2), ''),
	COALESCE(to_char(ID_Z_PRST_GL2), ''),
	COALESCE(to_char(ID_Z_ZMSP_GL2), ''),
	COALESCE(to_char(ID_Z_UZMS_GL2), ''),
	COALESCE(to_char(Z_LOKK_GL2), ''),
	COALESCE(to_char(ID_Z_PREK_VN), ''),
	COALESCE(to_char(ID_Z_DIS_REZ), ''),
	COALESCE(to_char(ID_Z_KVAR_REZ), ''),
	COALESCE(to_char(ID_Z_PRST_REZ), ''),
	COALESCE(to_char(ID_Z_ZMSP_REZ), ''),
	COALESCE(to_char(ID_Z_NEL1), ''),
	COALESCE(to_char(ID_Z_NEL2), ''),
	COALESCE(to_char(ID_Z_NEL3), ''),
	COALESCE(to_char(ID_Z_PREK_NN), ''),
	COALESCE(to_char(ID_Z_SABZ_SAB), ''),
	COALESCE(to_char(ID_Z_OTPR_SAB), ''),
	COALESCE(to_char(ID_S_VREM_USL), ''),
	COALESCE(UZROK_TEKST, ''),
	COALESCE(to_char(ID_Z_JPS_VN), ''),
	COALESCE(to_char(ID_Z_JPS_NN), ''),
	COALESCE(POSL_TEKST, ''),
	COALESCE(to_char(ID_Z_TELE_POC_GL1), ''),
	COALESCE(to_char(ID_Z_TELE_KRAJ_GL1), ''),
	COALESCE(to_char(ID_Z_TELE_POC_GL2), ''),
	COALESCE(to_char(ID_Z_TELE_KRAJ_GL2), ''),
	COALESCE(SYNSOFT_ID, '')
	from PI_PI_DDN_S where SYNSOFT_ID = :1`

	row := m.DB.QueryRowContext(ctx, query, synsoftId)

	var pipiddn models.PiPiDDN
	err := row.Scan(
		&pipiddn.Datizv,
		&pipiddn.IdSMrc,
		&pipiddn.IdSTipd,
		&pipiddn.IdSVrpd,
		&pipiddn.IdTipob,
		&pipiddn.ObId,
		&pipiddn.TrafoId,
		&pipiddn.Vrepoc,
		&pipiddn.PocPP,
		&pipiddn.Vrezav,
		&pipiddn.ZavPP,
		&pipiddn.Id1SGruzr,
		&pipiddn.Id1SUzrok,
		&pipiddn.IdSGrraz,
		&pipiddn.IdSRazlog,
		&pipiddn.Snaga,
		&pipiddn.Opis,
		&pipiddn.IdSNap,
		&pipiddn.P2TrafId,
		&pipiddn.KorUneo,
		&pipiddn.Status,
		&pipiddn.Datpri,
		&pipiddn.IdZDsdfGL1,
		&pipiddn.IdZKvarGL1,
		&pipiddn.IdZRapuGL1,
		&pipiddn.IdZPrstGL1,
		&pipiddn.IdZZmspGL1,
		&pipiddn.IdZUzmsGL1,
		&pipiddn.ZLokkGL1,
		&pipiddn.IdZDsdfGL2,
		&pipiddn.IdZKvarGL2,
		&pipiddn.IdZRapuGL2,
		&pipiddn.IdZPrstGL2,
		&pipiddn.IdZZmspGL2,
		&pipiddn.IdZUzmsGL2,
		&pipiddn.ZLokkGL2,
		&pipiddn.IdZPrekVN,
		&pipiddn.IdZDisREZ,
		&pipiddn.IdZKvarREZ,
		&pipiddn.IdZPrstREZ,
		&pipiddn.IdZZmspREZ,
		&pipiddn.IdZNel1,
		&pipiddn.IdZNel2,
		&pipiddn.IdZNel3,
		&pipiddn.IdZPrekNN,
		&pipiddn.IdZSabzSAB,
		&pipiddn.IdZOtprSAB,
		&pipiddn.IdSVremUSL,
		&pipiddn.UzrokTekst,
		&pipiddn.IdZJpsVN,
		&pipiddn.IdZJpsNN,
		&pipiddn.PoslTekst,
		&pipiddn.IdZTelePocGL1,
		&pipiddn.IdZTeleKrajGL1,
		&pipiddn.IdZTelePocGL2,
		&pipiddn.IdZTeleKrajGL2,
		&pipiddn.SynsoftId,
	)

	if err != nil {
		return nil, err
	}

	return &pipiddn, nil
}

func (m *OracleDBRepo) GetPiPiDDNByIDP(synsoftId string) (*models.PiPiDDN, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select to_char(DATIZV, 'dd.mm.yyyy'),
	COALESCE(to_char(ID_S_MRC), ''),
	COALESCE(to_char(ID_S_TIPD), ''),
	COALESCE(to_char(ID_S_VRPD), ''),
	COALESCE(to_char(ID_TIPOB), ''),
	COALESCE(to_char(OB_ID), ''),
	COALESCE(to_char(TRAFO_ID), ''),
	to_char(VREPOC, 'dd.mm.yyyy HH24:MI:SS'),
	POC_PP,
	to_char(VREZAV, 'dd.mm.yyyy HH24:MI:SS'),
	ZAV_PP,
	COALESCE(to_char(ID1_S_GRUZR), ''),
	COALESCE(to_char(ID1_S_UZROK), ''),
	COALESCE(to_char(ID_S_GRRAZ), ''),
	COALESCE(to_char(ID_S_RAZLOG), ''),
	COALESCE(to_char(SNAGA), ''),
	OPIS,
	COALESCE(to_char(ID_S_NAP), ''),
	COALESCE(to_char(P2_TRAF_ID), ''),
	COALESCE(PGI_KOR, ''),
	COALESCE(STATUS, ''),
	to_char(DATPRI, 'dd.mm.yyyy HH24:MI:SS'),
	COALESCE(to_char(ID_Z_DSDF_GL1), ''),
	COALESCE(to_char(ID_Z_KVAR_GL1), ''),
	COALESCE(to_char(ID_Z_RAPU_GL1), ''),
	COALESCE(to_char(ID_Z_PRST_GL1), ''),
	COALESCE(to_char(ID_Z_ZMSP_GL1), ''),
	COALESCE(to_char(ID_Z_UZMS_GL1), ''),
	COALESCE(to_char(Z_LOKK_GL1), ''),
	COALESCE(to_char(ID_Z_DSDF_GL2), ''),
	COALESCE(to_char(ID_Z_KVAR_GL2), ''),
	COALESCE(to_char(ID_Z_RAPU_GL2), ''),
	COALESCE(to_char(ID_Z_PRST_GL2), ''),
	COALESCE(to_char(ID_Z_ZMSP_GL2), ''),
	COALESCE(to_char(ID_Z_UZMS_GL2), ''),
	COALESCE(to_char(Z_LOKK_GL2), ''),
	COALESCE(to_char(ID_Z_PREK_VN), ''),
	COALESCE(to_char(ID_Z_DIS_REZ), ''),
	COALESCE(to_char(ID_Z_KVAR_REZ), ''),
	COALESCE(to_char(ID_Z_PRST_REZ), ''),
	COALESCE(to_char(ID_Z_ZMSP_REZ), ''),
	COALESCE(to_char(ID_Z_NEL1), ''),
	COALESCE(to_char(ID_Z_NEL2), ''),
	COALESCE(to_char(ID_Z_NEL3), ''),
	COALESCE(to_char(ID_Z_PREK_NN), ''),
	COALESCE(to_char(ID_Z_SABZ_SAB), ''),
	COALESCE(to_char(ID_Z_OTPR_SAB), ''),
	COALESCE(to_char(ID_S_VREM_USL), ''),
	COALESCE(UZROK_TEKST, ''),
	COALESCE(to_char(ID_Z_JPS_VN), ''),
	COALESCE(to_char(ID_Z_JPS_NN), ''),
	COALESCE(POSL_TEKST, ''),
	COALESCE(to_char(ID_Z_TELE_POC_GL1), ''),
	COALESCE(to_char(ID_Z_TELE_KRAJ_GL1), ''),
	COALESCE(to_char(ID_Z_TELE_POC_GL2), ''),
	COALESCE(to_char(ID_Z_TELE_KRAJ_GL2), ''),
	COALESCE(SYNSOFT_ID, '')
	from PI_PI_DDN where SYNSOFT_ID = :1`

	row := m.DB.QueryRowContext(ctx, query, synsoftId)

	var pipiddn models.PiPiDDN
	err := row.Scan(
		&pipiddn.Datizv,
		&pipiddn.IdSMrc,
		&pipiddn.IdSTipd,
		&pipiddn.IdSVrpd,
		&pipiddn.IdTipob,
		&pipiddn.ObId,
		&pipiddn.TrafoId,
		&pipiddn.Vrepoc,
		&pipiddn.PocPP,
		&pipiddn.Vrezav,
		&pipiddn.ZavPP,
		&pipiddn.Id1SGruzr,
		&pipiddn.Id1SUzrok,
		&pipiddn.IdSGrraz,
		&pipiddn.IdSRazlog,
		&pipiddn.Snaga,
		&pipiddn.Opis,
		&pipiddn.IdSNap,
		&pipiddn.P2TrafId,
		&pipiddn.KorUneo,
		&pipiddn.Status,
		&pipiddn.Datpri,
		&pipiddn.IdZDsdfGL1,
		&pipiddn.IdZKvarGL1,
		&pipiddn.IdZRapuGL1,
		&pipiddn.IdZPrstGL1,
		&pipiddn.IdZZmspGL1,
		&pipiddn.IdZUzmsGL1,
		&pipiddn.ZLokkGL1,
		&pipiddn.IdZDsdfGL2,
		&pipiddn.IdZKvarGL2,
		&pipiddn.IdZRapuGL2,
		&pipiddn.IdZPrstGL2,
		&pipiddn.IdZZmspGL2,
		&pipiddn.IdZUzmsGL2,
		&pipiddn.ZLokkGL2,
		&pipiddn.IdZPrekVN,
		&pipiddn.IdZDisREZ,
		&pipiddn.IdZKvarREZ,
		&pipiddn.IdZPrstREZ,
		&pipiddn.IdZZmspREZ,
		&pipiddn.IdZNel1,
		&pipiddn.IdZNel2,
		&pipiddn.IdZNel3,
		&pipiddn.IdZPrekNN,
		&pipiddn.IdZSabzSAB,
		&pipiddn.IdZOtprSAB,
		&pipiddn.IdSVremUSL,
		&pipiddn.UzrokTekst,
		&pipiddn.IdZJpsVN,
		&pipiddn.IdZJpsNN,
		&pipiddn.PoslTekst,
		&pipiddn.IdZTelePocGL1,
		&pipiddn.IdZTeleKrajGL1,
		&pipiddn.IdZTelePocGL2,
		&pipiddn.IdZTeleKrajGL2,
		&pipiddn.SynsoftId,
	)

	if err != nil {
		return nil, err
	}

	return &pipiddn, nil
}

func (m *OracleDBRepo) GetAllUnfinishedEventsNDC() ([]*models.UnfinishedEvents, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select to_char(DATIZV, 'dd.mm.yyyy'),
	COALESCE(to_char(ID_S_MRC), ''),
	COALESCE(to_char(ID_S_TIPD), ''),
	COALESCE(to_char(ID_S_VRPD), ''),
	COALESCE(to_char(ID_TIPOB), ''),
	COALESCE(to_char(OB_ID), ''),
	COALESCE(to_char(TRAFO_ID), ''),
   to_char(VREPOC, 'dd.mm.yyyy HH24:MI:SS'),
   to_char(VREZAV, 'dd.mm.yyyy HH24:MI:SS'),
   COALESCE(to_char(ID1_S_GRUZR), ''),
   COALESCE(to_char(ID1_S_UZROK), ''),
   COALESCE(to_char(SNAGA), ''),
   COALESCE(OPIS, ''),
   COALESCE(to_char(ID_S_NAP), ''),
   COALESCE(to_char(P2_TRAF_ID), ''),
   COALESCE(PGI_KOR, ''),
   COALESCE(to_char(ID_Z_DSDF_GL1), ''),
   COALESCE(to_char(ID_Z_KVAR_GL1), ''),
   COALESCE(to_char(ID_Z_RAPU_GL1), ''),
   COALESCE(to_char(ID_Z_PRST_GL1), ''),
   COALESCE(to_char(ID_Z_ZMSP_GL1), ''),
   COALESCE(to_char(ID_Z_UZMS_GL1), ''),
   COALESCE(to_char(Z_LOKK_GL1), ''),
   COALESCE(to_char(ID_Z_DSDF_GL2), ''),
   COALESCE(to_char(ID_Z_KVAR_GL2), ''),
   COALESCE(to_char(ID_Z_RAPU_GL2), ''),
   COALESCE(to_char(ID_Z_PRST_GL2), ''),
   COALESCE(to_char(ID_Z_ZMSP_GL2), ''),
   COALESCE(to_char(ID_Z_UZMS_GL2), ''),
   COALESCE(to_char(Z_LOKK_GL2), ''),
   COALESCE(to_char(ID_Z_PREK_VN), ''),
   COALESCE(to_char(ID_Z_DIS_REZ), ''),
   COALESCE(to_char(ID_Z_KVAR_REZ), ''),
   COALESCE(to_char(ID_Z_PRST_REZ), ''),
   COALESCE(to_char(ID_Z_ZMSP_REZ), ''),
   COALESCE(to_char(ID_Z_NEL1), ''),
   COALESCE(to_char(ID_Z_NEL2), ''),
   COALESCE(to_char(ID_Z_NEL3), ''),
   COALESCE(to_char(ID_Z_PREK_NN), ''),
   COALESCE(to_char(ID_Z_SABZ_SAB), ''),
   COALESCE(to_char(ID_Z_OTPR_SAB), ''),
   COALESCE(to_char(ID_S_VREM_USL), ''),
   COALESCE(UZROK_TEKST, ''),
   COALESCE(to_char(ID_Z_JPS_VN), ''),
   COALESCE(to_char(ID_Z_JPS_NN), ''),
   COALESCE(POSL_TEKST, ''),
   COALESCE(to_char(ID_Z_TELE_POC_GL1), ''),
   COALESCE(to_char(ID_Z_TELE_KRAJ_GL1), ''),
   COALESCE(to_char(ID_Z_TELE_POC_GL2), ''),
   COALESCE(to_char(ID_Z_TELE_KRAJ_GL2), ''),
   COALESCE(SYNSOFT_ID, '')
   from nezavrseni_dog_s
   where id_s_mrc=8`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var p []*models.UnfinishedEvents

	for rows.Next() {
		var ue models.UnfinishedEvents
		err := rows.Scan(
			&ue.Datizv,
			&ue.IdSMrc,
			&ue.IdSTipd,
			&ue.IdSVrpd,
			&ue.IdTipob,
			&ue.ObId,
			&ue.TrafoId,
			&ue.Vrepoc,
			&ue.Vrezav,
			&ue.Id1SGruzr,
			&ue.Id1SUzrok,
			&ue.Snaga,
			&ue.Opis,
			&ue.IdSNap,
			&ue.P2TrafId,
			&ue.PgiKor,
			&ue.IdZDsdfGL1,
			&ue.IdZKvarGL1,
			&ue.IdZRapuGL1,
			&ue.IdZPrstGL1,
			&ue.IdZZmspGL1,
			&ue.IdZUzmsGL1,
			&ue.ZLokkGL1,
			&ue.IdZDsdfGL2,
			&ue.IdZKvarGL2,
			&ue.IdZRapuGL2,
			&ue.IdZPrstGL2,
			&ue.IdZZmspGL2,
			&ue.IdZUzmsGL2,
			&ue.ZLokkGL2,
			&ue.IdZPrekVN,
			&ue.IdZDisREZ,
			&ue.IdZKvarREZ,
			&ue.IdZPrstREZ,
			&ue.IdZZmspREZ,
			&ue.IdZNel1,
			&ue.IdZNel2,
			&ue.IdZNel3,
			&ue.IdZPrekNN,
			&ue.IdZSabzSAB,
			&ue.IdZOtprSAB,
			&ue.IdSVremUSL,
			&ue.UzrokTekst,
			&ue.IdZJpsVN,
			&ue.IdZJpsNN,
			&ue.PoslTekst,
			&ue.IdZTelePocGL1,
			&ue.IdZTeleKrajGL1,
			&ue.IdZTelePocGL2,
			&ue.IdZTeleKrajGL2,
			&ue.SynsoftId,
		)

		if err != nil {
			return nil, err
		}

		p = append(p, &ue)
	}

	return p, nil
}

func (m *OracleDBRepo) GetAllUnfinishedEventsNDCP() ([]*models.UnfinishedEvents, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select to_char(DATIZV, 'dd.mm.yyyy'),
	COALESCE(to_char(ID_S_MRC), ''),
	COALESCE(to_char(ID_S_TIPD), ''),
	COALESCE(to_char(ID_S_VRPD), ''),
	COALESCE(to_char(ID_TIPOB), ''),
	COALESCE(to_char(OB_ID), ''),
	COALESCE(to_char(TRAFO_ID), ''),
   to_char(VREPOC, 'dd.mm.yyyy HH24:MI:SS'),
   to_char(VREZAV, 'dd.mm.yyyy HH24:MI:SS'),
   COALESCE(to_char(ID1_S_GRUZR), ''),
   COALESCE(to_char(ID1_S_UZROK), ''),
   COALESCE(to_char(SNAGA), ''),
   COALESCE(OPIS, ''),
   COALESCE(to_char(ID_S_NAP), ''),
   COALESCE(to_char(P2_TRAF_ID), ''),
   COALESCE(PGI_KOR, ''),
   COALESCE(to_char(ID_Z_DSDF_GL1), ''),
   COALESCE(to_char(ID_Z_KVAR_GL1), ''),
   COALESCE(to_char(ID_Z_RAPU_GL1), ''),
   COALESCE(to_char(ID_Z_PRST_GL1), ''),
   COALESCE(to_char(ID_Z_ZMSP_GL1), ''),
   COALESCE(to_char(ID_Z_UZMS_GL1), ''),
   COALESCE(to_char(Z_LOKK_GL1), ''),
   COALESCE(to_char(ID_Z_DSDF_GL2), ''),
   COALESCE(to_char(ID_Z_KVAR_GL2), ''),
   COALESCE(to_char(ID_Z_RAPU_GL2), ''),
   COALESCE(to_char(ID_Z_PRST_GL2), ''),
   COALESCE(to_char(ID_Z_ZMSP_GL2), ''),
   COALESCE(to_char(ID_Z_UZMS_GL2), ''),
   COALESCE(to_char(Z_LOKK_GL2), ''),
   COALESCE(to_char(ID_Z_PREK_VN), ''),
   COALESCE(to_char(ID_Z_DIS_REZ), ''),
   COALESCE(to_char(ID_Z_KVAR_REZ), ''),
   COALESCE(to_char(ID_Z_PRST_REZ), ''),
   COALESCE(to_char(ID_Z_ZMSP_REZ), ''),
   COALESCE(to_char(ID_Z_NEL1), ''),
   COALESCE(to_char(ID_Z_NEL2), ''),
   COALESCE(to_char(ID_Z_NEL3), ''),
   COALESCE(to_char(ID_Z_PREK_NN), ''),
   COALESCE(to_char(ID_Z_SABZ_SAB), ''),
   COALESCE(to_char(ID_Z_OTPR_SAB), ''),
   COALESCE(to_char(ID_S_VREM_USL), ''),
   COALESCE(UZROK_TEKST, ''),
   COALESCE(to_char(ID_Z_JPS_VN), ''),
   COALESCE(to_char(ID_Z_JPS_NN), ''),
   COALESCE(POSL_TEKST, ''),
   COALESCE(to_char(ID_Z_TELE_POC_GL1), ''),
   COALESCE(to_char(ID_Z_TELE_KRAJ_GL1), ''),
   COALESCE(to_char(ID_Z_TELE_POC_GL2), ''),
   COALESCE(to_char(ID_Z_TELE_KRAJ_GL2), ''),
   COALESCE(SYNSOFT_ID, '')
   from nezavrseni_dog
   where id_s_mrc=8`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var p []*models.UnfinishedEvents

	for rows.Next() {
		var ue models.UnfinishedEvents
		err := rows.Scan(
			&ue.Datizv,
			&ue.IdSMrc,
			&ue.IdSTipd,
			&ue.IdSVrpd,
			&ue.IdTipob,
			&ue.ObId,
			&ue.TrafoId,
			&ue.Vrepoc,
			&ue.Vrezav,
			&ue.Id1SGruzr,
			&ue.Id1SUzrok,
			&ue.Snaga,
			&ue.Opis,
			&ue.IdSNap,
			&ue.P2TrafId,
			&ue.PgiKor,
			&ue.IdZDsdfGL1,
			&ue.IdZKvarGL1,
			&ue.IdZRapuGL1,
			&ue.IdZPrstGL1,
			&ue.IdZZmspGL1,
			&ue.IdZUzmsGL1,
			&ue.ZLokkGL1,
			&ue.IdZDsdfGL2,
			&ue.IdZKvarGL2,
			&ue.IdZRapuGL2,
			&ue.IdZPrstGL2,
			&ue.IdZZmspGL2,
			&ue.IdZUzmsGL2,
			&ue.ZLokkGL2,
			&ue.IdZPrekVN,
			&ue.IdZDisREZ,
			&ue.IdZKvarREZ,
			&ue.IdZPrstREZ,
			&ue.IdZZmspREZ,
			&ue.IdZNel1,
			&ue.IdZNel2,
			&ue.IdZNel3,
			&ue.IdZPrekNN,
			&ue.IdZSabzSAB,
			&ue.IdZOtprSAB,
			&ue.IdSVremUSL,
			&ue.UzrokTekst,
			&ue.IdZJpsVN,
			&ue.IdZJpsNN,
			&ue.PoslTekst,
			&ue.IdZTelePocGL1,
			&ue.IdZTeleKrajGL1,
			&ue.IdZTelePocGL2,
			&ue.IdZTeleKrajGL2,
			&ue.SynsoftId,
		)

		if err != nil {
			return nil, err
		}

		p = append(p, &ue)
	}

	return p, nil
}

func (m *OracleDBRepo) GetUnfinishedEventsByID(synsoftId string) (*models.UnfinishedEvents, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select to_char(DATIZV, 'dd.mm.yyyy'),
	COALESCE(to_char(ID_S_MRC), ''),
	COALESCE(to_char(ID_S_TIPD), ''),
	COALESCE(to_char(ID_S_VRPD), ''),
	COALESCE(to_char(ID_TIPOB), ''),
	COALESCE(to_char(OB_ID), ''),
	COALESCE(to_char(TRAFO_ID), ''),
   to_char(VREPOC, 'dd.mm.yyyy HH24:MI:SS'),
   to_char(VREZAV, 'dd.mm.yyyy HH24:MI:SS'),
   COALESCE(to_char(ID1_S_GRUZR), ''),
   COALESCE(to_char(ID1_S_UZROK), ''),
   COALESCE(to_char(SNAGA), ''),
   COALESCE(OPIS, ''),
   COALESCE(to_char(ID_S_NAP), ''),
   COALESCE(to_char(P2_TRAF_ID), ''),
   COALESCE(PGI_KOR, ''),
   COALESCE(to_char(ID_Z_DSDF_GL1), ''),
   COALESCE(to_char(ID_Z_KVAR_GL1), ''),
   COALESCE(to_char(ID_Z_RAPU_GL1), ''),
   COALESCE(to_char(ID_Z_PRST_GL1), ''),
   COALESCE(to_char(ID_Z_ZMSP_GL1), ''),
   COALESCE(to_char(ID_Z_UZMS_GL1), ''),
   COALESCE(to_char(Z_LOKK_GL1), ''),
   COALESCE(to_char(ID_Z_DSDF_GL2), ''),
   COALESCE(to_char(ID_Z_KVAR_GL2), ''),
   COALESCE(to_char(ID_Z_RAPU_GL2), ''),
   COALESCE(to_char(ID_Z_PRST_GL2), ''),
   COALESCE(to_char(ID_Z_ZMSP_GL2), ''),
   COALESCE(to_char(ID_Z_UZMS_GL2), ''),
   COALESCE(to_char(Z_LOKK_GL2), ''),
   COALESCE(to_char(ID_Z_PREK_VN), ''),
   COALESCE(to_char(ID_Z_DIS_REZ), ''),
   COALESCE(to_char(ID_Z_KVAR_REZ), ''),
   COALESCE(to_char(ID_Z_PRST_REZ), ''),
   COALESCE(to_char(ID_Z_ZMSP_REZ), ''),
   COALESCE(to_char(ID_Z_NEL1), ''),
   COALESCE(to_char(ID_Z_NEL2), ''),
   COALESCE(to_char(ID_Z_NEL3), ''),
   COALESCE(to_char(ID_Z_PREK_NN), ''),
   COALESCE(to_char(ID_Z_SABZ_SAB), ''),
   COALESCE(to_char(ID_Z_OTPR_SAB), ''),
   COALESCE(to_char(ID_S_VREM_USL), ''),
   COALESCE(UZROK_TEKST, ''),
   COALESCE(to_char(ID_Z_JPS_VN), ''),
   COALESCE(to_char(ID_Z_JPS_NN), ''),
   COALESCE(POSL_TEKST, ''),
   COALESCE(to_char(ID_Z_TELE_POC_GL1), ''),
   COALESCE(to_char(ID_Z_TELE_KRAJ_GL1), ''),
   COALESCE(to_char(ID_Z_TELE_POC_GL2), ''),
   COALESCE(to_char(ID_Z_TELE_KRAJ_GL2), ''),
   COALESCE(SYNSOFT_ID, '')
   from nezavrseni_dog_s where SYNSOFT_ID = :1`

	row := m.DB.QueryRowContext(ctx, query, synsoftId)

	var ue models.UnfinishedEvents
	err := row.Scan(
		&ue.Datizv,
		&ue.IdSMrc,
		&ue.IdSTipd,
		&ue.IdSVrpd,
		&ue.IdTipob,
		&ue.ObId,
		&ue.TrafoId,
		&ue.Vrepoc,
		&ue.Vrezav,
		&ue.Id1SGruzr,
		&ue.Id1SUzrok,
		&ue.Snaga,
		&ue.Opis,
		&ue.IdSNap,
		&ue.P2TrafId,
		&ue.PgiKor,
		&ue.IdZDsdfGL1,
		&ue.IdZKvarGL1,
		&ue.IdZRapuGL1,
		&ue.IdZPrstGL1,
		&ue.IdZZmspGL1,
		&ue.IdZUzmsGL1,
		&ue.ZLokkGL1,
		&ue.IdZDsdfGL2,
		&ue.IdZKvarGL2,
		&ue.IdZRapuGL2,
		&ue.IdZPrstGL2,
		&ue.IdZZmspGL2,
		&ue.IdZUzmsGL2,
		&ue.ZLokkGL2,
		&ue.IdZPrekVN,
		&ue.IdZDisREZ,
		&ue.IdZKvarREZ,
		&ue.IdZPrstREZ,
		&ue.IdZZmspREZ,
		&ue.IdZNel1,
		&ue.IdZNel2,
		&ue.IdZNel3,
		&ue.IdZPrekNN,
		&ue.IdZSabzSAB,
		&ue.IdZOtprSAB,
		&ue.IdSVremUSL,
		&ue.UzrokTekst,
		&ue.IdZJpsVN,
		&ue.IdZJpsNN,
		&ue.PoslTekst,
		&ue.IdZTelePocGL1,
		&ue.IdZTeleKrajGL1,
		&ue.IdZTelePocGL2,
		&ue.IdZTeleKrajGL2,
		&ue.SynsoftId,
	)

	if err != nil {
		return nil, err
	}

	return &ue, nil
}

func (m *OracleDBRepo) GetUnfinishedEventsByIDP(synsoftId string) (*models.UnfinishedEvents, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select to_char(DATIZV, 'dd.mm.yyyy'),
	COALESCE(to_char(ID_S_MRC), ''),
	COALESCE(to_char(ID_S_TIPD), ''),
	COALESCE(to_char(ID_S_VRPD), ''),
	COALESCE(to_char(ID_TIPOB), ''),
	COALESCE(to_char(OB_ID), ''),
	COALESCE(to_char(TRAFO_ID), ''),
   to_char(VREPOC, 'dd.mm.yyyy HH24:MI:SS'),
   to_char(VREZAV, 'dd.mm.yyyy HH24:MI:SS'),
   COALESCE(to_char(ID1_S_GRUZR), ''),
   COALESCE(to_char(ID1_S_UZROK), ''),
   COALESCE(to_char(SNAGA), ''),
   COALESCE(OPIS, ''),
   COALESCE(to_char(ID_S_NAP), ''),
   COALESCE(to_char(P2_TRAF_ID), ''),
   COALESCE(PGI_KOR, ''),
   COALESCE(to_char(ID_Z_DSDF_GL1), ''),
   COALESCE(to_char(ID_Z_KVAR_GL1), ''),
   COALESCE(to_char(ID_Z_RAPU_GL1), ''),
   COALESCE(to_char(ID_Z_PRST_GL1), ''),
   COALESCE(to_char(ID_Z_ZMSP_GL1), ''),
   COALESCE(to_char(ID_Z_UZMS_GL1), ''),
   COALESCE(to_char(Z_LOKK_GL1), ''),
   COALESCE(to_char(ID_Z_DSDF_GL2), ''),
   COALESCE(to_char(ID_Z_KVAR_GL2), ''),
   COALESCE(to_char(ID_Z_RAPU_GL2), ''),
   COALESCE(to_char(ID_Z_PRST_GL2), ''),
   COALESCE(to_char(ID_Z_ZMSP_GL2), ''),
   COALESCE(to_char(ID_Z_UZMS_GL2), ''),
   COALESCE(to_char(Z_LOKK_GL2), ''),
   COALESCE(to_char(ID_Z_PREK_VN), ''),
   COALESCE(to_char(ID_Z_DIS_REZ), ''),
   COALESCE(to_char(ID_Z_KVAR_REZ), ''),
   COALESCE(to_char(ID_Z_PRST_REZ), ''),
   COALESCE(to_char(ID_Z_ZMSP_REZ), ''),
   COALESCE(to_char(ID_Z_NEL1), ''),
   COALESCE(to_char(ID_Z_NEL2), ''),
   COALESCE(to_char(ID_Z_NEL3), ''),
   COALESCE(to_char(ID_Z_PREK_NN), ''),
   COALESCE(to_char(ID_Z_SABZ_SAB), ''),
   COALESCE(to_char(ID_Z_OTPR_SAB), ''),
   COALESCE(to_char(ID_S_VREM_USL), ''),
   COALESCE(UZROK_TEKST, ''),
   COALESCE(to_char(ID_Z_JPS_VN), ''),
   COALESCE(to_char(ID_Z_JPS_NN), ''),
   COALESCE(POSL_TEKST, ''),
   COALESCE(to_char(ID_Z_TELE_POC_GL1), ''),
   COALESCE(to_char(ID_Z_TELE_KRAJ_GL1), ''),
   COALESCE(to_char(ID_Z_TELE_POC_GL2), ''),
   COALESCE(to_char(ID_Z_TELE_KRAJ_GL2), ''),
   COALESCE(SYNSOFT_ID, '')
   from nezavrseni_dog where SYNSOFT_ID = :1`

	row := m.DB.QueryRowContext(ctx, query, synsoftId)

	var ue models.UnfinishedEvents
	err := row.Scan(
		&ue.Datizv,
		&ue.IdSMrc,
		&ue.IdSTipd,
		&ue.IdSVrpd,
		&ue.IdTipob,
		&ue.ObId,
		&ue.TrafoId,
		&ue.Vrepoc,
		&ue.Vrezav,
		&ue.Id1SGruzr,
		&ue.Id1SUzrok,
		&ue.Snaga,
		&ue.Opis,
		&ue.IdSNap,
		&ue.P2TrafId,
		&ue.PgiKor,
		&ue.IdZDsdfGL1,
		&ue.IdZKvarGL1,
		&ue.IdZRapuGL1,
		&ue.IdZPrstGL1,
		&ue.IdZZmspGL1,
		&ue.IdZUzmsGL1,
		&ue.ZLokkGL1,
		&ue.IdZDsdfGL2,
		&ue.IdZKvarGL2,
		&ue.IdZRapuGL2,
		&ue.IdZPrstGL2,
		&ue.IdZZmspGL2,
		&ue.IdZUzmsGL2,
		&ue.ZLokkGL2,
		&ue.IdZPrekVN,
		&ue.IdZDisREZ,
		&ue.IdZKvarREZ,
		&ue.IdZPrstREZ,
		&ue.IdZZmspREZ,
		&ue.IdZNel1,
		&ue.IdZNel2,
		&ue.IdZNel3,
		&ue.IdZPrekNN,
		&ue.IdZSabzSAB,
		&ue.IdZOtprSAB,
		&ue.IdSVremUSL,
		&ue.UzrokTekst,
		&ue.IdZJpsVN,
		&ue.IdZJpsNN,
		&ue.PoslTekst,
		&ue.IdZTelePocGL1,
		&ue.IdZTeleKrajGL1,
		&ue.IdZTelePocGL2,
		&ue.IdZTeleKrajGL2,
		&ue.SynsoftId,
	)

	if err != nil {
		return nil, err
	}

	return &ue, nil
}

func (m *OracleDBRepo) UpdateUnfinishedEvents(ue models.UnfinishedEventsUpdate) error {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var status int
	var message string

	query := `begin  ddn.synsoft.nezavrseni_dog_update(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42); end;`
	//var int status
	//var string message
	_, err := m.DB.ExecContext(ctx, query,
		ue.DatSmene,
		ue.TipSmene,
		ue.Vrepoc,
		ue.Vrezav,
		ue.Id1SGruzr,
		ue.Id1SUzrok,
		ue.Opis,
		ue.IdZDsdfGL1,
		ue.IdZKvarGL1,
		ue.IdZRapuGL1,
		ue.IdZPrstGL1,
		ue.IdZZmspGL1,
		ue.IdZUzmsGL1,
		ue.ZLokkGL1,
		ue.IdZDsdfGL2,
		ue.IdZKvarGL2,
		ue.IdZRapuGL2,
		ue.IdZPrstGL2,
		ue.IdZZmspGL2,
		ue.IdZUzmsGL2,
		ue.ZLokkGL2,
		ue.IdZPrekVN,
		ue.IdZDisREZ,
		ue.IdZKvarREZ,
		ue.IdZPrstREZ,
		ue.IdZZmspREZ,
		ue.IdZNel1,
		ue.IdZNel2,
		ue.IdZNel3,
		ue.IdZPrekNN,
		ue.IdZSabzSAB,
		ue.IdZOtprSAB,
		ue.IdSVremUSL,
		ue.IdZJpsVN,
		ue.IdZJpsNN,
		ue.IdZTelePocGL1,
		ue.IdZTeleKrajGL1,
		ue.IdZTelePocGL2,
		ue.IdZTeleKrajGL2,
		ue.SynsoftId,
		sql.Out{Dest: &status},
		sql.Out{Dest: &message},
	)

	if err != nil {
		log.Println(err)
		return err
	}

	if status != 0 {
		return errors.New(message)
	} else {
		return nil
	}
}

func (m *OracleDBRepo) UpdateUnfinishedEventsP(ue models.UnfinishedEventsUpdate) error {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var status int
	var message string

	query := `begin  ddn.synsoft.p_nezavrseni_dog_update(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42); end;`
	//var int status
	//var string message
	_, err := m.DB.ExecContext(ctx, query,
		ue.DatSmene,
		ue.TipSmene,
		ue.Vrepoc,
		ue.Vrezav,
		ue.Id1SGruzr,
		ue.Id1SUzrok,
		ue.Opis,
		ue.IdZDsdfGL1,
		ue.IdZKvarGL1,
		ue.IdZRapuGL1,
		ue.IdZPrstGL1,
		ue.IdZZmspGL1,
		ue.IdZUzmsGL1,
		ue.ZLokkGL1,
		ue.IdZDsdfGL2,
		ue.IdZKvarGL2,
		ue.IdZRapuGL2,
		ue.IdZPrstGL2,
		ue.IdZZmspGL2,
		ue.IdZUzmsGL2,
		ue.ZLokkGL2,
		ue.IdZPrekVN,
		ue.IdZDisREZ,
		ue.IdZKvarREZ,
		ue.IdZPrstREZ,
		ue.IdZZmspREZ,
		ue.IdZNel1,
		ue.IdZNel2,
		ue.IdZNel3,
		ue.IdZPrekNN,
		ue.IdZSabzSAB,
		ue.IdZOtprSAB,
		ue.IdSVremUSL,
		ue.IdZJpsVN,
		ue.IdZJpsNN,
		ue.IdZTelePocGL1,
		ue.IdZTeleKrajGL1,
		ue.IdZTelePocGL2,
		ue.IdZTeleKrajGL2,
		ue.SynsoftId,
		sql.Out{Dest: &status},
		sql.Out{Dest: &message},
	)

	if err != nil {
		log.Println(err)
		return err
	}

	if status != 0 {
		return errors.New(message)
	} else {
		return nil
	}
}

func (m *OracleDBRepo) DeletePiPiDDN(synsoftId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	stmt := `delete from pi_pi_ddn_s where synsoft_id = :1`

	_, err := m.DB.ExecContext(ctx, stmt, synsoftId)
	if err != nil {
		return err
	}

	return nil
}

func (m *OracleDBRepo) DeletePiPiDDNP(synsoftId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	stmt := `delete from pi_pi_ddn where synsoft_id = :1`

	_, err := m.DB.ExecContext(ctx, stmt, synsoftId)
	if err != nil {
		return err
	}

	return nil
}

func (m *OracleDBRepo) InsertPiPiDDNIspad(pipiddn models.PiPiDDNIspad) error {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var status int
	var message string

	query := `begin  ddn.synsoft.pi_pi_ddn_ispad_insert(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54); end;`
	//var int status
	//var string message
	_, err := m.DB.ExecContext(ctx, query,
		pipiddn.DatSmene,
		8,
		pipiddn.IdSTipd,
		pipiddn.IdSVrpd,
		pipiddn.IdRadAPU,
		pipiddn.IdTipob,
		pipiddn.ObId,
		pipiddn.TrafoId,
		pipiddn.Vrepoc,
		pipiddn.Vrezav,
		pipiddn.Id1SGruzr,
		pipiddn.Id1SUzrok,
		pipiddn.Snaga,
		pipiddn.Opis,
		pipiddn.IdSNap,
		pipiddn.P2TrafId,
		pipiddn.KorUneo,
		pipiddn.IdZDsdfGL1,
		pipiddn.IdZKvarGL1,
		pipiddn.IdZRapuGL1,
		pipiddn.IdZPrstGL1,
		pipiddn.IdZZmspGL1,
		pipiddn.IdZUzmsGL1,
		pipiddn.ZLokkGL1,
		pipiddn.IdZDsdfGL2,
		pipiddn.IdZKvarGL2,
		pipiddn.IdZRapuGL2,
		pipiddn.IdZPrstGL2,
		pipiddn.IdZZmspGL2,
		pipiddn.IdZUzmsGL2,
		pipiddn.ZLokkGL2,
		pipiddn.IdZPrekVN,
		pipiddn.IdZDisREZ,
		pipiddn.IdZKvarREZ,
		pipiddn.IdZPrstREZ,
		pipiddn.IdZZmspREZ,
		pipiddn.IdZNel1,
		pipiddn.IdZNel2,
		pipiddn.IdZNel3,
		pipiddn.IdZPrekNN,
		pipiddn.IdZSabzSAB,
		pipiddn.IdZOtprSAB,
		pipiddn.IdSVremUSL,
		pipiddn.UzrokTekst,
		pipiddn.IdZJpsVN,
		pipiddn.IdZJpsNN,
		pipiddn.PoslTekst,
		pipiddn.IdZTelePocGL1,
		pipiddn.IdZTeleKrajGL1,
		pipiddn.IdZTelePocGL2,
		pipiddn.IdZTeleKrajGL2,
		pipiddn.SynsoftId,
		sql.Out{Dest: &status},
		sql.Out{Dest: &message},
	)

	if err != nil {
		log.Println(err)
		return err
	}

	if status != 0 {
		return errors.New(message)
	} else {
		return nil
	}
}

func (m *OracleDBRepo) InsertPiPiDDNIspadP(pipiddn models.PiPiDDNIspad) error {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var status int
	var message string

	query := `begin  ddn.synsoft.p_pi_pi_ddn_ispad_insert(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54); end;`
	//var int status
	//var string message
	_, err := m.DB.ExecContext(ctx, query,
		pipiddn.DatSmene,
		8,
		pipiddn.IdSTipd,
		pipiddn.IdSVrpd,
		pipiddn.IdRadAPU,
		pipiddn.IdTipob,
		pipiddn.ObId,
		pipiddn.TrafoId,
		pipiddn.Vrepoc,
		pipiddn.Vrezav,
		pipiddn.Id1SGruzr,
		pipiddn.Id1SUzrok,
		pipiddn.Snaga,
		pipiddn.Opis,
		pipiddn.IdSNap,
		pipiddn.P2TrafId,
		pipiddn.KorUneo,
		pipiddn.IdZDsdfGL1,
		pipiddn.IdZKvarGL1,
		pipiddn.IdZRapuGL1,
		pipiddn.IdZPrstGL1,
		pipiddn.IdZZmspGL1,
		pipiddn.IdZUzmsGL1,
		pipiddn.ZLokkGL1,
		pipiddn.IdZDsdfGL2,
		pipiddn.IdZKvarGL2,
		pipiddn.IdZRapuGL2,
		pipiddn.IdZPrstGL2,
		pipiddn.IdZZmspGL2,
		pipiddn.IdZUzmsGL2,
		pipiddn.ZLokkGL2,
		pipiddn.IdZPrekVN,
		pipiddn.IdZDisREZ,
		pipiddn.IdZKvarREZ,
		pipiddn.IdZPrstREZ,
		pipiddn.IdZZmspREZ,
		pipiddn.IdZNel1,
		pipiddn.IdZNel2,
		pipiddn.IdZNel3,
		pipiddn.IdZPrekNN,
		pipiddn.IdZSabzSAB,
		pipiddn.IdZOtprSAB,
		pipiddn.IdSVremUSL,
		pipiddn.UzrokTekst,
		pipiddn.IdZJpsVN,
		pipiddn.IdZJpsNN,
		pipiddn.PoslTekst,
		pipiddn.IdZTelePocGL1,
		pipiddn.IdZTeleKrajGL1,
		pipiddn.IdZTelePocGL2,
		pipiddn.IdZTeleKrajGL2,
		pipiddn.SynsoftId,
		sql.Out{Dest: &status},
		sql.Out{Dest: &message},
	)

	if err != nil {
		log.Println(err)
		return err
	}

	if status != 0 {
		return errors.New(message)
	} else {
		return nil
	}
}

func (m *OracleDBRepo) UpdatePiPiDDNIspad(pipiddn models.PiPiDDNIspad) error {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var status int
	var message string

	query := `begin  ddn.synsoft.pi_pi_ddn_ispad_update(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54); end;`
	//var int status
	//var string message
	_, err := m.DB.ExecContext(ctx, query,
		pipiddn.DatSmene,
		8,
		pipiddn.IdSTipd,
		pipiddn.IdSVrpd,
		pipiddn.IdRadAPU,
		pipiddn.IdTipob,
		pipiddn.ObId,
		pipiddn.TrafoId,
		pipiddn.Vrepoc,
		pipiddn.Vrezav,
		pipiddn.Id1SGruzr,
		pipiddn.Id1SUzrok,
		pipiddn.Snaga,
		pipiddn.Opis,
		pipiddn.IdSNap,
		pipiddn.P2TrafId,
		pipiddn.KorUneo,
		pipiddn.IdZDsdfGL1,
		pipiddn.IdZKvarGL1,
		pipiddn.IdZRapuGL1,
		pipiddn.IdZPrstGL1,
		pipiddn.IdZZmspGL1,
		pipiddn.IdZUzmsGL1,
		pipiddn.ZLokkGL1,
		pipiddn.IdZDsdfGL2,
		pipiddn.IdZKvarGL2,
		pipiddn.IdZRapuGL2,
		pipiddn.IdZPrstGL2,
		pipiddn.IdZZmspGL2,
		pipiddn.IdZUzmsGL2,
		pipiddn.ZLokkGL2,
		pipiddn.IdZPrekVN,
		pipiddn.IdZDisREZ,
		pipiddn.IdZKvarREZ,
		pipiddn.IdZPrstREZ,
		pipiddn.IdZZmspREZ,
		pipiddn.IdZNel1,
		pipiddn.IdZNel2,
		pipiddn.IdZNel3,
		pipiddn.IdZPrekNN,
		pipiddn.IdZSabzSAB,
		pipiddn.IdZOtprSAB,
		pipiddn.IdSVremUSL,
		pipiddn.UzrokTekst,
		pipiddn.IdZJpsVN,
		pipiddn.IdZJpsNN,
		pipiddn.PoslTekst,
		pipiddn.IdZTelePocGL1,
		pipiddn.IdZTeleKrajGL1,
		pipiddn.IdZTelePocGL2,
		pipiddn.IdZTeleKrajGL2,
		pipiddn.SynsoftId,
		sql.Out{Dest: &status},
		sql.Out{Dest: &message},
	)

	if err != nil {
		log.Println(err)
		return err
	}

	if status != 0 {
		return errors.New(message)
	} else {
		return nil
	}
}

func (m *OracleDBRepo) UpdatePiPiDDNIspadP(pipiddn models.PiPiDDNIspad) error {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var status int
	var message string

	query := `begin  ddn.synsoft.p_pi_pi_ddn_ispad_update(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54); end;`
	//var int status
	//var string message
	_, err := m.DB.ExecContext(ctx, query,
		pipiddn.DatSmene,
		8,
		pipiddn.IdSTipd,
		pipiddn.IdSVrpd,
		pipiddn.IdRadAPU,
		pipiddn.IdTipob,
		pipiddn.ObId,
		pipiddn.TrafoId,
		pipiddn.Vrepoc,
		pipiddn.Vrezav,
		pipiddn.Id1SGruzr,
		pipiddn.Id1SUzrok,
		pipiddn.Snaga,
		pipiddn.Opis,
		pipiddn.IdSNap,
		pipiddn.P2TrafId,
		pipiddn.KorUneo,
		pipiddn.IdZDsdfGL1,
		pipiddn.IdZKvarGL1,
		pipiddn.IdZRapuGL1,
		pipiddn.IdZPrstGL1,
		pipiddn.IdZZmspGL1,
		pipiddn.IdZUzmsGL1,
		pipiddn.ZLokkGL1,
		pipiddn.IdZDsdfGL2,
		pipiddn.IdZKvarGL2,
		pipiddn.IdZRapuGL2,
		pipiddn.IdZPrstGL2,
		pipiddn.IdZZmspGL2,
		pipiddn.IdZUzmsGL2,
		pipiddn.ZLokkGL2,
		pipiddn.IdZPrekVN,
		pipiddn.IdZDisREZ,
		pipiddn.IdZKvarREZ,
		pipiddn.IdZPrstREZ,
		pipiddn.IdZZmspREZ,
		pipiddn.IdZNel1,
		pipiddn.IdZNel2,
		pipiddn.IdZNel3,
		pipiddn.IdZPrekNN,
		pipiddn.IdZSabzSAB,
		pipiddn.IdZOtprSAB,
		pipiddn.IdSVremUSL,
		pipiddn.UzrokTekst,
		pipiddn.IdZJpsVN,
		pipiddn.IdZJpsNN,
		pipiddn.PoslTekst,
		pipiddn.IdZTelePocGL1,
		pipiddn.IdZTeleKrajGL1,
		pipiddn.IdZTelePocGL2,
		pipiddn.IdZTeleKrajGL2,
		pipiddn.SynsoftId,
		sql.Out{Dest: &status},
		sql.Out{Dest: &message},
	)

	if err != nil {
		log.Println(err)
		return err
	}

	if status != 0 {
		return errors.New(message)
	} else {
		return nil
	}
}

func (m *OracleDBRepo) GetAllPiPiDDNIspad() ([]*models.PiPiDDN, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select to_char(DATIZV, 'dd.mm.yyyy'),
				COALESCE(to_char(ID_S_MRC), ''),
				COALESCE(to_char(ID_S_TIPD), ''),
				COALESCE(to_char(ID_S_VRPD), ''),
				COALESCE(to_char(ID_TIPOB), ''),
				COALESCE(to_char(OB_ID), ''),
				COALESCE(to_char(TRAFO_ID), ''),
				to_char(VREPOC, 'dd.mm.yyyy HH24:MI:SS'),
				POC_PP,
				to_char(VREZAV, 'dd.mm.yyyy HH24:MI:SS'),
				ZAV_PP,
				COALESCE(to_char(ID_S_GRRAZ), ''),
				COALESCE(to_char(ID_S_RAZLOG), ''),
				OPIS,
				COALESCE(to_char(ID_S_NAP), ''),
				COALESCE(to_char(P2_TRAF_ID), ''),
				PGI_KOR,
				STATUS,
				to_char(DATPRI, 'dd.mm.yyyy HH24:MI:SS'),
				SYNSOFT_ID
				from PI_PI_DDN_S
				WHERE ID_S_TIPD=1`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var p []*models.PiPiDDN

	for rows.Next() {
		var pipiddn models.PiPiDDN
		err := rows.Scan(
			&pipiddn.Datizv,
			&pipiddn.IdSMrc,
			&pipiddn.IdSTipd,
			&pipiddn.IdSVrpd,
			&pipiddn.IdTipob,
			&pipiddn.ObId,
			&pipiddn.TrafoId,
			&pipiddn.Vrepoc,
			&pipiddn.PocPP,
			&pipiddn.Vrezav,
			&pipiddn.ZavPP,
			&pipiddn.IdSGrraz,
			&pipiddn.IdSRazlog,
			&pipiddn.Opis,
			&pipiddn.IdSNap,
			&pipiddn.P2TrafId,
			&pipiddn.KorUneo,
			&pipiddn.Status,
			&pipiddn.Datpri,
			&pipiddn.SynsoftId,
		)

		if err != nil {
			return nil, err
		}

		p = append(p, &pipiddn)
	}

	return p, nil
}

func (m *OracleDBRepo) GetAllPiPiDDNIspadP() ([]*models.PiPiDDN, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select to_char(DATIZV, 'dd.mm.yyyy'),
				COALESCE(to_char(ID_S_MRC), ''),
				COALESCE(to_char(ID_S_TIPD), ''),
				COALESCE(to_char(ID_S_VRPD), ''),
				COALESCE(to_char(ID_TIPOB), ''),
				COALESCE(to_char(OB_ID), ''),
				COALESCE(to_char(TRAFO_ID), ''),
				to_char(VREPOC, 'dd.mm.yyyy HH24:MI:SS'),
				POC_PP,
				to_char(VREZAV, 'dd.mm.yyyy HH24:MI:SS'),
				ZAV_PP,
				COALESCE(to_char(ID_S_GRRAZ), ''),
				COALESCE(to_char(ID_S_RAZLOG), ''),
				OPIS,
				COALESCE(to_char(ID_S_NAP), ''),
				COALESCE(to_char(P2_TRAF_ID), ''),
				PGI_KOR,
				STATUS,
				to_char(DATPRI, 'dd.mm.yyyy HH24:MI:SS'),
				SYNSOFT_ID
				from PI_PI_DDN
				WHERE ID_S_TIPD=1`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var p []*models.PiPiDDN

	for rows.Next() {
		var pipiddn models.PiPiDDN
		err := rows.Scan(
			&pipiddn.Datizv,
			&pipiddn.IdSMrc,
			&pipiddn.IdSTipd,
			&pipiddn.IdSVrpd,
			&pipiddn.IdTipob,
			&pipiddn.ObId,
			&pipiddn.TrafoId,
			&pipiddn.Vrepoc,
			&pipiddn.PocPP,
			&pipiddn.Vrezav,
			&pipiddn.ZavPP,
			&pipiddn.IdSGrraz,
			&pipiddn.IdSRazlog,
			&pipiddn.Opis,
			&pipiddn.IdSNap,
			&pipiddn.P2TrafId,
			&pipiddn.KorUneo,
			&pipiddn.Status,
			&pipiddn.Datpri,
			&pipiddn.SynsoftId,
		)

		if err != nil {
			return nil, err
		}

		p = append(p, &pipiddn)
	}

	return p, nil
}

func (m *OracleDBRepo) InsertDDNInterruptionOfDelivery(ddnintd models.DDNInterruptionOfDelivery) error {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var status int
	var message string

	query := `begin  ddn.synsoft.ddn_prekid_isp_insert(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24); end;`
	//var int status
	//var string message
	_, err := m.DB.ExecContext(ctx, query,
		ddnintd.IdSMrc,
		ddnintd.IdSTipd,
		ddnintd.IdSVrpd,
		ddnintd.IdTipob,
		ddnintd.ObId,
		ddnintd.Vrepoc,
		ddnintd.Vrezav,
		ddnintd.IdSVrPrek,
		ddnintd.IdSUzrokPrek,
		ddnintd.Snaga,
		ddnintd.Opis,
		ddnintd.KorUneo,
		ddnintd.IdSMernaMesta,
		ddnintd.BrojMesta,
		ddnintd.Ind,
		ddnintd.P2TrafId,
		ddnintd.Bi,
		ddnintd.IdSPoduzrokPrek,
		ddnintd.IdDogPrekidP,
		ddnintd.IdTipObjektaNdc,
		ddnintd.IdTipDogadjajaNdc,
		ddnintd.SynsoftId,
		sql.Out{Dest: &status},
		sql.Out{Dest: &message},
	)

	if err != nil {
		log.Println(err)
		return err
	}
	//fmt.Println(pipiddn.TipMan)
	//fmt.Println(pipiddn.DatSmene)
	//fmt.Println(status)
	//fmt.Println(message)
	if status != 0 {
		return errors.New(message)
	} else {
		return nil
	}
}

func (m *OracleDBRepo) InsertDDNInterruptionOfDeliveryP(ddnintd models.DDNInterruptionOfDelivery) error {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var status int
	var message string

	query := `begin  ddn.synsoft.p_ddn_prekid_isp_insert(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24); end;`
	//var int status
	//var string message
	_, err := m.DB.ExecContext(ctx, query,
		ddnintd.IdSMrc,
		ddnintd.IdSTipd,
		ddnintd.IdSVrpd,
		ddnintd.IdTipob,
		ddnintd.ObId,
		ddnintd.Vrepoc,
		ddnintd.Vrezav,
		ddnintd.IdSVrPrek,
		ddnintd.IdSUzrokPrek,
		ddnintd.Snaga,
		ddnintd.Opis,
		ddnintd.KorUneo,
		ddnintd.IdSMernaMesta,
		ddnintd.BrojMesta,
		ddnintd.Ind,
		ddnintd.P2TrafId,
		ddnintd.Bi,
		ddnintd.IdSPoduzrokPrek,
		ddnintd.IdDogPrekidP,
		ddnintd.IdTipObjektaNdc,
		ddnintd.IdTipDogadjajaNdc,
		ddnintd.SynsoftId,
		sql.Out{Dest: &status},
		sql.Out{Dest: &message},
	)

	if err != nil {
		log.Println(err)
		return err
	}
	//fmt.Println(pipiddn.TipMan)
	//fmt.Println(pipiddn.DatSmene)
	//fmt.Println(status)
	//fmt.Println(message)
	if status != 0 {
		return errors.New(message)
	} else {
		return nil
	}
}

func (m *OracleDBRepo) UpdateDDNInterruptionOfDelivery(ddnintd models.DDNInterruptionOfDelivery) error {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var status int
	var message string

	query := `begin  ddn.synsoft.ddn_prekid_isp_update(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24); end;`
	//var int status
	//var string message
	_, err := m.DB.ExecContext(ctx, query,
		ddnintd.IdSMrc,
		ddnintd.IdSTipd,
		ddnintd.IdSVrpd,
		ddnintd.IdTipob,
		ddnintd.ObId,
		ddnintd.Vrepoc,
		ddnintd.Vrezav,
		ddnintd.IdSVrPrek,
		ddnintd.IdSUzrokPrek,
		ddnintd.Snaga,
		ddnintd.Opis,
		ddnintd.KorUneo,
		ddnintd.IdSMernaMesta,
		ddnintd.BrojMesta,
		ddnintd.Ind,
		ddnintd.P2TrafId,
		ddnintd.Bi,
		ddnintd.IdSPoduzrokPrek,
		ddnintd.IdDogPrekidP,
		ddnintd.IdTipObjektaNdc,
		ddnintd.IdTipDogadjajaNdc,
		ddnintd.SynsoftId,
		sql.Out{Dest: &status},
		sql.Out{Dest: &message},
	)

	if err != nil {
		log.Println(err)
		return err
	}
	//fmt.Println(pipiddn.TipMan)
	//fmt.Println(pipiddn.DatSmene)
	//fmt.Println(status)
	//fmt.Println(message)

	if status != 0 {
		return errors.New(message)
	} else {
		return nil
	}
}

func (m *OracleDBRepo) UpdateDDNInterruptionOfDeliveryP(ddnintd models.DDNInterruptionOfDelivery) error {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var status int
	var message string

	query := `begin  ddn.synsoft.p_ddn_prekid_isp_update(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24); end;`
	//var int status
	//var string message
	_, err := m.DB.ExecContext(ctx, query,
		ddnintd.IdSMrc,
		ddnintd.IdSTipd,
		ddnintd.IdSVrpd,
		ddnintd.IdTipob,
		ddnintd.ObId,
		ddnintd.Vrepoc,
		ddnintd.Vrezav,
		ddnintd.IdSVrPrek,
		ddnintd.IdSUzrokPrek,
		ddnintd.Snaga,
		ddnintd.Opis,
		ddnintd.KorUneo,
		ddnintd.IdSMernaMesta,
		ddnintd.BrojMesta,
		ddnintd.Ind,
		ddnintd.P2TrafId,
		ddnintd.Bi,
		ddnintd.IdSPoduzrokPrek,
		ddnintd.IdDogPrekidP,
		ddnintd.IdTipObjektaNdc,
		ddnintd.IdTipDogadjajaNdc,
		ddnintd.SynsoftId,
		sql.Out{Dest: &status},
		sql.Out{Dest: &message},
	)

	if err != nil {
		log.Println(err)
		return err
	}
	//fmt.Println(pipiddn.TipMan)
	//fmt.Println(pipiddn.DatSmene)
	//fmt.Println(status)
	//fmt.Println(message)

	if status != 0 {
		return errors.New(message)
	} else {
		return nil
	}
}

func (m *OracleDBRepo) DeleteDDNInterruptionOfDelivery(synsoftId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	stmt := `delete from ddn_prekid_isp_s where synsoft_id = :1`

	_, err := m.DB.ExecContext(ctx, stmt, synsoftId)
	if err != nil {
		return err
	}

	return nil
}

func (m *OracleDBRepo) InsertUpdateDDNInterruptionOfDelivery(ddnintd models.DDNInterruptionOfDeliveryPayload) error {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var status int
	var message string

	query := `begin  ddn.synsoft.ddn_prekid_isp_iu(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16); end;`
	//var int status
	//var string message
	_, err := m.DB.ExecContext(ctx, query,
		ddnintd.IdSMrc,
		ddnintd.IdSTipd,
		ddnintd.IdTipob,
		ddnintd.ObId,
		ddnintd.Vrepoc,
		ddnintd.Vrezav,
		ddnintd.TipDogadjaja,
		ddnintd.Uzrok,
		ddnintd.Snaga,
		ddnintd.Opis,
		ddnintd.KorUneo,
		ddnintd.P2TrafId,
		ddnintd.TipObjekta,
		ddnintd.SynsoftId,
		sql.Out{Dest: &status},
		sql.Out{Dest: &message},
	)

	if err != nil {
		log.Println(err)
		return err
	}
	//fmt.Println(pipiddn.TipMan)
	//fmt.Println(pipiddn.DatSmene)
	//fmt.Println(status)
	//fmt.Println(message)
	if status != 0 {
		return errors.New(message)
	} else {
		return nil
	}
}

func (m *OracleDBRepo) InsertUpdateDDNInterruptionOfDeliveryP(ddnintd models.DDNInterruptionOfDeliveryPayload) error {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var status int
	var message string

	query := `begin  ddn.synsoft.p_ddn_prekid_isp_iu(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16); end;`
	//var int status
	//var string message
	_, err := m.DB.ExecContext(ctx, query,
		ddnintd.IdSMrc,
		ddnintd.IdSTipd,
		ddnintd.IdTipob,
		ddnintd.ObId,
		ddnintd.Vrepoc,
		ddnintd.Vrezav,
		ddnintd.TipDogadjaja,
		ddnintd.Uzrok,
		ddnintd.Snaga,
		ddnintd.Opis,
		ddnintd.KorUneo,
		ddnintd.P2TrafId,
		ddnintd.TipObjekta,
		ddnintd.SynsoftId,
		sql.Out{Dest: &status},
		sql.Out{Dest: &message},
	)

	if err != nil {
		log.Println(err)
		return err
	}
	//fmt.Println(pipiddn.TipMan)
	//fmt.Println(pipiddn.DatSmene)
	//fmt.Println(status)
	//fmt.Println(message)
	if status != 0 {
		return errors.New(message)
	} else {
		return nil
	}
}

func (m *OracleDBRepo) DeleteDDNInterruptionOfDeliveryP(synsoftId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	stmt := `delete from ddn_prekid_isp where synsoft_id = :1`

	_, err := m.DB.ExecContext(ctx, stmt, synsoftId)
	if err != nil {
		return err
	}

	return nil
}

func (m *OracleDBRepo) GetDDNInterruptionOfDeliveryNDC() ([]*models.DDNInterruptionOfDelivery, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select ID_S_MRC,
	COALESCE(to_char(ID_S_TIPD), ''),
	COALESCE(to_char(ID_S_VRPD), ''),
	COALESCE(to_char(ID_TIPOB), ''),
	COALESCE(to_char(OB_ID), ''),
   to_char(VREPOC, 'dd.mm.yyyy HH24:MI:SS'),
   to_char(VREZAV, 'dd.mm.yyyy HH24:MI:SS'),
   COALESCE(to_char(ID_S_VR_PREK), ''),
   COALESCE(to_char(ID_S_UZROK_PREK), ''),
   COALESCE(to_char(SNAGA), ''),
   COALESCE(OPIS, ''),
   COALESCE(DDN_KOR, ''),
   COALESCE(to_char(ID_S_MERNA_MESTA), ''),
   COALESCE(to_char(BROJ_MMESTA), ''),
   COALESCE(IND, ''),
   COALESCE(to_char(ID_P2_TRAF), ''),
   COALESCE(to_char(BI), ''),
   COALESCE(to_char(ID_S_PODUZROK_PREK), ''),
   COALESCE(to_char(ID_DOG_PREKID_P), ''),
   COALESCE(to_char(ID_TIP_OBJEKTA_NDC), ''),
   COALESCE(to_char(ID_TIP_DOGADJAJA_NDC), ''),
   COALESCE(SYNSOFT_ID, '')
   from ddn_prekid_isp_s
   where id_s_mrc=8`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var p []*models.DDNInterruptionOfDelivery

	for rows.Next() {
		var ue models.DDNInterruptionOfDelivery
		err := rows.Scan(
			&ue.IdSMrc,
			&ue.IdSTipd,
			&ue.IdSVrpd,
			&ue.IdTipob,
			&ue.ObId,
			&ue.Vrepoc,
			&ue.Vrezav,
			&ue.IdSVrPrek,
			&ue.IdSUzrokPrek,
			&ue.Snaga,
			&ue.Opis,
			&ue.KorUneo,
			&ue.IdSMernaMesta,
			&ue.BrojMesta,
			&ue.Ind,
			&ue.P2TrafId,
			&ue.Bi,
			&ue.IdSPoduzrokPrek,
			&ue.IdDogPrekidP,
			&ue.IdTipObjektaNdc,
			&ue.IdTipDogadjajaNdc,
			&ue.SynsoftId,
		)

		if err != nil {
			return nil, err
		}

		p = append(p, &ue)
	}

	return p, nil
}

func (m *OracleDBRepo) GetDDNInterruptionOfDeliveryNDCP() ([]*models.DDNInterruptionOfDelivery, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select ID_S_MRC,
	COALESCE(to_char(ID_S_TIPD), ''),
	COALESCE(to_char(ID_S_VRPD), ''),
	COALESCE(to_char(ID_TIPOB), ''),
	COALESCE(to_char(OB_ID), ''),
   to_char(VREPOC, 'dd.mm.yyyy HH24:MI:SS'),
   to_char(VREZAV, 'dd.mm.yyyy HH24:MI:SS'),
   COALESCE(to_char(ID_S_VR_PREK), ''),
   COALESCE(to_char(ID_S_UZROK_PREK), ''),
   COALESCE(to_char(SNAGA), ''),
   COALESCE(OPIS, ''),
   COALESCE(DDN_KOR, ''),
   COALESCE(to_char(ID_S_MERNA_MESTA), ''),
   COALESCE(to_char(BROJ_MMESTA), ''),
   COALESCE(IND, ''),
   COALESCE(to_char(ID_P2_TRAF), ''),
   COALESCE(to_char(BI), ''),
   COALESCE(to_char(ID_S_PODUZROK_PREK), ''),
   COALESCE(to_char(ID_DOG_PREKID_P), ''),
   COALESCE(to_char(ID_TIP_OBJEKTA_NDC), ''),
   COALESCE(to_char(ID_TIP_DOGADJAJA_NDC), ''),
   COALESCE(SYNSOFT_ID, '')
   from ddn_prekid_isp
   where id_s_mrc=8`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var p []*models.DDNInterruptionOfDelivery

	for rows.Next() {
		var ue models.DDNInterruptionOfDelivery
		err := rows.Scan(
			&ue.IdSMrc,
			&ue.IdSTipd,
			&ue.IdSVrpd,
			&ue.IdTipob,
			&ue.ObId,
			&ue.Vrepoc,
			&ue.Vrezav,
			&ue.IdSVrPrek,
			&ue.IdSUzrokPrek,
			&ue.Snaga,
			&ue.Opis,
			&ue.KorUneo,
			&ue.IdSMernaMesta,
			&ue.BrojMesta,
			&ue.Ind,
			&ue.P2TrafId,
			&ue.Bi,
			&ue.IdSPoduzrokPrek,
			&ue.IdDogPrekidP,
			&ue.IdTipObjektaNdc,
			&ue.IdTipDogadjajaNdc,
			&ue.SynsoftId,
		)

		if err != nil {
			return nil, err
		}

		p = append(p, &ue)
	}

	return p, nil
}

func (m *OracleDBRepo) GetDDNInterruptionOfDeliveryNDCByID(synsoftId string) (*models.DDNInterruptionOfDelivery, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select ID_S_MRC,
	COALESCE(to_char(ID_S_TIPD), ''),
	COALESCE(to_char(ID_S_VRPD), ''),
	COALESCE(to_char(ID_TIPOB), ''),
	COALESCE(to_char(OB_ID), ''),
   to_char(VREPOC,'dd.mm.yyyy HH24:MI:SS'),
   to_char(VREZAV,'dd.mm.yyyy HH24:MI:SS'),
   COALESCE(to_char(ID_S_VR_PREK), ''),
   COALESCE(to_char(ID_S_UZROK_PREK), ''),
   COALESCE(to_char(SNAGA), ''),
   COALESCE(OPIS, ''),
   COALESCE(DDN_KOR, ''),
   COALESCE(to_char(ID_S_MERNA_MESTA), ''),
   COALESCE(to_char(BROJ_MMESTA), ''),
   COALESCE(IND, ''),
   COALESCE(to_char(ID_P2_TRAF), ''),
   COALESCE(to_char(BI), ''),
   COALESCE(to_char(ID_S_PODUZROK_PREK), ''),
   COALESCE(to_char(ID_DOG_PREKID_P), ''),
   COALESCE(to_char(ID_TIP_OBJEKTA_NDC), ''),
   COALESCE(to_char(ID_TIP_DOGADJAJA_NDC), ''),
   COALESCE(SYNSOFT_ID, '')
   from ddn_prekid_isp_s
   where synsoft_id=:1`

	row := m.DB.QueryRowContext(ctx, query, synsoftId)

	var ue models.DDNInterruptionOfDelivery
	err := row.Scan(
		&ue.IdSMrc,
		&ue.IdSTipd,
		&ue.IdSVrpd,
		&ue.IdTipob,
		&ue.ObId,
		&ue.Vrepoc,
		&ue.Vrezav,
		&ue.IdSVrPrek,
		&ue.IdSUzrokPrek,
		&ue.Snaga,
		&ue.Opis,
		&ue.KorUneo,
		&ue.IdSMernaMesta,
		&ue.BrojMesta,
		&ue.Ind,
		&ue.P2TrafId,
		&ue.Bi,
		&ue.IdSPoduzrokPrek,
		&ue.IdDogPrekidP,
		&ue.IdTipObjektaNdc,
		&ue.IdTipDogadjajaNdc,
		&ue.SynsoftId,
	)

	if err != nil {
		return nil, err
	}

	return &ue, nil
}

func (m *OracleDBRepo) GetDDNInterruptionOfDeliveryNDCByIDP(synsoftId string) (*models.DDNInterruptionOfDelivery, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select ID_S_MRC,
	COALESCE(to_char(ID_S_TIPD), ''),
	COALESCE(to_char(ID_S_VRPD), ''),
	COALESCE(to_char(ID_TIPOB), ''),
	COALESCE(to_char(OB_ID), ''),
   to_char(VREPOC,'dd.mm.yyyy HH24:MI:SS'),
   to_char(VREZAV,'dd.mm.yyyy HH24:MI:SS'),
   COALESCE(to_char(ID_S_VR_PREK), ''),
   COALESCE(to_char(ID_S_UZROK_PREK), ''),
   COALESCE(to_char(SNAGA), ''),
   COALESCE(OPIS, ''),
   COALESCE(DDN_KOR, ''),
   COALESCE(to_char(ID_S_MERNA_MESTA), ''),
   COALESCE(to_char(BROJ_MMESTA), ''),
   COALESCE(IND, ''),
   COALESCE(to_char(ID_P2_TRAF), ''),
   COALESCE(to_char(BI), ''),
   COALESCE(to_char(ID_S_PODUZROK_PREK), ''),
   COALESCE(to_char(ID_DOG_PREKID_P), ''),
   COALESCE(to_char(ID_TIP_OBJEKTA_NDC), ''),
   COALESCE(to_char(ID_TIP_DOGADJAJA_NDC), ''),
   COALESCE(SYNSOFT_ID, '')
   from ddn_prekid_isp
   where synsoft_id=:1`

	row := m.DB.QueryRowContext(ctx, query, synsoftId)

	var ue models.DDNInterruptionOfDelivery
	err := row.Scan(
		&ue.IdSMrc,
		&ue.IdSTipd,
		&ue.IdSVrpd,
		&ue.IdTipob,
		&ue.ObId,
		&ue.Vrepoc,
		&ue.Vrezav,
		&ue.IdSVrPrek,
		&ue.IdSUzrokPrek,
		&ue.Snaga,
		&ue.Opis,
		&ue.KorUneo,
		&ue.IdSMernaMesta,
		&ue.BrojMesta,
		&ue.Ind,
		&ue.P2TrafId,
		&ue.Bi,
		&ue.IdSPoduzrokPrek,
		&ue.IdDogPrekidP,
		&ue.IdTipObjektaNdc,
		&ue.IdTipDogadjajaNdc,
		&ue.SynsoftId,
	)

	if err != nil {
		return nil, err
	}

	return &ue, nil
}

/** start Check for PI_DOK **/
func (m *OracleDBRepo) CheckForPiDokYesterdayP(datIzv string, idSMrc int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	var num int
	stmt := `select count(*) from pgi.pi_dok where id_s_mrc = :1 and datizv = to_date(:2,'dd.mm.yyyy')-1 and tx_rx = 'O'`

	err := m.DB.QueryRowContext(ctx, stmt, idSMrc, datIzv).Scan(&num)
	if err != nil {
		return -1, err
	}
	return num, nil
}

func (m *OracleDBRepo) CheckForPiDokTodayP(datIzv string, idSMrc int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	var num int
	stmt := `select count(*) from pgi.pi_dok where id_s_mrc = :1 and datizv = to_date(:2,'dd.mm.yyyy') and tx_rx = 'O'`

	err := m.DB.QueryRowContext(ctx, stmt, idSMrc, datIzv).Scan(&num)
	if err != nil {
		return -1, err
	}
	return num, nil
}

func (m *OracleDBRepo) ClosePgiP(datIzv string, idSMrc string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	var status int
	var message string

	query := `begin  ddn.synsoft.zatvori_pogizv(:1, :2, :3, :4); end;`

	_, err := m.DB.ExecContext(ctx, query,
		datIzv,
		idSMrc,
		sql.Out{Dest: &status},
		sql.Out{Dest: &message},
	)

	if err != nil {
		log.Println(err)
		return err
	}
	//fmt.Println(pipiddn.TipMan)
	//fmt.Println(pipiddn.DatSmene)
	//fmt.Println(status)
	//fmt.Println(message)
	if status != 0 {
		return errors.New(message)
	} else {
		return nil
	}
}

func (m *OracleDBRepo) TransferInPgiP(datIzv string, idSMrc string, Tip string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	var status int
	var message string

	query := `begin  ddn.synsoft.prenos_u_pi(:1, :2, :3, :4, :5); end;`

	_, err := m.DB.ExecContext(ctx, query,
		datIzv,
		idSMrc,
		Tip,
		sql.Out{Dest: &status},
		sql.Out{Dest: &message},
	)

	if err != nil {
		log.Println(err)
		return err
	}
	//fmt.Println(pipiddn.TipMan)
	//fmt.Println(pipiddn.DatSmene)
	//fmt.Println(status)
	//fmt.Println(message)
	if status != 0 {
		return errors.New(message)
	} else {
		return nil
	}
}

/** end Check for PI_DOK **/

/** start NOVITA ***/
func (m *OracleDBRepo) GetAllUnbalancedTrader() ([]*models.UnbalancedTrader, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select COALESCE(SIFRA_TRG, ''),
			COALESCE(to_char(ODSTUPANJE), '')
   			from SYNSOFT_NEIZB_TRG`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var p []*models.UnbalancedTrader

	for rows.Next() {
		var ue models.UnbalancedTrader
		err := rows.Scan(
			&ue.Code,
			&ue.Deviation,
		)

		if err != nil {
			return nil, err
		}

		p = append(p, &ue)
	}

	return p, nil
}

/** end NOVITA ***/

/*** start with GIS ***/

/**** end GIS ***/

/*** start TDN***/

func (m *OracleDBRepo) InsertD2D3Dozvola(d *models.D2D3Dozvola) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	tx, err := m.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Insert u AAMS_DOZVOLE
	queryDozvola := `
		INSERT INTO TDN.AAMS_D2D3_DOZVOLE (
             AAMS_D2D3_ID, BROJ_DOZVOLE, TIP_DOZVOLE , DATPOC, DATZAV, STRUCNO_LICE,
             STATUS , VOZILO , TELEFON , DATPRI, DATIZM, ID_OSNOVNE_DOZ , BROJ_OSNOVNE_DOZ
        ) VALUES (
            :1, :2, :3, TO_DATE(:4, 'DD.MM.YYYY HH24:MI'), TO_DATE(:5, 'DD.MM.YYYY HH24:MI'),
			:6, :7, :8, :9, SYSDATE, SYSDATE,:10,:11
        )
	`

	_, err = tx.ExecContext(ctx, queryDozvola,
		d.DozvolaID,
		d.BrojDozvole,
		d.TipDozvole,
		d.DatumPocetka,
		d.DatumZavrsetka,
		d.StrucnoLice,
		d.Status,
		d.RegistracijaVozilo,
		d.KontaktTelefon,
		d.OsnovnaDozvolaID,
		d.BrojOsnovneDozole,
	)
	if err != nil {
		return fmt.Errorf("insert dozvola: %w", err)
	}

	// Insert lica u AAMS_D2D3_LICA
	for _, lice := range d.Lica {

		queryLice := `
			INSERT INTO TDN.AAMS_D2D3_LICA (
				AAMS_LICA_ID, AAMS_D2D3_DOZ_ID, LICE_NAZIV, PREDUZECE, LICNA_KARTA
			) VALUES (:1, :2, :3, :4, :5)
		`

		_, err := tx.ExecContext(ctx, queryLice,
			lice.LiceID,
			d.DozvolaID,
			lice.Ime,
			lice.NazivPreduzeca,
			lice.BrLicneKarte,
		)
		if err != nil {
			return fmt.Errorf("insert lice: %w", err)
		}
	}

	// Insert objekata u AAMS_DOZ_OBJEKTI
	for _, obj := range d.Objekti {

		queryObj := `
			INSERT INTO TDN.AAMS_D2D3_OBJEKTI (
				AAMS_D2D3_DOZ_ID,  AAMS_OBJEKAT
			) VALUES (:1, :2)
		`

		_, err := tx.ExecContext(ctx, queryObj,
			d.DozvolaID,
			obj,
		)
		if err != nil {
			return fmt.Errorf("insert objekat: %w", err)
		}
	}

	// Commit
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit failed: %w", err)
	}

	return nil
}

func (m *OracleDBRepo) InsertDozvola(d *models.Dozvola) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	tx, err := m.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Insert u AAMS_DOZVOLE
	queryDozvola := `
		INSERT INTO TDN.AAMS_DOZVOLE (
            AAMS_DOZVOLA_ID,  BROJ_DOZVOLE, TIP_DOZVOLE, DATPOC, DATZAV, 
            DATUM_POSETE, PRIMA_LICE, VOZILO, TELEFON, DATPRI, DATIZM, RAZLOG_POSETE, STATUS
        ) VALUES (
            :1, :2, :3,  TO_DATE(:4, 'DD.MM.YYYY HH24:MI'), TO_DATE(:5, 'DD.MM.YYYY HH24:MI'), 
            TO_DATE(:6, 'DD.MM.YYYY'), :7, :8, :9, SYSDATE, SYSDATE, :10, :11
        )
	`

	_, err = tx.ExecContext(ctx, queryDozvola,
		d.DozvolaID,
		d.BrojDozvole,
		d.TipDozvole,
		d.DatumPocetka,
		d.DatumZavrsetka,
		d.DatumPosete,
		d.Primalac,
		d.RegistracijaVozilo,
		d.KontaktTelefon,
		d.RazlogPosete,
		d.Status,
	)
	if err != nil {
		return fmt.Errorf("insert dozvola: %w", err)
	}

	// Insert lica u AAMS_DOZ_LICA
	for _, lice := range d.Lica {

		queryLice := `
			INSERT INTO TDN.AAMS_DOZ_LICA (
				AAMS_LICA_ID, AAMS_DOZVOLA_ID, LICE_NAZIV, PREDUZECE, LICNA_KARTA
			) VALUES (:1, :2, :3, :4, :5)
		`

		_, err := tx.ExecContext(ctx, queryLice,
			lice.LiceID,
			d.DozvolaID,
			lice.Ime,
			lice.NazivPreduzeca,
			lice.BrLicneKarte,
		)
		if err != nil {
			return fmt.Errorf("insert lice: %w", err)
		}
	}

	// Insert objekata u AAMS_DOZ_OBJEKTI
	for _, obj := range d.Objekti {

		queryObj := `
			INSERT INTO TDN.AAMS_DOZ_OBJEKTI (
				AAMS_DOZVOLA_ID,  AAMS_OBJEKAT
			) VALUES (:1, :2)
		`

		_, err := tx.ExecContext(ctx, queryObj,
			d.DozvolaID,
			obj,
		)
		if err != nil {
			return fmt.Errorf("insert objekat: %w", err)
		}
	}

	// Commit
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit failed: %w", err)
	}

	return nil
}

func (m *OracleDBRepo) GetAllDozvola() ([]*models.Dozvola, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select COALESCE(AAMS_DOZVOLA_ID, ''),
  				COALESCE(BROJ_DOZVOLE, ''),
  				COALESCE(TIP_DOZVOLE, ''),
				to_char(DATPOC, 'dd.mm.yyyy HH24:MI:SS'),
  				to_char(DATZAV, 'dd.mm.yyyy HH24:MI:SS'),
  				to_char(DATUM_POSETE, 'dd.mm.yyyy'),
  				COALESCE(PRIMA_LICE, '') ,
  				COALESCE(VOZILO, ''),
  				COALESCE(TELEFON, '') ,
  				COALESCE(RAZLOG_POSETE, ''),
  				COALESCE(STATUS, '') 
				FROM TDN.AAMS_DOZVOLE`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		// fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var ds = []*models.Dozvola{}

	for rows.Next() {
		var d models.Dozvola
		err := rows.Scan(
			&d.DozvolaID,
			&d.BrojDozvole,
			&d.TipDozvole,
			&d.DatumPocetka,
			&d.DatumZavrsetka,
			&d.DatumPosete,
			&d.Primalac,
			&d.RegistracijaVozilo,
			&d.KontaktTelefon,
			&d.RazlogPosete,
			&d.Status,
		)

		if err != nil {
			return nil, err
		}

		/***** start part for objects ****/

		query1 := `select o.aams_objekat 
				from aams_doz_objekti o 
				where o.aams_dozvola_id  = :1
			`
		var objs = []string{}

		rows1, err := m.DB.QueryContext(ctx, query1, d.DozvolaID)
		if err != nil {
			return nil, err
		}

		for rows1.Next() {
			var ob string
			err := rows1.Scan(
				&ob,
			)

			if err != nil {
				return nil, err
			}
			objs = append(objs, ob)
		}
		rows1.Close()
		d.Objekti = objs

		/***** end part for objects ******/

		/***** start part for lica ****/

		query2 := `select l.aams_lica_id, l.lice_naziv,l.preduzece, l.licna_karta 
				from aams_doz_lica l 
				where l.aams_dozvola_id  = :1
			`
		var lica = []models.Lice{}

		rows2, err := m.DB.QueryContext(ctx, query2, d.DozvolaID)
		if err != nil {
			rows1.Close()
			return nil, err
		}

		for rows2.Next() {
			var l models.Lice
			err := rows2.Scan(
				&l.LiceID,
				&l.Ime,
				&l.NazivPreduzeca,
				&l.BrLicneKarte,
			)

			if err != nil {
				rows2.Close()
				return nil, err
			}
			lica = append(lica, l)
		}
		rows2.Close()
		d.Lica = lica

		/***** end part for lica ******/

		ds = append(ds, &d)
	}

	return ds, nil
}

func (m *OracleDBRepo) GetAllD2D3Dozvola() ([]*models.D2D3Dozvola, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select COALESCE(AAMS_D2D3_ID, ''),
                  COALESCE(BROJ_DOZVOLE, ''),
                  COALESCE(TIP_DOZVOLE, ''),
                  to_char(DATPOC, 'dd.mm.yyyy HH24:MI:SS'),
                  to_char(DATZAV, 'dd.mm.yyyy HH24:MI:SS'),
                  COALESCE(STRUCNO_LICE, '') ,
                  COALESCE(STATUS, '') ,
                  COALESCE(VOZILO, ''),
                  COALESCE(TELEFON, '') ,
                  COALESCE(ID_OSNOVNE_DOZ, ''),
                  COALESCE(BROJ_OSNOVNE_DOZ, '')
                FROM TDN.AAMS_D2D3_DOZVOLE`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		// fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var ds = []*models.D2D3Dozvola{}

	for rows.Next() {
		var d models.D2D3Dozvola
		err := rows.Scan(
			&d.DozvolaID,
			&d.BrojDozvole,
			&d.TipDozvole,
			&d.DatumPocetka,
			&d.DatumZavrsetka,
			&d.StrucnoLice,
			&d.Status,
			&d.RegistracijaVozilo,
			&d.KontaktTelefon,
			&d.OsnovnaDozvolaID,
			&d.BrojOsnovneDozole,
		)

		if err != nil {
			return nil, err
		}

		/***** start part for objects ****/

		query1 := `select o.aams_objekat 
				from aams_d2d3_objekti o 
				where o.aams_d2d3_doz_id  = :1
			`
		var objs = []string{}

		rows1, err := m.DB.QueryContext(ctx, query1, d.DozvolaID)
		if err != nil {
			return nil, err
		}

		for rows1.Next() {
			var ob string
			err := rows1.Scan(
				&ob,
			)

			if err != nil {
				return nil, err
			}
			objs = append(objs, ob)
		}
		rows1.Close()
		d.Objekti = objs

		/***** end part for objects ******/

		/***** start part for lica ****/

		query2 := `select l.aams_lica_id, l.lice_naziv,l.preduzece, l.licna_karta 
				from aams_d2d3_lica l 
				where l.aams_d2d3_doz_id  = :1
			`
		var lica = []models.Lice{}

		rows2, err := m.DB.QueryContext(ctx, query2, d.DozvolaID)
		if err != nil {
			rows1.Close()
			return nil, err
		}

		for rows2.Next() {
			var l models.Lice
			err := rows2.Scan(
				&l.LiceID,
				&l.Ime,
				&l.NazivPreduzeca,
				&l.BrLicneKarte,
			)

			if err != nil {
				rows2.Close()
				return nil, err
			}
			lica = append(lica, l)
		}
		rows2.Close()
		d.Lica = lica

		/***** end part for lica ******/

		ds = append(ds, &d)
	}

	return ds, nil
}

func (m *OracleDBRepo) GetByIdD2D3Dozvola(id string) (*models.D2D3Dozvola, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select COALESCE(AAMS_D2D3_ID, ''),
                  COALESCE(BROJ_DOZVOLE, ''),
                  COALESCE(TIP_DOZVOLE, ''),
                  to_char(DATPOC, 'dd.mm.yyyy HH24:MI:SS'),
                  to_char(DATZAV, 'dd.mm.yyyy HH24:MI:SS'),
                  COALESCE(STRUCNO_LICE, '') ,
                  COALESCE(STATUS, '') ,
                  COALESCE(VOZILO, ''),
                  COALESCE(TELEFON, '') ,
                  COALESCE(ID_OSNOVNE_DOZ, ''),
                  COALESCE(BROJ_OSNOVNE_DOZ, '')
                FROM TDN.AAMS_D2D3_DOZVOLE
				WHERE AAMS_D2D3_ID=:1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var d models.D2D3Dozvola
	err := row.Scan(
		&d.DozvolaID,
		&d.BrojDozvole,
		&d.TipDozvole,
		&d.DatumPocetka,
		&d.DatumZavrsetka,
		&d.StrucnoLice,
		&d.Status,
		&d.RegistracijaVozilo,
		&d.KontaktTelefon,
		&d.OsnovnaDozvolaID,
		&d.BrojOsnovneDozole,
	)
	if err != nil {
		// fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}

	/***** start part for objects ****/

	query1 := `select o.aams_objekat 
				from aams_d2d3_objekti o 
				where o.aams_d2d3_doz_id  = :1
			`
	var objs = []string{}

	rows1, err := m.DB.QueryContext(ctx, query1, d.DozvolaID)
	if err != nil {
		return nil, err
	}

	for rows1.Next() {
		var ob string
		err := rows1.Scan(
			&ob,
		)

		if err != nil {
			return nil, err
		}
		objs = append(objs, ob)
	}
	rows1.Close()
	d.Objekti = objs

	/***** end part for objects ******/

	/***** start part for lica ****/

	query2 := `select l.aams_lica_id, l.lice_naziv,l.preduzece, l.licna_karta 
				from aams_d2d3_lica l 
				where l.aams_d2d3_doz_id  = :1
			`
	var lica = []models.Lice{}

	rows2, err := m.DB.QueryContext(ctx, query2, d.DozvolaID)
	if err != nil {
		rows1.Close()
		return nil, err
	}

	for rows2.Next() {
		var l models.Lice
		err := rows2.Scan(
			&l.LiceID,
			&l.Ime,
			&l.NazivPreduzeca,
			&l.BrLicneKarte,
		)

		if err != nil {
			rows2.Close()
			return nil, err
		}
		lica = append(lica, l)
	}
	rows2.Close()
	d.Lica = lica

	/***** end part for lica ******/

	return &d, nil
}

func (m *OracleDBRepo) GetByIdDozvola(id string) (*models.Dozvola, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `select COALESCE(AAMS_DOZVOLA_ID, ''),
  				COALESCE(BROJ_DOZVOLE, ''),
  				COALESCE(TIP_DOZVOLE, ''),
				to_char(DATPOC, 'dd.mm.yyyy HH24:MI:SS'),
  				to_char(DATZAV, 'dd.mm.yyyy HH24:MI:SS'),
  				to_char(DATUM_POSETE, 'dd.mm.yyyy'),
  				COALESCE(PRIMA_LICE, '') ,
  				COALESCE(VOZILO, ''),
  				COALESCE(TELEFON, '') ,
  				COALESCE(RAZLOG_POSETE, ''),
  				COALESCE(STATUS, '') 
				FROM TDN.AAMS_DOZVOLE
				WHERE AAMS_DOZVOLA_ID=:1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var d models.Dozvola
	err := row.Scan(
		&d.DozvolaID,
		&d.BrojDozvole,
		&d.TipDozvole,
		&d.DatumPocetka,
		&d.DatumZavrsetka,
		&d.DatumPosete,
		&d.Primalac,
		&d.RegistracijaVozilo,
		&d.KontaktTelefon,
		&d.RazlogPosete,
		&d.Status,
	)
	if err != nil {
		// fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}

	/***** start part for objects ****/

	query1 := `select o.aams_objekat 
				from aams_doz_objekti o 
				where o.aams_dozvola_id  = :1
			`
	var objs = []string{}

	rows1, err := m.DB.QueryContext(ctx, query1, d.DozvolaID)
	if err != nil {
		return nil, err
	}

	for rows1.Next() {
		var ob string
		err := rows1.Scan(
			&ob,
		)

		if err != nil {
			return nil, err
		}
		objs = append(objs, ob)
	}
	rows1.Close()
	d.Objekti = objs

	/***** end part for objects ******/

	/***** start part for lica ****/

	query2 := `select l.aams_lica_id, l.lice_naziv,l.preduzece, l.licna_karta 
				from aams_doz_lica l 
				where l.aams_dozvola_id  = :1
			`
	var lica = []models.Lice{}

	rows2, err := m.DB.QueryContext(ctx, query2, d.DozvolaID)
	if err != nil {
		rows1.Close()
		return nil, err
	}

	for rows2.Next() {
		var l models.Lice
		err := rows2.Scan(
			&l.LiceID,
			&l.Ime,
			&l.NazivPreduzeca,
			&l.BrLicneKarte,
		)

		if err != nil {
			rows2.Close()
			return nil, err
		}
		lica = append(lica, l)
	}
	rows2.Close()
	d.Lica = lica

	/***** end part for lica ******/

	return &d, nil
}

func (m *OracleDBRepo) DeleteD2D3DozvolaByID(dozvolaID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	tx, err := m.DB.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}

	// 1. Briši lica
	_, err = tx.Exec(`DELETE FROM TDN.AAMS_D2D3_LICA WHERE AAMS_D2D3_DOZ_ID = :1`, dozvolaID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error deleting lica: %w", err)
	}

	// 2. Briši objekte
	_, err = tx.Exec(`DELETE FROM TDN.AAMS_D2D3_OBJEKTI WHERE AAMS_D2D3_DOZ_ID = :1`, dozvolaID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error deleting objekti: %w", err)
	}

	// 3. Briši dozvolu
	_, err = tx.Exec(`DELETE FROM TDN.AAMS_D2D3_DOZVOLE WHERE AAMS_D2D3_ID = :1`, dozvolaID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error deleting dozvola: %w", err)
	}

	// Commit transakcije
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	return nil
}

func (m *OracleDBRepo) DeleteDozvolaByID(dozvolaID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	tx, err := m.DB.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}

	// 1. Briši lica
	_, err = tx.Exec(`DELETE FROM TDN.AAMS_DOZ_LICA WHERE AAMS_DOZVOLA_ID = :1`, dozvolaID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error deleting lica: %w", err)
	}

	// 2. Briši objekte
	_, err = tx.Exec(`DELETE FROM TDN.AAMS_DOZ_OBJEKTI WHERE AAMS_DOZVOLA_ID = :1`, dozvolaID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error deleting objekti: %w", err)
	}

	// 3. Briši dozvolu
	_, err = tx.Exec(`DELETE FROM TDN.AAMS_DOZVOLE WHERE AAMS_DOZVOLA_ID = :1`, dozvolaID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error deleting dozvola: %w", err)
	}

	// Commit transakcije
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	return nil
}

func (m *OracleDBRepo) GetD2D3DozvolaById(id string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `SELECT COUNT(*) FROM TDN.AAMS_D2D3_DOZVOLE WHERE AAMS_D2D3_ID  = :1`

	var count int
	err := m.DB.QueryRowContext(ctx, query, id).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (m *OracleDBRepo) GetDozvolaById(id string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `SELECT COUNT(*) FROM TDN.AAMS_DOZVOLE WHERE AAMS_DOZVOLA_ID = :1`

	var count int
	err := m.DB.QueryRowContext(ctx, query, id).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (m *OracleDBRepo) InsertLog(operacija, status, poruka string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `
		INSERT INTO TDN.AAMS_DOZ_LOG (
			OPERACIJA, STATUS, PORUKA, CREATED_AT
		) VALUES (:1, :2, :3,  SYSDATE)
	`

	_, err := m.DB.ExecContext(ctx, query, operacija, status, poruka)
	if err != nil {
		// log.Println("Greška pri logovanju:", err)
		return fmt.Errorf("log insert error: %w", err)
	}

	return nil
}

/*** end TDN ***/
