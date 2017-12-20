package sql

import (
	"fmt"
	goast "go/ast"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/gokit/sqlkit/static"
	"github.com/influx6/faux/fmtwriter"
	"github.com/influx6/moz/ast"
	"github.com/influx6/moz/gen"
)

// APIGen implements the necessary logic to generated CRUD structures for using a sql based database.
func APIGen(toPackage string, an ast.AnnotationDeclaration, str ast.StructDeclaration, pkgDeclr ast.PackageDeclaration, pkg ast.Package) ([]gen.WriteDirective, error) {
	var hasPublicID bool

	// Validate we have a `PublicID` field.
	{
	fieldLoop:
		for _, field := range str.Struct.Fields.List {
			typeIdent, ok := field.Type.(*goast.Ident)
			if !ok {
				continue
			}

			// If typeName is not a string, skip.
			if typeIdent.Name != "string" {
				continue
			}

			for _, indent := range field.Names {
				if indent.Name == "PublicID" {
					hasPublicID = true
					break fieldLoop
				}
			}
		}
	}

	if !hasPublicID {
		return nil, fmt.Errorf(`Struct has no 'PublicID' field with 'string' type
		 Add 'PublicID string json:"public_id"' to struct %q
		`, str.Object.Name.Name)
	}

	packageName := fmt.Sprintf("%ssql", strings.ToLower(str.Object.Name.Name))
	packageFinalPath := filepath.Join(toPackage, packageName)

	//sqlTestGen := gen.Block(
	//	gen.Package(
	//		gen.Name(fmt.Sprintf("%s_test", packageName)),
	//		gen.Imports(
	//			gen.Import("os", ""),
	//			gen.Import("time", ""),
	//			gen.Import("testing", ""),
	//			gen.Import("encoding/json", ""),
	//			gen.Import("github.com/influx6/faux/db", ""),
	//			gen.Import("github.com/influx6/faux/tests", ""),
	//			gen.Import("github.com/influx6/faux/db/sql", ""),
	//			gen.Import("github.com/influx6/faux/metrics", ""),
	//			gen.Import("github.com/influx6/faux/context", ""),
	//			gen.Import("github.com/influx6/faux/metrics/custom", ""),
	//			gen.Import("github.com/go-sql-driver/mysql", "_"),
	//			gen.Import("github.com/lib/pq", "_"),
	//			gen.Import("github.com/mattn/go-sqlite3", "_"),
	//			gen.Import(filepath.Join(str.Path, toDir, packageName), "sqldb"),
	//			gen.Import(str.Path, ""),
	//		),
	//		gen.Block(
	//			gen.SourceTextWith(
	//				string(static.MustReadFile("sql-api-test.tml", true)),
	//				gen.ToTemplateFuncs(
	//					ast.ASTTemplatFuncs,
	//					template.FuncMap{
	//						"hasFunc": pkgDeclr.HasFunctionFor,
	//					},
	//				),
	//				struct {
	//					Pkg          *ast.PackageDeclaration
	//					Struct       ast.StructDeclaration
	//				}{
	//					Pkg:          &pkgDeclr,
	//					Struct:       str,
	//				},
	//			),
	//		),
	//	),
	//)

	sqlReadmeGen := gen.Block(
		gen.Block(
			gen.SourceText(
				string(static.MustReadFile("sql-api-readme.tml", true)),
				struct {
					Pkg         *ast.PackageDeclaration
					Struct      ast.StructDeclaration
					PackageName string
					PackagePath string
				}{
					PackagePath: packageFinalPath,
					PackageName: packageName,
					Pkg:         &pkgDeclr,
					Struct:      str,
				},
			),
		),
	)

	sqlJSONGen := gen.Block(
		gen.Package(
			gen.Name("fixtures"),
			gen.Imports(
				gen.Import("encoding/json", ""),
				gen.Import(str.Path, ""),
			),
			gen.Block(
				gen.SourceTextWith(
					string(static.MustReadFile("sql-api-json.tml", true)),
					template.FuncMap{
						"map":           ast.MapOutFields,
						"mapValues":     ast.MapOutValues,
						"mapJSON":       ast.MapOutFieldsToJSON,
						"mapRandomJSON": ast.MapOutFieldsWithRandomValuesToJSON,
						"hasFunc":       pkgDeclr.HasFunctionFor,
					},
					struct {
						Pkg    *ast.PackageDeclaration
						Struct ast.StructDeclaration
					}{
						Pkg:    &pkgDeclr,
						Struct: str,
					},
				),
			),
		),
	)

	sqlBackendGen := gen.Block(
		gen.Package(
			gen.Name("types"),
			gen.Imports(
				gen.Import("github.com/influx6/faux/context", ""),
				gen.Import(str.Path, ""),
			),
			gen.Block(
				gen.SourceTextWith(
					string(static.MustReadFile("sql-api-backend.tml", true)),
					gen.ToTemplateFuncs(
						ast.ASTTemplatFuncs,
						template.FuncMap{
							"hasFunc": pkgDeclr.HasFunctionFor,
						},
					),
					struct {
						Pkg    *ast.PackageDeclaration
						Struct ast.StructDeclaration
					}{
						Pkg:    &pkgDeclr,
						Struct: str,
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
				gen.Import("database/sql", "dsql"),
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
							"hasFunc": pkgDeclr.HasFunctionFor,
						},
					),
					struct {
						Pkg    *ast.PackageDeclaration
						Struct ast.StructDeclaration
					}{
						Pkg:    &pkgDeclr,
						Struct: str,
					},
				),
			),
		),
	)

	return []gen.WriteDirective{
		{
			Writer:       sqlJSONGen,
			FileName:     fmt.Sprintf("%s_fixtures.go", packageName),
			Dir:          filepath.Join(packageName, "fixtures"),
			DontOverride: true,
		},
		{
			Writer:   sqlReadmeGen,
			FileName: "README.md",
			Dir:      packageName,
		},
		{
			Writer:       fmtwriter.New(sqlBackendGen, true, true),
			FileName:     fmt.Sprintf("%s_backend.go", strings.ToLower(str.Object.Name.Name)),
			Dir:          "types",
			DontOverride: true,
		},
		{
			Writer:   fmtwriter.New(sqlGen, true, true),
			FileName: fmt.Sprintf("%s.go", packageName),
			Dir:      packageName,
		},
		//{
		//	Writer:   fmtwriter.New(sqlTestGen, true, true),
		//	FileName: fmt.Sprintf("%s_test.go", packageName),
		//	Dir:      packageName,
		//},
	}, nil
}
