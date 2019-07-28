package db

import (
	"database/sql"
	"fmt"
)

const certificateTemplateGet = `
select	"template"."templateId"
,		"template"."name"
,		"template"."countryCode"
,		"template"."state"
,		"template"."locality"
,		"template"."organizationName"
,		"template"."organizationalUnit"
,		"template"."emailAddress"
,		"template"."validFor"
,		"template"."keySize"
,		"template"."digest"
from	"CertificateTemplate" "template"
where	"template"."templateId" = $1;
`

const certificateTemplateList = `
select	"template"."templateId"
,		"template"."name"
,		"template"."countryCode"
,		"template"."state"
,		"template"."locality"
,		"template"."organizationName"
,		"template"."organizationalUnit"
,		"template"."emailAddress"
,		"template"."validFor"
,		"template"."keySize"
,		"template"."digest"
from	"CertificateTemplate" "template"
order by "template"."name"
`

const certificateTemplateCreate = `
insert into "CertificateTemplate"
(
    "name"
,   "countryCode"
,   "state"
,   "locality"
,   "organizationName"
,   "organizationalUnit"
,   "emailAddress"
,   "validFor"
,   "keySize"
,   "digest"
)
values
(
    $1
,   $2
,   $3
,   $4
,   $5
,   $6
,   $7
,   $8
,   $9
,   $10
)
returning "templateId";
`

const certificateTemplateUpdate = `
update	"CertificateTemplate" "template"
set		"name" = $2
,   	"countryCode" = $3
,   	"state" = $4
,   	"locality" = $5
,   	"organizationName" = $6
,   	"organizationalUnit" = $7
,   	"emailAddress" = $8
,   	"validFor" = $9
,   	"keySize" = $10
,		"digest" = $11
where	"template"."templateId" = $1;
`

const certificateTemplateDelete = `
delete
from	"CertificateTemplate" "template"
where	"template"."templateId" = $1;
`

type CertificateTemplate struct {
	CertificateTemplateData
	TemplateId int `json:"templateId"`
}

type CertificateTemplateData struct {
	Name               string `json:"name"`
	CountryCode        string `json:"countryCode"`
	State              string `json:"state"`
	Locality           string `json:"locality"`
	OrganizationName   string `json:"organizationName"`
	OrganizationalUnit string `json:"organizationalUnit"`
	EmailAddress       string `json:"emailAddress"`
	ValidFor           int    `json:"validFor"`
	KeySize            string `json:"keySize"`
	Digest             string `json:"digest"`
}

type CertificateTemplateDAO struct {
	database *Database
}

func NewCertificateTemplateDAO(database *Database) *CertificateTemplateDAO {
	return &CertificateTemplateDAO{database}
}

func (dao *CertificateTemplateDAO) Get(templateId int) (CertificateTemplate, error) {
	row := dao.database.db.QueryRow(certificateTemplateGet, templateId)
	var certificateTemplate CertificateTemplate
	if err := row.Scan(
		&certificateTemplate.TemplateId,
		&certificateTemplate.Name,
		&certificateTemplate.CountryCode,
		&certificateTemplate.State,
		&certificateTemplate.Locality,
		&certificateTemplate.OrganizationName,
		&certificateTemplate.OrganizationalUnit,
		&certificateTemplate.EmailAddress,
		&certificateTemplate.ValidFor,
		&certificateTemplate.KeySize,
		&certificateTemplate.Digest); err != nil {
		return CertificateTemplate{}, err
	}
	return certificateTemplate, nil
}

func (dao *CertificateTemplateDAO) List() ([]CertificateTemplate, error) {
	rows, err := dao.database.db.Query(certificateTemplateList)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var certificateTemplates []CertificateTemplate
	for rows.Next() {
		var certificateTemplate CertificateTemplate
		if err := rows.Scan(
			&certificateTemplate.TemplateId,
			&certificateTemplate.Name,
			&certificateTemplate.CountryCode,
			&certificateTemplate.State,
			&certificateTemplate.Locality,
			&certificateTemplate.OrganizationName,
			&certificateTemplate.OrganizationalUnit,
			&certificateTemplate.EmailAddress,
			&certificateTemplate.ValidFor,
			&certificateTemplate.KeySize,
			&certificateTemplate.Digest); err != nil {
			fmt.Println("error scanning type in list:", err)
			continue
		}
		certificateTemplates = append(certificateTemplates, certificateTemplate)
	}
	return certificateTemplates, nil
}

func (dao *CertificateTemplateDAO) Update(templateId int, data CertificateTemplateData) error {
	result, err := dao.database.db.Exec(certificateTemplateUpdate,
		templateId,
		data.Name,
		data.CountryCode,
		data.State,
		data.Locality,
		data.OrganizationName,
		data.OrganizationalUnit,
		data.EmailAddress,
		data.ValidFor,
		data.KeySize,
		data.Digest)
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

func (dao *CertificateTemplateDAO) Delete(templateId int) error {
	result, err := dao.database.db.Exec(certificateTemplateDelete, templateId)
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

func (dao *CertificateTemplateDAO) Create(data CertificateTemplateData) (int, error) {
	var templateId int
	if err := dao.database.db.QueryRow(certificateTemplateCreate,
		data.Name,
		data.CountryCode,
		data.State,
		data.Locality,
		data.OrganizationName,
		data.OrganizationalUnit,
		data.EmailAddress,
		data.ValidFor,
		data.KeySize,
		data.Digest).Scan(&templateId); err != nil {
		return 0, err
	}
	return templateId, nil
}
