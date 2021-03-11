package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestUnknowOperation(t *testing.T) {
	testString := "20^20"
	_, err := calculateString(testString)
	if err == nil {
		t.Errorf("Test for check unknow operation was failed - unknow operation was processed, string with data - %s", testString)
	}

}
func TestProblemWithParsing(t *testing.T) {
	testString := "(20)20"
	_, err := calculateString(testString)
	if err == nil {
		t.Errorf("Test for check failed of parsing is bad, string with data - %s", testString)
	}
}

func TestEachOperation(t *testing.T) {
	listOperation := []string{"+", "-", "/", "*"}

	for _, v := range listOperation {
		testString := "20" + v + "20"
		_, err := calculateString(testString)
		if err != nil {
			t.Errorf("Test for operation '%s' was failed, string with data - '%s', error from function '%s'", v, testString, err)
		}
	}
}

func TestDivideZero(t *testing.T) {
	testString := "20/0"
	res, err := calculateString(testString)
	fmt.Println(err)
	if err == nil {
		t.Errorf("Test for zero, string with data - '%s', result from function '%s'", testString, strconv.FormatFloat(res, 'f', 6, 64))
	}
}

func TestAnyVarianOfExpr(t *testing.T) {
	listOperation := []string{"1+1", "(1+1)", "(1+1)*2", "(1+1)*2/4", "(1+1)*2/-4", "1.5/4"}

	for _, v := range listOperation {
		res, err := calculateString(v)
		switch v {
		case "1+1", "(1+1)":
			if res != 2 {
				t.Errorf("Test was failed, calculate string with data - '%s' is incorrect", v)
			}
		case "(1+1)*2":
			if res != 4 {
				t.Errorf("Test was failed, calculate string with data - '%s' is incorrect", v)
			}
		case "(1+1)*2/4":
			if res != 1 {
				t.Errorf("Test was failed, calculate string with data - '%s' is incorrect", v)
			}
		case "(1+1)*2/-4":
			if res != -1 {
				t.Errorf("Test was failed, calculate string with data - '%s' is incorrect", v)
			}
		case "1.5/4":
			if res != 0.375 {
				t.Errorf("Test was failed, calculate string with data - '%s' is incorrect", v)
			}
		}
		if err != nil {
			t.Errorf("Test was failed, string with data - '%s', error from function '%s'", v, err)
		}
	}

}

type TestCase struct {
	expr       string
	Response   string
	StatusCode int
}

func TestGetUser(t *testing.T) {
	cases := []TestCase{
		TestCase{
			expr:       "notParam=1+1",
			Response:   `requet has no expr param`,
			StatusCode: http.StatusOK,
		},
		TestCase{
			expr:       "expr=1+1",
			Response:   `result:2.000000`,
			StatusCode: http.StatusOK,
		},
		TestCase{
			expr:       "expr=2/0",
			Response:   `calculate was wrong:divide zero error`,
			StatusCode: http.StatusOK,
		},
		TestCase{
			expr:       "expr=1^2",
			Response:   `calculate was wrong:unknow type of operation '^'`,
			StatusCode: http.StatusOK,
		},
		TestCase{
			expr:       "expr=1^2&test=1111",
			Response:   `calculate was wrong:unknow type of operation '^'`,
			StatusCode: http.StatusOK,
		},
	}
	for caseNum, item := range cases {
		url := "http://example.com:80/?" + item.expr
		req := httptest.NewRequest("GET", url, nil)
		w := httptest.NewRecorder()
		handler(w, req)
		if w.Code != item.StatusCode {
			t.Errorf("[%d] wrong StatusCode: got %d, expected %d", caseNum, w.Code, item.StatusCode)
		}
		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		bodyStr := string(body)
		if bodyStr != item.Response {
			t.Errorf("[%d] wrong Response: got %+v, expected %+v", caseNum, bodyStr, item.Response)
		}
	}
}
