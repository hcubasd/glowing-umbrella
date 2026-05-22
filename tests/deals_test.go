package tests

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testDB *sql.DB

var now = time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

func TestMain(m *testing.M) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PGHOST"), os.Getenv("PGPORT"), os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"), os.Getenv("PGDATABASE"),
	)
	var err error
	testDB, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	if err = testDB.Ping(); err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

type Team struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type User struct {
	ID       string  `json:"id"`
	FullName string  `json:"full_name"`
	Email    *string `json:"email"`
	Phone    *string `json:"phone"`
	Team     *Team   `json:"team"`
}

type Industry struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type Source struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
}

type Campaign struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
}

type LossReason struct {
	ID     string `json:"id"`
	Reason string `json:"reason"`
}

type Product struct {
	ID    string  `json:"id"`
	Title string  `json:"title"`
	Price float64 `json:"price"`
}

type Pipeline struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	DisplayOrder int    `json:"display_order"`
}

type PipelineStage struct {
	ID           string   `json:"id"`
	Title        string   `json:"title"`
	DisplayOrder int      `json:"display_order"`
	Pipeline     Pipeline `json:"pipeline"`
}

type Contact struct {
	ID       string  `json:"id"`
	FullName string  `json:"full_name"`
	JobTitle *string `json:"job_title"`
}

type Organization struct {
	ID         string     `json:"id"`
	Title      string     `json:"title"`
	Owner      *User      `json:"owner"`
	Industries []Industry `json:"industries"`
	Followers  []User     `json:"followers"`
	Contacts   []Contact  `json:"contacts"`
}

type Task struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	TaskType  string `json:"task_type"`
	Status    string `json:"status"`
	CreatedBy User   `json:"created_by"`
	Assignees []User `json:"assignees"`
}

type Deal struct {
	ID           string        `json:"id"`
	Title        string        `json:"title"`
	Stage        PipelineStage `json:"stage"`
	Owner        *User         `json:"owner"`
	Source       *Source       `json:"source"`
	Campaign     *Campaign     `json:"campaign"`
	LossReason   *LossReason   `json:"loss_reason"`
	Organization *Organization `json:"organization"`
	Amount       *float64      `json:"amount"`
	Rating       *int          `json:"rating"`
	Status       string        `json:"status"`
	Contacts     []Contact     `json:"contacts"`
	Products     []Product     `json:"products"`
	Tasks        []Task        `json:"tasks"`
}

const truncateSQL = `TRUNCATE
	sales.crm_tasks_users, sales.crm_deals_contacts, sales.crm_deals_products,
	sales.crm_organizations_industries, sales.crm_organizations_users, sales.crm_teams_users,
	sales.crm_tasks, sales.crm_deals, sales.crm_contacts, sales.crm_organizations,
	sales.crm_pipeline_stages, sales.crm_users, sales.crm_teams, sales.crm_pipelines,
	sales.crm_campaigns, sales.crm_sources, sales.crm_loss_reasons,
	sales.crm_industries, sales.crm_products CASCADE`

func truncateAll(t *testing.T) {
	t.Helper()
	_, err := testDB.Exec(truncateSQL)
	require.NoError(t, err)
}

func insertFixtures(t *testing.T) {
	t.Helper()
	exec := func(query string, args ...any) {
		t.Helper()
		_, err := testDB.Exec(query, args...)
		require.NoError(t, err)
	}

	exec(`INSERT INTO sales.crm_teams (id, title, created_at, updated_at) VALUES ('t1', 'Sales', $1, $1)`, now)
	exec(`INSERT INTO sales.crm_users (id, full_name, email, phone, created_at, updated_at) VALUES ('u1', 'Alice', 'alice@example.com', NULL, $1, $1)`, now)
	exec(`INSERT INTO sales.crm_teams_users (team_id, user_id) VALUES ('t1', 'u1')`)
	exec(`INSERT INTO sales.crm_pipelines (id, title, display_order, created_at, updated_at) VALUES ('pl1', 'Main', 1, $1, $1)`, now)
	exec(`INSERT INTO sales.crm_pipeline_stages (id, pipeline_id, title, description, objective, display_order, created_at, updated_at) VALUES ('ps1', 'pl1', 'Proposal', NULL, NULL, 1, $1, $1)`, now)
	exec(`INSERT INTO sales.crm_industries (id, title, created_at, updated_at) VALUES ('ind1', 'Tech', $1, $1)`, now)
	exec(`INSERT INTO sales.crm_products (id, title, description, price, created_at, updated_at) VALUES ('pr1', 'Widget', NULL, 99.99, $1, $1)`, now)
	exec(`INSERT INTO sales.crm_sources (id, title, description, created_at, updated_at) VALUES ('src1', 'Web', NULL, $1, $1)`, now)
	exec(`INSERT INTO sales.crm_campaigns (id, title, description, created_at, updated_at) VALUES ('cmp1', 'Q1', NULL, $1, $1)`, now)
	exec(`INSERT INTO sales.crm_loss_reasons (id, reason, created_at, updated_at) VALUES ('lr1', 'Price', $1, $1)`, now)
	exec(`INSERT INTO sales.crm_organizations (id, owner_id, title, description, website, address, created_at, updated_at) VALUES ('org1', 'u1', 'Acme', NULL, NULL, NULL, $1, $1)`, now)
	exec(`INSERT INTO sales.crm_contacts (id, organization_id, full_name, job_title, emails, phones, social_profiles, created_at, updated_at) VALUES ('ct1', 'org1', 'Bob', 'CEO', '[]', '[]', '[]', $1, $1)`, now)
	exec(`INSERT INTO sales.crm_organizations_industries (organization_id, industry_id) VALUES ('org1', 'ind1')`)
	exec(`INSERT INTO sales.crm_organizations_users (organization_id, user_id) VALUES ('org1', 'u1')`)
	exec(`INSERT INTO sales.crm_deals (id, stage_id, owner_id, source_id, campaign_id, loss_reason_id, organization_id, title, amount, expected_close_date, rating, status, closed_at, created_at, updated_at) VALUES ('d1', 'ps1', 'u1', 'src1', 'cmp1', 'lr1', 'org1', 'Big Deal', 10000.00, '2024-03-31', 3, 'ongoing', NULL, $1, $1)`, now)
	exec(`INSERT INTO sales.crm_deals_contacts (deal_id, contact_id) VALUES ('d1', 'ct1')`)
	exec(`INSERT INTO sales.crm_deals_products (deal_id, product_id) VALUES ('d1', 'pr1')`)
	exec(`INSERT INTO sales.crm_tasks (id, created_by_id, completed_by_id, deal_id, title, description, task_type, status, due_date, completed_at, created_at, updated_at) VALUES ('tk1', 'u1', NULL, 'd1', 'Follow up', NULL, 'call', 'pending', NULL, NULL, $1, $1)`, now)
	exec(`INSERT INTO sales.crm_tasks_users (task_id, user_id) VALUES ('tk1', 'u1')`)
}

func TestGetDeals(t *testing.T) {
	truncateAll(t)
	insertFixtures(t)

	resp, err := http.Get("http://localhost:8080/deals")
	require.NoError(t, err)
	defer resp.Body.Close()

	require.Equal(t, http.StatusOK, resp.StatusCode)

	var deals []Deal
	require.NoError(t, json.NewDecoder(resp.Body).Decode(&deals))
	require.Len(t, deals, 1)

	d := deals[0]
	assert.Equal(t, "d1", d.ID)
	assert.Equal(t, "Big Deal", d.Title)
	assert.Equal(t, "ongoing", d.Status)
	require.NotNil(t, d.Amount)
	assert.Equal(t, 10000.0, *d.Amount)
	require.NotNil(t, d.Rating)
	assert.Equal(t, 3, *d.Rating)

	assert.Equal(t, "ps1", d.Stage.ID)
	assert.Equal(t, "Proposal", d.Stage.Title)
	assert.Equal(t, "pl1", d.Stage.Pipeline.ID)
	assert.Equal(t, "Main", d.Stage.Pipeline.Title)

	require.NotNil(t, d.Owner)
	assert.Equal(t, "u1", d.Owner.ID)
	assert.Equal(t, "Alice", d.Owner.FullName)
	require.NotNil(t, d.Owner.Team)
	assert.Equal(t, "t1", d.Owner.Team.ID)

	require.NotNil(t, d.Source)
	assert.Equal(t, "src1", d.Source.ID)
	require.NotNil(t, d.Campaign)
	assert.Equal(t, "cmp1", d.Campaign.ID)
	require.NotNil(t, d.LossReason)
	assert.Equal(t, "lr1", d.LossReason.ID)

	require.NotNil(t, d.Organization)
	assert.Equal(t, "org1", d.Organization.ID)
	assert.Equal(t, "Acme", d.Organization.Title)
	require.NotNil(t, d.Organization.Owner)
	assert.Equal(t, "u1", d.Organization.Owner.ID)
	assert.Len(t, d.Organization.Industries, 1)
	assert.Equal(t, "ind1", d.Organization.Industries[0].ID)
	assert.Len(t, d.Organization.Followers, 1)
	assert.Equal(t, "u1", d.Organization.Followers[0].ID)
	assert.Len(t, d.Organization.Contacts, 1)
	assert.Equal(t, "ct1", d.Organization.Contacts[0].ID)

	require.Len(t, d.Contacts, 1)
	assert.Equal(t, "ct1", d.Contacts[0].ID)
	require.Len(t, d.Products, 1)
	assert.Equal(t, "pr1", d.Products[0].ID)
	assert.Equal(t, 99.99, d.Products[0].Price)

	require.Len(t, d.Tasks, 1)
	tk := d.Tasks[0]
	assert.Equal(t, "tk1", tk.ID)
	assert.Equal(t, "Follow up", tk.Title)
	assert.Equal(t, "call", tk.TaskType)
	assert.Equal(t, "pending", tk.Status)
	assert.Equal(t, "u1", tk.CreatedBy.ID)
	require.Len(t, tk.Assignees, 1)
	assert.Equal(t, "u1", tk.Assignees[0].ID)
}
