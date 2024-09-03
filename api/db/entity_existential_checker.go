package db

import (
	"fmt"
)

func EntityExists(conn Connection, entityID uuid_t, entityName string) bool {

	rows, err := conn.Gateway.Query(
		fmt.Sprintf(`SELECT %s_id FROM rewild.%ss WHERE %s_id=$1`, entityName, entityName, entityName),
		entityID,
	)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	doesExist := false
	for rows.Next() {
		doesExist = true
	}

	return doesExist

}
