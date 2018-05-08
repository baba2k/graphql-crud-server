// Code generated from /home/baba/dev/workspace/go/src/github.com/baba2k/graphql-rungen/parser/grammar/GraphQL.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // GraphQL

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseGraphQLListener is a complete listener for a parse tree produced by GraphQLParser.
type BaseGraphQLListener struct{}

var _ GraphQLListener = &BaseGraphQLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGraphQLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGraphQLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGraphQLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGraphQLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterDocument is called when production document is entered.
func (s *BaseGraphQLListener) EnterDocument(ctx *DocumentContext) {}

// ExitDocument is called when production document is exited.
func (s *BaseGraphQLListener) ExitDocument(ctx *DocumentContext) {}

// EnterDefinition is called when production definition is entered.
func (s *BaseGraphQLListener) EnterDefinition(ctx *DefinitionContext) {}

// ExitDefinition is called when production definition is exited.
func (s *BaseGraphQLListener) ExitDefinition(ctx *DefinitionContext) {}

// EnterOperationType is called when production operationType is entered.
func (s *BaseGraphQLListener) EnterOperationType(ctx *OperationTypeContext) {}

// ExitOperationType is called when production operationType is exited.
func (s *BaseGraphQLListener) ExitOperationType(ctx *OperationTypeContext) {}

// EnterDescription is called when production description is entered.
func (s *BaseGraphQLListener) EnterDescription(ctx *DescriptionContext) {}

// ExitDescription is called when production description is exited.
func (s *BaseGraphQLListener) ExitDescription(ctx *DescriptionContext) {}

// EnterEnumValue is called when production enumValue is entered.
func (s *BaseGraphQLListener) EnterEnumValue(ctx *EnumValueContext) {}

// ExitEnumValue is called when production enumValue is exited.
func (s *BaseGraphQLListener) ExitEnumValue(ctx *EnumValueContext) {}

// EnterArrayValue is called when production arrayValue is entered.
func (s *BaseGraphQLListener) EnterArrayValue(ctx *ArrayValueContext) {}

// ExitArrayValue is called when production arrayValue is exited.
func (s *BaseGraphQLListener) ExitArrayValue(ctx *ArrayValueContext) {}

// EnterArrayValueWithVariable is called when production arrayValueWithVariable is entered.
func (s *BaseGraphQLListener) EnterArrayValueWithVariable(ctx *ArrayValueWithVariableContext) {}

// ExitArrayValueWithVariable is called when production arrayValueWithVariable is exited.
func (s *BaseGraphQLListener) ExitArrayValueWithVariable(ctx *ArrayValueWithVariableContext) {}

// EnterObjectValue is called when production objectValue is entered.
func (s *BaseGraphQLListener) EnterObjectValue(ctx *ObjectValueContext) {}

// ExitObjectValue is called when production objectValue is exited.
func (s *BaseGraphQLListener) ExitObjectValue(ctx *ObjectValueContext) {}

// EnterObjectValueWithVariable is called when production objectValueWithVariable is entered.
func (s *BaseGraphQLListener) EnterObjectValueWithVariable(ctx *ObjectValueWithVariableContext) {}

// ExitObjectValueWithVariable is called when production objectValueWithVariable is exited.
func (s *BaseGraphQLListener) ExitObjectValueWithVariable(ctx *ObjectValueWithVariableContext) {}

// EnterObjectField is called when production objectField is entered.
func (s *BaseGraphQLListener) EnterObjectField(ctx *ObjectFieldContext) {}

// ExitObjectField is called when production objectField is exited.
func (s *BaseGraphQLListener) ExitObjectField(ctx *ObjectFieldContext) {}

// EnterObjectFieldWithVariable is called when production objectFieldWithVariable is entered.
func (s *BaseGraphQLListener) EnterObjectFieldWithVariable(ctx *ObjectFieldWithVariableContext) {}

// ExitObjectFieldWithVariable is called when production objectFieldWithVariable is exited.
func (s *BaseGraphQLListener) ExitObjectFieldWithVariable(ctx *ObjectFieldWithVariableContext) {}

// EnterDirectives is called when production directives is entered.
func (s *BaseGraphQLListener) EnterDirectives(ctx *DirectivesContext) {}

// ExitDirectives is called when production directives is exited.
func (s *BaseGraphQLListener) ExitDirectives(ctx *DirectivesContext) {}

// EnterDirective is called when production directive is entered.
func (s *BaseGraphQLListener) EnterDirective(ctx *DirectiveContext) {}

// ExitDirective is called when production directive is exited.
func (s *BaseGraphQLListener) ExitDirective(ctx *DirectiveContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseGraphQLListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseGraphQLListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterArgument is called when production argument is entered.
func (s *BaseGraphQLListener) EnterArgument(ctx *ArgumentContext) {}

// ExitArgument is called when production argument is exited.
func (s *BaseGraphQLListener) ExitArgument(ctx *ArgumentContext) {}

// EnterName is called when production name is entered.
func (s *BaseGraphQLListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BaseGraphQLListener) ExitName(ctx *NameContext) {}

// EnterValue is called when production value is entered.
func (s *BaseGraphQLListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseGraphQLListener) ExitValue(ctx *ValueContext) {}

// EnterValueWithVariable is called when production valueWithVariable is entered.
func (s *BaseGraphQLListener) EnterValueWithVariable(ctx *ValueWithVariableContext) {}

// ExitValueWithVariable is called when production valueWithVariable is exited.
func (s *BaseGraphQLListener) ExitValueWithVariable(ctx *ValueWithVariableContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseGraphQLListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseGraphQLListener) ExitVariable(ctx *VariableContext) {}

// EnterDefaultValue is called when production defaultValue is entered.
func (s *BaseGraphQLListener) EnterDefaultValue(ctx *DefaultValueContext) {}

// ExitDefaultValue is called when production defaultValue is exited.
func (s *BaseGraphQLListener) ExitDefaultValue(ctx *DefaultValueContext) {}

// EnterStringValue is called when production stringValue is entered.
func (s *BaseGraphQLListener) EnterStringValue(ctx *StringValueContext) {}

// ExitStringValue is called when production stringValue is exited.
func (s *BaseGraphQLListener) ExitStringValue(ctx *StringValueContext) {}

// EnterObjType is called when production objType is entered.
func (s *BaseGraphQLListener) EnterObjType(ctx *ObjTypeContext) {}

// ExitObjType is called when production objType is exited.
func (s *BaseGraphQLListener) ExitObjType(ctx *ObjTypeContext) {}

// EnterTypeName is called when production typeName is entered.
func (s *BaseGraphQLListener) EnterTypeName(ctx *TypeNameContext) {}

// ExitTypeName is called when production typeName is exited.
func (s *BaseGraphQLListener) ExitTypeName(ctx *TypeNameContext) {}

// EnterListType is called when production listType is entered.
func (s *BaseGraphQLListener) EnterListType(ctx *ListTypeContext) {}

// ExitListType is called when production listType is exited.
func (s *BaseGraphQLListener) ExitListType(ctx *ListTypeContext) {}

// EnterNonNullType is called when production nonNullType is entered.
func (s *BaseGraphQLListener) EnterNonNullType(ctx *NonNullTypeContext) {}

// ExitNonNullType is called when production nonNullType is exited.
func (s *BaseGraphQLListener) ExitNonNullType(ctx *NonNullTypeContext) {}

// EnterOperationDefinition is called when production operationDefinition is entered.
func (s *BaseGraphQLListener) EnterOperationDefinition(ctx *OperationDefinitionContext) {}

// ExitOperationDefinition is called when production operationDefinition is exited.
func (s *BaseGraphQLListener) ExitOperationDefinition(ctx *OperationDefinitionContext) {}

// EnterVariableDefinitions is called when production variableDefinitions is entered.
func (s *BaseGraphQLListener) EnterVariableDefinitions(ctx *VariableDefinitionsContext) {}

// ExitVariableDefinitions is called when production variableDefinitions is exited.
func (s *BaseGraphQLListener) ExitVariableDefinitions(ctx *VariableDefinitionsContext) {}

// EnterVariableDefinition is called when production variableDefinition is entered.
func (s *BaseGraphQLListener) EnterVariableDefinition(ctx *VariableDefinitionContext) {}

// ExitVariableDefinition is called when production variableDefinition is exited.
func (s *BaseGraphQLListener) ExitVariableDefinition(ctx *VariableDefinitionContext) {}

// EnterSelectionSet is called when production selectionSet is entered.
func (s *BaseGraphQLListener) EnterSelectionSet(ctx *SelectionSetContext) {}

// ExitSelectionSet is called when production selectionSet is exited.
func (s *BaseGraphQLListener) ExitSelectionSet(ctx *SelectionSetContext) {}

// EnterSelection is called when production selection is entered.
func (s *BaseGraphQLListener) EnterSelection(ctx *SelectionContext) {}

// ExitSelection is called when production selection is exited.
func (s *BaseGraphQLListener) ExitSelection(ctx *SelectionContext) {}

// EnterField is called when production field is entered.
func (s *BaseGraphQLListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseGraphQLListener) ExitField(ctx *FieldContext) {}

// EnterAlias is called when production alias is entered.
func (s *BaseGraphQLListener) EnterAlias(ctx *AliasContext) {}

// ExitAlias is called when production alias is exited.
func (s *BaseGraphQLListener) ExitAlias(ctx *AliasContext) {}

// EnterFragmentSpread is called when production fragmentSpread is entered.
func (s *BaseGraphQLListener) EnterFragmentSpread(ctx *FragmentSpreadContext) {}

// ExitFragmentSpread is called when production fragmentSpread is exited.
func (s *BaseGraphQLListener) ExitFragmentSpread(ctx *FragmentSpreadContext) {}

// EnterInlineFragment is called when production inlineFragment is entered.
func (s *BaseGraphQLListener) EnterInlineFragment(ctx *InlineFragmentContext) {}

// ExitInlineFragment is called when production inlineFragment is exited.
func (s *BaseGraphQLListener) ExitInlineFragment(ctx *InlineFragmentContext) {}

// EnterFragmentDefinition is called when production fragmentDefinition is entered.
func (s *BaseGraphQLListener) EnterFragmentDefinition(ctx *FragmentDefinitionContext) {}

// ExitFragmentDefinition is called when production fragmentDefinition is exited.
func (s *BaseGraphQLListener) ExitFragmentDefinition(ctx *FragmentDefinitionContext) {}

// EnterFragmentName is called when production fragmentName is entered.
func (s *BaseGraphQLListener) EnterFragmentName(ctx *FragmentNameContext) {}

// ExitFragmentName is called when production fragmentName is exited.
func (s *BaseGraphQLListener) ExitFragmentName(ctx *FragmentNameContext) {}

// EnterTypeCondition is called when production typeCondition is entered.
func (s *BaseGraphQLListener) EnterTypeCondition(ctx *TypeConditionContext) {}

// ExitTypeCondition is called when production typeCondition is exited.
func (s *BaseGraphQLListener) ExitTypeCondition(ctx *TypeConditionContext) {}

// EnterTypeSystemDefinition is called when production typeSystemDefinition is entered.
func (s *BaseGraphQLListener) EnterTypeSystemDefinition(ctx *TypeSystemDefinitionContext) {}

// ExitTypeSystemDefinition is called when production typeSystemDefinition is exited.
func (s *BaseGraphQLListener) ExitTypeSystemDefinition(ctx *TypeSystemDefinitionContext) {}

// EnterSchemaDefinition is called when production schemaDefinition is entered.
func (s *BaseGraphQLListener) EnterSchemaDefinition(ctx *SchemaDefinitionContext) {}

// ExitSchemaDefinition is called when production schemaDefinition is exited.
func (s *BaseGraphQLListener) ExitSchemaDefinition(ctx *SchemaDefinitionContext) {}

// EnterOperationTypeDefinition is called when production operationTypeDefinition is entered.
func (s *BaseGraphQLListener) EnterOperationTypeDefinition(ctx *OperationTypeDefinitionContext) {}

// ExitOperationTypeDefinition is called when production operationTypeDefinition is exited.
func (s *BaseGraphQLListener) ExitOperationTypeDefinition(ctx *OperationTypeDefinitionContext) {}

// EnterTypeDefinition is called when production typeDefinition is entered.
func (s *BaseGraphQLListener) EnterTypeDefinition(ctx *TypeDefinitionContext) {}

// ExitTypeDefinition is called when production typeDefinition is exited.
func (s *BaseGraphQLListener) ExitTypeDefinition(ctx *TypeDefinitionContext) {}

// EnterTypeExtension is called when production typeExtension is entered.
func (s *BaseGraphQLListener) EnterTypeExtension(ctx *TypeExtensionContext) {}

// ExitTypeExtension is called when production typeExtension is exited.
func (s *BaseGraphQLListener) ExitTypeExtension(ctx *TypeExtensionContext) {}

// EnterScalarTypeDefinition is called when production scalarTypeDefinition is entered.
func (s *BaseGraphQLListener) EnterScalarTypeDefinition(ctx *ScalarTypeDefinitionContext) {}

// ExitScalarTypeDefinition is called when production scalarTypeDefinition is exited.
func (s *BaseGraphQLListener) ExitScalarTypeDefinition(ctx *ScalarTypeDefinitionContext) {}

// EnterScalarTypeExtensionDefinition is called when production scalarTypeExtensionDefinition is entered.
func (s *BaseGraphQLListener) EnterScalarTypeExtensionDefinition(ctx *ScalarTypeExtensionDefinitionContext) {
}

// ExitScalarTypeExtensionDefinition is called when production scalarTypeExtensionDefinition is exited.
func (s *BaseGraphQLListener) ExitScalarTypeExtensionDefinition(ctx *ScalarTypeExtensionDefinitionContext) {
}

// EnterObjectTypeDefinition is called when production objectTypeDefinition is entered.
func (s *BaseGraphQLListener) EnterObjectTypeDefinition(ctx *ObjectTypeDefinitionContext) {}

// ExitObjectTypeDefinition is called when production objectTypeDefinition is exited.
func (s *BaseGraphQLListener) ExitObjectTypeDefinition(ctx *ObjectTypeDefinitionContext) {}

// EnterObjectTypeExtensionDefinition is called when production objectTypeExtensionDefinition is entered.
func (s *BaseGraphQLListener) EnterObjectTypeExtensionDefinition(ctx *ObjectTypeExtensionDefinitionContext) {
}

// ExitObjectTypeExtensionDefinition is called when production objectTypeExtensionDefinition is exited.
func (s *BaseGraphQLListener) ExitObjectTypeExtensionDefinition(ctx *ObjectTypeExtensionDefinitionContext) {
}

// EnterImplementsInterfaces is called when production implementsInterfaces is entered.
func (s *BaseGraphQLListener) EnterImplementsInterfaces(ctx *ImplementsInterfacesContext) {}

// ExitImplementsInterfaces is called when production implementsInterfaces is exited.
func (s *BaseGraphQLListener) ExitImplementsInterfaces(ctx *ImplementsInterfacesContext) {}

// EnterFieldsDefinition is called when production fieldsDefinition is entered.
func (s *BaseGraphQLListener) EnterFieldsDefinition(ctx *FieldsDefinitionContext) {}

// ExitFieldsDefinition is called when production fieldsDefinition is exited.
func (s *BaseGraphQLListener) ExitFieldsDefinition(ctx *FieldsDefinitionContext) {}

// EnterFieldDefinition is called when production fieldDefinition is entered.
func (s *BaseGraphQLListener) EnterFieldDefinition(ctx *FieldDefinitionContext) {}

// ExitFieldDefinition is called when production fieldDefinition is exited.
func (s *BaseGraphQLListener) ExitFieldDefinition(ctx *FieldDefinitionContext) {}

// EnterArgumentsDefinition is called when production argumentsDefinition is entered.
func (s *BaseGraphQLListener) EnterArgumentsDefinition(ctx *ArgumentsDefinitionContext) {}

// ExitArgumentsDefinition is called when production argumentsDefinition is exited.
func (s *BaseGraphQLListener) ExitArgumentsDefinition(ctx *ArgumentsDefinitionContext) {}

// EnterInputValueDefinition is called when production inputValueDefinition is entered.
func (s *BaseGraphQLListener) EnterInputValueDefinition(ctx *InputValueDefinitionContext) {}

// ExitInputValueDefinition is called when production inputValueDefinition is exited.
func (s *BaseGraphQLListener) ExitInputValueDefinition(ctx *InputValueDefinitionContext) {}

// EnterInterfaceTypeDefinition is called when production interfaceTypeDefinition is entered.
func (s *BaseGraphQLListener) EnterInterfaceTypeDefinition(ctx *InterfaceTypeDefinitionContext) {}

// ExitInterfaceTypeDefinition is called when production interfaceTypeDefinition is exited.
func (s *BaseGraphQLListener) ExitInterfaceTypeDefinition(ctx *InterfaceTypeDefinitionContext) {}

// EnterInterfaceTypeExtensionDefinition is called when production interfaceTypeExtensionDefinition is entered.
func (s *BaseGraphQLListener) EnterInterfaceTypeExtensionDefinition(ctx *InterfaceTypeExtensionDefinitionContext) {
}

// ExitInterfaceTypeExtensionDefinition is called when production interfaceTypeExtensionDefinition is exited.
func (s *BaseGraphQLListener) ExitInterfaceTypeExtensionDefinition(ctx *InterfaceTypeExtensionDefinitionContext) {
}

// EnterUnionTypeDefinition is called when production unionTypeDefinition is entered.
func (s *BaseGraphQLListener) EnterUnionTypeDefinition(ctx *UnionTypeDefinitionContext) {}

// ExitUnionTypeDefinition is called when production unionTypeDefinition is exited.
func (s *BaseGraphQLListener) ExitUnionTypeDefinition(ctx *UnionTypeDefinitionContext) {}

// EnterUnionTypeExtensionDefinition is called when production unionTypeExtensionDefinition is entered.
func (s *BaseGraphQLListener) EnterUnionTypeExtensionDefinition(ctx *UnionTypeExtensionDefinitionContext) {
}

// ExitUnionTypeExtensionDefinition is called when production unionTypeExtensionDefinition is exited.
func (s *BaseGraphQLListener) ExitUnionTypeExtensionDefinition(ctx *UnionTypeExtensionDefinitionContext) {
}

// EnterUnionMembership is called when production unionMembership is entered.
func (s *BaseGraphQLListener) EnterUnionMembership(ctx *UnionMembershipContext) {}

// ExitUnionMembership is called when production unionMembership is exited.
func (s *BaseGraphQLListener) ExitUnionMembership(ctx *UnionMembershipContext) {}

// EnterUnionMembers is called when production unionMembers is entered.
func (s *BaseGraphQLListener) EnterUnionMembers(ctx *UnionMembersContext) {}

// ExitUnionMembers is called when production unionMembers is exited.
func (s *BaseGraphQLListener) ExitUnionMembers(ctx *UnionMembersContext) {}

// EnterEnumTypeDefinition is called when production enumTypeDefinition is entered.
func (s *BaseGraphQLListener) EnterEnumTypeDefinition(ctx *EnumTypeDefinitionContext) {}

// ExitEnumTypeDefinition is called when production enumTypeDefinition is exited.
func (s *BaseGraphQLListener) ExitEnumTypeDefinition(ctx *EnumTypeDefinitionContext) {}

// EnterEnumTypeExtensionDefinition is called when production enumTypeExtensionDefinition is entered.
func (s *BaseGraphQLListener) EnterEnumTypeExtensionDefinition(ctx *EnumTypeExtensionDefinitionContext) {
}

// ExitEnumTypeExtensionDefinition is called when production enumTypeExtensionDefinition is exited.
func (s *BaseGraphQLListener) ExitEnumTypeExtensionDefinition(ctx *EnumTypeExtensionDefinitionContext) {
}

// EnterEnumValueDefinitions is called when production enumValueDefinitions is entered.
func (s *BaseGraphQLListener) EnterEnumValueDefinitions(ctx *EnumValueDefinitionsContext) {}

// ExitEnumValueDefinitions is called when production enumValueDefinitions is exited.
func (s *BaseGraphQLListener) ExitEnumValueDefinitions(ctx *EnumValueDefinitionsContext) {}

// EnterEnumValueDefinition is called when production enumValueDefinition is entered.
func (s *BaseGraphQLListener) EnterEnumValueDefinition(ctx *EnumValueDefinitionContext) {}

// ExitEnumValueDefinition is called when production enumValueDefinition is exited.
func (s *BaseGraphQLListener) ExitEnumValueDefinition(ctx *EnumValueDefinitionContext) {}

// EnterInputObjectTypeDefinition is called when production inputObjectTypeDefinition is entered.
func (s *BaseGraphQLListener) EnterInputObjectTypeDefinition(ctx *InputObjectTypeDefinitionContext) {}

// ExitInputObjectTypeDefinition is called when production inputObjectTypeDefinition is exited.
func (s *BaseGraphQLListener) ExitInputObjectTypeDefinition(ctx *InputObjectTypeDefinitionContext) {}

// EnterInputObjectTypeExtensionDefinition is called when production inputObjectTypeExtensionDefinition is entered.
func (s *BaseGraphQLListener) EnterInputObjectTypeExtensionDefinition(ctx *InputObjectTypeExtensionDefinitionContext) {
}

// ExitInputObjectTypeExtensionDefinition is called when production inputObjectTypeExtensionDefinition is exited.
func (s *BaseGraphQLListener) ExitInputObjectTypeExtensionDefinition(ctx *InputObjectTypeExtensionDefinitionContext) {
}

// EnterInputObjectValueDefinitions is called when production inputObjectValueDefinitions is entered.
func (s *BaseGraphQLListener) EnterInputObjectValueDefinitions(ctx *InputObjectValueDefinitionsContext) {
}

// ExitInputObjectValueDefinitions is called when production inputObjectValueDefinitions is exited.
func (s *BaseGraphQLListener) ExitInputObjectValueDefinitions(ctx *InputObjectValueDefinitionsContext) {
}

// EnterDirectiveDefinition is called when production directiveDefinition is entered.
func (s *BaseGraphQLListener) EnterDirectiveDefinition(ctx *DirectiveDefinitionContext) {}

// ExitDirectiveDefinition is called when production directiveDefinition is exited.
func (s *BaseGraphQLListener) ExitDirectiveDefinition(ctx *DirectiveDefinitionContext) {}

// EnterDirectiveLocation is called when production directiveLocation is entered.
func (s *BaseGraphQLListener) EnterDirectiveLocation(ctx *DirectiveLocationContext) {}

// ExitDirectiveLocation is called when production directiveLocation is exited.
func (s *BaseGraphQLListener) ExitDirectiveLocation(ctx *DirectiveLocationContext) {}

// EnterDirectiveLocations is called when production directiveLocations is entered.
func (s *BaseGraphQLListener) EnterDirectiveLocations(ctx *DirectiveLocationsContext) {}

// ExitDirectiveLocations is called when production directiveLocations is exited.
func (s *BaseGraphQLListener) ExitDirectiveLocations(ctx *DirectiveLocationsContext) {}
