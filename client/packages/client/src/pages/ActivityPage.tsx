import { Button, Modal, Space, Table } from 'antd';
import { useState } from 'react';
import Page from './Page';
import SimpleForm from '../components/SimpleForm';
import { deleteApiProxy, postApiProxy, useGetApiProxy } from '../proxy/proxy';

const ActivityPage: React.FC = () => {
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
			render: (text: any, record: any) => (
				<Space size="middle">
					<Button type="primary" danger onClick={() => onDelete(record.id)}>
						Delete
					</Button>
				</Space>
			),
		},
	];

	return (
		// <BasicTransition>
		<Page>
			<Button type="primary" onClick={showModal}>
				Open Modal
			</Button>
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
			<Table dataSource={data?.data} columns={columns} loading={isLoading} />
		</Page>
		// </BasicTransition >
	);
};

export default ActivityPage;
