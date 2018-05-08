grammar GraphQL;

document : definition+;

definition : operationDefinition | fragmentDefinition | typeSystemDefinition;

operationType : SUBSCRIPTION | MUTATION | QUERY;

description : stringValue;

enumValue : name;

arrayValue : '[' value* ']';
arrayValueWithVariable : '[' valueWithVariable* ']';

objectValue : '{' objectField* '}';
objectValueWithVariable : '{' objectFieldWithVariable* '}';
objectField : name ':' value;
objectFieldWithVariable : name ':' valueWithVariable;

directives : directive+;
directive :'@' name arguments?;

arguments : '(' argument+ ')';
argument : name ':' valueWithVariable;

name : NAME | FRAGMENT | QUERY | MUTATION | SUBSCRIPTION | SCHEMA | SCALAR | TYPE | INTERFACE | IMPLEMENTS | ENUM | UNION | INPUT | EXTEND | DIRECTIVE;

value : stringValue | IntValue | FloatValue | BooleanValue | NullValue | enumValue | arrayValue | objectValue;

valueWithVariable : variable | stringValue | IntValue | FloatValue | BooleanValue | NullValue | enumValue | arrayValueWithVariable | objectValueWithVariable;

variable : '$' name;

defaultValue : '=' value;

stringValue : TripleQuotedStringValue | StringValue;
objType : typeName | listType | nonNullType;

typeName : name;
listType : '[' objType ']';
nonNullType : typeName '!' | listType '!';

FRAGMENT : 'fragment';
QUERY : 'query';
MUTATION : 'mutation';
SUBSCRIPTION : 'subscription';
SCHEMA : 'schema';
SCALAR : 'scalar';
TYPE : 'type';
INTERFACE : 'interface';
IMPLEMENTS : 'implements';
ENUM : 'enum';
UNION : 'union';
INPUT : 'input';
EXTEND : 'extend';
DIRECTIVE : 'directive';
NAME : [_A-Za-z][_0-9A-Za-z]*;

BooleanValue : 'true' | 'false';

NullValue : 'null';

IntValue : Sign? IntegerPart;

FloatValue : Sign? IntegerPart ('.' Digit+)? ExponentPart?;

Sign : '-';

IntegerPart : '0' | NonZeroDigit | NonZeroDigit Digit+;

NonZeroDigit: '1'.. '9';

ExponentPart : ('e'|'E') Sign? Digit+;

Digit : '0'..'9';

StringValue : '"' ( ~["\\\n\r\u2028\u2029] | EscapedChar )* '"';

TripleQuotedStringValue : '"""' TripleQuotedStringPart? '"""';

fragment TripleQuotedStringPart : ( EscapedTripleQuote | SourceCharacter )+?;
fragment EscapedTripleQuote : '\\"""';
fragment SourceCharacter : [\u0009\u000A\u000D\u0020-\uFFFF];
fragment EscapedChar : '\\' (["\\/bfnrt] | Unicode);
fragment Unicode : 'u' Hex Hex Hex Hex;
fragment Hex : [0-9a-fA-F];
fragment LineTerminator: [\n\r\u2028\u2029];
fragment Whitespace : [\u0009\u0020];
fragment Comma : ',';
fragment UnicodeBOM : [\ufeff];
Comment: '#' ~[\n\r\u2028\u2029]* -> channel(2);
Ignored: (UnicodeBOM|Whitespace|LineTerminator|Comma) -> skip;

operationDefinition : selectionSet | operationType name? variableDefinitions? directives? selectionSet;

variableDefinitions : '(' variableDefinition+ ')';

variableDefinition : variable ':' objType defaultValue?;

selectionSet : '{' selection+ '}';

selection : field | fragmentSpread | inlineFragment;

field : alias? name arguments? directives? selectionSet?;

alias : name ':';

fragmentSpread : '...' fragmentName directives?;

inlineFragment : '...' typeCondition? directives? selectionSet;

fragmentDefinition : 'fragment' fragmentName typeCondition directives? selectionSet;

fragmentName : name;

typeCondition : 'on' typeName;

typeSystemDefinition : description? schemaDefinition | typeDefinition | typeExtension | directiveDefinition;

schemaDefinition : description? SCHEMA directives? '{' operationTypeDefinition+ '}';

operationTypeDefinition : description? operationType ':' typeName;

typeDefinition: scalarTypeDefinition | objectTypeDefinition | interfaceTypeDefinition | unionTypeDefinition | enumTypeDefinition | inputObjectTypeDefinition;

typeExtension : objectTypeExtensionDefinition | interfaceTypeExtensionDefinition | unionTypeExtensionDefinition | scalarTypeExtensionDefinition | enumTypeExtensionDefinition | inputObjectTypeExtensionDefinition;

scalarTypeDefinition : description? SCALAR name directives?;

scalarTypeExtensionDefinition : EXTEND SCALAR name directives?;

objectTypeDefinition : description? TYPE name implementsInterfaces? directives? fieldsDefinition?;

objectTypeExtensionDefinition : EXTEND TYPE name implementsInterfaces? directives? fieldsDefinition?;

implementsInterfaces : IMPLEMENTS '&'? typeName+ | implementsInterfaces '&' typeName;

fieldsDefinition : '{' fieldDefinition+ '}';

fieldDefinition : description? name argumentsDefinition? ':' objType directives?;

argumentsDefinition : '(' inputValueDefinition+ ')';

inputValueDefinition : description? name ':' objType defaultValue? directives?;

interfaceTypeDefinition : description? INTERFACE name directives? fieldsDefinition?;

interfaceTypeExtensionDefinition : EXTEND INTERFACE name directives? fieldsDefinition?;

unionTypeDefinition : description? UNION name directives? unionMembership;

unionTypeExtensionDefinition : EXTEND UNION name directives? unionMembership?;

unionMembership : '=' unionMembers;

unionMembers : '|'? typeName | unionMembers '|' typeName;

enumTypeDefinition : description? ENUM name directives? enumValueDefinitions;

enumTypeExtensionDefinition : EXTEND ENUM name directives? enumValueDefinitions?;

enumValueDefinitions : '{' enumValueDefinition+ '}';

enumValueDefinition : description? enumValue directives?;

inputObjectTypeDefinition : description? INPUT name directives? inputObjectValueDefinitions;

inputObjectTypeExtensionDefinition : EXTEND INPUT name directives? inputObjectValueDefinitions?;

inputObjectValueDefinitions : '{' inputValueDefinition+ '}';

directiveDefinition : description? DIRECTIVE '@' name argumentsDefinition? 'on' directiveLocations;

directiveLocation : name;

directiveLocations : directiveLocation | directiveLocations '|' directiveLocation;