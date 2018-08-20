package main

import (
	"database/sql"
	"fmt"
	"github.com/geekmdio/ehrprotorepo/v1/generated/goproto"
	"github.com/pkg/errors"
	"sort"
	"github.com/geekmdio/noted"
	"github.com/google/uuid"
)

// MockDb implements RDBMSAccessor, but the database is simply a slice of Note pointers. Used in unit testing.
type MockDb struct {
	db []*ehrpb.Note
}

// The database should be initialized after instantiation for all structs implementing the RDBMSAccessor interface.
func (m *MockDb) Initialize(config *Config) (*sql.DB, error) {
	var notes []*ehrpb.Note
	notes = append(notes, buildNote1(), buildNote2())
	m.db = notes

	return nil, nil
}

// Add a note to the mock database.
func (m *MockDb) AddNote(note *ehrpb.Note) (id int64, err error) {
	if note.Id > 0 {
		return 0, errors.New("note has index greater than 0 and is rejected")
	}
	note.Id = m.generateUniqueId()

	m.db = append(m.db, note)

	return note.Id, nil
}

// Update a note which already exists in the mock database.
func (m *MockDb) UpdateNote(note *ehrpb.Note) error {

	var noteIndex int
	found := false
	for k, v := range m.db {
		if v.Id == note.Id {
			noteIndex = k
			found = true
			break
		}
	}

	if !found {
		return errors.New("cannot update note because it could not be found")
	}
	m.db[noteIndex] = note

	return nil
}

// Delete a note from the mock database.
func (m *MockDb) DeleteNote(id int64) error {
	var index int
	var found bool
	for k, n := range m.db {
		if n.Id == id {
			index = k
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("note with id %v not located in database", id)
	}

	var newDb []*ehrpb.Note
	newDb = append(newDb, m.db[:index]...)
	newDb = append(newDb, m.db[index+1:]...)
	m.db = newDb
	return nil
}

// Returns all notes currently stored in the mock database.
func (m *MockDb) AllNotes() ([]*ehrpb.Note, error) {
	return m.db, nil
}

// Get's a Note by it's Id, which should be unique.
func (m *MockDb) GetNoteById(id int64) (*ehrpb.Note, error) {
	var foundNote *ehrpb.Note
	found := false
	for _, v := range m.db {
		if v.Id == id {
			foundNote = v
			found = true
		}
	}

	if !found {
		return nil, errors.New("unable to locate note with that id")
	}

	return foundNote, nil
}

// Find a note using a number of powerful search filters.
func (m *MockDb) FindNote(filter NoteFindFilter) ([]*ehrpb.Note, error) {
	var foundNotes []*ehrpb.Note
	for _, v := range m.db {
		if v.GetVisitGuid() == filter.VisitGuid ||
			v.GetPatientGuid() == filter.PatientGuid ||
			v.GetAuthorGuid() == filter.AuthorGuid {
			foundNotes = append(foundNotes, v)
		}
	}

	if len(foundNotes) == 0 {
		return nil, errors.New("unable to find notes matching query")
	}

	return foundNotes, nil
}

func (*MockDb) AddNoteFragment(note *ehrpb.NoteFragment) (id int64, guid string, err error) {
	panic("implement me")
}

func (*MockDb) UpdateNoteFragment(note *ehrpb.NoteFragment) error {
	panic("implement me")
}

func (*MockDb) AllNoteFragments() ([]*ehrpb.NoteFragment, error) {
	panic("implement me")
}

func (*MockDb) GetNoteFragmentsById(id int64) (*ehrpb.NoteFragment, error) {
	panic("implement me")
}

func (*MockDb) FindNoteFragments(filter NoteFragmentFindFilter) ([]*ehrpb.NoteFragment, error) {
	panic("implement me")
}

func (m *MockDb) AddNoteTag(noteGuid string, tag string) (id int64, err error) {
	panic("implement me")
}

func (m *MockDb) GetNoteTagsByNoteGuid(noteGuid string) (tag []string, err error) {
	panic("implement me")
}

func (m *MockDb) DeleteNoteFragment(noteFragmentGuid string) error {
	panic("implement me")
}

func (m *MockDb) GetNoteFragmentById(id int64) (*ehrpb.NoteFragment, error) {
	panic("implement me")
}

func (m *MockDb) GetNoteFragmentsByNoteGuid(noteGuid string) ([]*ehrpb.NoteFragment, error) {
	panic("implement me")
}

func (m *MockDb) AddNoteFragmentTag(noteGuid string, tag string) (id int64, err error) {
	panic("implement me")
}

func (m *MockDb) GetNoteFragmentTagsByNoteFragmentGuid(noteFragGuid string) (tag []string, err error) {
	panic("implement me")
}

func (*MockDb) createSchema() error {
	panic("implement me")
}

func (m *MockDb) generateUniqueId() int64 {
	var idList []int
	for _, v := range m.db {
		idList = append(idList, int(v.Id))
	}
	sort.Ints(idList)
	listLen := len(idList) - 1
	max := idList[listLen]
	generatedId := int64(max + 1)
	return generatedId
}

func buildNote1() *ehrpb.Note {
	nb := noted.NoteBuilder{}
	note := nb.Init().
		SetId(0).
		SetDateCreated(TimestampNow()).
		SetType(ehrpb.NoteType_HISTORY_AND_PHYSICAL).
		SetVisitGuid(uuid.New().String()).
		SetAuthorGuid(uuid.New().String()).
		SetPatientGuid(uuid.New().String()).
		Build()
	note.Tags = append(note.Tags, "note1tag1", "note1tag2")

	nfb := noted.NoteFragmentBuilder{}
	noteFragment := nfb.InitFromNote(note).
		SetDateCreated(TimestampNow()).
		SetId(0).
		SetDescription("Note 1 Fragment 1 Description").
		SetIcd10LongDescription("ICD10 long description").
		SetIcd10Code("ICD10 Code").
		SetIssueGuid(uuid.New().String()).
		SetTopic(ehrpb.FragmentType_SUBJECTIVE).
		SetPriority(ehrpb.RecordPriority_HIGH).
		SetStatus(ehrpb.RecordStatus_ACTIVE).
		SetMarkdownContent("This is the content of Note 1 Fragment 1.").
		Build()
	noteFragment.Tags = append(noteFragment.Tags, "noteFrag1Tag1", "noteFrag1Tag2")

	note.Fragments = append(note.Fragments, noteFragment)
	return note
}

func buildNote2() *ehrpb.Note {
	nb := noted.NoteBuilder{}
	note := nb.Init().
		SetId(1).
		SetDateCreated(TimestampNow()).
		SetType(ehrpb.NoteType_HISTORY_AND_PHYSICAL).
		SetVisitGuid(uuid.New().String()).
		SetAuthorGuid(uuid.New().String()).
		SetPatientGuid(uuid.New().String()).
		Build()
	note.Tags = append(note.Tags, "note2tag1", "note2tag2")

	nfb := noted.NoteFragmentBuilder{}
	noteFragment := nfb.InitFromNote(note).
		SetDateCreated(TimestampNow()).
		SetId(1).
		SetDescription("Note 2 Fragment 1 Description").
		SetIcd10LongDescription("ICD10 long description").
		SetIcd10Code("ICD10 Code").
		SetIssueGuid(uuid.New().String()).
		SetTopic(ehrpb.FragmentType_SUBJECTIVE).
		SetPriority(ehrpb.RecordPriority_HIGH).
		SetStatus(ehrpb.RecordStatus_ACTIVE).
		SetMarkdownContent("This is the content of Note 2 Fragment 1.").
		Build()
	noteFragment.Tags = append(noteFragment.Tags, "noteFrag2Tag1", "noteFrag2Tag2")

	note.Fragments = append(note.Fragments, noteFragment)
	return note
}
