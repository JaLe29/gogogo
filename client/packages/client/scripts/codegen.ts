const config: any = {
	overwrite: true,
	schema: `http://localhost:2001/graphql`,
	generates: {
		'./src/types/gql.ts': {
			plugins: ['typescript', 'typescript-operations'],
			config: {
				maybeValue: 'T',
			},
		},
	},
	hooks: {
		// afterAllFileWrite: 'yarn lint:fix',
	},
};

export default config;
