// Code generated (@generated) by entc, DO NOT EDIT.

package workspaceinvitation

import (
	"time"

	"github.com/adnaan/authzen/models/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/google/uuid"
)

// ID filters vertices based on their identifier.
func ID(id uuid.UUID) predicate.WorkspaceInvitation {
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.WorkspaceInvitation {
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.WorkspaceInvitation {
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.WorkspaceInvitation {
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.WorkspaceInvitation {
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.WorkspaceInvitation {
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.WorkspaceInvitation {
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.WorkspaceInvitation {
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.WorkspaceInvitation {
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// WorkspaceID applies equality check predicate on the "workspace_id" field. It's identical to WorkspaceIDEQ.
func WorkspaceID(v uuid.UUID) predicate.WorkspaceInvitation {
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWorkspaceID), v))
	})
}

// Role applies equality check predicate on the "role" field. It's identical to RoleEQ.
func Role(v string) predicate.WorkspaceInvitation {
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRole), v))
	})
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.WorkspaceInvitation {
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmail), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.WorkspaceInvitation {
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// WorkspaceIDEQ applies the EQ predicate on the "workspace_id" field.
func WorkspaceIDEQ(v uuid.UUID) predicate.WorkspaceInvitation {
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWorkspaceID), v))
	})
}

// WorkspaceIDNEQ applies the NEQ predicate on the "workspace_id" field.
func WorkspaceIDNEQ(v uuid.UUID) predicate.WorkspaceInvitation {
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldWorkspaceID), v))
	})
}

// WorkspaceIDIn applies the In predicate on the "workspace_id" field.
func WorkspaceIDIn(vs ...uuid.UUID) predicate.WorkspaceInvitation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldWorkspaceID), v...))
	})
}

// WorkspaceIDNotIn applies the NotIn predicate on the "workspace_id" field.
func WorkspaceIDNotIn(vs ...uuid.UUID) predicate.WorkspaceInvitation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldWorkspaceID), v...))
	})
}

// WorkspaceIDGT applies the GT predicate on the "workspace_id" field.
func WorkspaceIDGT(v uuid.UUID) predicate.WorkspaceInvitation {
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldWorkspaceID), v))
	})
}

// WorkspaceIDGTE applies the GTE predicate on the "workspace_id" field.
func WorkspaceIDGTE(v uuid.UUID) predicate.WorkspaceInvitation {
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldWorkspaceID), v))
	})
}

// WorkspaceIDLT applies the LT predicate on the "workspace_id" field.
func WorkspaceIDLT(v uuid.UUID) predicate.WorkspaceInvitation {
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldWorkspaceID), v))
	})
}

// WorkspaceIDLTE applies the LTE predicate on the "workspace_id" field.
func WorkspaceIDLTE(v uuid.UUID) predicate.WorkspaceInvitation {
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldWorkspaceID), v))
	})
}

// RoleEQ applies the EQ predicate on the "role" field.
func RoleEQ(v string) predicate.WorkspaceInvitation {
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRole), v))
	})
}

// RoleNEQ applies the NEQ predicate on the "role" field.
func RoleNEQ(v string) predicate.WorkspaceInvitation {
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRole), v))
	})
}

// RoleIn applies the In predicate on the "role" field.
func RoleIn(vs ...string) predicate.WorkspaceInvitation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRole), v...))
	})
}

// RoleNotIn applies the NotIn predicate on the "role" field.
func RoleNotIn(vs ...string) predicate.WorkspaceInvitation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRole), v...))
	})
}

// RoleGT applies the GT predicate on the "role" field.
func RoleGT(v string) predicate.WorkspaceInvitation {
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRole), v))
	})
}

// RoleGTE applies the GTE predicate on the "role" field.
func RoleGTE(v string) predicate.WorkspaceInvitation {
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRole), v))
	})
}

// RoleLT applies the LT predicate on the "role" field.
func RoleLT(v string) predicate.WorkspaceInvitation {
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRole), v))
	})
}

// RoleLTE applies the LTE predicate on the "role" field.
func RoleLTE(v string) predicate.WorkspaceInvitation {
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRole), v))
	})
}

// RoleContains applies the Contains predicate on the "role" field.
func RoleContains(v string) predicate.WorkspaceInvitation {
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRole), v))
	})
}

// RoleHasPrefix applies the HasPrefix predicate on the "role" field.
func RoleHasPrefix(v string) predicate.WorkspaceInvitation {
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRole), v))
	})
}

// RoleHasSuffix applies the HasSuffix predicate on the "role" field.
func RoleHasSuffix(v string) predicate.WorkspaceInvitation {
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRole), v))
	})
}

// RoleEqualFold applies the EqualFold predicate on the "role" field.
func RoleEqualFold(v string) predicate.WorkspaceInvitation {
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRole), v))
	})
}

// RoleContainsFold applies the ContainsFold predicate on the "role" field.
func RoleContainsFold(v string) predicate.WorkspaceInvitation {
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRole), v))
	})
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.WorkspaceInvitation {
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmail), v))
	})
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.WorkspaceInvitation {
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEmail), v))
	})
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.WorkspaceInvitation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEmail), v...))
	})
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.WorkspaceInvitation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEmail), v...))
	})
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.WorkspaceInvitation {
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEmail), v))
	})
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.WorkspaceInvitation {
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEmail), v))
	})
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.WorkspaceInvitation {
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEmail), v))
	})
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.WorkspaceInvitation {
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEmail), v))
	})
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.WorkspaceInvitation {
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEmail), v))
	})
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.WorkspaceInvitation {
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEmail), v))
	})
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.WorkspaceInvitation {
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEmail), v))
	})
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.WorkspaceInvitation {
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEmail), v))
	})
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.WorkspaceInvitation {
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEmail), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.WorkspaceInvitation {
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.WorkspaceInvitation {
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.WorkspaceInvitation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.WorkspaceInvitation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.WorkspaceInvitation {
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.WorkspaceInvitation {
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.WorkspaceInvitation {
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.WorkspaceInvitation {
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.WorkspaceInvitation) predicate.WorkspaceInvitation {
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.WorkspaceInvitation) predicate.WorkspaceInvitation {
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.WorkspaceInvitation) predicate.WorkspaceInvitation {
	return predicate.WorkspaceInvitation(func(s *sql.Selector) {
		p(s.Not())
	})
}
