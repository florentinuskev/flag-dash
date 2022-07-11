package repository

const (
	createUserQuery = `INSERT INTO users (email, password) VALUES ($1, $2) RETURNING *`
	createUserProfileQuery = `INSERT INTO user_profiles (user_id, first_name, last_name, phone_number) VALUES ($1, $2, $3, $4) RETURNING *`
	createUserRoleQuery = `INSERT INTO user_roles (user_id, role_id) VALUES ($1, 1)`
	findUserByEmailQuery = `SELECT users.id, users.email, users.password, user_profiles.first_name, user_profiles.last_name, role_types.name as role_name, role_types.level as role_level
							FROM users INNER JOIN user_profiles ON users.id = user_profiles.user_id 
							INNER JOIN user_roles ON users.id = user_roles.user_id 
							INNER JOIN role_types ON user_roles.role_id = role_types.id WHERE email = $1 LIMIT 1`
	findUserByIDQuery = `SELECT users.id, users.email, users.password, user_profiles.first_name, user_profiles.last_name, role_types.name as role_name, role_types.level as role_level
							FROM users INNER JOIN user_profiles ON users.id = user_profiles.user_id 
							INNER JOIN user_roles ON users.id = user_roles.user_id 
							INNER JOIN role_types ON user_roles.role_id = role_types.id WHERE users.id = $1 LIMIT 1`
	findUsersQuery = `SELECT * FROM users LIMIT $1 OFFSET $2`
)