import { pool } from "./db.js";
import type {
	Campaign,
	Contact,
	Deal,
	Industry,
	LossReason,
	Organization,
	Pipeline,
	PipelineStage,
	Product,
	Source,
	Task,
	Team,
	User,
} from "./schemas.js";

const iso = (d: Date) => d.toISOString();
const isoOrNull = (d: Date | null) => (d === null ? null : d.toISOString());
// expected_close_date comes back as a raw "YYYY-MM-DD" string (see db.ts) - anchor it to
// midnight UTC to match the full-datetime-string contract every other timestamp field uses.
const dateOnlyToIsoOrNull = (d: string | null) => (d === null ? null : `${d}T00:00:00.000Z`);

function lookupAll<T>(ids: string[], byId: Map<string, T>): T[] {
	const out: T[] = [];
	for (const id of ids) {
		const v = byId.get(id);
		if (v) out.push(v);
	}
	return out;
}

async function queryTeams(): Promise<Map<string, Team>> {
	const { rows } = await pool.query<{
		id: string;
		title: string;
		created_at: Date;
		updated_at: Date;
	}>("SELECT id, title, created_at, updated_at FROM sales.crm_teams");
	const teams = new Map<string, Team>();
	for (const r of rows) {
		teams.set(r.id, {
			id: r.id,
			title: r.title,
			created_at: iso(r.created_at),
			updated_at: iso(r.updated_at),
		});
	}
	return teams;
}

async function queryUserTeams(): Promise<Map<string, string>> {
	const { rows } = await pool.query<{ user_id: string; team_id: string }>(
		"SELECT user_id, team_id FROM sales.crm_teams_users",
	);
	const m = new Map<string, string>();
	for (const r of rows) {
		if (!m.has(r.user_id)) m.set(r.user_id, r.team_id);
	}
	return m;
}

async function queryUsers(
	teams: Map<string, Team>,
	userTeams: Map<string, string>,
): Promise<Map<string, User>> {
	const { rows } = await pool.query<{
		id: string;
		full_name: string;
		email: string | null;
		phone: string | null;
		created_at: Date;
		updated_at: Date;
	}>("SELECT id, full_name, email, phone, created_at, updated_at FROM sales.crm_users");
	const users = new Map<string, User>();
	for (const r of rows) {
		const teamId = userTeams.get(r.id);
		users.set(r.id, {
			id: r.id,
			full_name: r.full_name,
			email: r.email,
			phone: r.phone,
			team: teamId ? (teams.get(teamId) ?? null) : null,
			created_at: iso(r.created_at),
			updated_at: iso(r.updated_at),
		});
	}
	return users;
}

async function queryIndustries(): Promise<Map<string, Industry>> {
	const { rows } = await pool.query<{
		id: string;
		title: string;
		created_at: Date;
		updated_at: Date;
	}>("SELECT id, title, created_at, updated_at FROM sales.crm_industries");
	const m = new Map<string, Industry>();
	for (const r of rows) {
		m.set(r.id, {
			id: r.id,
			title: r.title,
			created_at: iso(r.created_at),
			updated_at: iso(r.updated_at),
		});
	}
	return m;
}

async function querySources(): Promise<Map<string, Source>> {
	const { rows } = await pool.query<{
		id: string;
		title: string;
		description: string | null;
		created_at: Date;
		updated_at: Date;
	}>("SELECT id, title, description, created_at, updated_at FROM sales.crm_sources");
	const m = new Map<string, Source>();
	for (const r of rows) {
		m.set(r.id, {
			id: r.id,
			title: r.title,
			description: r.description,
			created_at: iso(r.created_at),
			updated_at: iso(r.updated_at),
		});
	}
	return m;
}

async function queryCampaigns(): Promise<Map<string, Campaign>> {
	const { rows } = await pool.query<{
		id: string;
		title: string;
		description: string | null;
		created_at: Date;
		updated_at: Date;
	}>("SELECT id, title, description, created_at, updated_at FROM sales.crm_campaigns");
	const m = new Map<string, Campaign>();
	for (const r of rows) {
		m.set(r.id, {
			id: r.id,
			title: r.title,
			description: r.description,
			created_at: iso(r.created_at),
			updated_at: iso(r.updated_at),
		});
	}
	return m;
}

async function queryLossReasons(): Promise<Map<string, LossReason>> {
	const { rows } = await pool.query<{
		id: string;
		reason: string;
		created_at: Date;
		updated_at: Date;
	}>("SELECT id, reason, created_at, updated_at FROM sales.crm_loss_reasons");
	const m = new Map<string, LossReason>();
	for (const r of rows) {
		m.set(r.id, {
			id: r.id,
			reason: r.reason,
			created_at: iso(r.created_at),
			updated_at: iso(r.updated_at),
		});
	}
	return m;
}

async function queryProducts(): Promise<Map<string, Product>> {
	const { rows } = await pool.query<{
		id: string;
		title: string;
		description: string | null;
		price: number;
		created_at: Date;
		updated_at: Date;
	}>("SELECT id, title, description, price, created_at, updated_at FROM sales.crm_products");
	const m = new Map<string, Product>();
	for (const r of rows) {
		m.set(r.id, {
			id: r.id,
			title: r.title,
			description: r.description,
			price: r.price,
			created_at: iso(r.created_at),
			updated_at: iso(r.updated_at),
		});
	}
	return m;
}

async function queryPipelines(): Promise<Map<string, Pipeline>> {
	const { rows } = await pool.query<{
		id: string;
		title: string;
		display_order: number;
		created_at: Date;
		updated_at: Date;
	}>("SELECT id, title, display_order, created_at, updated_at FROM sales.crm_pipelines");
	const m = new Map<string, Pipeline>();
	for (const r of rows) {
		m.set(r.id, {
			id: r.id,
			title: r.title,
			display_order: r.display_order,
			created_at: iso(r.created_at),
			updated_at: iso(r.updated_at),
		});
	}
	return m;
}

async function queryPipelineStages(
	pipelines: Map<string, Pipeline>,
): Promise<Map<string, PipelineStage>> {
	const { rows } = await pool.query<{
		id: string;
		pipeline_id: string;
		title: string;
		description: string | null;
		objective: string | null;
		display_order: number;
		created_at: Date;
		updated_at: Date;
	}>(
		`SELECT id, pipeline_id, title, description, objective, display_order, created_at, updated_at
		 FROM sales.crm_pipeline_stages`,
	);
	const m = new Map<string, PipelineStage>();
	for (const r of rows) {
		// pipeline_id is NOT NULL with a FK to crm_pipelines, so this is always present.
		const pipeline = pipelines.get(r.pipeline_id);
		if (!pipeline) continue;
		m.set(r.id, {
			id: r.id,
			title: r.title,
			description: r.description,
			objective: r.objective,
			display_order: r.display_order,
			pipeline,
			created_at: iso(r.created_at),
			updated_at: iso(r.updated_at),
		});
	}
	return m;
}

async function queryContacts(): Promise<{
	contacts: Map<string, Contact>;
	orgContacts: Map<string, string[]>;
}> {
	const { rows } = await pool.query<{
		id: string;
		organization_id: string | null;
		full_name: string;
		job_title: string | null;
		emails: unknown;
		phones: unknown;
		social_profiles: unknown;
		created_at: Date;
		updated_at: Date;
	}>(
		`SELECT id, organization_id, full_name, job_title, emails, phones, social_profiles, created_at, updated_at
		 FROM sales.crm_contacts`,
	);
	const contacts = new Map<string, Contact>();
	const orgContacts = new Map<string, string[]>();
	for (const r of rows) {
		contacts.set(r.id, {
			id: r.id,
			full_name: r.full_name,
			job_title: r.job_title,
			emails: r.emails,
			phones: r.phones,
			social_profiles: r.social_profiles,
			created_at: iso(r.created_at),
			updated_at: iso(r.updated_at),
		});
		if (r.organization_id) {
			const list = orgContacts.get(r.organization_id) ?? [];
			list.push(r.id);
			orgContacts.set(r.organization_id, list);
		}
	}
	return { contacts, orgContacts };
}

/** table/keyCol/valCol are always internal call-site literals below, never user input. */
async function queryJoinList(
	table: string,
	keyCol: string,
	valCol: string,
): Promise<Map<string, string[]>> {
	const { rows } = await pool.query<Record<string, string>>(
		`SELECT ${keyCol}, ${valCol} FROM sales.${table}`,
	);
	const m = new Map<string, string[]>();
	for (const r of rows) {
		const list = m.get(r[keyCol]) ?? [];
		list.push(r[valCol]);
		m.set(r[keyCol], list);
	}
	return m;
}

async function queryOrganizations(
	users: Map<string, User>,
	industries: Map<string, Industry>,
	orgIndustries: Map<string, string[]>,
	orgFollowers: Map<string, string[]>,
	contacts: Map<string, Contact>,
	orgContacts: Map<string, string[]>,
): Promise<Map<string, Organization>> {
	const { rows } = await pool.query<{
		id: string;
		owner_id: string | null;
		title: string;
		description: string | null;
		website: string | null;
		address: unknown;
		created_at: Date;
		updated_at: Date;
	}>(
		`SELECT id, owner_id, title, description, website, address, created_at, updated_at
		 FROM sales.crm_organizations`,
	);
	const orgs = new Map<string, Organization>();
	for (const r of rows) {
		orgs.set(r.id, {
			id: r.id,
			title: r.title,
			description: r.description,
			website: r.website,
			address: r.address,
			owner: r.owner_id ? (users.get(r.owner_id) ?? null) : null,
			industries: lookupAll(orgIndustries.get(r.id) ?? [], industries),
			followers: lookupAll(orgFollowers.get(r.id) ?? [], users),
			contacts: lookupAll(orgContacts.get(r.id) ?? [], contacts),
			created_at: iso(r.created_at),
			updated_at: iso(r.updated_at),
		});
	}
	return orgs;
}

async function queryTasks(
	users: Map<string, User>,
	taskAssignees: Map<string, string[]>,
): Promise<Map<string, Task[]>> {
	const { rows } = await pool.query<{
		id: string;
		created_by_id: string;
		completed_by_id: string | null;
		deal_id: string | null;
		title: string;
		description: string | null;
		task_type: string;
		status: string;
		due_date: Date | null;
		completed_at: Date | null;
		created_at: Date;
		updated_at: Date;
	}>(
		`SELECT id, created_by_id, completed_by_id, deal_id, title, description, task_type, status,
		        due_date, completed_at, created_at, updated_at
		 FROM sales.crm_tasks`,
	);
	const dealTasks = new Map<string, Task[]>();
	for (const r of rows) {
		if (!r.deal_id) continue;
		// created_by_id is NOT NULL with a FK to crm_users, so this is always present.
		const createdBy = users.get(r.created_by_id);
		if (!createdBy) continue;
		const task: Task = {
			id: r.id,
			title: r.title,
			description: r.description,
			task_type: r.task_type,
			status: r.status,
			due_date: isoOrNull(r.due_date),
			completed_at: isoOrNull(r.completed_at),
			created_by: createdBy,
			completed_by: r.completed_by_id ? (users.get(r.completed_by_id) ?? null) : null,
			assignees: lookupAll(taskAssignees.get(r.id) ?? [], users),
			created_at: iso(r.created_at),
			updated_at: iso(r.updated_at),
		};
		const list = dealTasks.get(r.deal_id) ?? [];
		list.push(task);
		dealTasks.set(r.deal_id, list);
	}
	return dealTasks;
}

export async function getDeals(): Promise<Deal[]> {
	const teams = await queryTeams();
	const userTeams = await queryUserTeams();
	const users = await queryUsers(teams, userTeams);
	const industries = await queryIndustries();
	const orgIndustries = await queryJoinList(
		"crm_organizations_industries",
		"organization_id",
		"industry_id",
	);
	const orgFollowers = await queryJoinList("crm_organizations_users", "organization_id", "user_id");
	const { contacts, orgContacts } = await queryContacts();
	const orgs = await queryOrganizations(
		users,
		industries,
		orgIndustries,
		orgFollowers,
		contacts,
		orgContacts,
	);
	const sources = await querySources();
	const campaigns = await queryCampaigns();
	const lossReasons = await queryLossReasons();
	const products = await queryProducts();
	const pipelines = await queryPipelines();
	const stages = await queryPipelineStages(pipelines);
	const taskAssignees = await queryJoinList("crm_tasks_users", "task_id", "user_id");
	const dealTasks = await queryTasks(users, taskAssignees);
	const dealProducts = await queryJoinList("crm_deals_products", "deal_id", "product_id");
	const dealContacts = await queryJoinList("crm_deals_contacts", "deal_id", "contact_id");

	const { rows } = await pool.query<{
		id: string;
		stage_id: string;
		owner_id: string | null;
		source_id: string | null;
		campaign_id: string | null;
		loss_reason_id: string | null;
		organization_id: string | null;
		title: string;
		amount: number | null;
		expected_close_date: string | null;
		rating: number | null;
		status: string;
		closed_at: Date | null;
		created_at: Date;
		updated_at: Date;
	}>(
		`SELECT id, stage_id, owner_id, source_id, campaign_id, loss_reason_id,
		        organization_id, title, amount, expected_close_date, rating, status,
		        closed_at, created_at, updated_at
		 FROM sales.crm_deals`,
	);

	const deals: Deal[] = [];
	for (const r of rows) {
		// stage_id is NOT NULL with a FK to crm_pipeline_stages, so this is always present.
		const stage = stages.get(r.stage_id);
		if (!stage) continue;
		deals.push({
			id: r.id,
			title: r.title,
			stage,
			owner: r.owner_id ? (users.get(r.owner_id) ?? null) : null,
			source: r.source_id ? (sources.get(r.source_id) ?? null) : null,
			campaign: r.campaign_id ? (campaigns.get(r.campaign_id) ?? null) : null,
			loss_reason: r.loss_reason_id ? (lossReasons.get(r.loss_reason_id) ?? null) : null,
			organization: r.organization_id ? (orgs.get(r.organization_id) ?? null) : null,
			amount: r.amount,
			expected_close_date: dateOnlyToIsoOrNull(r.expected_close_date),
			rating: r.rating,
			status: r.status as Deal["status"],
			closed_at: isoOrNull(r.closed_at),
			contacts: lookupAll(dealContacts.get(r.id) ?? [], contacts),
			products: lookupAll(dealProducts.get(r.id) ?? [], products),
			tasks: dealTasks.get(r.id) ?? [],
			created_at: iso(r.created_at),
			updated_at: iso(r.updated_at),
		});
	}
	return deals;
}
