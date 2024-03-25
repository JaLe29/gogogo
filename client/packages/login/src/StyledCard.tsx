import { Card } from 'antd';
import styled from 'styled-components';

export const StyledCard = styled(Card)<{ $width?: string; $margin?: string; $padding?: string }>`
	width: ${({ $width }) => $width || 'initial'};
	box-shadow: 0px 4px 20px rgba(0, 0, 0, 0.07);
	margin: ${({ $margin }) => $margin};

	.ant-card-body {
		padding: ${({ $padding }) => $padding};
	}
`;