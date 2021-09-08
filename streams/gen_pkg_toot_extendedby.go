// Code generated by astool. DO NOT EDIT.

package streams

import (
	typeemoji "github.com/go-fed/activity/streams/impl/toot/type_emoji"
	typehashtag "github.com/go-fed/activity/streams/impl/toot/type_hashtag"
	typeidentityproof "github.com/go-fed/activity/streams/impl/toot/type_identityproof"
	vocab "github.com/go-fed/activity/streams/vocab"
)

// TootEmojiIsExtendedBy returns true if the other's type extends from Emoji. Note
// that it returns false if the types are the same; see the "IsOrExtends"
// variant instead.
func TootEmojiIsExtendedBy(other vocab.Type) bool {
	return typeemoji.EmojiIsExtendedBy(other)
}

// TootHashtagIsExtendedBy returns true if the other's type extends from Hashtag.
// Note that it returns false if the types are the same; see the "IsOrExtends"
// variant instead.
func TootHashtagIsExtendedBy(other vocab.Type) bool {
	return typehashtag.HashtagIsExtendedBy(other)
}

// TootIdentityProofIsExtendedBy returns true if the other's type extends from
// IdentityProof. Note that it returns false if the types are the same; see
// the "IsOrExtends" variant instead.
func TootIdentityProofIsExtendedBy(other vocab.Type) bool {
	return typeidentityproof.IdentityProofIsExtendedBy(other)
}
