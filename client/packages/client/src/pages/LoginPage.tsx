import { Button, Card } from 'antd';
import { Link } from 'react-router-dom';
import styled from 'styled-components';
import { ROUTES } from '../const/routes';

const Wrapper = styled.div`
	max-width: 500px;
`;

const LogoWrapper = styled.div`
	display: flex;
	align-items: center;
	justify-content: center;

	h3 {
		margin-bottom: 0;
	}

	img {
		margin-left: 0.5em;
	}
`;

const LinkPaddingButton = styled(Button)`
	padding: 0;
	height: 0px;
`;

const LoginPage: React.FC = () => (
	<Card>
		<Wrapper>
			<LogoWrapper>xxxxxxxx</LogoWrapper>
			{/* {error && (
							<Alert message={error} type="error" style={{ textAlign: 'center', margin: '0.5em' }} />
						)}
						<SimpleForm
							fields={FIELDS}
							onFinish={onFinishLocal}
							submitButtonTitle="Přihlásit se"
							loading={loading}
						/> */}
			<div style={{ textAlign: 'center', paddingTop: '14px' }}>
				<Link
					style={{
						fontSize: 'bold',
						fontWeight: '700',
					}}
					to={ROUTES.OUT_APP.REGISTER}
				>
					{'Nemáte účet? '}
					<LinkPaddingButton type="link">Zaregistrute se.</LinkPaddingButton>
				</Link>
			</div>
		</Wrapper>
	</Card>
);

export default LoginPage;
