package dbrepo

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"
	"time"

	_ "github.com/godror/godror"
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

/*func TestOracleDBRepoGetDvDidf(t *testing.T) {
	signals, err := testRepo.GetDvDidf()
	if err != nil {
		t.Errorf("all users reports an error: %s", err)
	}

	if len(signals) != 1 {
		t.Errorf("all users reports wrong size; expected 1, but got %d", len(users))
	}

	testUser := data.User{
		FirstName: "Jack",
		LastName:  "Smith",
		Email:     "jack@smith.com",
		Password:  "secret",
		IsAdmin:   1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, _ = testRepo.InsertUser(testUser)

	users, err = testRepo.AllUsers()
	if err != nil {
		t.Errorf("all users reports an error: %s", err)
	}

	if len(users) != 2 {
		t.Errorf("all users reports wrong size after insert; expected 2, but got %d", len(users))
	}
}*/

/*func TestOracleDBRepoInsertPiPiDDNIsklj(t *testing.T) {

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
}*/
