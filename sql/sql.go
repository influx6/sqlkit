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

// SimpleGen implements the necessary logic to generated CRUD structures for using a sql based database.
func SimpleGen(toPackage string, an ast.AnnotationDeclaration, pkgDeclr ast.PackageDeclaration, pkg ast.Package) ([]gen.WriteDirective, error) {
	sqlReadmeGen := gen.Block(
		gen.Block(
			gen.SourceText("sqlkit:simpleGen",
				string(static.MustReadFile("sql-simple-readme.tml", true)),
				struct {
					Pkg         *ast.PackageDeclaration
					Struct      ast.StructDeclaration
					PackageName string
					PackagePath string
				}{
					PackagePath: filepath.Join(pkgDeclr.Path, "sqldb"),
					PackageName: "sqldb",
					Pkg:         &pkgDeclr,
				},
			),
		),
	)

	sqlGen := gen.Block(
		gen.Package(
			gen.Name("sqldb"),
			gen.Imports(
				gen.Import("context", ""),
				gen.Import("fmt", ""),
				gen.Import("errors", ""),
				gen.Import("database/sql", "dsql"),
				gen.Import("github.com/influx6/faux/metrics", ""),
				gen.Import("github.com/jmoiron/sqlx", ""),
			),
			gen.Block(
				gen.SourceTextWith("sqlkit:simpleGen",
					string(static.MustReadFile("sql-simple.tml", true)),
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
						Pkg: &pkgDeclr,
					},
				),
			),
		),
	)

	return []gen.WriteDirective{
		{
			Writer:   sqlReadmeGen,
			FileName: "README.md",
			Dir:      "sqldb",
		},
		{
			Writer:   fmtwriter.New(sqlGen, true, true),
			FileName: "sqldb.go",
			Dir:      "sqldb",
		},
	}, nil
}

// MethodGen implements the necessary logic to generated method for CRUD operations for using a sql based database.
func MethodGen(toPackage string, an ast.AnnotationDeclaration, str ast.StructDeclaration, pkgDeclr ast.PackageDeclaration, pkg ast.Package) ([]gen.WriteDirective, error) {
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
	packageFinalFixturesPath := filepath.Join(toPackage, packageName, "fixtures")

	configName := an.Param("ENVName")
	if configName == "" {
		configName = strings.ToUpper(str.Package)
	}

	sqlTestGen := gen.Block(
		gen.Package(
			gen.Name(fmt.Sprintf("%s_test", packageName)),
			gen.Imports(
				gen.Import("os", ""),
				gen.Import("time", ""),
				gen.Import("testing", ""),
				gen.Import("context", ""),
				gen.Import("github.com/influx6/faux/tests", ""),
				gen.Import("github.com/influx6/faux/metrics", ""),
				gen.Import("github.com/influx6/faux/metrics/custom", ""),
				gen.Import("github.com/go-sql-driver/mysql", "_"),
				gen.Import("github.com/mattn/go-sqlite3", "_"),
				gen.Import(packageFinalPath, "sqldb"),
				gen.Import(packageFinalFixturesPath, "fixtures"),
			),
			gen.Block(
				gen.SourceTextWith("sqlkit:methodGen",
					string(static.MustReadFile("sql-functions-test.tml", true)),
					gen.ToTemplateFuncs(
						ast.ASTTemplatFuncs,
						template.FuncMap{
							"hasFunc": pkgDeclr.HasFunctionFor,
						},
					),
					struct {
						ENVName string
						Pkg     *ast.PackageDeclaration
						Struct  ast.StructDeclaration
					}{
						ENVName: configName,
						Pkg:     &pkgDeclr,
						Struct:  str,
					},
				),
			),
		),
	)

	sqlMakefileGen := gen.Block(
		gen.Block(
			gen.SourceText("sqlkit:methodGen",
				string(static.MustReadFile("makefile.tml", true)),
				struct {
					Pkg         *ast.PackageDeclaration
					Struct      ast.StructDeclaration
					PackageName string
					PackagePath string
					ENVName     string
				}{
					ENVName:     configName,
					PackagePath: packageFinalPath,
					PackageName: packageName,
					Pkg:         &pkgDeclr,
					Struct:      str,
				},
			),
		),
	)

	sqlDockerfileGen := gen.Block(
		gen.Block(
			gen.SourceText("sqlkit:methodGen",
				string(static.MustReadFile("dockerfile.tml", true)),
				struct {
					Pkg         *ast.PackageDeclaration
					Struct      ast.StructDeclaration
					PackageName string
					PackagePath string
					ENVName     string
				}{
					ENVName:     configName,
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
				gen.SourceTextWith("sqlkit:methodGen",
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

	sqlGen := gen.Block(
		gen.Package(
			gen.Name(packageName),
			gen.Imports(
				gen.Import("context", ""),
				gen.Import("fmt", ""),
				gen.Import("errors", ""),
				gen.Import("time", ""),
				gen.Import("strconv", ""),
				gen.Import("strings", ""),
				gen.Import("encoding/json", ""),
				gen.Import("database/sql", "dsql"),
				gen.Import("github.com/jmoiron/sqlx", ""),
				gen.Import("github.com/influx6/faux/db", ""),
				gen.Import("github.com/influx6/faux/metrics", ""),
				gen.Import("github.com/influx6/faux/db/sql/tables", ""),
				gen.Import(str.Path, ""),
			),
			gen.Block(
				gen.SourceTextWith("sqlkit:methodGen",
					string(static.MustReadFile("sql-functions.tml", true)),
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
			Writer:   sqlMakefileGen,
			FileName: "makefile",
			Dir:      packageName,
		},
		{
			Writer:   sqlDockerfileGen,
			FileName: "test.dockerfile",
			Dir:      packageName,
		},
		{
			Writer:   fmtwriter.New(sqlGen, true, true),
			FileName: fmt.Sprintf("%s.go", packageName),
			Dir:      packageName,
		},
		{
			Writer:   fmtwriter.New(sqlTestGen, true, true),
			FileName: fmt.Sprintf("%s_test.go", packageName),
			Dir:      packageName,
		},
	}, nil
}
