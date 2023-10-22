package dropdown_req

var GetAllDataInfo = `Select  id_data, name_data from devices.data where deleted_at is null`

var GetDataInfoId = `Select  id_data,name_data from devices.data where id_data = $1 and deleted_at is null`

var GetDataInfoName = `Select id_data,name_data from devices.data where deleted_at is null and LOWER(name_data) like lower('%'||$1||'%')`

var InsertDataInfo = `Insert into devices.data (name_data) values ($1) returning id_data`

var DeleteDataInfo = `Update devices.data set deleted_at = now()::timestamp, updated_at = now()::timestamp where id_data = $1`

var RestoreDataInfo = `Update devices.data set deleted_at = null, updated_at = now()::timestamp where id_data = $1`
var UpdateDataInfoName = `Update devices.data set name_data = $2, updated_at = now()::timestamp where id_data = $1`

var ShowDeleteDataInfo = `Select id_data, name_data from devices.data
                                where (deleted_at +  $1) > now()::timestamp`

var ShowDeleteDataInfoId = `Select id_data, name_data from devices.data
                                where (deleted_at +  $1) > now()::timestamp and id_data = $1`

var ExistsDataInfoName = `SELECT EXISTS(Select * from devices.data where lower(name_data) like lower($1))`
var ExistsDataInfoId = `SELECT EXISTS(Select * from devices.data where id_data = $1)`
