package main

// TODO: Consider moving to SQL files
const createNoteTable =
`CREATE TABLE IF NOT EXISTS note
(
  id                   serial            NOT NULL
    CONSTRAINT note_pkey
    PRIMARY KEY, 
  date_created_seconds integer           NOT NULL,
  date_created_nanos   integer default 0 NOT NULL,
  note_guid            varchar(38)       NOT NULL,
  visit_guid           varchar(38)       NOT NULL,
  author_guid          varchar(38)       NOT NULL,
  patient_guid         varchar(38)       NOT NULL,
  type                 integer           NOT NULL,
  status               integer           NOT NULL
);

ALTER TABLE note
  OWNER to postgres;

CREATE UNIQUE INDEX note_id_uindex
  ON note (id);

CREATE UNIQUE INDEX note_note_guid_uindex
  ON note (note_guid);
`

const createNoteFragmentTable =
`CREATE TABLE IF NOT EXISTS note_fragment
(
  id                   serial            NOT NULL
    CONSTRAINT note_fragment_pkey
    PRIMARY KEY, 
  date_created_seconds integer           NOT NULL,
  date_created_nanos   integer default 0 NOT NULL,
  note_fragment_guid   varchar(38)       NOT NULL,
  note_guid            varchar(38)       NOT NULL
    CONSTRAINT note_fragment_note_note_guid_fk
    REFERENCES note (note_guid)
	ON DELETE CASCADE,
  icd_10code           varchar(15)       NOT NULL,
  icd_10long           varchar(250)      NOT NULL,
  description          varchar(150)      NOT NULL,
  status               integer           NOT NULL,
  priority             integer           NOT NULL,
  topic                integer           NOT NULL,
  content     varchar(2500)     NOT NULL
);

ALTER TABLE note_fragment
  OWNER to postgres;

CREATE UNIQUE INDEX note_fragment_id_uindex
  ON note_fragment (id);

CREATE UNIQUE INDEX note_fragment_note_fragment_guid_uindex
  ON note_fragment (note_fragment_guid);

`

const createNoteTagTable =
`CREATE TABLE IF NOT EXISTS note_tag
(
	id serial NOT NULL
		CONSTRAINT note_tag_pkey
			PRIMARY KEY, 
	note_guid varchar(38) NOT NULL
		CONSTRAINT note_tag_note_note_guid_fk
			REFERENCES note (note_guid)
			ON DELETE CASCADE,
	tag varchar(55) NOT NULL
)
;

ALTER TABLE note_tag OWNER to postgres
;

CREATE UNIQUE INDEX note_tag_id_uindex
	ON note_tag (id)
;

`

const createNoteFragmentTagTable =
`CREATE TABLE IF NOT EXISTS note_fragment_tag
(
	id serial NOT NULL
		CONSTRAINT note_fragment_tag_pkey
			PRIMARY KEY, 
	note_fragment_guid varchar(38) NOT NULL
		CONSTRAINT note_fragment_tag_note_fragment_note_fragment_guid_fk
			REFERENCES note_fragment (note_fragment_guid)
			ON DELETE CASCADE,
	tag varchar(55) NOT NULL
)
;

ALTER TABLE note_fragment_tag OWNER to postgres
;

CREATE UNIQUE INDEX note_fragment_tag_id_uindex
	ON note_fragment_tag (id)
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
)
RETURNING id;`

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
)
RETURNING id;`

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
)
RETURNING id;`

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
)
RETURNING id;`

const getNoteTagByNoteGuid =
`SELECT * from note_tag WHERE note_guid = $1;`

const getNoteFragmentTagsByNoteFragmentGuid =
`SELECT * from note_fragment_tag WHERE note_fragment_guid = $1;`

const getNoteFragmentByNoteGuid =
`SELECT * from note_fragment WHERE note_guid = $1;`

const getNoteByIdQuery =
`SELECT * from note WHERE id = $1;`

const updateNoteFragmentStatusToStatusByNoteFragmentGuidQuery =
`UPDATE note_fragment
SET status = $1
WHERE note_fragment_guid = $2
RETURNING id;`

const updateNoteStatusToStatusByNoteIdQuery =
`UPDATE note
SET status = $1
WHERE id = $2
RETURNING id;`
