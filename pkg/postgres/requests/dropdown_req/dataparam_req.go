package dropdown_req

var GetAllDataParam = `Select  id_data_property, name_data_property from devices.data_property where deleted_at is null`

var GetDataParamId = `Select  id_data_property,name_data_property from devices.data_property where id_data_property = $1 and deleted_at is null`

var GetDataParamName = `Select id_data_property,name_data_property from devices.data_property where deleted_at is null and LOWER(name_data_property) like lower('%'||$1||'%')`

var GetDataParamReference = `Select id_data_property,name_data_property from devices.data_property where id_data = $1`

var InsertDataParam = `Insert into devices.data_property (name_data_property,id_data) values ($1,$2) returning id_data_property`

var DeleteDataParam = `Update devices.data_property set deleted_at = now()::timestamp, updated_at = now()::timestamp where id_data_property = $1`

var RestoreDataParam = `Update devices.data_property set deleted_at = null, updated_at = now()::timestamp where id_data_property = $1`
var UpdateDataParamName = `Update devices.data_property set name_data_property = $2,id_data = $3, updated_at = now()::timestamp where id_data_property = $1`

var ShowDeleteDataParam = `Select id_data_property,id_data, name_data_property from devices.data_property
                                where (deleted_at +  $1) > now()::timestamp`

var ShowDeleteDataParamId = `Select id_data_property,id_data, name_data_property from devices.data_property
                                where (deleted_at +  $1) > now()::timestamp and id_data_property = $1`

var ExistsDataParamName = `SELECT EXISTS(Select * from devices.data_property where lower(name_data_property) like lower($1))`
var ExistsDataParamId = `SELECT EXISTS(Select * from devices.data_property where id_data_property = $1)`
