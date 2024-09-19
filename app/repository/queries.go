package repository

const CreateUserQ = `
	INSERT INTO USERS(first_name, last_name, email, password, admin, active) VALUES (?, ?, ?, ?, ?, true)
`

const GetUserQ = `
	SELECT first_name, last_name, email, admin, active FROM
		USERS WHERE email = ?
`

const UpdateUserQ = `
	UPDATE TABLE USERS SET first_name = ?, last_name = ?, password = ?, admin = ?
		WHERE email = ?
`

const AuthenticateUserQ = `
	SELECT password FROM USERS where email = ?
`

const DeleteUserQ = `
	UPDATE USERS SET active = false WHERE email = ?
`
