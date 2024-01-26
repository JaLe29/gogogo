import { ConfigProvider } from 'antd';
import cs_CZ from 'antd/locale/cs_CZ';
import 'dayjs/locale/cs';
import { Route, RouterProvider, createBrowserRouter, createRoutesFromElements } from 'react-router-dom';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import axios from 'axios';
import InAppLayout from './components/Layouts/InAppLayout';
import OutAppLayout from './components/Layouts/OutAppLayout';
import HomePage from './pages/HomePage';
import LoginPage from './pages/LoginPage';
import ProxyPage from './pages/ProxyPage';
import { BASE_URL } from './const/client';
import ActivityPage from './pages/ActivityPage';
import BlockPage from './pages/BlockPage';

axios.defaults.baseURL = BASE_URL;
const queryClient = new QueryClient();

const App: React.FC = () => {
	const router = createBrowserRouter(
		createRoutesFromElements(
			<>
				<Route element={<InAppLayout />} path="/app">
					<Route path="" Component={HomePage} />
					<Route path="proxy" Component={ProxyPage} />
					<Route path="activity/:proxyId" Component={ActivityPage} />
					<Route path="block/:proxyId" Component={BlockPage} />
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
					<RouterProvider router={router} />
				</QueryClientProvider>
			</ConfigProvider>
		</ConfigProvider>
	);
};

export default App;
