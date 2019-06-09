package caps

/*
	WARNING: this file is autogenerated by `make capdefs`
	DO NOT EDIT MANUALLY.
*/

const (
	// number of recognized capabilities:
	numCapabs = 27
	// length of the uint64 array that represents the bitset:
	bitsetLen = 1
)

const (
	// Acc is the proposed IRCv3 capability named "draft/acc":
	// https://github.com/ircv3/ircv3-specifications/pull/276
	Acc Capability = iota

	// AccountNotify is the IRCv3 capability named "account-notify":
	// https://ircv3.net/specs/extensions/account-notify-3.1.html
	AccountNotify Capability = iota

	// AccountTag is the IRCv3 capability named "account-tag":
	// https://ircv3.net/specs/extensions/account-tag-3.2.html
	AccountTag Capability = iota

	// AwayNotify is the IRCv3 capability named "away-notify":
	// https://ircv3.net/specs/extensions/away-notify-3.1.html
	AwayNotify Capability = iota

	// Batch is the IRCv3 capability named "batch":
	// https://ircv3.net/specs/extensions/batch-3.2.html
	Batch Capability = iota

	// CapNotify is the IRCv3 capability named "cap-notify":
	// https://ircv3.net/specs/extensions/cap-notify-3.2.html
	CapNotify Capability = iota

	// ChgHost is the IRCv3 capability named "chghost":
	// https://ircv3.net/specs/extensions/chghost-3.2.html
	ChgHost Capability = iota

	// EchoMessage is the IRCv3 capability named "echo-message":
	// https://ircv3.net/specs/extensions/echo-message-3.2.html
	EchoMessage Capability = iota

	// ExtendedJoin is the IRCv3 capability named "extended-join":
	// https://ircv3.net/specs/extensions/extended-join-3.1.html
	ExtendedJoin Capability = iota

	// InviteNotify is the IRCv3 capability named "invite-notify":
	// https://ircv3.net/specs/extensions/invite-notify-3.2.html
	InviteNotify Capability = iota

	// LabeledResponse is the draft IRCv3 capability named "draft/labeled-response":
	// https://ircv3.net/specs/extensions/labeled-response.html
	LabeledResponse Capability = iota

	// Languages is the proposed IRCv3 capability named "draft/languages":
	// https://gist.github.com/DanielOaks/8126122f74b26012a3de37db80e4e0c6
	Languages Capability = iota

	// MaxLine is the Oragono-specific capability named "oragono.io/maxline-2":
	// https://oragono.io/maxline-2
	MaxLine Capability = iota

	// MessageTags is the IRCv3 capability named "message-tags":
	// https://ircv3.net/specs/extensions/message-tags.html
	MessageTags Capability = iota

	// MultiPrefix is the IRCv3 capability named "multi-prefix":
	// https://ircv3.net/specs/extensions/multi-prefix-3.1.html
	MultiPrefix Capability = iota

	// Rename is the proposed IRCv3 capability named "draft/rename":
	// https://github.com/SaberUK/ircv3-specifications/blob/rename/extensions/rename.md
	Rename Capability = iota

	// Resume is the proposed IRCv3 capability named "draft/resume-0.5":
	// https://github.com/DanielOaks/ircv3-specifications/blob/master+resume/extensions/resume.md
	Resume Capability = iota

	// SASL is the IRCv3 capability named "sasl":
	// https://ircv3.net/specs/extensions/sasl-3.2.html
	SASL Capability = iota

	// ServerTime is the IRCv3 capability named "server-time":
	// https://ircv3.net/specs/extensions/server-time-3.2.html
	ServerTime Capability = iota

	// SetName is the proposed IRCv3 capability named "draft/setname":
	// https://github.com/ircv3/ircv3-specifications/pull/361
	SetName Capability = iota

	// STS is the IRCv3 capability named "sts":
	// https://ircv3.net/specs/extensions/sts.html
	STS Capability = iota

	// UserhostInNames is the IRCv3 capability named "userhost-in-names":
	// https://ircv3.net/specs/extensions/userhost-in-names-3.2.html
	UserhostInNames Capability = iota

	// Bouncer is the Oragono-specific capability named "oragono.io/bnc":
	// https://oragono.io/bnc
	Bouncer Capability = iota

	// ZNCSelfMessage is the ZNC vendor capability named "znc.in/self-message":
	// https://wiki.znc.in/Query_buffers
	ZNCSelfMessage Capability = iota

	// EventPlayback is the Proposed IRCv3 capability named "draft/event-playback":
	// https://github.com/ircv3/ircv3-specifications/pull/362
	EventPlayback Capability = iota

	// ZNCPlayback is the ZNC vendor capability named "znc.in/playback":
	// https://wiki.znc.in/Playback
	ZNCPlayback Capability = iota

	// Nope is the Oragono vendor capability named "oragono.io/nope":
	// https://oragono.io/nope
	Nope Capability = iota
)

// `capabilityNames[capab]` is the string name of the capability `capab`
var (
	capabilityNames = [numCapabs]string{
		"draft/acc",
		"account-notify",
		"account-tag",
		"away-notify",
		"batch",
		"cap-notify",
		"chghost",
		"echo-message",
		"extended-join",
		"invite-notify",
		"draft/labeled-response",
		"draft/languages",
		"oragono.io/maxline-2",
		"message-tags",
		"multi-prefix",
		"draft/rename",
		"draft/resume-0.5",
		"sasl",
		"server-time",
		"draft/setname",
		"sts",
		"userhost-in-names",
		"oragono.io/bnc",
		"znc.in/self-message",
		"draft/event-playback",
		"znc.in/playback",
		"oragono.io/nope",
	}
)
