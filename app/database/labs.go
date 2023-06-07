package database

import "database/sql"

type Lab struct {
	UUID  string
	Title string
	Image string
	Port  uint16
}

func (d DatabaseClient) LabsGetAll(user string) ([]Lab, error) {
	// obtenemos todos los laboratorios disponibles
	query, err := d.client.Prepare("SELECT lab, description, image, COALESCE (port, 0) FROM user_labs WHERE \"user\" = $1 OR \"user\" IS NULL")
	if err != nil {
		return nil, err
	}
	defer query.Close()
	rows, err := query.Query(user)
	if err != nil {
		return nil, err
	}
	// creamos el array de laboratorios
	var labs []Lab
	for rows.Next() {
		var lab Lab
		if err := rows.Scan(&lab.UUID, &lab.Title, &lab.Image, &lab.Port); err != nil {
			continue
		}
		labs = append(labs, lab)
	}
	return labs, nil
}

func (d DatabaseClient) LabsGetImageById(id string) (string, error) {
	query, err := d.client.Prepare("SELECT image FROM images WHERE lab = $1 LIMIT 1")
	if err != nil {
		return "", err
	}
	defer query.Close()
	var image string
	err = query.QueryRow(id).Scan(&image)
	return image, err
}

func (d DatabaseClient) LabsGetContainerByUser(user string) (string, bool, error) {
	query, err := d.client.Prepare("SELECT container FROM users_labs WHERE \"user\" = $1 LIMIT 1")
	if err != nil && err != sql.ErrNoRows {
		return "", false, err
	} else if err == sql.ErrNoRows {
		return "", false, nil
	}
	defer query.Close()
	var container string
	if err := query.QueryRow(user).Scan(&container); err != nil && err != sql.ErrNoRows {
		return "", false, err
	} else if err == sql.ErrNoRows {
		return "", false, nil
	}
	
	return container, true, nil
}

func (d DatabaseClient) LabsUserStarted(user string, lab string, container string, port uint16) error {
	insert, err := d.client.Prepare("INSERT INTO users_labs (\"user\", lab, container, port) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return err
	}
	defer insert.Close()
	_, err = insert.Exec(user, lab, container, port)
	return err
}

func (d DatabaseClient) LabsUserStopped(user string) error {
	delete, err := d.client.Prepare("DELETE FROM users_labs WHERE \"user\" = $1")
	if err != nil {
		return err
	}
	defer delete.Close()
	_, err = delete.Exec(user)
	return err
}
