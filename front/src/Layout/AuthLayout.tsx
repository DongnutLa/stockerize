import { Layout } from "antd";
import { Content, Footer } from "antd/es/layout/layout";
import { Outlet } from "react-router";

  const contentStyle: React.CSSProperties = {
    textAlign: 'center',
    height: 'calc(100vh - 100px)',
    lineHeight: '120px',
    display: 'grid',
    placeContent: 'center'
  };

  const footerStyle: React.CSSProperties = {
    textAlign: 'center',
    height: 100,
    backgroundColor: '#001529',
  };
  
  const layoutStyle = {
    borderRadius: 8,
    overflow: 'hidden',
    width: '100vw',
    height: '100vh'
  };

function AuthLayout() {
    return (
        <Layout style={layoutStyle}>
            <Content style={contentStyle}>
              <Outlet />
            </Content>
            <Footer style={footerStyle} />
        </Layout>
    )
}

export default AuthLayout;