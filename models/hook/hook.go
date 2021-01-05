// Code generated (@generated) by entc, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/adnaan/authzen/models"
)

// The AccountFunc type is an adapter to allow the use of ordinary
// function as Account mutator.
type AccountFunc func(context.Context, *models.AccountMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f AccountFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.AccountMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.AccountMutation", m)
	}
	return f(ctx, mv)
}

// The AccountRoleFunc type is an adapter to allow the use of ordinary
// function as AccountRole mutator.
type AccountRoleFunc func(context.Context, *models.AccountRoleMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f AccountRoleFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.AccountRoleMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.AccountRoleMutation", m)
	}
	return f(ctx, mv)
}

// The GroupFunc type is an adapter to allow the use of ordinary
// function as Group mutator.
type GroupFunc func(context.Context, *models.GroupMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f GroupFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.GroupMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.GroupMutation", m)
	}
	return f(ctx, mv)
}

// The GroupRoleFunc type is an adapter to allow the use of ordinary
// function as GroupRole mutator.
type GroupRoleFunc func(context.Context, *models.GroupRoleMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f GroupRoleFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.GroupRoleMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.GroupRoleMutation", m)
	}
	return f(ctx, mv)
}

// The PermissionFunc type is an adapter to allow the use of ordinary
// function as Permission mutator.
type PermissionFunc func(context.Context, *models.PermissionMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f PermissionFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.PermissionMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.PermissionMutation", m)
	}
	return f(ctx, mv)
}

// The SessionFunc type is an adapter to allow the use of ordinary
// function as Session mutator.
type SessionFunc func(context.Context, *models.SessionMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f SessionFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.SessionMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.SessionMutation", m)
	}
	return f(ctx, mv)
}

// The WorkspaceFunc type is an adapter to allow the use of ordinary
// function as Workspace mutator.
type WorkspaceFunc func(context.Context, *models.WorkspaceMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f WorkspaceFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.WorkspaceMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.WorkspaceMutation", m)
	}
	return f(ctx, mv)
}

// The WorkspaceInvitationFunc type is an adapter to allow the use of ordinary
// function as WorkspaceInvitation mutator.
type WorkspaceInvitationFunc func(context.Context, *models.WorkspaceInvitationMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f WorkspaceInvitationFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.WorkspaceInvitationMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.WorkspaceInvitationMutation", m)
	}
	return f(ctx, mv)
}

// The WorkspaceRoleFunc type is an adapter to allow the use of ordinary
// function as WorkspaceRole mutator.
type WorkspaceRoleFunc func(context.Context, *models.WorkspaceRoleMutation) (models.Value, error)

// Mutate calls f(ctx, m).
func (f WorkspaceRoleFunc) Mutate(ctx context.Context, m models.Mutation) (models.Value, error) {
	mv, ok := m.(*models.WorkspaceRoleMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *models.WorkspaceRoleMutation", m)
	}
	return f(ctx, mv)
}

// Condition is a hook condition function.
type Condition func(context.Context, models.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m models.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m models.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m models.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op models.Op) Condition {
	return func(_ context.Context, m models.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m models.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m models.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m models.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
//
func If(hk models.Hook, cond Condition) models.Hook {
	return func(next models.Mutator) models.Mutator {
		return models.MutateFunc(func(ctx context.Context, m models.Mutation) (models.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, models.Delete|models.Create)
//
func On(hk models.Hook, op models.Op) models.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, models.Update|models.UpdateOne)
//
func Unless(hk models.Hook, op models.Op) models.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) models.Hook {
	return func(models.Mutator) models.Mutator {
		return models.MutateFunc(func(context.Context, models.Mutation) (models.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []models.Hook {
//		return []models.Hook{
//			Reject(models.Delete|models.Update),
//		}
//	}
//
func Reject(op models.Op) models.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []models.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...models.Hook) Chain {
	return Chain{append([]models.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() models.Hook {
	return func(mutator models.Mutator) models.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...models.Hook) Chain {
	newHooks := make([]models.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
