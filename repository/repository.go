package repository

import (
	"database/sql"

	"github.com/tijanadmi/tis-api/models"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	OneSignal(id int) (*models.Signal, error)
	GetDvDidf() ([]*models.Signal, error)
	GetDiffTr() ([]*models.Signal, error)
	GetDisTrRes() ([]*models.Signal, error)
	GetDisDiffSp() ([]*models.Signal, error)
	GetMalfunctionIn() ([]*models.MalfunctionIn, error)
	OneMalfunctionIn(id int) (*models.MalfunctionIn, error)
	GetAPU() ([]*models.Apu, error)
	OneAPU(id int) (*models.Apu, error)
	GetOCDV() ([]*models.Signal, error)
	GetOCTR12() ([]*models.Signal, error)
	GetOCTRR() ([]*models.Signal, error)
	GetOCSP() ([]*models.Signal, error)
	GetEarthfaultOCDV() ([]*models.Signal, error)
	GetEarthfaultOCTR() ([]*models.Signal, error)
	GetEarthfaultOCSP() ([]*models.Signal, error)
	GetDirEarthfaultOC() ([]*models.Signal, error)
	GetTPSendRcdv() ([]*models.Signal, error)
	GetCircuitbreaker() ([]*models.Signal, error)
	GetBBPBFtrip() ([]*models.Signal, error)
	GetNonElectrical() ([]*models.Signal, error)
	GetBBPBBtrip() ([]*models.Signal, error)
	GetBFtrip() ([]*models.Signal, error)
	GetGroupsOfCauses() ([]*models.GroupOfCause, error)
	OneGroupOfCauses(id int) (*models.GroupOfCause, error)
	GetCauses() ([]*models.Cause, error)
	OneCause(id int) (*models.Cause, error)
	GetGroupOfReasons() ([]*models.GroupOfReason, error)
	OneGroupOfReasons(id int) (*models.GroupOfReason, error)
	GetReasons() ([]*models.Reason, error)
	OneReason(id int) (*models.Reason, error)
	GetWeatherConditions() ([]*models.WeatherCondition, error)
	OneWeatherCondition(id int) (*models.WeatherCondition, error)
	GetCategoriesOfEvents() ([]*models.CategoryOfEvent, error)
	OneCategoryOfEvents(id int) (*models.CategoryOfEvent, error)
	GetOHL() ([]*models.OverheadLine, error)
	GetPowerCables() ([]*models.PowerCable, error)
	GetSubstations() ([]*models.Substation, error)
	GetFeeders() ([]*models.Feeder, error)
	GetProtectionDevices() ([]*models.ProtectionDevice, error)
	GetPowerTransformers() ([]*models.PowerTransformer, error)
	GetDisconnectors() ([]*models.Disconnector, error)
	GetWorkPermissions() ([]*models.WorkPermission, error)
	GetWorkPermissionsAll() ([]*models.WorkPermissionAll, error)
	GetRequest1Gr() ([]*models.Request1Gr, error)
	GetRequest2Gr() ([]*models.Request2Gr, error)
	GetWorkInEENetwork() ([]*models.WorkInEENetwork, error)
	GetWeather(year string) ([]*models.WeatherData, error)
	GetWeatherForecast() ([]*models.WeatherData, error)
	GetWeatherHistory(year string) ([]*models.WeatherDataHistory, error)
	GetPermissions1(year string) ([]*models.Permission, error)
	GetPermissions23(year string) ([]*models.Permission, error)
	GetRequests1(year string) ([]*models.Request, error)
	GetRequests23(year string) ([]*models.Request, error)
	GetOutages(year string) ([]*models.Outage, error)
	GetExclusions(year string) ([]*models.Exclusion, error)
	GetPlans(year string) ([]*models.Plan, error)
	GetUnopenedPermitForDay(day string, org string) ([]*models.UnopenedPermit, error)
	Authenticate(username, testPassword string) error
	GetUserByUsername(username string) (*models.User, error)
	GetUserByID(id int) (*models.User, error)
	InsertPiPiDDNIsklj(pipiddn models.PiPiDDNIsklj) error
	InsertPiPiDDNIskljP(pipiddn models.PiPiDDNIsklj) error
	UpdatePiPiDDNIsklj(pipiddn models.PiPiDDNIsklj) error
	UpdatePiPiDDNIskljP(pipiddn models.PiPiDDNIsklj) error
	GetAllPiPiDDN() ([]*models.PiPiDDN, error)
	GetAllPiPiDDNP() ([]*models.PiPiDDN, error)
	GetPiPiDDNByID(synsoftId string) (*models.PiPiDDN, error)
	GetPiPiDDNByIDP(synsoftId string) (*models.PiPiDDN, error)
	GetAllUnfinishedEventsNDC() ([]*models.UnfinishedEvents, error)
	GetAllUnfinishedEventsNDCP() ([]*models.UnfinishedEvents, error)
	GetUnfinishedEventsByID(synsoftId string) (*models.UnfinishedEvents, error)
	GetUnfinishedEventsByIDP(synsoftId string) (*models.UnfinishedEvents, error)
	UpdateUnfinishedEvents(ue models.UnfinishedEventsUpdate) error
	UpdateUnfinishedEventsP(ue models.UnfinishedEventsUpdate) error
	DeletePiPiDDN(synsoftId string) error
	DeletePiPiDDNP(synsoftId string) error
	InsertPiPiDDNIspad(pipiddn models.PiPiDDNIspad) error
	InsertPiPiDDNIspadP(pipiddn models.PiPiDDNIspad) error
	UpdatePiPiDDNIspad(pipiddn models.PiPiDDNIspad) error
	UpdatePiPiDDNIspadP(pipiddn models.PiPiDDNIspad) error
	GetAllPiPiDDNIspad() ([]*models.PiPiDDN, error)
	GetAllPiPiDDNIspadP() ([]*models.PiPiDDN, error)
	InsertDDNInterruptionOfDelivery(ddnintd models.DDNInterruptionOfDelivery) error
	InsertDDNInterruptionOfDeliveryP(ddnintd models.DDNInterruptionOfDelivery) error
	UpdateDDNInterruptionOfDelivery(ddnintd models.DDNInterruptionOfDelivery) error
	UpdateDDNInterruptionOfDeliveryP(ddnintd models.DDNInterruptionOfDelivery) error
	DeleteDDNInterruptionOfDelivery(synsoftId string) error
	DeleteDDNInterruptionOfDeliveryP(synsoftId string) error
	GetDDNInterruptionOfDeliveryNDC() ([]*models.DDNInterruptionOfDelivery, error)
	GetDDNInterruptionOfDeliveryNDCP() ([]*models.DDNInterruptionOfDelivery, error)
	GetDDNInterruptionOfDeliveryNDCByID(synsoftId string) (*models.DDNInterruptionOfDelivery, error)
	GetDDNInterruptionOfDeliveryNDCByIDP(synsoftId string) (*models.DDNInterruptionOfDelivery, error)
	CheckForPiDokYesterdayP(datIzv string, idSMrc int) (int,error)
	CheckForPiDokTodayP(datIzv string, idSMrc int) (int,error)
	GetAllUnbalancedTrader() ([]*models.UnbalancedTrader, error)
}
