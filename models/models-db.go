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
