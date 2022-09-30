import { Stack, App, StackProps, Api } from "@serverless-stack/resources";

export class ApiStack extends Stack {
    constructor(scope: App, id: string, props?: StackProps) {
        super(scope, id, props);

        const api = new Api(this, "api", {
            routes: {
                $default: "cmd/startup/lambda.go"
            }
        });

        this.addOutputs({
            endpoint: process.env.IS_LOCAL
                ? api.url
                : "https://*******/.execute-api.*******.amazonaws.com/api/",
        });
    }
}
