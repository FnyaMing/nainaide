package types

import (
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/FnyaMing/nainaide/types"
)

func TestMsgUnjailGetSignBytes(t *testing.T) {
	addr := sdk.AccAddress("abcd")
	msg := NewMsgUnjail(sdk.ValAddress(addr))
	bytes := msg.GetSignBytes()
	require.Equal(
		t,
		`{"type":"cosmos-sdk/MsgUnjail","value":{"address":"nainaidevaloper1v93xxeqtyq23u"}}`,
		string(bytes),
	)
}
