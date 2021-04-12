package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/twpayne/chezmoi/v2/internal/chezmoitest"
)

func Test_lastpassParseNote(t *testing.T) {
	for _, tc := range []struct {
		note     string
		expected map[string]string
	}{
		{
			note: "Foo:bar\n",
			expected: map[string]string{
				"foo": "bar\n",
			},
		},
		{
			note: "Foo:bar\nbaz\n",
			expected: map[string]string{
				"foo": "bar\nbaz\n",
			},
		},
		{
			note: chezmoitest.JoinLines(
				"NoteType:SSH Key",
				"Language:en-US",
				"Bit Strength:2048",
				"Format:RSA",
				"Passphrase:Passphrase",
				"Private Key:-----BEGIN OPENSSH PRIVATE KEY-----",
				"-----END OPENSSH PRIVATE KEY-----",
				"Public Key:ssh-rsa public-key you@example",
				"Hostname:Hostname",
				"Date:Date",
			) + "Notes:",
			expected: map[string]string{
				"noteType":    "SSH Key\n",
				"language":    "en-US\n",
				"bitStrength": "2048\n",
				"format":      "RSA\n",
				"passphrase":  "Passphrase\n",
				"privateKey":  "-----BEGIN OPENSSH PRIVATE KEY-----\n-----END OPENSSH PRIVATE KEY-----\n",
				"publicKey":   "ssh-rsa public-key you@example\n",
				"hostname":    "Hostname\n",
				"date":        "Date\n",
				"notes":       "\n",
			},
		},
	} {
		assert.Equal(t, tc.expected, lastpassParseNote(tc.note))
	}
}
