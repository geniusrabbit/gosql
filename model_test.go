//
// @project GeniusRabbit
// @author Dmitry Ponomarev <demdxx@gmail.com> 2016, 2020
//

package gosql

import (
	"database/sql"
	"encoding/json"
	"testing"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

type jsonTestItem struct {
	JValue int `json:"json"`
}

// Model for test
type Model struct {
	ID                       uint64
	Char                     Char
	NumberArray              NumberArray[int]
	NullableIntArray         NullableNumberArray[int]
	OrderedIntArray          OrderedNumberArray[int]
	NullableOrderedIntArray  NullableOrderedNumberArray[int]
	UintArray                NumberArray[uint]
	NullableUintArray        NullableNumberArray[uint]
	OrderedUintArray         OrderedNumberArray[uint]
	NullableOrderedUintArray NullableOrderedNumberArray[uint]
	StringArray              StringArray
	NullableStringArray      NullableStringArray
	JSON                     JSON[jsonTestItem]
	NullableJSON             NullableJSON[int]
}

func selectIt(db *sql.DB) (m *Model, err error) {
	m = &Model{}
	var rows *sql.Rows

	rows, err = db.Query("SELECT id, char, int_array, nullable_int_array, ordered_int_array, nullable_ordered_int_array," +
		/**/ "uint_array, nullable_uint_array, ordered_uint_array, nullable_ordered_uint_array," +
		/**/ "string_array, nullable_string_array, json, nullable_json FROM models")

	if err == nil {
		defer rows.Close()

		for rows.Next() {
			_ = rows.Scan(&m.ID, &m.Char, &m.NumberArray, &m.NullableIntArray, &m.OrderedIntArray, &m.NullableOrderedIntArray,
				/* * */ &m.UintArray, &m.NullableUintArray, &m.OrderedUintArray, &m.NullableOrderedUintArray,
				/* * */ &m.StringArray, &m.NullableStringArray, &m.JSON, &m.NullableJSON)
			break
		}
	}
	return
}

// TestModel data
func TestModel(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	columns := []string{
		"id", "char", "int_array", "nullable_int_array", "ordered_int_array", "nullable_ordered_int_array",
		"uint_array", "nullable_uint_array", "ordered_uint_array", "nullable_ordered_uint_array",
		"string_array", "nullable_string_array", "json", "nullable_json",
	}

	// expects
	rows := sqlmock.NewRows(columns).
		AddRow(1, "M", "{9,3,4,6,1}", nil, "{9,3,4,6,1}", nil,
			"{9,3,4,6,1}", nil, "{9,3,4,6,1}", nil,
			"{a,c,b}", nil, "{\"json\":1}", nil)
	mock.ExpectQuery("SELECT (.+) FROM models").
		WillReturnRows(rows)

	// now we execute our method
	if m, err := selectIt(db); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	} else {
		if data, err := json.MarshalIndent(m, "", "    "); nil == err {
			t.Log(string(data))
		} else {
			t.Error(err)
		}
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}
