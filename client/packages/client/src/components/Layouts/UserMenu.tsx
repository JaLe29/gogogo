/* eslint-disable jsx-a11y/anchor-is-valid */
import React from 'react';
import { DownOutlined } from '@ant-design/icons';
import type { MenuProps } from 'antd';
import { Dropdown, Space } from 'antd';
import { Link } from 'react-router-dom'; 

const items: MenuProps['items'] = [
	{
		key: '1',
		label: (
			<Link
				to="#"
				onClick={() => {
					// localStorage.removeItem(TOKEN_LOCAL_STORAGE_KEY);
					// window.location.href = ROUTES.OUT_APP.LOGIN;
				}}
			>
				Odhlásit se
			</Link>
		),
	},
];

const UserMenu: React.FC = () => {
	// const user = useSystemStore(state => state.user);

	return (
		<div>
			<Dropdown menu={{ items }}>
				<span onClick={e => e.preventDefault()}>
					<Space>
						{"ERROR"}
						<DownOutlined />
					</Space>
				</span>
			</Dropdown>
		</div>
	);
};

export default UserMenu;