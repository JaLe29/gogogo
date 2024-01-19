import { UserOutlined } from '@ant-design/icons';
import { Button, Layout } from 'antd';
import React from 'react';
import { Link } from 'react-router-dom';
import styled from 'styled-components'; 
import UserMenu from './UserMenu';

const { Header: AntdHeader } = Layout;

const StyledAntdHeader = styled(AntdHeader)`
	display: flex;
	flex-direction: row;
	justify-content: space-between;
	background: #ffffff;
	position: relative;
	box-shadow: 0px 4px 20px rgba(0, 0, 0, 0.07);

 
`;

const StyledLogoWrapper = styled.div`
	display: flex;
	flex-direction: row;
	align-items: center;
	justify-content: center;
 
`;

const StyledLogo = styled.div`
	padding-right: 4em;
	display: flex;
	align-items: center;
	justify-content: center;

	a {
		img {
			display: flex;
			align-items: center;
			justify-content: center;
		}
	}
`;

const StyledSearchWrapper = styled.div`
	display: flex;
	flex-direction: row;
	align-items: center;
	gap: 15px;
 
`;

interface Props {
	isLoginFlow?: boolean;
}

const Header: React.FC<Props> = ({ isLoginFlow }) => (
	<StyledAntdHeader>
		<StyledLogoWrapper>
			<StyledLogo>
				<Link to="/app">
					Logo
				</Link>
			</StyledLogo> 
		</StyledLogoWrapper>

		<UserMenu />
	</StyledAntdHeader>
);

export default Header;