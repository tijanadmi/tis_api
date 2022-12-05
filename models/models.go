package models

import (
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// Models is the wrapper for database
type Models struct {
	DB DBModel
}

// NewModels returns models with db pool
func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{DB: db},
	}
}

// Zastita is the type for all signals of protective devices
type Signal struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type MalfunctionIn struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Order string `json:"order"`
}

type Apu struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
	Code      string `json:"code"`
	Status    string `json:"status"`
}

type WeatherConditions struct {
	ID     int    `json:"id"`
	Code   string `json:"code"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type OverheadLine struct {
	ID                      int    `json:"tis_id_ohl"`
	IpsIdDV                 string `json:"ips_id_ohl"`
	SapIdDV                 string `json:"sap_id_ohl"`
	IdDal                   string `json:"code_ohl"`
	IPSName                 string `json:"ips_name_ohl"`
	CategoryNameDV          string `json:"category_name_ohl"`
	IdNN                    int    `json:"id_voltage_level"`
	NNName                  string `json:"name_voltage_level"`
	CategoryId              string `json:"category_id"`
	CategoryName            string `json:"category_name"`
	TisIdFirstFeeder        int    `json:"tis_id_first_feeder"`
	IpsIdFirstFeeder        string `json:"ips_id_first_feeder"`
	SapIdFirstFeeder        string `json:"sap_id_first_feeder"`
	FirstFeederNumber       string `json:"first_feeder_number"`
	FirstFeederName         string `json:"first_feeder_name"`
	FirstFeederCategoryName string `json:"first_feeder_category_name"`
	TisIdLastFeeder         int    `json:"tis_id_last_feeder"`
	IpsIdLastFeeder         string `json:"ips_id_last_feeder"`
	SapIdLastFeeder         string `json:"sap_id_last_feeder"`
	LastFeederNumber        string `json:"last_feeder_number"`
	LastFeederName          string `json:"last_feeder_name"`
	LastFeederCategoryName  string `json:"last_feeder_category_name"`
	ObjTypeId               int    `json:"object_type_id"`
	ObjTypeCode             string `json:"object_type_code"`
	ObjTypeName             string `json:"object_type_name"`
	Mrc1Id                  string `json:"rdc1_id"`
	Mrc1Name                string `json:"rdc1_name"`
	Mrc2Id                  string `json:"rdc2_id"`
	Mrc2Name                string `json:"rdc2_name"`
}

type PowerCable struct {
	ID                      int    `json:"tis_id_cabel"`
	IpsIdKB                 string `json:"ips_id_cabel"`
	SapIdKB                 string `json:"sap_id_cabel"`
	IdKab                   string `json:"code_cabel"`
	IPSName                 string `json:"ips_name_cabel"`
	CategoryNameDV          string `json:"category_name_cabel"`
	CategoryId              string `json:"category_id"`
	CategoryName            string `json:"category_name"`
	TisIdFirstFeeder        int    `json:"tis_id_first_feeder"`
	IpsIdFirstFeeder        string `json:"ips_id_first_feeder"`
	SapIdFirstFeeder        string `json:"sap_id_first_feeder"`
	FirstFeederNumber       string `json:"first_feeder_number"`
	FirstFeederName         string `json:"first_feeder_name"`
	FirstFeederCategoryName string `json:"first_feeder_category_name"`
	TisIdLastFeeder         int    `json:"tis_id_last_feeder"`
	IpsIdLastFeeder         string `json:"ips_id_last_feeder"`
	SapIdLastFeeder         string `json:"sap_id_last_feeder"`
	LastFeederNumber        string `json:"last_feeder_number"`
	LastFeederName          string `json:"last_feeder_name"`
	LastFeederCategoryName  string `json:"last_feeder_category_name"`
	ObjTypeId               int    `json:"object_type_id"`
	ObjTypeCode             string `json:"object_type_code"`
	ObjTypeName             string `json:"object_type_name"`
	Mrc1Id                  string `json:"rdc1_id"`
	Mrc1Name                string `json:"rdc1_name"`
	Mrc2Id                  string `json:"rdc2_id"`
	Mrc2Name                string `json:"rdc2_name"`
}

type Substation struct {
	ID              int    `json:"tis_id"`
	IpsId           string `json:"ips_id"`
	SapId           string `json:"sap_id"`
	TISName         string `json:"name_sub"`
	CategoryNameSub string `json:"category_name_sub"`
	ObjTypeId       int    `json:"object_type_id"`
	ObjTypeCode     string `json:"object_type_code"`
	ObjTypeName     string `json:"object_type_name"`
	Mrc1Id          int    `json:"rdc1_id"`
	Mrc1Name        string `json:"rdc1_name"`
	Mrc2Id          string `json:"rdc2_id"`
	Mrc2Name        string `json:"rdc2_name"`
	OrgId           int    `json:"org_id"`
	OrgName         string `json:"org_name"`
}

type Feeder struct {
	ID                 int    `json:"tis_id_feeder"`
	IpsIdFeeder        string `json:"ips_id_feeder"`
	SapIdFeeder        string `json:"sap_id_feeder"`
	TisIdSub           int    `json:"tis_id_substation"`
	IpsIdSub           string `json:"ips_id_substation"`
	SapIdSub           string `json:"sap_id_substation"`
	NameSub            string `json:"name_substation"`
	IdNN               int    `json:"id_voltage_level"`
	NNName             string `json:"name_voltage_level"`
	FeederNumber       string `json:"feeder_number"`
	FeederName         string `json:"feeder_name"`
	FeederCategoryName string `json:"feeder_category_name"`
	CategoryId         string `json:"category_id"`
	CategoryName       string `json:"category_name"`
	FeederFunId        string `json:"feeder_fun_id"`
	FeederFunName      string `json:"feeder_fun_name"`
	Equipped           string `json:"equipped"`
	Active             string `json:"active"`
	Completely         string `json:"completely"`
	FeederType         string `json:"feeder_type"`
}

type ProtectionDevice struct {
	ID           int    `json:"tis_id_prd"`
	IpsIdDV      string `json:"ips_id_prd"`
	TisIdFeeder  int    `json:"tis_id_feeder"`
	IpsIdFeeder  string `json:"ips_id_feeder"`
	SapIdFFeeder string `json:"sap_id_feeder"`
	Manufacturer string `json:"manufacturer"`
	Type         string `json:"type"`
	TypeOfEqu    int    `json:"type_of_equ"`
	Indicator    string `json:"indicator"`
	Technology   string `json:"tehnology"`
	StatusIPS    string `json:"status_ips"`
	Status       string `json:"status"`
	SerialNumber string `json:"serial_number"`
	DeviceTag    string `json:"device_tag"`
	APU          string `json:"apu"`
}

// User is the type for users
type User struct {
	ID       int
	Username string
	Password string
}

func (u *User) PasswordMatches(plainText string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainText))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			// invalid password
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}
