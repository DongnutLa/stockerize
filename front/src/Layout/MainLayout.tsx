import Layout, { Content, Footer, Header } from 'antd/es/layout/layout'
import Sider from 'antd/es/layout/Sider'
import MainMenu from '../Components/Menu';
import { Outlet, useLocation, useNavigate, useParams } from 'react-router';
import { AUTH_USER_KEY, PRIVATE_ROUTES, ROUTES } from '../utils/constants';
import { useEffect } from 'react';
import { getFromLocalStorage } from '../utils/functions';
import { AuthUser } from '../models';

const footerStyle: React.CSSProperties = {
  textAlign: 'center',
};

const layoutStyle = {
  borderRadius: 8,
  overflow: 'hidden',
  width: '100%',
  maxWidth: '100%',
  height: '100vh',
};

const contentStyle: React.CSSProperties = {
  textAlign: 'center',
  minHeight: 120,
  lineHeight: '120px',
};

const siderStyle: React.CSSProperties = {
  textAlign: 'center',
  lineHeight: '120px',
  color: '#fff',
};

const headerStyle: React.CSSProperties = {
  textAlign: 'center',
  height: 64,
  paddingInline: 48,
  lineHeight: '64px',
};

function MainLayout() {
  const location = useLocation();
  const navigate = useNavigate();
  const authUser = JSON.parse(getFromLocalStorage<AuthUser>(AUTH_USER_KEY) as string) as AuthUser | null;

  useEffect(() => {
    PRIVATE_ROUTES.forEach(route => {
      const path = location.pathname;
      
      const isPrivatePath = path.includes(route)
      const isLoggedIn = !!authUser && !!authUser.token

      if (isPrivatePath && !isLoggedIn) {
        return navigate(ROUTES.login)
      }
    })

  }, [location, authUser])

    return (
      <Layout style={layoutStyle}>
        <Sider width="25%" style={siderStyle}>
          <div style={{ height: "20%"}}></div>
          <MainMenu />
        </Sider>
        <Layout>
          <Header style={headerStyle}></Header>
          <Content style={contentStyle}>
            <Outlet />
          </Content>
          <Footer style={footerStyle}></Footer>
        </Layout>
      </Layout>
    )
}

export default MainLayout;