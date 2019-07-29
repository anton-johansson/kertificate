package db

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
