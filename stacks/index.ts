import { App } from "@serverless-stack/resources";
import { ApiStack } from "./api";

const stage = process.env.STAGE ?? ("local" as string);

export default function main(app: App) {
	app.setDefaultFunctionProps({
		runtime: "go1.x",

		environment: {
			STAGE: stage,
			MONGO_URI: "todo_add",
			REGION: app.region
		}
	});

	new ApiStack(app, "api");
}
