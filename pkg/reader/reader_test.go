package reader

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ReadJSONSuite struct {
	suite.Suite
}

type TestData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestReadJSONSuite(t *testing.T) {
	suite.Run(t, new(ReadJSONSuite))
}

func (suite *ReadJSONSuite) TestReadJSONSuccess() {
	var testData TestData

	jsonData := `{"name": "John Doe", "age": 30}`

	reader := bytes.NewReader([]byte(jsonData))

	err := ReadJSON(reader, &testData)

	suite.NoError(err)

	expectedData := TestData{Name: "John Doe", Age: 30}
	suite.Equal(expectedData, testData)
}

func (suite *ReadJSONSuite) TestReadJSONError() {
	var testData TestData

	jsonData := `{"name": "John Doe", "age": "invalid"}`

	reader := bytes.NewReader([]byte(jsonData))

	err := ReadJSON(reader, &testData)

	suite.Error(err)
}
