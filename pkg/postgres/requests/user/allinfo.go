package user

var UserAll = `Select fio,name_group,name,epoints,exp,type_role
	from hackerevolution.users 
    inner join hackerevolution.role r on r.id_role = users.id_role
    inner join hackerevolution.organizations o on o.id_org = users.id_org
    inner join hackerevolution.org_group og on og.id_org_group = users.id_org_group
    inner join hackerevolution.groups g on g.id_group = og.id_group
    where id_user = $1`
