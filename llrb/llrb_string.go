package llrb

import (
	"bytes"
	"fmt"
	"io"
)

func (t *LLRB) String() string {
	buf := bytes.NewBuffer(nil)
	fprintNodes(buf, t.root, edges{}, edgeRoot)
	return buf.String()
}

func fprintNodes(wr io.Writer, n *Node, pre edges, edge edge) {
	if n == nil {
		return
	}

	{
		es := edgeSpace
		if edge == edgeRight {
			es = edgeLink
		}
		fprintNodes(wr, n.Left, append(pre, es), edgeLeft)
	}

	fmt.Fprintf(wr, "%s %v\n", append(pre, edge), n.Item)

	{
		es := edgeSpace
		if edge == edgeLeft {
			es = edgeLink
		}
		fprintNodes(wr, n.Right, append(pre, es), edgeRight)
	}
}

type edges []edge

func (e edges) String() string {
	buf := make([]rune, 0, len(e))
	for _, v := range e {
		buf = append(buf, []rune(v.String())...)
	}
	return string(buf)
}

type edge uint

const (
	_ = edge(iota)
	edgeSpace
	edgeLink
	edgeRoot
	edgeLeft
	edgeRight
)

var edgeMap = map[edge]string{
	edgeSpace: `  `,
	edgeLink:  ` │`,
	edgeRoot:  `───`,
	edgeLeft:  ` ┌─`,
	edgeRight: ` └─`,
}

func (e edge) String() string {
	v, ok := edgeMap[e]
	if ok {
		return v
	}
	return ""
}
