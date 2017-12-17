package sql

import (
	"fmt"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/gokit/sqlkit/static"
	"github.com/influx6/faux/fmtwriter"
	"github.com/influx6/moz/ast"
	"github.com/influx6/moz/gen"
)

// APIGen implements the necessary logic to generated CRUD structures for using a sql based database.
func APIGen(toDir string, an ast.AnnotationDeclaration, str ast.StructDeclaration, pkgDeclr ast.PackageDeclaration, pkg ast.Package) ([]gen.WriteDirective, error) {
	updateAction := str
	createAction := str

	if newActionName := an.Param("New"); newActionName != "" {
		if action, err := ast.FindStructType(pkgDeclr, newActionName); err == nil && action.Declr != nil {
			createAction = action
		}
	}

	if updateActionName := an.Param("Update"); updateActionName != "" {
		if action, err := ast.FindStructType(pkgDeclr, updateActionName); err == nil && action.Declr != nil {
			updateAction = action
		}
	}

	packageName := fmt.Sprintf("%ssql", strings.ToLower(str.Object.Name.Name))
	sqlTestGen := gen.Block(
		gen.Package(
			gen.Name(fmt.Sprintf("%s_test", packageName)),
			gen.Imports(
				gen.Import("os", ""),
				gen.Import("time", ""),
				gen.Import("testing", ""),
				gen.Import("encoding/json", ""),
				gen.Import("github.com/influx6/faux/db", ""),
				gen.Import("github.com/influx6/faux/tests", ""),
				gen.Import("github.com/influx6/faux/db/sql", ""),
				gen.Import("github.com/influx6/faux/metrics", ""),
				gen.Import("github.com/influx6/faux/context", ""),
				gen.Import("github.com/influx6/faux/metrics/custom", ""),
				gen.Import("github.com/go-sql-driver/mysql", "_"),
				gen.Import("github.com/lib/pq", "_"),
				gen.Import("github.com/mattn/go-sqlite3", "_"),
				gen.Import(filepath.Join(str.Path, toDir, packageName), "sqldb"),
				gen.Import(str.Path, ""),
			),
			gen.Block(
				gen.SourceTextWith(
					string(static.MustReadFile("sql-api-test.tml", true)),
					gen.ToTemplateFuncs(
						ast.ASTTemplatFuncs,
						template.FuncMap{
							"hasFunc": ast.HasFunctionFor(pkgDeclr),
						},
					),
					struct {
						Pkg          *ast.PackageDeclaration
						Struct       ast.StructDeclaration
						CreateAction ast.StructDeclaration
						UpdateAction ast.StructDeclaration
					}{
						Pkg:          &pkgDeclr,
						Struct:       str,
						CreateAction: createAction,
						UpdateAction: updateAction,
					},
				),
			),
		),
	)

	sqlReadmeGen := gen.Block(
		gen.Block(
			gen.SourceText(
				string(static.MustReadFile("sql-api-readme.tml", true)),
				struct {
					Pkg          *ast.PackageDeclaration
					Struct       ast.StructDeclaration
					CreateAction ast.StructDeclaration
					UpdateAction ast.StructDeclaration
				}{
					Pkg:          &pkgDeclr,
					Struct:       str,
					CreateAction: createAction,
					UpdateAction: updateAction,
				},
			),
		),
	)

	sqlJSONGen := gen.Block(
		gen.Package(
			gen.Name(fmt.Sprintf("%s_test", packageName)),
			gen.Imports(
				gen.Import("encoding/json", ""),
				gen.Import(str.Path, ""),
			),
			gen.Block(
				gen.SourceTextWith(
					string(static.MustReadFile("sql-api-json.tml", true)),
					template.FuncMap{
						"map":       ast.MapOutFields,
						"mapValues": ast.MapOutValues,
						"mapJSON":   ast.MapOutFieldsToJSON,
						"hasFunc":   ast.HasFunctionFor(pkgDeclr),
					},
					struct {
						Pkg          *ast.PackageDeclaration
						Struct       ast.StructDeclaration
						CreateAction ast.StructDeclaration
						UpdateAction ast.StructDeclaration
					}{
						Pkg:          &pkgDeclr,
						Struct:       str,
						CreateAction: createAction,
						UpdateAction: updateAction,
					},
				),
			),
		),
	)

	sqlGen := gen.Block(
		gen.Commentary(
			gen.SourceText(`Package `+packageName+` provides a auto-generated package which contains a sql CRUD API for the specific {{.Object.Name}} struct in package {{.Package}}.`, str),
		),
		gen.Package(
			gen.Name(packageName),
			gen.Imports(
				gen.Import("fmt", ""),
				gen.Import("errors", ""),
				gen.Import("github.com/jmoiron/sqlx", ""),
				gen.Import("github.com/influx6/faux/db", ""),
				gen.Import("github.com/influx6/faux/context", ""),
				gen.Import("github.com/influx6/faux/metrics", ""),
				gen.Import("github.com/influx6/faux/db/sql", ""),
				gen.Import("github.com/influx6/faux/db/sql/tables", ""),
				gen.Import(str.Path, ""),
			),
			gen.Block(
				gen.SourceTextWith(
					string(static.MustReadFile("sql-api.tml", true)),
					gen.ToTemplateFuncs(
						ast.ASTTemplatFuncs,
						template.FuncMap{
							"hasFunc": ast.HasFunctionFor(pkgDeclr),
						},
					),
					struct {
						Pkg          *ast.PackageDeclaration
						Struct       ast.StructDeclaration
						CreateAction ast.StructDeclaration
						UpdateAction ast.StructDeclaration
					}{
						Pkg:          &pkgDeclr,
						Struct:       str,
						CreateAction: createAction,
						UpdateAction: updateAction,
					},
				),
			),
		),
	)

	return []gen.WriteDirective{
		{
			Writer:       sqlJSONGen,
			FileName:     fmt.Sprintf("%s_fixtures_test.go", packageName),
			Dir:          packageName,
			DontOverride: true,
		},
		{
			Writer:   sqlReadmeGen,
			FileName: "README.md",
			Dir:      packageName,
		},
		{
			Writer:   fmtwriter.New(sqlTestGen, true, true),
			FileName: fmt.Sprintf("%s_test.go", packageName),
			Dir:      packageName,
		},
		{
			Writer:   fmtwriter.New(sqlGen, true, true),
			FileName: fmt.Sprintf("%s.go", packageName),
			Dir:      packageName,
		},
	}, nil
}
