package main

import (
	"encoding/json"
	"time"
)

type Team struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	ID        string    `json:"id"`
	FullName  string    `json:"full_name"`
	Email     *string   `json:"email"`
	Phone     *string   `json:"phone"`
	Team      *Team     `json:"team"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Industry struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

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

type Product struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Pipeline struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	DisplayOrder int       `json:"display_order"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

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

type Contact struct {
	ID             string          `json:"id"`
	FullName       string          `json:"full_name"`
	JobTitle       *string         `json:"job_title"`
	Emails         json.RawMessage `json:"emails"`
	Phones         json.RawMessage `json:"phones"`
	SocialProfiles json.RawMessage `json:"social_profiles"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
}

type Organization struct {
	ID          string          `json:"id"`
	Title       string          `json:"title"`
	Description *string         `json:"description"`
	Website     *string         `json:"website"`
	Address     json.RawMessage `json:"address"`
	Owner       *User           `json:"owner"`
	Industries  []Industry      `json:"industries"`
	Followers   []User          `json:"followers"`
	Contacts    []Contact       `json:"contacts"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

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
