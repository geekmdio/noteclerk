package main

//TODO: outsource this to loading SQL files
const createNoteTable = `create table note
(
  id                   serial            not null
    constraint note_pkey
    primary key,
  date_created_seconds integer           not null,
  date_created_nanos   integer default 0 not null,
  note_guid            varchar(38)       not null,
  visit_guid           varchar(38)       not null,
  author_guid          varchar(38)       not null,
  patient_guid         varchar(38)       not null,
  type                 integer           not null,
  status               integer           not null
);

alter table note
  owner to postgres;

create unique index note_id_uindex
  on note (id);

create unique index note_note_guid_uindex
  on note (note_guid);
`

const createNoteFragmentTable = `create table note_fragment
(
  id                   serial            not null
    constraint note_fragment_pkey
    primary key,
  date_created_seconds integer           not null,
  date_created_nanos   integer default 0 not null,
  note_fragment_guid   varchar(38)       not null,
  note_guid            varchar(38)       not null
    constraint note_fragment_note_note_guid_fk
    references note (note_guid),
  icd_10code           varchar(15)       not null,
  icd_10long           varchar(250)      not null,
  description          varchar(150)      not null,
  status               integer           not null,
  priority             integer           not null,
  topic                integer           not null,
  content     varchar(2500)     not null
);

alter table note_fragment
  owner to postgres;

create unique index note_fragment_id_uindex
  on note_fragment (id);

create unique index note_fragment_note_fragment_guid_uindex
  on note_fragment (note_fragment_guid);

`

const createNoteTagTable = `create table note_tag
(
	id serial not null
		constraint note_tag_pkey
			primary key,
	note_guid varchar(38) not null
		constraint note_tag_note_note_guid_fk
			references note (note_guid),
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

const addNoteQuery =
`INSERT INTO "public"."note" 
(
	"id", 
	"date_created_seconds", 
	"date_created_nanos", 
	"note_guid", 
	"visit_guid", 
	"author_guid", 
	"patient_guid", 
	"type", 
	"status"
) 
VALUES 
(
	DEFAULT, 
	$1, 
	$2, 
	$3, 
	$4, 
	$5, 
	$6, 
	$7, 
	$8
);`

const addNoteFragmentQuery =
`INSERT INTO "public"."note_fragment" 
(
	"id", 
	"date_created_seconds", 
	"date_created_nanos", 
	"note_fragment_guid", 
	"note_guid", 
	"icd_10code", 
	"icd_10long", 
	"description", 
	"status", 
	"priority", 
	"topic", 
	"content"
) 
VALUES 
(
	DEFAULT, 
	$1, 
	$2, 
	$3, 
	$4, 
	$5, 
	$6, 
	$7, 
	$8, 
	$9, 
	$10,
	$11
);`

const addNoteTagQuery =
`INSERT INTO "public"."note_tag" 
(
	"id", 
	"note_guid", 
	"tag"
) 
VALUES 
(
	DEFAULT, 
	$1, 
	$2
);`

const addNoteFragmentTagQuery =
`INSERT INTO "public"."note_fragment_tag" 
(
	"id", 
	"note_fragment_guid", 
	"tag"
) 
VALUES 
(
	DEFAULT, 
	$1, 
	$2
);`