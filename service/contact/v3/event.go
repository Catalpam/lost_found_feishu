// Code generated by lark suite oapi sdk gen
package v3

import (
	"lost_found/core"
	"lost_found/core/config"
	"lost_found/event"
)

type DepartmentCreatedEventHandler struct {
	Fn func(*core.Context, *DepartmentCreatedEvent) error
}

func (h *DepartmentCreatedEventHandler) GetEvent() interface{} {
	return &DepartmentCreatedEvent{}
}

func (h *DepartmentCreatedEventHandler) Handle(ctx *core.Context, event interface{}) error {
	return h.Fn(ctx, event.(*DepartmentCreatedEvent))
}

func SetDepartmentCreatedEventHandler(conf *config.Config, fn func(ctx *core.Context, event *DepartmentCreatedEvent) error) {
	event.SetTypeHandler(conf, "contact.department.created_v3", &DepartmentCreatedEventHandler{Fn: fn})
}

type DepartmentDeletedEventHandler struct {
	Fn func(*core.Context, *DepartmentDeletedEvent) error
}

func (h *DepartmentDeletedEventHandler) GetEvent() interface{} {
	return &DepartmentDeletedEvent{}
}

func (h *DepartmentDeletedEventHandler) Handle(ctx *core.Context, event interface{}) error {
	return h.Fn(ctx, event.(*DepartmentDeletedEvent))
}

func SetDepartmentDeletedEventHandler(conf *config.Config, fn func(ctx *core.Context, event *DepartmentDeletedEvent) error) {
	event.SetTypeHandler(conf, "contact.department.deleted_v3", &DepartmentDeletedEventHandler{Fn: fn})
}

type DepartmentUpdatedEventHandler struct {
	Fn func(*core.Context, *DepartmentUpdatedEvent) error
}

func (h *DepartmentUpdatedEventHandler) GetEvent() interface{} {
	return &DepartmentUpdatedEvent{}
}

func (h *DepartmentUpdatedEventHandler) Handle(ctx *core.Context, event interface{}) error {
	return h.Fn(ctx, event.(*DepartmentUpdatedEvent))
}

func SetDepartmentUpdatedEventHandler(conf *config.Config, fn func(ctx *core.Context, event *DepartmentUpdatedEvent) error) {
	event.SetTypeHandler(conf, "contact.department.updated_v3", &DepartmentUpdatedEventHandler{Fn: fn})
}

type UserCreatedEventHandler struct {
	Fn func(*core.Context, *UserCreatedEvent) error
}

func (h *UserCreatedEventHandler) GetEvent() interface{} {
	return &UserCreatedEvent{}
}

func (h *UserCreatedEventHandler) Handle(ctx *core.Context, event interface{}) error {
	return h.Fn(ctx, event.(*UserCreatedEvent))
}

func SetUserCreatedEventHandler(conf *config.Config, fn func(ctx *core.Context, event *UserCreatedEvent) error) {
	event.SetTypeHandler(conf, "contact.user.created_v3", &UserCreatedEventHandler{Fn: fn})
}

type UserDeletedEventHandler struct {
	Fn func(*core.Context, *UserDeletedEvent) error
}

func (h *UserDeletedEventHandler) GetEvent() interface{} {
	return &UserDeletedEvent{}
}

func (h *UserDeletedEventHandler) Handle(ctx *core.Context, event interface{}) error {
	return h.Fn(ctx, event.(*UserDeletedEvent))
}

func SetUserDeletedEventHandler(conf *config.Config, fn func(ctx *core.Context, event *UserDeletedEvent) error) {
	event.SetTypeHandler(conf, "contact.user.deleted_v3", &UserDeletedEventHandler{Fn: fn})
}

type UserUpdatedEventHandler struct {
	Fn func(*core.Context, *UserUpdatedEvent) error
}

func (h *UserUpdatedEventHandler) GetEvent() interface{} {
	return &UserUpdatedEvent{}
}

func (h *UserUpdatedEventHandler) Handle(ctx *core.Context, event interface{}) error {
	return h.Fn(ctx, event.(*UserUpdatedEvent))
}

func SetUserUpdatedEventHandler(conf *config.Config, fn func(ctx *core.Context, event *UserUpdatedEvent) error) {
	event.SetTypeHandler(conf, "contact.user.updated_v3", &UserUpdatedEventHandler{Fn: fn})
}

type UserGroupCreatedEventHandler struct {
	Fn func(*core.Context, *UserGroupCreatedEvent) error
}

func (h *UserGroupCreatedEventHandler) GetEvent() interface{} {
	return &UserGroupCreatedEvent{}
}

func (h *UserGroupCreatedEventHandler) Handle(ctx *core.Context, event interface{}) error {
	return h.Fn(ctx, event.(*UserGroupCreatedEvent))
}

func SetUserGroupCreatedEventHandler(conf *config.Config, fn func(ctx *core.Context, event *UserGroupCreatedEvent) error) {
	event.SetTypeHandler(conf, "contact.user_group.created_v3", &UserGroupCreatedEventHandler{Fn: fn})
}

type UserGroupDeletedEventHandler struct {
	Fn func(*core.Context, *UserGroupDeletedEvent) error
}

func (h *UserGroupDeletedEventHandler) GetEvent() interface{} {
	return &UserGroupDeletedEvent{}
}

func (h *UserGroupDeletedEventHandler) Handle(ctx *core.Context, event interface{}) error {
	return h.Fn(ctx, event.(*UserGroupDeletedEvent))
}

func SetUserGroupDeletedEventHandler(conf *config.Config, fn func(ctx *core.Context, event *UserGroupDeletedEvent) error) {
	event.SetTypeHandler(conf, "contact.user_group.deleted_v3", &UserGroupDeletedEventHandler{Fn: fn})
}

type UserGroupUpdatedEventHandler struct {
	Fn func(*core.Context, *UserGroupUpdatedEvent) error
}

func (h *UserGroupUpdatedEventHandler) GetEvent() interface{} {
	return &UserGroupUpdatedEvent{}
}

func (h *UserGroupUpdatedEventHandler) Handle(ctx *core.Context, event interface{}) error {
	return h.Fn(ctx, event.(*UserGroupUpdatedEvent))
}

func SetUserGroupUpdatedEventHandler(conf *config.Config, fn func(ctx *core.Context, event *UserGroupUpdatedEvent) error) {
	event.SetTypeHandler(conf, "contact.user_group.updated_v3", &UserGroupUpdatedEventHandler{Fn: fn})
}

type ScopeUpdatedEventHandler struct {
	Fn func(*core.Context, *ScopeUpdatedEvent) error
}

func (h *ScopeUpdatedEventHandler) GetEvent() interface{} {
	return &ScopeUpdatedEvent{}
}

func (h *ScopeUpdatedEventHandler) Handle(ctx *core.Context, event interface{}) error {
	return h.Fn(ctx, event.(*ScopeUpdatedEvent))
}

func SetScopeUpdatedEventHandler(conf *config.Config, fn func(ctx *core.Context, event *ScopeUpdatedEvent) error) {
	event.SetTypeHandler(conf, "contact.scope.updated_v3", &ScopeUpdatedEventHandler{Fn: fn})
}

type UserGroupMemberChangedEventHandler struct {
	Fn func(*core.Context, *UserGroupMemberChangedEvent) error
}

func (h *UserGroupMemberChangedEventHandler) GetEvent() interface{} {
	return &UserGroupMemberChangedEvent{}
}

func (h *UserGroupMemberChangedEventHandler) Handle(ctx *core.Context, event interface{}) error {
	return h.Fn(ctx, event.(*UserGroupMemberChangedEvent))
}

func SetUserGroupMemberChangedEventHandler(conf *config.Config, fn func(ctx *core.Context, event *UserGroupMemberChangedEvent) error) {
	event.SetTypeHandler(conf, "contact.user_group.member.changed_v3", &UserGroupMemberChangedEventHandler{Fn: fn})
}

type CustomAttrEventUpdatedEventHandler struct {
	Fn func(*core.Context, *CustomAttrEventUpdatedEvent) error
}

func (h *CustomAttrEventUpdatedEventHandler) GetEvent() interface{} {
	return &CustomAttrEventUpdatedEvent{}
}

func (h *CustomAttrEventUpdatedEventHandler) Handle(ctx *core.Context, event interface{}) error {
	return h.Fn(ctx, event.(*CustomAttrEventUpdatedEvent))
}

func SetCustomAttrEventUpdatedEventHandler(conf *config.Config, fn func(ctx *core.Context, event *CustomAttrEventUpdatedEvent) error) {
	event.SetTypeHandler(conf, "contact.custom_attr_event.updated_v3", &CustomAttrEventUpdatedEventHandler{Fn: fn})
}

type EmployeeTypeEnumActivedEventHandler struct {
	Fn func(*core.Context, *EmployeeTypeEnumActivedEvent) error
}

func (h *EmployeeTypeEnumActivedEventHandler) GetEvent() interface{} {
	return &EmployeeTypeEnumActivedEvent{}
}

func (h *EmployeeTypeEnumActivedEventHandler) Handle(ctx *core.Context, event interface{}) error {
	return h.Fn(ctx, event.(*EmployeeTypeEnumActivedEvent))
}

func SetEmployeeTypeEnumActivedEventHandler(conf *config.Config, fn func(ctx *core.Context, event *EmployeeTypeEnumActivedEvent) error) {
	event.SetTypeHandler(conf, "contact.employee_type_enum.actived_v3", &EmployeeTypeEnumActivedEventHandler{Fn: fn})
}

type EmployeeTypeEnumCreatedEventHandler struct {
	Fn func(*core.Context, *EmployeeTypeEnumCreatedEvent) error
}

func (h *EmployeeTypeEnumCreatedEventHandler) GetEvent() interface{} {
	return &EmployeeTypeEnumCreatedEvent{}
}

func (h *EmployeeTypeEnumCreatedEventHandler) Handle(ctx *core.Context, event interface{}) error {
	return h.Fn(ctx, event.(*EmployeeTypeEnumCreatedEvent))
}

func SetEmployeeTypeEnumCreatedEventHandler(conf *config.Config, fn func(ctx *core.Context, event *EmployeeTypeEnumCreatedEvent) error) {
	event.SetTypeHandler(conf, "contact.employee_type_enum.created_v3", &EmployeeTypeEnumCreatedEventHandler{Fn: fn})
}

type EmployeeTypeEnumDeactivatedEventHandler struct {
	Fn func(*core.Context, *EmployeeTypeEnumDeactivatedEvent) error
}

func (h *EmployeeTypeEnumDeactivatedEventHandler) GetEvent() interface{} {
	return &EmployeeTypeEnumDeactivatedEvent{}
}

func (h *EmployeeTypeEnumDeactivatedEventHandler) Handle(ctx *core.Context, event interface{}) error {
	return h.Fn(ctx, event.(*EmployeeTypeEnumDeactivatedEvent))
}

func SetEmployeeTypeEnumDeactivatedEventHandler(conf *config.Config, fn func(ctx *core.Context, event *EmployeeTypeEnumDeactivatedEvent) error) {
	event.SetTypeHandler(conf, "contact.employee_type_enum.deactivated_v3", &EmployeeTypeEnumDeactivatedEventHandler{Fn: fn})
}

type EmployeeTypeEnumDeletedEventHandler struct {
	Fn func(*core.Context, *EmployeeTypeEnumDeletedEvent) error
}

func (h *EmployeeTypeEnumDeletedEventHandler) GetEvent() interface{} {
	return &EmployeeTypeEnumDeletedEvent{}
}

func (h *EmployeeTypeEnumDeletedEventHandler) Handle(ctx *core.Context, event interface{}) error {
	return h.Fn(ctx, event.(*EmployeeTypeEnumDeletedEvent))
}

func SetEmployeeTypeEnumDeletedEventHandler(conf *config.Config, fn func(ctx *core.Context, event *EmployeeTypeEnumDeletedEvent) error) {
	event.SetTypeHandler(conf, "contact.employee_type_enum.deleted_v3", &EmployeeTypeEnumDeletedEventHandler{Fn: fn})
}

type EmployeeTypeEnumUpdatedEventHandler struct {
	Fn func(*core.Context, *EmployeeTypeEnumUpdatedEvent) error
}

func (h *EmployeeTypeEnumUpdatedEventHandler) GetEvent() interface{} {
	return &EmployeeTypeEnumUpdatedEvent{}
}

func (h *EmployeeTypeEnumUpdatedEventHandler) Handle(ctx *core.Context, event interface{}) error {
	return h.Fn(ctx, event.(*EmployeeTypeEnumUpdatedEvent))
}

func SetEmployeeTypeEnumUpdatedEventHandler(conf *config.Config, fn func(ctx *core.Context, event *EmployeeTypeEnumUpdatedEvent) error) {
	event.SetTypeHandler(conf, "contact.employee_type_enum.updated_v3", &EmployeeTypeEnumUpdatedEventHandler{Fn: fn})
}
