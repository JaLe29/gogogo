/* eslint-disable consistent-return */
/* eslint-disable import/no-anonymous-default-export */
/* eslint-disable import/no-extraneous-dependencies */
import react from '@vitejs/plugin-react';
import checker from 'vite-plugin-checker';
import { resolve } from 'path';
import { viteSingleFile } from "vite-plugin-singlefile"

const SRC_DIR = resolve(__dirname, 'src');

export default (): any => ({
	server: {
		port: 3000,
	},
	plugins: [react(), checker({ typescript: true }), viteSingleFile()],
	resolve: {
		alias: {
			components: resolve(SRC_DIR, 'components'),
			const: resolve(SRC_DIR, 'const'),
			context: resolve(SRC_DIR, 'context'),
			contexts: resolve(SRC_DIR, 'contexts'),
			core: resolve(SRC_DIR, 'core'),
			hooks: resolve(SRC_DIR, 'hooks'),
			interfaces: resolve(SRC_DIR, 'interfaces'),
			layouts: resolve(SRC_DIR, 'layouts'),
			theme: resolve(SRC_DIR, 'theme'),
			types: resolve(SRC_DIR, 'types/types'),
			utils: resolve(SRC_DIR, 'utils'),
			pages: resolve(SRC_DIR, 'pages'),
		},
	},
	build: {
		rollupOptions: {
			external: ['@trpc/server'],
			// output: {
			// 	manualChunks(id: string) {
			// 		if (id.includes('/antd/')) {
			// 			return '@antd';
			// 		}

			// 		if (id.includes('/lodash/')) {
			// 			return '@lodash';
			// 		}

			// 		if (id.includes('react-router-dom') || id.includes('react-router')) {
			// 			return '@react-router';
			// 		}

			// 		if (id.includes('node_modules/rc-')) {
			// 			return '@rc';
			// 		}
			// 	},
			// },
		},
	},
});
