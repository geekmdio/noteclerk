package main


const createNoteTable = `create table createNoteTable
(
	id serial not null
		constraint note_pkey
			primary key,
	date_created_seconds integer not null,
	date_created_nanos integer default 0 not null,
	note_guid varchar(38) not null,
	visit_guid varchar(38) not null,
	author_guid varchar(38) not null,
	patient_guid varchar(38) not null,
	type integer not null
)
;

alter table createNoteTable owner to postgres
;

create unique index note_id_uindex
	on createNoteTable (id)
;

create unique index note_note_guid_uindex
	on createNoteTable (note_guid)
;


`

const createNoteFragmentTable = `create table note_fragment
(
	id serial not null
		constraint note_fragment_pkey
			primary key,
	date_created_seconds integer not null,
	date_created_nanos integer default 0 not null,
	note_fragment_guid varchar(38) not null,
	note_guid varchar(38) not null
		constraint note_fragment_note_note_guid_fk
			references createNoteTable (note_guid),
	icd_10code varchar(15) not null,
	icd_10long varchar(250) not null,
	description varchar(150) not null,
	status integer not null,
	priority integer not null,
	topic integer not null,
	markdown_content varchar(2500) not null
)
;

alter table note_fragment owner to postgres
;

create unique index note_fragment_id_uindex
	on note_fragment (id)
;

create unique index note_fragment_note_fragment_guid_uindex
	on note_fragment (note_fragment_guid)
;

`

const createNoteTagTable = `create table note_tag
(
	id serial not null
		constraint note_tag_pkey
			primary key,
	note_guid varchar(38) not null
		constraint note_tag_note_note_guid_fk
			references createNoteTable (note_guid),
	tag varchar(55) not null
)
;

alter table note_tag owner to postgres
;

create unique index note_tag_id_uindex
	on note_tag (id)
;


`

const createNoteFragmentTagTable = `create table note_fragment_tag
(
	id serial not null
		constraint note_fragment_tag_pkey
			primary key,
	note_fragment_guid varchar(38) not null
		constraint note_fragment_tag_note_fragment_note_fragment_guid_fk
			references note_fragment (note_fragment_guid),
	tag varchar(55) not null
)
;

alter table note_fragment_tag owner to postgres
;

create unique index note_fragment_tag_id_uindex
	on note_fragment_tag (id)
;



`