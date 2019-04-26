// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/RTradeLtd/database/v2/models"
)

type FakeUserManager struct {
	ChangePasswordStub        func(string, string, string) (bool, error)
	changePasswordMutex       sync.RWMutex
	changePasswordArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	changePasswordReturns struct {
		result1 bool
		result2 error
	}
	changePasswordReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	FindByEmailStub        func(string) (*models.User, error)
	findByEmailMutex       sync.RWMutex
	findByEmailArgsForCall []struct {
		arg1 string
	}
	findByEmailReturns struct {
		result1 *models.User
		result2 error
	}
	findByEmailReturnsOnCall map[int]struct {
		result1 *models.User
		result2 error
	}
	FindByUserNameStub        func(string) (*models.User, error)
	findByUserNameMutex       sync.RWMutex
	findByUserNameArgsForCall []struct {
		arg1 string
	}
	findByUserNameReturns struct {
		result1 *models.User
		result2 error
	}
	findByUserNameReturnsOnCall map[int]struct {
		result1 *models.User
		result2 error
	}
	GenerateEmailVerificationTokenStub        func(string) (*models.User, error)
	generateEmailVerificationTokenMutex       sync.RWMutex
	generateEmailVerificationTokenArgsForCall []struct {
		arg1 string
	}
	generateEmailVerificationTokenReturns struct {
		result1 *models.User
		result2 error
	}
	generateEmailVerificationTokenReturnsOnCall map[int]struct {
		result1 *models.User
		result2 error
	}
	NewUserAccountStub        func(string, string, string) (*models.User, error)
	newUserAccountMutex       sync.RWMutex
	newUserAccountArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	newUserAccountReturns struct {
		result1 *models.User
		result2 error
	}
	newUserAccountReturnsOnCall map[int]struct {
		result1 *models.User
		result2 error
	}
	ResetPasswordStub        func(string) (string, error)
	resetPasswordMutex       sync.RWMutex
	resetPasswordArgsForCall []struct {
		arg1 string
	}
	resetPasswordReturns struct {
		result1 string
		result2 error
	}
	resetPasswordReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	SignInStub        func(string, string) (bool, error)
	signInMutex       sync.RWMutex
	signInArgsForCall []struct {
		arg1 string
		arg2 string
	}
	signInReturns struct {
		result1 bool
		result2 error
	}
	signInReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	UpdateCustomerObjectHashStub        func(string, string) error
	updateCustomerObjectHashMutex       sync.RWMutex
	updateCustomerObjectHashArgsForCall []struct {
		arg1 string
		arg2 string
	}
	updateCustomerObjectHashReturns struct {
		result1 error
	}
	updateCustomerObjectHashReturnsOnCall map[int]struct {
		result1 error
	}
	ValidateEmailVerificationTokenStub        func(string, string) (*models.User, error)
	validateEmailVerificationTokenMutex       sync.RWMutex
	validateEmailVerificationTokenArgsForCall []struct {
		arg1 string
		arg2 string
	}
	validateEmailVerificationTokenReturns struct {
		result1 *models.User
		result2 error
	}
	validateEmailVerificationTokenReturnsOnCall map[int]struct {
		result1 *models.User
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUserManager) ChangePassword(arg1 string, arg2 string, arg3 string) (bool, error) {
	fake.changePasswordMutex.Lock()
	ret, specificReturn := fake.changePasswordReturnsOnCall[len(fake.changePasswordArgsForCall)]
	fake.changePasswordArgsForCall = append(fake.changePasswordArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("ChangePassword", []interface{}{arg1, arg2, arg3})
	fake.changePasswordMutex.Unlock()
	if fake.ChangePasswordStub != nil {
		return fake.ChangePasswordStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.changePasswordReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUserManager) ChangePasswordCallCount() int {
	fake.changePasswordMutex.RLock()
	defer fake.changePasswordMutex.RUnlock()
	return len(fake.changePasswordArgsForCall)
}

func (fake *FakeUserManager) ChangePasswordCalls(stub func(string, string, string) (bool, error)) {
	fake.changePasswordMutex.Lock()
	defer fake.changePasswordMutex.Unlock()
	fake.ChangePasswordStub = stub
}

func (fake *FakeUserManager) ChangePasswordArgsForCall(i int) (string, string, string) {
	fake.changePasswordMutex.RLock()
	defer fake.changePasswordMutex.RUnlock()
	argsForCall := fake.changePasswordArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeUserManager) ChangePasswordReturns(result1 bool, result2 error) {
	fake.changePasswordMutex.Lock()
	defer fake.changePasswordMutex.Unlock()
	fake.ChangePasswordStub = nil
	fake.changePasswordReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeUserManager) ChangePasswordReturnsOnCall(i int, result1 bool, result2 error) {
	fake.changePasswordMutex.Lock()
	defer fake.changePasswordMutex.Unlock()
	fake.ChangePasswordStub = nil
	if fake.changePasswordReturnsOnCall == nil {
		fake.changePasswordReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.changePasswordReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeUserManager) FindByEmail(arg1 string) (*models.User, error) {
	fake.findByEmailMutex.Lock()
	ret, specificReturn := fake.findByEmailReturnsOnCall[len(fake.findByEmailArgsForCall)]
	fake.findByEmailArgsForCall = append(fake.findByEmailArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("FindByEmail", []interface{}{arg1})
	fake.findByEmailMutex.Unlock()
	if fake.FindByEmailStub != nil {
		return fake.FindByEmailStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.findByEmailReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUserManager) FindByEmailCallCount() int {
	fake.findByEmailMutex.RLock()
	defer fake.findByEmailMutex.RUnlock()
	return len(fake.findByEmailArgsForCall)
}

func (fake *FakeUserManager) FindByEmailCalls(stub func(string) (*models.User, error)) {
	fake.findByEmailMutex.Lock()
	defer fake.findByEmailMutex.Unlock()
	fake.FindByEmailStub = stub
}

func (fake *FakeUserManager) FindByEmailArgsForCall(i int) string {
	fake.findByEmailMutex.RLock()
	defer fake.findByEmailMutex.RUnlock()
	argsForCall := fake.findByEmailArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeUserManager) FindByEmailReturns(result1 *models.User, result2 error) {
	fake.findByEmailMutex.Lock()
	defer fake.findByEmailMutex.Unlock()
	fake.FindByEmailStub = nil
	fake.findByEmailReturns = struct {
		result1 *models.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserManager) FindByEmailReturnsOnCall(i int, result1 *models.User, result2 error) {
	fake.findByEmailMutex.Lock()
	defer fake.findByEmailMutex.Unlock()
	fake.FindByEmailStub = nil
	if fake.findByEmailReturnsOnCall == nil {
		fake.findByEmailReturnsOnCall = make(map[int]struct {
			result1 *models.User
			result2 error
		})
	}
	fake.findByEmailReturnsOnCall[i] = struct {
		result1 *models.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserManager) FindByUserName(arg1 string) (*models.User, error) {
	fake.findByUserNameMutex.Lock()
	ret, specificReturn := fake.findByUserNameReturnsOnCall[len(fake.findByUserNameArgsForCall)]
	fake.findByUserNameArgsForCall = append(fake.findByUserNameArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("FindByUserName", []interface{}{arg1})
	fake.findByUserNameMutex.Unlock()
	if fake.FindByUserNameStub != nil {
		return fake.FindByUserNameStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.findByUserNameReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUserManager) FindByUserNameCallCount() int {
	fake.findByUserNameMutex.RLock()
	defer fake.findByUserNameMutex.RUnlock()
	return len(fake.findByUserNameArgsForCall)
}

func (fake *FakeUserManager) FindByUserNameCalls(stub func(string) (*models.User, error)) {
	fake.findByUserNameMutex.Lock()
	defer fake.findByUserNameMutex.Unlock()
	fake.FindByUserNameStub = stub
}

func (fake *FakeUserManager) FindByUserNameArgsForCall(i int) string {
	fake.findByUserNameMutex.RLock()
	defer fake.findByUserNameMutex.RUnlock()
	argsForCall := fake.findByUserNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeUserManager) FindByUserNameReturns(result1 *models.User, result2 error) {
	fake.findByUserNameMutex.Lock()
	defer fake.findByUserNameMutex.Unlock()
	fake.FindByUserNameStub = nil
	fake.findByUserNameReturns = struct {
		result1 *models.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserManager) FindByUserNameReturnsOnCall(i int, result1 *models.User, result2 error) {
	fake.findByUserNameMutex.Lock()
	defer fake.findByUserNameMutex.Unlock()
	fake.FindByUserNameStub = nil
	if fake.findByUserNameReturnsOnCall == nil {
		fake.findByUserNameReturnsOnCall = make(map[int]struct {
			result1 *models.User
			result2 error
		})
	}
	fake.findByUserNameReturnsOnCall[i] = struct {
		result1 *models.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserManager) GenerateEmailVerificationToken(arg1 string) (*models.User, error) {
	fake.generateEmailVerificationTokenMutex.Lock()
	ret, specificReturn := fake.generateEmailVerificationTokenReturnsOnCall[len(fake.generateEmailVerificationTokenArgsForCall)]
	fake.generateEmailVerificationTokenArgsForCall = append(fake.generateEmailVerificationTokenArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GenerateEmailVerificationToken", []interface{}{arg1})
	fake.generateEmailVerificationTokenMutex.Unlock()
	if fake.GenerateEmailVerificationTokenStub != nil {
		return fake.GenerateEmailVerificationTokenStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.generateEmailVerificationTokenReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUserManager) GenerateEmailVerificationTokenCallCount() int {
	fake.generateEmailVerificationTokenMutex.RLock()
	defer fake.generateEmailVerificationTokenMutex.RUnlock()
	return len(fake.generateEmailVerificationTokenArgsForCall)
}

func (fake *FakeUserManager) GenerateEmailVerificationTokenCalls(stub func(string) (*models.User, error)) {
	fake.generateEmailVerificationTokenMutex.Lock()
	defer fake.generateEmailVerificationTokenMutex.Unlock()
	fake.GenerateEmailVerificationTokenStub = stub
}

func (fake *FakeUserManager) GenerateEmailVerificationTokenArgsForCall(i int) string {
	fake.generateEmailVerificationTokenMutex.RLock()
	defer fake.generateEmailVerificationTokenMutex.RUnlock()
	argsForCall := fake.generateEmailVerificationTokenArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeUserManager) GenerateEmailVerificationTokenReturns(result1 *models.User, result2 error) {
	fake.generateEmailVerificationTokenMutex.Lock()
	defer fake.generateEmailVerificationTokenMutex.Unlock()
	fake.GenerateEmailVerificationTokenStub = nil
	fake.generateEmailVerificationTokenReturns = struct {
		result1 *models.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserManager) GenerateEmailVerificationTokenReturnsOnCall(i int, result1 *models.User, result2 error) {
	fake.generateEmailVerificationTokenMutex.Lock()
	defer fake.generateEmailVerificationTokenMutex.Unlock()
	fake.GenerateEmailVerificationTokenStub = nil
	if fake.generateEmailVerificationTokenReturnsOnCall == nil {
		fake.generateEmailVerificationTokenReturnsOnCall = make(map[int]struct {
			result1 *models.User
			result2 error
		})
	}
	fake.generateEmailVerificationTokenReturnsOnCall[i] = struct {
		result1 *models.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserManager) NewUserAccount(arg1 string, arg2 string, arg3 string) (*models.User, error) {
	fake.newUserAccountMutex.Lock()
	ret, specificReturn := fake.newUserAccountReturnsOnCall[len(fake.newUserAccountArgsForCall)]
	fake.newUserAccountArgsForCall = append(fake.newUserAccountArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("NewUserAccount", []interface{}{arg1, arg2, arg3})
	fake.newUserAccountMutex.Unlock()
	if fake.NewUserAccountStub != nil {
		return fake.NewUserAccountStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.newUserAccountReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUserManager) NewUserAccountCallCount() int {
	fake.newUserAccountMutex.RLock()
	defer fake.newUserAccountMutex.RUnlock()
	return len(fake.newUserAccountArgsForCall)
}

func (fake *FakeUserManager) NewUserAccountCalls(stub func(string, string, string) (*models.User, error)) {
	fake.newUserAccountMutex.Lock()
	defer fake.newUserAccountMutex.Unlock()
	fake.NewUserAccountStub = stub
}

func (fake *FakeUserManager) NewUserAccountArgsForCall(i int) (string, string, string) {
	fake.newUserAccountMutex.RLock()
	defer fake.newUserAccountMutex.RUnlock()
	argsForCall := fake.newUserAccountArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeUserManager) NewUserAccountReturns(result1 *models.User, result2 error) {
	fake.newUserAccountMutex.Lock()
	defer fake.newUserAccountMutex.Unlock()
	fake.NewUserAccountStub = nil
	fake.newUserAccountReturns = struct {
		result1 *models.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserManager) NewUserAccountReturnsOnCall(i int, result1 *models.User, result2 error) {
	fake.newUserAccountMutex.Lock()
	defer fake.newUserAccountMutex.Unlock()
	fake.NewUserAccountStub = nil
	if fake.newUserAccountReturnsOnCall == nil {
		fake.newUserAccountReturnsOnCall = make(map[int]struct {
			result1 *models.User
			result2 error
		})
	}
	fake.newUserAccountReturnsOnCall[i] = struct {
		result1 *models.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserManager) ResetPassword(arg1 string) (string, error) {
	fake.resetPasswordMutex.Lock()
	ret, specificReturn := fake.resetPasswordReturnsOnCall[len(fake.resetPasswordArgsForCall)]
	fake.resetPasswordArgsForCall = append(fake.resetPasswordArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ResetPassword", []interface{}{arg1})
	fake.resetPasswordMutex.Unlock()
	if fake.ResetPasswordStub != nil {
		return fake.ResetPasswordStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.resetPasswordReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUserManager) ResetPasswordCallCount() int {
	fake.resetPasswordMutex.RLock()
	defer fake.resetPasswordMutex.RUnlock()
	return len(fake.resetPasswordArgsForCall)
}

func (fake *FakeUserManager) ResetPasswordCalls(stub func(string) (string, error)) {
	fake.resetPasswordMutex.Lock()
	defer fake.resetPasswordMutex.Unlock()
	fake.ResetPasswordStub = stub
}

func (fake *FakeUserManager) ResetPasswordArgsForCall(i int) string {
	fake.resetPasswordMutex.RLock()
	defer fake.resetPasswordMutex.RUnlock()
	argsForCall := fake.resetPasswordArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeUserManager) ResetPasswordReturns(result1 string, result2 error) {
	fake.resetPasswordMutex.Lock()
	defer fake.resetPasswordMutex.Unlock()
	fake.ResetPasswordStub = nil
	fake.resetPasswordReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeUserManager) ResetPasswordReturnsOnCall(i int, result1 string, result2 error) {
	fake.resetPasswordMutex.Lock()
	defer fake.resetPasswordMutex.Unlock()
	fake.ResetPasswordStub = nil
	if fake.resetPasswordReturnsOnCall == nil {
		fake.resetPasswordReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.resetPasswordReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeUserManager) SignIn(arg1 string, arg2 string) (bool, error) {
	fake.signInMutex.Lock()
	ret, specificReturn := fake.signInReturnsOnCall[len(fake.signInArgsForCall)]
	fake.signInArgsForCall = append(fake.signInArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("SignIn", []interface{}{arg1, arg2})
	fake.signInMutex.Unlock()
	if fake.SignInStub != nil {
		return fake.SignInStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.signInReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUserManager) SignInCallCount() int {
	fake.signInMutex.RLock()
	defer fake.signInMutex.RUnlock()
	return len(fake.signInArgsForCall)
}

func (fake *FakeUserManager) SignInCalls(stub func(string, string) (bool, error)) {
	fake.signInMutex.Lock()
	defer fake.signInMutex.Unlock()
	fake.SignInStub = stub
}

func (fake *FakeUserManager) SignInArgsForCall(i int) (string, string) {
	fake.signInMutex.RLock()
	defer fake.signInMutex.RUnlock()
	argsForCall := fake.signInArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeUserManager) SignInReturns(result1 bool, result2 error) {
	fake.signInMutex.Lock()
	defer fake.signInMutex.Unlock()
	fake.SignInStub = nil
	fake.signInReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeUserManager) SignInReturnsOnCall(i int, result1 bool, result2 error) {
	fake.signInMutex.Lock()
	defer fake.signInMutex.Unlock()
	fake.SignInStub = nil
	if fake.signInReturnsOnCall == nil {
		fake.signInReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.signInReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeUserManager) UpdateCustomerObjectHash(arg1 string, arg2 string) error {
	fake.updateCustomerObjectHashMutex.Lock()
	ret, specificReturn := fake.updateCustomerObjectHashReturnsOnCall[len(fake.updateCustomerObjectHashArgsForCall)]
	fake.updateCustomerObjectHashArgsForCall = append(fake.updateCustomerObjectHashArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("UpdateCustomerObjectHash", []interface{}{arg1, arg2})
	fake.updateCustomerObjectHashMutex.Unlock()
	if fake.UpdateCustomerObjectHashStub != nil {
		return fake.UpdateCustomerObjectHashStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.updateCustomerObjectHashReturns
	return fakeReturns.result1
}

func (fake *FakeUserManager) UpdateCustomerObjectHashCallCount() int {
	fake.updateCustomerObjectHashMutex.RLock()
	defer fake.updateCustomerObjectHashMutex.RUnlock()
	return len(fake.updateCustomerObjectHashArgsForCall)
}

func (fake *FakeUserManager) UpdateCustomerObjectHashCalls(stub func(string, string) error) {
	fake.updateCustomerObjectHashMutex.Lock()
	defer fake.updateCustomerObjectHashMutex.Unlock()
	fake.UpdateCustomerObjectHashStub = stub
}

func (fake *FakeUserManager) UpdateCustomerObjectHashArgsForCall(i int) (string, string) {
	fake.updateCustomerObjectHashMutex.RLock()
	defer fake.updateCustomerObjectHashMutex.RUnlock()
	argsForCall := fake.updateCustomerObjectHashArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeUserManager) UpdateCustomerObjectHashReturns(result1 error) {
	fake.updateCustomerObjectHashMutex.Lock()
	defer fake.updateCustomerObjectHashMutex.Unlock()
	fake.UpdateCustomerObjectHashStub = nil
	fake.updateCustomerObjectHashReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserManager) UpdateCustomerObjectHashReturnsOnCall(i int, result1 error) {
	fake.updateCustomerObjectHashMutex.Lock()
	defer fake.updateCustomerObjectHashMutex.Unlock()
	fake.UpdateCustomerObjectHashStub = nil
	if fake.updateCustomerObjectHashReturnsOnCall == nil {
		fake.updateCustomerObjectHashReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateCustomerObjectHashReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserManager) ValidateEmailVerificationToken(arg1 string, arg2 string) (*models.User, error) {
	fake.validateEmailVerificationTokenMutex.Lock()
	ret, specificReturn := fake.validateEmailVerificationTokenReturnsOnCall[len(fake.validateEmailVerificationTokenArgsForCall)]
	fake.validateEmailVerificationTokenArgsForCall = append(fake.validateEmailVerificationTokenArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("ValidateEmailVerificationToken", []interface{}{arg1, arg2})
	fake.validateEmailVerificationTokenMutex.Unlock()
	if fake.ValidateEmailVerificationTokenStub != nil {
		return fake.ValidateEmailVerificationTokenStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.validateEmailVerificationTokenReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUserManager) ValidateEmailVerificationTokenCallCount() int {
	fake.validateEmailVerificationTokenMutex.RLock()
	defer fake.validateEmailVerificationTokenMutex.RUnlock()
	return len(fake.validateEmailVerificationTokenArgsForCall)
}

func (fake *FakeUserManager) ValidateEmailVerificationTokenCalls(stub func(string, string) (*models.User, error)) {
	fake.validateEmailVerificationTokenMutex.Lock()
	defer fake.validateEmailVerificationTokenMutex.Unlock()
	fake.ValidateEmailVerificationTokenStub = stub
}

func (fake *FakeUserManager) ValidateEmailVerificationTokenArgsForCall(i int) (string, string) {
	fake.validateEmailVerificationTokenMutex.RLock()
	defer fake.validateEmailVerificationTokenMutex.RUnlock()
	argsForCall := fake.validateEmailVerificationTokenArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeUserManager) ValidateEmailVerificationTokenReturns(result1 *models.User, result2 error) {
	fake.validateEmailVerificationTokenMutex.Lock()
	defer fake.validateEmailVerificationTokenMutex.Unlock()
	fake.ValidateEmailVerificationTokenStub = nil
	fake.validateEmailVerificationTokenReturns = struct {
		result1 *models.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserManager) ValidateEmailVerificationTokenReturnsOnCall(i int, result1 *models.User, result2 error) {
	fake.validateEmailVerificationTokenMutex.Lock()
	defer fake.validateEmailVerificationTokenMutex.Unlock()
	fake.ValidateEmailVerificationTokenStub = nil
	if fake.validateEmailVerificationTokenReturnsOnCall == nil {
		fake.validateEmailVerificationTokenReturnsOnCall = make(map[int]struct {
			result1 *models.User
			result2 error
		})
	}
	fake.validateEmailVerificationTokenReturnsOnCall[i] = struct {
		result1 *models.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.changePasswordMutex.RLock()
	defer fake.changePasswordMutex.RUnlock()
	fake.findByEmailMutex.RLock()
	defer fake.findByEmailMutex.RUnlock()
	fake.findByUserNameMutex.RLock()
	defer fake.findByUserNameMutex.RUnlock()
	fake.generateEmailVerificationTokenMutex.RLock()
	defer fake.generateEmailVerificationTokenMutex.RUnlock()
	fake.newUserAccountMutex.RLock()
	defer fake.newUserAccountMutex.RUnlock()
	fake.resetPasswordMutex.RLock()
	defer fake.resetPasswordMutex.RUnlock()
	fake.signInMutex.RLock()
	defer fake.signInMutex.RUnlock()
	fake.updateCustomerObjectHashMutex.RLock()
	defer fake.updateCustomerObjectHashMutex.RUnlock()
	fake.validateEmailVerificationTokenMutex.RLock()
	defer fake.validateEmailVerificationTokenMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeUserManager) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}
