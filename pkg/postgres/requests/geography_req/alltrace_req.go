package geography_req

var GetAllTrace = `select 
    geography.country.id_country, geography.country.name_country,geography.country.code_country,
     r.id_region,r.name_region,r.code_region,
     c.id_city,c.name_city,c.code_city,
     a.id_address,a.name_address,a.postcode,a.code_address,
       l.id_location,l.name_location, l.code_location,
w.id_workspace, w.name_workspace, w.code_workspace
from geography.country  inner join geography.region r on country.id_country = r.id_country
                        inner join geography.city c on r.id_region = c.id_region
                        inner join geography.address a on c.id_city = a.id_city
                        inner join geography.location l on a.id_address = l.id_address
                        inner join geography.workspace w on l.id_location = w.id_location where id_workspace = $1`
