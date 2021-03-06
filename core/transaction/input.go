package transaction

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/nknorg/nkn/common"
	"github.com/nknorg/nkn/common/serialization"
)

type UTXOTxInput struct {

	//Indicate the previous Tx which include the UTXO output for usage
	ReferTxID common.Uint256

	//The index of output in the referTx output list
	ReferTxOutputIndex uint16
}

func (ui *UTXOTxInput) Serialize(w io.Writer) error {
	_, err := ui.ReferTxID.Serialize(w)
	if err != nil {
		return err
	}
	err = serialization.WriteUint16(w, ui.ReferTxOutputIndex)
	if err != nil {
		return err
	}

	return nil
}

func (ui *UTXOTxInput) Deserialize(r io.Reader) error {
	//referTxID
	err := ui.ReferTxID.Deserialize(r)
	if err != nil {
		return err
	}

	//Output Index
	temp, err := serialization.ReadUint16(r)
	ui.ReferTxOutputIndex = uint16(temp)
	if err != nil {
		return err
	}

	return nil
}

func (ui *UTXOTxInput) Equal(ui2 *UTXOTxInput) bool {
	if ui.ReferTxID.CompareTo(ui2.ReferTxID) != 0 {
		return false
	}

	if ui.ReferTxOutputIndex != ui.ReferTxOutputIndex {
		return false
	}

	return true
}

func (ui *UTXOTxInput) MarshalJson() ([]byte, error) {

	inputInfo := UTXOTxInputInfo{
		ReferTxID:          common.BytesToHexString(ui.ReferTxID.ToArrayReverse()),
		ReferTxOutputIndex: ui.ReferTxOutputIndex,
	}

	data, err := json.Marshal(inputInfo)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (ui *UTXOTxInput) UnmarshalJson(data []byte) error {

	inputInfo := new(UTXOTxInputInfo)
	var err error
	if err = json.Unmarshal(data, &inputInfo); err != nil {
		return err
	}

	txid, err := common.HexStringToBytesReverse(inputInfo.ReferTxID)
	if err != nil {
		return err
	}
	ui.ReferTxID, err = common.Uint256ParseFromBytes(txid)
	if err != nil {
		return err
	}
	ui.ReferTxOutputIndex = inputInfo.ReferTxOutputIndex

	return nil
}

func (ui *UTXOTxInput) ToArray() []byte {
	b := new(bytes.Buffer)
	ui.Serialize(b)
	return b.Bytes()
}

func (ui *UTXOTxInput) ToString() string {
	return fmt.Sprintf("%x%x", ui.ReferTxID.ToString(), ui.ReferTxOutputIndex)
}

func (ui *UTXOTxInput) Equals(other *UTXOTxInput) bool {
	if ui == other {
		return true
	}
	if other == nil {
		return false
	}
	if ui.ReferTxID == other.ReferTxID && ui.ReferTxOutputIndex == other.ReferTxOutputIndex {
		return true
	} else {
		return false
	}
}
