package kschat

import (
	"encoding/json"
	"strings"

	"github.com/khulnasoft-bot/khulnasoft-bot/khulnasoft-chat-bot/kschat/types/khulnasoft1"
)

type kvstoreMethod string

type kvstoreOptions struct {
	Team       *string `json:"team"`
	Namespace  *string `json:"namespace,omitempty"`
	EntryKey   *string `json:"entryKey,omitempty"`
	EntryValue *string `json:"entryValue,omitempty"`
	Revision   *int    `json:"revision,omitempty"`
}

type kvstoreParams struct {
	Options kvstoreOptions `json:"options"`
}

type kvstoreAPIReq struct {
	Method kvstoreMethod `json:"method"`
	Params kvstoreParams `json:"params"`
}

type GetEntryRes struct {
	Result khulnasoft1.KVGetResult `json:"result"`
	Error  Error                `json:"error,omitempty"`
}

type PutEntryRes struct {
	Result khulnasoft1.KVPutResult `json:"result"`
	Error  Error                `json:"error,omitempty"`
}

type DeleteEntryRes struct {
	Result khulnasoft1.KVDeleteEntryResult `json:"result"`
	Error  Error                        `json:"error,omitempty"`
}

type ListNamespacesRes struct {
	Result khulnasoft1.KVListNamespaceResult `json:"result"`
	Error  Error                          `json:"error,omitempty"`
}

type ListEntryKeysRes struct {
	Result khulnasoft1.KVListEntryResult `json:"result"`
	Error  Error                      `json:"error,omitempty"`
}

type KVStoreAPI interface {
	PutEntry(teamName *string, namespace string, entryKey string, entryValue string) (khulnasoft1.KVPutResult, error)
	PutEntryWithRevision(teamName *string, namespace string, entryKey string, entryValue string, revision int) (khulnasoft1.KVPutResult, error)
	DeleteEntry(teamName *string, namespace string, entryKey string) (khulnasoft1.KVDeleteEntryResult, error)
	DeleteEntryWithRevision(teamName *string, namespace string, entryKey string, revision int) (khulnasoft1.KVDeleteEntryResult, error)
	GetEntry(teamName *string, namespace string, entryKey string) (khulnasoft1.KVGetResult, error)
	ListNamespaces(teamName *string) (khulnasoft1.KVListNamespaceResult, error)
	ListEntryKeys(teamName *string, namespace string) (khulnasoft1.KVListEntryResult, error)
}

func (a *API) PutEntry(teamName *string, namespace string, entryKey string, entryValue string) (result khulnasoft1.KVPutResult, err error) {
	return a.PutEntryWithRevision(teamName, namespace, entryKey, entryValue, 0)
}

func (a *API) PutEntryWithRevision(teamName *string, namespace string, entryKey string, entryValue string, revision int) (result khulnasoft1.KVPutResult, err error) {

	opts := kvstoreOptions{
		Team:       teamName,
		Namespace:  &namespace,
		EntryKey:   &entryKey,
		EntryValue: &entryValue,
	}
	if revision != 0 {
		opts.Revision = &revision
	}
	args := kvstoreAPIReq{Method: "put", Params: kvstoreParams{Options: opts}}
	apiInput, err := json.Marshal(args)
	if err != nil {
		return result, err
	}

	cmd := a.runOpts.Command("kvstore", "api")
	cmd.Stdin = strings.NewReader(string(apiInput))
	bytes, err := cmd.Output()
	if err != nil {
		return result, APIError{err}
	}

	entry := PutEntryRes{}
	err = json.Unmarshal(bytes, &entry)
	if err != nil {
		return result, UnmarshalError{err}
	}
	if entry.Error.Message != "" {
		return result, entry.Error
	}
	return entry.Result, nil
}

func (a *API) DeleteEntry(teamName *string, namespace string, entryKey string) (result khulnasoft1.KVDeleteEntryResult, err error) {
	return a.DeleteEntryWithRevision(teamName, namespace, entryKey, 0)
}

func (a *API) DeleteEntryWithRevision(teamName *string, namespace string, entryKey string, revision int) (result khulnasoft1.KVDeleteEntryResult, err error) {

	opts := kvstoreOptions{
		Team:      teamName,
		Namespace: &namespace,
		EntryKey:  &entryKey,
	}
	if revision != 0 {
		opts.Revision = &revision
	}
	args := kvstoreAPIReq{Method: "del", Params: kvstoreParams{Options: opts}}
	apiInput, err := json.Marshal(args)
	if err != nil {
		return result, err
	}

	cmd := a.runOpts.Command("kvstore", "api")
	cmd.Stdin = strings.NewReader(string(apiInput))
	bytes, err := cmd.Output()
	if err != nil {
		return result, APIError{err}
	}

	entry := DeleteEntryRes{}
	err = json.Unmarshal(bytes, &entry)
	if err != nil {
		return result, UnmarshalError{err}
	}
	if entry.Error.Message != "" {
		return result, entry.Error
	}
	return entry.Result, nil
}

func (a *API) GetEntry(teamName *string, namespace string, entryKey string) (result khulnasoft1.KVGetResult, err error) {

	opts := kvstoreOptions{
		Team:      teamName,
		Namespace: &namespace,
		EntryKey:  &entryKey,
	}
	args := kvstoreAPIReq{Method: "get", Params: kvstoreParams{Options: opts}}
	apiInput, err := json.Marshal(args)
	if err != nil {
		return result, err
	}
	cmd := a.runOpts.Command("kvstore", "api")
	cmd.Stdin = strings.NewReader(string(apiInput))
	bytes, err := cmd.Output()
	if err != nil {
		return result, APIError{err}
	}

	entry := GetEntryRes{}
	err = json.Unmarshal(bytes, &entry)
	if err != nil {
		return result, UnmarshalError{err}
	}
	if entry.Error.Message != "" {
		return result, entry.Error
	}
	return entry.Result, nil
}

func (a *API) ListNamespaces(teamName *string) (result khulnasoft1.KVListNamespaceResult, err error) {

	opts := kvstoreOptions{
		Team: teamName,
	}
	args := kvstoreAPIReq{Method: "list", Params: kvstoreParams{Options: opts}}
	apiInput, err := json.Marshal(args)
	if err != nil {
		return result, err
	}

	cmd := a.runOpts.Command("kvstore", "api")
	cmd.Stdin = strings.NewReader(string(apiInput))
	bytes, err := cmd.Output()
	if err != nil {
		return result, APIError{err}
	}

	var namespaces ListNamespacesRes
	err = json.Unmarshal(bytes, &namespaces)
	if err != nil {
		return result, UnmarshalError{err}
	}
	if namespaces.Error.Message != "" {
		return result, namespaces.Error
	}
	return namespaces.Result, nil
}

func (a *API) ListEntryKeys(teamName *string, namespace string) (result khulnasoft1.KVListEntryResult, err error) {

	opts := kvstoreOptions{
		Team:      teamName,
		Namespace: &namespace,
	}
	args := kvstoreAPIReq{Method: "list", Params: kvstoreParams{Options: opts}}
	apiInput, err := json.Marshal(args)
	if err != nil {
		return result, err
	}

	cmd := a.runOpts.Command("kvstore", "api")
	cmd.Stdin = strings.NewReader(string(apiInput))
	bytes, err := cmd.Output()
	if err != nil {
		return result, APIError{err}
	}

	entryKeys := ListEntryKeysRes{}
	err = json.Unmarshal(bytes, &entryKeys)
	if err != nil {
		return result, UnmarshalError{err}
	}
	if entryKeys.Error.Message != "" {
		return result, entryKeys.Error
	}
	return entryKeys.Result, nil
}
