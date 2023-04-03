
package model

import (
    "time"

    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Expense struct {
    ID          primitive.ObjectID `bson:"_id,omitempty"`
    Title       string             `bson:"title,omitempty" validate:"required,max=50"`
    Amount      float64            `bson:"amount,omitempty" validate:"required,max=20"`
    Type        string             `bson:"type,omitempty" default:"expense"`
    Date        string              `bson:"date,omitempty" validate:"required"`
    Category    string             `bson:"category,omitempty" validate:"required"`
    Description string             `bson:"description,omitempty" validate:"required,max=20"`
    CreatedAt   time.Time          `bson:"createdAt,omitempty"`
    UpdatedAt   time.Time          `bson:"updatedAt,omitempty"`
}
