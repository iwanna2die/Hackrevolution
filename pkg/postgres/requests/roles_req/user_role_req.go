package roles_req

var GetAllUserRole = `Select id_type_worker, name_type_worker from types.type_worker where deleted_at is null`

var GetUserRoleId = `Select id_type_worker, name_type_worker from types.type_worker where deleted_at is null and id_type_worker = $1`

var GetUserRoleName = `Select id_type_worker, name_type_worker from types.type_worker where deleted_at is null and LOWER(name_type_worker) like lower('%'||$1||'%')`

var InsertUserRole = `Insert into types.type_worker (name_type_worker) values ($1) returning id_type_worker`

var DeleteUserRole = `Update types.type_worker set deleted_at = now()::timestamp, updated_at = now()::timestamp where id_type_worker = $1`

var RestoreUserRole = `Update types.type_worker set deleted_at = null, updated_at = now()::timestamp where id_type_worker = $1`

var UpdateUserRoleCode = `Update types.type_worker set name_type_worker = $2, updated_at = now()::timestamp where id_type_worker = $1`

var UpdateUserRoleAll = `Update types.type_worker set name_type_worker = $2 = $3, updated_at = now()::timestamp where id_type_worker = $1`

var ShowDeleteUserRole = `Select	id_type_worker, name_type_worker from types.type_worker
                                where (deleted_at +  $1) > now()::timestamp`

var ShowDeleteUserRoleId = `Select id_type_worker, name_type_worker from types.type_worker
                                where (deleted_at +  $1) > now()::timestamp and id_type_worker = $1`

var ExistsUserRoleName = `SELECT EXISTS(Select * from types.type_worker where lower(name_type_worker) like lower($1))`

var ExistsUserRoleId = `SELECT EXISTS(Select * from types.type_worker where type_worker.id_type_worker = $1)`
