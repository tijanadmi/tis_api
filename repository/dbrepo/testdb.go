package dbrepo

import (
	"database/sql"
	"errors"

	"github.com/tijanadmi/tis-api/models"
)

type TestDBRepo struct{}

func (m *TestDBRepo) Connection() *sql.DB {
	return nil
}

func (m *TestDBRepo) OneSignal(id int) (*models.Signal, error) {

	var signal models.Signal

	return &signal, nil
}

// Get returns all zas_dv_didf_all_v and error, if any
func (m *TestDBRepo) GetDvDidf() ([]*models.Signal, error) {
	var signals []*models.Signal

	return signals, nil
}

// Get returns all zas_tr_gldif_all_v and error, if any
func (m *TestDBRepo) GetDiffTr() ([]*models.Signal, error) {
	var signals []*models.Signal

	return signals, nil
}

// Get returns all zas_tr_rez_dis_all_v and error, if any
func (m *TestDBRepo) GetDisTrRes() ([]*models.Signal, error) {
	var signals []*models.Signal

	return signals, nil

}

// Get returns all ted.zas_spp_didf_all_v and error, if any
func (m *TestDBRepo) GetDisDiffSp() ([]*models.Signal, error) {
	var signals []*models.Signal

	return signals, nil
}

// Get returns all pgd.zas_pd_kk_v and error, if any
func (m *TestDBRepo) GetMalfunctionIn() ([]*models.MalfunctionIn, error) {

	var ms []*models.MalfunctionIn

	return ms, nil
}

func (m *TestDBRepo) OneMalfunctionIn(id int) (*models.MalfunctionIn, error) {

	var mf models.MalfunctionIn

	return &mf, nil
}

// Get returns all ddn.s_rapu and error, if any
func (m *TestDBRepo) GetAPU() ([]*models.Apu, error) {

	var ms []*models.Apu

	return ms, nil
}

func (m *TestDBRepo) OneAPU(id int) (*models.Apu, error) {

	var apu models.Apu

	return &apu, nil
}

// Get returns all zas_dv_pres_all_v and error, if any
func (m *TestDBRepo) GetOCDV() ([]*models.Signal, error) {

	var signals []*models.Signal

	return signals, nil
}

// Get returns all zas_tr_glprs_all_v and error, if any
func (m *TestDBRepo) GetOCTR12() ([]*models.Signal, error) {

	var signals []*models.Signal

	return signals, nil
}

// Get returns all zas_tr_rez_prs_all_v and error, if any
func (m *TestDBRepo) GetOCTRR() ([]*models.Signal, error) {

	var signals []*models.Signal

	return signals, nil
}

// Get returns all zas_spp_prs_all_v and error, if any
func (m *TestDBRepo) GetOCSP() ([]*models.Signal, error) {

	var signals []*models.Signal

	return signals, nil
}

// Get returns all zas_dv_zmsp_all_v and error, if any
func (m *TestDBRepo) GetEarthfaultOCDV() ([]*models.Signal, error) {

	var signals []*models.Signal

	return signals, nil
}

// Get returns all zas_tr_glzms_all_v and error, if any
func (m *TestDBRepo) GetEarthfaultOCTR() ([]*models.Signal, error) {

	var signals []*models.Signal

	return signals, nil
}

// Get returns all zas_spp_zms_all_v and error, if any
func (m *TestDBRepo) GetEarthfaultOCSP() ([]*models.Signal, error) {

	var signals []*models.Signal

	return signals, nil
}

// Get returns all zas_tr_rez_zms_all_v and error, if any
func (m *TestDBRepo) GetEarthfaultOCTRR() ([]*models.Signal, error) {
	

	var signals []*models.Signal


	return signals, nil
}


// Get returns all zas_dv_uzms_all_v and error, if any
func (m *TestDBRepo) GetDirEarthfaultOC() ([]*models.Signal, error) {

	var signals []*models.Signal

	return signals, nil
}

// Get returns all ZAS_DV_TELE_ALL_V and error, if any
func (m *TestDBRepo) GetTPSendRcdv() ([]*models.Signal, error) {

	var signals []*models.Signal

	return signals, nil
}

// Get returns all zas_prek_all_v and error, if any
func (m *TestDBRepo) GetCircuitbreaker() ([]*models.Signal, error) {

	var signals []*models.Signal

	return signals, nil
}

// Get returns all zas_jps_all_v and error, if any
func (m *TestDBRepo) GetBBPBFtrip() ([]*models.Signal, error) {

	var signals []*models.Signal

	return signals, nil
}

// Get returns all zas_tr_neel_all_v and error, if any
func (m *TestDBRepo) GetNonElectrical() ([]*models.Signal, error) {

	var signals []*models.Signal

	return signals, nil
}

// Get returns all zas_sab_sbr_all_v and error, if any
func (m *TestDBRepo) GetBBPBBtrip() ([]*models.Signal, error) {

	var signals []*models.Signal

	return signals, nil
}

// Get returns all zas_sab_opr_all_v and error, if any
func (m *TestDBRepo) GetBFtrip() ([]*models.Signal, error) {

	var signals []*models.Signal

	return signals, nil
}

// Get returns all groups of causes and error, if any
func (m *TestDBRepo) GetGroupsOfCauses() ([]*models.GroupOfCause, error) {

	var grcs []*models.GroupOfCause

	return grcs, nil
}

func (m *TestDBRepo) OneGroupOfCauses(id int) (*models.GroupOfCause, error) {

	var grc models.GroupOfCause

	return &grc, nil
}

// Get returns all causes and error, if any
func (m *TestDBRepo) GetCauses() ([]*models.Cause, error) {

	var caus []*models.Cause

	return caus, nil
}

func (m *TestDBRepo) OneCause(id int) (*models.Cause, error) {

	var cau models.Cause

	return &cau, nil
}

// Get returns all groups of reasons and error, if any
func (m *TestDBRepo) GetGroupOfReasons() ([]*models.GroupOfReason, error) {

	var grs []*models.GroupOfReason

	return grs, nil
}

func (m *TestDBRepo) OneGroupOfReasons(id int) (*models.GroupOfReason, error) {

	var gr models.GroupOfReason

	return &gr, nil
}

// Get returns all  reasons and error, if any
func (m *TestDBRepo) GetReasons() ([]*models.Reason, error) {

	var rs []*models.Reason

	return rs, nil
}

func (m *TestDBRepo) OneReason(id int) (*models.Reason, error) {

	var r models.Reason
	return &r, nil
}

// Get returns all weather conditions and error, if any
func (m *TestDBRepo) GetWeatherConditions() ([]*models.WeatherCondition, error) {
	var signals []*models.WeatherCondition

	return signals, nil
}

func (m *TestDBRepo) OneWeatherCondition(id int) (*models.WeatherCondition, error) {

	var signal models.WeatherCondition

	return &signal, nil
}

// Get returns all categories of events and error, if any
func (m *TestDBRepo) GetCategoriesOfEvents() ([]*models.CategoryOfEvent, error) {

	var ces []*models.CategoryOfEvent

	return ces, nil
}

func (m *TestDBRepo) OneCategoryOfEvents(id int) (*models.CategoryOfEvent, error) {

	var ce models.CategoryOfEvent

	return &ce, nil
}

// Get returns all weather conditions and error, if any
func (m *TestDBRepo) GetOHL() ([]*models.OverheadLine, error) {

	var ohls []*models.OverheadLine

	return ohls, nil
}

// Get returns all power_cables and error, if any
func (m *TestDBRepo) GetPowerCables() ([]*models.PowerCable, error) {

	var cabs []*models.PowerCable

	return cabs, nil
}

// Get returns all substations and error, if any
func (m *TestDBRepo) GetSubstations() ([]*models.Substation, error) {

	var subs []*models.Substation

	return subs, nil
}

// Get returns all feeders and error, if any
func (m *TestDBRepo) GetFeeders() ([]*models.Feeder, error) {

	var feeds []*models.Feeder

	return feeds, nil
}

// Get returns all protection devices and error, if any
func (m *TestDBRepo) GetProtectionDevices() ([]*models.ProtectionDevice, error) {

	var prts []*models.ProtectionDevice

	return prts, nil
}

// Get returns all power transformers and error, if any
func (m *TestDBRepo) GetPowerTransformers() ([]*models.PowerTransformer, error) {

	var trs []*models.PowerTransformer

	return trs, nil
}

// Get returns all disconnectors and error, if any
func (m *TestDBRepo) GetDisconnectors() ([]*models.Disconnector, error) {

	var diss []*models.Disconnector

	return diss, nil
}

// Get returns all permissions and error, if any
func (m *TestDBRepo) GetWorkPermissions() ([]*models.WorkPermission, error) {

	var prms []*models.WorkPermission

	return prms, nil
}

// Get returns all permissionsAll and error, if any
func (m *TestDBRepo) GetWorkPermissionsAll() ([]*models.WorkPermissionAll, error) {

	var prms []*models.WorkPermissionAll

	return prms, nil
}

// Get returns all Request1Gr and error, if any
func (m *TestDBRepo) GetRequest1Gr() ([]*models.Request1Gr, error) {

	var prms []*models.Request1Gr

	return prms, nil
}

// Get returns all Request2Gr and error, if any
func (m *TestDBRepo) GetRequest2Gr() ([]*models.Request2Gr, error) {

	var prms []*models.Request2Gr

	return prms, nil
}

// Get returns all permissions and error, if any
func (m *TestDBRepo) GetWorkInEENetwork() ([]*models.WorkInEENetwork, error) {

	var prms []*models.WorkInEENetwork

	return prms, nil
}

func (m *TestDBRepo) GetWeather(year string) ([]*models.WeatherData, error) {

	var weatherData []*models.WeatherData

	return weatherData, nil
}

func (m *TestDBRepo) GetWeatherForecast() ([]*models.WeatherData, error) {

	var weatherData []*models.WeatherData

	return weatherData, nil
}

func (m *TestDBRepo) GetWeatherHistory(year string) ([]*models.WeatherDataHistory, error) {

	var weatherDataHistory []*models.WeatherDataHistory

	return weatherDataHistory, nil
}

func (m *TestDBRepo) GetPermissions1(year string) ([]*models.Permission, error) {

	var permissions []*models.Permission

	return permissions, nil
}

func (m *TestDBRepo) GetPermissions23(year string) ([]*models.Permission, error) {

	var permissions []*models.Permission

	return permissions, nil
}

func (m *TestDBRepo) GetRequests1(year string) ([]*models.Request, error) {

	var requests []*models.Request

	return requests, nil
}

func (m *TestDBRepo) GetRequests23(year string) ([]*models.Request, error) {

	var requests []*models.Request

	return requests, nil
}

func (m *TestDBRepo) GetOutages(year string) ([]*models.Outage, error) {

	var o []*models.Outage

	return o, nil
}

func (m *TestDBRepo) GetExclusions(year string) ([]*models.Exclusion, error) {

	var d []*models.Exclusion

	return d, nil
}

func (m *TestDBRepo) GetPlans(year string) ([]*models.Plan, error) {

	var p []*models.Plan

	return p, nil
}

func (m *TestDBRepo) GetUnopenedPermitForDay(day string, org string) ([]*models.UnopenedPermit, error) {

	var p []*models.UnopenedPermit

	return p, nil
}

// Authenticate authenticates a user
func (m *TestDBRepo) Authenticate(username, testPassword string) error {

	//var user models.User

	return nil
}

func (m *TestDBRepo) GetUserByUsername(username string) (*models.User, error) {

	if username == "test" {
		roles := []models.UserRole{
			{ID: 1, RoleName: "NDC"},
			{ID: 2, RoleName: "DWH"},
		}
		user := models.User{
			ID:       1,
			Username: "test",
			Password: "$2a$12$5ls4XYB69XyIJtl.qPeh0.AwdzQ7pfo3Py3J4Ur.0dtzJ8Sc.F0AS",
			UserRole: roles,
		}
		return &user, nil
	}
	return nil, errors.New("not found")
}

func (m *TestDBRepo) GetUserByID(id int) (*models.User, error) {

	var user = models.User{
		ID: 1,
	}

	return &user, nil
}

func (m *TestDBRepo) InsertPiPiDDNIsklj(pipiddn models.PiPiDDNIsklj) error {

	return nil
}
func (m *TestDBRepo) InsertPiPiDDNIskljP(pipiddn models.PiPiDDNIsklj) error {

	return nil
}

func (m *TestDBRepo) UpdatePiPiDDNIsklj(pipiddn models.PiPiDDNIsklj) error {
	if pipiddn.SynsoftId == "1" {
		return nil
	}
	return errors.New("update failed - no PiPiDDN found")
}

func (m *TestDBRepo) UpdatePiPiDDNIskljP(pipiddn models.PiPiDDNIsklj) error {
	if pipiddn.SynsoftId == "1" {
		return nil
	}
	return errors.New("update failed - no PiPiDDN found")
}

func (m *TestDBRepo) GetAllPiPiDDN() ([]*models.PiPiDDN, error) {

	var p []*models.PiPiDDN

	return p, nil
}

func (m *TestDBRepo) GetAllPiPiDDNP() ([]*models.PiPiDDN, error) {

	var p []*models.PiPiDDN

	return p, nil
}

func (m *TestDBRepo) GetPiPiDDNByID(synsoftId string) (*models.PiPiDDN, error) {

	var pipiddn models.PiPiDDN
	if synsoftId == "1" {
		pipiddn = models.PiPiDDN{
			Datizv:         "27.04.2023",
			IdSMrc:         "8",
			IdSTipd:        "1",
			IdSVrpd:        "1",
			IdTipob:        "1",
			ObId:           "354",
			TrafoId:        "248",
			Vrepoc:         "27.04.2023 00:04:00",
			Vrezav:         "",
			Id1SGruzr:      "9",
			Id1SUzrok:      "70",
			Snaga:          "",
			Opis:           "this is outage put api",
			IdSNap:         "7",
			P2TrafId:       "3625",
			KorUneo:        "DTOMIC",
			IdZDsdfGL1:     "54",
			IdZKvarGL1:     "",
			IdZRapuGL1:     "",
			IdZPrstGL1:     "",
			IdZZmspGL1:     "",
			IdZUzmsGL1:     "",
			ZLokkGL1:       "",
			IdZDsdfGL2:     "54",
			IdZKvarGL2:     "",
			IdZRapuGL2:     "",
			IdZPrstGL2:     "",
			IdZZmspGL2:     "",
			IdZUzmsGL2:     "",
			ZLokkGL2:       "",
			IdZPrekVN:      "",
			IdZDisREZ:      "",
			IdZKvarREZ:     "",
			IdZPrstREZ:     "",
			IdZZmspREZ:     "",
			IdZNel1:        "44",
			IdZNel2:        "",
			IdZNel3:        "",
			IdZPrekNN:      "",
			IdZSabzSAB:     "",
			IdZOtprSAB:     "",
			IdSVremUSL:     "",
			UzrokTekst:     "",
			IdZJpsVN:       "",
			IdZJpsNN:       "",
			PoslTekst:      "",
			IdZTelePocGL1:  "",
			IdZTeleKrajGL1: "",
			IdZTelePocGL2:  "",
			IdZTeleKrajGL2: "",
			SynsoftId:      "1119",
		}
		return &pipiddn, nil
	}

	return nil, errors.New("pipiddn not found")
}

func (m *TestDBRepo) GetPiPiDDNByIDP(synsoftId string) (*models.PiPiDDN, error) {

	var pipiddn models.PiPiDDN
	if synsoftId == "1" {
		pipiddn = models.PiPiDDN{
			Datizv:         "27.04.2023",
			IdSMrc:         "8",
			IdSTipd:        "1",
			IdSVrpd:        "1",
			IdTipob:        "1",
			ObId:           "354",
			TrafoId:        "248",
			Vrepoc:         "27.04.2023 00:04:00",
			Vrezav:         "",
			Id1SGruzr:      "9",
			Id1SUzrok:      "70",
			Snaga:          "",
			Opis:           "this is outage put api",
			IdSNap:         "7",
			P2TrafId:       "3625",
			KorUneo:        "DTOMIC",
			IdZDsdfGL1:     "54",
			IdZKvarGL1:     "",
			IdZRapuGL1:     "",
			IdZPrstGL1:     "",
			IdZZmspGL1:     "",
			IdZUzmsGL1:     "",
			ZLokkGL1:       "",
			IdZDsdfGL2:     "54",
			IdZKvarGL2:     "",
			IdZRapuGL2:     "",
			IdZPrstGL2:     "",
			IdZZmspGL2:     "",
			IdZUzmsGL2:     "",
			ZLokkGL2:       "",
			IdZPrekVN:      "",
			IdZDisREZ:      "",
			IdZKvarREZ:     "",
			IdZPrstREZ:     "",
			IdZZmspREZ:     "",
			IdZNel1:        "44",
			IdZNel2:        "",
			IdZNel3:        "",
			IdZPrekNN:      "",
			IdZSabzSAB:     "",
			IdZOtprSAB:     "",
			IdSVremUSL:     "",
			UzrokTekst:     "",
			IdZJpsVN:       "",
			IdZJpsNN:       "",
			PoslTekst:      "",
			IdZTelePocGL1:  "",
			IdZTeleKrajGL1: "",
			IdZTelePocGL2:  "",
			IdZTeleKrajGL2: "",
			SynsoftId:      "1119",
		}
		return &pipiddn, nil
	}

	return nil, errors.New("pipiddn not found")
}

func (m *TestDBRepo) GetAllUnfinishedEventsNDC() ([]*models.UnfinishedEvents, error) {

	var p []*models.UnfinishedEvents

	return p, nil
}

func (m *TestDBRepo) GetAllUnfinishedEventsNDCP() ([]*models.UnfinishedEvents, error) {

	var p []*models.UnfinishedEvents

	return p, nil
}

func (m *TestDBRepo) GetUnfinishedEventsByID(synsoftId string) (*models.UnfinishedEvents, error) {

	var ue models.UnfinishedEvents

	return &ue, nil
}

func (m *TestDBRepo) GetUnfinishedEventsByIDP(synsoftId string) (*models.UnfinishedEvents, error) {

	var ue models.UnfinishedEvents

	return &ue, nil
}

func (m *TestDBRepo) UpdateUnfinishedEvents(ue models.UnfinishedEventsUpdate) error {

	return nil
}

func (m *TestDBRepo) UpdateUnfinishedEventsP(ue models.UnfinishedEventsUpdate) error {

	return nil
}

func (m *TestDBRepo) DeletePiPiDDN(synsoftId string) error {

	return nil
}

func (m *TestDBRepo) DeletePiPiDDNP(synsoftId string) error {

	return nil
}

func (m *TestDBRepo) InsertPiPiDDNIspad(pipiddn models.PiPiDDNIspad) error {

	return nil
}

func (m *TestDBRepo) InsertPiPiDDNIspadP(pipiddn models.PiPiDDNIspad) error {

	return nil
}

func (m *TestDBRepo) UpdatePiPiDDNIspad(pipiddn models.PiPiDDNIspad) error {
	return nil
}

func (m *TestDBRepo) UpdatePiPiDDNIspadP(pipiddn models.PiPiDDNIspad) error {
	return nil
}

func (m *TestDBRepo) GetAllPiPiDDNIspad() ([]*models.PiPiDDN, error) {

	var p []*models.PiPiDDN

	return p, nil
}

func (m *TestDBRepo) GetAllPiPiDDNIspadP() ([]*models.PiPiDDN, error) {

	var p []*models.PiPiDDN

	return p, nil
}

func (m *TestDBRepo) InsertDDNInterruptionOfDelivery(ddnintd models.DDNInterruptionOfDelivery) error {
	return nil
}

func (m *TestDBRepo) InsertDDNInterruptionOfDeliveryP(ddnintd models.DDNInterruptionOfDelivery) error {
	return nil
}

func (m *TestDBRepo) UpdateDDNInterruptionOfDelivery(ddnintd models.DDNInterruptionOfDelivery) error {
	return nil
}

func (m *TestDBRepo) UpdateDDNInterruptionOfDeliveryP(ddnintd models.DDNInterruptionOfDelivery) error {
	return nil
}

func (m *TestDBRepo) DeleteDDNInterruptionOfDelivery(synsoftId string) error {

	return nil
}

func (m *TestDBRepo) DeleteDDNInterruptionOfDeliveryP(synsoftId string) error {

	return nil
}

func (m *TestDBRepo) GetDDNInterruptionOfDeliveryNDC() ([]*models.DDNInterruptionOfDelivery, error) {

	var p []*models.DDNInterruptionOfDelivery

	return p, nil
}

func (m *TestDBRepo) GetDDNInterruptionOfDeliveryNDCP() ([]*models.DDNInterruptionOfDelivery, error) {

	var p []*models.DDNInterruptionOfDelivery

	return p, nil
}

func (m *TestDBRepo) GetDDNInterruptionOfDeliveryNDCByID(synsoftId string) (*models.DDNInterruptionOfDelivery, error) {

	var ue models.DDNInterruptionOfDelivery

	return &ue, nil
}

func (m *TestDBRepo) GetDDNInterruptionOfDeliveryNDCByIDP(synsoftId string) (*models.DDNInterruptionOfDelivery, error) {

	var ue models.DDNInterruptionOfDelivery

	return &ue, nil
}

/** start Check for PI_DOK **/
func (m *TestDBRepo) CheckForPiDokYesterdayP(datIzv string, idSMrc int) (int, error) {

	var num int

	return num, nil
}

func (m *TestDBRepo) CheckForPiDokTodayP(datIzv string, idSMrc int) (int, error) {

	var num int

	return num, nil
}

func (m *TestDBRepo) ClosePgiP(datIzv string, idSMrc string) error {

	

	return nil
}

func (m *TestDBRepo) TransferInPgiP(datIzv string, idSMrc string, Tip string) error {
	
		return nil
	
}
/** start Check for PI_DOK **/

/** start NOVITA ***/
func (m *TestDBRepo) GetAllUnbalancedTrader() ([]*models.UnbalancedTrader, error) {

	var p []*models.UnbalancedTrader

	return p, nil
}

/** end NOVITA ***/
