import { useParams } from 'react-router-dom';
import {
	Chart as ChartJS,
	CategoryScale,
	LinearScale,
	PointElement,
	LineElement,
	Title,
	Tooltip,
	Legend,
} from 'chart.js';
import { Button, Modal, Popconfirm, Space, Table } from 'antd';
import { useState } from 'react';
import Page from './Page';
import { deleteApiGuardProxyId, postApiGuardProxyId, useGetApiGuardProxyId } from '../guard/guard';
import SimpleForm from '../components/SimpleForm';

ChartJS.register(CategoryScale, LinearScale, PointElement, LineElement, Title, Tooltip, Legend);

const GuardPage: React.FC = () => {
	const { proxyId } = useParams();
	const [isModalOpen, setIsModalOpen] = useState(false);

	const showModal = () => {
		setIsModalOpen(true);
	};

	const handleCancel = () => {
		setIsModalOpen(false);
	};

	const { data, isLoading, refetch } = useGetApiGuardProxyId(proxyId ?? 'error');

	const onCreate = async (values: { email: string; password: string }) => {
		await postApiGuardProxyId(proxyId ?? 'error', { password: values.password, email: values.email });
		refetch();
		handleCancel();
	};

	const onDelete = async (id: string) => {
		await deleteApiGuardProxyId(proxyId ?? 'error', { id });
		refetch();
	};

	if (isLoading) {
		return <Page>Loading...</Page>;
	}

	const columns = [
		{
			title: 'Id',
			dataIndex: 'id',
			key: 'id',
		},
		{
			title: 'createdAt',
			dataIndex: 'createdAt',
			key: 'createdAt',
		},
		{
			title: 'Email',
			dataIndex: 'email',
			key: 'email',
		},
		{
			title: 'Actions',
			key: 'actions',
			render: (text: any, record: any) => (
				<Space size="middle">
					<Popconfirm
						placement="topLeft"
						title="Are you sure?"
						okText="Yes"
						cancelText="No"
						onConfirm={() => onDelete(record.id)}
					>
						<Button type="primary" danger>
							Delete
						</Button>
					</Popconfirm>
				</Space>
			),
		},
	];

	return (
		<Page>
			<Button type="primary" onClick={showModal}>
				New Guard
			</Button>
			<br />
			<br />
			<Modal footer={null} title="Basic Modal" open={isModalOpen} onCancel={handleCancel}>
				<SimpleForm
					submitButtonTitle="Create"
					fields={[
						{ key: 'email', type: 'string', label: 'E-mail', rules: [{ required: true }] },
						{ key: 'password', type: 'string', label: 'Password', rules: [{ required: true }] },
					]}
					onFinish={onCreate}
				/>
			</Modal>
			<Table dataSource={data?.data} columns={columns} loading={isLoading} />
		</Page>
	);
};

export default GuardPage;
