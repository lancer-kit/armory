package metrics

import (
	"context"
	"encoding/json"
	"strings"

	"gitlab.inn4science.com/gophers/service-kit/routines"
)

const Separator = "."

type MKey string

func NewMKey(parts ...string) MKey {
	return MKey(strings.Join(parts, Separator))
}

func (key MKey) Split() []string {
	return strings.Split(string(key), Separator)
}

type SafeMetrics struct {
	Data        map[MKey]uint64 `json:"data"`
	PrettyPrint bool            `json:"-"`

	bus chan MKey
	ctx context.Context
}

func (m *SafeMetrics) New(ctx context.Context) *SafeMetrics {
	m.Data = make(map[MKey]uint64)
	m.bus = make(chan MKey, 16)
	m.ctx = ctx
	return m
}

func (m SafeMetrics) Add(name MKey) {
	m.bus <- name
}

func (m SafeMetrics) Collect() {
	for {
		select {
		case name := <-m.bus:
			m.Data[name]++
		case <-m.ctx.Done():
			close(m.bus)
			return
		}
	}

}

func (m *SafeMetrics) Init(ctx context.Context) routines.Worker {
	return m.New(ctx)
}

func (m SafeMetrics) RestartOnFail() bool {
	return true
}

func (m SafeMetrics) Run() {
	m.Collect()
}

func (m *SafeMetrics) MarshalJSON() ([]byte, error) {
	if !m.PrettyPrint {
		return json.Marshal(m.Data)
	}
	res := make(map[string]map[string]interface{})
	nodes := m.parseNodes()
	for key, n := range nodes {
		res[key] = n.toJSON()
	}

	return json.Marshal(res)
}

func (m SafeMetrics) parseNodes() map[string]node {
	result := make(map[string]node)
	for mKey, count := range m.Data {
		mKeyParts := mKey.Split()
		topName := mKeyParts[0]
		t := result[topName]
		t.name = topName
		result[topName] = *buildMetricsTree(&t, mKeyParts, count)
	}
	return result
}

type node struct {
	name     string
	level    int
	value    *uint64
	children map[string]*node
}

func (mn *node) toJSON() map[string]interface{} {
	res := make(map[string]interface{})
	if mn.value != nil {
		res["value"] = *mn.value
	}

	for _, value := range mn.children {
		res[value.name] = value.toJSON()
	}

	return res
}

func buildMetricsTree(parent *node, mKeyParts []string, value uint64) *node {
	parent.name = mKeyParts[parent.level]
	if parent.level+1 > len(mKeyParts) {
		parent.value = &value
		return nil
	}

	if parent.level+1 == len(mKeyParts) {
		parent.value = &value
		return parent
	}

	child := &node{
		name:  mKeyParts[parent.level+1],
		level: parent.level + 1,
	}

	child = buildMetricsTree(child, mKeyParts, value)
	if child == nil {
		return parent
	}

	if parent.children == nil {
		parent.children = map[string]*node{}
	}

	exChild, ok := parent.children[child.name]
	if ok {
		if exChild.children == nil {
			exChild.children = map[string]*node{}
		}
		for key, value := range child.children {
			exChild.children[key] = value
		}
		child = exChild
	}

	parent.children[child.name] = child
	return parent
}
