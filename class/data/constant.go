package data

// The ClassFile attribute names, in the order they are defined in
// https://docs.oracle.com/javase/specs/jvms/se11/html/jvms-4.html#jvms-4.7-300.
const CONSTANT_VALUE = "ConstantValue"
const CODE = "Code"
const STACK_MAP_TABLE = "StackMapTable"
const EXCEPTIONS = "Exceptions"
const INNER_CLASSES = "InnerClasses"
const ENCLOSING_METHOD = "EnclosingMethod"
const SYNTHETIC = "Synthetic"
const SIGNATURE = "Signature"
const SOURCE_FILE = "SourceFile"
const SOURCE_DEBUG_EXTENSION = "SourceDebugExtension"
const LINE_NUMBER_TABLE = "LineNumberTable"
const LOCAL_VARIABLE_TABLE = "LocalVariableTable"
const LOCAL_VARIABLE_TYPE_TABLE = "LocalVariableTypeTable"
const DEPRECATED = "Deprecated"
const RUNTIME_VISIBLE_ANNOTATIONS = "RuntimeVisibleAnnotations"
const RUNTIME_INVISIBLE_ANNOTATIONS = "RuntimeInvisibleAnnotations"
const RUNTIME_VISIBLE_PARAMETER_ANNOTATIONS = "RuntimeVisibleParameterAnnotations"
const RUNTIME_INVISIBLE_PARAMETER_ANNOTATIONS = "RuntimeInvisibleParameterAnnotations"
const RUNTIME_VISIBLE_TYPE_ANNOTATIONS = "RuntimeVisibleTypeAnnotations"
const RUNTIME_INVISIBLE_TYPE_ANNOTATIONS = "RuntimeInvisibleTypeAnnotations"
const ANNOTATION_DEFAULT = "AnnotationDefault"
const BOOTSTRAP_METHODS = "BootstrapMethods"
const METHOD_PARAMETERS = "MethodParameters"
const MODULE = "Module"
const MODULE_PACKAGES = "ModulePackages"
const MODULE_MAIN_CLASS = "ModuleMainClass"
const NEST_HOST = "NestHost"
const NEST_MEMBERS = "NestMembers"
const PERMITTED_SUBTYPES = "PermittedSubtypes"
const RECORD = "Record"

const (
	TYPE_SORT_VOID int = iota
	TYPE_SORT_BOOLEAN
	TYPE_SORT_HAR
	TYPE_SORT_BYTE
	TYPE_SORT_SHORT
	TYPE_SORT_INT
	TYPE_SORT_FLOAT
	TYPE_SORT_LONG
	TYPE_SORT_DOUBLE
	TYPE_SORT_ARRAY
	TYPE_SORT_OBJECT
	TYPE_SORT_METHOD
	TYPE_SORT_INTERNAL
)
