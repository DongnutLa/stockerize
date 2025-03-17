import { createRoot } from 'react-dom/client'
import App from './App.tsx'
import { ConfigProvider, theme } from 'antd'

createRoot(document.getElementById('root')!).render(
    <ConfigProvider
      theme={{
        algorithm: theme.darkAlgorithm,
        token: {
          colorBgBase: "#000000",
          colorPrimary: "#fa8c16",
          colorInfo: "#fa8c16",
          colorWarning: "#fadb14",
          colorTextBase: "#ffffff",
          fontSize: 20
        }
      }}
  >
    <App />
  </ConfigProvider>
)
