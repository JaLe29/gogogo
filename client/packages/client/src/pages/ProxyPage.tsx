/* eslint-disable jsx-a11y/control-has-associated-label */
import { Button, Modal, Popconfirm, Space, Table, Tooltip } from 'antd';
import { useState } from 'react';
import { Link } from 'react-router-dom';
import {
	CloudOutlined,
	CloudServerOutlined,
	DeleteOutlined,
	EyeOutlined,
	LineChartOutlined,
	LockOutlined,
	PauseCircleOutlined,
	PlayCircleOutlined,
	UnlockOutlined,
	VerifiedOutlined,
} from '@ant-design/icons';
import Page from './Page';
import SimpleForm from '../components/SimpleForm';
import { deleteApiProxy, patchApiProxy, postApiProxy, useGetApiProxy } from '../proxy/proxy';
import { PatchProxy, Proxy } from '../model/index';

const ProxyPage: React.FC = () => {
	const [isModalOpen, setIsModalOpen] = useState(false);

	const showModal = () => {
		setIsModalOpen(true);
	};

	const handleCancel = () => {
		setIsModalOpen(false);
	};

	const { data, isLoading, refetch } = useGetApiProxy();

	const onDelete = async (id: string) => {
		await deleteApiProxy({ id });
		refetch();
	};

	const onCreate = async (values: { source: string; target: string }) => {
		await postApiProxy({ source: values.source, target: values.target });

		refetch();
		handleCancel();
	};

	const onPatchProxy = async (id: string, payload: PatchProxy) => {
		await patchApiProxy(payload, { id });
		refetch();
	};

	const columns = [
		{
			title: 'Id',
			dataIndex: 'id',
			key: 'id',
		},
		{
			title: 'Source',
			dataIndex: 'source',
			key: 'source',
		},
		{
			title: 'Target',
			dataIndex: 'target',
			key: 'target',
		},
		{
			title: 'Actions',
			key: 'actions',
			render: (text: any, record: Proxy) => (
				<Space size="middle">
					<Tooltip title="Open">
						<a href={`https://${record.source}`} target="_blank" rel="noreferrer">
							<Button type="primary" icon={<EyeOutlined />} />
						</a>
					</Tooltip>
					<Tooltip title="Activity">
						<Link to={`/activity/${record.id}`}>
							<Button type="primary" icon={<LineChartOutlined />} />
						</Link>
					</Tooltip>
					<Tooltip title="Block">
						<Link to={`/block/${record.id}`}>
							<Button type="primary" icon={<LockOutlined />} />
						</Link>
					</Tooltip>
					<Tooltip title="Allow">
						<Link to={`/allow/${record.id}`}>
							<Button type="primary" icon={<UnlockOutlined />} />
						</Link>
					</Tooltip>
					<Tooltip title="Guard">
						<Link to={`/guard/${record.id}`}>
							<Button type="primary" icon={<VerifiedOutlined />} />
						</Link>
					</Tooltip>
					<Tooltip title="Pause">
						<Button
							onClick={() => onPatchProxy(record.id, { cache: record.cache, disable: !record.disable })}
							type={record.disable ? 'primary' : 'dashed'}
							icon={record.disable ? <PlayCircleOutlined /> : <PauseCircleOutlined />}
						/>
					</Tooltip>
					<Tooltip title="Cache">
						<Button
							onClick={() => onPatchProxy(record.id, { cache: !record.cache, disable: record.disable })}
							type={record.cache ? 'primary' : 'dashed'}
							icon={record.cache ? <CloudServerOutlined /> : <CloudOutlined />}
						/>
					</Tooltip>
					<Tooltip title="Delete">
						<Popconfirm
							placement="topLeft"
							title="Are you sure?"
							okText="Yes"
							cancelText="No"
							onConfirm={() => onDelete(record.id)}
						>
							<Button type="primary" danger icon={<DeleteOutlined />} />
						</Popconfirm>
					</Tooltip>
				</Space>
			),
		},
	];

	return (
		<Page>
			<Button type="primary" onClick={showModal}>
				New Proxy
			</Button>
			<br />
			<br />
			<Modal footer={null} title="Basic Modal" open={isModalOpen} onCancel={handleCancel}>
				<SimpleForm
					submitButtonTitle="Create"
					fields={[
						{ key: 'source', type: 'string', label: 'Source', rules: [{ required: true }] },
						{ key: 'target', type: 'string', label: 'Target', rules: [{ required: true }] },
					]}
					onFinish={onCreate}
				/>
			</Modal>
			<Table
				dataSource={data?.data}
				columns={columns}
				loading={isLoading}
				rowClassName={(record: Proxy) => (record.disable ? 'disabled-row' : '')}
			/>
		</Page>
	);
};

export default ProxyPage;
