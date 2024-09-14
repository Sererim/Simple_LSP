package analysis

import (
	"fmt"
	"simple_lsp/lsp"
	"strings"
)

type State struct {
	Documents map[string]string
}

func NewState() State {
	return State{Documents: map[string]string{}}
}

func (s *State) OpenDocument(uri, text string) {
	s.Documents[uri] = text
}

func (s *State) UpdateDocument(uri, text string) {
	s.Documents[uri] = text
}

func (s *State) Hover(
	id int, uri string, position lsp.Position) lsp.HoverResponse {

	document := s.Documents[uri]

	return lsp.HoverResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: lsp.HoverResult{
			Contents: fmt.Sprintf(
				"File: %s\nCharacters: %d", uri, len(document)),
		},
	}
}

func (s *State) Definition(
	id int, uri string, position lsp.Position) lsp.DefinitionResponse {

	return lsp.DefinitionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: lsp.Location{
			URI: uri,
			Range: lsp.Range{
				Start: lsp.Position{
					Line:      position.Line - 1,
					Character: 0,
				},
				End: lsp.Position{
					Line:      position.Line - 1,
					Character: 0,
				},
			},
		},
	}
}

func (s *State) TextDocumentCodeAction(id int, uri string) lsp.TextDocumentCodeActionResponse {
	text := s.Documents[uri]

	actions := []lsp.CodeAction{}
	for row, line := range strings.Split(text, "\n") {
		idx := strings.Index(line, "Vim")
		if idx >= 0 {
			replace_change := map[string][]lsp.TextEdit{}
			replace_change[uri] = []lsp.TextEdit{
				{
					Range:   LineRange(row, idx, idx+len("Vim")),
					NewText: "VS Code",
				},
			}
			actions = append(actions, lsp.CodeAction{
				Title: "Replace V*m with a superior editor.",
				Edit:  &lsp.WorkSpaceEdit{Changes: replace_change},
			})
		}
	}

	response := lsp.TextDocumentCodeActionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: actions,
	}

	return response
}

func (s *State) TextDocumentCompletion(id int, uri string) lsp.CompletionResponse {
	items := []lsp.CompletionItem{
		{
			Label:         "VS CODE",
			Detail:        "CHAD EDITOR",
			Documentation: "THANK YOU MICROSOFT GODS",
		},
	}

	response := lsp.CompletionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: items,
	}

	return response
}

func LineRange(line, start, end int) lsp.Range {
	return lsp.Range{
		Start: lsp.Position{
			Line:      line,
			Character: 0,
		},
		End: lsp.Position{
			Line:      line,
			Character: 0,
		},
	}
}
