package amocrmlib

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"github.com/zlietapki/amocrmlib/apimodel"
)

func (t *Token) GetNotes(entityType string, entityId int64) ([]*apimodel.Note, error) {
	url := fmt.Sprintf("/api/v4/%s/%d/notes", entityType, entityId)
	resp, err := t.DoRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	if len(resp) == 0 {
		return nil, nil
	}

	notesResp := new(apimodel.NotesListResponse)
	if err = notesResp.UnmarshalBinary(resp); err != nil {
		return nil, errors.WithStack(err)
	}

	return notesResp.Embedded.Notes, nil
}

func (t *Token) GetNote(entityType string, noteId int64) (*apimodel.Note, error) {
	url := fmt.Sprintf("/api/v4/%s/notes/%d", entityType, noteId)
	resp, err := t.DoRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	if len(resp) == 0 {
		return nil, errors.New("Note not found")
	}

	apiNote := new(apimodel.Note)
	if err = apiNote.UnmarshalBinary(resp); err != nil {
		return nil, errors.WithStack(err)
	}

	return apiNote, nil
}

func (t *Token) AddNoteCommon(entityType string, entityId int64, text string) (*apimodel.NoteAddResponse, error) {
	reqStruct := []apimodel.NoteAddRequest{
		{
			NoteType: "common",
			Params: &apimodel.NoteParams{
				Text: text,
			},
		},
	}
	bodyReq, err := json.Marshal(reqStruct)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	path := fmt.Sprintf("/api/v4/%s/%d/notes", entityType, entityId)
	resp, err := t.DoRequest(http.MethodPost, path, bodyReq)
	if err != nil {
		return nil, err
	}

	addNoteResp := new(apimodel.NoteAddResponse)
	if err = addNoteResp.UnmarshalBinary(resp); err != nil {
		return nil, errors.WithStack(err)
	}

	return addNoteResp, nil
}
