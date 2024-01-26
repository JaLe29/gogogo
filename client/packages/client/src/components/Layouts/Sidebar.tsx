import { ProfileOutlined } from '@ant-design/icons';
import { Button, Layout } from 'antd';
import React from 'react';
import { isMobile } from 'react-device-detect';
import { Link, useLocation } from 'react-router-dom';
import styled, { css } from 'styled-components';
// import { ROUTES } from '../const/routes';
// import { useUserFeatureFlags } from '../hooks/useUserFeatureFlags';

const ROUTES = { APP: { PROXY: '/app/proxy', ACTIVITY: '/app/activity' } };

const { Sider: AntdSidebar } = Layout;

const StyledAntdSidebar = styled(AntdSidebar)`
	background-color: #ffffff !important;
	box-shadow: 0px 4px 20px rgba(0, 0, 0, 0.07);
	clip-path: inset(0px -20px 0px 0px);
	border-top: 1px solid #f5f5f5;
`;

interface ExtendedButtonProps {
	collapsed?: boolean;
}

const StyledButton = styled(Button)<ExtendedButtonProps>`
	width: 100%;
	margin-top: 1em;
	${props =>
		!props.collapsed &&
		css`
			text-align: left;
		`}
`;

const SideBar: React.FC = () => {
	const { pathname } = useLocation();

	return (
		<>
			<StyledAntdSidebar width={176} collapsible={isMobile}>
				<div style={{ display: 'relative', padding: '20px' }}>
					<div style={{ display: 'flex', flexDirection: 'column', width: '100%' }}>
						<Link to={ROUTES.APP.PROXY}>
							<StyledButton
								type={pathname.startsWith(ROUTES.APP.PROXY) ? 'link' : 'text'}
								icon={<ProfileOutlined />}
							>
								Proxy
							</StyledButton>
						</Link>
						{/* <Link to={ROUTES.APP.ACTIVITY}>
							<StyledButton
								type={pathname.startsWith(ROUTES.APP.ACTIVITY) ? 'link' : 'text'}
								icon={<IdcardOutlined />}
							>
								Activity
							</StyledButton>
						</Link> */}
					</div>
				</div>
			</StyledAntdSidebar>
		</>
	);
};

export default SideBar;
