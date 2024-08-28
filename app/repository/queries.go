package repository

const CreateUserQ = `
	INSERT INTO USERS(first_name, last_name, password, admin, active) VALUES (?, ?, ?, ?, true)
`

const GetUserQ = `
	SELECT first_name, last_name, email, admin, active FROM
		USER WHERE email = ?
`

const UpdateUserQ = `
	UPDATE TABLE USER SET first_name = ?, last_name = ?, password = ?, admin = ?
		WHERE id = ?
`

const AuthenticateUserQ = `
	SELECT password FROM USER where email = ?
`

const DeleteUserQ = `
	UPDATE USER SET active = false WHERE email = ?
`
