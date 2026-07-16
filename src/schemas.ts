import { z } from "zod";

export const TeamSchema = z.object({
	id: z.string(),
	title: z.string(),
	created_at: z.string(),
	updated_at: z.string(),
});

export const UserSchema = z.object({
	id: z.string(),
	full_name: z.string(),
	email: z.string().nullable(),
	phone: z.string().nullable(),
	team: TeamSchema.nullable(),
	created_at: z.string(),
	updated_at: z.string(),
});

export const IndustrySchema = z.object({
	id: z.string(),
	title: z.string(),
	created_at: z.string(),
	updated_at: z.string(),
});

export const SourceSchema = z.object({
	id: z.string(),
	title: z.string(),
	description: z.string().nullable(),
	created_at: z.string(),
	updated_at: z.string(),
});

export const CampaignSchema = z.object({
	id: z.string(),
	title: z.string(),
	description: z.string().nullable(),
	created_at: z.string(),
	updated_at: z.string(),
});

export const LossReasonSchema = z.object({
	id: z.string(),
	reason: z.string(),
	created_at: z.string(),
	updated_at: z.string(),
});

export const ProductSchema = z.object({
	id: z.string(),
	title: z.string(),
	description: z.string().nullable(),
	price: z.number(),
	created_at: z.string(),
	updated_at: z.string(),
});

export const PipelineSchema = z.object({
	id: z.string(),
	title: z.string(),
	display_order: z.number().int(),
	created_at: z.string(),
	updated_at: z.string(),
});

export const PipelineStageSchema = z.object({
	id: z.string(),
	title: z.string(),
	description: z.string().nullable(),
	objective: z.string().nullable(),
	display_order: z.number().int(),
	pipeline: PipelineSchema,
	created_at: z.string(),
	updated_at: z.string(),
});

export const ContactSchema = z.object({
	id: z.string(),
	full_name: z.string(),
	job_title: z.string().nullable(),
	emails: z.unknown(),
	phones: z.unknown(),
	social_profiles: z.unknown(),
	created_at: z.string(),
	updated_at: z.string(),
});

export const OrganizationSchema = z.object({
	id: z.string(),
	title: z.string(),
	description: z.string().nullable(),
	website: z.string().nullable(),
	address: z.unknown(),
	owner: UserSchema.nullable(),
	industries: z.array(IndustrySchema),
	followers: z.array(UserSchema),
	contacts: z.array(ContactSchema),
	created_at: z.string(),
	updated_at: z.string(),
});

export const TaskSchema = z.object({
	id: z.string(),
	title: z.string(),
	description: z.string().nullable(),
	task_type: z.string(),
	status: z.string(),
	due_date: z.string().nullable(),
	completed_at: z.string().nullable(),
	created_by: UserSchema,
	completed_by: UserSchema.nullable(),
	assignees: z.array(UserSchema),
	created_at: z.string(),
	updated_at: z.string(),
});

export const DealSchema = z.object({
	id: z.string(),
	title: z.string(),
	stage: PipelineStageSchema,
	owner: UserSchema.nullable(),
	source: SourceSchema.nullable(),
	campaign: CampaignSchema.nullable(),
	loss_reason: LossReasonSchema.nullable(),
	organization: OrganizationSchema.nullable(),
	amount: z.number().nullable(),
	expected_close_date: z.string().nullable(),
	rating: z.number().int().nullable(),
	status: z.enum(["won", "lost", "ongoing"]),
	closed_at: z.string().nullable(),
	contacts: z.array(ContactSchema),
	products: z.array(ProductSchema),
	tasks: z.array(TaskSchema),
	created_at: z.string(),
	updated_at: z.string(),
});

export type Team = z.infer<typeof TeamSchema>;
export type User = z.infer<typeof UserSchema>;
export type Industry = z.infer<typeof IndustrySchema>;
export type Source = z.infer<typeof SourceSchema>;
export type Campaign = z.infer<typeof CampaignSchema>;
export type LossReason = z.infer<typeof LossReasonSchema>;
export type Product = z.infer<typeof ProductSchema>;
export type Pipeline = z.infer<typeof PipelineSchema>;
export type PipelineStage = z.infer<typeof PipelineStageSchema>;
export type Contact = z.infer<typeof ContactSchema>;
export type Organization = z.infer<typeof OrganizationSchema>;
export type Task = z.infer<typeof TaskSchema>;
export type Deal = z.infer<typeof DealSchema>;
