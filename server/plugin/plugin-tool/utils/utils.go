package utils

import (
	"go/token"
	"go/parser"
	"go/ast"
	"os"
	"path/filepath"
	"strings"

	"wvp-go/server/model/system"
)

// GetPluginData 获取插件数据（API、菜单、字典）
func GetPluginData(pluginName string) ([]system.SysApi, []system.SysBaseMenu, []system.SysDictionary) {
	var apis []system.SysApi
	var menus []system.SysBaseMenu
	var dicts []system.SysDictionary

	// 查找插件目录下的 initialize 目录中的 go 文件
	initializePath := filepath.Join("plugin", pluginName, "initialize")
	entries, err := os.ReadDir(initializePath)
	if err != nil {
		return apis, menus, dicts
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".go") {
			continue
		}
		filePath := filepath.Join(initializePath, entry.Name())
		src, err := os.ReadFile(filePath)
		if err != nil {
			continue
		}

		fileSet := token.NewFileSet()
		astFile, err := parser.ParseFile(fileSet, filePath, src, 0)
		if err != nil {
			continue
		}

		// 提取 API 数据
		findSysApiArray := findArrayElements(astFile, "SysApi")
		for _, name := range findSysApiArray {
			apis = append(apis, system.SysApi{
				Path:   "/" + pluginName + "/" + strings.ToLower(name),
				Method: "POST",
			})
		}

		// 提取菜单数据
		findSysBaseMenuArray := findArrayElements(astFile, "SysBaseMenu")
		for _, name := range findSysBaseMenuArray {
			menus = append(menus, system.SysBaseMenu{
				Name: strings.ToLower(name),
				Path: strings.ToLower(name),
			})
		}

		// 提取字典数据
		findSysDictionaryArray := findArrayElements(astFile, "SysDictionary")
		for _, name := range findSysDictionaryArray {
			dicts = append(dicts, system.SysDictionary{
				Name: name,
				Type: strings.ToLower(name),
			})
		}
	}

	return apis, menus, dicts
}

// findArrayElements 查找 AST 中的数组元素名称
func findArrayElements(file *ast.File, typeName string) []string {
	var names []string
	for _, decl := range file.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.VAR {
			continue
		}
		for _, spec := range genDecl.Specs {
			varSpec, ok := spec.(*ast.ValueSpec)
			if !ok {
				continue
			}
			for _, val := range varSpec.Values {
				compositeLit, ok := val.(*ast.CompositeLit)
				if !ok {
					continue
				}
				// 检查数组类型是否包含目标类型名
				// compositeLit.Type 是 []system.SysApi 这样的数组类型表达式
				if arrayType, ok := compositeLit.Type.(*ast.ArrayType); ok {
					if ident, ok := arrayType.Elt.(*ast.Ident); ok && ident.Name == typeName {
						for _, elt := range compositeLit.Elts {
							if cl, ok := elt.(*ast.CompositeLit); ok {
								for _, elt := range cl.Elts {
									if kv, ok := elt.(*ast.KeyValueExpr); ok {
										if ident, ok := kv.Key.(*ast.Ident); ok && ident.Name == "Name" {
											if bl, ok := kv.Value.(*ast.BasicLit); ok {
												name := strings.Trim(bl.Value, `"`)
												names = append(names, name)
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return names
}
