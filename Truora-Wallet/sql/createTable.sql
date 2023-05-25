CREATE TABLE solicitud (
	id_persona int PRIMARY KEY,
	dni_solicitud VARCHAR(250),
	fecha_solicitud DATE,
	pais varchar(50),
	estado varchar(50),
	codigo int
)

CREATE TABLE wallet (
	id_persona int PRIMARY KEY REFERENCES solicitud(id_persona),
	dni VARCHAR(250),
	pais_id VARCHAR(50),
	creacion VARCHAR(50)
)