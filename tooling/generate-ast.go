package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Println("Usage: generate_ast <output directory>")
		os.Exit(64)
	}

	outputDir := args[1]

	err := defineAst(outputDir, "Expr", []string{
		"Binary    : Left Expr[R], Operator token.Token, Right Expr[R]",
		"Grouping  : Expression Expr[R]",
		"Literal   : Value interface{}",
		"Unary     : Operator token.Token, Right Expr[R]",
	})
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(-1)
	}

}

func defineAst(outDir string, baseName string, types []string) error {
	path := fmt.Sprintf("%s/%s.go", outDir, baseName)
	if _, err := os.Stat(outDir); os.IsNotExist(err) {
		err := os.Mkdir(outDir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("defineAST: failed to create file... %w", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	_, err = fmt.Fprintf(writer, "package %s\n\nimport \"github.com/brewinski/crafting-interpreters/pkg/token\"\n", strings.ToLower(baseName))
	_, err = fmt.Fprintln(writer, "")
	_, err = fmt.Fprint(writer, "\n\n")

	_, err = fmt.Fprintf(writer, "type %s[R any] interface {\n", baseName)
	_, err = fmt.Fprintln(writer, "\tAccept(visitor Visitor[R]) R\n")
	_, err = fmt.Fprintf(writer, "}")

	_, err = fmt.Fprint(writer, "\n\n")

	defineVisitor(writer, baseName, types)

	_, err = fmt.Fprint(writer, "\n\n")

	for _, t := range types {
		structName := strings.Trim(strings.Split(t, ":")[0], " ")
		fields := strings.Trim(strings.Split(t, ":")[1], " ")
		defineStructure(writer, baseName, structName, fields)
	}

	err = writer.Flush()

	if err != nil {
		return err
	}

	return nil
}

//	type Binary[R any] struct {
//		left     Expr
//		operator token.Token
//		right    Expr
//	}
//
//	func NewBinary[R any](left Expr, operator token.Token, right Expr) Binary[R] {
//		return Binary[R]{
//			left,
//			operator,
//			right,
//		}
//	}
//
//	func (binary Binary[R]) Accept(visitor Visitor[R]) R {
//		return visitor.visitBinaryExpr(binary)
//	}
func defineStructure(w *bufio.Writer, name, structName, fields string) {
	fmt.Fprintf(w, "type %s[R any] struct {\n", structName)

	fieldsArr := strings.Split(fields, ",")
	for _, f := range fieldsArr {
		fmt.Fprintf(w, "\t%s\n", strings.Trim(f, " "))
	}

	fmt.Fprintf(w, "}\n\n")

	fmt.Fprintf(w, "func New%s[R any](%s) %s[R] { \n", structName, fields, structName)
	fmt.Fprintf(w, "\treturn %s[R]{\n", structName)

	for _, f := range fieldsArr {
		f = strings.Trim(f, " ")
		fParts := strings.Split(f, " ")
		fmt.Fprintf(w, "\t\t%s,\n", fParts[0])
	}

	fmt.Fprintf(w, "\t}\n")
	fmt.Fprintf(w, "}\n")

	fmt.Fprintf(w, "\n\n")

	// define the accept method
	fmt.Fprintf(w, "func (%s %s[R]) Accept(visitor Visitor[R]) R {\n\treturn visitor.Visit%s%s(%s)\n}", strings.ToLower(structName), structName, structName, name, strings.ToLower(structName))

	fmt.Fprintf(w, "\n\n")
}

//	type Visitor[R any] interface {
//		visitBinaryExpr(binary Binary) R
//	}
func defineVisitor(w *bufio.Writer, baseName string, typeFields []string) {
	fmt.Fprintf(w, "type Visitor[R any] interface {\n")

	for _, tf := range typeFields {
		structName := strings.Trim(strings.Split(tf, ":")[0], " ")
		fmt.Fprintf(w, "\tVisit%s%s(%s %s[R]) R\n", structName, baseName, strings.ToLower(structName), structName)
	}

	fmt.Fprintf(w, "}\n")
}
