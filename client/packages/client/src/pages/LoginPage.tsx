import { Alert, Button, Card } from 'antd';
import { useState } from 'react';
import { Link } from 'react-router-dom';
import styled from 'styled-components'; 
import { ROUTES } from '../const/routes';

const MissingPaswordWrapper = styled.div`
	font-family: 'Inter';
	font-style: normal;
	font-weight: 700;
	font-size: 14px;
	line-height: 22px;
	text-align: right;
	margin-top: 0.5em;

	button {
		padding: 0;
		color: #595959;
	}
`;

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

const LoginPage: React.FC = () => {
	// const setToken = useSetUserToken();
	// const [loading, setLoading] = useState(false);
	// const [error, setError] = useState<undefined | string>(undefined);

	// const loginUser = trpc.user.loginUser.useMutation();

	// const onFinishLocal = async (values: any): Promise<void> => {
	// 	setError(undefined);
	// 	setLoading(true);

	// 	try {
	// 		const r = await loginUser.mutateAsync({ email: values.email.toLowerCase(), password: values.password });

	// 		if (r.state === 'OK') {
	// 			setToken(r.token ?? undefined);
	// 			return;
	// 		}
	// 		if (r.state === 'NOT_ACTIVATED') {
	// 			setError('Účet není aktivován. Aktivujte ho prosím přes odkaz v e-mailu.');
	// 		} else {
	// 			setError('Nesprávné přihlašovací údaje.');
	// 		}
	// 	} catch {
	// 		setError('Nesprávné přihlašovací údaje.');
	// 	}

	// 	setLoading(false);
	// };
	return (
		<Card>
			<Wrapper>
				<LogoWrapper>
					xxxxxxxx
				</LogoWrapper>
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
};

export default LoginPage;
