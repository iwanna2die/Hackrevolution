package geography_req

var GetAllCountry = `Select id_country, code_country,name_country from geography.country where deleted_at is null`

var GetCountryId = `Select id_country, code_country,name_country from geography.country where deleted_at is null and id_country = $1`

var GetCountryName = `Select id_country, code_country,name_country from geography.country where deleted_at is null and LOWER(name_country) like lower('%'||$1||'%')`

var InsertCountry = `Insert into geography.country (code_country, name_country) values ($1,$2) returning id_country`

var DeleteCountry = `Update geography.country set deleted_at = now()::timestamp, updated_at = now()::timestamp where id_country = $1`

var RestoreCountry = `Update geography.country set deleted_at = null, updated_at = now()::timestamp where id_country = $1`

var UpdateCountryName = `Update geography.country set name_country = $2, updated_at = now()::timestamp where id_country = $1`

var UpdateCountryCode = `Update geography.country set code_country = $2, updated_at = now()::timestamp where id_country = $1`

var UpdateCountryAll = `Update geography.country set code_country = $2,name_country = $3, updated_at = now()::timestamp where id_country = $1`

var ShowDeleteCountry = `Select	id_country, code_country,name_country from geography.country
                                where (deleted_at +  $1) > now()::timestamp`

var ShowDeleteCountryId = `Select id_country, code_country,name_country from geography.country
                                where (deleted_at +  $1) > now()::timestamp and id_country = $1`

var ExistsCountryName = `SELECT EXISTS(Select * from geography.country where lower(name_country) like lower($1))`

var ExistsCountryId = `SELECT EXISTS(Select * from geography.country where country.id_country = $1)`

var ExistsCountryCode = `SELECT EXISTS(Select * from geography.country where  lower(code_country) like lower($1))`
