package carrier

import (
	"bytes"
	"context"
	"embed"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"text/template"
)

const (
	TypeStruct = iota
	TypeEnt
)

const (
	FieldTypeSimple = iota
	FieldTypeStruct
)

//go:embed template/*
var templateDir embed.FS

func writeTo(name string, tvar templateVar, path string) error {
	funcMap := template.FuncMap{
		"FirstLower": firstLower,
	}
	tmpl, err := template.New(name).Funcs(funcMap).ParseFS(
		templateDir, fmt.Sprintf("template/%s", name),
	)
	if err != nil {
		return err
	}
	b := &bytes.Buffer{}
	err = tmpl.Execute(b, tvar)
	if err != nil {
		return err
	}
	var buf []byte
	if buf, err = format.Source(b.Bytes()); err != nil {
		fmt.Println("formating:", err)
		fmt.Println(b.String())
		return err
	}

	if err = ioutil.WriteFile(path, buf, 0644); err != nil { //nolint
		return err
	}
	return nil
}

func writeToSingle(name string, tvar templateVarSingle, path string) error {
	funcMap := template.FuncMap{
		"FirstLower": firstLower,
	}
	tmpl, err := template.New(name).Funcs(funcMap).ParseFS(
		templateDir, fmt.Sprintf("template/%s", name),
	)
	if err != nil {
		return err
	}
	b := &bytes.Buffer{}
	err = tmpl.Execute(b, tvar)
	if err != nil {
		return err
	}
	var buf []byte
	if buf, err = format.Source(b.Bytes()); err != nil {
		fmt.Println("formating:", err)
		fmt.Println(b.String())
		return err
	}

	if err = ioutil.WriteFile(path, buf, 0644); err != nil { //nolint
		return err
	}
	return nil
}

type Schema interface {
	Name() string
	Type() string
	Import() string
	Fields() []Field
	PostFields() []PostField
	check() bool
	schemaType() int
	TraitList() []string
	EntPkg() string
}

type StructSchema struct {
	To     interface{}
	Alias  string
	Posts  []PostField
	rtype  reflect.Type
	Traits []string
}

func (s *StructSchema) schemaType() int { return TypeStruct }

func (s *StructSchema) check() bool {
	v := reflect.ValueOf(s.To)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		fmt.Println("schema not valid")
		return false
	}
	s.rtype = v.Type()
	return true
}

func (s *StructSchema) Name() string {
	if s.Alias != "" {
		return s.Alias
	}
	return s.rtype.Name()
}
func (s *StructSchema) Type() string {
	return s.rtype.String()
}
func (s *StructSchema) Import() string { return pkgPath(s.rtype) }
func (s *StructSchema) Fields() []Field {
	t := s.rtype
	var fields []Field
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.PkgPath != "" {
			continue
		}
		f := Field{
			Name:      field.Name,
			ValueType: field.Type.String(),
			Import:    pkgPath(field.Type),
		}
		fields = append(fields, f)
	}
	return fields
}
func (s *StructSchema) PostFields() []PostField { return s.Posts }
func (s *StructSchema) TraitList() []string     { return s.Traits }
func (s *StructSchema) EntPkg() string          { return "" }

type EntCreateInterface interface {
	Exec(ctx context.Context) error
}

type EntSchema struct {
	To     EntCreateInterface // ent model struct
	Posts  []PostField
	rtype  reflect.Type
	Alias  string
	Traits []string
	entPkg string
}

func (s *EntSchema) schemaType() int { return TypeEnt }

func (s *EntSchema) check() bool {
	v := reflect.ValueOf(s.To)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		fmt.Println("schema not valid")
		return false
	}

	p := reflect.PtrTo(v.Type())
	method, ok := p.MethodByName("Save")
	if !ok {
		fmt.Println("save method not found")
		return false
	}
	out := method.Type.Out(0)
	s.rtype = out.Elem()
	tp := strings.Split(v.Type().String(), ".")
	s.entPkg = tp[0]
	return true
}
func (s *EntSchema) Name() string {
	if s.Alias != "" {
		return s.Alias
	}
	return s.rtype.Name()
}
func (s *EntSchema) Type() string {
	return s.rtype.String()
}
func (s *EntSchema) Import() string { return pkgPath(s.rtype) }
func (s *EntSchema) Fields() []Field {
	v := reflect.ValueOf(s.To)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		fmt.Println("schema not valid")
		return nil
	}
	t := reflect.PtrTo(v.Type())
	var fields []Field
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		if method.PkgPath != "" {
			continue
		}
		if !strings.HasPrefix(method.Name, "Set") {
			continue
		}
		if strings.HasPrefix(method.Name, "SetNillable") {
			continue
		}
		fieldName := strings.TrimPrefix(method.Name, "Set")
		fieldInput := method.Type.In(1)
		f := Field{
			Name:      fieldName,
			ValueType: fieldInput.String(),
			Setter:    method.Name,
			Import:    pkgPath(fieldInput),
		}
		fields = append(fields, f)
	}
	return fields
}
func (s *EntSchema) PostFields() []PostField { return s.Posts }
func (s *EntSchema) TraitList() []string     { return s.Traits }
func (s *EntSchema) EntPkg() string          { return s.entPkg }

type PostField struct {
	Name  string
	Input interface{}
}

type Field struct {
	Name      string
	ValueType string
	Setter    string
	Import    string
}

type carrierInfo struct {
	Name       string
	Type       string
	Fields     []Field
	PostFields []Field
	imports    map[string]string
	Traits     []string
	EntPkg     string
}

type templateVar struct {
	Schemas    []carrierInfo
	EntSchemas []carrierInfo
	Imports    []string
}

type templateVarSingle struct {
	Schema  carrierInfo
	Imports []string
	Prefix  string
}

func pkgPath(t reflect.Type) string {
	pkg := t.PkgPath()
	if pkg != "" {
		return pkg
	}
	switch t.Kind() {
	case reflect.Slice, reflect.Array, reflect.Ptr:
		return pkgPath(t.Elem())

	case reflect.Map:
		return pkgPath(t.Key()) + "|" + pkgPath(t.Elem())
	}
	return pkg
}

func firstLower(s string) string {
	if len(s) == 0 {
		return s
	}

	return strings.ToLower(s[:1]) + s[1:]
}

func updateImport(m map[string]string, path string) {
	all := strings.Split(path, "|")
	for _, p := range all {
		if p != "" {
			m[p] = ""
		}
	}
}

func SchemaToMetaFactory(pkg string, path string, schemas []Schema) error {
	const duplicateError = "duplicate name found: %s, please add alias to schema"
	var structInfo []carrierInfo
	var entInfo []carrierInfo
	var factoryImports = map[string]bool{pkg: true}
	structNames := make(map[string]bool)
	entNames := make(map[string]bool)

	for _, s := range schemas {
		valid := s.check()
		if !valid {
			continue
		}
		// check name duplicate
		if s.schemaType() == TypeStruct {
			_, ok := structNames[s.Name()]
			if ok {
				return fmt.Errorf(duplicateError, s.Name())
			}
			structNames[s.Name()] = true
		} else {
			_, ok := entNames[s.Name()]
			if ok {
				return fmt.Errorf(duplicateError, s.Name())
			}
			entNames[s.Name()] = true
			factoryImports[s.Import()] = true
		}
		importMap := make(map[string]string)
		updateImport(importMap, s.Import())

		fields := s.Fields()
		for _, f := range fields {
			updateImport(importMap, f.Import)
		}

		var postFields []Field
		for _, post := range s.PostFields() {
			name := strings.Title(post.Name)
			t := reflect.TypeOf(post.Input)
			updateImport(importMap, pkgPath(t))
			postFields = append(
				postFields,
				Field{Name: name, ValueType: t.String()},
			)
		}

		var info *[]carrierInfo
		if s.schemaType() == TypeStruct {
			info = &structInfo
		} else {
			info = &entInfo
		}

		var traits []string
		for _, t := range s.TraitList() {
			traits = append(traits, strings.Title(t))
		}

		*info = append(*info, carrierInfo{
			Name:       s.Name(),
			Type:       s.Type(),
			Fields:     fields,
			PostFields: postFields,
			imports:    importMap,
			Traits:     traits,
			EntPkg:     s.EntPkg(),
		})
	}
	// prepare dir
	abs, err := filepath.Abs(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	dir := filepath.Dir(abs)
	err = os.RemoveAll(dir + "/factory")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = os.Mkdir(dir+"/factory", os.ModePerm)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// generate factory.go
	fi := []string{}
	for k := range factoryImports {
		fi = append(fi, k)
	}
	err = writeTo("factory.tmpl", templateVar{
		Schemas:    structInfo,
		EntSchemas: entInfo,
		Imports:    fi,
	}, dir+"/factory.go")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// generate base.go
	err = writeTo("base.tmpl", templateVar{}, dir+"/factory/base.go")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// generate struct factories
	for _, schema := range structInfo {
		imports := []string{}
		for k := range schema.imports {
			imports = append(imports, k)
		}
		fullPath := fmt.Sprintf("%s/factory/%s.go", dir, strings.ToLower(schema.Name))
		err = writeToSingle("carrier.tmpl", templateVarSingle{
			Schema:  schema,
			Imports: imports,
			Prefix:  "",
		}, fullPath)
		if err != nil {
			return err
		}
	}
	// generate ent factories
	for _, schema := range entInfo {
		imports := []string{}
		for k := range schema.imports {
			imports = append(imports, k)
		}
		fullPath := fmt.Sprintf("%s/factory/ent_%s.go", dir, strings.ToLower(schema.Name))
		err = writeToSingle("carrier.tmpl", templateVarSingle{
			Schema:  schema,
			Imports: imports,
			Prefix:  "Ent",
		}, fullPath)
		if err != nil {
			return err
		}
	}
	return nil
}
