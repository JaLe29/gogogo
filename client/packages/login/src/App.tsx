import { Alert, Button } from 'antd';
import { useState } from 'react';
import styled from 'styled-components';
import Center from './Center';
import { StyledCard } from './StyledCard';
import SimpleForm, { Field } from './SimpleForm';
import BasicTransition from './BasicTranslation';

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

const FIELDS: Field[] = [
	{
		key: 'email',
		label: 'E-mail',
		placeholder: 'email@gmail.com',
		required: true,
		rules: [{ required: true, type: 'email', message: 'Invalid email!' }],
	},
	{
		key: 'password',
		label: 'Password',
		placeholder: 'Password',
		required: true,
		type: 'password',
		rules: [
			{
				required: true,
				message: 'Please enter a password!',
			},
		],
	},
];

const App: React.FC = () => {
	// const setToken = useSetUserToken();
	const [loading, setLoading] = useState(false);
	const [error, setError] = useState<undefined | string>(undefined);

	// const loginUser = trpc.user.loginUser.useMutation();

	const onFinishLocal = async (values: any): Promise<void> => {
		setError(undefined);
		setLoading(true);
		fetch('/__system/login', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
			},
			body: JSON.stringify({ email: values.email.toLowerCase(), password: values.password }),
		})
			.then((r) => r.json())
			.then((r) => {
				if (r.error) {
					setError('Invalid login credentials');
					setLoading(false);
				} else {
					// setToken(r.token);
					location.reload();
				}
			})
			.catch(() => {
				setError('Invalid login credentials');
				setLoading(false);
			});

		// try {
		// const r = await loginUser.mutateAsync({ email: values.email.toLowerCase(), password: values.password });
		// console.log(r)
		// } catch {
		// setError('Nesprávné přihlašovací údaje.');
		// }

		// setLoading(false);
	};

	return (
		<BasicTransition>
			<Center hasPadding>
				<StyledCard>
					<Wrapper>
						<LogoWrapper>
							Bastard Proxy
						</LogoWrapper>
						{error && (
							<Alert message={error} type="error" style={{ textAlign: 'center', margin: '0.5em' }} />
						)}
						<SimpleForm
							fields={FIELDS}
							onFinish={onFinishLocal}
							submitButtonTitle="Přihlásit se"
							loading={loading}
						/>
					</Wrapper>
				</StyledCard>
			</Center>
		</BasicTransition>
	);
};

export default App;
