import { afterAll, beforeEach, describe, expect, it } from "vitest";
import { buildApp } from "./app.js";
import { pool } from "./db.js";

const NOW = new Date("2024-01-01T00:00:00.000Z");

async function truncateAll() {
	await pool.query(`
		TRUNCATE TABLE
			sales.crm_campaigns, sales.crm_loss_reasons, sales.crm_pipelines, sales.crm_products,
			sales.crm_industries, sales.crm_sources, sales.crm_teams, sales.crm_users,
			sales.crm_pipeline_stages, sales.crm_organizations, sales.crm_contacts,
			sales.crm_deals, sales.crm_tasks
		CASCADE
	`);
}

afterAll(async () => {
	await pool.end();
});

describe("GET /deals", () => {
	beforeEach(truncateAll);

	it("returns the fully nested deals graph", async () => {
		await pool.query(
			"INSERT INTO sales.crm_teams (id, title, created_at, updated_at) VALUES ($1, $2, $3, $3)",
			["team-1", "Sales", NOW],
		);
		await pool.query(
			`INSERT INTO sales.crm_users (id, full_name, email, phone, created_at, updated_at)
			 VALUES ($1, $2, $3, $4, $5, $5)`,
			["user-1", "Ada Lovelace", "ada@example.com", null, NOW],
		);
		await pool.query("INSERT INTO sales.crm_teams_users (team_id, user_id) VALUES ($1, $2)", [
			"team-1",
			"user-1",
		]);
		await pool.query(
			`INSERT INTO sales.crm_pipelines (id, title, display_order, created_at, updated_at)
			 VALUES ($1, $2, $3, $4, $4)`,
			["pipeline-1", "Default", 1, NOW],
		);
		await pool.query(
			`INSERT INTO sales.crm_pipeline_stages
			 (id, pipeline_id, title, description, objective, display_order, created_at, updated_at)
			 VALUES ($1, $2, $3, $4, $5, $6, $7, $7)`,
			["stage-1", "pipeline-1", "Qualified", null, null, 1, NOW],
		);
		await pool.query(
			`INSERT INTO sales.crm_products (id, title, description, price, created_at, updated_at)
			 VALUES ($1, $2, $3, $4, $5, $5)`,
			["product-1", "Widget", null, "199.90", NOW],
		);
		await pool.query(
			`INSERT INTO sales.crm_deals
			 (id, stage_id, owner_id, title, expected_close_date, rating, status, amount, created_at, updated_at)
			 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $9)`,
			["deal-1", "stage-1", "user-1", "Acme contract", "2024-06-15", 5, "ongoing", "199.90", NOW],
		);
		await pool.query("INSERT INTO sales.crm_deals_products (deal_id, product_id) VALUES ($1, $2)", [
			"deal-1",
			"product-1",
		]);
		await pool.query(
			`INSERT INTO sales.crm_tasks
			 (id, created_by_id, deal_id, title, task_type, status, created_at, updated_at)
			 VALUES ($1, $2, $3, $4, $5, $6, $7, $7)`,
			["task-1", "user-1", "deal-1", "Follow up", "call", "pending", NOW],
		);
		await pool.query("INSERT INTO sales.crm_tasks_users (task_id, user_id) VALUES ($1, $2)", [
			"task-1",
			"user-1",
		]);

		const app = await buildApp();
		const res = await app.inject({ method: "GET", url: "/deals" });

		expect(res.statusCode).toBe(200);
		const deals = res.json();
		expect(deals).toHaveLength(1);

		const [deal] = deals;
		expect(deal).toMatchObject({
			id: "deal-1",
			title: "Acme contract",
			amount: 199.9,
			expected_close_date: "2024-06-15T00:00:00.000Z",
			rating: 5,
			status: "ongoing",
			stage: {
				id: "stage-1",
				title: "Qualified",
				pipeline: { id: "pipeline-1", title: "Default" },
			},
			owner: {
				id: "user-1",
				full_name: "Ada Lovelace",
				team: { id: "team-1", title: "Sales" },
			},
		});
		expect(deal.contacts).toEqual([]);
		expect(deal.products).toMatchObject([{ id: "product-1", price: 199.9 }]);
		expect(deal.tasks).toMatchObject([{ id: "task-1", assignees: [{ id: "user-1" }] }]);

		await app.close();
	});
});
