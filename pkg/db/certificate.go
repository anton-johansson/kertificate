package db

import "time"

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
returning "certificateId";
`

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
