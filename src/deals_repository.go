package main

import "fmt"

func queryTeams() (map[string]Team, error) {
	rows, err := db.Query("SELECT id, title, created_at, updated_at FROM sales.crm_teams")
	if err != nil {
		return nil, fmt.Errorf("queryTeams: %w", err)
	}
	defer rows.Close()
	teams := map[string]Team{}
	for rows.Next() {
		var t Team
		if err := rows.Scan(&t.ID, &t.Title, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, fmt.Errorf("queryTeams: %w", err)
		}
		teams[t.ID] = t
	}
	return teams, rows.Err()
}

func queryUserTeams() (map[string]string, error) {
	rows, err := db.Query("SELECT user_id, team_id FROM sales.crm_teams_users")
	if err != nil {
		return nil, fmt.Errorf("queryUserTeams: %w", err)
	}
	defer rows.Close()
	m := map[string]string{}
	for rows.Next() {
		var userID, teamID string
		if err := rows.Scan(&userID, &teamID); err != nil {
			return nil, fmt.Errorf("queryUserTeams: %w", err)
		}
		if _, exists := m[userID]; !exists {
			m[userID] = teamID
		}
	}
	return m, rows.Err()
}

func queryUsers(teams map[string]Team, userTeams map[string]string) (map[string]User, error) {
	rows, err := db.Query("SELECT id, full_name, email, phone, created_at, updated_at FROM sales.crm_users")
	if err != nil {
		return nil, fmt.Errorf("queryUsers: %w", err)
	}
	defer rows.Close()
	users := map[string]User{}
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.FullName, &u.Email, &u.Phone, &u.CreatedAt, &u.UpdatedAt); err != nil {
			return nil, fmt.Errorf("queryUsers: %w", err)
		}
		if teamID, ok := userTeams[u.ID]; ok {
			if team, ok := teams[teamID]; ok {
				u.Team = &team
			}
		}
		users[u.ID] = u
	}
	return users, rows.Err()
}

func queryIndustries() (map[string]Industry, error) {
	rows, err := db.Query("SELECT id, title, created_at, updated_at FROM sales.crm_industries")
	if err != nil {
		return nil, fmt.Errorf("queryIndustries: %w", err)
	}
	defer rows.Close()
	industries := map[string]Industry{}
	for rows.Next() {
		var i Industry
		if err := rows.Scan(&i.ID, &i.Title, &i.CreatedAt, &i.UpdatedAt); err != nil {
			return nil, fmt.Errorf("queryIndustries: %w", err)
		}
		industries[i.ID] = i
	}
	return industries, rows.Err()
}

func querySources() (map[string]Source, error) {
	rows, err := db.Query("SELECT id, title, description, created_at, updated_at FROM sales.crm_sources")
	if err != nil {
		return nil, fmt.Errorf("querySources: %w", err)
	}
	defer rows.Close()
	sources := map[string]Source{}
	for rows.Next() {
		var s Source
		if err := rows.Scan(&s.ID, &s.Title, &s.Description, &s.CreatedAt, &s.UpdatedAt); err != nil {
			return nil, fmt.Errorf("querySources: %w", err)
		}
		sources[s.ID] = s
	}
	return sources, rows.Err()
}

func queryCampaigns() (map[string]Campaign, error) {
	rows, err := db.Query("SELECT id, title, description, created_at, updated_at FROM sales.crm_campaigns")
	if err != nil {
		return nil, fmt.Errorf("queryCampaigns: %w", err)
	}
	defer rows.Close()
	campaigns := map[string]Campaign{}
	for rows.Next() {
		var c Campaign
		if err := rows.Scan(&c.ID, &c.Title, &c.Description, &c.CreatedAt, &c.UpdatedAt); err != nil {
			return nil, fmt.Errorf("queryCampaigns: %w", err)
		}
		campaigns[c.ID] = c
	}
	return campaigns, rows.Err()
}

func queryLossReasons() (map[string]LossReason, error) {
	rows, err := db.Query("SELECT id, reason, created_at, updated_at FROM sales.crm_loss_reasons")
	if err != nil {
		return nil, fmt.Errorf("queryLossReasons: %w", err)
	}
	defer rows.Close()
	lossReasons := map[string]LossReason{}
	for rows.Next() {
		var lr LossReason
		if err := rows.Scan(&lr.ID, &lr.Reason, &lr.CreatedAt, &lr.UpdatedAt); err != nil {
			return nil, fmt.Errorf("queryLossReasons: %w", err)
		}
		lossReasons[lr.ID] = lr
	}
	return lossReasons, rows.Err()
}

func queryProducts() (map[string]Product, error) {
	rows, err := db.Query("SELECT id, title, description, price, created_at, updated_at FROM sales.crm_products")
	if err != nil {
		return nil, fmt.Errorf("queryProducts: %w", err)
	}
	defer rows.Close()
	products := map[string]Product{}
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Title, &p.Description, &p.Price, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, fmt.Errorf("queryProducts: %w", err)
		}
		products[p.ID] = p
	}
	return products, rows.Err()
}

func queryPipelines() (map[string]Pipeline, error) {
	rows, err := db.Query("SELECT id, title, display_order, created_at, updated_at FROM sales.crm_pipelines")
	if err != nil {
		return nil, fmt.Errorf("queryPipelines: %w", err)
	}
	defer rows.Close()
	pipelines := map[string]Pipeline{}
	for rows.Next() {
		var p Pipeline
		if err := rows.Scan(&p.ID, &p.Title, &p.DisplayOrder, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, fmt.Errorf("queryPipelines: %w", err)
		}
		pipelines[p.ID] = p
	}
	return pipelines, rows.Err()
}

func queryPipelineStages(pipelines map[string]Pipeline) (map[string]PipelineStage, error) {
	rows, err := db.Query("SELECT id, pipeline_id, title, description, objective, display_order, created_at, updated_at FROM sales.crm_pipeline_stages")
	if err != nil {
		return nil, fmt.Errorf("queryPipelineStages: %w", err)
	}
	defer rows.Close()
	stages := map[string]PipelineStage{}
	for rows.Next() {
		var s PipelineStage
		var pipelineID string
		if err := rows.Scan(&s.ID, &pipelineID, &s.Title, &s.Description, &s.Objective, &s.DisplayOrder, &s.CreatedAt, &s.UpdatedAt); err != nil {
			return nil, fmt.Errorf("queryPipelineStages: %w", err)
		}
		s.Pipeline = pipelines[pipelineID]
		stages[s.ID] = s
	}
	return stages, rows.Err()
}

func queryContacts() (map[string]Contact, map[string][]string, error) {
	rows, err := db.Query("SELECT id, organization_id, full_name, job_title, emails, phones, social_profiles, created_at, updated_at FROM sales.crm_contacts")
	if err != nil {
		return nil, nil, fmt.Errorf("queryContacts: %w", err)
	}
	defer rows.Close()
	contacts := map[string]Contact{}
	orgContacts := map[string][]string{}
	for rows.Next() {
		var c Contact
		var orgID *string
		var emails, phones, socialProfiles []byte
		if err := rows.Scan(&c.ID, &orgID, &c.FullName, &c.JobTitle, &emails, &phones, &socialProfiles, &c.CreatedAt, &c.UpdatedAt); err != nil {
			return nil, nil, fmt.Errorf("queryContacts: %w", err)
		}
		c.Emails, c.Phones, c.SocialProfiles = emails, phones, socialProfiles
		contacts[c.ID] = c
		if orgID != nil {
			orgContacts[*orgID] = append(orgContacts[*orgID], c.ID)
		}
	}
	return contacts, orgContacts, rows.Err()
}

func queryOrganizations(
	users map[string]User,
	industries map[string]Industry,
	orgIndustries map[string][]string,
	orgFollowers map[string][]string,
	contacts map[string]Contact,
	orgContacts map[string][]string,
) (map[string]Organization, error) {
	rows, err := db.Query("SELECT id, owner_id, title, description, website, address, created_at, updated_at FROM sales.crm_organizations")
	if err != nil {
		return nil, fmt.Errorf("queryOrganizations: %w", err)
	}
	defer rows.Close()
	orgs := map[string]Organization{}
	for rows.Next() {
		var o Organization
		var ownerID *string
		var addr []byte
		if err := rows.Scan(&o.ID, &ownerID, &o.Title, &o.Description, &o.Website, &addr, &o.CreatedAt, &o.UpdatedAt); err != nil {
			return nil, fmt.Errorf("queryOrganizations: %w", err)
		}
		o.Address = addr
		if ownerID != nil {
			if user, ok := users[*ownerID]; ok {
				o.Owner = &user
			}
		}
		for _, indID := range orgIndustries[o.ID] {
			if ind, ok := industries[indID]; ok {
				o.Industries = append(o.Industries, ind)
			}
		}
		for _, userID := range orgFollowers[o.ID] {
			if user, ok := users[userID]; ok {
				o.Followers = append(o.Followers, user)
			}
		}
		for _, contactID := range orgContacts[o.ID] {
			if contact, ok := contacts[contactID]; ok {
				o.Contacts = append(o.Contacts, contact)
			}
		}
		if o.Industries == nil {
			o.Industries = []Industry{}
		}
		if o.Followers == nil {
			o.Followers = []User{}
		}
		if o.Contacts == nil {
			o.Contacts = []Contact{}
		}
		orgs[o.ID] = o
	}
	return orgs, rows.Err()
}

func queryTasks(users map[string]User, taskAssignees map[string][]string) (map[string][]Task, error) {
	rows, err := db.Query("SELECT id, created_by_id, completed_by_id, deal_id, title, description, task_type, status, due_date, completed_at, created_at, updated_at FROM sales.crm_tasks")
	if err != nil {
		return nil, fmt.Errorf("queryTasks: %w", err)
	}
	defer rows.Close()
	dealTasks := map[string][]Task{}
	for rows.Next() {
		var t Task
		var createdByID string
		var completedByID, dealID *string
		if err := rows.Scan(&t.ID, &createdByID, &completedByID, &dealID, &t.Title, &t.Description, &t.TaskType, &t.Status, &t.DueDate, &t.CompletedAt, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, fmt.Errorf("queryTasks: %w", err)
		}
		t.CreatedBy = users[createdByID]
		if completedByID != nil {
			if user, ok := users[*completedByID]; ok {
				t.CompletedBy = &user
			}
		}
		for _, userID := range taskAssignees[t.ID] {
			if user, ok := users[userID]; ok {
				t.Assignees = append(t.Assignees, user)
			}
		}
		if t.Assignees == nil {
			t.Assignees = []User{}
		}
		if dealID != nil {
			dealTasks[*dealID] = append(dealTasks[*dealID], t)
		}
	}
	return dealTasks, rows.Err()
}

func queryJoinList(table, keyCol, valCol string) (map[string][]string, error) {
	rows, err := db.Query(fmt.Sprintf("SELECT %s, %s FROM sales.%s", keyCol, valCol, table))
	if err != nil {
		return nil, fmt.Errorf("queryJoinList %s: %w", table, err)
	}
	defer rows.Close()
	m := map[string][]string{}
	for rows.Next() {
		var k, v string
		if err := rows.Scan(&k, &v); err != nil {
			return nil, fmt.Errorf("queryJoinList %s: %w", table, err)
		}
		m[k] = append(m[k], v)
	}
	return m, rows.Err()
}

func GetDeals() ([]Deal, error) {
	teams, err := queryTeams()
	if err != nil {
		return nil, err
	}
	userTeams, err := queryUserTeams()
	if err != nil {
		return nil, err
	}
	users, err := queryUsers(teams, userTeams)
	if err != nil {
		return nil, err
	}
	industries, err := queryIndustries()
	if err != nil {
		return nil, err
	}
	orgIndustries, err := queryJoinList("crm_organizations_industries", "organization_id", "industry_id")
	if err != nil {
		return nil, err
	}
	orgFollowers, err := queryJoinList("crm_organizations_users", "organization_id", "user_id")
	if err != nil {
		return nil, err
	}
	contacts, orgContacts, err := queryContacts()
	if err != nil {
		return nil, err
	}
	orgs, err := queryOrganizations(users, industries, orgIndustries, orgFollowers, contacts, orgContacts)
	if err != nil {
		return nil, err
	}
	sources, err := querySources()
	if err != nil {
		return nil, err
	}
	campaigns, err := queryCampaigns()
	if err != nil {
		return nil, err
	}
	lossReasons, err := queryLossReasons()
	if err != nil {
		return nil, err
	}
	products, err := queryProducts()
	if err != nil {
		return nil, err
	}
	pipelines, err := queryPipelines()
	if err != nil {
		return nil, err
	}
	stages, err := queryPipelineStages(pipelines)
	if err != nil {
		return nil, err
	}
	taskAssignees, err := queryJoinList("crm_tasks_users", "task_id", "user_id")
	if err != nil {
		return nil, err
	}
	dealTasks, err := queryTasks(users, taskAssignees)
	if err != nil {
		return nil, err
	}
	dealProducts, err := queryJoinList("crm_deals_products", "deal_id", "product_id")
	if err != nil {
		return nil, err
	}
	dealContacts, err := queryJoinList("crm_deals_contacts", "deal_id", "contact_id")
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(`
		SELECT id, stage_id, owner_id, source_id, campaign_id, loss_reason_id,
		       organization_id, title, amount, expected_close_date, rating, status,
		       closed_at, created_at, updated_at
		FROM sales.crm_deals
	`)
	if err != nil {
		return nil, fmt.Errorf("GetDeals: %w", err)
	}
	defer rows.Close()

	var deals []Deal
	for rows.Next() {
		var d Deal
		var stageID string
		var ownerID, sourceID, campaignID, lossReasonID, organizationID *string
		if err := rows.Scan(
			&d.ID, &stageID, &ownerID, &sourceID, &campaignID, &lossReasonID,
			&organizationID, &d.Title, &d.Amount, &d.ExpectedCloseDate, &d.Rating,
			&d.Status, &d.ClosedAt, &d.CreatedAt, &d.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("GetDeals: %w", err)
		}
		d.Stage = stages[stageID]
		if ownerID != nil {
			if u, ok := users[*ownerID]; ok {
				d.Owner = &u
			}
		}
		if sourceID != nil {
			if s, ok := sources[*sourceID]; ok {
				d.Source = &s
			}
		}
		if campaignID != nil {
			if c, ok := campaigns[*campaignID]; ok {
				d.Campaign = &c
			}
		}
		if lossReasonID != nil {
			if lr, ok := lossReasons[*lossReasonID]; ok {
				d.LossReason = &lr
			}
		}
		if organizationID != nil {
			if org, ok := orgs[*organizationID]; ok {
				d.Organization = &org
			}
		}
		for _, productID := range dealProducts[d.ID] {
			if p, ok := products[productID]; ok {
				d.Products = append(d.Products, p)
			}
		}
		for _, contactID := range dealContacts[d.ID] {
			if c, ok := contacts[contactID]; ok {
				d.Contacts = append(d.Contacts, c)
			}
		}
		d.Tasks = dealTasks[d.ID]
		if d.Products == nil {
			d.Products = []Product{}
		}
		if d.Contacts == nil {
			d.Contacts = []Contact{}
		}
		if d.Tasks == nil {
			d.Tasks = []Task{}
		}
		deals = append(deals, d)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetDeals: %w", err)
	}
	if deals == nil {
		deals = []Deal{}
	}
	return deals, nil
}
