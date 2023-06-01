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
	creacion DATE
)

CREATE TABLE transaciones (
	nro_transaction varchar(250) primary key,
	sender_id integer not null,
	receiver_id integer not null,
	amount float not null,
	movement varchar (250) not null,
	time_transaction timestamp not null,
	FOREIGN KEY (sender_id) REFERENCES wallet (id_persona),
	FOREIGN KEY (receiver_id) REFERENCES wallet (id_persona)
);