// Code generated by astool. DO NOT EDIT.

package streams

import (
	typeemoji "github.com/go-fed/activity/streams/impl/toot/type_emoji"
	typehashtag "github.com/go-fed/activity/streams/impl/toot/type_hashtag"
	typeidentityproof "github.com/go-fed/activity/streams/impl/toot/type_identityproof"
	vocab "github.com/go-fed/activity/streams/vocab"
)

// TootTootEmojiExtends returns true if Emoji extends from the other's type.
func TootTootEmojiExtends(other vocab.Type) bool {
	return typeemoji.TootEmojiExtends(other)
}

// TootTootHashtagExtends returns true if Hashtag extends from the other's type.
func TootTootHashtagExtends(other vocab.Type) bool {
	return typehashtag.TootHashtagExtends(other)
}

// TootTootIdentityProofExtends returns true if IdentityProof extends from the
// other's type.
func TootTootIdentityProofExtends(other vocab.Type) bool {
	return typeidentityproof.TootIdentityProofExtends(other)
}
