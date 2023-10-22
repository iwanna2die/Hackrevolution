// при ненаходе пустой список
package geography_req

var GetAllAddress = `Select id_Address, code_Address,name_Address,postcode from geography.Address where deleted_at is null`

var GetAddressId = `Select id_Address, code_Address,name_Address,postcode from geography.Address where deleted_at is null and id_Address = $1`

var GetAddressIdCity = `Select id_Address, code_Address, name_Address,postcode from geography.Address where deleted_at is null and id_city = $1`

var GetAddressName = `Select id_Address, code_Address,name_Address,postcode from geography.Address where deleted_at is null and LOWER(name_Address) like lower('%'||$1||'%')`

var InsertAddress = `Insert into geography.Address (code_Address, name_Address,id_city,postcode) values ($1,$2,$3,$4) returning id_Address`

var DeleteAddress = `Update geography.Address set deleted_at = now()::timestamp, updated_at = now()::timestamp where id_Address = $1`

var RestoreAddress = `Update geography.Address set deleted_at = null, updated_at = now()::timestamp where id_Address = $1`

var UpdateAddressName = `Update geography.Address set name_Address = $2	, updated_at = now()::timestamp where id_address = $1`

var UpdateAddressCode = `Update geography.Address set code_Address = $2, updated_at = now()::timestamp where id_address = $1`

var UpdateAddressReference = `Update geography.Address set id_city = $2, updated_at = now()::timestamp where id_address = $1`

var UpdateAddressPostCode = `Update geography.Address set name_Address = $2, updated_at = now()::timestamp where id_address = $1`

var UpdateAddressAll = `Update geography.Address set name_Address = $2,code_Address = $3,id_city = $4,postcode = $5, updated_at = now()::timestamp where id_address = $1`

var ShowDeleteAddress = `Select	id_Address, code_Address,name_Address,postcode from geography.Address
                                where (deleted_at +  $1) > now()::timestamp`

var ShowDeleteAddressId = `Select id_Address, code_Address,name_Address,postcode from geography.Address
                                where (deleted_at +  $1) > now()::timestamp and id_Address = $1`

var ExistsAddressName = `SELECT EXISTS(Select * from geography.Address where lower(name_Address) like $1)`
var ExistsAddressId = `SELECT EXISTS(Select * from geography.Address where id_address = $1)`
var ExistsAddressCode = `SELECT EXISTS(Select * from geography.Address where lower(code_address) like $1)`

//при ненаходе пустой список
