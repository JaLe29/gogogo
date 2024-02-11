import { Button, Space, Table, Tooltip } from 'antd';
import { Link, useParams } from 'react-router-dom';
import { EyeOutlined } from '@ant-design/icons';
import Page from './Page';
import { useGetApiActivityProxyIdAggregateIp } from '../activity/activity';
import { ActivityIpAggregate } from '../model';

const ActivityPage: React.FC = () => {
	const { proxyId } = useParams();

	const { data, isLoading } = useGetApiActivityProxyIdAggregateIp(proxyId ?? 'error');

	const columns = [
		{
			title: 'Ip',
			dataIndex: 'ip',
			key: 'ip',
		},
		{
			title: 'Sum',
			dataIndex: 'sum',
			key: 'sum',
		},
		{
			title: 'Actions',
			key: 'actions',
			render: (text: any, record: ActivityIpAggregate) => (
				<Space size="middle">
					<Tooltip title="Detail">
						<Link to={`/activity/${proxyId}/${record.ip}`}>
							<Button type="primary" icon={<EyeOutlined />} />
						</Link>
					</Tooltip>
				</Space>
			),
		},
	];

	return (
		<Page>
			<Table dataSource={data?.data} columns={columns} loading={isLoading} />
		</Page>
	);
};

export default ActivityPage;
