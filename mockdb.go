package main

import (
	"database/sql"
	"github.com/geekmdio/ehrprotorepo/goproto"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/google/uuid"
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
		NoteGuid:    uuid.New().String(),
		VisitGuid:   uuid.New().String(),
		AuthorGuid:  uuid.New().String(),
		PatientGuid: uuid.New().String(),
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
		NoteFragmentGuid: uuid.New().String(),
		NoteGuid:         note1.GetNoteGuid(),
		IssueGuid:        uuid.New().String(),
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
		NoteGuid:    uuid.New().String(),
		VisitGuid:   uuid.New().String(),
		AuthorGuid:  uuid.New().String(),
		PatientGuid: uuid.New().String(),
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
		NoteFragmentGuid: uuid.New().String(),
		NoteGuid:         note1.GetNoteGuid(),
		IssueGuid:        uuid.New().String(),
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