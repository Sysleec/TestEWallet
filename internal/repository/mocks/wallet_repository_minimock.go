package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i github.com/Sysleec/TestEWallet/internal/repository.WalletRepository -o wallet_repository_minimock.go -n WalletRepositoryMock

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/Sysleec/TestEWallet/internal/model"
	"github.com/gojuno/minimock/v3"
)

// WalletRepositoryMock implements repository.WalletRepository
type WalletRepositoryMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcCreateWallet          func(ctx context.Context) (wp1 *model.Wallet, err error)
	inspectFuncCreateWallet   func(ctx context.Context)
	afterCreateWalletCounter  uint64
	beforeCreateWalletCounter uint64
	CreateWalletMock          mWalletRepositoryMockCreateWallet

	funcGetHistory          func(ctx context.Context, id string) (ta1 []model.Transfer, err error)
	inspectFuncGetHistory   func(ctx context.Context, id string)
	afterGetHistoryCounter  uint64
	beforeGetHistoryCounter uint64
	GetHistoryMock          mWalletRepositoryMockGetHistory

	funcGetWallet          func(ctx context.Context, id string) (wp1 *model.Wallet, err error)
	inspectFuncGetWallet   func(ctx context.Context, id string)
	afterGetWalletCounter  uint64
	beforeGetWalletCounter uint64
	GetWalletMock          mWalletRepositoryMockGetWallet

	funcSendMoney          func(ctx context.Context, trans *model.Transfer) (err error)
	inspectFuncSendMoney   func(ctx context.Context, trans *model.Transfer)
	afterSendMoneyCounter  uint64
	beforeSendMoneyCounter uint64
	SendMoneyMock          mWalletRepositoryMockSendMoney
}

// NewWalletRepositoryMock returns a mock for repository.WalletRepository
func NewWalletRepositoryMock(t minimock.Tester) *WalletRepositoryMock {
	m := &WalletRepositoryMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CreateWalletMock = mWalletRepositoryMockCreateWallet{mock: m}
	m.CreateWalletMock.callArgs = []*WalletRepositoryMockCreateWalletParams{}

	m.GetHistoryMock = mWalletRepositoryMockGetHistory{mock: m}
	m.GetHistoryMock.callArgs = []*WalletRepositoryMockGetHistoryParams{}

	m.GetWalletMock = mWalletRepositoryMockGetWallet{mock: m}
	m.GetWalletMock.callArgs = []*WalletRepositoryMockGetWalletParams{}

	m.SendMoneyMock = mWalletRepositoryMockSendMoney{mock: m}
	m.SendMoneyMock.callArgs = []*WalletRepositoryMockSendMoneyParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mWalletRepositoryMockCreateWallet struct {
	mock               *WalletRepositoryMock
	defaultExpectation *WalletRepositoryMockCreateWalletExpectation
	expectations       []*WalletRepositoryMockCreateWalletExpectation

	callArgs []*WalletRepositoryMockCreateWalletParams
	mutex    sync.RWMutex
}

// WalletRepositoryMockCreateWalletExpectation specifies expectation struct of the WalletRepository.CreateWallet
type WalletRepositoryMockCreateWalletExpectation struct {
	mock    *WalletRepositoryMock
	params  *WalletRepositoryMockCreateWalletParams
	results *WalletRepositoryMockCreateWalletResults
	Counter uint64
}

// WalletRepositoryMockCreateWalletParams contains parameters of the WalletRepository.CreateWallet
type WalletRepositoryMockCreateWalletParams struct {
	ctx context.Context
}

// WalletRepositoryMockCreateWalletResults contains results of the WalletRepository.CreateWallet
type WalletRepositoryMockCreateWalletResults struct {
	wp1 *model.Wallet
	err error
}

// Expect sets up expected params for WalletRepository.CreateWallet
func (mmCreateWallet *mWalletRepositoryMockCreateWallet) Expect(ctx context.Context) *mWalletRepositoryMockCreateWallet {
	if mmCreateWallet.mock.funcCreateWallet != nil {
		mmCreateWallet.mock.t.Fatalf("WalletRepositoryMock.CreateWallet mock is already set by Set")
	}

	if mmCreateWallet.defaultExpectation == nil {
		mmCreateWallet.defaultExpectation = &WalletRepositoryMockCreateWalletExpectation{}
	}

	mmCreateWallet.defaultExpectation.params = &WalletRepositoryMockCreateWalletParams{ctx}
	for _, e := range mmCreateWallet.expectations {
		if minimock.Equal(e.params, mmCreateWallet.defaultExpectation.params) {
			mmCreateWallet.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCreateWallet.defaultExpectation.params)
		}
	}

	return mmCreateWallet
}

// Inspect accepts an inspector function that has same arguments as the WalletRepository.CreateWallet
func (mmCreateWallet *mWalletRepositoryMockCreateWallet) Inspect(f func(ctx context.Context)) *mWalletRepositoryMockCreateWallet {
	if mmCreateWallet.mock.inspectFuncCreateWallet != nil {
		mmCreateWallet.mock.t.Fatalf("Inspect function is already set for WalletRepositoryMock.CreateWallet")
	}

	mmCreateWallet.mock.inspectFuncCreateWallet = f

	return mmCreateWallet
}

// Return sets up results that will be returned by WalletRepository.CreateWallet
func (mmCreateWallet *mWalletRepositoryMockCreateWallet) Return(wp1 *model.Wallet, err error) *WalletRepositoryMock {
	if mmCreateWallet.mock.funcCreateWallet != nil {
		mmCreateWallet.mock.t.Fatalf("WalletRepositoryMock.CreateWallet mock is already set by Set")
	}

	if mmCreateWallet.defaultExpectation == nil {
		mmCreateWallet.defaultExpectation = &WalletRepositoryMockCreateWalletExpectation{mock: mmCreateWallet.mock}
	}
	mmCreateWallet.defaultExpectation.results = &WalletRepositoryMockCreateWalletResults{wp1, err}
	return mmCreateWallet.mock
}

// Set uses given function f to mock the WalletRepository.CreateWallet method
func (mmCreateWallet *mWalletRepositoryMockCreateWallet) Set(f func(ctx context.Context) (wp1 *model.Wallet, err error)) *WalletRepositoryMock {
	if mmCreateWallet.defaultExpectation != nil {
		mmCreateWallet.mock.t.Fatalf("Default expectation is already set for the WalletRepository.CreateWallet method")
	}

	if len(mmCreateWallet.expectations) > 0 {
		mmCreateWallet.mock.t.Fatalf("Some expectations are already set for the WalletRepository.CreateWallet method")
	}

	mmCreateWallet.mock.funcCreateWallet = f
	return mmCreateWallet.mock
}

// When sets expectation for the WalletRepository.CreateWallet which will trigger the result defined by the following
// Then helper
func (mmCreateWallet *mWalletRepositoryMockCreateWallet) When(ctx context.Context) *WalletRepositoryMockCreateWalletExpectation {
	if mmCreateWallet.mock.funcCreateWallet != nil {
		mmCreateWallet.mock.t.Fatalf("WalletRepositoryMock.CreateWallet mock is already set by Set")
	}

	expectation := &WalletRepositoryMockCreateWalletExpectation{
		mock:   mmCreateWallet.mock,
		params: &WalletRepositoryMockCreateWalletParams{ctx},
	}
	mmCreateWallet.expectations = append(mmCreateWallet.expectations, expectation)
	return expectation
}

// Then sets up WalletRepository.CreateWallet return parameters for the expectation previously defined by the When method
func (e *WalletRepositoryMockCreateWalletExpectation) Then(wp1 *model.Wallet, err error) *WalletRepositoryMock {
	e.results = &WalletRepositoryMockCreateWalletResults{wp1, err}
	return e.mock
}

// CreateWallet implements repository.WalletRepository
func (mmCreateWallet *WalletRepositoryMock) CreateWallet(ctx context.Context) (wp1 *model.Wallet, err error) {
	mm_atomic.AddUint64(&mmCreateWallet.beforeCreateWalletCounter, 1)
	defer mm_atomic.AddUint64(&mmCreateWallet.afterCreateWalletCounter, 1)

	if mmCreateWallet.inspectFuncCreateWallet != nil {
		mmCreateWallet.inspectFuncCreateWallet(ctx)
	}

	mm_params := WalletRepositoryMockCreateWalletParams{ctx}

	// Record call args
	mmCreateWallet.CreateWalletMock.mutex.Lock()
	mmCreateWallet.CreateWalletMock.callArgs = append(mmCreateWallet.CreateWalletMock.callArgs, &mm_params)
	mmCreateWallet.CreateWalletMock.mutex.Unlock()

	for _, e := range mmCreateWallet.CreateWalletMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.wp1, e.results.err
		}
	}

	if mmCreateWallet.CreateWalletMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCreateWallet.CreateWalletMock.defaultExpectation.Counter, 1)
		mm_want := mmCreateWallet.CreateWalletMock.defaultExpectation.params
		mm_got := WalletRepositoryMockCreateWalletParams{ctx}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCreateWallet.t.Errorf("WalletRepositoryMock.CreateWallet got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCreateWallet.CreateWalletMock.defaultExpectation.results
		if mm_results == nil {
			mmCreateWallet.t.Fatal("No results are set for the WalletRepositoryMock.CreateWallet")
		}
		return (*mm_results).wp1, (*mm_results).err
	}
	if mmCreateWallet.funcCreateWallet != nil {
		return mmCreateWallet.funcCreateWallet(ctx)
	}
	mmCreateWallet.t.Fatalf("Unexpected call to WalletRepositoryMock.CreateWallet. %v", ctx)
	return
}

// CreateWalletAfterCounter returns a count of finished WalletRepositoryMock.CreateWallet invocations
func (mmCreateWallet *WalletRepositoryMock) CreateWalletAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreateWallet.afterCreateWalletCounter)
}

// CreateWalletBeforeCounter returns a count of WalletRepositoryMock.CreateWallet invocations
func (mmCreateWallet *WalletRepositoryMock) CreateWalletBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreateWallet.beforeCreateWalletCounter)
}

// Calls returns a list of arguments used in each call to WalletRepositoryMock.CreateWallet.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCreateWallet *mWalletRepositoryMockCreateWallet) Calls() []*WalletRepositoryMockCreateWalletParams {
	mmCreateWallet.mutex.RLock()

	argCopy := make([]*WalletRepositoryMockCreateWalletParams, len(mmCreateWallet.callArgs))
	copy(argCopy, mmCreateWallet.callArgs)

	mmCreateWallet.mutex.RUnlock()

	return argCopy
}

// MinimockCreateWalletDone returns true if the count of the CreateWallet invocations corresponds
// the number of defined expectations
func (m *WalletRepositoryMock) MinimockCreateWalletDone() bool {
	for _, e := range m.CreateWalletMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateWalletMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateWalletCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreateWallet != nil && mm_atomic.LoadUint64(&m.afterCreateWalletCounter) < 1 {
		return false
	}
	return true
}

// MinimockCreateWalletInspect logs each unmet expectation
func (m *WalletRepositoryMock) MinimockCreateWalletInspect() {
	for _, e := range m.CreateWalletMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to WalletRepositoryMock.CreateWallet with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateWalletMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateWalletCounter) < 1 {
		if m.CreateWalletMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to WalletRepositoryMock.CreateWallet")
		} else {
			m.t.Errorf("Expected call to WalletRepositoryMock.CreateWallet with params: %#v", *m.CreateWalletMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreateWallet != nil && mm_atomic.LoadUint64(&m.afterCreateWalletCounter) < 1 {
		m.t.Error("Expected call to WalletRepositoryMock.CreateWallet")
	}
}

type mWalletRepositoryMockGetHistory struct {
	mock               *WalletRepositoryMock
	defaultExpectation *WalletRepositoryMockGetHistoryExpectation
	expectations       []*WalletRepositoryMockGetHistoryExpectation

	callArgs []*WalletRepositoryMockGetHistoryParams
	mutex    sync.RWMutex
}

// WalletRepositoryMockGetHistoryExpectation specifies expectation struct of the WalletRepository.GetHistory
type WalletRepositoryMockGetHistoryExpectation struct {
	mock    *WalletRepositoryMock
	params  *WalletRepositoryMockGetHistoryParams
	results *WalletRepositoryMockGetHistoryResults
	Counter uint64
}

// WalletRepositoryMockGetHistoryParams contains parameters of the WalletRepository.GetHistory
type WalletRepositoryMockGetHistoryParams struct {
	ctx context.Context
	id  string
}

// WalletRepositoryMockGetHistoryResults contains results of the WalletRepository.GetHistory
type WalletRepositoryMockGetHistoryResults struct {
	ta1 []model.Transfer
	err error
}

// Expect sets up expected params for WalletRepository.GetHistory
func (mmGetHistory *mWalletRepositoryMockGetHistory) Expect(ctx context.Context, id string) *mWalletRepositoryMockGetHistory {
	if mmGetHistory.mock.funcGetHistory != nil {
		mmGetHistory.mock.t.Fatalf("WalletRepositoryMock.GetHistory mock is already set by Set")
	}

	if mmGetHistory.defaultExpectation == nil {
		mmGetHistory.defaultExpectation = &WalletRepositoryMockGetHistoryExpectation{}
	}

	mmGetHistory.defaultExpectation.params = &WalletRepositoryMockGetHistoryParams{ctx, id}
	for _, e := range mmGetHistory.expectations {
		if minimock.Equal(e.params, mmGetHistory.defaultExpectation.params) {
			mmGetHistory.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetHistory.defaultExpectation.params)
		}
	}

	return mmGetHistory
}

// Inspect accepts an inspector function that has same arguments as the WalletRepository.GetHistory
func (mmGetHistory *mWalletRepositoryMockGetHistory) Inspect(f func(ctx context.Context, id string)) *mWalletRepositoryMockGetHistory {
	if mmGetHistory.mock.inspectFuncGetHistory != nil {
		mmGetHistory.mock.t.Fatalf("Inspect function is already set for WalletRepositoryMock.GetHistory")
	}

	mmGetHistory.mock.inspectFuncGetHistory = f

	return mmGetHistory
}

// Return sets up results that will be returned by WalletRepository.GetHistory
func (mmGetHistory *mWalletRepositoryMockGetHistory) Return(ta1 []model.Transfer, err error) *WalletRepositoryMock {
	if mmGetHistory.mock.funcGetHistory != nil {
		mmGetHistory.mock.t.Fatalf("WalletRepositoryMock.GetHistory mock is already set by Set")
	}

	if mmGetHistory.defaultExpectation == nil {
		mmGetHistory.defaultExpectation = &WalletRepositoryMockGetHistoryExpectation{mock: mmGetHistory.mock}
	}
	mmGetHistory.defaultExpectation.results = &WalletRepositoryMockGetHistoryResults{ta1, err}
	return mmGetHistory.mock
}

// Set uses given function f to mock the WalletRepository.GetHistory method
func (mmGetHistory *mWalletRepositoryMockGetHistory) Set(f func(ctx context.Context, id string) (ta1 []model.Transfer, err error)) *WalletRepositoryMock {
	if mmGetHistory.defaultExpectation != nil {
		mmGetHistory.mock.t.Fatalf("Default expectation is already set for the WalletRepository.GetHistory method")
	}

	if len(mmGetHistory.expectations) > 0 {
		mmGetHistory.mock.t.Fatalf("Some expectations are already set for the WalletRepository.GetHistory method")
	}

	mmGetHistory.mock.funcGetHistory = f
	return mmGetHistory.mock
}

// When sets expectation for the WalletRepository.GetHistory which will trigger the result defined by the following
// Then helper
func (mmGetHistory *mWalletRepositoryMockGetHistory) When(ctx context.Context, id string) *WalletRepositoryMockGetHistoryExpectation {
	if mmGetHistory.mock.funcGetHistory != nil {
		mmGetHistory.mock.t.Fatalf("WalletRepositoryMock.GetHistory mock is already set by Set")
	}

	expectation := &WalletRepositoryMockGetHistoryExpectation{
		mock:   mmGetHistory.mock,
		params: &WalletRepositoryMockGetHistoryParams{ctx, id},
	}
	mmGetHistory.expectations = append(mmGetHistory.expectations, expectation)
	return expectation
}

// Then sets up WalletRepository.GetHistory return parameters for the expectation previously defined by the When method
func (e *WalletRepositoryMockGetHistoryExpectation) Then(ta1 []model.Transfer, err error) *WalletRepositoryMock {
	e.results = &WalletRepositoryMockGetHistoryResults{ta1, err}
	return e.mock
}

// GetHistory implements repository.WalletRepository
func (mmGetHistory *WalletRepositoryMock) GetHistory(ctx context.Context, id string) (ta1 []model.Transfer, err error) {
	mm_atomic.AddUint64(&mmGetHistory.beforeGetHistoryCounter, 1)
	defer mm_atomic.AddUint64(&mmGetHistory.afterGetHistoryCounter, 1)

	if mmGetHistory.inspectFuncGetHistory != nil {
		mmGetHistory.inspectFuncGetHistory(ctx, id)
	}

	mm_params := WalletRepositoryMockGetHistoryParams{ctx, id}

	// Record call args
	mmGetHistory.GetHistoryMock.mutex.Lock()
	mmGetHistory.GetHistoryMock.callArgs = append(mmGetHistory.GetHistoryMock.callArgs, &mm_params)
	mmGetHistory.GetHistoryMock.mutex.Unlock()

	for _, e := range mmGetHistory.GetHistoryMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.ta1, e.results.err
		}
	}

	if mmGetHistory.GetHistoryMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetHistory.GetHistoryMock.defaultExpectation.Counter, 1)
		mm_want := mmGetHistory.GetHistoryMock.defaultExpectation.params
		mm_got := WalletRepositoryMockGetHistoryParams{ctx, id}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetHistory.t.Errorf("WalletRepositoryMock.GetHistory got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetHistory.GetHistoryMock.defaultExpectation.results
		if mm_results == nil {
			mmGetHistory.t.Fatal("No results are set for the WalletRepositoryMock.GetHistory")
		}
		return (*mm_results).ta1, (*mm_results).err
	}
	if mmGetHistory.funcGetHistory != nil {
		return mmGetHistory.funcGetHistory(ctx, id)
	}
	mmGetHistory.t.Fatalf("Unexpected call to WalletRepositoryMock.GetHistory. %v %v", ctx, id)
	return
}

// GetHistoryAfterCounter returns a count of finished WalletRepositoryMock.GetHistory invocations
func (mmGetHistory *WalletRepositoryMock) GetHistoryAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetHistory.afterGetHistoryCounter)
}

// GetHistoryBeforeCounter returns a count of WalletRepositoryMock.GetHistory invocations
func (mmGetHistory *WalletRepositoryMock) GetHistoryBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetHistory.beforeGetHistoryCounter)
}

// Calls returns a list of arguments used in each call to WalletRepositoryMock.GetHistory.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetHistory *mWalletRepositoryMockGetHistory) Calls() []*WalletRepositoryMockGetHistoryParams {
	mmGetHistory.mutex.RLock()

	argCopy := make([]*WalletRepositoryMockGetHistoryParams, len(mmGetHistory.callArgs))
	copy(argCopy, mmGetHistory.callArgs)

	mmGetHistory.mutex.RUnlock()

	return argCopy
}

// MinimockGetHistoryDone returns true if the count of the GetHistory invocations corresponds
// the number of defined expectations
func (m *WalletRepositoryMock) MinimockGetHistoryDone() bool {
	for _, e := range m.GetHistoryMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetHistoryMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetHistoryCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetHistory != nil && mm_atomic.LoadUint64(&m.afterGetHistoryCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetHistoryInspect logs each unmet expectation
func (m *WalletRepositoryMock) MinimockGetHistoryInspect() {
	for _, e := range m.GetHistoryMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to WalletRepositoryMock.GetHistory with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetHistoryMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetHistoryCounter) < 1 {
		if m.GetHistoryMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to WalletRepositoryMock.GetHistory")
		} else {
			m.t.Errorf("Expected call to WalletRepositoryMock.GetHistory with params: %#v", *m.GetHistoryMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetHistory != nil && mm_atomic.LoadUint64(&m.afterGetHistoryCounter) < 1 {
		m.t.Error("Expected call to WalletRepositoryMock.GetHistory")
	}
}

type mWalletRepositoryMockGetWallet struct {
	mock               *WalletRepositoryMock
	defaultExpectation *WalletRepositoryMockGetWalletExpectation
	expectations       []*WalletRepositoryMockGetWalletExpectation

	callArgs []*WalletRepositoryMockGetWalletParams
	mutex    sync.RWMutex
}

// WalletRepositoryMockGetWalletExpectation specifies expectation struct of the WalletRepository.GetWallet
type WalletRepositoryMockGetWalletExpectation struct {
	mock    *WalletRepositoryMock
	params  *WalletRepositoryMockGetWalletParams
	results *WalletRepositoryMockGetWalletResults
	Counter uint64
}

// WalletRepositoryMockGetWalletParams contains parameters of the WalletRepository.GetWallet
type WalletRepositoryMockGetWalletParams struct {
	ctx context.Context
	id  string
}

// WalletRepositoryMockGetWalletResults contains results of the WalletRepository.GetWallet
type WalletRepositoryMockGetWalletResults struct {
	wp1 *model.Wallet
	err error
}

// Expect sets up expected params for WalletRepository.GetWallet
func (mmGetWallet *mWalletRepositoryMockGetWallet) Expect(ctx context.Context, id string) *mWalletRepositoryMockGetWallet {
	if mmGetWallet.mock.funcGetWallet != nil {
		mmGetWallet.mock.t.Fatalf("WalletRepositoryMock.GetWallet mock is already set by Set")
	}

	if mmGetWallet.defaultExpectation == nil {
		mmGetWallet.defaultExpectation = &WalletRepositoryMockGetWalletExpectation{}
	}

	mmGetWallet.defaultExpectation.params = &WalletRepositoryMockGetWalletParams{ctx, id}
	for _, e := range mmGetWallet.expectations {
		if minimock.Equal(e.params, mmGetWallet.defaultExpectation.params) {
			mmGetWallet.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetWallet.defaultExpectation.params)
		}
	}

	return mmGetWallet
}

// Inspect accepts an inspector function that has same arguments as the WalletRepository.GetWallet
func (mmGetWallet *mWalletRepositoryMockGetWallet) Inspect(f func(ctx context.Context, id string)) *mWalletRepositoryMockGetWallet {
	if mmGetWallet.mock.inspectFuncGetWallet != nil {
		mmGetWallet.mock.t.Fatalf("Inspect function is already set for WalletRepositoryMock.GetWallet")
	}

	mmGetWallet.mock.inspectFuncGetWallet = f

	return mmGetWallet
}

// Return sets up results that will be returned by WalletRepository.GetWallet
func (mmGetWallet *mWalletRepositoryMockGetWallet) Return(wp1 *model.Wallet, err error) *WalletRepositoryMock {
	if mmGetWallet.mock.funcGetWallet != nil {
		mmGetWallet.mock.t.Fatalf("WalletRepositoryMock.GetWallet mock is already set by Set")
	}

	if mmGetWallet.defaultExpectation == nil {
		mmGetWallet.defaultExpectation = &WalletRepositoryMockGetWalletExpectation{mock: mmGetWallet.mock}
	}
	mmGetWallet.defaultExpectation.results = &WalletRepositoryMockGetWalletResults{wp1, err}
	return mmGetWallet.mock
}

// Set uses given function f to mock the WalletRepository.GetWallet method
func (mmGetWallet *mWalletRepositoryMockGetWallet) Set(f func(ctx context.Context, id string) (wp1 *model.Wallet, err error)) *WalletRepositoryMock {
	if mmGetWallet.defaultExpectation != nil {
		mmGetWallet.mock.t.Fatalf("Default expectation is already set for the WalletRepository.GetWallet method")
	}

	if len(mmGetWallet.expectations) > 0 {
		mmGetWallet.mock.t.Fatalf("Some expectations are already set for the WalletRepository.GetWallet method")
	}

	mmGetWallet.mock.funcGetWallet = f
	return mmGetWallet.mock
}

// When sets expectation for the WalletRepository.GetWallet which will trigger the result defined by the following
// Then helper
func (mmGetWallet *mWalletRepositoryMockGetWallet) When(ctx context.Context, id string) *WalletRepositoryMockGetWalletExpectation {
	if mmGetWallet.mock.funcGetWallet != nil {
		mmGetWallet.mock.t.Fatalf("WalletRepositoryMock.GetWallet mock is already set by Set")
	}

	expectation := &WalletRepositoryMockGetWalletExpectation{
		mock:   mmGetWallet.mock,
		params: &WalletRepositoryMockGetWalletParams{ctx, id},
	}
	mmGetWallet.expectations = append(mmGetWallet.expectations, expectation)
	return expectation
}

// Then sets up WalletRepository.GetWallet return parameters for the expectation previously defined by the When method
func (e *WalletRepositoryMockGetWalletExpectation) Then(wp1 *model.Wallet, err error) *WalletRepositoryMock {
	e.results = &WalletRepositoryMockGetWalletResults{wp1, err}
	return e.mock
}

// GetWallet implements repository.WalletRepository
func (mmGetWallet *WalletRepositoryMock) GetWallet(ctx context.Context, id string) (wp1 *model.Wallet, err error) {
	mm_atomic.AddUint64(&mmGetWallet.beforeGetWalletCounter, 1)
	defer mm_atomic.AddUint64(&mmGetWallet.afterGetWalletCounter, 1)

	if mmGetWallet.inspectFuncGetWallet != nil {
		mmGetWallet.inspectFuncGetWallet(ctx, id)
	}

	mm_params := WalletRepositoryMockGetWalletParams{ctx, id}

	// Record call args
	mmGetWallet.GetWalletMock.mutex.Lock()
	mmGetWallet.GetWalletMock.callArgs = append(mmGetWallet.GetWalletMock.callArgs, &mm_params)
	mmGetWallet.GetWalletMock.mutex.Unlock()

	for _, e := range mmGetWallet.GetWalletMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.wp1, e.results.err
		}
	}

	if mmGetWallet.GetWalletMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetWallet.GetWalletMock.defaultExpectation.Counter, 1)
		mm_want := mmGetWallet.GetWalletMock.defaultExpectation.params
		mm_got := WalletRepositoryMockGetWalletParams{ctx, id}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetWallet.t.Errorf("WalletRepositoryMock.GetWallet got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetWallet.GetWalletMock.defaultExpectation.results
		if mm_results == nil {
			mmGetWallet.t.Fatal("No results are set for the WalletRepositoryMock.GetWallet")
		}
		return (*mm_results).wp1, (*mm_results).err
	}
	if mmGetWallet.funcGetWallet != nil {
		return mmGetWallet.funcGetWallet(ctx, id)
	}
	mmGetWallet.t.Fatalf("Unexpected call to WalletRepositoryMock.GetWallet. %v %v", ctx, id)
	return
}

// GetWalletAfterCounter returns a count of finished WalletRepositoryMock.GetWallet invocations
func (mmGetWallet *WalletRepositoryMock) GetWalletAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetWallet.afterGetWalletCounter)
}

// GetWalletBeforeCounter returns a count of WalletRepositoryMock.GetWallet invocations
func (mmGetWallet *WalletRepositoryMock) GetWalletBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetWallet.beforeGetWalletCounter)
}

// Calls returns a list of arguments used in each call to WalletRepositoryMock.GetWallet.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetWallet *mWalletRepositoryMockGetWallet) Calls() []*WalletRepositoryMockGetWalletParams {
	mmGetWallet.mutex.RLock()

	argCopy := make([]*WalletRepositoryMockGetWalletParams, len(mmGetWallet.callArgs))
	copy(argCopy, mmGetWallet.callArgs)

	mmGetWallet.mutex.RUnlock()

	return argCopy
}

// MinimockGetWalletDone returns true if the count of the GetWallet invocations corresponds
// the number of defined expectations
func (m *WalletRepositoryMock) MinimockGetWalletDone() bool {
	for _, e := range m.GetWalletMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetWalletMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetWalletCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetWallet != nil && mm_atomic.LoadUint64(&m.afterGetWalletCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetWalletInspect logs each unmet expectation
func (m *WalletRepositoryMock) MinimockGetWalletInspect() {
	for _, e := range m.GetWalletMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to WalletRepositoryMock.GetWallet with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetWalletMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetWalletCounter) < 1 {
		if m.GetWalletMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to WalletRepositoryMock.GetWallet")
		} else {
			m.t.Errorf("Expected call to WalletRepositoryMock.GetWallet with params: %#v", *m.GetWalletMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetWallet != nil && mm_atomic.LoadUint64(&m.afterGetWalletCounter) < 1 {
		m.t.Error("Expected call to WalletRepositoryMock.GetWallet")
	}
}

type mWalletRepositoryMockSendMoney struct {
	mock               *WalletRepositoryMock
	defaultExpectation *WalletRepositoryMockSendMoneyExpectation
	expectations       []*WalletRepositoryMockSendMoneyExpectation

	callArgs []*WalletRepositoryMockSendMoneyParams
	mutex    sync.RWMutex
}

// WalletRepositoryMockSendMoneyExpectation specifies expectation struct of the WalletRepository.SendMoney
type WalletRepositoryMockSendMoneyExpectation struct {
	mock    *WalletRepositoryMock
	params  *WalletRepositoryMockSendMoneyParams
	results *WalletRepositoryMockSendMoneyResults
	Counter uint64
}

// WalletRepositoryMockSendMoneyParams contains parameters of the WalletRepository.SendMoney
type WalletRepositoryMockSendMoneyParams struct {
	ctx   context.Context
	trans *model.Transfer
}

// WalletRepositoryMockSendMoneyResults contains results of the WalletRepository.SendMoney
type WalletRepositoryMockSendMoneyResults struct {
	err error
}

// Expect sets up expected params for WalletRepository.SendMoney
func (mmSendMoney *mWalletRepositoryMockSendMoney) Expect(ctx context.Context, trans *model.Transfer) *mWalletRepositoryMockSendMoney {
	if mmSendMoney.mock.funcSendMoney != nil {
		mmSendMoney.mock.t.Fatalf("WalletRepositoryMock.SendMoney mock is already set by Set")
	}

	if mmSendMoney.defaultExpectation == nil {
		mmSendMoney.defaultExpectation = &WalletRepositoryMockSendMoneyExpectation{}
	}

	mmSendMoney.defaultExpectation.params = &WalletRepositoryMockSendMoneyParams{ctx, trans}
	for _, e := range mmSendMoney.expectations {
		if minimock.Equal(e.params, mmSendMoney.defaultExpectation.params) {
			mmSendMoney.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmSendMoney.defaultExpectation.params)
		}
	}

	return mmSendMoney
}

// Inspect accepts an inspector function that has same arguments as the WalletRepository.SendMoney
func (mmSendMoney *mWalletRepositoryMockSendMoney) Inspect(f func(ctx context.Context, trans *model.Transfer)) *mWalletRepositoryMockSendMoney {
	if mmSendMoney.mock.inspectFuncSendMoney != nil {
		mmSendMoney.mock.t.Fatalf("Inspect function is already set for WalletRepositoryMock.SendMoney")
	}

	mmSendMoney.mock.inspectFuncSendMoney = f

	return mmSendMoney
}

// Return sets up results that will be returned by WalletRepository.SendMoney
func (mmSendMoney *mWalletRepositoryMockSendMoney) Return(err error) *WalletRepositoryMock {
	if mmSendMoney.mock.funcSendMoney != nil {
		mmSendMoney.mock.t.Fatalf("WalletRepositoryMock.SendMoney mock is already set by Set")
	}

	if mmSendMoney.defaultExpectation == nil {
		mmSendMoney.defaultExpectation = &WalletRepositoryMockSendMoneyExpectation{mock: mmSendMoney.mock}
	}
	mmSendMoney.defaultExpectation.results = &WalletRepositoryMockSendMoneyResults{err}
	return mmSendMoney.mock
}

// Set uses given function f to mock the WalletRepository.SendMoney method
func (mmSendMoney *mWalletRepositoryMockSendMoney) Set(f func(ctx context.Context, trans *model.Transfer) (err error)) *WalletRepositoryMock {
	if mmSendMoney.defaultExpectation != nil {
		mmSendMoney.mock.t.Fatalf("Default expectation is already set for the WalletRepository.SendMoney method")
	}

	if len(mmSendMoney.expectations) > 0 {
		mmSendMoney.mock.t.Fatalf("Some expectations are already set for the WalletRepository.SendMoney method")
	}

	mmSendMoney.mock.funcSendMoney = f
	return mmSendMoney.mock
}

// When sets expectation for the WalletRepository.SendMoney which will trigger the result defined by the following
// Then helper
func (mmSendMoney *mWalletRepositoryMockSendMoney) When(ctx context.Context, trans *model.Transfer) *WalletRepositoryMockSendMoneyExpectation {
	if mmSendMoney.mock.funcSendMoney != nil {
		mmSendMoney.mock.t.Fatalf("WalletRepositoryMock.SendMoney mock is already set by Set")
	}

	expectation := &WalletRepositoryMockSendMoneyExpectation{
		mock:   mmSendMoney.mock,
		params: &WalletRepositoryMockSendMoneyParams{ctx, trans},
	}
	mmSendMoney.expectations = append(mmSendMoney.expectations, expectation)
	return expectation
}

// Then sets up WalletRepository.SendMoney return parameters for the expectation previously defined by the When method
func (e *WalletRepositoryMockSendMoneyExpectation) Then(err error) *WalletRepositoryMock {
	e.results = &WalletRepositoryMockSendMoneyResults{err}
	return e.mock
}

// SendMoney implements repository.WalletRepository
func (mmSendMoney *WalletRepositoryMock) SendMoney(ctx context.Context, trans *model.Transfer) (err error) {
	mm_atomic.AddUint64(&mmSendMoney.beforeSendMoneyCounter, 1)
	defer mm_atomic.AddUint64(&mmSendMoney.afterSendMoneyCounter, 1)

	if mmSendMoney.inspectFuncSendMoney != nil {
		mmSendMoney.inspectFuncSendMoney(ctx, trans)
	}

	mm_params := WalletRepositoryMockSendMoneyParams{ctx, trans}

	// Record call args
	mmSendMoney.SendMoneyMock.mutex.Lock()
	mmSendMoney.SendMoneyMock.callArgs = append(mmSendMoney.SendMoneyMock.callArgs, &mm_params)
	mmSendMoney.SendMoneyMock.mutex.Unlock()

	for _, e := range mmSendMoney.SendMoneyMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmSendMoney.SendMoneyMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSendMoney.SendMoneyMock.defaultExpectation.Counter, 1)
		mm_want := mmSendMoney.SendMoneyMock.defaultExpectation.params
		mm_got := WalletRepositoryMockSendMoneyParams{ctx, trans}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmSendMoney.t.Errorf("WalletRepositoryMock.SendMoney got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmSendMoney.SendMoneyMock.defaultExpectation.results
		if mm_results == nil {
			mmSendMoney.t.Fatal("No results are set for the WalletRepositoryMock.SendMoney")
		}
		return (*mm_results).err
	}
	if mmSendMoney.funcSendMoney != nil {
		return mmSendMoney.funcSendMoney(ctx, trans)
	}
	mmSendMoney.t.Fatalf("Unexpected call to WalletRepositoryMock.SendMoney. %v %v", ctx, trans)
	return
}

// SendMoneyAfterCounter returns a count of finished WalletRepositoryMock.SendMoney invocations
func (mmSendMoney *WalletRepositoryMock) SendMoneyAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSendMoney.afterSendMoneyCounter)
}

// SendMoneyBeforeCounter returns a count of WalletRepositoryMock.SendMoney invocations
func (mmSendMoney *WalletRepositoryMock) SendMoneyBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSendMoney.beforeSendMoneyCounter)
}

// Calls returns a list of arguments used in each call to WalletRepositoryMock.SendMoney.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmSendMoney *mWalletRepositoryMockSendMoney) Calls() []*WalletRepositoryMockSendMoneyParams {
	mmSendMoney.mutex.RLock()

	argCopy := make([]*WalletRepositoryMockSendMoneyParams, len(mmSendMoney.callArgs))
	copy(argCopy, mmSendMoney.callArgs)

	mmSendMoney.mutex.RUnlock()

	return argCopy
}

// MinimockSendMoneyDone returns true if the count of the SendMoney invocations corresponds
// the number of defined expectations
func (m *WalletRepositoryMock) MinimockSendMoneyDone() bool {
	for _, e := range m.SendMoneyMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SendMoneyMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSendMoneyCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSendMoney != nil && mm_atomic.LoadUint64(&m.afterSendMoneyCounter) < 1 {
		return false
	}
	return true
}

// MinimockSendMoneyInspect logs each unmet expectation
func (m *WalletRepositoryMock) MinimockSendMoneyInspect() {
	for _, e := range m.SendMoneyMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to WalletRepositoryMock.SendMoney with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SendMoneyMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSendMoneyCounter) < 1 {
		if m.SendMoneyMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to WalletRepositoryMock.SendMoney")
		} else {
			m.t.Errorf("Expected call to WalletRepositoryMock.SendMoney with params: %#v", *m.SendMoneyMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSendMoney != nil && mm_atomic.LoadUint64(&m.afterSendMoneyCounter) < 1 {
		m.t.Error("Expected call to WalletRepositoryMock.SendMoney")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *WalletRepositoryMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockCreateWalletInspect()

			m.MinimockGetHistoryInspect()

			m.MinimockGetWalletInspect()

			m.MinimockSendMoneyInspect()
			m.t.FailNow()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *WalletRepositoryMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *WalletRepositoryMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCreateWalletDone() &&
		m.MinimockGetHistoryDone() &&
		m.MinimockGetWalletDone() &&
		m.MinimockSendMoneyDone()
}