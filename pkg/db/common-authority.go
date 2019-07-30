package db

import (
	"database/sql"
	"fmt"
)

const commonAuthorityLoad = `
select  "commonAuthority"."commonAuthorityId"
,       "commonAuthority"."name"
,       "commonAuthority"."privateKeyData"
,       "commonAuthority"."certificateData"
,       "user"."username"
from    "CommonAuthority" "commonAuthority"
inner join "User" "user"
        on      "user"."userId" = "commonAuthority"."createdBy"
where	"commonAuthority"."commonAuthorityId" = $1
`

const commonAuthoritySave = `
insert into "CommonAuthority"
(
    "name"
,   "privateKeyData"
,   "certificateData"
,   "createdBy"
)
values
(
    $1
,   $2
,   $3
,   $4
)
returning "commonAuthorityId";
`

const commonAuthorityList = `
select  "commonAuthority"."commonAuthorityId"
,       "commonAuthority"."name"
,       "commonAuthority"."certificateData"
,       "user"."username"
from    "CommonAuthority" "commonAuthority"
inner join "User" "user"
        on      "user"."userId" = "commonAuthority"."createdBy"
`

const commonAuthorityDelete = `
delete
from	"CommonAuthority" "commonAuthority"
where	"commonAuthority"."commonAuthorityId" = $1;
`

type CommonAuthority struct {
	CommonAuthorityData
	CommonAuthorityId int
}

type CommonAuthorityData struct {
	Name            string
	PrivateKeyData  []byte
	CertificateData []byte
}

func (data CommonAuthorityData) GetPrivateKeyData() []byte {
	return data.PrivateKeyData
}

func (data CommonAuthorityData) GetCertificateData() []byte {
	return data.CertificateData
}

type CommonAuthorityInfo struct {
	CommonAuthorityData
	CommonAuthorityId int
	CreatedBy         UserInfo
}

type UserInfo struct {
	Username  string
	FirstName string
	LastName  string
}

type CommonAuthorityDAO struct {
	database *Database
}

func NewCommonAuthorityDAO(database *Database) *CommonAuthorityDAO {
	return &CommonAuthorityDAO{database}
}

func (dao *CommonAuthorityDAO) SaveCommonAuthority(name string, privateKeyData, certificateData []byte, userId int) (int, error) {
	var commonAuthorityId int
	if err := dao.database.db.QueryRow(commonAuthoritySave, name, privateKeyData, certificateData, userId).Scan(&commonAuthorityId); err != nil {
		return 0, err
	}
	return commonAuthorityId, nil
}

func (dao *CommonAuthorityDAO) LoadCommonAuthority(commonAuthorityId int) (CommonAuthorityInfo, error) {
	row := dao.database.db.QueryRow(commonAuthorityLoad, commonAuthorityId)
	var commonAuthority CommonAuthorityInfo
	if err := row.Scan(
		&commonAuthority.CommonAuthorityId,
		&commonAuthority.Name,
		&commonAuthority.PrivateKeyData,
		&commonAuthority.CertificateData,
		&commonAuthority.CreatedBy.Username); err != nil {
		return CommonAuthorityInfo{}, err
	}

	// TODO: Load these
	commonAuthority.CreatedBy.FirstName = "Anton"
	commonAuthority.CreatedBy.LastName = "Johansson"
	return commonAuthority, nil
}

func (dao *CommonAuthorityDAO) ListCommonAuthorities() ([]CommonAuthorityInfo, error) {
	rows, err := dao.database.db.Query(commonAuthorityList)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var commonAuthorities []CommonAuthorityInfo
	for rows.Next() {
		var commonAuthority CommonAuthorityInfo
		if err := rows.Scan(
			&commonAuthority.CommonAuthorityId,
			&commonAuthority.Name,
			&commonAuthority.CertificateData,
			&commonAuthority.CreatedBy.Username); err != nil {
			fmt.Println("error scanning type in list:", err)
			continue
		}
		// TODO: Load these
		commonAuthority.CreatedBy.FirstName = "Anton"
		commonAuthority.CreatedBy.LastName = "Johansson"
		commonAuthorities = append(commonAuthorities, commonAuthority)
	}
	return commonAuthorities, nil
}

func (dao *CommonAuthorityDAO) DeleteCommonAuthority(commonAuthorityId int) error {
	result, err := dao.database.db.Exec(commonAuthorityDelete, commonAuthorityId)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}
