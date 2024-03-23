// Code generated -- DO NOT EDIT
package gen_mock

import (
	its "github.com/youta-t/its"
	itskit "github.com/youta-t/its/itskit"
	mockkit "github.com/youta-t/its/mocker/mockkit"
	testee "github.com/youta-t/its/mocker/internal/example"
	
)
type _SessionStoreCallSpec struct {
	cookie its.Matcher[string]
	
	
}

type _SessionStoreCall struct {
	name itskit.Label
	spec _SessionStoreCallSpec
}

func SessionStore_Expects(
	cookie its.Matcher[string],
	
) _SessionStoreCall {
	cancel := itskit.SkipStack()
	defer cancel()

	spec := _SessionStoreCallSpec {}
	spec.cookie = itskit.Named(
		"cookie",
		cookie,
	)
	
	
	return _SessionStoreCall{
		name: itskit.NewLabelWithLocation("func SessionStore"),
		spec: spec,
	}
}

type _SessionStoreBehavior  struct {
	name itskit.Label
	spec _SessionStoreCallSpec
	effect func(cookie string) (userId string, ok bool)
}

func (b *_SessionStoreBehavior) Fn(t mockkit.TestLike) func(cookie string) (userId string, ok bool) {
	return func (
		
		arg0 string,
		
		
	) (
		string,
		bool,
		
	) {
		if h, ok := t.(interface { Helper() }); ok {
			h.Helper()
		}
		ok := 0
		matches := []itskit.Match{}
		
		{
			matcher := b.spec.cookie
			if matcher == nil {
				matcher = its.Never[string]()
			}
			m := matcher.Match(arg0)
			if m.Ok() {
				ok += 1
			}
			matches = append(matches, m)
		}
		
		itskit.NewMatch(
			ok == len(matches),
			b.name.Fill(itskit.Missing),
			matches...,
		).OrError(t)
		return b.effect(
			
			arg0,
			
			
		)
	}
}

func (c _SessionStoreCall) ThenReturn(

	ret0 string,

	ret1 bool,

) mockkit.FuncBehavior[ func (cookie string) (userId string, ok bool)  ] {
	return c.ThenEffect(func(
		
		string,
		
		
	)(
		string,
		bool,
		
	){
		
		return ret0,  ret1
		
	})
}

func (c _SessionStoreCall) ThenEffect(effect func(cookie string) (userId string, ok bool)) mockkit.FuncBehavior[ func (cookie string) (userId string, ok bool) ] {
	return &_SessionStoreBehavior {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}


type _UserRegistry_GetCallSpec struct {
	userId its.Matcher[string]
	
	
}

type _UserRegistry_GetCall struct {
	name itskit.Label
	spec _UserRegistry_GetCallSpec
}

func UserRegistry_Get_Expects(
	userId its.Matcher[string],
	
) _UserRegistry_GetCall {
	cancel := itskit.SkipStack()
	defer cancel()

	spec := _UserRegistry_GetCallSpec {}
	spec.userId = itskit.Named(
		"userId",
		userId,
	)
	
	
	return _UserRegistry_GetCall{
		name: itskit.NewLabelWithLocation("func UserRegistry_Get"),
		spec: spec,
	}
}

type _UserRegistry_GetBehavior  struct {
	name itskit.Label
	spec _UserRegistry_GetCallSpec
	effect func(userId string) ( testee.User,  error)
}

func (b *_UserRegistry_GetBehavior) Fn(t mockkit.TestLike) func(userId string) ( testee.User,  error) {
	return func (
		
		arg0 string,
		
		
	) (
		testee.User,
		error,
		
	) {
		if h, ok := t.(interface { Helper() }); ok {
			h.Helper()
		}
		ok := 0
		matches := []itskit.Match{}
		
		{
			matcher := b.spec.userId
			if matcher == nil {
				matcher = its.Never[string]()
			}
			m := matcher.Match(arg0)
			if m.Ok() {
				ok += 1
			}
			matches = append(matches, m)
		}
		
		itskit.NewMatch(
			ok == len(matches),
			b.name.Fill(itskit.Missing),
			matches...,
		).OrError(t)
		return b.effect(
			
			arg0,
			
			
		)
	}
}

func (c _UserRegistry_GetCall) ThenReturn(

	ret0 testee.User,

	ret1 error,

) mockkit.FuncBehavior[ func (userId string) ( testee.User,  error)  ] {
	return c.ThenEffect(func(
		
		string,
		
		
	)(
		testee.User,
		error,
		
	){
		
		return ret0,  ret1
		
	})
}

func (c _UserRegistry_GetCall) ThenEffect(effect func(userId string) ( testee.User,  error)) mockkit.FuncBehavior[ func (userId string) ( testee.User,  error) ] {
	return &_UserRegistry_GetBehavior {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}


type _UserRegistry_UpdateCallSpec struct {
	arg0 its.Matcher[testee.User]
	
	
}

type _UserRegistry_UpdateCall struct {
	name itskit.Label
	spec _UserRegistry_UpdateCallSpec
}

func UserRegistry_Update_Expects(
	arg0 its.Matcher[testee.User],
	
) _UserRegistry_UpdateCall {
	cancel := itskit.SkipStack()
	defer cancel()

	spec := _UserRegistry_UpdateCallSpec {}
	spec.arg0 = itskit.Named(
		"arg0",
		arg0,
	)
	
	
	return _UserRegistry_UpdateCall{
		name: itskit.NewLabelWithLocation("func UserRegistry_Update"),
		spec: spec,
	}
}

type _UserRegistry_UpdateBehavior  struct {
	name itskit.Label
	spec _UserRegistry_UpdateCallSpec
	effect func(arg0 testee.User) error
}

func (b *_UserRegistry_UpdateBehavior) Fn(t mockkit.TestLike) func(arg0 testee.User) error {
	return func (
		
		arg0 testee.User,
		
		
	) (
		error,
		
	) {
		if h, ok := t.(interface { Helper() }); ok {
			h.Helper()
		}
		ok := 0
		matches := []itskit.Match{}
		
		{
			matcher := b.spec.arg0
			if matcher == nil {
				matcher = its.Never[testee.User]()
			}
			m := matcher.Match(arg0)
			if m.Ok() {
				ok += 1
			}
			matches = append(matches, m)
		}
		
		itskit.NewMatch(
			ok == len(matches),
			b.name.Fill(itskit.Missing),
			matches...,
		).OrError(t)
		return b.effect(
			
			arg0,
			
			
		)
	}
}

func (c _UserRegistry_UpdateCall) ThenReturn(

	ret0 error,

) mockkit.FuncBehavior[ func (arg0 testee.User) error  ] {
	return c.ThenEffect(func(
		
		testee.User,
		
		
	)(
		error,
		
	){
		
		return ret0
		
	})
}

func (c _UserRegistry_UpdateCall) ThenEffect(effect func(arg0 testee.User) error) mockkit.FuncBehavior[ func (arg0 testee.User) error ] {
	return &_UserRegistry_UpdateBehavior {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}


type _UserRegistry_DeleteCallSpec struct {
	arg0 its.Matcher[testee.User]
	
	
}

type _UserRegistry_DeleteCall struct {
	name itskit.Label
	spec _UserRegistry_DeleteCallSpec
}

func UserRegistry_Delete_Expects(
	arg0 its.Matcher[testee.User],
	
) _UserRegistry_DeleteCall {
	cancel := itskit.SkipStack()
	defer cancel()

	spec := _UserRegistry_DeleteCallSpec {}
	spec.arg0 = itskit.Named(
		"arg0",
		arg0,
	)
	
	
	return _UserRegistry_DeleteCall{
		name: itskit.NewLabelWithLocation("func UserRegistry_Delete"),
		spec: spec,
	}
}

type _UserRegistry_DeleteBehavior  struct {
	name itskit.Label
	spec _UserRegistry_DeleteCallSpec
	effect func(arg0 testee.User) error
}

func (b *_UserRegistry_DeleteBehavior) Fn(t mockkit.TestLike) func(arg0 testee.User) error {
	return func (
		
		arg0 testee.User,
		
		
	) (
		error,
		
	) {
		if h, ok := t.(interface { Helper() }); ok {
			h.Helper()
		}
		ok := 0
		matches := []itskit.Match{}
		
		{
			matcher := b.spec.arg0
			if matcher == nil {
				matcher = its.Never[testee.User]()
			}
			m := matcher.Match(arg0)
			if m.Ok() {
				ok += 1
			}
			matches = append(matches, m)
		}
		
		itskit.NewMatch(
			ok == len(matches),
			b.name.Fill(itskit.Missing),
			matches...,
		).OrError(t)
		return b.effect(
			
			arg0,
			
			
		)
	}
}

func (c _UserRegistry_DeleteCall) ThenReturn(

	ret0 error,

) mockkit.FuncBehavior[ func (arg0 testee.User) error  ] {
	return c.ThenEffect(func(
		
		testee.User,
		
		
	)(
		error,
		
	){
		
		return ret0
		
	})
}

func (c _UserRegistry_DeleteCall) ThenEffect(effect func(arg0 testee.User) error) mockkit.FuncBehavior[ func (arg0 testee.User) error ] {
	return &_UserRegistry_DeleteBehavior {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}




type _UserRegistryImpl struct {
	
	Get func(userId string) ( testee.User,  error)
	Update func(arg0 testee.User) error
	Delete func(arg0 testee.User) error
}

func UserRegistry_Build(t mockkit.TestLike, spec UserRegistry_Spec) testee.UserRegistry {
	impl := _UserRegistryImpl{}

	
	if spec.Get != nil {
		impl.Get = spec.Get.Fn(t)
	}
	
	if spec.Update != nil {
		impl.Update = spec.Update.Fn(t)
	}
	
	if spec.Delete != nil {
		impl.Delete = spec.Delete.Fn(t)
	}
	

	return _UserRegistryMock { t: t, impl: impl }
}

type _UserRegistryMock struct {
	t mockkit.TestLike
	impl _UserRegistryImpl
}

func (m _UserRegistryMock) Get (
	userId string,
) (
	testee.User,
	error,
) {
	cancel := itskit.SkipStack()
	defer cancel()

	if m.impl.Get == nil {
		itskit.NG(
			itskit.NewLabelWithLocation("UserRegistry.Get is not mocked").String(),
		).OrFatal(m.t)
	}

	return m.impl.Get(
		userId,
	)
}


func (m _UserRegistryMock) Update (
	arg0 testee.User,
) (
	error,
) {
	cancel := itskit.SkipStack()
	defer cancel()

	if m.impl.Update == nil {
		itskit.NG(
			itskit.NewLabelWithLocation("UserRegistry.Update is not mocked").String(),
		).OrFatal(m.t)
	}

	return m.impl.Update(
		arg0,
	)
}


func (m _UserRegistryMock) Delete (
	arg0 testee.User,
) (
	error,
) {
	cancel := itskit.SkipStack()
	defer cancel()

	if m.impl.Delete == nil {
		itskit.NG(
			itskit.NewLabelWithLocation("UserRegistry.Delete is not mocked").String(),
		).OrFatal(m.t)
	}

	return m.impl.Delete(
		arg0,
	)
}


type UserRegistry_Spec struct {
	
	Get mockkit.FuncBehavior[func (userId string) ( testee.User,  error)]
	
	Update mockkit.FuncBehavior[func (arg0 testee.User) error]
	
	Delete mockkit.FuncBehavior[func (arg0 testee.User) error]
	
}

