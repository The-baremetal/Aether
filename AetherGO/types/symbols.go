package types

import (
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

type SymbolScope struct {
	Parent  *SymbolScope
	Symbols map[string]*Symbol
}

type Symbol struct {
	Name     string
	Type     types.Type
	Value    value.Value
	Constant bool
}

func NewSymbolScope(parent *SymbolScope) *SymbolScope {
	return &SymbolScope{
		Parent:  parent,
		Symbols: make(map[string]*Symbol),
	}
}

func (s *SymbolScope) Lookup(name string) *Symbol {
	if sym, ok := s.Symbols[name]; ok {
		return sym
	}
	if s.Parent != nil {
		return s.Parent.Lookup(name)
	}
	return nil
}

func (s *SymbolScope) Define(sym *Symbol) {
	s.Symbols[sym.Name] = sym
}
