package context

// BusinessActionContext store the tcc branch transaction context
type BusinessActionContext struct {
	*RootContext
	XID                    string
	BranchID               int64
	TCCAsyncPhaseTwoEnable bool
	ActionName             string
	ActionContext          map[string]interface{}
}
