package workplace_req

var GetAllTypeWorkplace = `Select id_type_workplace,name_type_workplace from types.type_workplace where deleted_at is null`

var GetTypeWorkplaceId = `Select id_type_workplace,name_type_workplace from types.type_workplace where type_workplace.id_type_workplace = $1 and deleted_at is null`

var GetTypeWorkplaceName = `Select id_type_workplace,name_type_workplace from types.type_workplace where deleted_at is null and LOWER(name_type_workplace) like lower('%'||$1||'%')`

var InsertTypeWorkplace = `Insert into types.type_workplace (name_type_workplace) values ($1) returning id_type_workplace`

var DeleteTypeWorkplace = `Update types.type_workplace set deleted_at = now()::timestamp, updated_at = now()::timestamp where id_type_workplace = $1`

var RestoreTypeWorkplace = `Update types.type_workplace set deleted_at = null, updated_at = now()::timestamp where id_type_workplace = $1`
var UpdateTypeWorkplaceName = `Update types.type_workplace set name_type_workplace = $2, updated_at = now()::timestamp where id_type_workplace = $1`

var ShowDeleteTypeWorkplace = `Select	id_type_workplace,name_type_workplace from types.type_workplace
                                where (deleted_at +  $1) > now()::timestamp`

var ShowDeleteTypeWorkplaceId = `Select id_type_workplace, name_type_workplace from types.type_workplace
                                where (deleted_at +  $1) > now()::timestamp and id_type_workplace = $1`

var ExistsTypeWorkplaceName = `SELECT EXISTS(Select * from types.type_workplace where lower(name_type_workplace) like lower($1))`
var ExistsTypeWorkplaceId = `SELECT EXISTS(Select * from types.type_workplace where id_type_workplace = $1)`
