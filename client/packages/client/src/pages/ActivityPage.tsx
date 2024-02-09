import { Table } from 'antd';
import { useParams } from 'react-router-dom';
import Page from './Page';
import { useGetApiActivityProxyId } from '../activity/activity';

const ActivityPage: React.FC = () => {
	const { proxyId } = useParams();

	const { data, isLoading } = useGetApiActivityProxyId(proxyId ?? 'error');

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
	];

	return (
		<Page>
			<Table dataSource={data?.data} columns={columns} loading={isLoading} />
		</Page>
	);
};

export default ActivityPage;
