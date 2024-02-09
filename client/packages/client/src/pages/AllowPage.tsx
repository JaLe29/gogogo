import { Button, Modal, Popconfirm, Space, Table } from 'antd';
import { useState } from 'react';
import { useParams } from 'react-router-dom';
import Page from './Page';
import SimpleForm from '../components/SimpleForm';
import { deleteApiAllowProxyId, postApiAllowProxyId, useGetApiAllowProxyId } from '../allow/allow';

const AllowPage: React.FC = () => {
	const { proxyId } = useParams();
	const [isModalOpen, setIsModalOpen] = useState(false);

	const showModal = () => {
		setIsModalOpen(true);
	};

	const handleCancel = () => {
		setIsModalOpen(false);
	};

	const { data, isLoading, refetch } = useGetApiAllowProxyId(proxyId ?? 'error');

	const onCreate = async (values: { ip: string }) => {
		await postApiAllowProxyId(proxyId ?? 'error', { ip: values.ip });
		refetch();
		handleCancel();
	};

	const onDelete = async (id: string) => {
		await deleteApiAllowProxyId(proxyId ?? 'error', { id });
		refetch();
	};

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
			title: 'Ip',
			dataIndex: 'ip',
			key: 'ip',
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
				New Allow
			</Button>
			<br />
			<br />
			<Modal footer={null} title="Basic Modal" open={isModalOpen} onCancel={handleCancel}>
				<SimpleForm
					submitButtonTitle="Create"
					fields={[{ key: 'ip', type: 'string', label: 'Ip', rules: [{ required: true }] }]}
					onFinish={onCreate}
				/>
			</Modal>
			<Table dataSource={data?.data} columns={columns} loading={isLoading} />
		</Page>
	);
};

export default AllowPage;
