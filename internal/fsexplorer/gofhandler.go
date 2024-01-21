package fsexplorer

import (
	"bufio"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strconv"
	"strings"

	fsinfo "github.com/Assifar-Karim/cyclomatix/internal/fctinfo"
	"github.com/Assifar-Karim/cyclomatix/internal/utils"
)

type GoFileHandler struct {
	indirectionLvl int32
}

type GoFctVisitor struct {
	cfg            *utils.Graph
	fset           *token.FileSet
	lines          []string
	visitedNodes   []ast.Node
	deferredCalls  []ast.Node
	returnBlocks   []int32
	loopStack      []int32
	afterLoopStack []int32
	breakQueue     []int32
	callList       map[string]int
}

func (fh GoFileHandler) HandleFile(path string, fctTable *[]fsinfo.FctInfo) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path, nil, 0)
	if err != nil {
		log.Fatalf("[ERROR] Couldn't parse %q file: %s\n", path, err)
	}
	lines := readSourceFile(path)

	for _, decl := range file.Decls {
		switch d := decl.(type) {
		case *ast.FuncDecl:
			visitor := GoFctVisitor{
				cfg:      utils.NewCFGraph(),
				fset:     fset,
				lines:    lines,
				callList: map[string]int{},
			}
			visitor.cfg.Append("")
			ast.Walk(&visitor, d)
			// Reinitialize visited nodes for deferred calls visit
			visitor.visitedNodes = []ast.Node{}
			for i := len(visitor.deferredCalls) - 1; i >= 0; i-- {
				deferredCall := visitor.deferredCalls[i]
				ast.Walk(&visitor, deferredCall)
			}
			if visitor.cfg.NodesNames[visitor.cfg.LatestNode] == "" {
				visitor.cfg.AddStatement("End")
			} else {
				visitor.cfg.Append("End")
			}
			// Link return blocks to the end block
			for _, value := range visitor.returnBlocks {
				visitor.cfg.AdjList[value] = []int32{visitor.cfg.LatestNode}
			}
			//visitor.cfg.Print()
			*fctTable = append(*fctTable, fsinfo.NewFctInfo(
				file.Name.Name,
				d.Name.Name,
				path,
				*visitor.cfg,
				visitor.callList,
			))
		}
	}
}

func (fh GoFileHandler) ComputeComplexities(fctTable *[]fsinfo.FctInfo) {
}

func (v *GoFctVisitor) Visit(node ast.Node) ast.Visitor {
	// Check if the current node hasn't been visited before
	for _, visitedNode := range v.visitedNodes {
		if node == visitedNode {
			return v
		}
	}

	switch n := node.(type) {
	case *ast.DeclStmt, *ast.AssignStmt, *ast.IncDecStmt:
		stmt := v.getLine(n.Pos(), n.End())
		v.cfg.AddStatement(stmt)
	case *ast.DeferStmt:
		for _, args := range n.Call.Args {
			switch a := args.(type) {
			case *ast.CallExpr:
				v.visitedNodes = append(v.visitedNodes, a)
			}
		}
		v.visitedNodes = append(v.visitedNodes, n.Call)
		v.deferredCalls = append(v.deferredCalls, n.Call)
	case *ast.CallExpr:
		stmt := v.getLine(n.Pos(), n.End())
		switch s := n.Fun.(type) {
		case *ast.SelectorExpr:
			key := fmt.Sprintf("%v.%v", s.X, s.Sel)
			v.callList[key]++
		}
		for _, arg := range n.Args {
			switch a := arg.(type) {
			case *ast.CallExpr:
				ast.Walk(v, a)
			}
		}

		if v.cfg.NodesNames[v.cfg.LatestNode] == "" {
			v.cfg.AddStatement(stmt)
		} else {
			v.cfg.Append(stmt)
		}
		v.cfg.Append("")
	case *ast.IfStmt:
		if n.Init != nil {
			ast.Walk(v, n.Init)
		}
		cond := v.getLine(n.Cond.Pos(), n.Cond.End())
		stmt := "if " + cond[:len(cond)-1]
		v.cfg.AddStatement(stmt)
		condBlockIdx := v.cfg.LatestNode

		// If body block
		v.cfg.Append("")
		ast.Walk(v, n.Body)
		ifBlockEndIdx := v.cfg.LatestNode
		if v.cfg.NodesNames[ifBlockEndIdx] == "" {
			v.cfg.DeleteNode(ifBlockEndIdx)
			ifBlockEndIdx = v.cfg.LatestNode
		}

		var elseBlockEndIdx int32 = -1
		if n.Else != nil {
			// Else body block
			v.cfg.AddNode("")
			v.cfg.LinkNodes(condBlockIdx, v.cfg.LatestNode)
			ast.Walk(v, n.Else)
			elseBlockEndIdx = v.cfg.LatestNode
		}
		// After if block
		afterIfBlockIdx := elseBlockEndIdx
		if elseBlockEndIdx == -1 || v.cfg.NodesNames[elseBlockEndIdx] != "" {
			v.cfg.AddNode("If End")
			afterIfBlockIdx = v.cfg.LatestNode
		} else {
			v.cfg.AddStatement("If End")
		}
		v.cfg.Append("")
		ifBlockEndStmt := strings.TrimSuffix(v.cfg.NodesNames[ifBlockEndIdx], "\n")

		if ifBlockEndStmt != "break" && ifBlockEndStmt != "continue" {
			v.cfg.LinkNodes(ifBlockEndIdx, afterIfBlockIdx)
		}
		if elseBlockEndIdx != -1 && afterIfBlockIdx != elseBlockEndIdx {
			elseBlockEndStmt := strings.TrimSuffix(v.cfg.NodesNames[elseBlockEndIdx], "\n")
			if elseBlockEndStmt != "break" && elseBlockEndStmt != "continue" {
				v.cfg.LinkNodes(elseBlockEndIdx, afterIfBlockIdx)
			}
		} else if elseBlockEndIdx == -1 {
			v.cfg.LinkNodes(condBlockIdx, afterIfBlockIdx)
		}
	case *ast.SwitchStmt:
		stmt := "switch "
		if n.Init != nil {
			ast.Walk(v, n.Init)
			init := v.getLine(n.Init.Pos(), n.Init.End())
			stmt += init[:len(init)-1]
			v.visitedNodes = append(v.visitedNodes, n.Init)
		} else if n.Tag != nil {
			tag := v.getLine(n.Tag.Pos(), n.Tag.End())
			stmt += tag[:len(tag)-1]
			v.visitedNodes = append(v.visitedNodes, n.Tag)
		}
		v.cfg.AddStatement(stmt)
		condBlockIdx := v.cfg.LatestNode
		// Case statements handling
		caseBlockEndIdxs := []int32{}
		for _, caseStmt := range n.Body.List {
			switch c := caseStmt.(type) {
			case *ast.CaseClause:
				if c.List == nil {
					stmt = "default: "
				} else {
					stmt = "case "
					stmt += v.getLine(c.List[0].Pos(), c.List[0].End())
				}
				v.cfg.AddNode(stmt)
				v.cfg.LinkNodes(condBlockIdx, v.cfg.LatestNode)
				for _, b := range c.Body {
					ast.Walk(v, b)
				}
				if v.cfg.NodesNames[v.cfg.LatestNode] == "" {
					v.cfg.DeleteNode(v.cfg.LatestNode)
				}
				caseBlockEndIdxs = append(caseBlockEndIdxs, v.cfg.LatestNode)
			}
		}
		// After switch block
		v.cfg.AddNode("Switch End")
		for _, caseBlockEndIdx := range caseBlockEndIdxs {
			v.cfg.LinkNodes(caseBlockEndIdx, v.cfg.LatestNode)
		}
		v.cfg.Append("")

	case *ast.ForStmt, *ast.RangeStmt:
		stmt := "for "
		var body *ast.BlockStmt
		switch f := n.(type) {
		case *ast.ForStmt:
			body = f.Body
			if f.Init != nil {
				stmt += v.getLine(f.Init.Pos(), f.Init.End())
				v.visitedNodes = append(v.visitedNodes, f.Init)
			} else if f.Cond != nil {
				stmt += v.getLine(f.Cond.Pos(), f.Cond.End())
			}
			if f.Post != nil {
				v.visitedNodes = append(v.visitedNodes, f.Post)
			}
		case *ast.RangeStmt:
			body = f.Body
			key := v.getLine(f.Key.Pos(), f.Key.End())
			stmt += key
			switch fc := f.X.(type) {
			case *ast.CallExpr:
				ast.Walk(v, fc)
			}
		}

		stmt = stmt[:len(stmt)-1]
		if v.cfg.NodesNames[v.cfg.LatestNode] == "" {
			v.cfg.AddStatement(stmt)
		} else {
			v.cfg.Append(stmt + "\n")
		}
		v.loopStack = append(v.loopStack, v.cfg.LatestNode)

		// For body block
		hasBreak := false
		v.cfg.Append("")
		ast.Walk(v, body)
		for _, inst := range body.List {
			switch i := inst.(type) {
			case *ast.BranchStmt:
				if i.Tok == token.BREAK {
					hasBreak = true
					break
				}
			}
		}

		if !hasBreak {
			v.cfg.LinkNodes(v.cfg.LatestNode, v.loopStack[len(v.loopStack)-1])
		}

		// After for block
		if v.cfg.NodesNames[v.cfg.LatestNode] == "" {
			v.cfg.AddStatement("Loop End")
		} else {
			v.cfg.AddNode("Loop End")
		}
		v.cfg.LinkNodes(v.loopStack[len(v.loopStack)-1], v.cfg.LatestNode)
		v.afterLoopStack = append(v.afterLoopStack, v.cfg.LatestNode)
		v.cfg.Append("")
		for _, breakBlockIdx := range v.breakQueue {
			v.cfg.LinkNodes(breakBlockIdx, v.afterLoopStack[len(v.afterLoopStack)-1])
		}
		v.breakQueue = []int32{}

		v.loopStack = v.loopStack[:len(v.loopStack)-1]
		v.afterLoopStack = v.afterLoopStack[:len(v.afterLoopStack)-1]

	case *ast.BranchStmt:
		stmt := v.getLine(n.Pos(), n.End())
		if v.cfg.NodesNames[v.cfg.LatestNode] == "" {
			v.cfg.AddStatement(stmt)
		} else {
			v.cfg.Append(stmt + "\n")
		}

		if n.Tok == token.BREAK {
			v.breakQueue = append(v.breakQueue, v.cfg.LatestNode)
		} else if n.Tok == token.CONTINUE {
			v.cfg.LinkNodes(v.cfg.LatestNode, v.loopStack[len(v.loopStack)-1])
		}

	case *ast.ReturnStmt:
		stmt := v.getLine(n.Pos(), n.End())
		if v.cfg.NodesNames[v.cfg.LatestNode] == "" {
			v.cfg.AddStatement(stmt)
		} else {
			v.cfg.Append(stmt + "\n")
		}
		v.returnBlocks = append(v.returnBlocks, v.cfg.LatestNode)
	}
	v.visitedNodes = append(v.visitedNodes, node)
	return v
}

func (v *GoFctVisitor) getLine(start token.Pos, end token.Pos) string {
	startStr := fmt.Sprintf("%v", v.fset.PositionFor(start, false))
	endStr := fmt.Sprintf("%v", v.fset.PositionFor(end, false))
	startInfo := strings.Split(startStr, ":")
	endInfo := strings.Split(endStr, ":")

	startLn := atoi(startInfo[1])
	startIdx := atoi(startInfo[2])
	endLn := atoi(endInfo[1])
	endIdx := atoi(endInfo[2])

	output := ""

	for i := startLn - 1; i < endLn; i++ {
		if i == startLn-1 {
			output += v.lines[i][startIdx-1:]
		} else if i == endLn-1 {
			output += v.lines[i][:endIdx]
		} else {
			output += v.lines[i]
		}
	}
	return output
}

func readSourceFile(path string) []string {
	lines := []string{}
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("[ERROR] Couldn't open %q file: %s\n", path, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func atoi(input string) int {
	output, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	return output
}

func NewGoFileHandler(indirectionLvl int32) GoFileHandler {
	return GoFileHandler{
		indirectionLvl: indirectionLvl,
	}
}
