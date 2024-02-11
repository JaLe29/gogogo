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
import { Line } from 'react-chartjs-2';
import Page from './Page';
import { useGetApiActivityProxyIdTimelineIp } from '../activity/activity';

ChartJS.register(CategoryScale, LinearScale, PointElement, LineElement, Title, Tooltip, Legend);

const ActivityDetailIpPage: React.FC = () => {
	const { proxyId, ip } = useParams();

	const { data, isLoading } = useGetApiActivityProxyIdTimelineIp(proxyId ?? 'error', { ip: ip ?? 'error' });

	if (isLoading) {
		return <Page>Loading...</Page>;
	}

	return (
		<Page>
			<Line
				options={{
					responsive: true,
					normalized: true,
					plugins: {
						legend: {
							position: 'bottom' as const,
						},
						title: {
							display: true,
							text: ip,
						},
					},
				}}
				data={{
					labels: data?.data.map(d => d.createdAt) ?? [],
					datasets: [
						{
							label: 'Sum',
							data: data?.data.map(d => d.sum) ?? [],
							fill: false,
							backgroundColor: 'rgb(75, 192, 192)',
							borderColor: 'rgba(75, 192, 192, 0.2)',
						},
					],
				}}
			/>
		</Page>
	);
};

export default ActivityDetailIpPage;
