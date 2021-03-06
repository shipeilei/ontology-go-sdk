/*
 * Copyright (C) 2018 The ontology Authors
 * This file is part of The ontology library.
 *
 * The ontology is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The ontology is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The ontology.  If not, see <http://www.gnu.org/licenses/>.
 */
//Provide some utils for ontology-go-sdk
package utils

import (
	"encoding/hex"
	"fmt"
	sdkcom "github.com/ontio/ontology-go-sdk/common"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology/core/types"
	nvutils "github.com/ontio/ontology/smartcontract/service/native/utils"
	"math/big"
	"os"
	"strings"
)

func GetContractAddress(contractCode string) (common.Address, error) {
	code, err := hex.DecodeString(contractCode)
	if err != nil {
		return common.ADDRESS_EMPTY, fmt.Errorf("hex.DecodeString error:%s", err)
	}
	return types.AddressFromVmCode(code), nil
}

func GetAssetAddress(asset string) (common.Address, error) {
	var contractAddress common.Address
	switch strings.ToUpper(asset) {
	case "ONT":
		contractAddress = nvutils.OntContractAddress
	case "ONG":
		contractAddress = nvutils.OngContractAddress
	default:
		return common.ADDRESS_EMPTY, fmt.Errorf("asset:%s not equal ont or ong", asset)
	}
	return contractAddress, nil
}

//IsFileExist return is file is exist
func IsFileExist(file string) bool {
	_, err := os.Stat(file)
	return err == nil || os.IsExist(err)
}

//ParseNeoVMContractReturnType return value for result of smart contract execute code.
func ParseNeoVMContractReturnType(value interface{}, returnType sdkcom.NeoVMReturnType) (interface{}, error) {
	switch returnType {
	case sdkcom.NEOVM_TYPE_BOOL:
		return ParseNeoVMContractReturnTypeBool(value)
	case sdkcom.NEOVM_TYPE_INTEGER:
		return ParseNeoVMContractReturnTypeInteger(value)
	case sdkcom.NEOVM_TYPE_STRING:
		return ParseNeoVMContractReturnTypeString(value)
	case sdkcom.NEOVM_TYPE_BYTE_ARRAY:
		return ParseNeoVMContractReturnTypeByteArray(value)
	case sdkcom.NEOVM_TYPE_ARRAY:
		return value, nil
	}
	return value, nil
}

//ParseNeoVMContractReturnTypeBool return bool value of smart contract execute code.
func ParseNeoVMContractReturnTypeBool(val interface{}) (bool, error) {
	hexStr, ok := val.(string)
	if !ok {
		return false, fmt.Errorf("asset to string failed")
	}
	return hexStr == "01", nil
}

//ParseNeoVMContractReturnTypeInteger return integer value of smart contract execute code.
func ParseNeoVMContractReturnTypeInteger(val interface{}) (*big.Int, error) {
	hexStr, ok := val.(string)
	if !ok {
		return nil, fmt.Errorf("asset to string failed")
	}
	data, err := hex.DecodeString(hexStr)
	if err != nil {
		return nil, fmt.Errorf("hex.DecodeString error:%s", err)
	}
	return common.BigIntFromNeoBytes(data), nil
}

//ParseNeoVMContractReturnTypeByteArray return []byte value of smart contract execute code.
func ParseNeoVMContractReturnTypeByteArray(val interface{}) ([]byte, error) {
	hexStr, ok := val.(string)
	if !ok {
		return nil, fmt.Errorf("asset to string failed")
	}
	data, err := hex.DecodeString(hexStr)
	if err != nil {
		return nil, fmt.Errorf("hex.DecodeString error:%s", err)
	}
	return data, nil
}

//ParseNeoVMContractReturnTypeString return string value of smart contract execute code.
func ParseNeoVMContractReturnTypeString(val interface{}) (string, error) {
	data, err := ParseNeoVMContractReturnTypeByteArray(val)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

//BytesReverse return the reverse of []byte
func BytesReverse(u []byte) []byte {
	for i, j := 0, len(u)-1; i < j; i, j = i+1, j-1 {
		u[i], u[j] = u[j], u[i]
	}
	return u
}
