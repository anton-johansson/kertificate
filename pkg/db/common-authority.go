// Copyright 2019 Anton Johansson
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package db

import (
	"database/sql"
	"fmt"
	"time"
)

const commonAuthorityLoad = `
select  "commonAuthority"."commonAuthorityId"
,       "commonAuthority"."name"
,       "data"."privateKeyData"
,       "data"."certificateData"
,       "user"."username"
from    "CommonAuthority" "commonAuthority"
inner join "CertificateData" "data"
        on      "data"."certificateDataId" = "commonAuthority"."certificateDataId"
inner join "User" "user"
        on      "user"."userId" = "commonAuthority"."createdBy"
where	"commonAuthority"."commonAuthorityId" = $1
`

const commonAuthoritySave = `
with "data" as
(
    insert into "CertificateData"
    (
            "privateKeyData"
    ,       "certificateData"
    ,       "expiresAt"
    )
    values
    (
            $1
	,       $2
    ,       $3
    )
	returning "certificateDataId"
)
insert into "CommonAuthority"
(
        "certificateDataId"
,       "name"
,       "createdBy"
)
select  "certificateDataId"
,       $4
,       $5
from    "data"
returning "commonAuthorityId";
`

const commonAuthorityList = `
select  "commonAuthority"."commonAuthorityId"
,       "commonAuthority"."name"
,       "data"."certificateData"
,       "user"."username"
from    "CommonAuthority" "commonAuthority"
inner join "CertificateData" "data"
        on      "data"."certificateDataId" = "commonAuthority"."certificateDataId"
inner join "User" "user"
        on      "user"."userId" = "commonAuthority"."createdBy"
`

const commonAuthorityDelete = `
with "deleted" as
(
        delete
        from	"CommonAuthority" "commonAuthority"
		where	"commonAuthority"."commonAuthorityId" = $1
		returning "certificateDataId"
)
delete
from    "CertificateData" "data"
where   "data"."certificateDataId" =
(
    select  "certificateDataId"
    from    "deleted"
);`

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

func (dao *CommonAuthorityDAO) SaveCommonAuthority(name string, privateKeyData, certificateData []byte, expiresAt time.Time, userId int) (int, error) {
	var commonAuthorityId int
	if err := dao.database.db.QueryRow(
		commonAuthoritySave,
		privateKeyData,
		certificateData,
		expiresAt,
		name,
		userId).Scan(&commonAuthorityId); err != nil {
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
