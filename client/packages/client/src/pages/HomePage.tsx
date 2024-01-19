import { Button, Col, Row, Space, Typography } from 'antd';
import styled from 'styled-components';
import { Link } from 'react-router-dom';
import { useQuery } from '@apollo/client'; 
import Page from './Page';


const HomePage: React.FC = () => {
	return (
		// <BasicTransition>
		<Page>
			home page
		</Page>
		// </BasicTransition >
	);
};

export default HomePage;
