package lexer

type Token interface {}

type LToken struct {
	Filename string
	Line, Col int
	v Token
}

// Special
type TkNewline struct {} // \n or ;;
type TkOpenR struct {} // (
type TkCloseR struct {} // )
type TkOpenC struct {} // {
type TkCloseC struct {} // }
type TkOpenS struct {} // [
type TkCloseS struct {} // ]
type TkComma struct {} // ,
type TkSep struct {} // ;
// Word
type TkLit struct { Raw string } // int/float/rune/string
type TkId struct { Name string } // identifier
type TkLabel struct { Label string }
// Symbol
type TkOp struct { Name string } // operator
type TkFn struct {} // =>
type TkAssign struct {} // =
type TkTypeDef struct {} // :=
type TkTypeDecl struct {} // :
type TkTypeCast struct {} // ::
type TkTypeAsrt struct {} // :?
type TkMatch struct {} // ?
type TkThen struct {} // ->
type TkConst struct {} // $
type TkStruct struct {} // @
type TkItf struct {} // @@
type TkThen struct {} // ->
// Keywords
type TkNil struct {}
type TkTrue struct {}
type TkFalse struct {}
type TkPackage struct {}
type TkImport struct {}
type TkIf struct {}
type TkElse struct {}
type TkSwitch struct {}
type TkFor struct {}
type TkReturn struct {}
type TkDefer struct {}
type TkBreak struct {}
type TkContinue struct {}
type TkFallthrough struct {}
type TkGoto struct {}
type TkGo struct {}
type TkSelect struct {}
type TkMap struct {}
type TkChan struct {}
