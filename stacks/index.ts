import { App } from "@serverless-stack/resources";
import { ApiStack } from "./api";

const stage = process.env.STAGE ?? ("local" as string);

export default function main(app: App) {
	app.setDefaultFunctionProps({
		runtime: "go1.x",

		environment: {
			STAGE: stage,
			MONGO_URI: process.env.MONGO_URI ?? "mongodb://localhost:27017/go-startup",
			DATABASE: process.env.DATABASE ?? "go-startup",
			REGION: app.region
		}
	});

	new ApiStack(app, "api");
}
