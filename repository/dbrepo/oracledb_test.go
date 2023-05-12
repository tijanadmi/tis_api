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

func TestOracleDBRepoInsertUser(t *testing.T) {

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
