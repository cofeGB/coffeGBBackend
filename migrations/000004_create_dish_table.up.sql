CREATE TABLE dish(
    id uuid NOT NULL primary key,
	created_at DATE not null,
	updated_at DATE not null,
	category  uuid not null,
	creator_gu_guid uuid not null,
	title         varchar(256) not null,
	description   text not null ,
	weight        float not null,
	volume        float not null,
	price        float not null,
	quantity      integer not null,
	availability integer not null,
	warnings     Text not null default 'no warning' ,
	is_run boolean not null default true
);
