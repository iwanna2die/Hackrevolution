package roles_req

var DeleteRouteRole = `Delete from workers.connect_roles where id_type_worker = $1`

var PostRouteRole = `Insert into workers.connect_roles ( id_site_role, id_type_worker) values ($1,$2)`
