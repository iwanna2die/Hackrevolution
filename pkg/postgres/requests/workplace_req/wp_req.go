package workplace_req

var GetAllWorkplace = `Select id_workplace,description,id_type_workplace,id_workspace from workplaces.workplace where deleted_at is null`

var GetWorkplaceId = `Select id_workplace,description,id_type_workplace,id_workspace from workplaces.workplace where workplace.id_workplace = $1 and deleted_at is null`

var GetWorkplaceName = `Select id_workplace,description,id_type_workplace,id_workspace from workplaces.workplace where deleted_at is null and LOWER(description) like lower('%'||$1||'%')`

var InsertWorkplace = `Insert into workplaces.workplace (description,id_type_workplace,id_workspace) values ($1,$2,$3) returning id_workplace`

var DeleteWorkplace = `Update workplaces.workplace set deleted_at = now()::timestamp, updated_at = now()::timestamp where id_workplace = $1`

var RestoreWorkplace = `Update workplaces.workplace set deleted_at = null, updated_at = now()::timestamp where id_workplace = $1`
var UpdateWorkplaceName = `Update workplaces.workplace set description = $2, updated_at = now()::timestamp where id_workplace = $1`
var UpdateWorkplaceType = `Update workplaces.workplace set id_type_workplace = $2, updated_at = now()::timestamp where id_workplace = $1`
var UpdateWorkplaceWorkspace = `Update workplaces.workplace set id_workspace = $2, updated_at = now()::timestamp where id_workplace = $1`
var UpdateWorkplaceAll = `Update workplaces.workplace set  description = $2,id_type_workplace = $3, id_workspace = $4, updated_at = now()::timestamp where id_workplace = $1`

var ShowDeleteWorkplace = `Select	id_workplace,description,id_type_workplace,id_workspace from workplaces.workplace
                                where (deleted_at +  $1) > now()::timestamp`

var ShowDeleteWorkplaceId = `Select id_workplace,description,id_type_workplace,id_workspace from workplaces.workplace
                                where (deleted_at +  $1) > now()::timestamp and id_workplace = $1`

var ExistsWorkplaceId = `SELECT EXISTS(Select * from workplaces.workplace where id_workplace = $1)`
