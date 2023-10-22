package geography_req

var GetAllRegion = `Select id_Region, code_Region,name_Region from geography.Region where deleted_at is null`

var GetRegionId = `Select id_Region, code_Region,name_Region from geography.Region where deleted_at is null and id_Region = $1`

var GetRegionIdCountry = `Select id_Region, code_Region, name_Region from geography.region where deleted_at is null and id_country = $1`

var GetRegionName = `Select id_Region, code_Region,name_Region from geography.Region where deleted_at is null and LOWER(name_Region) like lower('%'||$1||'%')`

var InsertRegion = `Insert into geography.Region (code_Region, name_Region,id_country) values ($1,$2,$3) returning id_Region`

var DeleteRegion = `Update geography.Region set deleted_at = now()::timestamp, updated_at = now()::timestamp where id_Region = $1`

var RestoreRegion = `Update geography.Region set deleted_at = null, updated_at = now()::timestamp where id_Region = $1`

var UpdateRegionName = `Update geography.Region set name_Region = $2, updated_at = now()::timestamp where id_region = $1`

var UpdateRegionCode = `Update geography.Region set code_Region = $1, updated_at = now()::timestamp where id_region = $1`

var UpdateRegionAll = `Update geography.Region set code_Region = $2,name_region = $3,id_country = $4, updated_at = now()::timestamp where id_region = $1`

var UpdateRegionReference = `Update geography.Region set id_country = $2, updated_at = now()::timestamp where id_region = $1`

var ShowDeleteRegion = `Select	id_Region, code_Region,name_Region from geography.Region
                                where (deleted_at +  $1) > now()::timestamp`

var ShowDeleteRegionId = `Select id_Region, code_Region,name_Region from geography.Region
                                where (deleted_at +  $1) > now()::timestamp and id_Region = $1`

var ExistsRegionName = `SELECT EXISTS(Select * from geography.Region where lower(name_Region) like lower($1))`
var ExistsRegionId = `SELECT EXISTS(Select * from geography.Region where id_region = $1)`
var ExistsRegionCode = `SELECT EXISTS(Select * from geography.Region where lower(code_Region) like lower($1))`

//при ненаходе пустой список
