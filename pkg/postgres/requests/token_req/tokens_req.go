package token_req

var CreateToken = `INSERT INTO hackerevolution.tokens (id_user, token,update_time)
		VALUES ($1, $2, $3)
		ON CONFLICT (id_user)
		DO UPDATE SET token = EXCLUDED.token,update_time = EXCLUDED.update_time;`
