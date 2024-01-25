// import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { ConfigProvider } from 'antd';
import cs_CZ from 'antd/locale/cs_CZ';
import 'dayjs/locale/cs';
import { Route, RouterProvider, createBrowserRouter, createRoutesFromElements } from 'react-router-dom';
// import DebugPage from './components/DebugPage';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import axios from 'axios';
import InAppLayout from './components/Layouts/InAppLayout';
import OutAppLayout from './components/Layouts/OutAppLayout';
import HomePage from './pages/HomePage';
import LoginPage from './pages/LoginPage';
import ProxyPage from './pages/ProxyPage';
import { BASE_URL } from './const/client';
import ActivityPage from './pages/ActivityPage';

// const queryClient = new QueryClient({
// 	defaultOptions: {
// 		queries: {
// 			refetchOnWindowFocus: false,
// 		},
// 	},
// });
axios.defaults.baseURL = BASE_URL;
const queryClient = new QueryClient();

const App: React.FC = () => {
	const router = createBrowserRouter(
		createRoutesFromElements(
			<>
				<Route element={<InAppLayout />} path="/app">
					<Route path="" Component={HomePage} />
					<Route path="proxy" Component={ProxyPage} />
					<Route path="activity" Component={ActivityPage} />
				</Route>
				<Route element={<OutAppLayout />} path="/">
					<Route path="" Component={LoginPage} />
				</Route>
			</>,
		),
	);

	return (
		<ConfigProvider
			theme={{
				token: {
					colorPrimary: '#2875ea',
					colorSuccess: '#52c41a',
					colorWarning: '#faad14',
					colorError: '#ff4d4f',
					colorInfo: '#2875ea',
				},
			}}
		>
			<ConfigProvider locale={cs_CZ}>
				<QueryClientProvider client={queryClient}>
					{/* <trpc.Provider client={trpcClient} queryClient={queryClient}>
					<QueryClientProvider client={queryClient}>
						<ApolloProvider client={apolloClient}> */}
					{/* <BrowserRouter> */}
					{/* <System> */}
					<RouterProvider router={router} />
					{/* </System> */}
					{/* </BrowserRouter> */}
					{/* </ApolloProvider>
					</QueryClientProvider>
				</trpc.Provider> */}
				</QueryClientProvider>
			</ConfigProvider>
		</ConfigProvider>
	);
};

export default App;
