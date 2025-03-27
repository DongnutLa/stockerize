import { createRoot } from 'react-dom/client'
import { ConfigProvider } from 'antd'
import App from './App.tsx'

import '@ant-design/v5-patch-for-react-19';

createRoot(document.getElementById('root')!).render(
    <ConfigProvider
      theme={{
        token: {
          colorPrimary: "#fa8c16",
          colorInfo: "#fa8c16",
          colorWarning: "#fadb14",
          fontSize: 20
        }
      }}
  >
    <App />
  </ConfigProvider>
)
