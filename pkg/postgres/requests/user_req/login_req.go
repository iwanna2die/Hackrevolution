package user_req

var ExistLoginPwd = `select Exists(Select * from hackerevolution.users where login = $1 and password = $2)`

var TakeUserInfo = ` select id_user,fio,name,name_group from hackerevolution.users 
     inner join hackerevolution.organizations o on o.id_org = users.id_org 
    inner join hackerevolution.org_group og on og.id_org_group = users.id_org_group
    inner join hackerevolution.groups g on g.id_group = og.id_group
	where login = $1 and password = $2`

var TakeUserRole = `select r.id_role, type_role 
from hackerevolution.users inner join hackerevolution.role r on r.id_role = users.id_role where id_user = $1`
