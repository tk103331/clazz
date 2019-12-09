package data

// Java ClassFile versions (the minor version is stored in the 16 most significant bits, and the
// major version in the 16 least significant bits).
const V1_1 = 3<<16 | 45
const V1_2 = 0<<16 | 46
const V1_3 = 0<<16 | 47
const V1_4 = 0<<16 | 48
const V1_5 = 0<<16 | 49
const V1_6 = 0<<16 | 50
const V1_7 = 0<<16 | 51
const V1_8 = 0<<16 | 52
const V9 = 0<<16 | 53
const V10 = 0<<16 | 54
const V11 = 0<<16 | 55
const V12 = 0<<16 | 56
const V13 = 0<<16 | 57
const V14 = 0<<16 | 58

// Access flags values, defined in
// - https://docs.oracle.com/javase/specs/jvms/se9/html/jvms-4.html#jvms-4.1-200-E.1
// - https://docs.oracle.com/javase/specs/jvms/se9/html/jvms-4.html#jvms-4.5-200-A.1
// - https://docs.oracle.com/javase/specs/jvms/se9/html/jvms-4.html#jvms-4.6-200-A.1
// - https://docs.oracle.com/javase/specs/jvms/se9/html/jvms-4.html#jvms-4.7.25
const ACC_PUBLIC uint16 = 0x0001       // class, field, method
const ACC_PRIVATE uint16 = 0x0002      // class, field, method
const ACC_PROTECTED uint16 = 0x0004    // class, field, method
const ACC_STATIC uint16 = 0x0008       // field, method
const ACC_FINAL uint16 = 0x0010        // class, field, method, parameter
const ACC_SUPER uint16 = 0x0020        // class
const ACC_SYNCHRONIZED uint16 = 0x0020 // method
const ACC_OPEN uint16 = 0x0020         // module
const ACC_TRANSITIVE uint16 = 0x0020   // module requires
const ACC_VOLATILE uint16 = 0x0040     // field
const ACC_BRIDGE uint16 = 0x0040       // method
const ACC_STATIC_PHASE uint16 = 0x0040 // module requires
const ACC_VARARGS uint16 = 0x0080      // method
const ACC_TRANSIENT uint16 = 0x0080    // field
const ACC_NATIVE uint16 = 0x0100       // method
const ACC_constERFACE uint16 = 0x0200  // class
const ACC_ABSTRACT uint16 = 0x0400     // class, method
const ACC_STRICT uint16 = 0x0800       // method
const ACC_SYNTHETIC uint16 = 0x1000    // class, field, method, parameter, module *
const ACC_ANNOTATION uint16 = 0x2000   // class
const ACC_ENUM uint16 = 0x4000         // class(?) field inner
const ACC_MANDATED uint16 = 0x8000     // field, method, parameter, module, module *
const ACC_MODULE uint16 = 0x8000       // class

// The JVM opcode values (with the MethodVisitor method name used to visit them in comment, and
// where '-' means 'same method name as on the previous line').
// See https://docs.oracle.com/javase/specs/jvms/se9/html/jvms-6.html.
const NOP = 0                 // visitInsn
const ACONST_NULL = 1         // -
const ICONST_M1 = 2           // -
const ICONST_0 = 3            // -
const ICONST_1 = 4            // -
const ICONST_2 = 5            // -
const ICONST_3 = 6            // -
const ICONST_4 = 7            // -
const ICONST_5 = 8            // -
const LCONST_0 = 9            // -
const LCONST_1 = 10           // -
const FCONST_0 = 11           // -
const FCONST_1 = 12           // -
const FCONST_2 = 13           // -
const DCONST_0 = 14           // -
const DCONST_1 = 15           // -
const BIPUSH = 16             // visitconstInsn
const SIPUSH = 17             // -
const LDC = 18                // visitLdcInsn
const ILOAD = 21              // visitVarInsn
const LLOAD = 22              // -
const FLOAD = 23              // -
const DLOAD = 24              // -
const ALOAD = 25              // -
const IALOAD = 46             // visitInsn
const LALOAD = 47             // -
const FALOAD = 48             // -
const DALOAD = 49             // -
const AALOAD = 50             // -
const BALOAD = 51             // -
const CALOAD = 52             // -
const SALOAD = 53             // -
const ISTORE = 54             // visitVarInsn
const LSTORE = 55             // -
const FSTORE = 56             // -
const DSTORE = 57             // -
const ASTORE = 58             // -
const IASTORE = 79            // visitInsn
const LASTORE = 80            // -
const FASTORE = 81            // -
const DASTORE = 82            // -
const AASTORE = 83            // -
const BASTORE = 84            // -
const CASTORE = 85            // -
const SASTORE = 86            // -
const POP = 87                // -
const POP2 = 88               // -
const DUP = 89                // -
const DUP_X1 = 90             // -
const DUP_X2 = 91             // -
const DUP2 = 92               // -
const DUP2_X1 = 93            // -
const DUP2_X2 = 94            // -
const SWAP = 95               // -
const IADD = 96               // -
const LADD = 97               // -
const FADD = 98               // -
const DADD = 99               // -
const ISUB = 100              // -
const LSUB = 101              // -
const FSUB = 102              // -
const DSUB = 103              // -
const IMUL = 104              // -
const LMUL = 105              // -
const FMUL = 106              // -
const DMUL = 107              // -
const IDIV = 108              // -
const LDIV = 109              // -
const FDIV = 110              // -
const DDIV = 111              // -
const IREM = 112              // -
const LREM = 113              // -
const FREM = 114              // -
const DREM = 115              // -
const INEG = 116              // -
const LNEG = 117              // -
const FNEG = 118              // -
const DNEG = 119              // -
const ISHL = 120              // -
const LSHL = 121              // -
const ISHR = 122              // -
const LSHR = 123              // -
const IUSHR = 124             // -
const LUSHR = 125             // -
const IAND = 126              // -
const LAND = 127              // -
const IOR = 128               // -
const LOR = 129               // -
const IXOR = 130              // -
const LXOR = 131              // -
const IINC = 132              // visitIincInsn
const I2L = 133               // visitInsn
const I2F = 134               // -
const I2D = 135               // -
const L2I = 136               // -
const L2F = 137               // -
const L2D = 138               // -
const F2I = 139               // -
const F2L = 140               // -
const F2D = 141               // -
const D2I = 142               // -
const D2L = 143               // -
const D2F = 144               // -
const I2B = 145               // -
const I2C = 146               // -
const I2S = 147               // -
const LCMP = 148              // -
const FCMPL = 149             // -
const FCMPG = 150             // -
const DCMPL = 151             // -
const DCMPG = 152             // -
const IFEQ = 153              // visitJumpInsn
const IFNE = 154              // -
const IFLT = 155              // -
const IFGE = 156              // -
const IFGT = 157              // -
const IFLE = 158              // -
const IF_ICMPEQ = 159         // -
const IF_ICMPNE = 160         // -
const IF_ICMPLT = 161         // -
const IF_ICMPGE = 162         // -
const IF_ICMPGT = 163         // -
const IF_ICMPLE = 164         // -
const IF_ACMPEQ = 165         // -
const IF_ACMPNE = 166         // -
const GOTO = 167              // -
const JSR = 168               // -
const RET = 169               // visitVarInsn
const TABLESWITCH = 170       // visiTableSwitchInsn
const LOOKUPSWITCH = 171      // visitLookupSwitch
const IRETURN = 172           // visitInsn
const LRETURN = 173           // -
const FRETURN = 174           // -
const DRETURN = 175           // -
const ARETURN = 176           // -
const RETURN = 177            // -
const GETSTATIC = 178         // visitFieldInsn
const PUTSTATIC = 179         // -
const GETFIELD = 180          // -
const PUTFIELD = 181          // -
const INVOKEVIRTUAL = 182     // visitMethodInsn
const INVOKESPECIAL = 183     // -
const INVOKESTATIC = 184      // -
const INVOKEconstERFACE = 185 // -
const INVOKEDYNAMIC = 186     // visitInvokeDynamicInsn
const NEW = 187               // visitTypeInsn
const NEWARRAY = 188          // visitconstInsn
const ANEWARRAY = 189         // visitTypeInsn
const ARRAYLENGTH = 190       // visitInsn
const ATHROW = 191            // -
const CHECKCAST = 192         // visitTypeInsn
const INSTANCEOF = 193        // -
const MONITORENTER = 194      // visitInsn
const MONITOREXIT = 195       // -
const MULTIANEWARRAY = 197    // visitMultiANewArrayInsn
const IFNULL = 198            // visitJumpInsn
const IFNONNULL = 199         // -

const (
	HANDLE_GETFIELD uint8 = iota
	HANDLE_GETSTATIC
	HANDLE_PUTFIELD
	HANDLE_PUTSTATIC
	HANDLE_INVOKEVIRTUAL
	HANDLE_INVOKESTATIC
	HANDLE_INVOKESPECIAL
	HANDLE_NEWINVOKESPECIAL
	HANDLE_INVOKEINTERFACE
)
