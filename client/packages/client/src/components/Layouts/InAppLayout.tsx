import React, { useState } from 'react';
import { Button, Layout } from 'antd';
import { Outlet, useNavigate } from 'react-router-dom';
import { RollbackOutlined } from '@ant-design/icons';
import styled from 'styled-components';
import Header from './Header';
import Sidebar from './Sidebar';

const shadow = 'rgba(61, 107, 177, 0.5)';
const shadow2 = 'rgba(61, 107, 177, 0)';

const StyledSwipeButton = styled.div<{ $opacity: number }>`
	pointer-events: none;
	top: 50%;
	left: 5%;
	position: fixed;
	z-index: 10001;
	opacity: ${({ $opacity }) => `${$opacity}%`};

	transition-duration: 1s;
	transition-property: transform;
	transform: ${({ $opacity }) => `${$opacity >= 100 ? 'rotate(-80deg)' : 'none'}`};

	@keyframes spin {
		from {
			transform: rotate(0deg);
		}
		to {
			transform: rotate(360deg);
		}
	}
`;

const ShadowSwipeArea = styled.div<{ $opacity: number }>`
	pointer-events: none;
	top: calc(50% - 180px);
	left: -200px;
	position: fixed;
	z-index: 10000;

	opacity: ${({ $opacity }) => {
		const max = 20;
		const o = $opacity / 5;

		if (o <= max) {
			return `${o}%`;
		}
		return `${max}%`;
	}};
	height: 400px;
	width: 200px;
	background-color: rgb(61, 107, 177);
	border-radius: 50%;

	/* transition-duration: 1s; */
	transition-property: transform;
	transform: ${({ $opacity }) => `translateX(${(Math.min($opacity, 100) / 100) * 50}px)`};

	@keyframes spin {
		from {
			transform: translateX(0);
		}
		to {
			transform: translateX(50px);
		}
	}

	/* animation: pulse 1s;
	box-shadow: 0px 0px 0px 0px ${shadow};
	animation-iteration-count: 100;
	@keyframes pulse {
		0% {
			transform: scale(0.85);
		}
		70% {
			transform: scale(1);
			box-shadow: 0 0 0 15px ${shadow2};
		}
		100% {
			transform: scale(1);
			box-shadow: 0 0 0 0 ${shadow2};
		}
	} */
`;

const InAppLayout: React.FC = () => {
	const navigate = useNavigate();

	const [swipePoint, setSwipePoint] = useState(0);
	const fired = React.useRef(false);

	return (
		<>
			<StyledSwipeButton $opacity={swipePoint}>
				<Button icon={<RollbackOutlined />} shape="circle" size="large" type="primary" />
			</StyledSwipeButton>
			<ShadowSwipeArea $opacity={swipePoint} />
			<Layout style={{ height: '100vh' }}>
				<Header />
				<Layout>
					<Sidebar />
					<Layout>
						<Outlet />
					</Layout>
				</Layout>
			</Layout>
		</>
	);
};

export default InAppLayout;
