package roles_req

var GetAllSiteRoute = `Select id_site_role, name_route,description_role from workers.site_role where deleted_at is null`

var GetSiteRouteId = `Select id_site_role, name_route,description_role from workers.site_role where deleted_at is null and id_site_role = $1`

var GetSiteRouteName = `Select id_site_role, name_route,description_role from workers.site_role where deleted_at is null and LOWER(name_route) like lower('%'||$1||'%')`

var InsertSiteRoute = `Insert into workers.site_role (name_route, description_role) values ($1,$2) returning id_site_role`

var DeleteSiteRoute = `Update workers.site_role set deleted_at = now()::timestamp, updated_at = now()::timestamp where id_site_role = $1`

var RestoreSiteRoute = `Update workers.site_role set deleted_at = null, updated_at = now()::timestamp where id_site_role = $1`

var UpdateSiteRouteName = `Update workers.site_role set description_role = $2, updated_at = now()::timestamp where id_site_role = $1`

var UpdateSiteRouteCode = `Update workers.site_role set name_route = $2, updated_at = now()::timestamp where id_site_role = $1`

var UpdateSiteRouteAll = `Update workers.site_role set name_route = $2,description_role = $3, updated_at = now()::timestamp where id_site_role = $1`

var ShowDeleteSiteRoute = `Select	id_site_role, name_route,description_role from workers.site_role
                                where (deleted_at +  $1) > now()::timestamp`

var ShowDeleteSiteRouteId = `Select id_site_role, name_route,description_role from workers.site_role
                                where (deleted_at +  $1) > now()::timestamp and id_site_role = $1`

var ExistsSiteRouteName = `SELECT EXISTS(Select * from workers.site_role where lower(name_route) like lower($1))`

var ExistsSiteRouteId = `SELECT EXISTS(Select * from workers.site_role where site_role.id_site_role = $1)`
