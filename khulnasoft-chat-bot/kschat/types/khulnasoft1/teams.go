// Auto-generated to Go types using avdl-compiler v1.4.10 (https://github.com/khulnasoft/node-avdl-compiler)
//   Input file: ../client/protocol/avdl/khulnasoft1/teams.avdl

package khulnasoft1

import (
	"errors"
	"fmt"
)

type TeamRole int

const (
	TeamRole_NONE          TeamRole = 0
	TeamRole_READER        TeamRole = 1
	TeamRole_WRITER        TeamRole = 2
	TeamRole_ADMIN         TeamRole = 3
	TeamRole_OWNER         TeamRole = 4
	TeamRole_BOT           TeamRole = 5
	TeamRole_RESTRICTEDBOT TeamRole = 6
)

func (o TeamRole) DeepCopy() TeamRole { return o }

var TeamRoleMap = map[string]TeamRole{
	"NONE":          0,
	"READER":        1,
	"WRITER":        2,
	"ADMIN":         3,
	"OWNER":         4,
	"BOT":           5,
	"RESTRICTEDBOT": 6,
}

var TeamRoleRevMap = map[TeamRole]string{
	0: "NONE",
	1: "READER",
	2: "WRITER",
	3: "ADMIN",
	4: "OWNER",
	5: "BOT",
	6: "RESTRICTEDBOT",
}

func (e TeamRole) String() string {
	if v, ok := TeamRoleRevMap[e]; ok {
		return v
	}
	return fmt.Sprintf("%v", int(e))
}

type TeamApplication int

const (
	TeamApplication_KBFS                TeamApplication = 1
	TeamApplication_CHAT                TeamApplication = 2
	TeamApplication_SALTPACK            TeamApplication = 3
	TeamApplication_GIT_METADATA        TeamApplication = 4
	TeamApplication_SEITAN_INVITE_TOKEN TeamApplication = 5
	TeamApplication_STELLAR_RELAY       TeamApplication = 6
	TeamApplication_KVSTORE             TeamApplication = 7
)

func (o TeamApplication) DeepCopy() TeamApplication { return o }

var TeamApplicationMap = map[string]TeamApplication{
	"KBFS":                1,
	"CHAT":                2,
	"SALTPACK":            3,
	"GIT_METADATA":        4,
	"SEITAN_INVITE_TOKEN": 5,
	"STELLAR_RELAY":       6,
	"KVSTORE":             7,
}

var TeamApplicationRevMap = map[TeamApplication]string{
	1: "KBFS",
	2: "CHAT",
	3: "SALTPACK",
	4: "GIT_METADATA",
	5: "SEITAN_INVITE_TOKEN",
	6: "STELLAR_RELAY",
	7: "KVSTORE",
}

func (e TeamApplication) String() string {
	if v, ok := TeamApplicationRevMap[e]; ok {
		return v
	}
	return fmt.Sprintf("%v", int(e))
}

type TeamStatus int

const (
	TeamStatus_NONE      TeamStatus = 0
	TeamStatus_LIVE      TeamStatus = 1
	TeamStatus_DELETED   TeamStatus = 2
	TeamStatus_ABANDONED TeamStatus = 3
)

func (o TeamStatus) DeepCopy() TeamStatus { return o }

var TeamStatusMap = map[string]TeamStatus{
	"NONE":      0,
	"LIVE":      1,
	"DELETED":   2,
	"ABANDONED": 3,
}

var TeamStatusRevMap = map[TeamStatus]string{
	0: "NONE",
	1: "LIVE",
	2: "DELETED",
	3: "ABANDONED",
}

func (e TeamStatus) String() string {
	if v, ok := TeamStatusRevMap[e]; ok {
		return v
	}
	return fmt.Sprintf("%v", int(e))
}

type AuditMode int

const (
	AuditMode_STANDARD           AuditMode = 0
	AuditMode_JUST_CREATED       AuditMode = 1
	AuditMode_SKIP               AuditMode = 2
	AuditMode_STANDARD_NO_HIDDEN AuditMode = 3
)

func (o AuditMode) DeepCopy() AuditMode { return o }

var AuditModeMap = map[string]AuditMode{
	"STANDARD":           0,
	"JUST_CREATED":       1,
	"SKIP":               2,
	"STANDARD_NO_HIDDEN": 3,
}

var AuditModeRevMap = map[AuditMode]string{
	0: "STANDARD",
	1: "JUST_CREATED",
	2: "SKIP",
	3: "STANDARD_NO_HIDDEN",
}

func (e AuditMode) String() string {
	if v, ok := AuditModeRevMap[e]; ok {
		return v
	}
	return fmt.Sprintf("%v", int(e))
}

type PerTeamKeyGeneration int

func (o PerTeamKeyGeneration) DeepCopy() PerTeamKeyGeneration {
	return o
}

type PTKType int

const (
	PTKType_READER PTKType = 0
)

func (o PTKType) DeepCopy() PTKType { return o }

var PTKTypeMap = map[string]PTKType{
	"READER": 0,
}

var PTKTypeRevMap = map[PTKType]string{
	0: "READER",
}

func (e PTKType) String() string {
	if v, ok := PTKTypeRevMap[e]; ok {
		return v
	}
	return fmt.Sprintf("%v", int(e))
}

type PerTeamSeedCheckVersion int

const (
	PerTeamSeedCheckVersion_V1 PerTeamSeedCheckVersion = 1
)

func (o PerTeamSeedCheckVersion) DeepCopy() PerTeamSeedCheckVersion { return o }

var PerTeamSeedCheckVersionMap = map[string]PerTeamSeedCheckVersion{
	"V1": 1,
}

var PerTeamSeedCheckVersionRevMap = map[PerTeamSeedCheckVersion]string{
	1: "V1",
}

func (e PerTeamSeedCheckVersion) String() string {
	if v, ok := PerTeamSeedCheckVersionRevMap[e]; ok {
		return v
	}
	return fmt.Sprintf("%v", int(e))
}

type PerTeamSeedCheck struct {
	Version PerTeamSeedCheckVersion `codec:"version" json:"version"`
	Value   PerTeamSeedCheckValue   `codec:"value" json:"value"`
}

func (o PerTeamSeedCheck) DeepCopy() PerTeamSeedCheck {
	return PerTeamSeedCheck{
		Version: o.Version.DeepCopy(),
		Value:   o.Value.DeepCopy(),
	}
}

type PerTeamSeedCheckValue []byte

func (o PerTeamSeedCheckValue) DeepCopy() PerTeamSeedCheckValue {
	return (func(x []byte) []byte {
		if x == nil {
			return nil
		}
		return append([]byte{}, x...)
	})(o)
}

type PerTeamSeedCheckValuePostImage []byte

func (o PerTeamSeedCheckValuePostImage) DeepCopy() PerTeamSeedCheckValuePostImage {
	return (func(x []byte) []byte {
		if x == nil {
			return nil
		}
		return append([]byte{}, x...)
	})(o)
}

type PerTeamSeedCheckPostImage struct {
	Value   PerTeamSeedCheckValuePostImage `codec:"h" json:"h"`
	Version PerTeamSeedCheckVersion        `codec:"v" json:"v"`
}

func (o PerTeamSeedCheckPostImage) DeepCopy() PerTeamSeedCheckPostImage {
	return PerTeamSeedCheckPostImage{
		Value:   o.Value.DeepCopy(),
		Version: o.Version.DeepCopy(),
	}
}

type TeamApplicationKey struct {
	Application   TeamApplication      `codec:"application" json:"application"`
	KeyGeneration PerTeamKeyGeneration `codec:"keyGeneration" json:"keyGeneration"`
	Key           Bytes32              `codec:"key" json:"key"`
}

func (o TeamApplicationKey) DeepCopy() TeamApplicationKey {
	return TeamApplicationKey{
		Application:   o.Application.DeepCopy(),
		KeyGeneration: o.KeyGeneration.DeepCopy(),
		Key:           o.Key.DeepCopy(),
	}
}

type MaskB64 []byte

func (o MaskB64) DeepCopy() MaskB64 {
	return (func(x []byte) []byte {
		if x == nil {
			return nil
		}
		return append([]byte{}, x...)
	})(o)
}

type TeamInviteID string

func (o TeamInviteID) DeepCopy() TeamInviteID {
	return o
}

type TeamInviteMaxUses int

func (o TeamInviteMaxUses) DeepCopy() TeamInviteMaxUses {
	return o
}

type ReaderKeyMask struct {
	Application TeamApplication      `codec:"application" json:"application"`
	Generation  PerTeamKeyGeneration `codec:"generation" json:"generation"`
	Mask        MaskB64              `codec:"mask" json:"mask"`
}

func (o ReaderKeyMask) DeepCopy() ReaderKeyMask {
	return ReaderKeyMask{
		Application: o.Application.DeepCopy(),
		Generation:  o.Generation.DeepCopy(),
		Mask:        o.Mask.DeepCopy(),
	}
}

type PerTeamKey struct {
	Gen    PerTeamKeyGeneration `codec:"gen" json:"gen"`
	Seqno  Seqno                `codec:"seqno" json:"seqno"`
	SigKID KID                  `codec:"sigKID" json:"sigKID"`
	EncKID KID                  `codec:"encKID" json:"encKID"`
}

func (o PerTeamKey) DeepCopy() PerTeamKey {
	return PerTeamKey{
		Gen:    o.Gen.DeepCopy(),
		Seqno:  o.Seqno.DeepCopy(),
		SigKID: o.SigKID.DeepCopy(),
		EncKID: o.EncKID.DeepCopy(),
	}
}

type PerTeamKeyAndCheck struct {
	Ptk   PerTeamKey                `codec:"ptk" json:"ptk"`
	Check PerTeamSeedCheckPostImage `codec:"check" json:"check"`
}

func (o PerTeamKeyAndCheck) DeepCopy() PerTeamKeyAndCheck {
	return PerTeamKeyAndCheck{
		Ptk:   o.Ptk.DeepCopy(),
		Check: o.Check.DeepCopy(),
	}
}

type PerTeamKeySeed [32]byte

func (o PerTeamKeySeed) DeepCopy() PerTeamKeySeed {
	var ret PerTeamKeySeed
	copy(ret[:], o[:])
	return ret
}

type PerTeamKeySeedItem struct {
	Seed       PerTeamKeySeed       `codec:"seed" json:"seed"`
	Generation PerTeamKeyGeneration `codec:"generation" json:"generation"`
	Seqno      Seqno                `codec:"seqno" json:"seqno"`
	Check      *PerTeamSeedCheck    `codec:"check,omitempty" json:"check,omitempty"`
}

func (o PerTeamKeySeedItem) DeepCopy() PerTeamKeySeedItem {
	return PerTeamKeySeedItem{
		Seed:       o.Seed.DeepCopy(),
		Generation: o.Generation.DeepCopy(),
		Seqno:      o.Seqno.DeepCopy(),
		Check: (func(x *PerTeamSeedCheck) *PerTeamSeedCheck {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Check),
	}
}

type TeamMember struct {
	Uid         UID              `codec:"uid" json:"uid"`
	Role        TeamRole         `codec:"role" json:"role"`
	EldestSeqno Seqno            `codec:"eldestSeqno" json:"eldestSeqno"`
	Status      TeamMemberStatus `codec:"status" json:"status"`
	BotSettings *TeamBotSettings `codec:"botSettings,omitempty" json:"botSettings,omitempty"`
}

func (o TeamMember) DeepCopy() TeamMember {
	return TeamMember{
		Uid:         o.Uid.DeepCopy(),
		Role:        o.Role.DeepCopy(),
		EldestSeqno: o.EldestSeqno.DeepCopy(),
		Status:      o.Status.DeepCopy(),
		BotSettings: (func(x *TeamBotSettings) *TeamBotSettings {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.BotSettings),
	}
}

type TeamMembers struct {
	Owners         []UserVersion `codec:"owners" json:"owners"`
	Admins         []UserVersion `codec:"admins" json:"admins"`
	Writers        []UserVersion `codec:"writers" json:"writers"`
	Readers        []UserVersion `codec:"readers" json:"readers"`
	Bots           []UserVersion `codec:"bots" json:"bots"`
	RestrictedBots []UserVersion `codec:"restrictedBots" json:"restrictedBots"`
}

func (o TeamMembers) DeepCopy() TeamMembers {
	return TeamMembers{
		Owners: (func(x []UserVersion) []UserVersion {
			if x == nil {
				return nil
			}
			ret := make([]UserVersion, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Owners),
		Admins: (func(x []UserVersion) []UserVersion {
			if x == nil {
				return nil
			}
			ret := make([]UserVersion, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Admins),
		Writers: (func(x []UserVersion) []UserVersion {
			if x == nil {
				return nil
			}
			ret := make([]UserVersion, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Writers),
		Readers: (func(x []UserVersion) []UserVersion {
			if x == nil {
				return nil
			}
			ret := make([]UserVersion, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Readers),
		Bots: (func(x []UserVersion) []UserVersion {
			if x == nil {
				return nil
			}
			ret := make([]UserVersion, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Bots),
		RestrictedBots: (func(x []UserVersion) []UserVersion {
			if x == nil {
				return nil
			}
			ret := make([]UserVersion, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.RestrictedBots),
	}
}

type TeamMemberStatus int

const (
	TeamMemberStatus_ACTIVE  TeamMemberStatus = 0
	TeamMemberStatus_RESET   TeamMemberStatus = 1
	TeamMemberStatus_DELETED TeamMemberStatus = 2
)

func (o TeamMemberStatus) DeepCopy() TeamMemberStatus { return o }

var TeamMemberStatusMap = map[string]TeamMemberStatus{
	"ACTIVE":  0,
	"RESET":   1,
	"DELETED": 2,
}

var TeamMemberStatusRevMap = map[TeamMemberStatus]string{
	0: "ACTIVE",
	1: "RESET",
	2: "DELETED",
}

func (e TeamMemberStatus) String() string {
	if v, ok := TeamMemberStatusRevMap[e]; ok {
		return v
	}
	return fmt.Sprintf("%v", int(e))
}

type TeamMemberDetails struct {
	Uv       UserVersion      `codec:"uv" json:"uv"`
	Username string           `codec:"username" json:"username"`
	FullName FullName         `codec:"fullName" json:"fullName"`
	NeedsPUK bool             `codec:"needsPUK" json:"needsPUK"`
	Status   TeamMemberStatus `codec:"status" json:"status"`
	JoinTime *Time            `codec:"joinTime,omitempty" json:"joinTime,omitempty"`
	Role     TeamRole         `codec:"role" json:"role"`
}

func (o TeamMemberDetails) DeepCopy() TeamMemberDetails {
	return TeamMemberDetails{
		Uv:       o.Uv.DeepCopy(),
		Username: o.Username,
		FullName: o.FullName.DeepCopy(),
		NeedsPUK: o.NeedsPUK,
		Status:   o.Status.DeepCopy(),
		JoinTime: (func(x *Time) *Time {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.JoinTime),
		Role: o.Role.DeepCopy(),
	}
}

type TeamMembersDetails struct {
	Owners         []TeamMemberDetails `codec:"owners" json:"owners"`
	Admins         []TeamMemberDetails `codec:"admins" json:"admins"`
	Writers        []TeamMemberDetails `codec:"writers" json:"writers"`
	Readers        []TeamMemberDetails `codec:"readers" json:"readers"`
	Bots           []TeamMemberDetails `codec:"bots" json:"bots"`
	RestrictedBots []TeamMemberDetails `codec:"restrictedBots" json:"restrictedBots"`
}

func (o TeamMembersDetails) DeepCopy() TeamMembersDetails {
	return TeamMembersDetails{
		Owners: (func(x []TeamMemberDetails) []TeamMemberDetails {
			if x == nil {
				return nil
			}
			ret := make([]TeamMemberDetails, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Owners),
		Admins: (func(x []TeamMemberDetails) []TeamMemberDetails {
			if x == nil {
				return nil
			}
			ret := make([]TeamMemberDetails, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Admins),
		Writers: (func(x []TeamMemberDetails) []TeamMemberDetails {
			if x == nil {
				return nil
			}
			ret := make([]TeamMemberDetails, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Writers),
		Readers: (func(x []TeamMemberDetails) []TeamMemberDetails {
			if x == nil {
				return nil
			}
			ret := make([]TeamMemberDetails, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Readers),
		Bots: (func(x []TeamMemberDetails) []TeamMemberDetails {
			if x == nil {
				return nil
			}
			ret := make([]TeamMemberDetails, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Bots),
		RestrictedBots: (func(x []TeamMemberDetails) []TeamMemberDetails {
			if x == nil {
				return nil
			}
			ret := make([]TeamMemberDetails, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.RestrictedBots),
	}
}

type TeamDetails struct {
	Name                   string                               `codec:"name" json:"name"`
	Members                TeamMembersDetails                   `codec:"members" json:"members"`
	KeyGeneration          PerTeamKeyGeneration                 `codec:"keyGeneration" json:"keyGeneration"`
	AnnotatedActiveInvites map[TeamInviteID]AnnotatedTeamInvite `codec:"annotatedActiveInvites" json:"annotatedActiveInvites"`
	Settings               TeamSettings                         `codec:"settings" json:"settings"`
	Showcase               TeamShowcase                         `codec:"showcase" json:"showcase"`
}

func (o TeamDetails) DeepCopy() TeamDetails {
	return TeamDetails{
		Name:          o.Name,
		Members:       o.Members.DeepCopy(),
		KeyGeneration: o.KeyGeneration.DeepCopy(),
		AnnotatedActiveInvites: (func(x map[TeamInviteID]AnnotatedTeamInvite) map[TeamInviteID]AnnotatedTeamInvite {
			if x == nil {
				return nil
			}
			ret := make(map[TeamInviteID]AnnotatedTeamInvite, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.AnnotatedActiveInvites),
		Settings: o.Settings.DeepCopy(),
		Showcase: o.Showcase.DeepCopy(),
	}
}

type TeamMemberRole struct {
	Uid      UID      `codec:"uid" json:"uid"`
	Username string   `codec:"username" json:"username"`
	FullName FullName `codec:"fullName" json:"fullName"`
	Role     TeamRole `codec:"role" json:"role"`
}

func (o TeamMemberRole) DeepCopy() TeamMemberRole {
	return TeamMemberRole{
		Uid:      o.Uid.DeepCopy(),
		Username: o.Username,
		FullName: o.FullName.DeepCopy(),
		Role:     o.Role.DeepCopy(),
	}
}

type UntrustedTeamInfo struct {
	Name          TeamName         `codec:"name" json:"name"`
	InTeam        bool             `codec:"inTeam" json:"inTeam"`
	Open          bool             `codec:"open" json:"open"`
	Description   string           `codec:"description" json:"description"`
	PublicAdmins  []string         `codec:"publicAdmins" json:"publicAdmins"`
	NumMembers    int              `codec:"numMembers" json:"numMembers"`
	PublicMembers []TeamMemberRole `codec:"publicMembers" json:"publicMembers"`
}

func (o UntrustedTeamInfo) DeepCopy() UntrustedTeamInfo {
	return UntrustedTeamInfo{
		Name:        o.Name.DeepCopy(),
		InTeam:      o.InTeam,
		Open:        o.Open,
		Description: o.Description,
		PublicAdmins: (func(x []string) []string {
			if x == nil {
				return nil
			}
			ret := make([]string, len(x))
			for i, v := range x {
				vCopy := v
				ret[i] = vCopy
			}
			return ret
		})(o.PublicAdmins),
		NumMembers: o.NumMembers,
		PublicMembers: (func(x []TeamMemberRole) []TeamMemberRole {
			if x == nil {
				return nil
			}
			ret := make([]TeamMemberRole, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.PublicMembers),
	}
}

type UserVersionPercentForm string

func (o UserVersionPercentForm) DeepCopy() UserVersionPercentForm {
	return o
}

type TeamUsedInvite struct {
	InviteID TeamInviteID           `codec:"inviteID" json:"inviteID"`
	Uv       UserVersionPercentForm `codec:"uv" json:"uv"`
}

func (o TeamUsedInvite) DeepCopy() TeamUsedInvite {
	return TeamUsedInvite{
		InviteID: o.InviteID.DeepCopy(),
		Uv:       o.Uv.DeepCopy(),
	}
}

type TeamChangeReq struct {
	Owners           []UserVersion                           `codec:"owners" json:"owners"`
	Admins           []UserVersion                           `codec:"admins" json:"admins"`
	Writers          []UserVersion                           `codec:"writers" json:"writers"`
	Readers          []UserVersion                           `codec:"readers" json:"readers"`
	Bots             []UserVersion                           `codec:"bots" json:"bots"`
	RestrictedBots   map[UserVersion]TeamBotSettings         `codec:"restrictedBots" json:"restrictedBots"`
	None             []UserVersion                           `codec:"none" json:"none"`
	CompletedInvites map[TeamInviteID]UserVersionPercentForm `codec:"completedInvites" json:"completedInvites"`
	UsedInvites      []TeamUsedInvite                        `codec:"usedInvites" json:"usedInvites"`
}

func (o TeamChangeReq) DeepCopy() TeamChangeReq {
	return TeamChangeReq{
		Owners: (func(x []UserVersion) []UserVersion {
			if x == nil {
				return nil
			}
			ret := make([]UserVersion, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Owners),
		Admins: (func(x []UserVersion) []UserVersion {
			if x == nil {
				return nil
			}
			ret := make([]UserVersion, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Admins),
		Writers: (func(x []UserVersion) []UserVersion {
			if x == nil {
				return nil
			}
			ret := make([]UserVersion, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Writers),
		Readers: (func(x []UserVersion) []UserVersion {
			if x == nil {
				return nil
			}
			ret := make([]UserVersion, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Readers),
		Bots: (func(x []UserVersion) []UserVersion {
			if x == nil {
				return nil
			}
			ret := make([]UserVersion, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Bots),
		RestrictedBots: (func(x map[UserVersion]TeamBotSettings) map[UserVersion]TeamBotSettings {
			if x == nil {
				return nil
			}
			ret := make(map[UserVersion]TeamBotSettings, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.RestrictedBots),
		None: (func(x []UserVersion) []UserVersion {
			if x == nil {
				return nil
			}
			ret := make([]UserVersion, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.None),
		CompletedInvites: (func(x map[TeamInviteID]UserVersionPercentForm) map[TeamInviteID]UserVersionPercentForm {
			if x == nil {
				return nil
			}
			ret := make(map[TeamInviteID]UserVersionPercentForm, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.CompletedInvites),
		UsedInvites: (func(x []TeamUsedInvite) []TeamUsedInvite {
			if x == nil {
				return nil
			}
			ret := make([]TeamUsedInvite, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.UsedInvites),
	}
}

type TeamPlusApplicationKeys struct {
	Id                 TeamID               `codec:"id" json:"id"`
	Name               string               `codec:"name" json:"name"`
	Implicit           bool                 `codec:"implicit" json:"implicit"`
	Public             bool                 `codec:"public" json:"public"`
	Application        TeamApplication      `codec:"application" json:"application"`
	Writers            []UserVersion        `codec:"writers" json:"writers"`
	OnlyReaders        []UserVersion        `codec:"onlyReaders" json:"onlyReaders"`
	OnlyRestrictedBots []UserVersion        `codec:"onlyRestrictedBots" json:"onlyRestrictedBots"`
	ApplicationKeys    []TeamApplicationKey `codec:"applicationKeys" json:"applicationKeys"`
}

func (o TeamPlusApplicationKeys) DeepCopy() TeamPlusApplicationKeys {
	return TeamPlusApplicationKeys{
		Id:          o.Id.DeepCopy(),
		Name:        o.Name,
		Implicit:    o.Implicit,
		Public:      o.Public,
		Application: o.Application.DeepCopy(),
		Writers: (func(x []UserVersion) []UserVersion {
			if x == nil {
				return nil
			}
			ret := make([]UserVersion, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Writers),
		OnlyReaders: (func(x []UserVersion) []UserVersion {
			if x == nil {
				return nil
			}
			ret := make([]UserVersion, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.OnlyReaders),
		OnlyRestrictedBots: (func(x []UserVersion) []UserVersion {
			if x == nil {
				return nil
			}
			ret := make([]UserVersion, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.OnlyRestrictedBots),
		ApplicationKeys: (func(x []TeamApplicationKey) []TeamApplicationKey {
			if x == nil {
				return nil
			}
			ret := make([]TeamApplicationKey, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.ApplicationKeys),
	}
}

type TeamData struct {
	Subversion                int                                                  `codec:"v" json:"v"`
	Frozen                    bool                                                 `codec:"frozen" json:"frozen"`
	Tombstoned                bool                                                 `codec:"tombstoned" json:"tombstoned"`
	Secretless                bool                                                 `codec:"secretless" json:"secretless"`
	Name                      TeamName                                             `codec:"name" json:"name"`
	Chain                     TeamSigChainState                                    `codec:"chain" json:"chain"`
	PerTeamKeySeedsUnverified map[PerTeamKeyGeneration]PerTeamKeySeedItem          `codec:"perTeamKeySeeds" json:"perTeamKeySeedsUnverified"`
	ReaderKeyMasks            map[TeamApplication]map[PerTeamKeyGeneration]MaskB64 `codec:"readerKeyMasks" json:"readerKeyMasks"`
	LatestSeqnoHint           Seqno                                                `codec:"latestSeqnoHint" json:"latestSeqnoHint"`
	CachedAt                  Time                                                 `codec:"cachedAt" json:"cachedAt"`
	TlfCryptKeys              map[TeamApplication][]CryptKey                       `codec:"tlfCryptKeys" json:"tlfCryptKeys"`
}

func (o TeamData) DeepCopy() TeamData {
	return TeamData{
		Subversion: o.Subversion,
		Frozen:     o.Frozen,
		Tombstoned: o.Tombstoned,
		Secretless: o.Secretless,
		Name:       o.Name.DeepCopy(),
		Chain:      o.Chain.DeepCopy(),
		PerTeamKeySeedsUnverified: (func(x map[PerTeamKeyGeneration]PerTeamKeySeedItem) map[PerTeamKeyGeneration]PerTeamKeySeedItem {
			if x == nil {
				return nil
			}
			ret := make(map[PerTeamKeyGeneration]PerTeamKeySeedItem, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.PerTeamKeySeedsUnverified),
		ReaderKeyMasks: (func(x map[TeamApplication]map[PerTeamKeyGeneration]MaskB64) map[TeamApplication]map[PerTeamKeyGeneration]MaskB64 {
			if x == nil {
				return nil
			}
			ret := make(map[TeamApplication]map[PerTeamKeyGeneration]MaskB64, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := (func(x map[PerTeamKeyGeneration]MaskB64) map[PerTeamKeyGeneration]MaskB64 {
					if x == nil {
						return nil
					}
					ret := make(map[PerTeamKeyGeneration]MaskB64, len(x))
					for k, v := range x {
						kCopy := k.DeepCopy()
						vCopy := v.DeepCopy()
						ret[kCopy] = vCopy
					}
					return ret
				})(v)
				ret[kCopy] = vCopy
			}
			return ret
		})(o.ReaderKeyMasks),
		LatestSeqnoHint: o.LatestSeqnoHint.DeepCopy(),
		CachedAt:        o.CachedAt.DeepCopy(),
		TlfCryptKeys: (func(x map[TeamApplication][]CryptKey) map[TeamApplication][]CryptKey {
			if x == nil {
				return nil
			}
			ret := make(map[TeamApplication][]CryptKey, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := (func(x []CryptKey) []CryptKey {
					if x == nil {
						return nil
					}
					ret := make([]CryptKey, len(x))
					for i, v := range x {
						vCopy := v.DeepCopy()
						ret[i] = vCopy
					}
					return ret
				})(v)
				ret[kCopy] = vCopy
			}
			return ret
		})(o.TlfCryptKeys),
	}
}

type FastTeamData struct {
	Frozen                     bool                                                 `codec:"frozen" json:"frozen"`
	Subversion                 int                                                  `codec:"subversion" json:"subversion"`
	Tombstoned                 bool                                                 `codec:"tombstoned" json:"tombstoned"`
	Name                       TeamName                                             `codec:"name" json:"name"`
	Chain                      FastTeamSigChainState                                `codec:"chain" json:"chain"`
	PerTeamKeySeedsUnverified  map[PerTeamKeyGeneration]PerTeamKeySeed              `codec:"perTeamKeySeeds" json:"perTeamKeySeedsUnverified"`
	MaxContinuousPTKGeneration PerTeamKeyGeneration                                 `codec:"maxContinuousPTKGeneration" json:"maxContinuousPTKGeneration"`
	SeedChecks                 map[PerTeamKeyGeneration]PerTeamSeedCheck            `codec:"seedChecks" json:"seedChecks"`
	LatestKeyGeneration        PerTeamKeyGeneration                                 `codec:"latestKeyGeneration" json:"latestKeyGeneration"`
	ReaderKeyMasks             map[TeamApplication]map[PerTeamKeyGeneration]MaskB64 `codec:"readerKeyMasks" json:"readerKeyMasks"`
	LatestSeqnoHint            Seqno                                                `codec:"latestSeqnoHint" json:"latestSeqnoHint"`
	CachedAt                   Time                                                 `codec:"cachedAt" json:"cachedAt"`
	LoadedLatest               bool                                                 `codec:"loadedLatest" json:"loadedLatest"`
}

func (o FastTeamData) DeepCopy() FastTeamData {
	return FastTeamData{
		Frozen:     o.Frozen,
		Subversion: o.Subversion,
		Tombstoned: o.Tombstoned,
		Name:       o.Name.DeepCopy(),
		Chain:      o.Chain.DeepCopy(),
		PerTeamKeySeedsUnverified: (func(x map[PerTeamKeyGeneration]PerTeamKeySeed) map[PerTeamKeyGeneration]PerTeamKeySeed {
			if x == nil {
				return nil
			}
			ret := make(map[PerTeamKeyGeneration]PerTeamKeySeed, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.PerTeamKeySeedsUnverified),
		MaxContinuousPTKGeneration: o.MaxContinuousPTKGeneration.DeepCopy(),
		SeedChecks: (func(x map[PerTeamKeyGeneration]PerTeamSeedCheck) map[PerTeamKeyGeneration]PerTeamSeedCheck {
			if x == nil {
				return nil
			}
			ret := make(map[PerTeamKeyGeneration]PerTeamSeedCheck, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.SeedChecks),
		LatestKeyGeneration: o.LatestKeyGeneration.DeepCopy(),
		ReaderKeyMasks: (func(x map[TeamApplication]map[PerTeamKeyGeneration]MaskB64) map[TeamApplication]map[PerTeamKeyGeneration]MaskB64 {
			if x == nil {
				return nil
			}
			ret := make(map[TeamApplication]map[PerTeamKeyGeneration]MaskB64, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := (func(x map[PerTeamKeyGeneration]MaskB64) map[PerTeamKeyGeneration]MaskB64 {
					if x == nil {
						return nil
					}
					ret := make(map[PerTeamKeyGeneration]MaskB64, len(x))
					for k, v := range x {
						kCopy := k.DeepCopy()
						vCopy := v.DeepCopy()
						ret[kCopy] = vCopy
					}
					return ret
				})(v)
				ret[kCopy] = vCopy
			}
			return ret
		})(o.ReaderKeyMasks),
		LatestSeqnoHint: o.LatestSeqnoHint.DeepCopy(),
		CachedAt:        o.CachedAt.DeepCopy(),
		LoadedLatest:    o.LoadedLatest,
	}
}

type RatchetType int

const (
	RatchetType_MAIN        RatchetType = 0
	RatchetType_BLINDED     RatchetType = 1
	RatchetType_SELF        RatchetType = 2
	RatchetType_UNCOMMITTED RatchetType = 3
)

func (o RatchetType) DeepCopy() RatchetType { return o }

var RatchetTypeMap = map[string]RatchetType{
	"MAIN":        0,
	"BLINDED":     1,
	"SELF":        2,
	"UNCOMMITTED": 3,
}

var RatchetTypeRevMap = map[RatchetType]string{
	0: "MAIN",
	1: "BLINDED",
	2: "SELF",
	3: "UNCOMMITTED",
}

func (e RatchetType) String() string {
	if v, ok := RatchetTypeRevMap[e]; ok {
		return v
	}
	return fmt.Sprintf("%v", int(e))
}

type HiddenTeamChainRatchetSet struct {
	Ratchets map[RatchetType]LinkTripleAndTime `codec:"ratchets" json:"ratchets"`
}

func (o HiddenTeamChainRatchetSet) DeepCopy() HiddenTeamChainRatchetSet {
	return HiddenTeamChainRatchetSet{
		Ratchets: (func(x map[RatchetType]LinkTripleAndTime) map[RatchetType]LinkTripleAndTime {
			if x == nil {
				return nil
			}
			ret := make(map[RatchetType]LinkTripleAndTime, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.Ratchets),
	}
}

type HiddenTeamChain struct {
	Id                 TeamID                         `codec:"id" json:"id"`
	Subversion         int                            `codec:"subversion" json:"subversion"`
	Public             bool                           `codec:"public" json:"public"`
	Frozen             bool                           `codec:"frozen" json:"frozen"`
	Tombstoned         bool                           `codec:"tombstoned" json:"tombstoned"`
	Last               Seqno                          `codec:"last" json:"last"`
	LastFull           Seqno                          `codec:"lastFull" json:"lastFull"`
	LatestSeqnoHint    Seqno                          `codec:"latestSeqnoHint" json:"latestSeqnoHint"`
	LastCommittedSeqno Seqno                          `codec:"lastCommittedSeqno" json:"lastCommittedSeqno"`
	LinkReceiptTimes   map[Seqno]Time                 `codec:"linkReceiptTimes" json:"linkReceiptTimes"`
	LastPerTeamKeys    map[PTKType]Seqno              `codec:"lastPerTeamKeys" json:"lastPerTeamKeys"`
	Outer              map[Seqno]LinkID               `codec:"outer" json:"outer"`
	Inner              map[Seqno]HiddenTeamChainLink  `codec:"inner" json:"inner"`
	ReaderPerTeamKeys  map[PerTeamKeyGeneration]Seqno `codec:"readerPerTeamKeys" json:"readerPerTeamKeys"`
	RatchetSet         HiddenTeamChainRatchetSet      `codec:"ratchetSet" json:"ratchetSet"`
	CachedAt           Time                           `codec:"cachedAt" json:"cachedAt"`
	NeedRotate         bool                           `codec:"needRotate" json:"needRotate"`
	MerkleRoots        map[Seqno]MerkleRootV2         `codec:"merkleRoots" json:"merkleRoots"`
}

func (o HiddenTeamChain) DeepCopy() HiddenTeamChain {
	return HiddenTeamChain{
		Id:                 o.Id.DeepCopy(),
		Subversion:         o.Subversion,
		Public:             o.Public,
		Frozen:             o.Frozen,
		Tombstoned:         o.Tombstoned,
		Last:               o.Last.DeepCopy(),
		LastFull:           o.LastFull.DeepCopy(),
		LatestSeqnoHint:    o.LatestSeqnoHint.DeepCopy(),
		LastCommittedSeqno: o.LastCommittedSeqno.DeepCopy(),
		LinkReceiptTimes: (func(x map[Seqno]Time) map[Seqno]Time {
			if x == nil {
				return nil
			}
			ret := make(map[Seqno]Time, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.LinkReceiptTimes),
		LastPerTeamKeys: (func(x map[PTKType]Seqno) map[PTKType]Seqno {
			if x == nil {
				return nil
			}
			ret := make(map[PTKType]Seqno, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.LastPerTeamKeys),
		Outer: (func(x map[Seqno]LinkID) map[Seqno]LinkID {
			if x == nil {
				return nil
			}
			ret := make(map[Seqno]LinkID, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.Outer),
		Inner: (func(x map[Seqno]HiddenTeamChainLink) map[Seqno]HiddenTeamChainLink {
			if x == nil {
				return nil
			}
			ret := make(map[Seqno]HiddenTeamChainLink, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.Inner),
		ReaderPerTeamKeys: (func(x map[PerTeamKeyGeneration]Seqno) map[PerTeamKeyGeneration]Seqno {
			if x == nil {
				return nil
			}
			ret := make(map[PerTeamKeyGeneration]Seqno, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.ReaderPerTeamKeys),
		RatchetSet: o.RatchetSet.DeepCopy(),
		CachedAt:   o.CachedAt.DeepCopy(),
		NeedRotate: o.NeedRotate,
		MerkleRoots: (func(x map[Seqno]MerkleRootV2) map[Seqno]MerkleRootV2 {
			if x == nil {
				return nil
			}
			ret := make(map[Seqno]MerkleRootV2, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.MerkleRoots),
	}
}

type LinkTriple struct {
	Seqno   Seqno   `codec:"seqno" json:"seqno"`
	SeqType SeqType `codec:"seqType" json:"seqType"`
	LinkID  LinkID  `codec:"linkID" json:"linkID"`
}

func (o LinkTriple) DeepCopy() LinkTriple {
	return LinkTriple{
		Seqno:   o.Seqno.DeepCopy(),
		SeqType: o.SeqType.DeepCopy(),
		LinkID:  o.LinkID.DeepCopy(),
	}
}

type LinkTripleAndTime struct {
	Triple LinkTriple `codec:"triple" json:"triple"`
	Time   Time       `codec:"time" json:"time"`
}

func (o LinkTripleAndTime) DeepCopy() LinkTripleAndTime {
	return LinkTripleAndTime{
		Triple: o.Triple.DeepCopy(),
		Time:   o.Time.DeepCopy(),
	}
}

type UpPointer struct {
	OurSeqno    Seqno  `codec:"ourSeqno" json:"ourSeqno"`
	ParentID    TeamID `codec:"parentID" json:"parentID"`
	ParentSeqno Seqno  `codec:"parentSeqno" json:"parentSeqno"`
	Deletion    bool   `codec:"deletion" json:"deletion"`
}

func (o UpPointer) DeepCopy() UpPointer {
	return UpPointer{
		OurSeqno:    o.OurSeqno.DeepCopy(),
		ParentID:    o.ParentID.DeepCopy(),
		ParentSeqno: o.ParentSeqno.DeepCopy(),
		Deletion:    o.Deletion,
	}
}

type DownPointer struct {
	Id            TeamID `codec:"id" json:"id"`
	NameComponent string `codec:"nameComponent" json:"nameComponent"`
	IsDeleted     bool   `codec:"isDeleted" json:"isDeleted"`
}

func (o DownPointer) DeepCopy() DownPointer {
	return DownPointer{
		Id:            o.Id.DeepCopy(),
		NameComponent: o.NameComponent,
		IsDeleted:     o.IsDeleted,
	}
}

type Signer struct {
	E Seqno `codec:"e" json:"e"`
	K KID   `codec:"k" json:"k"`
	U UID   `codec:"u" json:"u"`
}

func (o Signer) DeepCopy() Signer {
	return Signer{
		E: o.E.DeepCopy(),
		K: o.K.DeepCopy(),
		U: o.U.DeepCopy(),
	}
}

type HiddenTeamChainLink struct {
	MerkleRoot  MerkleRootV2                   `codec:"m" json:"m"`
	ParentChain LinkTriple                     `codec:"p" json:"p"`
	Signer      Signer                         `codec:"s" json:"s"`
	Ptk         map[PTKType]PerTeamKeyAndCheck `codec:"k" json:"k"`
}

func (o HiddenTeamChainLink) DeepCopy() HiddenTeamChainLink {
	return HiddenTeamChainLink{
		MerkleRoot:  o.MerkleRoot.DeepCopy(),
		ParentChain: o.ParentChain.DeepCopy(),
		Signer:      o.Signer.DeepCopy(),
		Ptk: (func(x map[PTKType]PerTeamKeyAndCheck) map[PTKType]PerTeamKeyAndCheck {
			if x == nil {
				return nil
			}
			ret := make(map[PTKType]PerTeamKeyAndCheck, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.Ptk),
	}
}

type FastTeamSigChainState struct {
	ID                      TeamID                                  `codec:"ID" json:"ID"`
	Public                  bool                                    `codec:"public" json:"public"`
	RootAncestor            TeamName                                `codec:"rootAncestor" json:"rootAncestor"`
	NameDepth               int                                     `codec:"nameDepth" json:"nameDepth"`
	Last                    *LinkTriple                             `codec:"last,omitempty" json:"last,omitempty"`
	PerTeamKeys             map[PerTeamKeyGeneration]PerTeamKey     `codec:"perTeamKeys" json:"perTeamKeys"`
	PerTeamKeySeedsVerified map[PerTeamKeyGeneration]PerTeamKeySeed `codec:"perTeamKeySeedsVerified" json:"perTeamKeySeedsVerified"`
	DownPointers            map[Seqno]DownPointer                   `codec:"downPointers" json:"downPointers"`
	LastUpPointer           *UpPointer                              `codec:"lastUpPointer,omitempty" json:"lastUpPointer,omitempty"`
	PerTeamKeyCTime         UnixTime                                `codec:"perTeamKeyCTime" json:"perTeamKeyCTime"`
	LinkIDs                 map[Seqno]LinkID                        `codec:"linkIDs" json:"linkIDs"`
	MerkleInfo              map[Seqno]MerkleRootV2                  `codec:"merkleInfo" json:"merkleInfo"`
}

func (o FastTeamSigChainState) DeepCopy() FastTeamSigChainState {
	return FastTeamSigChainState{
		ID:           o.ID.DeepCopy(),
		Public:       o.Public,
		RootAncestor: o.RootAncestor.DeepCopy(),
		NameDepth:    o.NameDepth,
		Last: (func(x *LinkTriple) *LinkTriple {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Last),
		PerTeamKeys: (func(x map[PerTeamKeyGeneration]PerTeamKey) map[PerTeamKeyGeneration]PerTeamKey {
			if x == nil {
				return nil
			}
			ret := make(map[PerTeamKeyGeneration]PerTeamKey, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.PerTeamKeys),
		PerTeamKeySeedsVerified: (func(x map[PerTeamKeyGeneration]PerTeamKeySeed) map[PerTeamKeyGeneration]PerTeamKeySeed {
			if x == nil {
				return nil
			}
			ret := make(map[PerTeamKeyGeneration]PerTeamKeySeed, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.PerTeamKeySeedsVerified),
		DownPointers: (func(x map[Seqno]DownPointer) map[Seqno]DownPointer {
			if x == nil {
				return nil
			}
			ret := make(map[Seqno]DownPointer, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.DownPointers),
		LastUpPointer: (func(x *UpPointer) *UpPointer {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.LastUpPointer),
		PerTeamKeyCTime: o.PerTeamKeyCTime.DeepCopy(),
		LinkIDs: (func(x map[Seqno]LinkID) map[Seqno]LinkID {
			if x == nil {
				return nil
			}
			ret := make(map[Seqno]LinkID, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.LinkIDs),
		MerkleInfo: (func(x map[Seqno]MerkleRootV2) map[Seqno]MerkleRootV2 {
			if x == nil {
				return nil
			}
			ret := make(map[Seqno]MerkleRootV2, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.MerkleInfo),
	}
}

type Audit struct {
	Time           Time  `codec:"time" json:"time"`
	MaxMerkleSeqno Seqno `codec:"mms" json:"mms"`
	MaxChainSeqno  Seqno `codec:"mcs" json:"mcs"`
	MaxHiddenSeqno Seqno `codec:"mhs" json:"mhs"`
	MaxMerkleProbe Seqno `codec:"mmp" json:"mmp"`
}

func (o Audit) DeepCopy() Audit {
	return Audit{
		Time:           o.Time.DeepCopy(),
		MaxMerkleSeqno: o.MaxMerkleSeqno.DeepCopy(),
		MaxChainSeqno:  o.MaxChainSeqno.DeepCopy(),
		MaxHiddenSeqno: o.MaxHiddenSeqno.DeepCopy(),
		MaxMerkleProbe: o.MaxMerkleProbe.DeepCopy(),
	}
}

type Probe struct {
	Index           int   `codec:"i" json:"i"`
	TeamSeqno       Seqno `codec:"s" json:"t"`
	TeamHiddenSeqno Seqno `codec:"h" json:"h"`
}

func (o Probe) DeepCopy() Probe {
	return Probe{
		Index:           o.Index,
		TeamSeqno:       o.TeamSeqno.DeepCopy(),
		TeamHiddenSeqno: o.TeamHiddenSeqno.DeepCopy(),
	}
}

type AuditVersion int

const (
	AuditVersion_V0 AuditVersion = 0
	AuditVersion_V1 AuditVersion = 1
	AuditVersion_V2 AuditVersion = 2
	AuditVersion_V3 AuditVersion = 3
	AuditVersion_V4 AuditVersion = 4
)

func (o AuditVersion) DeepCopy() AuditVersion { return o }

var AuditVersionMap = map[string]AuditVersion{
	"V0": 0,
	"V1": 1,
	"V2": 2,
	"V3": 3,
	"V4": 4,
}

var AuditVersionRevMap = map[AuditVersion]string{
	0: "V0",
	1: "V1",
	2: "V2",
	3: "V3",
	4: "V4",
}

func (e AuditVersion) String() string {
	if v, ok := AuditVersionRevMap[e]; ok {
		return v
	}
	return fmt.Sprintf("%v", int(e))
}

type AuditHistory struct {
	ID                TeamID           `codec:"ID" json:"ID"`
	Public            bool             `codec:"public" json:"public"`
	PriorMerkleSeqno  Seqno            `codec:"priorMerkleSeqno" json:"priorMerkleSeqno"`
	Version           AuditVersion     `codec:"version" json:"version"`
	Audits            []Audit          `codec:"audits" json:"audits"`
	PreProbes         map[Seqno]Probe  `codec:"preProbes" json:"preProbes"`
	PostProbes        map[Seqno]Probe  `codec:"postProbes" json:"postProbes"`
	Tails             map[Seqno]LinkID `codec:"tails" json:"tails"`
	HiddenTails       map[Seqno]LinkID `codec:"hiddenTails" json:"hiddenTails"`
	PreProbesToRetry  []Seqno          `codec:"preProbesToRetry" json:"preProbesToRetry"`
	PostProbesToRetry []Seqno          `codec:"postProbesToRetry" json:"postProbesToRetry"`
	SkipUntil         Time             `codec:"skipUntil" json:"skipUntil"`
}

func (o AuditHistory) DeepCopy() AuditHistory {
	return AuditHistory{
		ID:               o.ID.DeepCopy(),
		Public:           o.Public,
		PriorMerkleSeqno: o.PriorMerkleSeqno.DeepCopy(),
		Version:          o.Version.DeepCopy(),
		Audits: (func(x []Audit) []Audit {
			if x == nil {
				return nil
			}
			ret := make([]Audit, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Audits),
		PreProbes: (func(x map[Seqno]Probe) map[Seqno]Probe {
			if x == nil {
				return nil
			}
			ret := make(map[Seqno]Probe, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.PreProbes),
		PostProbes: (func(x map[Seqno]Probe) map[Seqno]Probe {
			if x == nil {
				return nil
			}
			ret := make(map[Seqno]Probe, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.PostProbes),
		Tails: (func(x map[Seqno]LinkID) map[Seqno]LinkID {
			if x == nil {
				return nil
			}
			ret := make(map[Seqno]LinkID, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.Tails),
		HiddenTails: (func(x map[Seqno]LinkID) map[Seqno]LinkID {
			if x == nil {
				return nil
			}
			ret := make(map[Seqno]LinkID, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.HiddenTails),
		PreProbesToRetry: (func(x []Seqno) []Seqno {
			if x == nil {
				return nil
			}
			ret := make([]Seqno, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.PreProbesToRetry),
		PostProbesToRetry: (func(x []Seqno) []Seqno {
			if x == nil {
				return nil
			}
			ret := make([]Seqno, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.PostProbesToRetry),
		SkipUntil: o.SkipUntil.DeepCopy(),
	}
}

type TeamInviteCategory int

const (
	TeamInviteCategory_NONE       TeamInviteCategory = 0
	TeamInviteCategory_UNKNOWN    TeamInviteCategory = 1
	TeamInviteCategory_KHULNASOFT    TeamInviteCategory = 2
	TeamInviteCategory_EMAIL      TeamInviteCategory = 3
	TeamInviteCategory_SBS        TeamInviteCategory = 4
	TeamInviteCategory_SEITAN     TeamInviteCategory = 5
	TeamInviteCategory_PHONE      TeamInviteCategory = 6
	TeamInviteCategory_INVITELINK TeamInviteCategory = 7
)

func (o TeamInviteCategory) DeepCopy() TeamInviteCategory { return o }

var TeamInviteCategoryMap = map[string]TeamInviteCategory{
	"NONE":       0,
	"UNKNOWN":    1,
	"KHULNASOFT":    2,
	"EMAIL":      3,
	"SBS":        4,
	"SEITAN":     5,
	"PHONE":      6,
	"INVITELINK": 7,
}

var TeamInviteCategoryRevMap = map[TeamInviteCategory]string{
	0: "NONE",
	1: "UNKNOWN",
	2: "KHULNASOFT",
	3: "EMAIL",
	4: "SBS",
	5: "SEITAN",
	6: "PHONE",
	7: "INVITELINK",
}

func (e TeamInviteCategory) String() string {
	if v, ok := TeamInviteCategoryRevMap[e]; ok {
		return v
	}
	return fmt.Sprintf("%v", int(e))
}

type TeamInviteType struct {
	C__       TeamInviteCategory       `codec:"c" json:"c"`
	Unknown__ *string                  `codec:"unknown,omitempty" json:"unknown,omitempty"`
	Sbs__     *TeamInviteSocialNetwork `codec:"sbs,omitempty" json:"sbs,omitempty"`
}

func (o *TeamInviteType) C() (ret TeamInviteCategory, err error) {
	switch o.C__ {
	case TeamInviteCategory_UNKNOWN:
		if o.Unknown__ == nil {
			err = errors.New("unexpected nil value for Unknown__")
			return ret, err
		}
	case TeamInviteCategory_SBS:
		if o.Sbs__ == nil {
			err = errors.New("unexpected nil value for Sbs__")
			return ret, err
		}
	}
	return o.C__, nil
}

func (o TeamInviteType) Unknown() (res string) {
	if o.C__ != TeamInviteCategory_UNKNOWN {
		panic("wrong case accessed")
	}
	if o.Unknown__ == nil {
		return
	}
	return *o.Unknown__
}

func (o TeamInviteType) Sbs() (res TeamInviteSocialNetwork) {
	if o.C__ != TeamInviteCategory_SBS {
		panic("wrong case accessed")
	}
	if o.Sbs__ == nil {
		return
	}
	return *o.Sbs__
}

func NewTeamInviteTypeWithUnknown(v string) TeamInviteType {
	return TeamInviteType{
		C__:       TeamInviteCategory_UNKNOWN,
		Unknown__: &v,
	}
}

func NewTeamInviteTypeWithSbs(v TeamInviteSocialNetwork) TeamInviteType {
	return TeamInviteType{
		C__:   TeamInviteCategory_SBS,
		Sbs__: &v,
	}
}

func NewTeamInviteTypeDefault(c TeamInviteCategory) TeamInviteType {
	return TeamInviteType{
		C__: c,
	}
}

func (o TeamInviteType) DeepCopy() TeamInviteType {
	return TeamInviteType{
		C__: o.C__.DeepCopy(),
		Unknown__: (func(x *string) *string {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.Unknown__),
		Sbs__: (func(x *TeamInviteSocialNetwork) *TeamInviteSocialNetwork {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Sbs__),
	}
}

type TeamInviteSocialNetwork string

func (o TeamInviteSocialNetwork) DeepCopy() TeamInviteSocialNetwork {
	return o
}

type TeamInviteName string

func (o TeamInviteName) DeepCopy() TeamInviteName {
	return o
}

type TeamInviteDisplayName string

func (o TeamInviteDisplayName) DeepCopy() TeamInviteDisplayName {
	return o
}

type TeamInvite struct {
	Role    TeamRole           `codec:"role" json:"role"`
	Id      TeamInviteID       `codec:"id" json:"id"`
	Type    TeamInviteType     `codec:"type" json:"type"`
	Name    TeamInviteName     `codec:"name" json:"name"`
	Inviter UserVersion        `codec:"inviter" json:"inviter"`
	MaxUses *TeamInviteMaxUses `codec:"maxUses,omitempty" json:"maxUses,omitempty"`
	Etime   *UnixTime          `codec:"etime,omitempty" json:"etime,omitempty"`
}

func (o TeamInvite) DeepCopy() TeamInvite {
	return TeamInvite{
		Role:    o.Role.DeepCopy(),
		Id:      o.Id.DeepCopy(),
		Type:    o.Type.DeepCopy(),
		Name:    o.Name.DeepCopy(),
		Inviter: o.Inviter.DeepCopy(),
		MaxUses: (func(x *TeamInviteMaxUses) *TeamInviteMaxUses {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.MaxUses),
		Etime: (func(x *UnixTime) *UnixTime {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Etime),
	}
}

type AnnotatedTeamInvite struct {
	InviteMetadata      TeamInviteMetadata     `codec:"inviteMetadata" json:"inviteMetadata"`
	DisplayName         TeamInviteDisplayName  `codec:"displayName" json:"displayName"`
	InviterUsername     string                 `codec:"inviterUsername" json:"inviterUsername"`
	TeamName            string                 `codec:"teamName" json:"teamName"`
	IsValid             bool                   `codec:"isValid" json:"isValid"`
	ValidityDescription string                 `codec:"validityDescription" json:"validityDescription"`
	InviteExt           AnnotatedTeamInviteExt `codec:"inviteExt" json:"inviteExt"`
}

func (o AnnotatedTeamInvite) DeepCopy() AnnotatedTeamInvite {
	return AnnotatedTeamInvite{
		InviteMetadata:      o.InviteMetadata.DeepCopy(),
		DisplayName:         o.DisplayName.DeepCopy(),
		InviterUsername:     o.InviterUsername,
		TeamName:            o.TeamName,
		IsValid:             o.IsValid,
		ValidityDescription: o.ValidityDescription,
		InviteExt:           o.InviteExt.DeepCopy(),
	}
}

type KhulnasoftInviteExt struct {
	InviteeUv UserVersion      `codec:"inviteeUv" json:"inviteeUv"`
	Status    TeamMemberStatus `codec:"status" json:"status"`
	FullName  FullName         `codec:"fullName" json:"fullName"`
	Username  string           `codec:"username" json:"username"`
}

func (o KhulnasoftInviteExt) DeepCopy() KhulnasoftInviteExt {
	return KhulnasoftInviteExt{
		InviteeUv: o.InviteeUv.DeepCopy(),
		Status:    o.Status.DeepCopy(),
		FullName:  o.FullName.DeepCopy(),
		Username:  o.Username,
	}
}

type InvitelinkInviteExt struct {
	AnnotatedUsedInvites []AnnotatedTeamUsedInviteLogPoint `codec:"annotatedUsedInvites" json:"annotatedUsedInvites"`
}

func (o InvitelinkInviteExt) DeepCopy() InvitelinkInviteExt {
	return InvitelinkInviteExt{
		AnnotatedUsedInvites: (func(x []AnnotatedTeamUsedInviteLogPoint) []AnnotatedTeamUsedInviteLogPoint {
			if x == nil {
				return nil
			}
			ret := make([]AnnotatedTeamUsedInviteLogPoint, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.AnnotatedUsedInvites),
	}
}

type AnnotatedTeamInviteExt struct {
	C__          TeamInviteCategory   `codec:"c" json:"c"`
	Khulnasoft__    *KhulnasoftInviteExt    `codec:"khulnasoft,omitempty" json:"khulnasoft,omitempty"`
	Invitelink__ *InvitelinkInviteExt `codec:"invitelink,omitempty" json:"invitelink,omitempty"`
}

func (o *AnnotatedTeamInviteExt) C() (ret TeamInviteCategory, err error) {
	switch o.C__ {
	case TeamInviteCategory_KHULNASOFT:
		if o.Khulnasoft__ == nil {
			err = errors.New("unexpected nil value for Khulnasoft__")
			return ret, err
		}
	case TeamInviteCategory_INVITELINK:
		if o.Invitelink__ == nil {
			err = errors.New("unexpected nil value for Invitelink__")
			return ret, err
		}
	}
	return o.C__, nil
}

func (o AnnotatedTeamInviteExt) Khulnasoft() (res KhulnasoftInviteExt) {
	if o.C__ != TeamInviteCategory_KHULNASOFT {
		panic("wrong case accessed")
	}
	if o.Khulnasoft__ == nil {
		return
	}
	return *o.Khulnasoft__
}

func (o AnnotatedTeamInviteExt) Invitelink() (res InvitelinkInviteExt) {
	if o.C__ != TeamInviteCategory_INVITELINK {
		panic("wrong case accessed")
	}
	if o.Invitelink__ == nil {
		return
	}
	return *o.Invitelink__
}

func NewAnnotatedTeamInviteExtWithKhulnasoft(v KhulnasoftInviteExt) AnnotatedTeamInviteExt {
	return AnnotatedTeamInviteExt{
		C__:       TeamInviteCategory_KHULNASOFT,
		Khulnasoft__: &v,
	}
}

func NewAnnotatedTeamInviteExtWithInvitelink(v InvitelinkInviteExt) AnnotatedTeamInviteExt {
	return AnnotatedTeamInviteExt{
		C__:          TeamInviteCategory_INVITELINK,
		Invitelink__: &v,
	}
}

func NewAnnotatedTeamInviteExtDefault(c TeamInviteCategory) AnnotatedTeamInviteExt {
	return AnnotatedTeamInviteExt{
		C__: c,
	}
}

func (o AnnotatedTeamInviteExt) DeepCopy() AnnotatedTeamInviteExt {
	return AnnotatedTeamInviteExt{
		C__: o.C__.DeepCopy(),
		Khulnasoft__: (func(x *KhulnasoftInviteExt) *KhulnasoftInviteExt {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Khulnasoft__),
		Invitelink__: (func(x *InvitelinkInviteExt) *InvitelinkInviteExt {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Invitelink__),
	}
}

type TeamEncryptedKBFSKeyset struct {
	V int    `codec:"v" json:"v"`
	E []byte `codec:"e" json:"e"`
	N []byte `codec:"n" json:"n"`
}

func (o TeamEncryptedKBFSKeyset) DeepCopy() TeamEncryptedKBFSKeyset {
	return TeamEncryptedKBFSKeyset{
		V: o.V,
		E: (func(x []byte) []byte {
			if x == nil {
				return nil
			}
			return append([]byte{}, x...)
		})(o.E),
		N: (func(x []byte) []byte {
			if x == nil {
				return nil
			}
			return append([]byte{}, x...)
		})(o.N),
	}
}

type TeamGetLegacyTLFUpgrade struct {
	EncryptedKeyset  string               `codec:"encryptedKeyset" json:"encrypted_keyset"`
	TeamGeneration   PerTeamKeyGeneration `codec:"teamGeneration" json:"team_generation"`
	LegacyGeneration int                  `codec:"legacyGeneration" json:"legacy_generation"`
	AppType          TeamApplication      `codec:"appType" json:"app_type"`
}

func (o TeamGetLegacyTLFUpgrade) DeepCopy() TeamGetLegacyTLFUpgrade {
	return TeamGetLegacyTLFUpgrade{
		EncryptedKeyset:  o.EncryptedKeyset,
		TeamGeneration:   o.TeamGeneration.DeepCopy(),
		LegacyGeneration: o.LegacyGeneration,
		AppType:          o.AppType.DeepCopy(),
	}
}

type TeamEncryptedKBFSKeysetHash string

func (o TeamEncryptedKBFSKeysetHash) DeepCopy() TeamEncryptedKBFSKeysetHash {
	return o
}

type TeamLegacyTLFUpgradeChainInfo struct {
	KeysetHash       TeamEncryptedKBFSKeysetHash `codec:"keysetHash" json:"keysetHash"`
	TeamGeneration   PerTeamKeyGeneration        `codec:"teamGeneration" json:"teamGeneration"`
	LegacyGeneration int                         `codec:"legacyGeneration" json:"legacyGeneration"`
	AppType          TeamApplication             `codec:"appType" json:"appType"`
}

func (o TeamLegacyTLFUpgradeChainInfo) DeepCopy() TeamLegacyTLFUpgradeChainInfo {
	return TeamLegacyTLFUpgradeChainInfo{
		KeysetHash:       o.KeysetHash.DeepCopy(),
		TeamGeneration:   o.TeamGeneration.DeepCopy(),
		LegacyGeneration: o.LegacyGeneration,
		AppType:          o.AppType.DeepCopy(),
	}
}

type TeamSignatureMetadata struct {
	SigMeta SignatureMetadata `codec:"sigMeta" json:"sigMeta"`
	Uv      UserVersion       `codec:"uv" json:"uv"`
}

func (o TeamSignatureMetadata) DeepCopy() TeamSignatureMetadata {
	return TeamSignatureMetadata{
		SigMeta: o.SigMeta.DeepCopy(),
		Uv:      o.Uv.DeepCopy(),
	}
}

type TeamInviteMetadataCancel struct {
	TeamSigMeta TeamSignatureMetadata `codec:"teamSigMeta" json:"teamSigMeta"`
}

func (o TeamInviteMetadataCancel) DeepCopy() TeamInviteMetadataCancel {
	return TeamInviteMetadataCancel{
		TeamSigMeta: o.TeamSigMeta.DeepCopy(),
	}
}

type TeamInviteMetadataCompleted struct {
	TeamSigMeta TeamSignatureMetadata `codec:"teamSigMeta" json:"teamSigMeta"`
}

func (o TeamInviteMetadataCompleted) DeepCopy() TeamInviteMetadataCompleted {
	return TeamInviteMetadataCompleted{
		TeamSigMeta: o.TeamSigMeta.DeepCopy(),
	}
}

type TeamInviteMetadataStatusCode int

const (
	TeamInviteMetadataStatusCode_ACTIVE    TeamInviteMetadataStatusCode = 0
	TeamInviteMetadataStatusCode_OBSOLETE  TeamInviteMetadataStatusCode = 1
	TeamInviteMetadataStatusCode_CANCELLED TeamInviteMetadataStatusCode = 2
	TeamInviteMetadataStatusCode_COMPLETED TeamInviteMetadataStatusCode = 3
)

func (o TeamInviteMetadataStatusCode) DeepCopy() TeamInviteMetadataStatusCode { return o }

var TeamInviteMetadataStatusCodeMap = map[string]TeamInviteMetadataStatusCode{
	"ACTIVE":    0,
	"OBSOLETE":  1,
	"CANCELLED": 2,
	"COMPLETED": 3,
}

var TeamInviteMetadataStatusCodeRevMap = map[TeamInviteMetadataStatusCode]string{
	0: "ACTIVE",
	1: "OBSOLETE",
	2: "CANCELLED",
	3: "COMPLETED",
}

func (e TeamInviteMetadataStatusCode) String() string {
	if v, ok := TeamInviteMetadataStatusCodeRevMap[e]; ok {
		return v
	}
	return fmt.Sprintf("%v", int(e))
}

type TeamInviteMetadataStatus struct {
	Code__      TeamInviteMetadataStatusCode `codec:"code" json:"code"`
	Cancelled__ *TeamInviteMetadataCancel    `codec:"cancelled,omitempty" json:"cancelled,omitempty"`
	Completed__ *TeamInviteMetadataCompleted `codec:"completed,omitempty" json:"completed,omitempty"`
}

func (o *TeamInviteMetadataStatus) Code() (ret TeamInviteMetadataStatusCode, err error) {
	switch o.Code__ {
	case TeamInviteMetadataStatusCode_CANCELLED:
		if o.Cancelled__ == nil {
			err = errors.New("unexpected nil value for Cancelled__")
			return ret, err
		}
	case TeamInviteMetadataStatusCode_COMPLETED:
		if o.Completed__ == nil {
			err = errors.New("unexpected nil value for Completed__")
			return ret, err
		}
	}
	return o.Code__, nil
}

func (o TeamInviteMetadataStatus) Cancelled() (res TeamInviteMetadataCancel) {
	if o.Code__ != TeamInviteMetadataStatusCode_CANCELLED {
		panic("wrong case accessed")
	}
	if o.Cancelled__ == nil {
		return
	}
	return *o.Cancelled__
}

func (o TeamInviteMetadataStatus) Completed() (res TeamInviteMetadataCompleted) {
	if o.Code__ != TeamInviteMetadataStatusCode_COMPLETED {
		panic("wrong case accessed")
	}
	if o.Completed__ == nil {
		return
	}
	return *o.Completed__
}

func NewTeamInviteMetadataStatusWithActive() TeamInviteMetadataStatus {
	return TeamInviteMetadataStatus{
		Code__: TeamInviteMetadataStatusCode_ACTIVE,
	}
}

func NewTeamInviteMetadataStatusWithObsolete() TeamInviteMetadataStatus {
	return TeamInviteMetadataStatus{
		Code__: TeamInviteMetadataStatusCode_OBSOLETE,
	}
}

func NewTeamInviteMetadataStatusWithCancelled(v TeamInviteMetadataCancel) TeamInviteMetadataStatus {
	return TeamInviteMetadataStatus{
		Code__:      TeamInviteMetadataStatusCode_CANCELLED,
		Cancelled__: &v,
	}
}

func NewTeamInviteMetadataStatusWithCompleted(v TeamInviteMetadataCompleted) TeamInviteMetadataStatus {
	return TeamInviteMetadataStatus{
		Code__:      TeamInviteMetadataStatusCode_COMPLETED,
		Completed__: &v,
	}
}

func (o TeamInviteMetadataStatus) DeepCopy() TeamInviteMetadataStatus {
	return TeamInviteMetadataStatus{
		Code__: o.Code__.DeepCopy(),
		Cancelled__: (func(x *TeamInviteMetadataCancel) *TeamInviteMetadataCancel {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Cancelled__),
		Completed__: (func(x *TeamInviteMetadataCompleted) *TeamInviteMetadataCompleted {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Completed__),
	}
}

type TeamInviteMetadata struct {
	Invite      TeamInvite               `codec:"invite" json:"invite"`
	TeamSigMeta TeamSignatureMetadata    `codec:"teamSigMeta" json:"teamSigMeta"`
	Status      TeamInviteMetadataStatus `codec:"status" json:"status"`
	UsedInvites []TeamUsedInviteLogPoint `codec:"usedInvites" json:"usedInvites"`
}

func (o TeamInviteMetadata) DeepCopy() TeamInviteMetadata {
	return TeamInviteMetadata{
		Invite:      o.Invite.DeepCopy(),
		TeamSigMeta: o.TeamSigMeta.DeepCopy(),
		Status:      o.Status.DeepCopy(),
		UsedInvites: (func(x []TeamUsedInviteLogPoint) []TeamUsedInviteLogPoint {
			if x == nil {
				return nil
			}
			ret := make([]TeamUsedInviteLogPoint, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.UsedInvites),
	}
}

type TeamSigChainState struct {
	Reader                  UserVersion                                       `codec:"reader" json:"reader"`
	Id                      TeamID                                            `codec:"id" json:"id"`
	Implicit                bool                                              `codec:"implicit" json:"implicit"`
	Public                  bool                                              `codec:"public" json:"public"`
	RootAncestor            TeamName                                          `codec:"rootAncestor" json:"rootAncestor"`
	NameDepth               int                                               `codec:"nameDepth" json:"nameDepth"`
	NameLog                 []TeamNameLogPoint                                `codec:"nameLog" json:"nameLog"`
	LastSeqno               Seqno                                             `codec:"lastSeqno" json:"lastSeqno"`
	LastLinkID              LinkID                                            `codec:"lastLinkID" json:"lastLinkID"`
	LastHighSeqno           Seqno                                             `codec:"lastHighSeqno" json:"lastHighSeqno"`
	LastHighLinkID          LinkID                                            `codec:"lastHighLinkID" json:"lastHighLinkID"`
	ParentID                *TeamID                                           `codec:"parentID,omitempty" json:"parentID,omitempty"`
	UserLog                 map[UserVersion][]UserLogPoint                    `codec:"userLog" json:"userLog"`
	SubteamLog              map[TeamID][]SubteamLogPoint                      `codec:"subteamLog" json:"subteamLog"`
	PerTeamKeys             map[PerTeamKeyGeneration]PerTeamKey               `codec:"perTeamKeys" json:"perTeamKeys"`
	MaxPerTeamKeyGeneration PerTeamKeyGeneration                              `codec:"maxPerTeamKeyGeneration" json:"maxPerTeamKeyGeneration"`
	PerTeamKeyCTime         UnixTime                                          `codec:"perTeamKeyCTime" json:"perTeamKeyCTime"`
	LinkIDs                 map[Seqno]LinkID                                  `codec:"linkIDs" json:"linkIDs"`
	StubbedLinks            map[Seqno]bool                                    `codec:"stubbedLinks" json:"stubbedLinks"`
	InviteMetadatas         map[TeamInviteID]TeamInviteMetadata               `codec:"inviteMetadatas" json:"inviteMetadatas"`
	Open                    bool                                              `codec:"open" json:"open"`
	OpenTeamJoinAs          TeamRole                                          `codec:"openTeamJoinAs" json:"openTeamJoinAs"`
	Bots                    map[UserVersion]TeamBotSettings                   `codec:"bots" json:"bots"`
	TlfIDs                  []TLFID                                           `codec:"tlfIDs" json:"tlfIDs"`
	TlfLegacyUpgrade        map[TeamApplication]TeamLegacyTLFUpgradeChainInfo `codec:"tlfLegacyUpgrade" json:"tlfLegacyUpgrade"`
	HeadMerkle              *MerkleRootV2                                     `codec:"headMerkle,omitempty" json:"headMerkle,omitempty"`
	MerkleRoots             map[Seqno]MerkleRootV2                            `codec:"merkleRoots" json:"merkleRoots"`
}

func (o TeamSigChainState) DeepCopy() TeamSigChainState {
	return TeamSigChainState{
		Reader:       o.Reader.DeepCopy(),
		Id:           o.Id.DeepCopy(),
		Implicit:     o.Implicit,
		Public:       o.Public,
		RootAncestor: o.RootAncestor.DeepCopy(),
		NameDepth:    o.NameDepth,
		NameLog: (func(x []TeamNameLogPoint) []TeamNameLogPoint {
			if x == nil {
				return nil
			}
			ret := make([]TeamNameLogPoint, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.NameLog),
		LastSeqno:      o.LastSeqno.DeepCopy(),
		LastLinkID:     o.LastLinkID.DeepCopy(),
		LastHighSeqno:  o.LastHighSeqno.DeepCopy(),
		LastHighLinkID: o.LastHighLinkID.DeepCopy(),
		ParentID: (func(x *TeamID) *TeamID {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.ParentID),
		UserLog: (func(x map[UserVersion][]UserLogPoint) map[UserVersion][]UserLogPoint {
			if x == nil {
				return nil
			}
			ret := make(map[UserVersion][]UserLogPoint, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := (func(x []UserLogPoint) []UserLogPoint {
					if x == nil {
						return nil
					}
					ret := make([]UserLogPoint, len(x))
					for i, v := range x {
						vCopy := v.DeepCopy()
						ret[i] = vCopy
					}
					return ret
				})(v)
				ret[kCopy] = vCopy
			}
			return ret
		})(o.UserLog),
		SubteamLog: (func(x map[TeamID][]SubteamLogPoint) map[TeamID][]SubteamLogPoint {
			if x == nil {
				return nil
			}
			ret := make(map[TeamID][]SubteamLogPoint, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := (func(x []SubteamLogPoint) []SubteamLogPoint {
					if x == nil {
						return nil
					}
					ret := make([]SubteamLogPoint, len(x))
					for i, v := range x {
						vCopy := v.DeepCopy()
						ret[i] = vCopy
					}
					return ret
				})(v)
				ret[kCopy] = vCopy
			}
			return ret
		})(o.SubteamLog),
		PerTeamKeys: (func(x map[PerTeamKeyGeneration]PerTeamKey) map[PerTeamKeyGeneration]PerTeamKey {
			if x == nil {
				return nil
			}
			ret := make(map[PerTeamKeyGeneration]PerTeamKey, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.PerTeamKeys),
		MaxPerTeamKeyGeneration: o.MaxPerTeamKeyGeneration.DeepCopy(),
		PerTeamKeyCTime:         o.PerTeamKeyCTime.DeepCopy(),
		LinkIDs: (func(x map[Seqno]LinkID) map[Seqno]LinkID {
			if x == nil {
				return nil
			}
			ret := make(map[Seqno]LinkID, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.LinkIDs),
		StubbedLinks: (func(x map[Seqno]bool) map[Seqno]bool {
			if x == nil {
				return nil
			}
			ret := make(map[Seqno]bool, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v
				ret[kCopy] = vCopy
			}
			return ret
		})(o.StubbedLinks),
		InviteMetadatas: (func(x map[TeamInviteID]TeamInviteMetadata) map[TeamInviteID]TeamInviteMetadata {
			if x == nil {
				return nil
			}
			ret := make(map[TeamInviteID]TeamInviteMetadata, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.InviteMetadatas),
		Open:           o.Open,
		OpenTeamJoinAs: o.OpenTeamJoinAs.DeepCopy(),
		Bots: (func(x map[UserVersion]TeamBotSettings) map[UserVersion]TeamBotSettings {
			if x == nil {
				return nil
			}
			ret := make(map[UserVersion]TeamBotSettings, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.Bots),
		TlfIDs: (func(x []TLFID) []TLFID {
			if x == nil {
				return nil
			}
			ret := make([]TLFID, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.TlfIDs),
		TlfLegacyUpgrade: (func(x map[TeamApplication]TeamLegacyTLFUpgradeChainInfo) map[TeamApplication]TeamLegacyTLFUpgradeChainInfo {
			if x == nil {
				return nil
			}
			ret := make(map[TeamApplication]TeamLegacyTLFUpgradeChainInfo, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.TlfLegacyUpgrade),
		HeadMerkle: (func(x *MerkleRootV2) *MerkleRootV2 {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.HeadMerkle),
		MerkleRoots: (func(x map[Seqno]MerkleRootV2) map[Seqno]MerkleRootV2 {
			if x == nil {
				return nil
			}
			ret := make(map[Seqno]MerkleRootV2, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.MerkleRoots),
	}
}

type BoxSummaryHash string

func (o BoxSummaryHash) DeepCopy() BoxSummaryHash {
	return o
}

type TeamNameLogPoint struct {
	LastPart TeamNamePart `codec:"lastPart" json:"lastPart"`
	Seqno    Seqno        `codec:"seqno" json:"seqno"`
}

func (o TeamNameLogPoint) DeepCopy() TeamNameLogPoint {
	return TeamNameLogPoint{
		LastPart: o.LastPart.DeepCopy(),
		Seqno:    o.Seqno.DeepCopy(),
	}
}

type UserLogPoint struct {
	Role    TeamRole          `codec:"role" json:"role"`
	SigMeta SignatureMetadata `codec:"sigMeta" json:"sigMeta"`
}

func (o UserLogPoint) DeepCopy() UserLogPoint {
	return UserLogPoint{
		Role:    o.Role.DeepCopy(),
		SigMeta: o.SigMeta.DeepCopy(),
	}
}

type AnnotatedTeamUsedInviteLogPoint struct {
	Username               string                 `codec:"username" json:"username"`
	TeamUsedInviteLogPoint TeamUsedInviteLogPoint `codec:"teamUsedInviteLogPoint" json:"teamUsedInviteLogPoint"`
}

func (o AnnotatedTeamUsedInviteLogPoint) DeepCopy() AnnotatedTeamUsedInviteLogPoint {
	return AnnotatedTeamUsedInviteLogPoint{
		Username:               o.Username,
		TeamUsedInviteLogPoint: o.TeamUsedInviteLogPoint.DeepCopy(),
	}
}

type TeamUsedInviteLogPoint struct {
	Uv       UserVersion `codec:"uv" json:"uv"`
	LogPoint int         `codec:"logPoint" json:"logPoint"`
}

func (o TeamUsedInviteLogPoint) DeepCopy() TeamUsedInviteLogPoint {
	return TeamUsedInviteLogPoint{
		Uv:       o.Uv.DeepCopy(),
		LogPoint: o.LogPoint,
	}
}

type SubteamLogPoint struct {
	Name  TeamName `codec:"name" json:"name"`
	Seqno Seqno    `codec:"seqno" json:"seqno"`
}

func (o SubteamLogPoint) DeepCopy() SubteamLogPoint {
	return SubteamLogPoint{
		Name:  o.Name.DeepCopy(),
		Seqno: o.Seqno.DeepCopy(),
	}
}

type TeamNamePart string

func (o TeamNamePart) DeepCopy() TeamNamePart {
	return o
}

type TeamName struct {
	Parts []TeamNamePart `codec:"parts" json:"parts"`
}

func (o TeamName) DeepCopy() TeamName {
	return TeamName{
		Parts: (func(x []TeamNamePart) []TeamNamePart {
			if x == nil {
				return nil
			}
			ret := make([]TeamNamePart, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Parts),
	}
}

type TeamCLKRResetUser struct {
	Uid               UID   `codec:"uid" json:"uid"`
	UserEldestSeqno   Seqno `codec:"userEldestSeqno" json:"user_eldest"`
	MemberEldestSeqno Seqno `codec:"memberEldestSeqno" json:"member_eldest"`
}

func (o TeamCLKRResetUser) DeepCopy() TeamCLKRResetUser {
	return TeamCLKRResetUser{
		Uid:               o.Uid.DeepCopy(),
		UserEldestSeqno:   o.UserEldestSeqno.DeepCopy(),
		MemberEldestSeqno: o.MemberEldestSeqno.DeepCopy(),
	}
}

type TeamCLKRMsg struct {
	TeamID              TeamID               `codec:"teamID" json:"team_id"`
	Generation          PerTeamKeyGeneration `codec:"generation" json:"generation"`
	Score               int                  `codec:"score" json:"score"`
	ResetUsersUntrusted []TeamCLKRResetUser  `codec:"resetUsersUntrusted" json:"reset_users"`
}

func (o TeamCLKRMsg) DeepCopy() TeamCLKRMsg {
	return TeamCLKRMsg{
		TeamID:     o.TeamID.DeepCopy(),
		Generation: o.Generation.DeepCopy(),
		Score:      o.Score,
		ResetUsersUntrusted: (func(x []TeamCLKRResetUser) []TeamCLKRResetUser {
			if x == nil {
				return nil
			}
			ret := make([]TeamCLKRResetUser, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.ResetUsersUntrusted),
	}
}

type TeamResetUser struct {
	Username    string `codec:"username" json:"username"`
	Uid         UID    `codec:"uid" json:"uid"`
	EldestSeqno Seqno  `codec:"eldestSeqno" json:"eldest_seqno"`
	IsDelete    bool   `codec:"isDelete" json:"is_delete"`
}

func (o TeamResetUser) DeepCopy() TeamResetUser {
	return TeamResetUser{
		Username:    o.Username,
		Uid:         o.Uid.DeepCopy(),
		EldestSeqno: o.EldestSeqno.DeepCopy(),
		IsDelete:    o.IsDelete,
	}
}

type TeamMemberOutFromReset struct {
	TeamID    TeamID        `codec:"teamID" json:"team_id"`
	TeamName  string        `codec:"teamName" json:"team_name"`
	ResetUser TeamResetUser `codec:"resetUser" json:"reset_user"`
}

func (o TeamMemberOutFromReset) DeepCopy() TeamMemberOutFromReset {
	return TeamMemberOutFromReset{
		TeamID:    o.TeamID.DeepCopy(),
		TeamName:  o.TeamName,
		ResetUser: o.ResetUser.DeepCopy(),
	}
}

type TeamChangeRow struct {
	Id                  TeamID `codec:"id" json:"id"`
	Name                string `codec:"name" json:"name"`
	KeyRotated          bool   `codec:"keyRotated" json:"key_rotated"`
	MembershipChanged   bool   `codec:"membershipChanged" json:"membership_changed"`
	LatestSeqno         Seqno  `codec:"latestSeqno" json:"latest_seqno"`
	LatestHiddenSeqno   Seqno  `codec:"latestHiddenSeqno" json:"latest_hidden_seqno"`
	LatestOffchainSeqno Seqno  `codec:"latestOffchainSeqno" json:"latest_offchain_version"`
	ImplicitTeam        bool   `codec:"implicitTeam" json:"implicit_team"`
	Misc                bool   `codec:"misc" json:"misc"`
	RemovedResetUsers   bool   `codec:"removedResetUsers" json:"removed_reset_users"`
}

func (o TeamChangeRow) DeepCopy() TeamChangeRow {
	return TeamChangeRow{
		Id:                  o.Id.DeepCopy(),
		Name:                o.Name,
		KeyRotated:          o.KeyRotated,
		MembershipChanged:   o.MembershipChanged,
		LatestSeqno:         o.LatestSeqno.DeepCopy(),
		LatestHiddenSeqno:   o.LatestHiddenSeqno.DeepCopy(),
		LatestOffchainSeqno: o.LatestOffchainSeqno.DeepCopy(),
		ImplicitTeam:        o.ImplicitTeam,
		Misc:                o.Misc,
		RemovedResetUsers:   o.RemovedResetUsers,
	}
}

type TeamExitRow struct {
	Id TeamID `codec:"id" json:"id"`
}

func (o TeamExitRow) DeepCopy() TeamExitRow {
	return TeamExitRow{
		Id: o.Id.DeepCopy(),
	}
}

type TeamNewlyAddedRow struct {
	Id   TeamID `codec:"id" json:"id"`
	Name string `codec:"name" json:"name"`
}

func (o TeamNewlyAddedRow) DeepCopy() TeamNewlyAddedRow {
	return TeamNewlyAddedRow{
		Id:   o.Id.DeepCopy(),
		Name: o.Name,
	}
}

type TeamInvitee struct {
	InviteID    TeamInviteID `codec:"inviteID" json:"invite_id"`
	Uid         UID          `codec:"uid" json:"uid"`
	EldestSeqno Seqno        `codec:"eldestSeqno" json:"eldest_seqno"`
	Role        TeamRole     `codec:"role" json:"role"`
}

func (o TeamInvitee) DeepCopy() TeamInvitee {
	return TeamInvitee{
		InviteID:    o.InviteID.DeepCopy(),
		Uid:         o.Uid.DeepCopy(),
		EldestSeqno: o.EldestSeqno.DeepCopy(),
		Role:        o.Role.DeepCopy(),
	}
}

type TeamSBSMsg struct {
	TeamID   TeamID        `codec:"teamID" json:"team_id"`
	Score    int           `codec:"score" json:"score"`
	Invitees []TeamInvitee `codec:"invitees" json:"invitees"`
}

func (o TeamSBSMsg) DeepCopy() TeamSBSMsg {
	return TeamSBSMsg{
		TeamID: o.TeamID.DeepCopy(),
		Score:  o.Score,
		Invitees: (func(x []TeamInvitee) []TeamInvitee {
			if x == nil {
				return nil
			}
			ret := make([]TeamInvitee, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Invitees),
	}
}

type TeamAccessRequest struct {
	Uid         UID   `codec:"uid" json:"uid"`
	EldestSeqno Seqno `codec:"eldestSeqno" json:"eldest_seqno"`
}

func (o TeamAccessRequest) DeepCopy() TeamAccessRequest {
	return TeamAccessRequest{
		Uid:         o.Uid.DeepCopy(),
		EldestSeqno: o.EldestSeqno.DeepCopy(),
	}
}

type TeamOpenReqMsg struct {
	TeamID TeamID              `codec:"teamID" json:"team_id"`
	Tars   []TeamAccessRequest `codec:"tars" json:"tars"`
}

func (o TeamOpenReqMsg) DeepCopy() TeamOpenReqMsg {
	return TeamOpenReqMsg{
		TeamID: o.TeamID.DeepCopy(),
		Tars: (func(x []TeamAccessRequest) []TeamAccessRequest {
			if x == nil {
				return nil
			}
			ret := make([]TeamAccessRequest, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Tars),
	}
}

type SeitanAKey string

func (o SeitanAKey) DeepCopy() SeitanAKey {
	return o
}

type SeitanIKey string

func (o SeitanIKey) DeepCopy() SeitanIKey {
	return o
}

type SeitanIKeyInvitelink string

func (o SeitanIKeyInvitelink) DeepCopy() SeitanIKeyInvitelink {
	return o
}

type SeitanPubKey string

func (o SeitanPubKey) DeepCopy() SeitanPubKey {
	return o
}

type SeitanIKeyV2 string

func (o SeitanIKeyV2) DeepCopy() SeitanIKeyV2 {
	return o
}

type SeitanKeyAndLabelVersion int

const (
	SeitanKeyAndLabelVersion_V1         SeitanKeyAndLabelVersion = 1
	SeitanKeyAndLabelVersion_V2         SeitanKeyAndLabelVersion = 2
	SeitanKeyAndLabelVersion_Invitelink SeitanKeyAndLabelVersion = 3
)

func (o SeitanKeyAndLabelVersion) DeepCopy() SeitanKeyAndLabelVersion { return o }

var SeitanKeyAndLabelVersionMap = map[string]SeitanKeyAndLabelVersion{
	"V1":         1,
	"V2":         2,
	"Invitelink": 3,
}

var SeitanKeyAndLabelVersionRevMap = map[SeitanKeyAndLabelVersion]string{
	1: "V1",
	2: "V2",
	3: "Invitelink",
}

func (e SeitanKeyAndLabelVersion) String() string {
	if v, ok := SeitanKeyAndLabelVersionRevMap[e]; ok {
		return v
	}
	return fmt.Sprintf("%v", int(e))
}

type SeitanKeyAndLabel struct {
	V__          SeitanKeyAndLabelVersion     `codec:"v" json:"v"`
	V1__         *SeitanKeyAndLabelVersion1   `codec:"v1,omitempty" json:"v1,omitempty"`
	V2__         *SeitanKeyAndLabelVersion2   `codec:"v2,omitempty" json:"v2,omitempty"`
	Invitelink__ *SeitanKeyAndLabelInvitelink `codec:"invitelink,omitempty" json:"invitelink,omitempty"`
}

func (o *SeitanKeyAndLabel) V() (ret SeitanKeyAndLabelVersion, err error) {
	switch o.V__ {
	case SeitanKeyAndLabelVersion_V1:
		if o.V1__ == nil {
			err = errors.New("unexpected nil value for V1__")
			return ret, err
		}
	case SeitanKeyAndLabelVersion_V2:
		if o.V2__ == nil {
			err = errors.New("unexpected nil value for V2__")
			return ret, err
		}
	case SeitanKeyAndLabelVersion_Invitelink:
		if o.Invitelink__ == nil {
			err = errors.New("unexpected nil value for Invitelink__")
			return ret, err
		}
	}
	return o.V__, nil
}

func (o SeitanKeyAndLabel) V1() (res SeitanKeyAndLabelVersion1) {
	if o.V__ != SeitanKeyAndLabelVersion_V1 {
		panic("wrong case accessed")
	}
	if o.V1__ == nil {
		return
	}
	return *o.V1__
}

func (o SeitanKeyAndLabel) V2() (res SeitanKeyAndLabelVersion2) {
	if o.V__ != SeitanKeyAndLabelVersion_V2 {
		panic("wrong case accessed")
	}
	if o.V2__ == nil {
		return
	}
	return *o.V2__
}

func (o SeitanKeyAndLabel) Invitelink() (res SeitanKeyAndLabelInvitelink) {
	if o.V__ != SeitanKeyAndLabelVersion_Invitelink {
		panic("wrong case accessed")
	}
	if o.Invitelink__ == nil {
		return
	}
	return *o.Invitelink__
}

func NewSeitanKeyAndLabelWithV1(v SeitanKeyAndLabelVersion1) SeitanKeyAndLabel {
	return SeitanKeyAndLabel{
		V__:  SeitanKeyAndLabelVersion_V1,
		V1__: &v,
	}
}

func NewSeitanKeyAndLabelWithV2(v SeitanKeyAndLabelVersion2) SeitanKeyAndLabel {
	return SeitanKeyAndLabel{
		V__:  SeitanKeyAndLabelVersion_V2,
		V2__: &v,
	}
}

func NewSeitanKeyAndLabelWithInvitelink(v SeitanKeyAndLabelInvitelink) SeitanKeyAndLabel {
	return SeitanKeyAndLabel{
		V__:          SeitanKeyAndLabelVersion_Invitelink,
		Invitelink__: &v,
	}
}

func NewSeitanKeyAndLabelDefault(v SeitanKeyAndLabelVersion) SeitanKeyAndLabel {
	return SeitanKeyAndLabel{
		V__: v,
	}
}

func (o SeitanKeyAndLabel) DeepCopy() SeitanKeyAndLabel {
	return SeitanKeyAndLabel{
		V__: o.V__.DeepCopy(),
		V1__: (func(x *SeitanKeyAndLabelVersion1) *SeitanKeyAndLabelVersion1 {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.V1__),
		V2__: (func(x *SeitanKeyAndLabelVersion2) *SeitanKeyAndLabelVersion2 {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.V2__),
		Invitelink__: (func(x *SeitanKeyAndLabelInvitelink) *SeitanKeyAndLabelInvitelink {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Invitelink__),
	}
}

type SeitanKeyAndLabelVersion1 struct {
	I SeitanIKey     `codec:"i" json:"i"`
	L SeitanKeyLabel `codec:"l" json:"l"`
}

func (o SeitanKeyAndLabelVersion1) DeepCopy() SeitanKeyAndLabelVersion1 {
	return SeitanKeyAndLabelVersion1{
		I: o.I.DeepCopy(),
		L: o.L.DeepCopy(),
	}
}

type SeitanKeyAndLabelVersion2 struct {
	K SeitanPubKey   `codec:"k" json:"k"`
	L SeitanKeyLabel `codec:"l" json:"l"`
}

func (o SeitanKeyAndLabelVersion2) DeepCopy() SeitanKeyAndLabelVersion2 {
	return SeitanKeyAndLabelVersion2{
		K: o.K.DeepCopy(),
		L: o.L.DeepCopy(),
	}
}

type SeitanKeyAndLabelInvitelink struct {
	I SeitanIKeyInvitelink `codec:"i" json:"i"`
	L SeitanKeyLabel       `codec:"l" json:"l"`
}

func (o SeitanKeyAndLabelInvitelink) DeepCopy() SeitanKeyAndLabelInvitelink {
	return SeitanKeyAndLabelInvitelink{
		I: o.I.DeepCopy(),
		L: o.L.DeepCopy(),
	}
}

type SeitanKeyLabelType int

const (
	SeitanKeyLabelType_SMS     SeitanKeyLabelType = 1
	SeitanKeyLabelType_GENERIC SeitanKeyLabelType = 2
)

func (o SeitanKeyLabelType) DeepCopy() SeitanKeyLabelType { return o }

var SeitanKeyLabelTypeMap = map[string]SeitanKeyLabelType{
	"SMS":     1,
	"GENERIC": 2,
}

var SeitanKeyLabelTypeRevMap = map[SeitanKeyLabelType]string{
	1: "SMS",
	2: "GENERIC",
}

func (e SeitanKeyLabelType) String() string {
	if v, ok := SeitanKeyLabelTypeRevMap[e]; ok {
		return v
	}
	return fmt.Sprintf("%v", int(e))
}

type SeitanKeyLabel struct {
	T__       SeitanKeyLabelType     `codec:"t" json:"t"`
	Sms__     *SeitanKeyLabelSms     `codec:"sms,omitempty" json:"sms,omitempty"`
	Generic__ *SeitanKeyLabelGeneric `codec:"generic,omitempty" json:"generic,omitempty"`
}

func (o *SeitanKeyLabel) T() (ret SeitanKeyLabelType, err error) {
	switch o.T__ {
	case SeitanKeyLabelType_SMS:
		if o.Sms__ == nil {
			err = errors.New("unexpected nil value for Sms__")
			return ret, err
		}
	case SeitanKeyLabelType_GENERIC:
		if o.Generic__ == nil {
			err = errors.New("unexpected nil value for Generic__")
			return ret, err
		}
	}
	return o.T__, nil
}

func (o SeitanKeyLabel) Sms() (res SeitanKeyLabelSms) {
	if o.T__ != SeitanKeyLabelType_SMS {
		panic("wrong case accessed")
	}
	if o.Sms__ == nil {
		return
	}
	return *o.Sms__
}

func (o SeitanKeyLabel) Generic() (res SeitanKeyLabelGeneric) {
	if o.T__ != SeitanKeyLabelType_GENERIC {
		panic("wrong case accessed")
	}
	if o.Generic__ == nil {
		return
	}
	return *o.Generic__
}

func NewSeitanKeyLabelWithSms(v SeitanKeyLabelSms) SeitanKeyLabel {
	return SeitanKeyLabel{
		T__:   SeitanKeyLabelType_SMS,
		Sms__: &v,
	}
}

func NewSeitanKeyLabelWithGeneric(v SeitanKeyLabelGeneric) SeitanKeyLabel {
	return SeitanKeyLabel{
		T__:       SeitanKeyLabelType_GENERIC,
		Generic__: &v,
	}
}

func NewSeitanKeyLabelDefault(t SeitanKeyLabelType) SeitanKeyLabel {
	return SeitanKeyLabel{
		T__: t,
	}
}

func (o SeitanKeyLabel) DeepCopy() SeitanKeyLabel {
	return SeitanKeyLabel{
		T__: o.T__.DeepCopy(),
		Sms__: (func(x *SeitanKeyLabelSms) *SeitanKeyLabelSms {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Sms__),
		Generic__: (func(x *SeitanKeyLabelGeneric) *SeitanKeyLabelGeneric {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Generic__),
	}
}

type SeitanKeyLabelSms struct {
	F string `codec:"f" json:"f"`
	N string `codec:"n" json:"n"`
}

func (o SeitanKeyLabelSms) DeepCopy() SeitanKeyLabelSms {
	return SeitanKeyLabelSms{
		F: o.F,
		N: o.N,
	}
}

type SeitanKeyLabelGeneric struct {
	L string `codec:"l" json:"l"`
}

func (o SeitanKeyLabelGeneric) DeepCopy() SeitanKeyLabelGeneric {
	return SeitanKeyLabelGeneric{
		L: o.L,
	}
}

type TeamSeitanRequest struct {
	InviteID    TeamInviteID `codec:"inviteID" json:"invite_id"`
	Uid         UID          `codec:"uid" json:"uid"`
	EldestSeqno Seqno        `codec:"eldestSeqno" json:"eldest_seqno"`
	Akey        SeitanAKey   `codec:"akey" json:"akey"`
	Role        TeamRole     `codec:"role" json:"role"`
	UnixCTime   int64        `codec:"unixCTime" json:"ctime"`
}

func (o TeamSeitanRequest) DeepCopy() TeamSeitanRequest {
	return TeamSeitanRequest{
		InviteID:    o.InviteID.DeepCopy(),
		Uid:         o.Uid.DeepCopy(),
		EldestSeqno: o.EldestSeqno.DeepCopy(),
		Akey:        o.Akey.DeepCopy(),
		Role:        o.Role.DeepCopy(),
		UnixCTime:   o.UnixCTime,
	}
}

type TeamSeitanMsg struct {
	TeamID  TeamID              `codec:"teamID" json:"team_id"`
	Seitans []TeamSeitanRequest `codec:"seitans" json:"seitans"`
}

func (o TeamSeitanMsg) DeepCopy() TeamSeitanMsg {
	return TeamSeitanMsg{
		TeamID: o.TeamID.DeepCopy(),
		Seitans: (func(x []TeamSeitanRequest) []TeamSeitanRequest {
			if x == nil {
				return nil
			}
			ret := make([]TeamSeitanRequest, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Seitans),
	}
}

type TeamOpenSweepMsg struct {
	TeamID              TeamID              `codec:"teamID" json:"team_id"`
	ResetUsersUntrusted []TeamCLKRResetUser `codec:"resetUsersUntrusted" json:"reset_users"`
}

func (o TeamOpenSweepMsg) DeepCopy() TeamOpenSweepMsg {
	return TeamOpenSweepMsg{
		TeamID: o.TeamID.DeepCopy(),
		ResetUsersUntrusted: (func(x []TeamCLKRResetUser) []TeamCLKRResetUser {
			if x == nil {
				return nil
			}
			ret := make([]TeamCLKRResetUser, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.ResetUsersUntrusted),
	}
}

type TeamKBFSKeyRefresher struct {
	Generation int             `codec:"generation" json:"generation"`
	AppType    TeamApplication `codec:"appType" json:"appType"`
}

func (o TeamKBFSKeyRefresher) DeepCopy() TeamKBFSKeyRefresher {
	return TeamKBFSKeyRefresher{
		Generation: o.Generation,
		AppType:    o.AppType.DeepCopy(),
	}
}

// * TeamRefreshData are needed or wanted data requirements that, if unmet, will cause
// * a refresh of the cache.
type TeamRefreshers struct {
	NeedKeyGeneration                     PerTeamKeyGeneration                       `codec:"needKeyGeneration" json:"needKeyGeneration"`
	NeedApplicationsAtGenerations         map[PerTeamKeyGeneration][]TeamApplication `codec:"needApplicationsAtGenerations" json:"needApplicationsAtGenerations"`
	NeedApplicationsAtGenerationsWithKBFS map[PerTeamKeyGeneration][]TeamApplication `codec:"needApplicationsAtGenerationsWithKBFS" json:"needApplicationsAtGenerationsWithKBFS"`
	WantMembers                           []UserVersion                              `codec:"wantMembers" json:"wantMembers"`
	WantMembersRole                       TeamRole                                   `codec:"wantMembersRole" json:"wantMembersRole"`
	NeedKBFSKeyGeneration                 TeamKBFSKeyRefresher                       `codec:"needKBFSKeyGeneration" json:"needKBFSKeyGeneration"`
}

func (o TeamRefreshers) DeepCopy() TeamRefreshers {
	return TeamRefreshers{
		NeedKeyGeneration: o.NeedKeyGeneration.DeepCopy(),
		NeedApplicationsAtGenerations: (func(x map[PerTeamKeyGeneration][]TeamApplication) map[PerTeamKeyGeneration][]TeamApplication {
			if x == nil {
				return nil
			}
			ret := make(map[PerTeamKeyGeneration][]TeamApplication, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := (func(x []TeamApplication) []TeamApplication {
					if x == nil {
						return nil
					}
					ret := make([]TeamApplication, len(x))
					for i, v := range x {
						vCopy := v.DeepCopy()
						ret[i] = vCopy
					}
					return ret
				})(v)
				ret[kCopy] = vCopy
			}
			return ret
		})(o.NeedApplicationsAtGenerations),
		NeedApplicationsAtGenerationsWithKBFS: (func(x map[PerTeamKeyGeneration][]TeamApplication) map[PerTeamKeyGeneration][]TeamApplication {
			if x == nil {
				return nil
			}
			ret := make(map[PerTeamKeyGeneration][]TeamApplication, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := (func(x []TeamApplication) []TeamApplication {
					if x == nil {
						return nil
					}
					ret := make([]TeamApplication, len(x))
					for i, v := range x {
						vCopy := v.DeepCopy()
						ret[i] = vCopy
					}
					return ret
				})(v)
				ret[kCopy] = vCopy
			}
			return ret
		})(o.NeedApplicationsAtGenerationsWithKBFS),
		WantMembers: (func(x []UserVersion) []UserVersion {
			if x == nil {
				return nil
			}
			ret := make([]UserVersion, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.WantMembers),
		WantMembersRole:       o.WantMembersRole.DeepCopy(),
		NeedKBFSKeyGeneration: o.NeedKBFSKeyGeneration.DeepCopy(),
	}
}

type LoadTeamArg struct {
	ID                        TeamID         `codec:"ID" json:"ID"`
	Name                      string         `codec:"name" json:"name"`
	Public                    bool           `codec:"public" json:"public"`
	NeedAdmin                 bool           `codec:"needAdmin" json:"needAdmin"`
	RefreshUIDMapper          bool           `codec:"refreshUIDMapper" json:"refreshUIDMapper"`
	Refreshers                TeamRefreshers `codec:"refreshers" json:"refreshers"`
	ForceFullReload           bool           `codec:"forceFullReload" json:"forceFullReload"`
	ForceRepoll               bool           `codec:"forceRepoll" json:"forceRepoll"`
	StaleOK                   bool           `codec:"staleOK" json:"staleOK"`
	AllowNameLookupBurstCache bool           `codec:"allowNameLookupBurstCache" json:"allowNameLookupBurstCache"`
	SkipNeedHiddenRotateCheck bool           `codec:"skipNeedHiddenRotateCheck" json:"skipNeedHiddenRotateCheck"`
	AuditMode                 AuditMode      `codec:"auditMode" json:"auditMode"`
}

func (o LoadTeamArg) DeepCopy() LoadTeamArg {
	return LoadTeamArg{
		ID:                        o.ID.DeepCopy(),
		Name:                      o.Name,
		Public:                    o.Public,
		NeedAdmin:                 o.NeedAdmin,
		RefreshUIDMapper:          o.RefreshUIDMapper,
		Refreshers:                o.Refreshers.DeepCopy(),
		ForceFullReload:           o.ForceFullReload,
		ForceRepoll:               o.ForceRepoll,
		StaleOK:                   o.StaleOK,
		AllowNameLookupBurstCache: o.AllowNameLookupBurstCache,
		SkipNeedHiddenRotateCheck: o.SkipNeedHiddenRotateCheck,
		AuditMode:                 o.AuditMode.DeepCopy(),
	}
}

type FastTeamLoadArg struct {
	ID                    TeamID                 `codec:"ID" json:"ID"`
	Public                bool                   `codec:"public" json:"public"`
	AssertTeamName        *TeamName              `codec:"assertTeamName,omitempty" json:"assertTeamName,omitempty"`
	Applications          []TeamApplication      `codec:"applications" json:"applications"`
	KeyGenerationsNeeded  []PerTeamKeyGeneration `codec:"keyGenerationsNeeded" json:"keyGenerationsNeeded"`
	NeedLatestKey         bool                   `codec:"needLatestKey" json:"needLatestKey"`
	ForceRefresh          bool                   `codec:"forceRefresh" json:"forceRefresh"`
	HiddenChainIsOptional bool                   `codec:"hiddenChainIsOptional" json:"hiddenChainIsOptional"`
}

func (o FastTeamLoadArg) DeepCopy() FastTeamLoadArg {
	return FastTeamLoadArg{
		ID:     o.ID.DeepCopy(),
		Public: o.Public,
		AssertTeamName: (func(x *TeamName) *TeamName {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.AssertTeamName),
		Applications: (func(x []TeamApplication) []TeamApplication {
			if x == nil {
				return nil
			}
			ret := make([]TeamApplication, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Applications),
		KeyGenerationsNeeded: (func(x []PerTeamKeyGeneration) []PerTeamKeyGeneration {
			if x == nil {
				return nil
			}
			ret := make([]PerTeamKeyGeneration, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.KeyGenerationsNeeded),
		NeedLatestKey:         o.NeedLatestKey,
		ForceRefresh:          o.ForceRefresh,
		HiddenChainIsOptional: o.HiddenChainIsOptional,
	}
}

type FastTeamLoadRes struct {
	Name            TeamName             `codec:"name" json:"name"`
	ApplicationKeys []TeamApplicationKey `codec:"applicationKeys" json:"applicationKeys"`
}

func (o FastTeamLoadRes) DeepCopy() FastTeamLoadRes {
	return FastTeamLoadRes{
		Name: o.Name.DeepCopy(),
		ApplicationKeys: (func(x []TeamApplicationKey) []TeamApplicationKey {
			if x == nil {
				return nil
			}
			ret := make([]TeamApplicationKey, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.ApplicationKeys),
	}
}

type ImplicitRole struct {
	Role     TeamRole `codec:"role" json:"role"`
	Ancestor TeamID   `codec:"ancestor" json:"ancestor"`
}

func (o ImplicitRole) DeepCopy() ImplicitRole {
	return ImplicitRole{
		Role:     o.Role.DeepCopy(),
		Ancestor: o.Ancestor.DeepCopy(),
	}
}

type MemberInfo struct {
	UserID              UID           `codec:"userID" json:"uid"`
	TeamID              TeamID        `codec:"teamID" json:"team_id"`
	FqName              string        `codec:"fqName" json:"fq_name"`
	IsImplicitTeam      bool          `codec:"isImplicitTeam" json:"is_implicit_team"`
	IsOpenTeam          bool          `codec:"isOpenTeam" json:"is_open_team"`
	Role                TeamRole      `codec:"role" json:"role"`
	Implicit            *ImplicitRole `codec:"implicit,omitempty" json:"implicit,omitempty"`
	MemberCount         int           `codec:"memberCount" json:"member_count"`
	AllowProfilePromote bool          `codec:"allowProfilePromote" json:"allow_profile_promote"`
	IsMemberShowcased   bool          `codec:"isMemberShowcased" json:"is_member_showcased"`
}

func (o MemberInfo) DeepCopy() MemberInfo {
	return MemberInfo{
		UserID:         o.UserID.DeepCopy(),
		TeamID:         o.TeamID.DeepCopy(),
		FqName:         o.FqName,
		IsImplicitTeam: o.IsImplicitTeam,
		IsOpenTeam:     o.IsOpenTeam,
		Role:           o.Role.DeepCopy(),
		Implicit: (func(x *ImplicitRole) *ImplicitRole {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Implicit),
		MemberCount:         o.MemberCount,
		AllowProfilePromote: o.AllowProfilePromote,
		IsMemberShowcased:   o.IsMemberShowcased,
	}
}

type TeamList struct {
	Teams []MemberInfo `codec:"teams" json:"teams"`
}

func (o TeamList) DeepCopy() TeamList {
	return TeamList{
		Teams: (func(x []MemberInfo) []MemberInfo {
			if x == nil {
				return nil
			}
			ret := make([]MemberInfo, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Teams),
	}
}

type AnnotatedMemberInfo struct {
	UserID              UID              `codec:"userID" json:"uid"`
	TeamID              TeamID           `codec:"teamID" json:"team_id"`
	Username            string           `codec:"username" json:"username"`
	FullName            string           `codec:"fullName" json:"full_name"`
	FqName              string           `codec:"fqName" json:"fq_name"`
	IsImplicitTeam      bool             `codec:"isImplicitTeam" json:"is_implicit_team"`
	ImpTeamDisplayName  string           `codec:"impTeamDisplayName" json:"implicit_team_display_name"`
	IsOpenTeam          bool             `codec:"isOpenTeam" json:"is_open_team"`
	Role                TeamRole         `codec:"role" json:"role"`
	Implicit            *ImplicitRole    `codec:"implicit,omitempty" json:"implicit,omitempty"`
	NeedsPUK            bool             `codec:"needsPUK" json:"needsPUK"`
	MemberCount         int              `codec:"memberCount" json:"member_count"`
	EldestSeqno         Seqno            `codec:"eldestSeqno" json:"member_eldest_seqno"`
	AllowProfilePromote bool             `codec:"allowProfilePromote" json:"allow_profile_promote"`
	IsMemberShowcased   bool             `codec:"isMemberShowcased" json:"is_member_showcased"`
	Status              TeamMemberStatus `codec:"status" json:"status"`
}

func (o AnnotatedMemberInfo) DeepCopy() AnnotatedMemberInfo {
	return AnnotatedMemberInfo{
		UserID:             o.UserID.DeepCopy(),
		TeamID:             o.TeamID.DeepCopy(),
		Username:           o.Username,
		FullName:           o.FullName,
		FqName:             o.FqName,
		IsImplicitTeam:     o.IsImplicitTeam,
		ImpTeamDisplayName: o.ImpTeamDisplayName,
		IsOpenTeam:         o.IsOpenTeam,
		Role:               o.Role.DeepCopy(),
		Implicit: (func(x *ImplicitRole) *ImplicitRole {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Implicit),
		NeedsPUK:            o.NeedsPUK,
		MemberCount:         o.MemberCount,
		EldestSeqno:         o.EldestSeqno.DeepCopy(),
		AllowProfilePromote: o.AllowProfilePromote,
		IsMemberShowcased:   o.IsMemberShowcased,
		Status:              o.Status.DeepCopy(),
	}
}

type AnnotatedTeamList struct {
	Teams []AnnotatedMemberInfo `codec:"teams" json:"teams"`
}

func (o AnnotatedTeamList) DeepCopy() AnnotatedTeamList {
	return AnnotatedTeamList{
		Teams: (func(x []AnnotatedMemberInfo) []AnnotatedMemberInfo {
			if x == nil {
				return nil
			}
			ret := make([]AnnotatedMemberInfo, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Teams),
	}
}

type TeamAddMemberResult struct {
	Invited     bool  `codec:"invited" json:"invited"`
	User        *User `codec:"user,omitempty" json:"user,omitempty"`
	ChatSending bool  `codec:"chatSending" json:"chatSending"`
}

func (o TeamAddMemberResult) DeepCopy() TeamAddMemberResult {
	return TeamAddMemberResult{
		Invited: o.Invited,
		User: (func(x *User) *User {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.User),
		ChatSending: o.ChatSending,
	}
}

type TeamAddMembersResult struct {
	NotAdded []User `codec:"notAdded" json:"notAdded"`
}

func (o TeamAddMembersResult) DeepCopy() TeamAddMembersResult {
	return TeamAddMembersResult{
		NotAdded: (func(x []User) []User {
			if x == nil {
				return nil
			}
			ret := make([]User, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.NotAdded),
	}
}

type TeamJoinRequest struct {
	Name     string   `codec:"name" json:"name"`
	Username string   `codec:"username" json:"username"`
	FullName FullName `codec:"fullName" json:"fullName"`
	Ctime    UnixTime `codec:"ctime" json:"ctime"`
}

func (o TeamJoinRequest) DeepCopy() TeamJoinRequest {
	return TeamJoinRequest{
		Name:     o.Name,
		Username: o.Username,
		FullName: o.FullName.DeepCopy(),
		Ctime:    o.Ctime.DeepCopy(),
	}
}

type TeamTreeResult struct {
	Entries []TeamTreeEntry `codec:"entries" json:"entries"`
}

func (o TeamTreeResult) DeepCopy() TeamTreeResult {
	return TeamTreeResult{
		Entries: (func(x []TeamTreeEntry) []TeamTreeEntry {
			if x == nil {
				return nil
			}
			ret := make([]TeamTreeEntry, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Entries),
	}
}

type TeamTreeEntry struct {
	Name  TeamName `codec:"name" json:"name"`
	Admin bool     `codec:"admin" json:"admin"`
}

func (o TeamTreeEntry) DeepCopy() TeamTreeEntry {
	return TeamTreeEntry{
		Name:  o.Name.DeepCopy(),
		Admin: o.Admin,
	}
}

type SubteamListEntry struct {
	Name        TeamName `codec:"name" json:"name"`
	TeamID      TeamID   `codec:"teamID" json:"teamID"`
	MemberCount int      `codec:"memberCount" json:"memberCount"`
}

func (o SubteamListEntry) DeepCopy() SubteamListEntry {
	return SubteamListEntry{
		Name:        o.Name.DeepCopy(),
		TeamID:      o.TeamID.DeepCopy(),
		MemberCount: o.MemberCount,
	}
}

type SubteamListResult struct {
	Entries []SubteamListEntry `codec:"entries" json:"entries"`
}

func (o SubteamListResult) DeepCopy() SubteamListResult {
	return SubteamListResult{
		Entries: (func(x []SubteamListEntry) []SubteamListEntry {
			if x == nil {
				return nil
			}
			ret := make([]SubteamListEntry, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Entries),
	}
}

type TeamCreateResult struct {
	TeamID       TeamID `codec:"teamID" json:"teamID"`
	ChatSent     bool   `codec:"chatSent" json:"chatSent"`
	CreatorAdded bool   `codec:"creatorAdded" json:"creatorAdded"`
}

func (o TeamCreateResult) DeepCopy() TeamCreateResult {
	return TeamCreateResult{
		TeamID:       o.TeamID.DeepCopy(),
		ChatSent:     o.ChatSent,
		CreatorAdded: o.CreatorAdded,
	}
}

type TeamSettings struct {
	Open   bool     `codec:"open" json:"open"`
	JoinAs TeamRole `codec:"joinAs" json:"joinAs"`
}

func (o TeamSettings) DeepCopy() TeamSettings {
	return TeamSettings{
		Open:   o.Open,
		JoinAs: o.JoinAs.DeepCopy(),
	}
}

type TeamBotSettings struct {
	Cmds     bool     `codec:"cmds" json:"cmds"`
	Mentions bool     `codec:"mentions" json:"mentions"`
	Triggers []string `codec:"triggers" json:"triggers"`
	Convs    []string `codec:"convs" json:"convs"`
}

func (o TeamBotSettings) DeepCopy() TeamBotSettings {
	return TeamBotSettings{
		Cmds:     o.Cmds,
		Mentions: o.Mentions,
		Triggers: (func(x []string) []string {
			if x == nil {
				return nil
			}
			ret := make([]string, len(x))
			for i, v := range x {
				vCopy := v
				ret[i] = vCopy
			}
			return ret
		})(o.Triggers),
		Convs: (func(x []string) []string {
			if x == nil {
				return nil
			}
			ret := make([]string, len(x))
			for i, v := range x {
				vCopy := v
				ret[i] = vCopy
			}
			return ret
		})(o.Convs),
	}
}

type TeamRequestAccessResult struct {
	Open bool `codec:"open" json:"open"`
}

func (o TeamRequestAccessResult) DeepCopy() TeamRequestAccessResult {
	return TeamRequestAccessResult{
		Open: o.Open,
	}
}

type TeamAcceptOrRequestResult struct {
	WasToken    bool `codec:"wasToken" json:"wasToken"`
	WasSeitan   bool `codec:"wasSeitan" json:"wasSeitan"`
	WasTeamName bool `codec:"wasTeamName" json:"wasTeamName"`
	WasOpenTeam bool `codec:"wasOpenTeam" json:"wasOpenTeam"`
}

func (o TeamAcceptOrRequestResult) DeepCopy() TeamAcceptOrRequestResult {
	return TeamAcceptOrRequestResult{
		WasToken:    o.WasToken,
		WasSeitan:   o.WasSeitan,
		WasTeamName: o.WasTeamName,
		WasOpenTeam: o.WasOpenTeam,
	}
}

type TeamShowcase struct {
	IsShowcased       bool    `codec:"isShowcased" json:"is_showcased"`
	Description       *string `codec:"description,omitempty" json:"description,omitempty"`
	SetByUID          *UID    `codec:"setByUID,omitempty" json:"set_by_uid,omitempty"`
	AnyMemberShowcase bool    `codec:"anyMemberShowcase" json:"any_member_showcase"`
}

func (o TeamShowcase) DeepCopy() TeamShowcase {
	return TeamShowcase{
		IsShowcased: o.IsShowcased,
		Description: (func(x *string) *string {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.Description),
		SetByUID: (func(x *UID) *UID {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.SetByUID),
		AnyMemberShowcase: o.AnyMemberShowcase,
	}
}

type TeamAndMemberShowcase struct {
	TeamShowcase      TeamShowcase `codec:"teamShowcase" json:"teamShowcase"`
	IsMemberShowcased bool         `codec:"isMemberShowcased" json:"isMemberShowcased"`
}

func (o TeamAndMemberShowcase) DeepCopy() TeamAndMemberShowcase {
	return TeamAndMemberShowcase{
		TeamShowcase:      o.TeamShowcase.DeepCopy(),
		IsMemberShowcased: o.IsMemberShowcased,
	}
}

type TeamAvatar struct {
	AvatarFilename string         `codec:"avatarFilename" json:"avatarFilename"`
	Crop           *ImageCropRect `codec:"crop,omitempty" json:"crop,omitempty"`
}

func (o TeamAvatar) DeepCopy() TeamAvatar {
	return TeamAvatar{
		AvatarFilename: o.AvatarFilename,
		Crop: (func(x *ImageCropRect) *ImageCropRect {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Crop),
	}
}

type TeamCreateFancyInfo struct {
	Name               string         `codec:"name" json:"name"`
	Description        string         `codec:"description" json:"description"`
	JoinSubteam        bool           `codec:"joinSubteam" json:"joinSubteam"`
	OpenSettings       TeamSettings   `codec:"openSettings" json:"openSettings"`
	ProfileShowcase    bool           `codec:"profileShowcase" json:"profileShowcase"`
	Avatar             *TeamAvatar    `codec:"avatar,omitempty" json:"avatar,omitempty"`
	ChatChannels       []string       `codec:"chatChannels" json:"chatChannels"`
	Subteams           []string       `codec:"subteams" json:"subteams"`
	Users              []UserRolePair `codec:"users" json:"users"`
	EmailInviteMessage *string        `codec:"emailInviteMessage,omitempty" json:"emailInviteMessage,omitempty"`
}

func (o TeamCreateFancyInfo) DeepCopy() TeamCreateFancyInfo {
	return TeamCreateFancyInfo{
		Name:            o.Name,
		Description:     o.Description,
		JoinSubteam:     o.JoinSubteam,
		OpenSettings:    o.OpenSettings.DeepCopy(),
		ProfileShowcase: o.ProfileShowcase,
		Avatar: (func(x *TeamAvatar) *TeamAvatar {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Avatar),
		ChatChannels: (func(x []string) []string {
			if x == nil {
				return nil
			}
			ret := make([]string, len(x))
			for i, v := range x {
				vCopy := v
				ret[i] = vCopy
			}
			return ret
		})(o.ChatChannels),
		Subteams: (func(x []string) []string {
			if x == nil {
				return nil
			}
			ret := make([]string, len(x))
			for i, v := range x {
				vCopy := v
				ret[i] = vCopy
			}
			return ret
		})(o.Subteams),
		Users: (func(x []UserRolePair) []UserRolePair {
			if x == nil {
				return nil
			}
			ret := make([]UserRolePair, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Users),
		EmailInviteMessage: (func(x *string) *string {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.EmailInviteMessage),
	}
}

type UserRolePair struct {
	Assertion   string           `codec:"assertion" json:"assertion"`
	Role        TeamRole         `codec:"role" json:"role"`
	BotSettings *TeamBotSettings `codec:"botSettings,omitempty" json:"botSettings,omitempty"`
}

func (o UserRolePair) DeepCopy() UserRolePair {
	return UserRolePair{
		Assertion: o.Assertion,
		Role:      o.Role.DeepCopy(),
		BotSettings: (func(x *TeamBotSettings) *TeamBotSettings {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.BotSettings),
	}
}

type AssertionTeamMemberToRemove struct {
	Assertion         string `codec:"assertion" json:"assertion"`
	RemoveFromSubtree bool   `codec:"removeFromSubtree" json:"removeFromSubtree"`
}

func (o AssertionTeamMemberToRemove) DeepCopy() AssertionTeamMemberToRemove {
	return AssertionTeamMemberToRemove{
		Assertion:         o.Assertion,
		RemoveFromSubtree: o.RemoveFromSubtree,
	}
}

type InviteTeamMemberToRemove struct {
	InviteID TeamInviteID `codec:"inviteID" json:"inviteID"`
}

func (o InviteTeamMemberToRemove) DeepCopy() InviteTeamMemberToRemove {
	return InviteTeamMemberToRemove{
		InviteID: o.InviteID.DeepCopy(),
	}
}

type TeamMemberToRemoveType int

const (
	TeamMemberToRemoveType_ASSERTION TeamMemberToRemoveType = 0
	TeamMemberToRemoveType_INVITEID  TeamMemberToRemoveType = 1
)

func (o TeamMemberToRemoveType) DeepCopy() TeamMemberToRemoveType { return o }

var TeamMemberToRemoveTypeMap = map[string]TeamMemberToRemoveType{
	"ASSERTION": 0,
	"INVITEID":  1,
}

var TeamMemberToRemoveTypeRevMap = map[TeamMemberToRemoveType]string{
	0: "ASSERTION",
	1: "INVITEID",
}

func (e TeamMemberToRemoveType) String() string {
	if v, ok := TeamMemberToRemoveTypeRevMap[e]; ok {
		return v
	}
	return fmt.Sprintf("%v", int(e))
}

type TeamMemberToRemove struct {
	Type__      TeamMemberToRemoveType       `codec:"type" json:"type"`
	Assertion__ *AssertionTeamMemberToRemove `codec:"assertion,omitempty" json:"assertion,omitempty"`
	Inviteid__  *InviteTeamMemberToRemove    `codec:"inviteid,omitempty" json:"inviteid,omitempty"`
}

func (o *TeamMemberToRemove) Type() (ret TeamMemberToRemoveType, err error) {
	switch o.Type__ {
	case TeamMemberToRemoveType_ASSERTION:
		if o.Assertion__ == nil {
			err = errors.New("unexpected nil value for Assertion__")
			return ret, err
		}
	case TeamMemberToRemoveType_INVITEID:
		if o.Inviteid__ == nil {
			err = errors.New("unexpected nil value for Inviteid__")
			return ret, err
		}
	}
	return o.Type__, nil
}

func (o TeamMemberToRemove) Assertion() (res AssertionTeamMemberToRemove) {
	if o.Type__ != TeamMemberToRemoveType_ASSERTION {
		panic("wrong case accessed")
	}
	if o.Assertion__ == nil {
		return
	}
	return *o.Assertion__
}

func (o TeamMemberToRemove) Inviteid() (res InviteTeamMemberToRemove) {
	if o.Type__ != TeamMemberToRemoveType_INVITEID {
		panic("wrong case accessed")
	}
	if o.Inviteid__ == nil {
		return
	}
	return *o.Inviteid__
}

func NewTeamMemberToRemoveWithAssertion(v AssertionTeamMemberToRemove) TeamMemberToRemove {
	return TeamMemberToRemove{
		Type__:      TeamMemberToRemoveType_ASSERTION,
		Assertion__: &v,
	}
}

func NewTeamMemberToRemoveWithInviteid(v InviteTeamMemberToRemove) TeamMemberToRemove {
	return TeamMemberToRemove{
		Type__:     TeamMemberToRemoveType_INVITEID,
		Inviteid__: &v,
	}
}

func (o TeamMemberToRemove) DeepCopy() TeamMemberToRemove {
	return TeamMemberToRemove{
		Type__: o.Type__.DeepCopy(),
		Assertion__: (func(x *AssertionTeamMemberToRemove) *AssertionTeamMemberToRemove {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Assertion__),
		Inviteid__: (func(x *InviteTeamMemberToRemove) *InviteTeamMemberToRemove {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Inviteid__),
	}
}

type RemoveTeamMemberFailure struct {
	TeamMember     TeamMemberToRemove `codec:"teamMember" json:"teamMember"`
	ErrorAtTarget  *string            `codec:"errorAtTarget,omitempty" json:"errorAtTarget,omitempty"`
	ErrorAtSubtree *string            `codec:"errorAtSubtree,omitempty" json:"errorAtSubtree,omitempty"`
}

func (o RemoveTeamMemberFailure) DeepCopy() RemoveTeamMemberFailure {
	return RemoveTeamMemberFailure{
		TeamMember: o.TeamMember.DeepCopy(),
		ErrorAtTarget: (func(x *string) *string {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.ErrorAtTarget),
		ErrorAtSubtree: (func(x *string) *string {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.ErrorAtSubtree),
	}
}

type TeamRemoveMembersResult struct {
	Failures []RemoveTeamMemberFailure `codec:"failures" json:"failures"`
}

func (o TeamRemoveMembersResult) DeepCopy() TeamRemoveMembersResult {
	return TeamRemoveMembersResult{
		Failures: (func(x []RemoveTeamMemberFailure) []RemoveTeamMemberFailure {
			if x == nil {
				return nil
			}
			ret := make([]RemoveTeamMemberFailure, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Failures),
	}
}

type TeamEditMembersResult struct {
	Failures []UserRolePair `codec:"failures" json:"failures"`
}

func (o TeamEditMembersResult) DeepCopy() TeamEditMembersResult {
	return TeamEditMembersResult{
		Failures: (func(x []UserRolePair) []UserRolePair {
			if x == nil {
				return nil
			}
			ret := make([]UserRolePair, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Failures),
	}
}

type UntrustedTeamExistsResult struct {
	Exists bool       `codec:"exists" json:"exists"`
	Status StatusCode `codec:"status" json:"status"`
}

func (o UntrustedTeamExistsResult) DeepCopy() UntrustedTeamExistsResult {
	return UntrustedTeamExistsResult{
		Exists: o.Exists,
		Status: o.Status.DeepCopy(),
	}
}

type Invitelink struct {
	Ikey SeitanIKeyInvitelink `codec:"ikey" json:"ikey"`
	Url  string               `codec:"url" json:"url"`
}

func (o Invitelink) DeepCopy() Invitelink {
	return Invitelink{
		Ikey: o.Ikey.DeepCopy(),
		Url:  o.Url,
	}
}

type BulkRes struct {
	Malformed []string `codec:"malformed" json:"malformed"`
}

func (o BulkRes) DeepCopy() BulkRes {
	return BulkRes{
		Malformed: (func(x []string) []string {
			if x == nil {
				return nil
			}
			ret := make([]string, len(x))
			for i, v := range x {
				vCopy := v
				ret[i] = vCopy
			}
			return ret
		})(o.Malformed),
	}
}

type InviteLinkDetails struct {
	InviteID          TeamInviteID               `codec:"inviteID" json:"inviteID"`
	InviterResetOrDel bool                       `codec:"inviterResetOrDel" json:"inviterResetOrDel"`
	InviterUID        UID                        `codec:"inviterUID" json:"inviterUID"`
	InviterUsername   string                     `codec:"inviterUsername" json:"inviterUsername"`
	IsMember          bool                       `codec:"isMember" json:"isMember"`
	TeamAvatars       map[AvatarFormat]AvatarUrl `codec:"teamAvatars" json:"teamAvatars"`
	TeamDesc          string                     `codec:"teamDesc" json:"teamDesc"`
	TeamID            TeamID                     `codec:"teamID" json:"teamID"`
	TeamIsOpen        bool                       `codec:"teamIsOpen" json:"teamIsOpen"`
	TeamName          TeamName                   `codec:"teamName" json:"teamName"`
	TeamNumMembers    int                        `codec:"teamNumMembers" json:"teamNumMembers"`
}

func (o InviteLinkDetails) DeepCopy() InviteLinkDetails {
	return InviteLinkDetails{
		InviteID:          o.InviteID.DeepCopy(),
		InviterResetOrDel: o.InviterResetOrDel,
		InviterUID:        o.InviterUID.DeepCopy(),
		InviterUsername:   o.InviterUsername,
		IsMember:          o.IsMember,
		TeamAvatars: (func(x map[AvatarFormat]AvatarUrl) map[AvatarFormat]AvatarUrl {
			if x == nil {
				return nil
			}
			ret := make(map[AvatarFormat]AvatarUrl, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.TeamAvatars),
		TeamDesc:       o.TeamDesc,
		TeamID:         o.TeamID.DeepCopy(),
		TeamIsOpen:     o.TeamIsOpen,
		TeamName:       o.TeamName.DeepCopy(),
		TeamNumMembers: o.TeamNumMembers,
	}
}

type ImplicitTeamUserSet struct {
	KhulnasoftUsers    []string          `codec:"khulnasoftUsers" json:"khulnasoftUsers"`
	UnresolvedUsers []SocialAssertion `codec:"unresolvedUsers" json:"unresolvedUsers"`
}

func (o ImplicitTeamUserSet) DeepCopy() ImplicitTeamUserSet {
	return ImplicitTeamUserSet{
		KhulnasoftUsers: (func(x []string) []string {
			if x == nil {
				return nil
			}
			ret := make([]string, len(x))
			for i, v := range x {
				vCopy := v
				ret[i] = vCopy
			}
			return ret
		})(o.KhulnasoftUsers),
		UnresolvedUsers: (func(x []SocialAssertion) []SocialAssertion {
			if x == nil {
				return nil
			}
			ret := make([]SocialAssertion, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.UnresolvedUsers),
	}
}

// * iTeams
type ImplicitTeamDisplayName struct {
	IsPublic     bool                      `codec:"isPublic" json:"isPublic"`
	Writers      ImplicitTeamUserSet       `codec:"writers" json:"writers"`
	Readers      ImplicitTeamUserSet       `codec:"readers" json:"readers"`
	ConflictInfo *ImplicitTeamConflictInfo `codec:"conflictInfo,omitempty" json:"conflictInfo,omitempty"`
}

func (o ImplicitTeamDisplayName) DeepCopy() ImplicitTeamDisplayName {
	return ImplicitTeamDisplayName{
		IsPublic: o.IsPublic,
		Writers:  o.Writers.DeepCopy(),
		Readers:  o.Readers.DeepCopy(),
		ConflictInfo: (func(x *ImplicitTeamConflictInfo) *ImplicitTeamConflictInfo {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.ConflictInfo),
	}
}

type ConflictGeneration int

func (o ConflictGeneration) DeepCopy() ConflictGeneration {
	return o
}

type ImplicitTeamConflictInfo struct {
	Generation ConflictGeneration `codec:"generation" json:"generation"`
	Time       Time               `codec:"time" json:"time"`
}

func (o ImplicitTeamConflictInfo) DeepCopy() ImplicitTeamConflictInfo {
	return ImplicitTeamConflictInfo{
		Generation: o.Generation.DeepCopy(),
		Time:       o.Time.DeepCopy(),
	}
}

type LookupImplicitTeamRes struct {
	TeamID      TeamID                  `codec:"teamID" json:"teamID"`
	Name        TeamName                `codec:"name" json:"name"`
	DisplayName ImplicitTeamDisplayName `codec:"displayName" json:"displayName"`
	TlfID       TLFID                   `codec:"tlfID" json:"tlfID"`
}

func (o LookupImplicitTeamRes) DeepCopy() LookupImplicitTeamRes {
	return LookupImplicitTeamRes{
		TeamID:      o.TeamID.DeepCopy(),
		Name:        o.Name.DeepCopy(),
		DisplayName: o.DisplayName.DeepCopy(),
		TlfID:       o.TlfID.DeepCopy(),
	}
}

type TeamOperation struct {
	ManageMembers          bool `codec:"manageMembers" json:"manageMembers"`
	ManageSubteams         bool `codec:"manageSubteams" json:"manageSubteams"`
	CreateChannel          bool `codec:"createChannel" json:"createChannel"`
	Chat                   bool `codec:"chat" json:"chat"`
	DeleteChannel          bool `codec:"deleteChannel" json:"deleteChannel"`
	RenameChannel          bool `codec:"renameChannel" json:"renameChannel"`
	RenameTeam             bool `codec:"renameTeam" json:"renameTeam"`
	EditChannelDescription bool `codec:"editChannelDescription" json:"editChannelDescription"`
	EditTeamDescription    bool `codec:"editTeamDescription" json:"editTeamDescription"`
	SetTeamShowcase        bool `codec:"setTeamShowcase" json:"setTeamShowcase"`
	SetMemberShowcase      bool `codec:"setMemberShowcase" json:"setMemberShowcase"`
	SetRetentionPolicy     bool `codec:"setRetentionPolicy" json:"setRetentionPolicy"`
	SetMinWriterRole       bool `codec:"setMinWriterRole" json:"setMinWriterRole"`
	ChangeOpenTeam         bool `codec:"changeOpenTeam" json:"changeOpenTeam"`
	LeaveTeam              bool `codec:"leaveTeam" json:"leaveTeam"`
	JoinTeam               bool `codec:"joinTeam" json:"joinTeam"`
	SetPublicityAny        bool `codec:"setPublicityAny" json:"setPublicityAny"`
	ListFirst              bool `codec:"listFirst" json:"listFirst"`
	ChangeTarsDisabled     bool `codec:"changeTarsDisabled" json:"changeTarsDisabled"`
	DeleteChatHistory      bool `codec:"deleteChatHistory" json:"deleteChatHistory"`
	DeleteOtherEmojis      bool `codec:"deleteOtherEmojis" json:"deleteOtherEmojis"`
	DeleteOtherMessages    bool `codec:"deleteOtherMessages" json:"deleteOtherMessages"`
	DeleteTeam             bool `codec:"deleteTeam" json:"deleteTeam"`
	PinMessage             bool `codec:"pinMessage" json:"pinMessage"`
	ManageBots             bool `codec:"manageBots" json:"manageBots"`
	ManageEmojis           bool `codec:"manageEmojis" json:"manageEmojis"`
}

func (o TeamOperation) DeepCopy() TeamOperation {
	return TeamOperation{
		ManageMembers:          o.ManageMembers,
		ManageSubteams:         o.ManageSubteams,
		CreateChannel:          o.CreateChannel,
		Chat:                   o.Chat,
		DeleteChannel:          o.DeleteChannel,
		RenameChannel:          o.RenameChannel,
		RenameTeam:             o.RenameTeam,
		EditChannelDescription: o.EditChannelDescription,
		EditTeamDescription:    o.EditTeamDescription,
		SetTeamShowcase:        o.SetTeamShowcase,
		SetMemberShowcase:      o.SetMemberShowcase,
		SetRetentionPolicy:     o.SetRetentionPolicy,
		SetMinWriterRole:       o.SetMinWriterRole,
		ChangeOpenTeam:         o.ChangeOpenTeam,
		LeaveTeam:              o.LeaveTeam,
		JoinTeam:               o.JoinTeam,
		SetPublicityAny:        o.SetPublicityAny,
		ListFirst:              o.ListFirst,
		ChangeTarsDisabled:     o.ChangeTarsDisabled,
		DeleteChatHistory:      o.DeleteChatHistory,
		DeleteOtherEmojis:      o.DeleteOtherEmojis,
		DeleteOtherMessages:    o.DeleteOtherMessages,
		DeleteTeam:             o.DeleteTeam,
		PinMessage:             o.PinMessage,
		ManageBots:             o.ManageBots,
		ManageEmojis:           o.ManageEmojis,
	}
}

type ProfileTeamLoadRes struct {
	LoadTimeNsec int64 `codec:"loadTimeNsec" json:"loadTimeNsec"`
}

func (o ProfileTeamLoadRes) DeepCopy() ProfileTeamLoadRes {
	return ProfileTeamLoadRes{
		LoadTimeNsec: o.LoadTimeNsec,
	}
}

type RotationType int

const (
	RotationType_VISIBLE RotationType = 0
	RotationType_HIDDEN  RotationType = 1
	RotationType_CLKR    RotationType = 2
)

func (o RotationType) DeepCopy() RotationType { return o }

var RotationTypeMap = map[string]RotationType{
	"VISIBLE": 0,
	"HIDDEN":  1,
	"CLKR":    2,
}

var RotationTypeRevMap = map[RotationType]string{
	0: "VISIBLE",
	1: "HIDDEN",
	2: "CLKR",
}

func (e RotationType) String() string {
	if v, ok := RotationTypeRevMap[e]; ok {
		return v
	}
	return fmt.Sprintf("%v", int(e))
}

type TeamDebugRes struct {
	Chain TeamSigChainState `codec:"chain" json:"chain"`
}

func (o TeamDebugRes) DeepCopy() TeamDebugRes {
	return TeamDebugRes{
		Chain: o.Chain.DeepCopy(),
	}
}

type TeamProfileAddEntry struct {
	TeamID         TeamID   `codec:"teamID" json:"teamID"`
	TeamName       TeamName `codec:"teamName" json:"teamName"`
	Open           bool     `codec:"open" json:"open"`
	DisabledReason string   `codec:"disabledReason" json:"disabledReason"`
}

func (o TeamProfileAddEntry) DeepCopy() TeamProfileAddEntry {
	return TeamProfileAddEntry{
		TeamID:         o.TeamID.DeepCopy(),
		TeamName:       o.TeamName.DeepCopy(),
		Open:           o.Open,
		DisabledReason: o.DisabledReason,
	}
}

type MemberEmail struct {
	Email string `codec:"email" json:"email"`
	Role  string `codec:"role" json:"role"`
}

func (o MemberEmail) DeepCopy() MemberEmail {
	return MemberEmail{
		Email: o.Email,
		Role:  o.Role,
	}
}

type MemberUsername struct {
	Username string `codec:"username" json:"username"`
	Role     string `codec:"role" json:"role"`
}

func (o MemberUsername) DeepCopy() MemberUsername {
	return MemberUsername{
		Username: o.Username,
		Role:     o.Role,
	}
}

type TeamRolePair struct {
	Role         TeamRole `codec:"role" json:"role"`
	ImplicitRole TeamRole `codec:"implicitRole" json:"implicit_role"`
}

func (o TeamRolePair) DeepCopy() TeamRolePair {
	return TeamRolePair{
		Role:         o.Role.DeepCopy(),
		ImplicitRole: o.ImplicitRole.DeepCopy(),
	}
}

type TeamRoleMapAndVersion struct {
	Teams   map[TeamID]TeamRolePair `codec:"teams" json:"teams"`
	Version UserTeamVersion         `codec:"version" json:"user_team_version"`
}

func (o TeamRoleMapAndVersion) DeepCopy() TeamRoleMapAndVersion {
	return TeamRoleMapAndVersion{
		Teams: (func(x map[TeamID]TeamRolePair) map[TeamID]TeamRolePair {
			if x == nil {
				return nil
			}
			ret := make(map[TeamID]TeamRolePair, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.Teams),
		Version: o.Version.DeepCopy(),
	}
}

type TeamRoleMapStored struct {
	Data     TeamRoleMapAndVersion `codec:"data" json:"data"`
	CachedAt Time                  `codec:"cachedAt" json:"cachedAt"`
}

func (o TeamRoleMapStored) DeepCopy() TeamRoleMapStored {
	return TeamRoleMapStored{
		Data:     o.Data.DeepCopy(),
		CachedAt: o.CachedAt.DeepCopy(),
	}
}

type UserTeamVersion int

func (o UserTeamVersion) DeepCopy() UserTeamVersion {
	return o
}

type UserTeamVersionUpdate struct {
	Version UserTeamVersion `codec:"version" json:"version"`
}

func (o UserTeamVersionUpdate) DeepCopy() UserTeamVersionUpdate {
	return UserTeamVersionUpdate{
		Version: o.Version.DeepCopy(),
	}
}

type AnnotatedTeam struct {
	TeamID                       TeamID                `codec:"teamID" json:"teamID"`
	Name                         string                `codec:"name" json:"name"`
	TransitiveSubteamsUnverified SubteamListResult     `codec:"transitiveSubteamsUnverified" json:"transitiveSubteamsUnverified"`
	Members                      []TeamMemberDetails   `codec:"members" json:"members"`
	Invites                      []AnnotatedTeamInvite `codec:"invites" json:"invites"`
	Settings                     TeamSettings          `codec:"settings" json:"settings"`
	KeyGeneration                PerTeamKeyGeneration  `codec:"keyGeneration" json:"keyGeneration"`
	Showcase                     TeamShowcase          `codec:"showcase" json:"showcase"`
	JoinRequests                 []TeamJoinRequest     `codec:"joinRequests" json:"joinRequests"`
	TarsDisabled                 bool                  `codec:"tarsDisabled" json:"tarsDisabled"`
}

func (o AnnotatedTeam) DeepCopy() AnnotatedTeam {
	return AnnotatedTeam{
		TeamID:                       o.TeamID.DeepCopy(),
		Name:                         o.Name,
		TransitiveSubteamsUnverified: o.TransitiveSubteamsUnverified.DeepCopy(),
		Members: (func(x []TeamMemberDetails) []TeamMemberDetails {
			if x == nil {
				return nil
			}
			ret := make([]TeamMemberDetails, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Members),
		Invites: (func(x []AnnotatedTeamInvite) []AnnotatedTeamInvite {
			if x == nil {
				return nil
			}
			ret := make([]AnnotatedTeamInvite, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Invites),
		Settings:      o.Settings.DeepCopy(),
		KeyGeneration: o.KeyGeneration.DeepCopy(),
		Showcase:      o.Showcase.DeepCopy(),
		JoinRequests: (func(x []TeamJoinRequest) []TeamJoinRequest {
			if x == nil {
				return nil
			}
			ret := make([]TeamJoinRequest, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.JoinRequests),
		TarsDisabled: o.TarsDisabled,
	}
}

type TeamTreeMembershipValue struct {
	Role     TeamRole `codec:"role" json:"role"`
	JoinTime *Time    `codec:"joinTime,omitempty" json:"joinTime,omitempty"`
	TeamID   TeamID   `codec:"teamID" json:"teamID"`
}

func (o TeamTreeMembershipValue) DeepCopy() TeamTreeMembershipValue {
	return TeamTreeMembershipValue{
		Role: o.Role.DeepCopy(),
		JoinTime: (func(x *Time) *Time {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.JoinTime),
		TeamID: o.TeamID.DeepCopy(),
	}
}

type TeamTreeMembershipStatus int

const (
	TeamTreeMembershipStatus_OK     TeamTreeMembershipStatus = 0
	TeamTreeMembershipStatus_ERROR  TeamTreeMembershipStatus = 1
	TeamTreeMembershipStatus_HIDDEN TeamTreeMembershipStatus = 2
)

func (o TeamTreeMembershipStatus) DeepCopy() TeamTreeMembershipStatus { return o }

var TeamTreeMembershipStatusMap = map[string]TeamTreeMembershipStatus{
	"OK":     0,
	"ERROR":  1,
	"HIDDEN": 2,
}

var TeamTreeMembershipStatusRevMap = map[TeamTreeMembershipStatus]string{
	0: "OK",
	1: "ERROR",
	2: "HIDDEN",
}

func (e TeamTreeMembershipStatus) String() string {
	if v, ok := TeamTreeMembershipStatusRevMap[e]; ok {
		return v
	}
	return fmt.Sprintf("%v", int(e))
}

type TeamTreeError struct {
	Message           string `codec:"message" json:"message"`
	WillSkipSubtree   bool   `codec:"willSkipSubtree" json:"willSkipSubtree"`
	WillSkipAncestors bool   `codec:"willSkipAncestors" json:"willSkipAncestors"`
}

func (o TeamTreeError) DeepCopy() TeamTreeError {
	return TeamTreeError{
		Message:           o.Message,
		WillSkipSubtree:   o.WillSkipSubtree,
		WillSkipAncestors: o.WillSkipAncestors,
	}
}

type TeamTreeMembershipResult struct {
	S__     TeamTreeMembershipStatus `codec:"s" json:"s"`
	Ok__    *TeamTreeMembershipValue `codec:"ok,omitempty" json:"ok,omitempty"`
	Error__ *TeamTreeError           `codec:"error,omitempty" json:"error,omitempty"`
}

func (o *TeamTreeMembershipResult) S() (ret TeamTreeMembershipStatus, err error) {
	switch o.S__ {
	case TeamTreeMembershipStatus_OK:
		if o.Ok__ == nil {
			err = errors.New("unexpected nil value for Ok__")
			return ret, err
		}
	case TeamTreeMembershipStatus_ERROR:
		if o.Error__ == nil {
			err = errors.New("unexpected nil value for Error__")
			return ret, err
		}
	}
	return o.S__, nil
}

func (o TeamTreeMembershipResult) Ok() (res TeamTreeMembershipValue) {
	if o.S__ != TeamTreeMembershipStatus_OK {
		panic("wrong case accessed")
	}
	if o.Ok__ == nil {
		return
	}
	return *o.Ok__
}

func (o TeamTreeMembershipResult) Error() (res TeamTreeError) {
	if o.S__ != TeamTreeMembershipStatus_ERROR {
		panic("wrong case accessed")
	}
	if o.Error__ == nil {
		return
	}
	return *o.Error__
}

func NewTeamTreeMembershipResultWithOk(v TeamTreeMembershipValue) TeamTreeMembershipResult {
	return TeamTreeMembershipResult{
		S__:  TeamTreeMembershipStatus_OK,
		Ok__: &v,
	}
}

func NewTeamTreeMembershipResultWithError(v TeamTreeError) TeamTreeMembershipResult {
	return TeamTreeMembershipResult{
		S__:     TeamTreeMembershipStatus_ERROR,
		Error__: &v,
	}
}

func NewTeamTreeMembershipResultWithHidden() TeamTreeMembershipResult {
	return TeamTreeMembershipResult{
		S__: TeamTreeMembershipStatus_HIDDEN,
	}
}

func (o TeamTreeMembershipResult) DeepCopy() TeamTreeMembershipResult {
	return TeamTreeMembershipResult{
		S__: o.S__.DeepCopy(),
		Ok__: (func(x *TeamTreeMembershipValue) *TeamTreeMembershipValue {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Ok__),
		Error__: (func(x *TeamTreeError) *TeamTreeError {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Error__),
	}
}

type TeamTreeMembership struct {
	TeamName       string                   `codec:"teamName" json:"teamName"`
	Result         TeamTreeMembershipResult `codec:"result" json:"result"`
	TargetTeamID   TeamID                   `codec:"targetTeamID" json:"targetTeamID"`
	TargetUsername string                   `codec:"targetUsername" json:"targetUsername"`
	Guid           int                      `codec:"guid" json:"guid"`
}

func (o TeamTreeMembership) DeepCopy() TeamTreeMembership {
	return TeamTreeMembership{
		TeamName:       o.TeamName,
		Result:         o.Result.DeepCopy(),
		TargetTeamID:   o.TargetTeamID.DeepCopy(),
		TargetUsername: o.TargetUsername,
		Guid:           o.Guid,
	}
}

type TeamTreeMembershipsDoneResult struct {
	ExpectedCount  int    `codec:"expectedCount" json:"expectedCount"`
	TargetTeamID   TeamID `codec:"targetTeamID" json:"targetTeamID"`
	TargetUsername string `codec:"targetUsername" json:"targetUsername"`
	Guid           int    `codec:"guid" json:"guid"`
}

func (o TeamTreeMembershipsDoneResult) DeepCopy() TeamTreeMembershipsDoneResult {
	return TeamTreeMembershipsDoneResult{
		ExpectedCount:  o.ExpectedCount,
		TargetTeamID:   o.TargetTeamID.DeepCopy(),
		TargetUsername: o.TargetUsername,
		Guid:           o.Guid,
	}
}

type TeamTreeInitial struct {
	Guid int `codec:"guid" json:"guid"`
}

func (o TeamTreeInitial) DeepCopy() TeamTreeInitial {
	return TeamTreeInitial{
		Guid: o.Guid,
	}
}
