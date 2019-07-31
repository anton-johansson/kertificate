package db

import (
	"database/sql"
	"time"
)

const certificateSave = `
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
insert into "Certificate"
(
        "certificateDataId"
,       "commonAuthorityId"
,       "name"
,       "createdBy"
)
select  "certificateDataId"
,       $4
,       $5
,       $6
from    "data"
returning "certificateId";`

const certificateDelete = `
with "deleted" as
(
        delete
        from	"Certificate" "certificate"
		where	"certificate"."certificateId" = $1
		returning "certificateDataId"
)
delete
from    "CertificateData" "data"
where   "data"."certificateDataId" =
(
    select  "certificateDataId"
    from    "deleted"
);`

type CertificateDAO struct {
	database *Database
}

func NewCertificateDAO(database *Database) *CertificateDAO {
	return &CertificateDAO{database}
}

func (dao *CertificateDAO) SaveCertificate(commonAuthorityId int, name string, privateKeyData, certificateData []byte, expiresAt time.Time, userId int) (int, error) {
	var certificateId int
	if err := dao.database.db.QueryRow(
		certificateSave,
		privateKeyData,
		certificateData,
		expiresAt,
		commonAuthorityId,
		name,
		userId).Scan(&certificateId); err != nil {
		return 0, err
	}
	return certificateId, nil
}

func (dao *CertificateDAO) DeleteCertificate(certificateId int) error {
	result, err := dao.database.db.Exec(certificateDelete, certificateId)
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
