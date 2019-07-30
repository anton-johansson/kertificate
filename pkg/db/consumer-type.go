package db

import (
	"database/sql"
	"fmt"
)

const consumerTypeGet = `
select	"type"."typeId"
,		"type"."name"
from	"ConsumerType" "type"
where	"type"."typeId" = $1;
`

const consumerTypeList = `
select	"type"."typeId"
,		"type"."name"
from	"ConsumerType" "type"
order by "type"."name"
`

const consumerTypeCreate = `
insert into "ConsumerType" ("name")
values ($1)
returning "typeId";
`

const consumerTypeUpdate = `
update	"ConsumerType" "type"
set		"name" = $2
where	"type"."typeId" = $1;
`

const consumerTypeDelete = `
delete
from	"ConsumerType" "type"
where	"type"."typeId" = $1;
`

type ConsumerType struct {
	ConsumerTypeData
	TypeId int `json:"typeId"`
}

type ConsumerTypeData struct {
	Name string `json:"name"`
}

type ConsumerTypeDAO struct {
	database *Database
}

func NewConsumerTypeDAO(database *Database) *ConsumerTypeDAO {
	return &ConsumerTypeDAO{
		database,
	}
}

func (dao *ConsumerTypeDAO) Get(typeId int) (ConsumerType, error) {
	row := dao.database.db.QueryRow(consumerTypeGet, typeId)
	var consumerType ConsumerType
	err := row.Scan(&consumerType.TypeId, &consumerType.Name)
	if err != nil {
		return ConsumerType{}, err
	}
	return consumerType, nil
}

func (dao *ConsumerTypeDAO) List() ([]ConsumerType, error) {
	rows, err := dao.database.db.Query(consumerTypeList)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var consumerTypes []ConsumerType
	for rows.Next() {
		var consumerType ConsumerType
		if err := rows.Scan(&consumerType.TypeId, &consumerType.Name); err != nil {
			fmt.Println("error scanning type in list:", err)
			continue
		}
		consumerTypes = append(consumerTypes, consumerType)
	}
	return consumerTypes, nil
}

func (dao *ConsumerTypeDAO) Update(typeId int, data ConsumerTypeData) error {
	result, err := dao.database.db.Exec(consumerTypeUpdate, typeId, data.Name)
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

func (dao *ConsumerTypeDAO) Delete(typeId int) error {
	result, err := dao.database.db.Exec(consumerTypeDelete, typeId)
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

func (dao *ConsumerTypeDAO) Create(data ConsumerTypeData) (int, error) {
	var typeId int
	if err := dao.database.db.QueryRow(consumerTypeCreate, data.Name).Scan(&typeId); err != nil {
		return 0, err
	}
	return typeId, nil
}
