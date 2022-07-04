package backend

import "GoReact/backend/secure"

func Start(){
	SetDBConnection(secure.NewDBOptions())
	router := setRouter()

	router.Run(":8080")
}