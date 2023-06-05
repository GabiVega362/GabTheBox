package database

import "fmt"

type Lab struct {
	UUID  string
	Title string
	Image string
}

func (d DatabaseClient) GetAllLabs() ([]Lab, error) {
	// obtenemos todos los laboratorios disponibles
	rows, err := d.client.Query("SELECT id, org, enviroment, release, description FROM labs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//creamos el array de laboratorios
	var labs []Lab
	for rows.Next() {
		var data struct {
			UUID        string
			Org         string
			Enviroment  string
			Release     string
			Description string
		}
		if err := rows.Scan(&data.UUID, &data.Org, &data.Enviroment, &data.Release, &data.Description); err != nil {
			continue
		}
		labs = append(labs, Lab{
			UUID:  data.UUID,
			Title: data.Description,
			Image: fmt.Sprintf("%s/%s:%s", data.Org, data.Enviroment, data.Release),
		})
	}
	return labs, nil
}
