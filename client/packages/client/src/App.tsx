import { ApolloProvider } from '@apollo/client';
// import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { ConfigProvider } from 'antd';
import cs_CZ from 'antd/locale/cs_CZ';
import 'dayjs/locale/cs';
import { Route, RouterProvider, createBrowserRouter, createRoutesFromElements } from 'react-router-dom';
import { useEffect } from 'react';
// import DebugPage from './components/DebugPage';
import InAppLayout from './components/Layouts/InAppLayout';
import OutAppLayout from './components/Layouts/OutAppLayout';
// import System from './components/System';
import { useApolloClient } from './const/ApolloClient';
import { ROUTES } from './const/routes';
import HomePage from './pages/HomePage';
import LoginPage from './pages/LoginPage';
import { useSystemStore } from './stores/SystemStore';
import { useDeviceStore } from './stores/DeviceStore';
import ProxyPage from './pages/ProxyPage';

// const queryClient = new QueryClient({
// 	defaultOptions: {
// 		queries: {
// 			refetchOnWindowFocus: false,
// 		},
// 	},
// });

const App: React.FC = () => {

	const router = createBrowserRouter(
		createRoutesFromElements(
			<>
				<Route
					element={
						<InAppLayout />
					}
					path="/app"
				>
					<Route path="" Component={HomePage} />
					<Route path="proxy" Component={ProxyPage} />
					
				</Route >
				<Route
					element={
						<OutAppLayout />
					}
					path="/"
				>
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
			</ConfigProvider>
		</ConfigProvider>
	);
};

export default App;
