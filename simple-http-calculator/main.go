package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	exprVal := queryParams.Get("expr")
	if exprVal == "" {
		w.WriteHeader(200)
		w.Write([]byte("requet has no expr param1"))
		return
	}
	exprVal = strings.Replace(exprVal, " ", "+", 1)

	res, err := calculateString(exprVal)
	if err != nil {
		w.WriteHeader(200)
		w.Write([]byte("calculate was wrong:" + err.Error()))
		return
	}
	w.WriteHeader(200)
	w.Write([]byte("result:" + strconv.FormatFloat(res, 'f', 6, 64)))
	return
}

func calculateString(expr string) (float64, error) {
	tr, errParse := parser.ParseExpr(expr)
	if errParse != nil {
		return 0, fmt.Errorf("error parsing expression to ast object %s", errParse)
	}
	res, errCalculate := calculateData(tr)
	if errCalculate != nil {
		return 0, errCalculate
	}
	return res, nil
}

func calculateData(exp interface{}) (float64, error) {
	switch vr := exp.(type) {
	case *ast.ParenExpr:
		return calculateData(vr.X)

	case *ast.BinaryExpr:
		x, errX := calculateData(vr.X)
		if errX != nil {
			return 0, errX
		}

		y, errY := calculateData(vr.Y)
		if errY != nil {
			return 0, errY
		}
		return makeResult(float64(x), float64(y), vr.Op.String())

	case *ast.UnaryExpr:
		val := vr.X.(*ast.BasicLit)
		res, errUnaryExpr := strconv.ParseFloat(vr.Op.String()+val.Value, 64)
		if errUnaryExpr != nil {
			return 0, errUnaryExpr
		}
		return res, nil

	case *ast.BasicLit:
		res, errBasicLit := strconv.ParseFloat(vr.Value, 64)
		if errBasicLit != nil {
			return 0, errBasicLit
		}
		return res, nil
	default:
		return 0, fmt.Errorf("no one variant was chosen")
	}
}

func makeResult(x, y float64, oper string) (float64, error) {
	switch oper {
	case "+":
		return x + y, nil
	case "-":
		return x - y, nil
	case "*":
		return x * y, nil
	case "/":
		if y == 0 {
			return 0, fmt.Errorf("divide zero error")
		}
		return x / y, nil
	default:
		return 0, fmt.Errorf("unknow type of operation '%s'", oper)
	}
}
