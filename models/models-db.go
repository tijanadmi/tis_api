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

func (m *DBModel) OneSignal(id int) (*Signal, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, naziv, status
			  from zsi_info
			  where id=:1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var signal Signal

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

func (m *DBModel) OneMalfunctionIn(id int) (*MalfunctionIn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id_pd_kk,naziv_edd,RED
			  from zas_pd_kk_v
			  where id_pd_kk=:1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var mf MalfunctionIn

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

func (m *DBModel) OneAPU(id int) (*Apu, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id,naziv,SKR_NAZ,SIFRA, status
			  from s_rapu
			  where id=:1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var apu Apu

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

// Get returns all groups of causes and error, if any
func (m *DBModel) GetGroupsOfCauses() ([]*GroupOfCause, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
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

	var grcs []*GroupOfCause

	for rows.Next() {
		var grc GroupOfCause
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

func (m *DBModel) OneGroupOfCauses(id int) (*GroupOfCause, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, sifra, naziv, skr, status, sif_ddn, sortr
			  from pgi.s_gruzr	
			  where id=:1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var grc GroupOfCause

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
func (m *DBModel) GetCauses() ([]*Cause, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
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

	var caus []*Cause

	for rows.Next() {
		var cau Cause
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

func (m *DBModel) OneCause(id int) (*Cause, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select u.id, u.id_s_gruzr, u.sifra, u.naziv, u.skr, u.status, u.sortr, gu.id, gu.sifra, gu.naziv, gu.skr, gu.status, gu.sif_ddn, gu.sortr
			  from pgi.s_uzrok u, pgi.s_gruzr gu
			  where u.id_s_gruzr=gu.id
			  and u.id=:1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var cau Cause

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
func (m *DBModel) GetGroupOfReasons() ([]*GroupOfReason, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
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

	var grs []*GroupOfReason

	for rows.Next() {
		var gr GroupOfReason
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

func (m *DBModel) OneGroupOfReasons(id int) (*GroupOfReason, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, sifra, naziv, skr, status, sortr
	          from PGI.S_GRRAZ		
			  where id=:1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var gr GroupOfReason

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
func (m *DBModel) GetReasons() ([]*Reason, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
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

	var rs []*Reason

	for rows.Next() {
		var r Reason
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

func (m *DBModel) OneReason(id int) (*Reason, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select r.id, r.id_s_grraz, r.sifra, r.naziv, r.skr, r.status, r.sortr,gr.id, gr.sifra, gr.naziv, gr.skr, gr.status, gr.sortr
			  from pgi.s_razlog r, PGI.S_GRRAZ gr
	 		  where r.id_s_grraz=gr.id
			  and r.id=:1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var r Reason

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
func (m *DBModel) GetWeatherConditions() ([]*WeatherCondition, error) {
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

	var signals []*WeatherCondition

	for rows.Next() {
		var signal WeatherCondition
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

func (m *DBModel) OneWeatherCondition(id int) (*WeatherCondition, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, sifra, naziv, status
			  from DDN.S_VREM_USL			
			  where id=:1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var signal WeatherCondition

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
func (m *DBModel) GetCategoriesOfEvents() ([]*CategoryOfEvent, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
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

	var ces []*CategoryOfEvent

	for rows.Next() {
		var ce CategoryOfEvent
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

func (m *DBModel) OneCategoryOfEvents(id int) (*CategoryOfEvent, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, id_s_tipd, sifra, naziv, skr, status, sortr
			  from s_vrpd				
			  where id=:1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var ce CategoryOfEvent

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
func (m *DBModel) GetOHL() ([]*OverheadLine, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
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

	var ohls []*OverheadLine

	for rows.Next() {
		var ohl OverheadLine
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
func (m *DBModel) GetPowerCables() ([]*PowerCable, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
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

	var cabs []*PowerCable

	for rows.Next() {
		var cab PowerCable
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
func (m *DBModel) GetSubstations() ([]*Substation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
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

	var subs []*Substation

	for rows.Next() {
		var sub Substation
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
func (m *DBModel) GetFeeders() ([]*Feeder, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
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

	var feeds []*Feeder

	for rows.Next() {
		var feed Feeder
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
func (m *DBModel) GetProtectionDevices() ([]*ProtectionDevice, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
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

	var prts []*ProtectionDevice

	for rows.Next() {
		var prt ProtectionDevice
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
func (m *DBModel) GetPowerTransformers() ([]*PowerTransformer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select TIS_ID_TRANSFORMATORA, IPS_ID_TRANSFORMATORA, SAP_ID_TRANSFORMATORA, TIS_ID_TRAFOSTANICE, IPS_ID_TRAFOSTANICE, SAP_ID_TRAFOSTANICE, NAZIV_TRAFOSTANICE,
	NAZIV_TRANSFORMATORA, OPIS_TRANSF_PO_KATEGOR, COALESCE(to_char(KATEGORIJA_ID), ''),COALESCE(KATEGORIJA, ''), PRENOSNI_ODNOS, TIS_ID_VN_POLJA, IPD_ID_VN_POLA, SAP_ID_VN_POLA, NAZIV_VN_POLJA,
	COALESCE(to_char(TIS_ID_SN1_POLJA), ''), COALESCE(IPD_ID_SN1_POLA, ''), COALESCE(SAP_ID_SN1_POLA, ''), COALESCE(NAZIV_SN1_POLJA, ''), COALESCE(to_char(TIS_ID_SN2_POLJA), ''), COALESCE(IPD_ID_SN2_POLA, ''), COALESCE(SAP_ID_SN2_POLA, ''), COALESCE(NAZIV_SN2_POLJA, ''), COALESCE(to_char(TIS_ID_T_POLJA), ''),
	COALESCE(IPD_ID_T_POLA, ''), COALESCE(SAP_ID_T_POLA, ''), COALESCE(NAZIV_T_POLJA, '')
 	          from SYNSOFT_TR_A
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var trs []*PowerTransformer

	for rows.Next() {
		var tr PowerTransformer
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
func (m *DBModel) GetDisconnectors() ([]*Disconnector, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
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

	var diss []*Disconnector

	for rows.Next() {
		var dis Disconnector
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
func (m *DBModel) GetWorkPermissions() ([]*WorkPermission, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
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

	var prms []*WorkPermission

	for rows.Next() {
		var prm WorkPermission
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
