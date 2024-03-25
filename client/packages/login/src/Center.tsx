import styled from 'styled-components';

const StyledDiv = styled.div<{ $hasPadding?: boolean }>`
	display: flex;
	align-items: ${props => (props.$hasPadding ? undefined : 'center')};
	justify-content: center;
	height: inherit;
	padding-top: ${props => (props.$hasPadding ? '6em' : '0')};
`;

interface Props {
	children: any;
	hasPadding?: boolean;
}

const Center: React.FC<Props> = ({ children, hasPadding }) => (
	<StyledDiv $hasPadding={hasPadding}>{children}</StyledDiv>
);

export default Center;