/* eslint-disable no-console */
// import styled from 'styled-components';
import styled from 'styled-components';

const PageStyled = styled.div<{ $width?: string }>`
	width: ${({ $width }) => $width || 'initial'};
	padding-top: 2em;
	padding: 2em;
	background: #f5f5f5;
`;

const Loader = styled.div`
	display: flex;
	align-items: center;
	justify-content: center;
	flex-direction: column;
	padding-top: 5em;

	img {
		padding-bottom: 2em;
	}

	h3 {
		padding-bottom: 2em;
	}
`;
interface Props {
	loading?: boolean;
	children: AnyReactElement;
	width?: string;
}

const Page: React.FC<Props> = ({ children, loading, width }) => {
	if (loading) {
		return <Loader>LOADING</Loader>;
	}

	return <PageStyled $width={width}>{children}</PageStyled>;
};

export default Page;
