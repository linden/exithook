package exithook

import _ "unsafe"

//go:linkname runtime_addExitHook runtime.addExitHook
func runtime_addExitHook(f func(), runOnNonZeroExit bool)

func New(fn func()) {
	runtime_addExitHook(fn, true)
}
