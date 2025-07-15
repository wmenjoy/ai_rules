package core

import (
	"fmt"
	"go/token"
	"strings"
	"text/scanner"
	"unicode"
)

// JavaToken represents different types of Java tokens
type JavaToken int

const (
	TOKEN_EOF JavaToken = iota
	TOKEN_IDENTIFIER
	TOKEN_KEYWORD
	TOKEN_OPERATOR
	TOKEN_DELIMITER
	TOKEN_LITERAL
	TOKEN_COMMENT
	TOKEN_WHITESPACE
)

// JavaLexer tokenizes Java source code
type JavaLexer struct {
	scanner scanner.Scanner
	current rune
	position token.Pos
	text     string
}

// JavaKeywords contains Java language keywords
var JavaKeywords = map[string]bool{
	"abstract":     true,
	"assert":       true,
	"boolean":      true,
	"break":        true,
	"byte":         true,
	"case":         true,
	"catch":        true,
	"char":         true,
	"class":        true,
	"const":        true,
	"continue":     true,
	"default":      true,
	"do":           true,
	"double":       true,
	"else":         true,
	"enum":         true,
	"extends":      true,
	"final":        true,
	"finally":      true,
	"float":        true,
	"for":          true,
	"goto":         true,
	"if":           true,
	"implements":   true,
	"import":       true,
	"instanceof":   true,
	"int":          true,
	"interface":    true,
	"long":         true,
	"native":       true,
	"new":          true,
	"package":      true,
	"private":      true,
	"protected":    true,
	"public":       true,
	"return":       true,
	"short":        true,
	"static":       true,
	"strictfp":     true,
	"super":        true,
	"switch":       true,
	"synchronized": true,
	"this":         true,
	"throw":        true,
	"throws":       true,
	"transient":    true,
	"try":          true,
	"void":         true,
	"volatile":     true,
	"while":        true,
}

// ClassDeclaration represents a Java class declaration
type ClassDeclaration struct {
	Name           string
	ParentClass    string
	Interfaces     []string
	Modifiers      []string
	IsInterface    bool
	IsAbstract     bool
	StartPosition  int
	EndPosition    int
}

// NewJavaLexer creates a new Java lexer
func NewJavaLexer(source string) *JavaLexer {
	lexer := &JavaLexer{
		text: source,
	}
	lexer.scanner.Init(strings.NewReader(source))
	lexer.scanner.Mode = scanner.ScanIdents | scanner.ScanFloats | scanner.ScanChars | scanner.ScanStrings | scanner.ScanComments
	return lexer
}

// NextToken returns the next token from the source
func (l *JavaLexer) NextToken() (JavaToken, string) {
	tok := l.scanner.Scan()
	text := l.scanner.TokenText()
	
	switch tok {
	case scanner.EOF:
		return TOKEN_EOF, ""
	case scanner.Ident:
		if JavaKeywords[text] {
			return TOKEN_KEYWORD, text
		}
		return TOKEN_IDENTIFIER, text
	case scanner.String, scanner.Char, scanner.Int, scanner.Float:
		return TOKEN_LITERAL, text
	case scanner.Comment:
		return TOKEN_COMMENT, text
	default:
		if unicode.IsSpace(rune(tok)) {
			return TOKEN_WHITESPACE, text
		}
		if strings.ContainsRune("{}()[];,.<>=!&|+-*/", rune(tok)) {
			return TOKEN_OPERATOR, text
		}
		return TOKEN_DELIMITER, text
	}
}

// JavaASTParser parses Java source code to extract class declarations
type JavaASTParser struct {
	lexer       *JavaLexer
	currentTok  JavaToken
	currentText string
}

// NewJavaASTParser creates a new Java AST parser
func NewJavaASTParser(source string) *JavaASTParser {
	parser := &JavaASTParser{
		lexer: NewJavaLexer(source),
	}
	parser.nextToken() // Initialize with first token
	return parser
}

// nextToken advances to the next token
func (p *JavaASTParser) nextToken() {
	p.currentTok, p.currentText = p.lexer.NextToken()
	// Skip whitespace and comments
	for p.currentTok == TOKEN_WHITESPACE || p.currentTok == TOKEN_COMMENT {
		p.currentTok, p.currentText = p.lexer.NextToken()
	}
}

// expect checks if current token matches expected token and advances
func (p *JavaASTParser) expect(expectedToken JavaToken, expectedText string) error {
	if p.currentTok != expectedToken || (expectedText != "" && p.currentText != expectedText) {
		return fmt.Errorf("expected %v '%s', got %v '%s'", expectedToken, expectedText, p.currentTok, p.currentText)
	}
	p.nextToken()
	return nil
}

// parseIdentifier parses a Java identifier (possibly qualified)
func (p *JavaASTParser) parseIdentifier() (string, error) {
	if p.currentTok != TOKEN_IDENTIFIER {
		return "", fmt.Errorf("expected identifier, got %v '%s'", p.currentTok, p.currentText)
	}
	
	identifier := p.currentText
	p.nextToken()
	
	// Handle qualified names (e.g., java.util.List)
	for p.currentTok == TOKEN_OPERATOR && p.currentText == "." {
		p.nextToken() // consume '.'
		if p.currentTok != TOKEN_IDENTIFIER {
			return "", fmt.Errorf("expected identifier after '.', got %v '%s'", p.currentTok, p.currentText)
		}
		identifier += "." + p.currentText
		p.nextToken()
	}
	
	return identifier, nil
}

// parseTypeParameters skips generic type parameters
func (p *JavaASTParser) parseTypeParameters() error {
	if p.currentTok == TOKEN_OPERATOR && p.currentText == "<" {
		p.nextToken() // consume '<'
		depth := 1
		for depth > 0 && p.currentTok != TOKEN_EOF {
			if p.currentTok == TOKEN_OPERATOR {
				if p.currentText == "<" {
					depth++
				} else if p.currentText == ">" {
					depth--
				}
			}
			p.nextToken()
		}
	}
	return nil
}

// parseInterfaceList parses a comma-separated list of interfaces
func (p *JavaASTParser) parseInterfaceList() ([]string, error) {
	var interfaces []string
	
	for {
		interfaceName, err := p.parseIdentifier()
		if err != nil {
			return nil, err
		}
		
		// Skip type parameters if present
		err = p.parseTypeParameters()
		if err != nil {
			return nil, err
		}
		
		interfaces = append(interfaces, interfaceName)
		
		// Check for comma to continue
		if p.currentTok == TOKEN_OPERATOR && p.currentText == "," {
			p.nextToken() // consume ','
			continue
		}
		break
	}
	
	return interfaces, nil
}

// parseClassDeclaration parses a Java class or interface declaration
func (p *JavaASTParser) parseClassDeclaration() (*ClassDeclaration, error) {
	decl := &ClassDeclaration{}
	
	// Parse modifiers
	for p.currentTok == TOKEN_KEYWORD {
		switch p.currentText {
		case "public", "private", "protected", "static", "final", "abstract":
			decl.Modifiers = append(decl.Modifiers, p.currentText)
			if p.currentText == "abstract" {
				decl.IsAbstract = true
			}
			p.nextToken()
		case "class", "interface":
			if p.currentText == "interface" {
				decl.IsInterface = true
			}
			p.nextToken()
			goto parseClassName
		default:
			return nil, fmt.Errorf("unexpected keyword: %s", p.currentText)
		}
	}
	
parseClassName:
	// Parse class/interface name
	className, err := p.parseIdentifier()
	if err != nil {
		return nil, fmt.Errorf("failed to parse class name: %v", err)
	}
	decl.Name = className
	
	// Skip type parameters if present
	err = p.parseTypeParameters()
	if err != nil {
		return nil, err
	}
	
	// Parse extends clause
	if p.currentTok == TOKEN_KEYWORD && p.currentText == "extends" {
		p.nextToken() // consume 'extends'
		parentClass, err := p.parseIdentifier()
		if err != nil {
			return nil, fmt.Errorf("failed to parse parent class: %v", err)
		}
		decl.ParentClass = parentClass
		
		// Skip type parameters if present
		err = p.parseTypeParameters()
		if err != nil {
			return nil, err
		}
	}
	
	// Parse implements clause
	if p.currentTok == TOKEN_KEYWORD && p.currentText == "implements" {
		p.nextToken() // consume 'implements'
		interfaces, err := p.parseInterfaceList()
		if err != nil {
			return nil, fmt.Errorf("failed to parse interfaces: %v", err)
		}
		decl.Interfaces = interfaces
	}
	
	return decl, nil
}

// ParseClassDeclarations extracts all class and interface declarations from Java source
func (p *JavaASTParser) ParseClassDeclarations() ([]*ClassDeclaration, error) {
	var declarations []*ClassDeclaration
	
	for p.currentTok != TOKEN_EOF {
		// Look for class or interface keywords
		if p.currentTok == TOKEN_KEYWORD && (p.currentText == "class" || p.currentText == "interface") {
			// Backtrack to capture modifiers
			// For simplicity, we'll parse from current position
			decl, err := p.parseClassDeclaration()
			if err != nil {
				return nil, err
			}
			declarations = append(declarations, decl)
			
			// Skip to next potential declaration
			for p.currentTok != TOKEN_EOF && !(p.currentTok == TOKEN_KEYWORD && (p.currentText == "class" || p.currentText == "interface")) {
				p.nextToken()
			}
		} else {
			// Look for modifiers that might precede class/interface
			if p.currentTok == TOKEN_KEYWORD && 
				(p.currentText == "public" || p.currentText == "private" || p.currentText == "protected" || 
				 p.currentText == "static" || p.currentText == "final" || p.currentText == "abstract") {
				// This might be a class declaration with modifiers
				decl, err := p.parseClassDeclaration()
				if err != nil {
					// If parsing fails, just skip this token
					p.nextToken()
					continue
				}
				declarations = append(declarations, decl)
				
				// Skip to next potential declaration
				for p.currentTok != TOKEN_EOF && !(p.currentTok == TOKEN_KEYWORD && (p.currentText == "class" || p.currentText == "interface")) {
					p.nextToken()
				}
			} else {
				p.nextToken()
			}
		}
	}
	
	return declarations, nil
}

// ParseJavaInheritance extracts inheritance information from Java source code using AST parsing
func ParseJavaInheritance(source string) (map[string]string, map[string][]string, error) {
	parser := NewJavaASTParser(source)
	declarations, err := parser.ParseClassDeclarations()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse class declarations: %v", err)
	}
	
	inheritanceTree := make(map[string]string)
	interfaceImpls := make(map[string][]string)
	
	for _, decl := range declarations {
		if decl.ParentClass != "" {
			inheritanceTree[decl.Name] = decl.ParentClass
		}
		if len(decl.Interfaces) > 0 {
			interfaceImpls[decl.Name] = decl.Interfaces
		}
	}
	
	return inheritanceTree, interfaceImpls, nil
}