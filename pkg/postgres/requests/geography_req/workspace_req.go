package geography_req

var GetAllWorkspace = `Select id_Workspace, code_Workspace,name_Workspace from geography.Workspace where deleted_at is null`

var GetWorkspaceId = `Select id_Workspace, code_Workspace,name_Workspace from geography.Workspace where deleted_at is null and id_Workspace = $1`

var GetWorkspaceIdLocation = `Select id_Workspace, code_Workspace, name_Workspace from geography.Workspace where deleted_at is null and id_Location = $1`

var GetWorkspaceName = `Select id_Workspace, code_Workspace,name_Workspace from geography.Workspace where deleted_at is null and LOWER(name_Workspace) like lower('%'||$1||'%')`

var InsertWorkspace = `Insert into geography.Workspace (code_Workspace, name_Workspace,id_location) values ($1,$2,$3) returning id_Workspace`

var DeleteWorkspace = `Update geography.Workspace set deleted_at = now()::timestamp, updated_at = now()::timestamp where id_Workspace = $1`

var RestoreWorkspace = `Update geography.Workspace set deleted_at = null, updated_at = now()::timestamp where id_Workspace = $1`

var UpdateWorkspaceName = `Update geography.Workspace set name_Workspace = $2, updated_at = now()::timestamp where id_Workspace = $1`

var UpdateWorkspaceAll = `Update geography.Workspace set name_Workspace = $2,code_workspace=$3,id_location=$4,updated_at = now()::timestamp where id_Workspace = $1`
var UpdateWorkspaceReference = `Update geography.Workspace set id_location = $2, updated_at = now()::timestamp where id_Workspace = $1`

var UpdateWorkspaceCode = `Update geography.Workspace set code_Workspace = $2, updated_at = now()::timestamp where id_Workspace = $1`

var ShowDeleteWorkspace = `Select id_Workspace, code_Workspace,name_Workspace from geography.Workspace
                                where (deleted_at +  $1) > now()::timestamp`

var ShowDeleteWorkspaceId = `Select id_Workspace, code_Workspace,name_Workspace from geography.Workspace
                                where (deleted_at +  $1) > now()::timestamp and id_Workspace = $1`

var ExistsWorkspaceName = `SELECT EXISTS(Select * from geography.Workspace where lower(name_Workspace) like lower($1))`
var ExistsWorkspaceId = `SELECT EXISTS(Select * from geography.Workspace where id_workspace = $1)`
var ExistsWorkspaceCode = `SELECT EXISTS(Select * from geography.Workspace where lower(code_Workspace) like lower($1))`

//при ненаходе пустой список
