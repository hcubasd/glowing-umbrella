import cors from "@fastify/cors";
import swagger from "@fastify/swagger";
import swaggerUi from "@fastify/swagger-ui";
import Fastify from "fastify";
import {
	jsonSchemaTransform,
	serializerCompiler,
	validatorCompiler,
	type ZodTypeProvider,
} from "fastify-type-provider-zod";
import { z } from "zod";
import { getDeals } from "./repository.js";
import { DealSchema } from "./schemas.js";

export async function buildApp() {
	const app = Fastify({ logger: true }).withTypeProvider<ZodTypeProvider>();

	app.setValidatorCompiler(validatorCompiler);
	app.setSerializerCompiler(serializerCompiler);

	await app.register(cors, {
		origin: "https://dashboard.mlclogistica.app",
		credentials: true,
		methods: ["GET", "OPTIONS"],
		allowedHeaders: ["Authorization", "Content-Type"],
	});

	await app.register(swagger, {
		openapi: {
			info: {
				title: "Glowing Umbrella API",
				version: "2.0.0",
				description: "CRM deals read API",
			},
		},
		transform: jsonSchemaTransform,
	});
	await app.register(swaggerUi, { routePrefix: "/documentation" });

	app.get(
		"/deals",
		{
			schema: {
				summary: "List deals",
				description: "Returns all deals with fully nested graph",
				response: { 200: z.array(DealSchema) },
			},
		},
		async () => getDeals(),
	);

	return app;
}
