package translator

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"modernc.org/cc/v4"
)

func NewTestTranslator(t *testing.T, code string, useSimplePredefs bool) *Translator {
	config, _ := cc.NewConfig(runtime.GOOS, runtime.GOARCH)
	var sources []cc.Source
	if useSimplePredefs {
		// easier to debug and sufficient for most tests
		sources = append(sources, cc.Source{Name: "<builtin>", Value: "typedef unsigned size_t; int __predefined_declarator;"})
	} else {
		sources = append(sources, cc.Source{Name: "<predefined>", Value: config.Predefined})
		sources = append(sources, cc.Source{Name: "<builtin>", Value: cc.Builtin})
	}
	sources = append(sources, cc.Source{Value: code})
	ast, err := cc.Translate(config, sources)
	assert.Equal(t, nil, err)
	translator, err := New(
		&Config{
			Rules: Rules(map[RuleTarget][]RuleSpec{"const": {{From: "^TEST", Action: "accept"}}}),
		},
	)
	assert.Equal(t, nil, err)
	translator.Learn(ast)
	return translator
}

func TestFunctionVoidParam(t *testing.T) {
	translator := NewTestTranslator(t, "void testfunc(void);", true)
	declarationFound := false
	for _, d := range translator.declares {
		if d.Name == "testfunc" {
			assert.False(t, declarationFound)
			declarationFound = true
			assert.Equal(t, false, d.IsStatic)
			assert.Equal(t, false, d.IsTypedef)
			assert.Equal(t, false, d.IsDefine)
			assert.Equal(t, "testfunc", d.Spec.GetTag())
			assert.Equal(t, "", d.Spec.GetBase())
			assert.Equal(t, FunctionKind, d.Spec.Kind())
			assert.Equal(t, false, d.Spec.IsConst())
			assert.Equal(t, true, d.Spec.IsComplete())
			assert.Equal(t, true, d.Spec.IsOpaque())
			funcSpec, ok := d.Spec.(*CFunctionSpec)
			if assert.True(t, ok) {
				assert.Equal(t, 0, len(funcSpec.Params))
				assert.Equal(t, "", funcSpec.Typedef)
				assert.Nil(t, funcSpec.Return)
			}
		}
	}
	assert.Equal(t, true, declarationFound)
}

func TestFunctionNoParam(t *testing.T) {
	translator := NewTestTranslator(t, "void testfunc();", true)
	declarationFound := false
	for _, d := range translator.declares {
		if d.Name == "testfunc" {
			assert.False(t, declarationFound)
			declarationFound = true
			assert.Equal(t, false, d.IsStatic)
			assert.Equal(t, false, d.IsTypedef)
			assert.Equal(t, false, d.IsDefine)
			assert.Equal(t, "testfunc", d.Spec.GetTag())
			assert.Equal(t, "", d.Spec.GetBase())
			assert.Equal(t, FunctionKind, d.Spec.Kind())
			assert.Equal(t, false, d.Spec.IsConst())
			assert.Equal(t, true, d.Spec.IsComplete())
			assert.Equal(t, true, d.Spec.IsOpaque())
			funcSpec, ok := d.Spec.(*CFunctionSpec)
			if assert.True(t, ok) {
				assert.Equal(t, 0, len(funcSpec.Params))
				assert.Equal(t, "", funcSpec.Typedef)
				assert.Nil(t, funcSpec.Return)
			}
		}
	}
	assert.Equal(t, true, declarationFound)
}

func TestFunctionVoidPtrParam(t *testing.T) {
	translator := NewTestTranslator(t, "void testfunc(void *);", true)
	declarationFound := false
	for _, d := range translator.declares {
		if d.Name == "testfunc" {
			assert.False(t, declarationFound)
			declarationFound = true
			assert.Equal(t, false, d.IsStatic)
			assert.Equal(t, false, d.IsTypedef)
			assert.Equal(t, false, d.IsDefine)
			assert.Equal(t, "testfunc", d.Spec.GetTag())
			assert.Equal(t, "", d.Spec.GetBase())
			assert.Equal(t, FunctionKind, d.Spec.Kind())
			assert.Equal(t, false, d.Spec.IsConst())
			assert.Equal(t, true, d.Spec.IsComplete())
			assert.Equal(t, false, d.Spec.IsOpaque())
			funcSpec, ok := d.Spec.(*CFunctionSpec)
			if assert.True(t, ok) {
				assert.Equal(t, 1, len(funcSpec.Params))
				assert.Equal(t, "", funcSpec.Typedef)
				assert.Nil(t, funcSpec.Return)
				param := funcSpec.Params[0]
				assert.Equal(t, "arg0", param.Name)
				assert.Equal(t, false, param.IsStatic)
				assert.Equal(t, false, param.IsTypedef)
				assert.Equal(t, false, param.IsDefine)
				assert.Equal(t, TypeKind, param.Spec.Kind())
				paramSpec, ok := param.Spec.(*CTypeSpec)
				if assert.True(t, ok) {
					assert.Equal(t, false, paramSpec.IsConst())
					assert.Equal(t, true, paramSpec.IsComplete())
					assert.Equal(t, false, paramSpec.IsOpaque())
					assert.Equal(t, "void*", paramSpec.GetBase())
					// void pointer is special
					assert.Equal(t, uint8(0), paramSpec.Pointers)
					assert.Equal(t, "", paramSpec.Raw)
				}
			}
		}
	}
	assert.Equal(t, true, declarationFound)
}

func TestFunctionVoidPtrTypedefParam(t *testing.T) {
	translator := NewTestTranslator(t, "typedef void* testID; void testfunc(testID id);", true)
	declarationFound := false
	for _, d := range translator.declares {
		if d.Name == "testfunc" {
			assert.False(t, declarationFound)
			declarationFound = true
			assert.Equal(t, false, d.IsStatic)
			assert.Equal(t, false, d.IsTypedef)
			assert.Equal(t, false, d.IsDefine)
			assert.Equal(t, "testfunc", d.Spec.GetTag())
			assert.Equal(t, "", d.Spec.GetBase())
			assert.Equal(t, FunctionKind, d.Spec.Kind())
			assert.Equal(t, false, d.Spec.IsConst())
			assert.Equal(t, true, d.Spec.IsComplete())
			assert.Equal(t, false, d.Spec.IsOpaque())
			funcSpec, ok := d.Spec.(*CFunctionSpec)
			if assert.True(t, ok) {
				assert.Equal(t, 1, len(funcSpec.Params))
				assert.Equal(t, "", funcSpec.Typedef)
				assert.Nil(t, funcSpec.Return)
				param := funcSpec.Params[0]
				assert.Equal(t, "id", param.Name)
				assert.Equal(t, false, param.IsStatic)
				assert.Equal(t, false, param.IsTypedef)
				assert.Equal(t, false, param.IsDefine)
				assert.Equal(t, TypeKind, param.Spec.Kind())
				paramSpec, ok := param.Spec.(*CTypeSpec)
				if assert.True(t, ok) {
					assert.Equal(t, false, paramSpec.IsConst())
					assert.Equal(t, true, paramSpec.IsComplete())
					assert.Equal(t, false, paramSpec.IsOpaque())
					assert.Equal(t, "void*", paramSpec.GetBase())
					// void pointer is special
					assert.Equal(t, uint8(0), paramSpec.Pointers)
					assert.Equal(t, "testID", paramSpec.Raw)
				}
			}
		}
	}
	assert.Equal(t, true, declarationFound)
}

func TestFunctionIntParam(t *testing.T) {
	translator := NewTestTranslator(t, "void testfunc(int testParam);", true)
	declarationFound := false
	for _, d := range translator.declares {
		if d.Name == "testfunc" {
			assert.False(t, declarationFound)
			declarationFound = true
			assert.Equal(t, false, d.IsStatic)
			assert.Equal(t, false, d.IsTypedef)
			assert.Equal(t, false, d.IsDefine)
			assert.Equal(t, "testfunc", d.Spec.GetTag())
			assert.Equal(t, "", d.Spec.GetBase())
			assert.Equal(t, FunctionKind, d.Spec.Kind())
			assert.Equal(t, false, d.Spec.IsConst())
			assert.Equal(t, true, d.Spec.IsComplete())
			assert.Equal(t, false, d.Spec.IsOpaque())
			funcSpec, ok := d.Spec.(*CFunctionSpec)
			if assert.True(t, ok) {
				assert.Equal(t, "", funcSpec.Typedef)
				assert.Nil(t, funcSpec.Return)
				if assert.Equal(t, 1, len(funcSpec.Params)) {
					param := funcSpec.Params[0]
					assert.Equal(t, "testParam", param.Name)
					assert.Equal(t, false, param.IsStatic)
					assert.Equal(t, false, param.IsTypedef)
					assert.Equal(t, false, param.IsDefine)
					assert.Equal(t, TypeKind, param.Spec.Kind())
					paramSpec, ok := param.Spec.(*CTypeSpec)
					if assert.True(t, ok) {
						assert.Equal(t, false, paramSpec.IsConst())
						assert.Equal(t, true, paramSpec.IsComplete())
						assert.Equal(t, false, paramSpec.IsOpaque())
						assert.Equal(t, "int", paramSpec.GetBase())
						assert.Equal(t, uint8(0), paramSpec.Pointers)
						assert.Equal(t, "", paramSpec.Raw)
					}
				}
			}
		}
	}
	assert.Equal(t, true, declarationFound)
}

func TestFunctionStringReturn(t *testing.T) {
	translator := NewTestTranslator(t, "const char* testfunc();", true)
	declarationFound := false
	for _, d := range translator.declares {
		if d.Name == "testfunc" {
			assert.False(t, declarationFound)
			declarationFound = true
			assert.Equal(t, false, d.IsStatic)
			assert.Equal(t, false, d.IsTypedef)
			assert.Equal(t, false, d.IsDefine)
			assert.Equal(t, "testfunc", d.Spec.GetTag())
			assert.Equal(t, "", d.Spec.GetBase())
			assert.Equal(t, FunctionKind, d.Spec.Kind())
			assert.Equal(t, false, d.Spec.IsConst())
			assert.Equal(t, true, d.Spec.IsComplete())
			assert.Equal(t, true, d.Spec.IsOpaque())
			funcSpec, ok := d.Spec.(*CFunctionSpec)
			if assert.True(t, ok) {
				assert.Equal(t, 0, len(funcSpec.Params))
				assert.Equal(t, "", funcSpec.Typedef)
				if assert.NotNil(t, funcSpec.Return) {
					returnSpec, ok := funcSpec.Return.(*CTypeSpec)
					if assert.True(t, ok) {
						assert.Equal(t, "char", returnSpec.Base)
						assert.Equal(t, "", returnSpec.Raw)
						assert.Equal(t, true, returnSpec.IsConst())
						assert.Equal(t, uint8(1), returnSpec.Pointers)
					}
				}
			}
		}
	}
	assert.Equal(t, true, declarationFound)
}

func TestTypedefInt(t *testing.T) {
	translator := NewTestTranslator(t, "typedef int TestVal;", true)
	typedefFound := false
	for _, d := range translator.declares {
		assert.NotEqual(t, "TestVal", d.Name, "Typedefs should not be in the list of declarations.")
	}
	for _, d := range translator.typedefs {
		if d.Name == "TestVal" {
			assert.False(t, typedefFound)
			typedefFound = true
			assert.Equal(t, false, d.IsStatic)
			assert.Equal(t, true, d.IsTypedef)
			assert.Equal(t, false, d.IsDefine)
			assert.Equal(t, "", d.Spec.GetTag())
			assert.Equal(t, "int", d.Spec.GetBase())
			assert.Equal(t, TypeKind, d.Spec.Kind())
			assert.Equal(t, false, d.Spec.IsConst())
			assert.Equal(t, true, d.Spec.IsComplete())
			assert.Equal(t, false, d.Spec.IsOpaque())
			typeSpec, ok := d.Spec.(*CTypeSpec)
			if assert.True(t, ok) {
				assert.Equal(t, "TestVal", typeSpec.Raw)
				assert.Equal(t, false, typeSpec.Signed) // not explicitly signed
				assert.Equal(t, false, typeSpec.Unsigned)
				assert.Equal(t, false, typeSpec.Short)
				assert.Equal(t, false, typeSpec.Long)
				assert.Equal(t, false, typeSpec.Complex)
				assert.Equal(t, false, typeSpec.Opaque)
			}
		}
	}
	assert.Equal(t, true, typedefFound)
}

func TestStructBoolField(t *testing.T) {
	translator := NewTestTranslator(t, "#include <stdbool.h>\nstruct TestStruct { bool TestField; };", true)
	declarationFound := false
	for _, d := range translator.declares {
		if d.Spec.GetTag() == "TestStruct" {
			assert.False(t, declarationFound)
			declarationFound = true
			assert.Equal(t, "", d.Name) // Name is for typedef, upper struct has Tag only
			assert.Equal(t, false, d.IsStatic)
			assert.Equal(t, false, d.IsTypedef)
			assert.Equal(t, false, d.IsDefine)
			assert.Equal(t, "TestStruct", d.Spec.GetBase())
			assert.Equal(t, StructKind, d.Spec.Kind())
			assert.Equal(t, false, d.Spec.IsConst())
			assert.Equal(t, true, d.Spec.IsComplete())
			assert.Equal(t, false, d.Spec.IsOpaque())
			structSpec, ok := d.Spec.(*CStructSpec)
			if assert.True(t, ok) {
				assert.Equal(t, "TestStruct", structSpec.Tag)
				assert.Equal(t, false, structSpec.IsUnion)
				assert.Equal(t, "", structSpec.Typedef)
				if assert.Equal(t, 1, len(structSpec.Members)) {
					assert.Equal(t, "TestField", structSpec.Members[0].Name)
					assert.Equal(t, "_Bool", structSpec.Members[0].Spec.GetBase())
					typeSpec, ok := structSpec.Members[0].Spec.(*CTypeSpec)
					if assert.True(t, ok) {
						// _Bool is unsigned, but the flag is not set in this case
						assert.Equal(t, false, typeSpec.Unsigned)
					}
				}
			}
		}
	}
	assert.Equal(t, true, declarationFound)
}

func TestStructUnsignedField(t *testing.T) {
	translator := NewTestTranslator(t, "struct TestStruct { unsigned TestField1; int TestField2; };", true)
	declarationFound := false
	for _, d := range translator.declares {
		if d.Spec.GetTag() == "TestStruct" {
			assert.False(t, declarationFound)
			declarationFound = true
			assert.Equal(t, "", d.Name) // Name is for typedef, upper struct has Tag only
			assert.Equal(t, false, d.IsStatic)
			assert.Equal(t, false, d.IsTypedef)
			assert.Equal(t, false, d.IsDefine)
			assert.Equal(t, "TestStruct", d.Spec.GetBase())
			assert.Equal(t, StructKind, d.Spec.Kind())
			assert.Equal(t, false, d.Spec.IsConst())
			assert.Equal(t, true, d.Spec.IsComplete())
			assert.Equal(t, false, d.Spec.IsOpaque())
			structSpec, ok := d.Spec.(*CStructSpec)
			if assert.True(t, ok) {
				assert.Equal(t, "TestStruct", structSpec.Tag)
				assert.Equal(t, false, structSpec.IsUnion)
				assert.Equal(t, "", structSpec.Typedef)
				if assert.Equal(t, 2, len(structSpec.Members)) {
					assert.Equal(t, "TestField1", structSpec.Members[0].Name)
					assert.Equal(t, "int", structSpec.Members[0].Spec.GetBase())
					typeSpec1, ok := structSpec.Members[0].Spec.(*CTypeSpec)
					if assert.True(t, ok) {
						assert.Equal(t, true, typeSpec1.Unsigned)
					}
					assert.Equal(t, "TestField2", structSpec.Members[1].Name)
					assert.Equal(t, "int", structSpec.Members[1].Spec.GetBase())
					typeSpec2, ok := structSpec.Members[1].Spec.(*CTypeSpec)
					if assert.True(t, ok) {
						assert.Equal(t, false, typeSpec2.Unsigned)
					}
				}
			}
		}
	}
	assert.Equal(t, true, declarationFound)
}

func TestFunctionTypedefReturn(t *testing.T) {
	translator := NewTestTranslator(t, "typedef unsigned int RetInt; RetInt* testfunc();", true)
	typedefFound := false
	for _, d := range translator.typedefs {
		if d.Name == "RetInt" {
			assert.False(t, typedefFound)
			typedefFound = true
			assert.Equal(t, false, d.IsStatic)
			assert.Equal(t, true, d.IsTypedef)
			assert.Equal(t, false, d.IsDefine)
			assert.Equal(t, "", d.Spec.GetTag())
			assert.Equal(t, "int", d.Spec.GetBase())
			assert.Equal(t, TypeKind, d.Spec.Kind())
			assert.Equal(t, false, d.Spec.IsConst())
			assert.Equal(t, true, d.Spec.IsComplete())
			assert.Equal(t, false, d.Spec.IsOpaque())
			typeSpec, ok := d.Spec.(*CTypeSpec)
			if assert.True(t, ok) {
				assert.Equal(t, "RetInt", typeSpec.Raw)
				assert.Equal(t, false, typeSpec.Signed)
				assert.Equal(t, true, typeSpec.Unsigned)
				assert.Equal(t, false, typeSpec.Short)
				assert.Equal(t, false, typeSpec.Long)
				assert.Equal(t, false, typeSpec.Complex)
				assert.Equal(t, false, typeSpec.Opaque)
			}
		}
	}
	declarationFound := false
	for _, d := range translator.declares {
		if d.Name == "testfunc" {
			assert.False(t, declarationFound)
			declarationFound = true
			assert.Equal(t, false, d.IsStatic)
			assert.Equal(t, false, d.IsTypedef)
			assert.Equal(t, false, d.IsDefine)
			assert.Equal(t, "testfunc", d.Spec.GetTag())
			assert.Equal(t, "", d.Spec.GetBase())
			assert.Equal(t, FunctionKind, d.Spec.Kind())
			assert.Equal(t, false, d.Spec.IsConst())
			assert.Equal(t, true, d.Spec.IsComplete())
			assert.Equal(t, true, d.Spec.IsOpaque())
			funcSpec, ok := d.Spec.(*CFunctionSpec)
			if assert.True(t, ok) {
				assert.Equal(t, 0, len(funcSpec.Params))
				assert.Equal(t, "RetInt", funcSpec.Typedef)
				if assert.NotNil(t, funcSpec.Return) {
					returnSpec, ok := funcSpec.Return.(*CTypeSpec)
					if assert.True(t, ok) {
						assert.Equal(t, "int", returnSpec.Base)
						assert.Equal(t, "RetInt", returnSpec.Raw)
						assert.Equal(t, false, returnSpec.IsConst())
						assert.Equal(t, uint8(1), returnSpec.Pointers)
					}
				}
			}
		}
	}
	assert.Equal(t, true, typedefFound)
	assert.Equal(t, true, declarationFound)
}

func TestFunctionTypedefParam(t *testing.T) {
	translator := NewTestTranslator(t, "typedef float ParamFloat; void testfunc(ParamFloat testparam);", true)
	typedefFound := false
	for _, d := range translator.typedefs {
		if d.Name == "ParamFloat" {
			assert.False(t, typedefFound)
			typedefFound = true
			assert.Equal(t, false, d.IsStatic)
			assert.Equal(t, true, d.IsTypedef)
			assert.Equal(t, false, d.IsDefine)
			assert.Equal(t, "", d.Spec.GetTag())
			assert.Equal(t, "float", d.Spec.GetBase())
			assert.Equal(t, TypeKind, d.Spec.Kind())
			assert.Equal(t, false, d.Spec.IsConst())
			assert.Equal(t, true, d.Spec.IsComplete())
			assert.Equal(t, false, d.Spec.IsOpaque())
			typeSpec, ok := d.Spec.(*CTypeSpec)
			if assert.True(t, ok) {
				assert.Equal(t, "ParamFloat", typeSpec.Raw)
				assert.Equal(t, false, typeSpec.Signed)
				assert.Equal(t, false, typeSpec.Unsigned)
				assert.Equal(t, false, typeSpec.Short)
				assert.Equal(t, false, typeSpec.Long)
				assert.Equal(t, false, typeSpec.Complex)
				assert.Equal(t, false, typeSpec.Opaque)
			}
		}
	}
	declarationFound := false
	for _, d := range translator.declares {
		if d.Name == "testfunc" {
			assert.False(t, declarationFound)
			declarationFound = true
			assert.Equal(t, false, d.IsStatic)
			assert.Equal(t, false, d.IsTypedef)
			assert.Equal(t, false, d.IsDefine)
			assert.Equal(t, "testfunc", d.Spec.GetTag())
			assert.Equal(t, "", d.Spec.GetBase())
			assert.Equal(t, FunctionKind, d.Spec.Kind())
			assert.Equal(t, false, d.Spec.IsConst())
			assert.Equal(t, true, d.Spec.IsComplete())
			assert.Equal(t, false, d.Spec.IsOpaque())
			funcSpec, ok := d.Spec.(*CFunctionSpec)
			if assert.True(t, ok) {
				assert.Equal(t, "", funcSpec.Typedef)
				assert.Nil(t, funcSpec.Return)
				if assert.Equal(t, 1, len(funcSpec.Params)) {
					param := funcSpec.Params[0]
					assert.Equal(t, "testparam", param.Name)
					assert.Equal(t, false, param.IsStatic)
					assert.Equal(t, false, param.IsTypedef)
					assert.Equal(t, false, param.IsDefine)
					assert.Equal(t, TypeKind, param.Spec.Kind())
					paramSpec, ok := param.Spec.(*CTypeSpec)
					if assert.True(t, ok) {
						assert.Equal(t, false, paramSpec.IsConst())
						assert.Equal(t, true, paramSpec.IsComplete())
						assert.Equal(t, false, paramSpec.IsOpaque())
						assert.Equal(t, "float", paramSpec.GetBase())
						assert.Equal(t, uint8(0), paramSpec.Pointers)
						assert.Equal(t, "ParamFloat", paramSpec.Raw)
					}
				}
			}
		}
	}
	assert.Equal(t, true, typedefFound)
	assert.Equal(t, true, declarationFound)
}

func TestFunctionStdFileParam(t *testing.T) {
	translator := NewTestTranslator(t, "#include <stdio.h>\nvoid testfunc(FILE *fileparam);", false)
	declarationFound := false
	for _, d := range translator.declares {
		if d.Name == "testfunc" {
			assert.False(t, declarationFound)
			declarationFound = true
			assert.Equal(t, false, d.IsStatic)
			assert.Equal(t, false, d.IsTypedef)
			assert.Equal(t, false, d.IsDefine)
			assert.Equal(t, "testfunc", d.Spec.GetTag())
			assert.Equal(t, "", d.Spec.GetBase())
			assert.Equal(t, FunctionKind, d.Spec.Kind())
			assert.Equal(t, false, d.Spec.IsConst())
			assert.Equal(t, true, d.Spec.IsComplete())
			assert.Equal(t, false, d.Spec.IsOpaque())
			funcSpec, ok := d.Spec.(*CFunctionSpec)
			if assert.True(t, ok) {
				assert.Equal(t, "", funcSpec.Typedef)
				assert.Nil(t, funcSpec.Return)
				if assert.Equal(t, 1, len(funcSpec.Params)) {
					param := funcSpec.Params[0]
					assert.Equal(t, "fileparam", param.Name)
					assert.Equal(t, false, param.IsStatic)
					assert.Equal(t, false, param.IsTypedef)
					assert.Equal(t, false, param.IsDefine)
					assert.Equal(t, StructKind, param.Spec.Kind())
					structSpec, ok := param.Spec.(*CStructSpec)
					if assert.True(t, ok) {
						assert.Equal(t, "FILE", structSpec.Typedef)
						assert.Equal(t, false, structSpec.IsUnion)
						assert.Equal(t, uint8(1), structSpec.Pointers)
					}
				}
			}
		}
	}
	assert.Equal(t, true, declarationFound)
}

func TestFunctionTypedef(t *testing.T) {
	translator := NewTestTranslator(t, "typedef void (*TestCallback)(const int* testlist);\nvoid TestSetCallback(TestCallback callback);", true)
	typedefFound := false
	for _, d := range translator.typedefs {
		if d.Name == "TestCallback" {
			assert.False(t, typedefFound)
			typedefFound = true
			assert.Equal(t, false, d.IsStatic)
			assert.Equal(t, true, d.IsTypedef)
			assert.Equal(t, false, d.IsDefine)
			assert.Equal(t, "TestCallback", d.Spec.GetTag())
			assert.Equal(t, "", d.Spec.GetBase())
			assert.Equal(t, FunctionKind, d.Spec.Kind())
			assert.Equal(t, false, d.Spec.IsConst())
			assert.Equal(t, true, d.Spec.IsComplete())
			assert.Equal(t, false, d.Spec.IsOpaque())
			funcSpec, ok := d.Spec.(*CFunctionSpec)
			if assert.True(t, ok) {
				assert.Equal(t, "TestCallback", funcSpec.Raw)
				assert.Equal(t, "", funcSpec.Typedef)
				assert.Equal(t, uint8(1), funcSpec.Pointers)
				assert.Equal(t, nil, funcSpec.Return)
				assert.Equal(t, 1, len(funcSpec.Params))
			}
		}
	}
	declarationFound := false
	for _, d := range translator.declares {
		if d.Name == "TestSetCallback" {
			assert.False(t, declarationFound)
			declarationFound = true
			assert.Equal(t, false, d.IsStatic)
			assert.Equal(t, false, d.IsTypedef)
			assert.Equal(t, false, d.IsDefine)
			assert.Equal(t, "TestSetCallback", d.Spec.GetTag())
			assert.Equal(t, "", d.Spec.GetBase())
			assert.Equal(t, FunctionKind, d.Spec.Kind())
			assert.Equal(t, false, d.Spec.IsConst())
			assert.Equal(t, true, d.Spec.IsComplete())
			assert.Equal(t, false, d.Spec.IsOpaque())
			funcSpec, ok := d.Spec.(*CFunctionSpec)
			if assert.True(t, ok) {
				assert.Equal(t, "", funcSpec.Typedef)
				assert.Nil(t, funcSpec.Return)
				if assert.Equal(t, 1, len(funcSpec.Params)) {
					param := funcSpec.Params[0]
					assert.Equal(t, "callback", param.Name)
					assert.Equal(t, false, param.IsStatic)
					assert.Equal(t, false, param.IsTypedef)
					assert.Equal(t, false, param.IsDefine)
					assert.Equal(t, false, param.Spec.IsConst())
					assert.Equal(t, FunctionKind, param.Spec.Kind())
					functionSpec, ok := param.Spec.(*CFunctionSpec)
					if assert.True(t, ok) {
						assert.Equal(t, "TestCallback", functionSpec.Raw)
						assert.Equal(t, "TestCallback", functionSpec.Typedef)
					}
				}
			}
		}
	}
	assert.Equal(t, true, typedefFound)
	assert.Equal(t, true, declarationFound)
}

func TestFunctionStructTypedef(t *testing.T) {
	translator := NewTestTranslator(t, "typedef struct TestStruct TestStruct;\nstruct TestStruct { float x; };\nvoid TestFunc(TestStruct* testStruct);", true)
	typedefFound := false
	for _, d := range translator.typedefs {
		if d.Name == "TestStruct" {
			assert.False(t, typedefFound)
			typedefFound = true
			assert.Equal(t, false, d.IsStatic)
			assert.Equal(t, true, d.IsTypedef)
			assert.Equal(t, false, d.IsDefine)
			assert.Equal(t, "TestStruct", d.Spec.GetTag())
			assert.Equal(t, "TestStruct", d.Spec.GetBase())
			assert.Equal(t, StructKind, d.Spec.Kind())
			assert.Equal(t, false, d.Spec.IsConst())
			assert.Equal(t, true, d.Spec.IsComplete())
			assert.Equal(t, false, d.Spec.IsOpaque())
			structSpec, ok := d.Spec.(*CStructSpec)
			if assert.True(t, ok) {
				assert.Equal(t, "TestStruct", structSpec.Tag)
				assert.Equal(t, "TestStruct", structSpec.Typedef)
				assert.Equal(t, uint8(0), structSpec.Pointers)
				if assert.Equal(t, 1, len(structSpec.Members)) {
					assert.Equal(t, "x", structSpec.Members[0].Name)
					assert.Equal(t, "float", structSpec.Members[0].Spec.GetBase())
					typeSpec, ok := structSpec.Members[0].Spec.(*CTypeSpec)
					if assert.True(t, ok) {
						assert.Equal(t, false, typeSpec.Const)
					}
				}
			}
		}
	}
	declarationFound := false
	for _, d := range translator.declares {
		if d.Spec.GetTag() == "TestStruct" {
			assert.False(t, declarationFound)
			declarationFound = true
			assert.Equal(t, "", d.Name) // Name is for typedef
			assert.Equal(t, false, d.IsStatic)
			assert.Equal(t, false, d.IsTypedef)
			assert.Equal(t, false, d.IsDefine)
			assert.Equal(t, "TestStruct", d.Spec.GetBase())
			assert.Equal(t, StructKind, d.Spec.Kind())
			assert.Equal(t, false, d.Spec.IsConst())
			assert.Equal(t, true, d.Spec.IsComplete())
			assert.Equal(t, false, d.Spec.IsOpaque())
			structSpec, ok := d.Spec.(*CStructSpec)
			if assert.True(t, ok) {
				assert.Equal(t, "TestStruct", structSpec.Tag)
				assert.Equal(t, false, structSpec.IsUnion)
				assert.Equal(t, "", structSpec.Typedef)
				if assert.Equal(t, 1, len(structSpec.Members)) {
					assert.Equal(t, "x", structSpec.Members[0].Name)
					assert.Equal(t, "float", structSpec.Members[0].Spec.GetBase())
					typeSpec, ok := structSpec.Members[0].Spec.(*CTypeSpec)
					if assert.True(t, ok) {
						assert.Equal(t, false, typeSpec.Const)
					}
				}
			}
		}
	}
	functionFound := false
	for _, d := range translator.declares {
		if d.Name == "TestFunc" {
			assert.False(t, functionFound)
			functionFound = true
			assert.Equal(t, false, d.IsStatic)
			assert.Equal(t, false, d.IsTypedef)
			assert.Equal(t, false, d.IsDefine)
			assert.Equal(t, "TestFunc", d.Spec.GetTag())
			assert.Equal(t, "", d.Spec.GetBase())
			assert.Equal(t, FunctionKind, d.Spec.Kind())
			assert.Equal(t, false, d.Spec.IsConst())
			assert.Equal(t, true, d.Spec.IsComplete())
			assert.Equal(t, false, d.Spec.IsOpaque())
			funcSpec, ok := d.Spec.(*CFunctionSpec)
			if assert.True(t, ok) {
				assert.Equal(t, "", funcSpec.Typedef)
				assert.Nil(t, funcSpec.Return)
				if assert.Equal(t, 1, len(funcSpec.Params)) {
					param := funcSpec.Params[0]
					assert.Equal(t, "testStruct", param.Name)
					assert.Equal(t, false, param.IsStatic)
					assert.Equal(t, false, param.IsTypedef)
					assert.Equal(t, false, param.IsDefine)
					assert.Equal(t, false, param.Spec.IsConst())
					assert.Equal(t, StructKind, param.Spec.Kind())
					structSpec, ok := param.Spec.(*CStructSpec)
					if assert.True(t, ok) {
						assert.Equal(t, "TestStruct", structSpec.Typedef)
						assert.Equal(t, "TestStruct", structSpec.Tag)
						assert.Equal(t, uint8(1), structSpec.Pointers)
					}
				}
			}
		}
	}
	assert.Equal(t, true, typedefFound)
	assert.Equal(t, true, declarationFound)
	assert.Equal(t, true, functionFound)
}

func TestStructNestedTypedef(t *testing.T) {
	translator := NewTestTranslator(t, "typedef struct TestLongNameStruct { int TestArray[8]; } TestLongNameStruct;\ntypedef TestLongNameStruct TestStruct;", true)
	typedefFound := false
	for _, d := range translator.typedefs {
		if d.Name == "TestStruct" {
			assert.False(t, typedefFound)
			typedefFound = true
			assert.Equal(t, false, d.IsStatic)
			assert.Equal(t, true, d.IsTypedef)
			assert.Equal(t, false, d.IsDefine)
			assert.Equal(t, "TestLongNameStruct", d.Spec.GetTag())
			assert.Equal(t, "TestStruct", d.Spec.GetBase())
			assert.Equal(t, StructKind, d.Spec.Kind())
			assert.Equal(t, false, d.Spec.IsConst())
			assert.Equal(t, true, d.Spec.IsComplete())
			assert.Equal(t, false, d.Spec.IsOpaque())
			structSpec, ok := d.Spec.(*CStructSpec)
			if assert.True(t, ok) {
				assert.Equal(t, "TestLongNameStruct", structSpec.Tag)
				assert.Equal(t, "TestStruct", structSpec.Typedef) // this is different from v1 but I think it's more correct.
				assert.Equal(t, uint8(0), structSpec.Pointers)
				if assert.Equal(t, 1, len(structSpec.Members)) {
					assert.Equal(t, "TestArray", structSpec.Members[0].Name)
					assert.Equal(t, "int", structSpec.Members[0].Spec.GetBase())
					typeSpec, ok := structSpec.Members[0].Spec.(*CTypeSpec)
					if assert.True(t, ok) {
						assert.Equal(t, false, typeSpec.Const)
						assert.Equal(t, "", typeSpec.Raw)
						assert.Equal(t, "", typeSpec.InnerArr.String())
						assert.Equal(t, "[8]", typeSpec.OuterArr.String())
					}
				}
			}
		}
	}
	baseTypedefFound := false
	for _, d := range translator.typedefs {
		if d.Name == "TestLongNameStruct" {
			assert.False(t, baseTypedefFound)
			baseTypedefFound = true
			assert.Equal(t, false, d.IsStatic)
			assert.Equal(t, true, d.IsTypedef)
			assert.Equal(t, false, d.IsDefine)
			assert.Equal(t, "TestLongNameStruct", d.Spec.GetTag())
			assert.Equal(t, "TestLongNameStruct", d.Spec.GetBase())
			assert.Equal(t, StructKind, d.Spec.Kind())
			assert.Equal(t, false, d.Spec.IsConst())
			assert.Equal(t, true, d.Spec.IsComplete())
			assert.Equal(t, false, d.Spec.IsOpaque())
			structSpec, ok := d.Spec.(*CStructSpec)
			if assert.True(t, ok) {
				assert.Equal(t, "TestLongNameStruct", structSpec.Tag)
				assert.Equal(t, "TestLongNameStruct", structSpec.Typedef)
				assert.Equal(t, uint8(0), structSpec.Pointers)
				if assert.Equal(t, 1, len(structSpec.Members)) {
					assert.Equal(t, "TestArray", structSpec.Members[0].Name)
					assert.Equal(t, "int", structSpec.Members[0].Spec.GetBase())
					typeSpec, ok := structSpec.Members[0].Spec.(*CTypeSpec)
					if assert.True(t, ok) {
						assert.Equal(t, false, typeSpec.Const)
						assert.Equal(t, "", typeSpec.Raw)
						assert.Equal(t, "", typeSpec.InnerArr.String())
						assert.Equal(t, "[8]", typeSpec.OuterArr.String())
					}
				}
			}
		}
	}
	assert.Equal(t, true, typedefFound)
	assert.Equal(t, true, baseTypedefFound)
}

func TestDefineConst(t *testing.T) {
	translator := NewTestTranslator(t, "#define TEST_ABC 10", true)
	defineFound := false
	for _, d := range translator.defines {
		if d.Name == "TEST_ABC" {
			assert.False(t, defineFound)
			defineFound = true
			assert.Equal(t, false, d.IsStatic)
			assert.Equal(t, false, d.IsTypedef)
			assert.Equal(t, true, d.IsDefine)
			assert.Equal(t, "10", d.Expression)
			assert.Nil(t, d.Spec) // there is no spec for macros
		}
	}
	assert.Equal(t, true, defineFound)
}
