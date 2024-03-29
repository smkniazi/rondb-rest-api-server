/*
 * This file is part of the RonDB REST API Server
 * Copyright (c) 2022 Hopsworks AB
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, version 3.
 *
 * This program is distributed in the hope that it will be useful, but
 * WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
 * General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program. If not, see <http://www.gnu.org/licenses/>.
 */

package pkread

import (
	"net/http"
	"testing"

	_ "github.com/ianlancetaylor/cgosymbolizer"
	"hopsworks.ai/rdrs/internal/common"
	ds "hopsworks.ai/rdrs/internal/datastructs"
	tu "hopsworks.ai/rdrs/internal/router/handler/utils"
)

// INT TESTS
// Test signed and unsigned int data type
func TestDataTypesInt(t *testing.T) {

	testTable := "int_table"
	testDb := "DB004"
	validateColumns := []interface{}{"col0", "col1"}
	tests := map[string]ds.PKTestInfo{
		"notfound": {
			PkReq: ds.PKReadBody{Filters: tu.NewFiltersKVs("id0", 100, "id1", 100),
				ReadColumns: tu.NewReadColumns("col", 2),
				OperationID: tu.NewOperationID(64),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusNotFound,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"simple1": {
			PkReq: ds.PKReadBody{Filters: tu.NewFiltersKVs("id0", 0, "id1", 0),
				ReadColumns: tu.NewReadColumns("col", 2),
				OperationID: tu.NewOperationID(64),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"simple2": { //with out operation ID
			PkReq: ds.PKReadBody{Filters: tu.NewFiltersKVs("id0", 0, "id1", 0),
				ReadColumns: tu.NewReadColumns("col", 2),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"simple3": { //without read columns.
			PkReq:        ds.PKReadBody{Filters: tu.NewFiltersKVs("id0", 0, "id1", 0)},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"simple4": { //Table with only primary keys
			PkReq: ds.PKReadBody{Filters: tu.NewFiltersKVs("id0", 0, "id1", 0),
				OperationID: tu.NewOperationID(64),
			},
			Table:        "int_table1",
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      []interface{}{},
		},

		"maxValues": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", 2147483647, "id1", 4294967295),
				ReadColumns: tu.NewReadColumns("col", 2),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"minValues": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", -2147483648, "id1", 0),
				ReadColumns: tu.NewReadColumns("col", 2),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"assignNegativeValToUnsignedCol": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", 1, "id1", -1), //id1 is unsigned
				ReadColumns: tu.NewReadColumns("col", 2),
				OperationID: tu.NewOperationID(64),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusBadRequest,
			BodyContains: common.ERROR_015(),
			RespKVs:      []interface{}{},
		},

		"assigningBiggerVals": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", 2147483648, "id1", 4294967295), //bigger than the range
				ReadColumns: tu.NewReadColumns("col", 2),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusBadRequest,
			BodyContains: common.ERROR_015(),
			RespKVs:      []interface{}{},
		},

		"assigningSmallerVals": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", -2147483649, "id1", 0), //smaller than range
				ReadColumns: tu.NewReadColumns("col", 2),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusBadRequest,
			BodyContains: common.ERROR_015(),
			RespKVs:      []interface{}{},
		},

		"nullVals": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", 1, "id1", 1),
				ReadColumns: tu.NewReadColumns("col", 2),
				OperationID: tu.NewOperationID(64),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},
	}

	tu.PkTest(t, tests, false, RegisterPKTestHandler)
}

func TestDataTypesBigInt(t *testing.T) {

	testTable := "bigint_table"
	testDb := "DB005"

	validateColumns := []interface{}{"col0", "col1"}
	tests := map[string]ds.PKTestInfo{

		"simple": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", 0, "id1", 0),
				ReadColumns: tu.NewReadColumns("col", 2),
				OperationID: tu.NewOperationID(64),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"maxValues": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", 9223372036854775807, "id1", uint64(18446744073709551615)),
				ReadColumns: tu.NewReadColumns("col", 2),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"minValues": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", -9223372036854775808, "id1", 0),
				ReadColumns: tu.NewReadColumns("col", 2),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"assignNegativeValToUnsignedCol": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", 0, "id1", -1), //id1 is unsigned
				ReadColumns: tu.NewReadColumns("col", 2),
				OperationID: tu.NewOperationID(64),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusBadRequest,
			BodyContains: common.ERROR_015(),
			RespKVs:      validateColumns,
		},

		"assigningBiggerVals": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", 9223372036854775807, "id1", "18446744073709551616"), //18446744073709551615+1
				ReadColumns: tu.NewReadColumns("col", 2),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusBadRequest,
			BodyContains: common.ERROR_015(),
			RespKVs:      validateColumns,
		},

		"assigningSmallerVals": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", "-9223372036854775809", "id1", 0), //-9223372036854775808-1
				ReadColumns: tu.NewReadColumns("col", 2),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusBadRequest,
			BodyContains: common.ERROR_015(),
			RespKVs:      validateColumns,
		},

		"nullVals": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", 1, "id1", 1),
				ReadColumns: tu.NewReadColumns("col", 2),
				OperationID: tu.NewOperationID(64),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},
	}
	tu.PkTest(t, tests, false, RegisterPKTestHandler)
}

func TestDataTypesTinyInt(t *testing.T) {

	testTable := "tinyint_table"
	testDb := "DB006"
	validateColumns := []interface{}{"col0", "col1"}
	tests := map[string]ds.PKTestInfo{

		"simple": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", 0, "id1", 0),
				ReadColumns: tu.NewReadColumns("col", 2),
				OperationID: tu.NewOperationID(64),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"maxValues": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", 127, "id1", 255),
				ReadColumns: tu.NewReadColumns("col", 2),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"minValues": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", -128, "id1", 0),
				ReadColumns: tu.NewReadColumns("col", 2),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"assignNegativeValToUnsignedCol": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", 0, "id1", -1), //id1 is unsigned
				ReadColumns: tu.NewReadColumns("col", 2),
				OperationID: tu.NewOperationID(64),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusBadRequest,
			BodyContains: common.ERROR_015(),
			RespKVs:      validateColumns,
		},

		"assigningBiggerVals": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", 127, "id1", 256), //255+1
				ReadColumns: tu.NewReadColumns("col", 2),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusBadRequest,
			BodyContains: common.ERROR_015(),
			RespKVs:      validateColumns,
		},

		"assigningSmallerVals": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", -129, "id1", 0), //-128-1
				ReadColumns: tu.NewReadColumns("col", 2),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusBadRequest,
			BodyContains: common.ERROR_015(),
			RespKVs:      validateColumns,
		},

		"nullVals": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", 1, "id1", 1),
				ReadColumns: tu.NewReadColumns("col", 2),
				OperationID: tu.NewOperationID(64),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},
	}
	tu.PkTest(t, tests, false, RegisterPKTestHandler)
}

func TestDataTypesSmallInt(t *testing.T) {

	testTable := "smallint_table"
	testDb := "DB007"
	validateColumns := []interface{}{"col0", "col1"}
	tests := map[string]ds.PKTestInfo{

		"simple": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", 0, "id1", 0),
				ReadColumns: tu.NewReadColumns("col", 2),
				OperationID: tu.NewOperationID(64),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"maxValues": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", 32767, "id1", 65535),
				ReadColumns: tu.NewReadColumns("col", 2),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"minValues": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", -32768, "id1", 0),
				ReadColumns: tu.NewReadColumns("col", 2),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"assignNegativeValToUnsignedCol": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", 0, "id1", -1), //id1 is unsigned
				ReadColumns: tu.NewReadColumns("col", 2),
				OperationID: tu.NewOperationID(64),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusBadRequest,
			BodyContains: common.ERROR_015(),
			RespKVs:      validateColumns,
		},

		"assigningBiggerVals": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", 32768, "id1", 256), //32767+1
				ReadColumns: tu.NewReadColumns("col", 2),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusBadRequest,
			BodyContains: common.ERROR_015(),
			RespKVs:      validateColumns,
		},

		"assigningSmallerVals": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", -32769, "id1", 0), //-32768-1
				ReadColumns: tu.NewReadColumns("col", 2),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusBadRequest,
			BodyContains: common.ERROR_015(),
			RespKVs:      validateColumns,
		},

		"nullVals": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", 1, "id1", 1),
				ReadColumns: tu.NewReadColumns("col", 2),
				OperationID: tu.NewOperationID(64),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},
	}
	tu.PkTest(t, tests, false, RegisterPKTestHandler)
}

func TestDataTypesMediumInt(t *testing.T) {

	testTable := "mediumint_table"
	testDb := "DB008"
	validateColumns := []interface{}{"col0", "col1"}
	tests := map[string]ds.PKTestInfo{

		"simple": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", 0, "id1", 0),
				ReadColumns: tu.NewReadColumns("col", 2),
				OperationID: tu.NewOperationID(64),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"maxValues": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", 8388607, "id1", 16777215),
				ReadColumns: tu.NewReadColumns("col", 2),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"minValues": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", -8388608, "id1", 0),
				ReadColumns: tu.NewReadColumns("col", 2),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"assignNegativeValToUnsignedCol": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", 0, "id1", -1), //id1 is unsigned
				ReadColumns: tu.NewReadColumns("col", 2),
				OperationID: tu.NewOperationID(64),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusBadRequest,
			BodyContains: common.ERROR_015(),
			RespKVs:      validateColumns,
		},

		"assigningBiggerVals": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", 8388608, "id1", 256), //8388607+1
				ReadColumns: tu.NewReadColumns("col", 2),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusBadRequest,
			BodyContains: common.ERROR_015(),
			RespKVs:      validateColumns,
		},

		"assigningSmallerVals": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", -8388609, "id1", 0), //-8388608-1
				ReadColumns: tu.NewReadColumns("col", 2),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusBadRequest,
			BodyContains: common.ERROR_015(),
			RespKVs:      validateColumns,
		},

		"nullVals": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", 1, "id1", 1),
				ReadColumns: tu.NewReadColumns("col", 2),
				OperationID: tu.NewOperationID(64),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},
	}
	tu.PkTest(t, tests, false, RegisterPKTestHandler)
}

func TestDataTypesFloat(t *testing.T) {

	// testTable := "float_table"
	testDb := "DB009"
	validateColumns := []interface{}{"col0", "col1"}
	tests := map[string]ds.PKTestInfo{

		"floatPK": { // NDB does not support floats PKs
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", 0),
				ReadColumns: tu.NewReadColumns("col", 2),
				OperationID: tu.NewOperationID(64),
			},
			Table:        "float_table2",
			Db:           testDb,
			HttpCode:     http.StatusBadRequest,
			BodyContains: common.ERROR_017(),
		},

		"simple": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", 0),
				ReadColumns: tu.NewReadColumns("col", 2),
			},
			Table:        "float_table1",
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"simple2": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", "1"),
				ReadColumns: tu.NewReadColumns("col", 2),
			},
			Table:        "float_table1",
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"nullVals": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", 2),
				ReadColumns: tu.NewReadColumns("col", 2),
				OperationID: tu.NewOperationID(64),
			},
			Table:        "float_table1",
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},
	}
	tu.PkTest(t, tests, false, RegisterPKTestHandler)
}

func TestDataTypesDouble(t *testing.T) {

	// testTable := "float_table"
	testDb := "DB010"
	validateColumns := []interface{}{"col0", "col1"}
	tests := map[string]ds.PKTestInfo{

		"floatPK": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", 0),
				ReadColumns: tu.NewReadColumns("col", 2),
				OperationID: tu.NewOperationID(64),
			},
			Table:        "double_table2",
			Db:           testDb,
			HttpCode:     http.StatusBadRequest,
			BodyContains: common.ERROR_017(),
		},

		"simple": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", 0),
				ReadColumns: tu.NewReadColumns("col", 2),
			},
			Table:        "double_table1",
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"simple2": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", 1),
				ReadColumns: tu.NewReadColumns("col", 2),
			},
			Table:        "double_table1",
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"nullVals": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", 2),
				ReadColumns: tu.NewReadColumns("col", 2),
				OperationID: tu.NewOperationID(64),
			},
			Table:        "double_table1",
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},
	}
	tu.PkTest(t, tests, false, RegisterPKTestHandler)
}

func TestDataTypesDecimal(t *testing.T) {

	testTable := "decimal_table"
	testDb := "DB011"
	validateColumns := []interface{}{"col0", "col1"}
	tests := map[string]ds.PKTestInfo{

		"simple": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", -12345.12345, "id1", 12345.12345),
				ReadColumns: tu.NewReadColumns("col", 2),
				OperationID: tu.NewOperationID(64),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"nullVals": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", -67890.12345, "id1", 67890.12345),
				ReadColumns: tu.NewReadColumns("col", 2),
				OperationID: tu.NewOperationID(64),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"assignNegativeValToUnsignedCol": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", -12345.12345, "id1", -12345.12345),
				ReadColumns: tu.NewReadColumns("col", 2),
				OperationID: tu.NewOperationID(64),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusBadRequest,
			BodyContains: common.ERROR_015(),
			RespKVs:      validateColumns,
		},

		"assigningBiggerVals": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", -12345.12345, "id1", 123456789.12345),
				ReadColumns: tu.NewReadColumns("col", 2),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusBadRequest,
			BodyContains: common.ERROR_015(),
			RespKVs:      validateColumns,
		},
	}
	tu.PkTest(t, tests, false, RegisterPKTestHandler)
}

func TestDataTypesBlobs(t *testing.T) {

	testDb := "DB013"
	tests := map[string]ds.PKTestInfo{

		"blob1": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", "1"),
				ReadColumns: tu.NewReadColumns("col", 2),
				OperationID: tu.NewOperationID(5),
			},
			Table:        "blob_table",
			Db:           testDb,
			HttpCode:     http.StatusInternalServerError,
			BodyContains: common.ERROR_026(),
			RespKVs:      []interface{}{},
		},

		"blob2": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", "1"),
				ReadColumns: tu.NewReadColumn("col1"),
				OperationID: tu.NewOperationID(5),
			},
			Table:        "blob_table",
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      []interface{}{"col1"},
		},

		"text1": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", "1"),
				ReadColumns: tu.NewReadColumns("col", 2),
				OperationID: tu.NewOperationID(5),
			},
			Table:        "text_table",
			Db:           testDb,
			HttpCode:     http.StatusInternalServerError,
			BodyContains: "",
			RespKVs:      []interface{}{},
		},

		"text2": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", "1"),
				ReadColumns: tu.NewReadColumn("col1"),
				OperationID: tu.NewOperationID(5),
			},
			Table:        "text_table",
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      []interface{}{"col1"},
		},
	}

	tu.PkTest(t, tests, false, RegisterPKTestHandler)
}

func TestDataTypesChar(t *testing.T) {
	ArrayColumnTest(t, "table1", "DB012", false, 100, true)
}

func TestDataTypesVarchar(t *testing.T) {
	ArrayColumnTest(t, "table1", "DB014", false, 50, false)
}

func TestDataTypesLongVarchar(t *testing.T) {
	ArrayColumnTest(t, "table1", "DB015", false, 256, false)
}

func TestDataTypesBinary(t *testing.T) {
	ArrayColumnTest(t, "table1", "DB016", true, 100, true)
}

func TestDataTypesVarbinary(t *testing.T) {
	ArrayColumnTest(t, "table1", "DB017", true, 100, false)
}

func TestDataTypesLongVarbinary(t *testing.T) {
	ArrayColumnTest(t, "table1", "DB018", true, 256, false)
}

func ArrayColumnTest(t *testing.T, table string, database string, isBinary bool, colWidth int, padding bool) {
	t.Helper()
	testTable := table
	testDb := database
	validateColumns := []interface{}{"col0"}
	tests := map[string]ds.PKTestInfo{

		"notfound1": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", tu.Encode("-1", isBinary, colWidth, padding)),
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusNotFound,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"notfound2": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", tu.Encode(*tu.NewOperationID(colWidth*4 + 1), isBinary, colWidth, padding)),
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusBadRequest,
			BodyContains: common.ERROR_008(),
			RespKVs:      validateColumns,
		},

		"simple1": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", tu.Encode("1", isBinary, colWidth, padding)),
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"simple2": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", tu.Encode("2", isBinary, colWidth, padding)),
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"simple3": { // new line char in string
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", tu.Encode("3", isBinary, colWidth, padding)),
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"simple4": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", tu.Encode("4", isBinary, colWidth, padding)),
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"simple5": { //unicode pk
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", tu.Encode("这是一个测验", isBinary, colWidth, padding)),
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"nulltest": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", tu.Encode("5", isBinary, colWidth, padding)),
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"escapedChars": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", tu.Encode("6", isBinary, colWidth, padding)),
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},
	}

	tu.PkTest(t, tests, isBinary, RegisterPKTestHandler)
}

func TestDataTypesDateColumn(t *testing.T) {
	t.Helper()
	testTable := "date_table"
	testDb := "DB019"
	validateColumns := []interface{}{"col0"}
	tests := map[string]ds.PKTestInfo{

		"validpk1": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", "1111-11-11"),
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"validpk2": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", "1111-11-11 00:00:00"),
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"invalidpk": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", "1111-11-11 11:00:00"),
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusBadRequest,
			BodyContains: common.ERROR_008(),
			RespKVs:      []interface{}{},
		},

		"invalidpk2": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", "1111-11-11 00:00:00.123123"),
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusBadRequest,
			BodyContains: common.ERROR_008(),
			RespKVs:      []interface{}{},
		},

		"nulltest1": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", "1111-11-12"),
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"error": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", "1111-13-11"),
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusBadRequest,
			BodyContains: common.ERROR_027(),
			RespKVs:      validateColumns,
		},
	}
	tu.PkTest(t, tests, false, RegisterPKTestHandler)
}

func TestDataTypesDatetimeColumn(t *testing.T) {
	t.Helper()
	testDb := "DB020"
	validateColumns := []interface{}{"col0"}
	tests := map[string]ds.PKTestInfo{

		"validpk1_pre0": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", "1111-11-11 11:11:11"),
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        "date_table0",
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},
		"validpk1_pre3": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", "1111-11-11 11:11:11.123"),
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        "date_table3",
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},
		"validpk1_pre6": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", "1111-11-11 11:11:11.123456"),
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        "date_table6",
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"validpk2_pre0": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", "1111-11-11 11:11:11.123123"), // nanoseconds should be ignored
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        "date_table0",
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"validpk2_pre3": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", "1111-11-11 11:11:11.123000"),
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        "date_table3",
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"validpk2_pre6": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", "1111-11-11 -11:11:11.123456"), //-iv sign should be ignored
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        "date_table6",
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"nulltest_pre0": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", "1111-11-12 11:11:11"),
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        "date_table0",
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},
		"nulltest_pre3": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", "1111-11-12 11:11:11.123"),
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        "date_table3",
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},
		"nulltest_pre6": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", "1111-11-12 11:11:11.123456"),
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        "date_table6",
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"wrongdate_pre0": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", "1111-13-11 11:11:11"),
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        "date_table0",
			Db:           testDb,
			HttpCode:     http.StatusBadRequest,
			BodyContains: common.ERROR_027(),
			RespKVs:      validateColumns,
		},
	}
	tu.PkTest(t, tests, false, RegisterPKTestHandler)
}

func TestDataTypesTimeColumn(t *testing.T) {
	t.Helper()
	testDb := "DB021"
	validateColumns := []interface{}{"col0"}
	tests := map[string]ds.PKTestInfo{

		"validpk1_pre0": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", "11:11:11"),
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        "time_table0",
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},
		"validpk1_pre3": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", "11:11:11.123"),
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        "time_table3",
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},
		"validpk1_pre6": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", "11:11:11.123456"),
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        "time_table6",
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"validpk2_pre0": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", "11:11:11.123123"), // nanoseconds should be ignored
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        "time_table0",
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"validpk2_pre3": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", "11:11:11.123000"),
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        "time_table3",
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"nulltest_pre0": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", "12:11:11"),
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        "time_table0",
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},
		"nulltest_pre3": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", "12:11:11.123"),
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        "time_table3",
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},
		"nulltest_pre6": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", "12:11:11.123456"),
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        "time_table6",
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"wrongtime_pre0": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", "11:61:11"),
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        "time_table0",
			Db:           testDb,
			HttpCode:     http.StatusBadRequest,
			BodyContains: common.ERROR_027(),
			RespKVs:      validateColumns,
		},
	}
	tu.PkTest(t, tests, false, RegisterPKTestHandler)
}

func TestDataTypesTimestampColumn(t *testing.T) {
	t.Helper()
	testDb := "DB022"
	validateColumns := []interface{}{"col0"}
	tests := map[string]ds.PKTestInfo{

		"badts_1": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", "1111-11-11 11:11:11"),
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        "ts_table0",
			Db:           testDb,
			HttpCode:     http.StatusBadRequest,
			BodyContains: common.ERROR_027(),
			RespKVs:      validateColumns,
		},

		"badts_2": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", "1970-01-01 00:00:00"),
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        "ts_table0",
			Db:           testDb,
			HttpCode:     http.StatusBadRequest,
			BodyContains: common.ERROR_027(),
			RespKVs:      validateColumns,
		},

		"badts_3": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", "2038-01-19 03:14:08"),
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        "ts_table0",
			Db:           testDb,
			HttpCode:     http.StatusBadRequest,
			BodyContains: common.ERROR_027(),
			RespKVs:      validateColumns,
		},

		"validpk1_pre0": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", "2022-11-11 11:11:11"),
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        "ts_table0",
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},
		"validpk1_pre3": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", "2022-11-11 11:11:11.123"),
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        "ts_table3",
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},
		"validpk1_pre6": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", "2022-11-11 11:11:11.123456"),
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        "ts_table6",
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"validpk2_pre0": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", "2022-11-11 11:11:11.123123"), // nanoseconds should be ignored
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        "ts_table0",
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"validpk2_pre3": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", "2022-11-11 11:11:11.123000"),
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        "ts_table3",
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"validpk2_pre6": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", "2022-11-11 -11:11:11.123456"), //-iv sign should be ignored
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        "ts_table6",
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"nulltest_pre0": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", "2022-11-12 11:11:11"),
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        "ts_table0",
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},
		"nulltest_pre3": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", "2022-11-12 11:11:11.123"),
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        "ts_table3",
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},
		"nulltest_pre6": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", "2022-11-12 11:11:11.123456"),
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        "ts_table6",
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"wrongdate_pre0": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", "2022-13-11 11:11:11"),
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        "ts_table0",
			Db:           testDb,
			HttpCode:     http.StatusBadRequest,
			BodyContains: common.ERROR_027(),
			RespKVs:      validateColumns,
		},
	}
	tu.PkTest(t, tests, false, RegisterPKTestHandler)
}

func TestDataTypesYearColumn(t *testing.T) {
	///< Year 1901-2155 (1 byte)
	t.Helper()
	testDb := "DB023"
	testTable := "year_table"
	validateColumns := []interface{}{"col0"}
	tests := map[string]ds.PKTestInfo{

		"simple1": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", "2022"),
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"notfound1": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", "1901"),
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusNotFound,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"notfound2": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", "2155"),
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusNotFound,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"nulltest": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", "2023"),
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"baddate1": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", "1900"),
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusBadRequest,
			BodyContains: common.ERROR_015(),
			RespKVs:      validateColumns,
		},

		"baddate2": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", "2156"),
				ReadColumns: tu.NewReadColumns("col", 1),
				OperationID: tu.NewOperationID(5),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusBadRequest,
			BodyContains: common.ERROR_015(),
			RespKVs:      validateColumns,
		},
	}
	tu.PkTest(t, tests, false, RegisterPKTestHandler)
}

func TestDataTypesBitColumn(t *testing.T) {
	t.Helper()
	testDb := "DB024"
	testTable := "bit_table"
	validateColumns := []interface{}{"col0"}
	tests := map[string]ds.PKTestInfo{

		"simple1": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", tu.Encode("1", true, 100, true)),
				ReadColumns: tu.NewReadColumns("col", 5),
				OperationID: tu.NewOperationID(5),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},
		"simple2": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", tu.Encode("2", true, 100, true)),
				ReadColumns: tu.NewReadColumns("col", 5),
				OperationID: tu.NewOperationID(5),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},

		"null": {
			PkReq: ds.PKReadBody{
				Filters:     tu.NewFiltersKVs("id0", tu.Encode("3", true, 100, true)),
				ReadColumns: tu.NewReadColumns("col", 5),
				OperationID: tu.NewOperationID(5),
			},
			Table:        testTable,
			Db:           testDb,
			HttpCode:     http.StatusOK,
			BodyContains: "",
			RespKVs:      validateColumns,
		},
	}
	tu.PkTest(t, tests, true, RegisterPKTestHandler)
}
