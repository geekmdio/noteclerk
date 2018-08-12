package main

import (
	"database/sql"
	"github.com/geekmdio/ehrprotorepo/goproto"
	"github.com/golang/protobuf/ptypes/timestamp"
)

type MockDb struct {
	db []*ehrpb.Note

}

func (m *MockDb) Init() (*sql.DB, error) {
	var notes []*ehrpb.Note

	notes = append(notes, buildNote1(), buildNote2())
	m.db = notes
	return nil, nil
}

func (*MockDb) AddNote(note *ehrpb.Note) (id int32, guid string, err error) {
	panic("implement me")
}

func (*MockDb) UpdateNote(note *ehrpb.Note) error {
	panic("implement me")
}

func (*MockDb) DeleteNote(id int32) error {
	panic("implement me")
}

func (m *MockDb) AllNotes() ([]*ehrpb.Note, error) {
	return m.db, nil
}

func (*MockDb) GetNoteById(id int32) (*ehrpb.Note, error) {
	panic("implement me")
}

func (*MockDb) FindNote(filter NoteFindFilter) ([]*ehrpb.Note, error) {
	panic("implement me")
}

func (*MockDb) AddNoteFragment(note *ehrpb.NoteFragment) (id int32, guid string, err error) {
	panic("implement me")
}

func (*MockDb) UpdateNoteFragment(note *ehrpb.NoteFragment) error {
	panic("implement me")
}

func (*MockDb) DeleteNoteFragment(id int32) error {
	panic("implement me")
}

func (*MockDb) AllNoteFragments() ([]*ehrpb.NoteFragment, error) {
	panic("implement me")
}

func (*MockDb) GetNoteFragmentsById(id int32) (*ehrpb.NoteFragment, error) {
	panic("implement me")
}

func (*MockDb) FindNoteFragments(filter NoteFragmentFindFilter) ([]*ehrpb.NoteFragment, error) {
	panic("implement me")
}

func buildNote1() *ehrpb.Note {
	note1 := &ehrpb.Note{
		Id: 0,
		DateCreated: &timestamp.Timestamp{
			Seconds: 1235114,
			Nanos:   323,
		},
		NoteGuid:    "7218e354-9e55-11e8-98d0-529269fb1459",
		VisitGuid:   "7218e7b4-9e55-11e8-98d0-529269fb1459",
		AuthorGuid:  "7218ea84-9e55-11e8-98d0-529269fb1459",
		PatientGuid: "7218eebc-9e55-11e8-98d0-529269fb1459",
		Type:        ehrpb.NoteType_HISTORY_AND_PHYSICAL,
		Fragments:   nil,
		Tags:        nil,
	}
	note1.Tags = append(note1.Tags, "note1tag1", "note1tag2")
	note1.Fragments = append(note1.Fragments, &ehrpb.NoteFragment{
		Id: 0,
		DateCreated: &timestamp.Timestamp{
			Seconds: note1.DateCreated.Seconds + 500,
			Nanos:   note1.DateCreated.Nanos + 550,
		},
		NoteFragmentGuid: "c2af597e-9e55-11e8-98d0-529269fb1459",
		NoteGuid:         note1.GetNoteGuid(),
		IssueGuid:        "7218f15a-9e55-11e8-98d0-529269fb1459",
		Icd_10Code:       "ICD10_Code",
		Icd_10Long:       "ICD10 Long Description",
		Description:      "My consumable description.",
		Status:           ehrpb.NoteFragmentStatus_ACTIVE,
		Priority:         ehrpb.FragmentPriority_HIGH,
		Topic:            ehrpb.FragmentTopic_SUBJECTIVE,
		MarkdownContent:  "This would be markdown content.",
		Tags:             []string{"noteFrag1Tag1"},
	})
	return note1
}

func buildNote2() *ehrpb.Note {
	note1 := &ehrpb.Note{
		Id: 1,
		DateCreated: &timestamp.Timestamp{
			Seconds: 1435114,
			Nanos:   523,
		},
		NoteGuid:    "a0ceabfc-9e55-11e8-98d0-529269fb1459",
		VisitGuid:   "a0ceb25a-9e55-11e8-98d0-529269fb1459",
		AuthorGuid:  "a0ceb502-9e55-11e8-98d0-529269fb1459",
		PatientGuid: "a0ceba84-9e55-11e8-98d0-529269fb1459",
		Type:        ehrpb.NoteType_HISTORY_AND_PHYSICAL,
		Fragments:   nil,
		Tags:        nil,
	}
	note1.Tags = append(note1.Tags, "note2tag1", "note2tag2")
	note1.Fragments = append(note1.Fragments, &ehrpb.NoteFragment{
		Id: 0,
		DateCreated: &timestamp.Timestamp{
			Seconds: note1.DateCreated.Seconds + 500,
			Nanos:   note1.DateCreated.Nanos + 550,
		},
		NoteFragmentGuid: "c2af5c80-9e55-11e8-98d0-529269fb1459",
		NoteGuid:         note1.GetNoteGuid(),
		IssueGuid:        "a0cebd04-9e55-11e8-98d0-529269fb1459",
		Icd_10Code:       "ICD10_Code",
		Icd_10Long:       "ICD10 Long Description",
		Description:      "My consumable description.",
		Status:           ehrpb.NoteFragmentStatus_ACTIVE,
		Priority:         ehrpb.FragmentPriority_HIGH,
		Topic:            ehrpb.FragmentTopic_SUBJECTIVE,
		MarkdownContent:  "This would be markdown content.",
		Tags:             []string{"noteFrag2Tag1"},
	})
	return note1
}