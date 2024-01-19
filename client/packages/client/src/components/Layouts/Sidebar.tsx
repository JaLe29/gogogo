import {
    BugOutlined,
    DownOutlined,
    FileTextOutlined,
    IdcardOutlined,
    MenuOutlined,
    ProfileOutlined,
    QuestionCircleOutlined,
    SettingOutlined,
    UpOutlined,
    UserOutlined,
} from '@ant-design/icons';
import { Button, Layout } from 'antd';
import React, { useState } from 'react';
import { useCollapse } from 'react-collapsed';
import { isMobile } from 'react-device-detect';
import { Link, useLocation } from 'react-router-dom';
import styled, { css } from 'styled-components';
// import { ROUTES } from '../const/routes';
// import { useUserFeatureFlags } from '../hooks/useUserFeatureFlags';

const ROUTES = { APP: {}, SETTINGS: {} } as any;

const { Sider: AntdSidebar } = Layout;

const StyledSettingsSection = styled.div`
	background-color: #f9f9f9;
	margin-left: 0.5em;

	button {
		padding-left: 1em;
		padding-bottom: 1em;
		max-width: 100px;
	}
`;

const StyledAntdSidebar = styled(AntdSidebar)`
	background-color: #ffffff !important;
	box-shadow: 0px 4px 20px rgba(0, 0, 0, 0.07);
	clip-path: inset(0px -20px 0px 0px);
	border-top: 1px solid #f5f5f5;
`;

interface ExtendedButtonProps {
    collapsed?: boolean;
}

const StyledButton = styled(Button) <ExtendedButtonProps>`
	width: 100%;
	margin-top: 1em;
	${props =>
        !props.collapsed &&
        css`
			text-align: left;
		`}
`;

const SideBar: React.FC = () => {
    const { pathname } = useLocation();
    const { getCollapseProps, getToggleProps, isExpanded } = useCollapse({
        defaultExpanded: pathname.startsWith(ROUTES.APP.SETTINGS),
    });
    // const featureFlags = useUserFeatureFlags();
    // const [collapsed, setCollapsed] = useState(isMobile);

    // useEffect(() => {
    // 	isExpanded && setCollapsed(false);
    // }, [isExpanded]);

    // useEffect(() => {
    // 	collapsed && setExpanded(false);
    // }, [collapsed]);

    const [displayedSidebar, setDisplayedSidebar] = useState(!isMobile);

    return (
        <>
            {isMobile && (
                <Button
                    type={pathname.startsWith(ROUTES.APP.EXAMS) ? 'link' : 'text'}
                    ghost
                    icon={<MenuOutlined />}
                    onClick={() => setDisplayedSidebar(!displayedSidebar)}
                    style={{ paddingTop: 0, paddingLeft: 0, position: 'absolute', zIndex: 50 }}
                />
            )}
            {displayedSidebar && (
                <StyledAntdSidebar width={176} collapsible={isMobile}>
                    <div style={{ display: 'relative', padding: '20px' }}>
                        <div style={{ display: 'flex', flexDirection: 'column', width: '100%' }}>
                            <Link to={ROUTES.APP.NEW_STUDENT}>
                                <StyledButton type="primary" ghost icon={<UserOutlined />}>
                                    Nový žák
                                </StyledButton>
                            </Link>
                            <Link to={ROUTES.APP.TRAININGS}>
                                <StyledButton
                                    type={pathname.startsWith(ROUTES.APP.TRAININGS) ? 'link' : 'text'}
                                    icon={<ProfileOutlined />}
                                >
                                    Výcviky
                                </StyledButton>
                            </Link>
                            <Link to={ROUTES.APP.REPORTING}>
                                <StyledButton
                                    type={pathname.startsWith(ROUTES.APP.REPORTING) ? 'link' : 'text'}
                                    icon={<IdcardOutlined />}
                                >
                                    Hlášenky
                                </StyledButton>
                            </Link>
                            <Link to={ROUTES.APP.EXAMS}>
                                <StyledButton
                                    type={pathname.startsWith(ROUTES.APP.EXAMS) ? 'link' : 'text'}
                                    icon={<FileTextOutlined />}
                                >
                                    Zkoušky
                                </StyledButton>
                            </Link>
                        </div>
                        <div
                            style={{
                                display: 'flex',
                                flexDirection: 'column',
                                position: 'absolute',
                                bottom: 0,
                                width: '100%',
                                left: 0,
                                padding: '20px',
                            }}
                        >
                            <Link to={ROUTES.APP.SUPPORT}>
                                <StyledButton
                                    type={pathname.startsWith(ROUTES.APP.SUPPORT) ? 'link' : 'text'}
                                    icon={<QuestionCircleOutlined />}
                                >
                                    Podpora
                                </StyledButton>
                            </Link>

                            <StyledButton type="text" icon={<SettingOutlined />} {...(getToggleProps() as any)}>
                                Nastavení
                                {!isExpanded ? <DownOutlined /> : <UpOutlined />}
                            </StyledButton>
                            <StyledSettingsSection {...getCollapseProps()}>
                                <Link to={`${ROUTES.APP.SETTINGS}/info`}>
                                    <StyledButton
                                        type={pathname.startsWith(`${ROUTES.APP.SETTINGS}/info`) ? 'link' : 'text'}
                                        size="small"
                                    >
                                        Informace
                                    </StyledButton>
                                </Link>

                                <Link to={`${ROUTES.APP.SETTINGS}/adress`}>
                                    <StyledButton
                                        type={pathname.startsWith(`${ROUTES.APP.SETTINGS}/adress`) ? 'link' : 'text'}
                                        size="small"
                                    >
                                        Adresa
                                    </StyledButton>
                                </Link>
                                <Link to={`${ROUTES.APP.SETTINGS}/cars`}>
                                    <StyledButton
                                        type={pathname.startsWith(`${ROUTES.APP.SETTINGS}/cars`) ? 'link' : 'text'}
                                        size="small"
                                    >
                                        Vozidla
                                    </StyledButton>
                                </Link>
                                <Link to={`${ROUTES.APP.SETTINGS}/contact`}>
                                    <StyledButton
                                        type={pathname.startsWith(`${ROUTES.APP.SETTINGS}/contact`) ? 'link' : 'text'}
                                        size="small"
                                    >
                                        Kontakt
                                    </StyledButton>
                                </Link>
                                <Link to={`${ROUTES.APP.SETTINGS}/account`}>
                                    <StyledButton
                                        type={pathname.startsWith(`${ROUTES.APP.SETTINGS}/account`) ? 'link' : 'text'}
                                        size="small"
                                    >
                                        Účet
                                    </StyledButton>
                                </Link>
                            </StyledSettingsSection>
                        </div>
                    </div>
                </StyledAntdSidebar>
            )}
        </>
    );
};

export default SideBar;