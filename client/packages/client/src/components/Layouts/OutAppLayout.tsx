import React from 'react';
import { Layout } from 'antd';
import { Outlet } from 'react-router-dom';
import { useSystemStore } from '../../stores/SystemStore';
// import Header from '../Header';

const OutAppLayout: React.FC = () => {
	// const loginFlowVisible = useSystemStore(state => state.loginFlowVisible);
	return (
		<Layout style={{ height: '100vh' }}>
			{/* {loginFlowVisible && <Header isLoginFlow />} */}
			<Layout>
				<Layout>
					<Outlet />
				</Layout>
			</Layout>
		</Layout>
	);
};

export default OutAppLayout;
