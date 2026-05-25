package main

import (
	"encoding/json"
	"time"
)

// Team represents a CRM team (group of users). It is included in API responses to indicate
// the team a user belongs to.
type Team struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// User represents a CRM user (person). Email and Phone are optional; Team may be nil.
type User struct {
	ID        string    `json:"id"`
	FullName  string    `json:"full_name"`
	Email     *string   `json:"email"`
	Phone     *string   `json:"phone"`
	Team      *Team     `json:"team"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Industry represents a business sector classification for organizations.
type Industry struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Source represents the origin of a deal (e.g., Web, Referral). Description is optional.
type Source struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Campaign struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type LossReason struct {
	ID        string    `json:"id"`
	Reason    string    `json:"reason"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Product represents a sellable item associated with deals.
// Price is a decimal value expressed as float64 here for JSON serialization.
// @name Product
type Product struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Pipeline groups pipeline stages for deal progression.
// @name Pipeline
type Pipeline struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	DisplayOrder int       `json:"display_order"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// PipelineStage represents a stage within a sales pipeline. Pipeline is embedded to
// provide the parent pipeline metadata.
// @name PipelineStage
type PipelineStage struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	Description  *string   `json:"description"`
	Objective    *string   `json:"objective"`
	DisplayOrder int       `json:"display_order"`
	Pipeline     Pipeline  `json:"pipeline"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// Contact represents an individual contact record. Emails/Phones/SocialProfiles are
// stored as JSON blobs and surfaced as opaque objects in the API.
// @name Contact
type Contact struct {
	ID             string          `json:"id"`
	FullName       string          `json:"full_name"`
	JobTitle       *string         `json:"job_title"`
	Emails         json.RawMessage `json:"emails" swaggertype:"object"`
	Phones         json.RawMessage `json:"phones" swaggertype:"object"`
	SocialProfiles json.RawMessage `json:"social_profiles" swaggertype:"object"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
}

// Organization represents a company or organization. Owner may be nil for anonymous orgs.
// Industries, Followers and Contacts are populated from join tables.
// @name Organization
type Organization struct {
	ID          string          `json:"id"`
	Title       string          `json:"title"`
	Description *string         `json:"description"`
	Website     *string         `json:"website"`
	Address     json.RawMessage `json:"address" swaggertype:"object"`
	Owner       *User           `json:"owner"`
	Industries  []Industry      `json:"industries"`
	Followers   []User          `json:"followers"`
	Contacts    []Contact       `json:"contacts"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

// Task represents an action item associated with a deal. CreatedBy is required; CompletedBy is optional.
// Assignees lists users assigned to the task.
// @name Task
type Task struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	TaskType    string     `json:"task_type"`
	Status      string     `json:"status"`
	DueDate     *time.Time `json:"due_date"`
	CompletedAt *time.Time `json:"completed_at"`
	CreatedBy   User       `json:"created_by"`
	CompletedBy *User      `json:"completed_by"`
	Assignees   []User     `json:"assignees"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// Deal is the aggregate root representing a sales opportunity. It nests related
// objects such as Stage, Organization, Contacts, Products and Tasks.
// @name Deal
type Deal struct {
	ID                string        `json:"id"`
	Title             string        `json:"title"`
	Stage             PipelineStage `json:"stage"`
	Owner             *User         `json:"owner"`
	Source            *Source       `json:"source"`
	Campaign          *Campaign     `json:"campaign"`
	LossReason        *LossReason   `json:"loss_reason"`
	Organization      *Organization `json:"organization"`
	Amount            *float64      `json:"amount"`
	ExpectedCloseDate *time.Time    `json:"expected_close_date"`
	Rating            *int          `json:"rating"`
	Status            string        `json:"status"`
	ClosedAt          *time.Time    `json:"closed_at"`
	Contacts          []Contact     `json:"contacts"`
	Products          []Product     `json:"products"`
	Tasks             []Task        `json:"tasks"`
	CreatedAt         time.Time     `json:"created_at"`
	UpdatedAt         time.Time     `json:"updated_at"`
}
