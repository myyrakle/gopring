package ast_util

import (
	"go/ast"
)

func GetSelectorNameFromType(typeExpr ast.Expr) *string {
	currentExpr := typeExpr

	if field, ok := currentExpr.(*ast.StarExpr); ok {
		currentExpr = field.X
	}

	if expr, ok := typeExpr.(*ast.SelectorExpr); ok {
		currentExpr = expr.X
	}

	if ident, ok := currentExpr.(*ast.Ident); ok {
		return &ident.Name
	} else {
		return nil
	}
}

func GetTypeNameFromType(typeExpr ast.Expr) *string {
	currentExpr := typeExpr

	if field, ok := currentExpr.(*ast.StarExpr); ok {
		currentExpr = field.X
	}

	if expr, ok := currentExpr.(*ast.SelectorExpr); ok {
		selectorName := expr.Sel.String()
		return &selectorName
	}

	return nil
}
