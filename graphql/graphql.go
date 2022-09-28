package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/sirupsen/logrus"
)

type QueryObject struct {
	graphql.Object
}

type MutationObject struct {
	graphql.Object
}

type SubscriptionObject struct {
	graphql.Object
}

func NewQuery() *QueryObject {
	obj := graphql.NewObject(graphql.ObjectConfig{
		Name:        "Query",
		Description: "The application's root query object",
		Fields:      graphql.Fields{},
	})

	return &QueryObject{
		*obj,
	}
}

func NewMutation() *MutationObject {
	obj := graphql.NewObject(graphql.ObjectConfig{
		Name:        "Mutation",
		Description: "The application's root mutation object",
		Fields:      graphql.Fields{},
	})

	return &MutationObject{
		*obj,
	}
}

func NewSubscribtion() *SubscriptionObject {
	obj := graphql.NewObject(graphql.ObjectConfig{
		Name:        "Subscription",
		Description: "The application's root subscription object",
		Fields:      graphql.Fields{},
	})

	return &SubscriptionObject{
		*obj,
	}
}

type NewSchemaOptions struct {
	Query        *QueryObject
	Mutation     *MutationObject
	Subscription *SubscriptionObject
	Logger       *logrus.Logger
}

func NewSchema(options *NewSchemaOptions) (*graphql.Schema, error) {
	query := options.Query
	mutation := options.Mutation
	subscription := options.Subscription

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: func() *graphql.Object {
			if len(query.Fields()) == 0 {
				return nil
			}
			return &query.Object
		}(),
		Mutation: func() *graphql.Object {
			if len(mutation.Fields()) == 0 {
				return nil
			}
			return &mutation.Object
		}(),
		Subscription: func() *graphql.Object {
			if len(subscription.Fields()) == 0 {
				return nil
			}
			return &subscription.Object
		}(),
	})
	return &schema, err
}
