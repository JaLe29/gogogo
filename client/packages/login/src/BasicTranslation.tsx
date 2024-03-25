import { animated, useSpring } from '@react-spring/web';
import { useEffect, useState } from 'react';

interface Props {
	children: any;
}

const BasicTransition: React.FC<Props> = ({ children }) => {
	const [active, setActive] = useState(false);

	const style = useSpring({
		opacity: active ? 1 : 0,
		y: active ? 0 : 24,
	});

	useEffect(() => setActive(true), []);

	return <animated.div style={style}>{children}</animated.div>;
};

export default BasicTransition;