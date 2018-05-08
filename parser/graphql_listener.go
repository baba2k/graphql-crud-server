package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/baba2k/graphql-rungen/parser/antlr"
	"fmt"
)

type GraphQLListener struct {
	*GraphQLListener
}

func NewGraphqlListener() *GraphQLListener {
	return new(GraphQLListener)
}

func (this *GraphQLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	//fmt.Println(ctx.GetText())
}

func (this *GraphQLListener) VisitErrorNode(node antlr.ErrorNode) {
	panic("could no parse graphql schema")
}

// VisitTerminal is called when a terminal node is visited.
func (s *GraphQLListener) VisitTerminal(node antlr.TerminalNode) {}

// ExitEveryRule is called when any rule is exited.
func (s *GraphQLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterDocument is called when production document is entered.
func (s *GraphQLListener) EnterDocument(ctx *parser.DocumentContext) { fmt.Println("EnterDocument") }

// ExitDocument is called when production document is exited.
func (s *GraphQLListener) ExitDocument(ctx *parser.DocumentContext) { fmt.Println("ExitDocument") }

// EnterDefinition is called when production definition is entered.
func (s *GraphQLListener) EnterDefinition(ctx *parser.DefinitionContext) {}

// ExitDefinition is called when production definition is exited.
func (s *GraphQLListener) ExitDefinition(ctx *parser.DefinitionContext) {}

// EnterOperationType is called when production operationType is entered.
func (s *GraphQLListener) EnterOperationType(ctx *parser.OperationTypeContext) { fmt.Println("EnterOperationType", ctx.GetText()) }

// ExitOperationType is called when production operationType is exited.
func (s *GraphQLListener) ExitOperationType(ctx *parser.OperationTypeContext) { fmt.Println("ExitOperationType", ctx.GetText()) }

// EnterDescription is called when production description is entered.
func (s *GraphQLListener) EnterDescription(ctx *parser.DescriptionContext) {}

// ExitDescription is called when production description is exited.
func (s *GraphQLListener) ExitDescription(ctx *parser.DescriptionContext) {}

// EnterEnumValue is called when production enumValue is entered.
func (s *GraphQLListener) EnterEnumValue(ctx *parser.EnumValueContext) {}

// ExitEnumValue is called when production enumValue is exited.
func (s *GraphQLListener) ExitEnumValue(ctx *parser.EnumValueContext) {}

// EnterArrayValue is called when production arrayValue is entered.
func (s *GraphQLListener) EnterArrayValue(ctx *parser.ArrayValueContext) {}

// ExitArrayValue is called when production arrayValue is exited.
func (s *GraphQLListener) ExitArrayValue(ctx *parser.ArrayValueContext) {}

// EnterArrayValueWithVariable is called when production arrayValueWithVariable is entered.
func (s *GraphQLListener) EnterArrayValueWithVariable(ctx *parser.ArrayValueWithVariableContext) {}

// ExitArrayValueWithVariable is called when production arrayValueWithVariable is exited.
func (s *GraphQLListener) ExitArrayValueWithVariable(ctx *parser.ArrayValueWithVariableContext) {}

// EnterObjectValue is called when production objectValue is entered.
func (s *GraphQLListener) EnterObjectValue(ctx *parser.ObjectValueContext) {}

// ExitObjectValue is called when production objectValue is exited.
func (s *GraphQLListener) ExitObjectValue(ctx *parser.ObjectValueContext) {}

// EnterObjectValueWithVariable is called when production objectValueWithVariable is entered.
func (s *GraphQLListener) EnterObjectValueWithVariable(ctx *parser.ObjectValueWithVariableContext) {}

// ExitObjectValueWithVariable is called when production objectValueWithVariable is exited.
func (s *GraphQLListener) ExitObjectValueWithVariable(ctx *parser.ObjectValueWithVariableContext) {}

// EnterObjectField is called when production objectField is entered.
func (s *GraphQLListener) EnterObjectField(ctx *parser.ObjectFieldContext) {}

// ExitObjectField is called when production objectField is exited.
func (s *GraphQLListener) ExitObjectField(ctx *parser.ObjectFieldContext) {}

// EnterObjectFieldWithVariable is called when production objectFieldWithVariable is entered.
func (s *GraphQLListener) EnterObjectFieldWithVariable(ctx *parser.ObjectFieldWithVariableContext) {}

// ExitObjectFieldWithVariable is called when production objectFieldWithVariable is exited.
func (s *GraphQLListener) ExitObjectFieldWithVariable(ctx *parser.ObjectFieldWithVariableContext) {}

// EnterDirectives is called when production directives is entered.
func (s *GraphQLListener) EnterDirectives(ctx *parser.DirectivesContext) {}

// ExitDirectives is called when production directives is exited.
func (s *GraphQLListener) ExitDirectives(ctx *parser.DirectivesContext) {}

// EnterDirective is called when production directive is entered.
func (s *GraphQLListener) EnterDirective(ctx *parser.DirectiveContext) {}

// ExitDirective is called when production directive is exited.
func (s *GraphQLListener) ExitDirective(ctx *parser.DirectiveContext) {}

// EnterArguments is called when production arguments is entered.
func (s *GraphQLListener) EnterArguments(ctx *parser.ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *GraphQLListener) ExitArguments(ctx *parser.ArgumentsContext) {}

// EnterArgument is called when production argument is entered.
func (s *GraphQLListener) EnterArgument(ctx *parser.ArgumentContext) {}

// ExitArgument is called when production argument is exited.
func (s *GraphQLListener) ExitArgument(ctx *parser.ArgumentContext) {}

// EnterName is called when production name is entered.
func (s *GraphQLListener) EnterName(ctx *parser.NameContext) {}

// ExitName is called when production name is exited.
func (s *GraphQLListener) ExitName(ctx *parser.NameContext) {}

// EnterValue is called when production value is entered.
func (s *GraphQLListener) EnterValue(ctx *parser.ValueContext) {}

// ExitValue is called when production value is exited.
func (s *GraphQLListener) ExitValue(ctx *parser.ValueContext) {}

// EnterValueWithVariable is called when production valueWithVariable is entered.
func (s *GraphQLListener) EnterValueWithVariable(ctx *parser.ValueWithVariableContext) {}

// ExitValueWithVariable is called when production valueWithVariable is exited.
func (s *GraphQLListener) ExitValueWithVariable(ctx *parser.ValueWithVariableContext) {}

// EnterVariable is called when production variable is entered.
func (s *GraphQLListener) EnterVariable(ctx *parser.VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *GraphQLListener) ExitVariable(ctx *parser.VariableContext) {}

// EnterDefaultValue is called when production defaultValue is entered.
func (s *GraphQLListener) EnterDefaultValue(ctx *parser.DefaultValueContext) {}

// ExitDefaultValue is called when production defaultValue is exited.
func (s *GraphQLListener) ExitDefaultValue(ctx *parser.DefaultValueContext) {}

// EnterStringValue is called when production stringValue is entered.
func (s *GraphQLListener) EnterStringValue(ctx *parser.StringValueContext) {}

// ExitStringValue is called when production stringValue is exited.
func (s *GraphQLListener) ExitStringValue(ctx *parser.StringValueContext) {}

// EnterObjType is called when production objType is entered.
func (s *GraphQLListener) EnterObjType(ctx *parser.ObjTypeContext) {}

// ExitObjType is called when production objType is exited.
func (s *GraphQLListener) ExitObjType(ctx *parser.ObjTypeContext) {}

// EnterTypeName is called when production typeName is entered.
func (s *GraphQLListener) EnterTypeName(ctx *parser.TypeNameContext) {}

// ExitTypeName is called when production typeName is exited.
func (s *GraphQLListener) ExitTypeName(ctx *parser.TypeNameContext) {}

// EnterListType is called when production listType is entered.
func (s *GraphQLListener) EnterListType(ctx *parser.ListTypeContext) {}

// ExitListType is called when production listType is exited.
func (s *GraphQLListener) ExitListType(ctx *parser.ListTypeContext) {}

// EnterNonNullType is called when production nonNullType is entered.
func (s *GraphQLListener) EnterNonNullType(ctx *parser.NonNullTypeContext) {}

// ExitNonNullType is called when production nonNullType is exited.
func (s *GraphQLListener) ExitNonNullType(ctx *parser.NonNullTypeContext) {}

// EnterOperationDefinition is called when production operationDefinition is entered.
func (s *GraphQLListener) EnterOperationDefinition(ctx *parser.OperationDefinitionContext) {}

// ExitOperationDefinition is called when production operationDefinition is exited.
func (s *GraphQLListener) ExitOperationDefinition(ctx *parser.OperationDefinitionContext) {}

// EnterVariableDefinitions is called when production variableDefinitions is entered.
func (s *GraphQLListener) EnterVariableDefinitions(ctx *parser.VariableDefinitionsContext) {}

// ExitVariableDefinitions is called when production variableDefinitions is exited.
func (s *GraphQLListener) ExitVariableDefinitions(ctx *parser.VariableDefinitionsContext) {}

// EnterVariableDefinition is called when production variableDefinition is entered.
func (s *GraphQLListener) EnterVariableDefinition(ctx *parser.VariableDefinitionContext) {}

// ExitVariableDefinition is called when production variableDefinition is exited.
func (s *GraphQLListener) ExitVariableDefinition(ctx *parser.VariableDefinitionContext) {}

// EnterSelectionSet is called when production selectionSet is entered.
func (s *GraphQLListener) EnterSelectionSet(ctx *parser.SelectionSetContext) {}

// ExitSelectionSet is called when production selectionSet is exited.
func (s *GraphQLListener) ExitSelectionSet(ctx *parser.SelectionSetContext) {}

// EnterSelection is called when production selection is entered.
func (s *GraphQLListener) EnterSelection(ctx *parser.SelectionContext) {}

// ExitSelection is called when production selection is exited.
func (s *GraphQLListener) ExitSelection(ctx *parser.SelectionContext) {}

// EnterField is called when production field is entered.
func (s *GraphQLListener) EnterField(ctx *parser.FieldContext) {}

// ExitField is called when production field is exited.
func (s *GraphQLListener) ExitField(ctx *parser.FieldContext) {}

// EnterAlias is called when production alias is entered.
func (s *GraphQLListener) EnterAlias(ctx *parser.AliasContext) {}

// ExitAlias is called when production alias is exited.
func (s *GraphQLListener) ExitAlias(ctx *parser.AliasContext) {}

// EnterFragmentSpread is called when production fragmentSpread is entered.
func (s *GraphQLListener) EnterFragmentSpread(ctx *parser.FragmentSpreadContext) {}

// ExitFragmentSpread is called when production fragmentSpread is exited.
func (s *GraphQLListener) ExitFragmentSpread(ctx *parser.FragmentSpreadContext) {}

// EnterInlineFragment is called when production inlineFragment is entered.
func (s *GraphQLListener) EnterInlineFragment(ctx *parser.InlineFragmentContext) {}

// ExitInlineFragment is called when production inlineFragment is exited.
func (s *GraphQLListener) ExitInlineFragment(ctx *parser.InlineFragmentContext) {}

// EnterFragmentDefinition is called when production fragmentDefinition is entered.
func (s *GraphQLListener) EnterFragmentDefinition(ctx *parser.FragmentDefinitionContext) {}

// ExitFragmentDefinition is called when production fragmentDefinition is exited.
func (s *GraphQLListener) ExitFragmentDefinition(ctx *parser.FragmentDefinitionContext) {}

// EnterFragmentName is called when production fragmentName is entered.
func (s *GraphQLListener) EnterFragmentName(ctx *parser.FragmentNameContext) {}

// ExitFragmentName is called when production fragmentName is exited.
func (s *GraphQLListener) ExitFragmentName(ctx *parser.FragmentNameContext) {}

// EnterTypeCondition is called when production typeCondition is entered.
func (s *GraphQLListener) EnterTypeCondition(ctx *parser.TypeConditionContext) {}

// ExitTypeCondition is called when production typeCondition is exited.
func (s *GraphQLListener) ExitTypeCondition(ctx *parser.TypeConditionContext) {}

// EnterTypeSystemDefinition is called when production typeSystemDefinition is entered.
func (s *GraphQLListener) EnterTypeSystemDefinition(ctx *parser.TypeSystemDefinitionContext) {}

// ExitTypeSystemDefinition is called when production typeSystemDefinition is exited.
func (s *GraphQLListener) ExitTypeSystemDefinition(ctx *parser.TypeSystemDefinitionContext) {}

// EnterSchemaDefinition is called when production schemaDefinition is entered.
func (s *GraphQLListener) EnterSchemaDefinition(ctx *parser.SchemaDefinitionContext) { fmt.Println("EnterSchemaDefinition") }

// ExitSchemaDefinition is called when production schemaDefinition is exited.
func (s *GraphQLListener) ExitSchemaDefinition(ctx *parser.SchemaDefinitionContext) { fmt.Println("ExitSchemaDefinition") }

// EnterOperationTypeDefinition is called when production operationTypeDefinition is entered.
func (s *GraphQLListener) EnterOperationTypeDefinition(ctx *parser.OperationTypeDefinitionContext) { fmt.Println("EnterOperationTypeDefinition") }

// ExitOperationTypeDefinition is called when production operationTypeDefinition is exited.
func (s *GraphQLListener) ExitOperationTypeDefinition(ctx *parser.OperationTypeDefinitionContext) { fmt.Println("ExitOperationTypeDefinition") }

// EnterTypeDefinition is called when production typeDefinition is entered.
func (s *GraphQLListener) EnterTypeDefinition(ctx *parser.TypeDefinitionContext) {}

// ExitTypeDefinition is called when production typeDefinition is exited.
func (s *GraphQLListener) ExitTypeDefinition(ctx *parser.TypeDefinitionContext) {}

// EnterTypeExtension is called when production typeExtension is entered.
func (s *GraphQLListener) EnterTypeExtension(ctx *parser.TypeExtensionContext) {}

// ExitTypeExtension is called when production typeExtension is exited.
func (s *GraphQLListener) ExitTypeExtension(ctx *parser.TypeExtensionContext) {}

// EnterScalarTypeDefinition is called when production scalarTypeDefinition is entered.
func (s *GraphQLListener) EnterScalarTypeDefinition(ctx *parser.ScalarTypeDefinitionContext) {}

// ExitScalarTypeDefinition is called when production scalarTypeDefinition is exited.
func (s *GraphQLListener) ExitScalarTypeDefinition(ctx *parser.ScalarTypeDefinitionContext) {}

// EnterScalarTypeExtensionDefinition is called when production scalarTypeExtensionDefinition is entered.
func (s *GraphQLListener) EnterScalarTypeExtensionDefinition(ctx *parser.ScalarTypeExtensionDefinitionContext) {
}

// ExitScalarTypeExtensionDefinition is called when production scalarTypeExtensionDefinition is exited.
func (s *GraphQLListener) ExitScalarTypeExtensionDefinition(ctx *parser.ScalarTypeExtensionDefinitionContext) {
}

// EnterObjectTypeDefinition is called when production objectTypeDefinition is entered.
func (s *GraphQLListener) EnterObjectTypeDefinition(ctx *parser.ObjectTypeDefinitionContext) {}

// ExitObjectTypeDefinition is called when production objectTypeDefinition is exited.
func (s *GraphQLListener) ExitObjectTypeDefinition(ctx *parser.ObjectTypeDefinitionContext) {}

// EnterObjectTypeExtensionDefinition is called when production objectTypeExtensionDefinition is entered.
func (s *GraphQLListener) EnterObjectTypeExtensionDefinition(ctx *parser.ObjectTypeExtensionDefinitionContext) {
}

// ExitObjectTypeExtensionDefinition is called when production objectTypeExtensionDefinition is exited.
func (s *GraphQLListener) ExitObjectTypeExtensionDefinition(ctx *parser.ObjectTypeExtensionDefinitionContext) {
}

// EnterImplementsInterfaces is called when production implementsInterfaces is entered.
func (s *GraphQLListener) EnterImplementsInterfaces(ctx *parser.ImplementsInterfacesContext) {}

// ExitImplementsInterfaces is called when production implementsInterfaces is exited.
func (s *GraphQLListener) ExitImplementsInterfaces(ctx *parser.ImplementsInterfacesContext) {}

// EnterFieldsDefinition is called when production fieldsDefinition is entered.
func (s *GraphQLListener) EnterFieldsDefinition(ctx *parser.FieldsDefinitionContext) {}

// ExitFieldsDefinition is called when production fieldsDefinition is exited.
func (s *GraphQLListener) ExitFieldsDefinition(ctx *parser.FieldsDefinitionContext) {}

// EnterFieldDefinition is called when production fieldDefinition is entered.
func (s *GraphQLListener) EnterFieldDefinition(ctx *parser.FieldDefinitionContext) {}

// ExitFieldDefinition is called when production fieldDefinition is exited.
func (s *GraphQLListener) ExitFieldDefinition(ctx *parser.FieldDefinitionContext) {}

// EnterArgumentsDefinition is called when production argumentsDefinition is entered.
func (s *GraphQLListener) EnterArgumentsDefinition(ctx *parser.ArgumentsDefinitionContext) {}

// ExitArgumentsDefinition is called when production argumentsDefinition is exited.
func (s *GraphQLListener) ExitArgumentsDefinition(ctx *parser.ArgumentsDefinitionContext) {}

// EnterInputValueDefinition is called when production inputValueDefinition is entered.
func (s *GraphQLListener) EnterInputValueDefinition(ctx *parser.InputValueDefinitionContext) {}

// ExitInputValueDefinition is called when production inputValueDefinition is exited.
func (s *GraphQLListener) ExitInputValueDefinition(ctx *parser.InputValueDefinitionContext) {}

// EnterInterfaceTypeDefinition is called when production interfaceTypeDefinition is entered.
func (s *GraphQLListener) EnterInterfaceTypeDefinition(ctx *parser.InterfaceTypeDefinitionContext) {}

// ExitInterfaceTypeDefinition is called when production interfaceTypeDefinition is exited.
func (s *GraphQLListener) ExitInterfaceTypeDefinition(ctx *parser.InterfaceTypeDefinitionContext) {}

// EnterInterfaceTypeExtensionDefinition is called when production interfaceTypeExtensionDefinition is entered.
func (s *GraphQLListener) EnterInterfaceTypeExtensionDefinition(ctx *parser.InterfaceTypeExtensionDefinitionContext) {
}

// ExitInterfaceTypeExtensionDefinition is called when production interfaceTypeExtensionDefinition is exited.
func (s *GraphQLListener) ExitInterfaceTypeExtensionDefinition(ctx *parser.InterfaceTypeExtensionDefinitionContext) {
}

// EnterUnionTypeDefinition is called when production unionTypeDefinition is entered.
func (s *GraphQLListener) EnterUnionTypeDefinition(ctx *parser.UnionTypeDefinitionContext) {}

// ExitUnionTypeDefinition is called when production unionTypeDefinition is exited.
func (s *GraphQLListener) ExitUnionTypeDefinition(ctx *parser.UnionTypeDefinitionContext) {}

// EnterUnionTypeExtensionDefinition is called when production unionTypeExtensionDefinition is entered.
func (s *GraphQLListener) EnterUnionTypeExtensionDefinition(ctx *parser.UnionTypeExtensionDefinitionContext) {
}

// ExitUnionTypeExtensionDefinition is called when production unionTypeExtensionDefinition is exited.
func (s *GraphQLListener) ExitUnionTypeExtensionDefinition(ctx *parser.UnionTypeExtensionDefinitionContext) {
}

// EnterUnionMembership is called when production unionMembership is entered.
func (s *GraphQLListener) EnterUnionMembership(ctx *parser.UnionMembershipContext) {}

// ExitUnionMembership is called when production unionMembership is exited.
func (s *GraphQLListener) ExitUnionMembership(ctx *parser.UnionMembershipContext) {}

// EnterUnionMembers is called when production unionMembers is entered.
func (s *GraphQLListener) EnterUnionMembers(ctx *parser.UnionMembersContext) {}

// ExitUnionMembers is called when production unionMembers is exited.
func (s *GraphQLListener) ExitUnionMembers(ctx *parser.UnionMembersContext) {}

// EnterEnumTypeDefinition is called when production enumTypeDefinition is entered.
func (s *GraphQLListener) EnterEnumTypeDefinition(ctx *parser.EnumTypeDefinitionContext) {}

// ExitEnumTypeDefinition is called when production enumTypeDefinition is exited.
func (s *GraphQLListener) ExitEnumTypeDefinition(ctx *parser.EnumTypeDefinitionContext) {}

// EnterEnumTypeExtensionDefinition is called when production enumTypeExtensionDefinition is entered.
func (s *GraphQLListener) EnterEnumTypeExtensionDefinition(ctx *parser.EnumTypeExtensionDefinitionContext) {
}

// ExitEnumTypeExtensionDefinition is called when production enumTypeExtensionDefinition is exited.
func (s *GraphQLListener) ExitEnumTypeExtensionDefinition(ctx *parser.EnumTypeExtensionDefinitionContext) {
}

// EnterEnumValueDefinitions is called when production enumValueDefinitions is entered.
func (s *GraphQLListener) EnterEnumValueDefinitions(ctx *parser.EnumValueDefinitionsContext) {}

// ExitEnumValueDefinitions is called when production enumValueDefinitions is exited.
func (s *GraphQLListener) ExitEnumValueDefinitions(ctx *parser.EnumValueDefinitionsContext) {}

// EnterEnumValueDefinition is called when production enumValueDefinition is entered.
func (s *GraphQLListener) EnterEnumValueDefinition(ctx *parser.EnumValueDefinitionContext) {}

// ExitEnumValueDefinition is called when production enumValueDefinition is exited.
func (s *GraphQLListener) ExitEnumValueDefinition(ctx *parser.EnumValueDefinitionContext) {}

// EnterInputObjectTypeDefinition is called when production inputObjectTypeDefinition is entered.
func (s *GraphQLListener) EnterInputObjectTypeDefinition(ctx *parser.InputObjectTypeDefinitionContext) {}

// ExitInputObjectTypeDefinition is called when production inputObjectTypeDefinition is exited.
func (s *GraphQLListener) ExitInputObjectTypeDefinition(ctx *parser.InputObjectTypeDefinitionContext) {}

// EnterInputObjectTypeExtensionDefinition is called when production inputObjectTypeExtensionDefinition is entered.
func (s *GraphQLListener) EnterInputObjectTypeExtensionDefinition(ctx *parser.InputObjectTypeExtensionDefinitionContext) {
}

// ExitInputObjectTypeExtensionDefinition is called when production inputObjectTypeExtensionDefinition is exited.
func (s *GraphQLListener) ExitInputObjectTypeExtensionDefinition(ctx *parser.InputObjectTypeExtensionDefinitionContext) {
}

// EnterInputObjectValueDefinitions is called when production inputObjectValueDefinitions is entered.
func (s *GraphQLListener) EnterInputObjectValueDefinitions(ctx *parser.InputObjectValueDefinitionsContext) {
}

// ExitInputObjectValueDefinitions is called when production inputObjectValueDefinitions is exited.
func (s *GraphQLListener) ExitInputObjectValueDefinitions(ctx *parser.InputObjectValueDefinitionsContext) {
}

// EnterDirectiveDefinition is called when production directiveDefinition is entered.
func (s *GraphQLListener) EnterDirectiveDefinition(ctx *parser.DirectiveDefinitionContext) {}

// ExitDirectiveDefinition is called when production directiveDefinition is exited.
func (s *GraphQLListener) ExitDirectiveDefinition(ctx *parser.DirectiveDefinitionContext) {}

// EnterDirectiveLocation is called when production directiveLocation is entered.
func (s *GraphQLListener) EnterDirectiveLocation(ctx *parser.DirectiveLocationContext) {}

// ExitDirectiveLocation is called when production directiveLocation is exited.
func (s *GraphQLListener) ExitDirectiveLocation(ctx *parser.DirectiveLocationContext) {}

// EnterDirectiveLocations is called when production directiveLocations is entered.
func (s *GraphQLListener) EnterDirectiveLocations(ctx *parser.DirectiveLocationsContext) {}

// ExitDirectiveLocations is called when production directiveLocations is exited.
func (s *GraphQLListener) ExitDirectiveLocations(ctx *parser.DirectiveLocationsContext) {}
