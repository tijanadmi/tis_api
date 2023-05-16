//go:build integration
package dbrepo

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"
	"time"

	_ "github.com/godror/godror"
	"github.com/tijanadmi/tis-api/models"
	"github.com/tijanadmi/tis-api/repository"
)

var dns = ""
var testDB *sql.DB
var testRepo repository.DatabaseRepo

func TestMain(m *testing.M) {
	testDB, err := sql.Open("godror", dns)
	if err != nil {
		log.Panicln(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = testDB.PingContext(ctx)
	if err != nil {
		log.Panicln(err)
	}

	testRepo = &OracleDBRepo{DB: testDB}

	// run tests
	code := m.Run()

	os.Exit(code)

}

func TestOracleDBRepoOneSignal(t *testing.T) {
	signal, err := testRepo.OneSignal(262)
	if err != nil {
		t.Errorf("error getting user by id: %s", err)
	}

	if signal.Name != "Sistem za hlađenje - kvar" {
		t.Errorf("wrong email returned by OneSignal; expected Sistem za hlađenje - kvar but got %s", signal.Name)
	}

	_, err = testRepo.OneSignal(500)
	if err == nil {
		t.Error("no error reported when getting non existent user by id")
	}

}

func TestOracleDBRepoOneMalfunctionIn(t *testing.T) {
	mf, err := testRepo.OneMalfunctionIn(18)
	if err != nil {
		t.Errorf("error getting user by id: %s", err)
	}

	if mf.Name != "Faza “0”" {
		t.Errorf("wrong email returned by OneSignal; expected Sistem za hlađenje - kvar but got %s", mf.Name)
	}

	_, err = testRepo.OneMalfunctionIn(1)
	if err == nil {
		t.Error("no error reported when getting non existent user by id")
	}
}

func TestOracleDBRepoOneAPU(t *testing.T) {
	mf, err := testRepo.OneAPU(1)
	if err != nil {
		t.Errorf("error getting user by id: %s", err)
	}

	if mf.Name != "Uspešno" {
		t.Errorf("wrong email returned by OneSignal; expected Sistem za hlađenje - kvar but got %s", mf.Name)
	}

	_, err = testRepo.OneAPU(201)
	if err == nil {
		t.Error("no error reported when getting non existent user by id")
	}
}

func TestOracleDBRepoOneGroupOfCauses(t *testing.T) {
	grc, err := testRepo.OneGroupOfCauses(1)
	if err != nil {
		t.Errorf("error getting user by id: %s", err)
	}

	if grc.Name != "DV - elektromontažni deo" {
		t.Errorf("wrong email returned by OneSignal; expected Sistem za hlađenje - kvar but got %s", grc.Name)
	}

	_, err = testRepo.OneGroupOfCauses(20)
	if err == nil {
		t.Error("no error reported when getting non existent user by id")
	}
}

func TestOracleDBRepoOneCause(t *testing.T) {
	cau, err := testRepo.OneCause(1)
	if err != nil {
		t.Errorf("error getting user by id: %s", err)
	}

	if cau.Name != "Avionska nezgoda" {
		t.Errorf("wrong email returned by OneSignal; expected Sistem za hlađenje - kvar but got %s", cau.Name)
	}

	_, err = testRepo.OneCause(95)
	if err == nil {
		t.Error("no error reported when getting non existent user by id")
	}
}

func TestOracleDBRepoOneGroupOfReasons(t *testing.T) {
	gr, err := testRepo.OneGroupOfReasons(1)
	if err != nil {
		t.Errorf("error getting user by id: %s", err)
	}

	if gr.Name != "Preventivno održavanje" {
		t.Errorf("wrong email returned by OneSignal; expected Sistem za hlađenje - kvar but got %s", gr.Name)
	}

	_, err = testRepo.OneGroupOfReasons(15)
	if err == nil {
		t.Error("no error reported when getting non existent user by id")
	}
}

func TestOracleDBRepoOneReason(t *testing.T) {
	r, err := testRepo.OneReason(1)
	if err != nil {
		t.Errorf("error getting user by id: %s", err)
	}

	if r.Name != "Pregled" {
		t.Errorf("wrong email returned by OneSignal; expected Sistem za hlađenje - kvar but got %s", r.Name)
	}

	_, err = testRepo.OneReason(25)
	if err == nil {
		t.Error("no error reported when getting non existent user by id")
	}
}
func TestOracleDBRepoOneWeatherCondition(t *testing.T) {
	w, err := testRepo.OneWeatherCondition(1)
	if err != nil {
		t.Errorf("error getting user by id: %s", err)
	}

	if w.Name != "Suvo vreme" {
		t.Errorf("wrong email returned by OneSignal; expected Sistem za hlađenje - kvar but got %s", w.Name)
	}

	_, err = testRepo.OneWeatherCondition(15)
	if err == nil {
		t.Error("no error reported when getting non existent user by id")
	}
}

func TestOracleDBRepoOneCategoryOfEvents(t *testing.T) {
	ce, err := testRepo.OneCategoryOfEvents(1)
	if err != nil {
		t.Errorf("error getting user by id: %s", err)
	}

	if ce.Name != "APU Uspešno" {
		t.Errorf("wrong email returned by OneSignal; expected Sistem za hlađenje - kvar but got %s", ce.Name)
	}

	_, err = testRepo.OneCategoryOfEvents(95)
	if err == nil {
		t.Error("no error reported when getting non existent user by id")
	}
}

func TestOracleDBRepoGetUserByUsername(t *testing.T) {
	user, err := testRepo.GetUserByUsername("tijana")
	if err != nil {
		t.Errorf("error getting user by username: %s", err)
	}

	if user.ID != 1 {
		t.Errorf("wrong id returned by GetUserByUsername; expected 1 but got %d", user.ID)
	}

	_, err = testRepo.GetUserByUsername("ana")
	if err == nil {
		t.Error("no error reported when getting non existent user by username")
	}
}
func TestOracleDBRepoGetUserByID(t *testing.T) {
	user, err := testRepo.GetUserByID(1)
	if err != nil {
		t.Errorf("error getting user by id: %s", err)
	}

	if user.Username != "tijana" {
		t.Errorf("wrong email returned by OneSignal; expected tijana but got %s", user.Username)
	}

	_, err = testRepo.GetUserByID(5)
	if err == nil {
		t.Error("no error reported when getting non existent user by id")
	}
}

func TestOracleDBRepoInsertPiPiDDNIsklj(t *testing.T) {

	pipiddn := models.PiPiDDNIsklj{
		DatSmene:  "07.04.2023",
		IdSMrc:    "8",
		TipMan:    "1",
		IdTipob:   "4",
		ObId:      "105",
		TrafoId:   "",
		Vrepoc:    "08.04.2023 19:35:00",
		Vrezav:    "08.04.2023 23:25:00",
		IdSGrraz:  "3",
		IdSRazlog: "21",
		ManTekst:  "Od 23.03.23 07:45 na snazi planirani radovi broj: R-67/1 (AKZ metalnih nosača VN opreme u TRP 400kV od TR1 u  TS Beograd 8 Vasiljević Aleksandar). NA SNAZI;",
		IdSNap:    "",
		P2TrafId:  "",
		KorUneo:   "pi1100",
		SynsoftId: "12399",
	}

	err := testRepo.InsertPiPiDDNIsklj(pipiddn)
	if err != nil {
		t.Errorf("insert user returned an error: %s", err)
	}
}

func TestOracleDBRepoGetPiPiDDNByID(t *testing.T) {
	pipiddn, err := testRepo.GetPiPiDDNByID("12399")
	if err != nil {
		t.Errorf("error getting user by id: %s", err)
	}

	if pipiddn.IdSMrc != "8" {
		t.Errorf("wrong  returned by GetPiPiDDNByID; expected 8 but got %s", pipiddn.IdSMrc)
	}

	_, err = testRepo.GetPiPiDDNByID("12398")
	if err == nil {
		t.Error("no error reported when getting non existent user by id")
	}
}

func TestOracleDBRepoUpdatePiPiDDNIsklj(t *testing.T) {

	pipiddn := models.PiPiDDNIsklj{
		DatSmene:  "08.04.2023",
		IdSMrc:    "8",
		TipMan:    "1",
		IdTipob:   "4",
		ObId:      "105",
		TrafoId:   "",
		Vrepoc:    "08.04.2023 19:35:00",
		Vrezav:    "08.04.2023 23:25:00",
		IdSGrraz:  "3",
		IdSRazlog: "21",
		ManTekst:  "Od 23.03.23 07:45 na snazi planirani radovi broj: R-67/1 (AKZ metalnih nosača VN opreme u TRP 400kV od TR1 u  TS Beograd 8 Vasiljević Aleksandar). NA SNAZI;",
		IdSNap:    "",
		P2TrafId:  "",
		KorUneo:   "pi1100",
		SynsoftId: "12399",
	}

	err := testRepo.UpdatePiPiDDNIsklj(pipiddn)
	if err != nil {
		t.Errorf("error updating PiPiDDNIsklj  with SynsoftId=12399 returned an error: %s", err)
	}

	pipiddn_up, _ := testRepo.GetPiPiDDNByID(pipiddn.SynsoftId)
	if pipiddn_up.Vrezav != "08.04.2023 23:25:00" {
		t.Errorf("expected updated record to have vrezav 08.04.2023 23:25:00, but got %s", pipiddn_up.Vrezav)
	}
}

func TestOracleDBRepoDeletePiPiDDN(t *testing.T) {
	err := testRepo.DeletePiPiDDN("12399")
	if err != nil {
		t.Errorf("error deleting user id 12399: %s", err)
	}

	_, err = testRepo.GetPiPiDDNByID("12399")
	if err == nil {
		t.Error("retrieved user id 12399, who should have been deleted")
	}
}

func TestOracleDBRepoInsertPiPiDDNIspad(t *testing.T) {

	pipiddn := models.PiPiDDNIspad{
		DatSmene:       "27.04.2023",
		IdSMrc:         "8",
		IdSTipd:        "1",
		IdSVrpd:        "1",
		IdRadAPU:       "1",
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

	err := testRepo.InsertPiPiDDNIspad(pipiddn)
	if err != nil {
		t.Errorf("insert user returned an error: %s", err)
	}
}

func TestOracleDBRepoUpdatePiPiDDNIspad(t *testing.T) {

	pipiddn := models.PiPiDDNIspad{
		DatSmene:       "27.04.2023",
		IdSMrc:         "8",
		IdSTipd:        "1",
		IdSVrpd:        "1",
		IdRadAPU:       "1",
		IdTipob:        "1",
		ObId:           "354",
		TrafoId:        "248",
		Vrepoc:         "27.04.2023 21:04:00",
		Vrezav:         "28.04.2023 00:04:00",
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

	err := testRepo.UpdatePiPiDDNIspad(pipiddn)
	if err != nil {
		t.Errorf("error updating PiPiDDNIsklj  with SynsoftId=12399 returned an error: %s", err)
	}

	pipiddn_up, _ := testRepo.GetPiPiDDNByID(pipiddn.SynsoftId)
	if pipiddn_up.Vrezav != "28.04.2023 00:04:00" {
		t.Errorf("expected updated record to have vrezav 28.04.2023 00:04:00, but got %s", pipiddn_up.Vrezav)
	}
}

func TestOracleDBRepoDeletePiPiDDN1(t *testing.T) {
	err := testRepo.DeletePiPiDDN("1119")
	if err != nil {
		t.Errorf("error deleting user id 1119: %s", err)
	}

	_, err = testRepo.GetPiPiDDNByID("1119")
	if err == nil {
		t.Error("retrieved user id 1119, who should have been deleted")
	}
}

func TestOracleDBRepoGetUnfinishedEventsByID(t *testing.T) {
	pipiddn, err := testRepo.GetUnfinishedEventsByID("123")
	if err != nil {
		t.Errorf("error getting user by id: %s", err)
	}

	if pipiddn.IdSMrc != "8" {
		t.Errorf("wrong  returned by GetUnfinishedEventsByID; expected 8 but got %s", pipiddn.IdSMrc)
	}

	_, err = testRepo.GetUnfinishedEventsByID("124")
	if err == nil {
		t.Error("no error reported when getting non existent user by id")
	}
}

func TestOracleDBRepoUpdateUnfinishedEvents(t *testing.T) {

	pipiddn := models.UnfinishedEventsUpdate{
		DatSmene:       "26.05.2023",
		TipSmene:       "D",
		Vrezav:         "16.05.2023 00:04:00",
		Id1SGruzr:      "9",
		Id1SUzrok:      "70",
		Opis:           "this is outage put api",
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
		IdZJpsVN:       "",
		IdZJpsNN:       "",
		IdZTelePocGL1:  "",
		IdZTeleKrajGL1: "",
		IdZTelePocGL2:  "",
		IdZTeleKrajGL2: "",
		SynsoftId:      "123",
	}

	err := testRepo.UpdateUnfinishedEvents(pipiddn)
	if err != nil {
		t.Errorf("error updating PiPiDDNIsklj  with SynsoftId=12399 returned an error: %s", err)
	}

	pipiddn_up, _ := testRepo.GetUnfinishedEventsByID(pipiddn.SynsoftId)
	if pipiddn_up.Vrezav != "16.05.2023 00:04:00" {
		t.Errorf("expected updated record to have vrezav 16.05.2023 00:04:00, but got %s", pipiddn_up.Vrezav)
	}
}

func TestOracleDBRepoInsertDDNInterruptionOfDelivery(t *testing.T) {

	pipiddn := models.DDNInterruptionOfDelivery{
		IdSMrc:            "8",
		IdSTipd:           "12",
		IdSVrpd:           "",
		IdTipob:           "6",
		ObId:              "149",
		Vrepoc:            "26.03.2023 12:46",
		Vrezav:            "",
		IdSVrPrek:         "2",
		IdSUzrokPrek:      "2",
		Snaga:             "",
		Opis:              "Problem sa ležajevima podbudilice.",
		KorUneo:           "NBRALUSIC",
		IdSMernaMesta:     "",
		BrojMesta:         "",
		Ind:               "P",
		P2TrafId:          "5391",
		Bi:                "",
		IdSPoduzrokPrek:   "",
		IdDogPrekidP:      "",
		IdTipObjektaNdc:   "",
		IdTipDogadjajaNdc: "",
		SynsoftId:         "12361",
	}

	err := testRepo.InsertDDNInterruptionOfDelivery(pipiddn)
	if err != nil {
		t.Errorf("insert InterruptionOfDelivery returned an error: %s", err)
	}
}

func TestOracleDBRepoGetDDNInterruptionOfDeliveryNDCByID(t *testing.T) {
	pipiddn, err := testRepo.GetDDNInterruptionOfDeliveryNDCByID("12361")
	if err != nil {
		t.Errorf("error getting user by id: %s", err)
	}

	if pipiddn.IdSMrc != "8" {
		t.Errorf("wrong  returned by GetUnfinishedEventsByID; expected 8 but got %s", pipiddn.IdSMrc)
	}

	_, err = testRepo.GetDDNInterruptionOfDeliveryNDCByID("124")
	if err == nil {
		t.Error("no error reported when getting non existent user by id")
	}
}

func TestOracleDBRepoUpdateDDNInterruptionOfDelivery(t *testing.T) {

	pipiddn := models.DDNInterruptionOfDelivery{
		IdSMrc:            "8",
		IdSTipd:           "12",
		IdSVrpd:           "",
		IdTipob:           "6",
		ObId:              "149",
		Vrepoc:            "26.03.2023 12:46",
		Vrezav:            "27.03.2023 12:46",
		IdSVrPrek:         "2",
		IdSUzrokPrek:      "2",
		Snaga:             "",
		Opis:              "Problem sa ležajevima podbudilice.",
		KorUneo:           "NBRALUSIC",
		IdSMernaMesta:     "",
		BrojMesta:         "",
		Ind:               "P",
		P2TrafId:          "5391",
		Bi:                "",
		IdSPoduzrokPrek:   "",
		IdDogPrekidP:      "",
		IdTipObjektaNdc:   "",
		IdTipDogadjajaNdc: "",
		SynsoftId:         "12361",
	}

	err := testRepo.UpdateDDNInterruptionOfDelivery(pipiddn)
	if err != nil {
		t.Errorf("error updating InterruptionOfDelivery  with SynsoftId=12399 returned an error: %s", err)
	}

	pipiddn_up, _ := testRepo.GetDDNInterruptionOfDeliveryNDCByID(pipiddn.SynsoftId)
	if pipiddn_up.Vrezav != "27.03.2023 12:46:00" {
		t.Errorf("expected updated record to have vrezav 28.04.2023 00:04:00, but got %s", pipiddn_up.Vrezav)
	}
}

func TestOracleDBRepoDeleteDDNInterruptionOfDelivery(t *testing.T) {
	err := testRepo.DeleteDDNInterruptionOfDelivery("12361")
	if err != nil {
		t.Errorf("error deleting user id 12361: %s", err)
	}

	_, err = testRepo.GetDDNInterruptionOfDeliveryNDCByID("12361")
	if err == nil {
		t.Error("retrieved user id 12361, who should have been deleted")
	}
}
