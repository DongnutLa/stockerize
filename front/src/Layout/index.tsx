import Layout, { Content, Footer, Header } from 'antd/es/layout/layout'
import {Space, Button } from 'antd'
import Sider from 'antd/es/layout/Sider'
import MainMenu from '../Components/Menu';

function MainLayout() {
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
        backgroundColor: '#000',
      };
    
      const headerStyle: React.CSSProperties = {
        textAlign: 'center',
        height: 64,
        paddingInline: 48,
        lineHeight: '64px',
        backgroundColor: '#000',
      };
    
      return (
        <Layout style={layoutStyle}>
          <Sider width="25%" style={siderStyle}>
            <div style={{ height: "20%"}}></div>
            <MainMenu />
          </Sider>
          <Layout>
            <Header style={headerStyle}>Header</Header>
            <Content style={contentStyle}>
            <Space>
          <Button type="primary">Primary</Button>
          <Button>Default</Button>
    
          asdfdsafsdaf
        </Space>
            </Content>
            <Footer style={footerStyle}>Footer</Footer>
          </Layout>
        </Layout>
      )
}

export default MainLayout;